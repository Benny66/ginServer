package websocket

import (
	"encoding/json"
	"fmt"
	"ginServer/config"
	"ginServer/utils/log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	SVR_KEY_BUSINESS = "business" //业务服务
	SVR_KEY_LOGIC    = "logic"    //逻辑服务
)

type conn struct {
	c   *websocket.Conn
	mtx *sync.Mutex
}

var (
	addr   string                               //websocket地址
	client = conn{mtx: new(sync.Mutex), c: nil} // websocket client

	httpReqMsgReceiver = receiver{
		cmdMtx:    new(sync.Mutex),
		curMaxCmd: 0,

		receiverMtx:       new(sync.RWMutex),
		maxReceiver:       100,
		receiver:          make(map[int]chan []byte),
		receiveMsgTimeout: 6 * time.Second,
	}

	clientArr = make(map[string]*websocket.Conn)
)

type receiver struct {
	curMaxCmd int
	cmdMtx    *sync.Mutex

	receiver    map[int]chan []byte
	receiverMtx *sync.RWMutex
	maxReceiver int

	receiveMsgTimeout time.Duration
}

func (w *receiver) pushMap(k int, c chan []byte) int {
	w.receiverMtx.Lock()
	if _, ok := w.receiver[k]; ok {
		return 51003
	}
	w.receiver[k] = c
	w.receiverMtx.Unlock()

	return 200
}

func (w *receiver) deleteMap(k int) {
	w.receiverMtx.Lock()
	if _, ok := w.receiver[k]; ok {
		if len(w.receiver[k]) > 0 {
			<-w.receiver[k]
		}
		close(w.receiver[k])
		delete(w.receiver, k)
	}
	w.receiverMtx.Unlock()
}

func (w *receiver) getCmd() (v int) {
	w.cmdMtx.Lock()
	w.curMaxCmd += 1
	v = w.curMaxCmd
	w.cmdMtx.Unlock()
	return
}

func (w *receiver) recoverCmd(v int) {
	if len(w.receiver) == 0 {
		w.cmdMtx.Lock()
		w.curMaxCmd = 0
		w.cmdMtx.Unlock()
		return
	}

	if v == w.curMaxCmd {
		w.cmdMtx.Lock()
		w.curMaxCmd -= 1
		w.cmdMtx.Unlock()
	}
}

var clientManager = map[string]string{
	"business": config.Config.BusinessAddr,
	"logic":    config.Config.LogicAddr,
}

func Start() {
	for k, v := range clientManager {
		go connServer(k, v)
	}

}

type baseMsg struct {
	Company       string      `json:"company"`
	Actioncode    string      `json:"actioncode"`
	Data          interface{} `json:"data"`
	ModID         string      `json:"modID"`
	Token         string      `json:"token"`
	Result        int         `json:"result"`
	ResultMessage string      `json:"result_message"`
	CmdSequence   int         `json:"CmdSequence"`
}

func connServer(key, Addr string) {
	var err error
	defer func() {
		client.mtx.Lock()
		if clientArr[key] != nil {
			_ = clientArr[key].Close()
		}

		clientArr[key] = nil
		client.mtx.Unlock()
		log.WSLog(Addr + " 自动重连")
		//自动重连机制
		time.Sleep(3 * time.Second)
		connServer(key, Addr)
	}()
	client.c, _, err = websocket.DefaultDialer.Dial(Addr, nil)

	if err != nil {
		return
	}
	clientArr[key] = client.c

	for {
		var res []byte
		_, res, err := clientArr[key].ReadMessage()
		if err != nil {
			log.WSLog("数据读取失败")
			return
		}
		var msg baseMsg
		err = json.Unmarshal(res, &msg)
		if err != nil {
			log.WSLog("数据解析失败")
			return
		}
		fmt.Printf("服务器推送数据：%+v\n", msg)
		httpReqMsgReceiver.ClientCodeToFunc(msg)
		//httpReqMsgReceiver.receiveMsg(msg.CmdSequence, res)
	}
}

//发送给固定客户端
func (w *receiver) receiveMsg(cmd int, msg []byte) {
	w.receiverMtx.RLock()
	defer w.receiverMtx.RUnlock()
	fmt.Println(9)
	if _, ok := w.receiver[cmd]; ok {
		go func() {
			select {
			case w.receiver[cmd] <- msg:
			case <-time.After(1 * time.Second):
			}
		}()
	}
}

/*
* description: 发送数据到各端
* author: jiangjm
* created on: 2021/4/8 11.03
* param key: business--表示发送给业务服务器，logic--表示逻辑服务器
* param Actioncode: 发送指令
* param ModID: 站点ID（3位前补零）+服务器模块类型（2位前补零），示例为1号站点服务器的数据服务
* param Token: 登录token
* param Data: 发送的具体数据
* return return_1:
 */
func SendAndReceive(key, Actioncode, ModID, Token string, Data interface{}) (data baseMsg, code int) {
	if httpReqMsgReceiver.maxReceiver < len(httpReqMsgReceiver.receiver) {
		return data, 51002
	}

	cmd := httpReqMsgReceiver.getCmd()
	ch := make(chan []byte)
	pErr := httpReqMsgReceiver.pushMap(cmd, ch)

	defer func() {
		httpReqMsgReceiver.deleteMap(cmd)
		httpReqMsgReceiver.recoverCmd(cmd)
	}()

	if pErr != 200 {
		return data, 51004
	}

	msg := baseMsg{
		Company:    "BL",
		Actioncode: Actioncode,
		Data:       Data,
		ModID:      ModID,
		Token:      Token,
	}

	msgByte, _ := json.Marshal(msg)
	client.mtx.Lock()
	fmt.Println(clientArr[key].LocalAddr())
	if clientArr[key] != nil {
		sErr := clientArr[key].WriteMessage(websocket.BinaryMessage, msgByte)
		client.mtx.Unlock()
		if sErr != nil {
			return data, 51004
		}

		select {
		case res := <-ch:
			var msg baseMsg
			err := json.Unmarshal([]byte(res), &msg)
			if err != nil {
				log.WSLog("数据解析失败")
				return data, 51004
			}
			fmt.Printf("接收发送数据：%+v\n", msg)
			return msg, msg.Result
		case <-time.After(httpReqMsgReceiver.receiveMsgTimeout):
			fmt.Printf("接收发送数据：连接超时\n")
			log.WSLog("连接超时")
			return data, 51005
		}
	}

	client.mtx.Unlock()
	return data, 51006
}
