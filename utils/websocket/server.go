package websocket

import (
	"encoding/json"
	"fmt"
	"ginServer/config"
	"ginServer/utils/function"
	"ginServer/utils/language"
	"ginServer/utils/log"
	"sync"

	"github.com/gorilla/websocket"
)

// websocket 服务器：定义一个 websocket 服务器和客户端连接。

//Client:单个websocket
type Client struct {
	Id      string
	Socket  *websocket.Conn
	Message chan []byte
}

// 广播发送数据信息
type BroadCastMessageData struct {
	Message     []byte
	IsBroadCast bool
	ClientIDs   []string
}

//Manager:所有websocket 信息
type Manager struct {
	//client.id => Client
	Group                map[string]*Client
	Lock                 sync.Mutex
	Register, UnRegister chan *Client
	BroadCastMessage     chan *BroadCastMessageData
	clientCount          uint //分组及客户端数量
}

// 初始化 wsManager 管理器
var WebsocketManager = Manager{
	Group:            make(map[string]*Client),
	Register:         make(chan *Client, 128),
	UnRegister:       make(chan *Client, 128),
	BroadCastMessage: make(chan *BroadCastMessageData, 128),
	clientCount:      0,
}

//启动websocket管理器
func (manager *Manager) Start() {
	log.WSLog("websocket 服务器启动")
	for {
		select {
		case client := <-manager.Register:
			//注册客户端
			manager.Lock.Lock()
			manager.Group[client.Id] = client
			manager.clientCount += 1
			log.WSLog(fmt.Sprintf("客户端注册: 客户端id为%s", client.Id))
			manager.Lock.Unlock()
		case client := <-manager.UnRegister:
			//注销客户端
			manager.Lock.Lock()
			if _, ok := manager.Group[client.Id]; ok {
				//关闭消息通道
				close(client.Message)
				//删除分组中客户
				delete(manager.Group, client.Id)
				//客户端数量减1
				manager.clientCount -= 1
				log.WSLog(fmt.Sprintf("客户端注销: 客户端id为%s", client.Id))
			}
			manager.Lock.Unlock()
		case data := <-manager.BroadCastMessage:
			//将数据广播给所有客户端
			for _, conn := range manager.Group {
				if data.IsBroadCast {
					conn.Message <- data.Message
				} else {
					if function.InSliceStr(conn.Id, data.ClientIDs) {
						conn.Message <- data.Message
					}
				}

			}

		}
	}
}

type ReadData struct {
	Company    string `json:"company"`
	Actioncode string `json:"actioncode"`
	Data       struct {
		// NodeID   uint   `json:"node_id"`
		// NodeName string `json:"node_name"`
	} `json:"data"`
	Token       string   `json:"token"`
	IsBroadCast bool     `json:"is_broadcast"`
	ClientIDs   []string `json:"client_ids"`
}

//从websocket中直接读取数据
func (c *Client) Read() {
	defer func() {
		//客户端关闭
		if err := c.Socket.Close(); err != nil {
			log.WSLog(fmt.Sprintf("client [%s] disconnect err: %s", c.Id, err))
		}
		//关闭后直接注销客户端
		WebsocketManager.UnRegisterClient(c)

		log.WSLog(fmt.Sprintf("client [%s],客户端关闭：[%s]", c.Id, websocket.CloseMessage))
	}()

	for {
		messageType, message, err := c.Socket.ReadMessage()
		//读取数据失败
		if err != nil || messageType == websocket.CloseMessage {
			log.WSLog(fmt.Sprintf("client [%s],数据读取失败或通道关闭：[%s],客户端连接状态：[%s]", c.Id, err.Error(), websocket.CloseMessage))
			break
		}
		//解析发送过来的参数
		var data ReadData
		err = json.Unmarshal(message, &data)
		if err != nil {
			log.WSLog("数据解析失败")
			return
		}

		//前端请求返回数据到指定客户端
		data.ClientIDs = append(data.ClientIDs, c.Id)
		WebsocketManager.ServerCodeToFunc(data)

	}
}

