package language

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-07 09:20:20
 * @LastEditors: shahao
 * @LastEditTime: 2021-07-23 17:20:39
 */

var chineseMsgMap = map[int]string{
	SUCCESS:           "成功",
	INVALID_PARMAS:    "非法参数",
	DB_ERROR:          "数据库操作失败",
	RECORD_NOT_EXISTS: "记录不存在",
	SERVER_PANIC:      "系统异常",
	METHOD_NOT_FOUND:  "方法不存在",
	TOKEN_INVALID:     "非法token",
	TOKEN_EXPIRE:      "token已过期",
	TOKEN_EMPTY:       "token不能为空",
	PERMISSION_DENY:   "权限不足",

	//用户模块
	USER_NAME_EMPTY:             "用户名不能为空",
	USER_NAME_EXISTS:            "用户名【%s】已存在",
	USER_PASS_EMPTY:             "密码不能为空",
	USER_LOGIN_ERROR:            "账号或密码错误",
	USER_PASS_NOT_MATCH:         "原密码错误",
	USER_CONFIRM_PASS_NOT_MATCH: "输入新密码不一致",
	USER_INVALID_PRIORITY:       "优先级格式不正确",
	USER_NOT_EXISTS:             "用户不存在",
	USER_TOKEN_CREATE_ERROR:     "登录失败，token生成有误",
}
