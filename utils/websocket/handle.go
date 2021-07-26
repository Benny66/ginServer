package websocket

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-22 10:04:26
 * @LastEditors: shahao
 * @LastEditTime: 2021-07-26 11:42:25
 */

import (
	"reflect"
	"strings"
)

//定义函数Map类型，便于后续快捷使用
type ControllerMapsType map[string]reflect.Value

/**
 * Description: websocket服务器接收数据指令调用对应函数
 * author: 	shahao
 * create on:	2021-04-16 18:05:21
 */
func (manager *Manager) ServerCodeToFunc(data ReadData) {
	funcName := case2Camel(data.Actioncode)
	vft := manager.serverReturnFunc()
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(data)
	if vft[funcName].IsValid() {
		vft[funcName].Call(params)
	}
}

/**
 * Description: 下划线写法转为驼峰写法
 * author: 	shahao
 * param: 	name
 * create on:	2021-04-17 08:56:25
 * return: 	string
 */
func case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

/**
 * Description: 查询结构体中的方法
 * author: 	shahao
 * create on:	2021-04-17 11:47:12
 * return: 	ControllerMapsType
 */
func (manager *Manager) serverReturnFunc() ControllerMapsType {
	var m ServerMethod
	vf := reflect.ValueOf(&m)
	vft := vf.Type()
	//读取方法数量
	mNum := vf.NumMethod()
	crMap := make(ControllerMapsType, 0)

	//遍历所有的方法，并将其存入映射变量中
	for i := 0; i < mNum; i++ {
		mName := vft.Method(i).Name
		crMap[mName] = vf.Method(i)
	}
	return crMap
}

/**
 * Description: websocket客户端接收数据指令调用对应函数
 * author: 	shahao
 * create on:	2021-04-16 18:05:21
 */
func (w *receiver) ClientCodeToFunc(data baseMsg) {
	funcName := case2Camel(data.Actioncode)
	vft := w.serverReturnFunc()

	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(data)
	if vft[funcName].IsValid() {
		vft[funcName].Call(params)
	}
}

/**
 * Description: 查询结构体中的方法
 * author: 	shahao
 * create on:	2021-04-17 11:47:12
 * return: 	ControllerMapsType
 */
func (w *receiver) serverReturnFunc() ControllerMapsType {
	var m ClientMethod
	vf := reflect.ValueOf(&m)
	vft := vf.Type()
	//读取方法数量
	mNum := vf.NumMethod()
	crMap := make(ControllerMapsType, 0)

	//遍历所有的方法，并将其存入映射变量中
	for i := 0; i < mNum; i++ {
		mName := vft.Method(i).Name
		crMap[mName] = vf.Method(i)
	}
	return crMap
}
