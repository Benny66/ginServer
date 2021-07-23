package language

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-07 09:20:20
 * @LastEditors: shahao
 * @LastEditTime: 2021-07-23 17:20:54
 */

const (
	SUCCESS           = 0
	INVALID_PARMAS    = 1
	DB_ERROR          = 2
	RECORD_NOT_EXISTS = 3
	SERVER_PANIC      = 4
	METHOD_NOT_FOUND  = 5
	TOKEN_INVALID     = 6
	TOKEN_EXPIRE      = 7
	TOKEN_EMPTY       = 8
	PERMISSION_DENY   = 9

	//用户模块
	USER_NAME_EMPTY             = 1000
	USER_NAME_EXISTS            = 1001
	USER_PASS_EMPTY             = 1002
	USER_LOGIN_ERROR            = 1003
	USER_INVALID_PRIORITY       = 1004
	USER_NOT_EXISTS             = 1005
	USER_TOKEN_CREATE_ERROR     = 1006
	USER_PASS_NOT_MATCH         = 1007
	USER_CONFIRM_PASS_NOT_MATCH = 1008
)
