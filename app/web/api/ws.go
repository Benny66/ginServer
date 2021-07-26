package api

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-22 11:35:25
 * @LastEditors: shahao
 * @LastEditTime: 2021-07-26 11:17:49
 */

import (
	"fmt"
	"ginServer/utils/format"
	"ginServer/utils/log"
	ws "ginServer/utils/websocket"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

// 客户端连接
func WsClient(context *gin.Context) {
	upGrande := websocket.Upgrader{
		//设置允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		//设置请求协议
		Subprotocols: []string{context.GetHeader("Sec-WebSocket-Protocol")},
	}
	//创建连接
	conn, err := upGrande.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		log.WSLog(fmt.Sprintf("websocket connect error: %s", context.Param("channel")))
		format.NewResponseJson(context).Error(51001)
		return
	}
	//生成唯一标识client_id
	var uuid = uuid.NewV4().String()
	client := &ws.Client{
		Id:      uuid,
		Socket:  conn,
		Message: make(chan []byte, 1024),
	}
	//注册
	ws.WebsocketManager.RegisterClient(client)

	//起协程，实时接收和回复数据
	go client.Read()
	go client.Write()
}
