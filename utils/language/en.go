package language

var englishMsgMap = map[int]string{
	SUCCESS:           "success",
	INVALID_PARMAS:    "invalid params",
	DB_ERROR:          "database error",
	RECORD_NOT_EXISTS: "record not exists",
	SERVER_PANIC:      "system exception",
	METHOD_NOT_FOUND:  "method not found",

	//用户模块
	USER_NAME_EMPTY:       "user name must not be empty",
	USER_NAME_EXISTS:      "user name【%s】exists",
	USER_PASS_EMPTY:       "password must not be empty",
	USER_PASS_NOT_MATCH:   "password not match",
	USER_INVALID_PRIORITY: "invalid priority",
}
