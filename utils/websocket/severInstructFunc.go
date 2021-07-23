/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-22 11:42:19
 * @LastEditors: shahao
 * @LastEditTime: 2021-04-27 13:53:18
 */
package websocket

type ServerMethod struct {
}

//设备状态
// func (m *ServerMethod) EquipmentStatus(params ReadData) {
// 	WebsocketManager.Success(params.Actioncode, 21, params.IsBroadCast, params.ClientIDs)
// }

//心跳包
func (m *ServerMethod) HeartBeat(params ReadData) {
	WebsocketManager.Success(params.Actioncode, true, params.IsBroadCast, params.ClientIDs)
}
