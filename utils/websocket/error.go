package websocket

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-22 09:50:42
 * @LastEditors: shahao
 * @LastEditTime: 2021-04-22 09:50:46
 */

type IServiceError interface {
	GetErrorCode() int
	GetParams() []interface{}
}

func NewServiceError(errorCode int, params ...interface{}) *serviceError {
	return &serviceError{
		errorCode: errorCode,
		params:    params,
	}
}

type serviceError struct {
	errorCode int
	params    []interface{}
}

func (se serviceError) GetErrorCode() int {
	return se.errorCode
}

func (se serviceError) GetParams() []interface{} {
	return se.params
}
