package service

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-07 09:20:20
 * @LastEditors: shahao
 * @LastEditTime: 2021-04-07 09:55:05
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