//写入数据到websocket中
func (c *Client) Write() {
	defer func() {
		//客户端关闭
		if err := c.Socket.Close(); err != nil {
			log.WSLog(fmt.Sprintf("client [%s] disconnect err: %s", c.Id, err))
			return
		}
		//关闭后直接注销客户端
		WebsocketManager.UnRegisterClient(c)
		log.WSLog(fmt.Sprintf("client [%s],客户端关闭：[%s]", c.Id, websocket.CloseMessage))
	}()
	for {
		select {
		case message, ok := <-c.Message:

			if !ok {
				//数据写入失败，关闭通道
				log.WSLog(fmt.Sprintf("client [%s],客户端连接状态：[%s]", c.Id, websocket.CloseMessage))
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				//消息通道关闭后直接注销客户端
				return
			}

			err := c.Socket.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.WSLog(fmt.Sprintf("client [%s] write message err: %s", c.Id, err))
				return
			}
		}
	}
}

//注册
func (manager *Manager) RegisterClient(client *Client) {
	manager.Register <- client
}

//注销
func (manager *Manager) UnRegisterClient(client *Client) {
	manager.UnRegister <- client
}

type ResultData struct {
	Company       string      `json:"company"`
	DeviceName    string      `json:"device_name"`
	Result        int         `json:"result"`
	ResultMessage string      `json:"result_message"`
	Version       string      `json:"version"`
	DBVersion     string      `json:"db_version"`
	Language      string      `json:"language"`
	Data          interface{} `json:"data"`
	ActionCode    string      `json:"actioncode"`
}

/*
* description: 各端推送数据处理
* author: shahao
* created on: 2021/4/8 11.03
* param data: 发送的数据
* param isBroadCast: 是否广播（true-是,false-否）
* param ClientIDs: 当isBroadCast=false,该参数才有效，表示要推送到指定的客户端
* return return_1:
 */
func (manager *Manager) Success(ActionCode string, data interface{}, isBroadCast bool, ClientIDs []string) {
	var result = ResultData{
		Company:       "BL",
		DeviceName:    config.Config.GetAppName(),
		Result:        language.SUCCESS,
		ResultMessage: language.Lang.Msg(language.SUCCESS),
		Version:       config.Config.GetAppVersion(),
		DBVersion:     config.Config.GetDatabaseVersion(),
		Language:      config.Config.GetLanguage(),
		Data:          data,
		ActionCode:    ActionCode,
	}
	msg, err := json.Marshal(result)
	if err != nil {
		log.WSLog("数据转换失败")
		return
	}
	//数据写入通道
	WebsocketManager.BroadCastMessage <- &BroadCastMessageData{Message: msg, IsBroadCast: isBroadCast, ClientIDs: ClientIDs}
}

/*
* description: 各端推送数据处理
* author: shahao
* created on: 2021/4/8 11.03
* param data: 发送的数据
* param isBroadCast: 是否广播（true-是,false-否）
* param ClientIDs: 当isBroadCast=false,该参数才有效，表示要推送到指定的客户端
* return return_1:
 */
func (manager *Manager) Error(errorCode int, ActionCode string, param []interface{}, isBroadCast bool, ClientIDs []string) {
	var result = ResultData{
		Company:       "BL",
		DeviceName:    config.Config.GetAppName(),
		Result:        language.SUCCESS,
		ResultMessage: language.Lang.Msg(errorCode, param...),
		Version:       config.Config.GetAppVersion(),
		DBVersion:     config.Config.GetDatabaseVersion(),
		Language:      config.Config.GetLanguage(),
		Data:          "",
		ActionCode:    ActionCode,
	}
	msg, err := json.Marshal(result)
	if err != nil {
		log.WSLog("数据转换失败")
		return
	}
	//数据写入通道
	WebsocketManager.BroadCastMessage <- &BroadCastMessageData{Message: msg, IsBroadCast: isBroadCast, ClientIDs: ClientIDs}
}
