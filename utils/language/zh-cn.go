package language

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-07 09:20:20
 * @LastEditors: shahao
 * @LastEditTime: 2021-07-14 14:21:00
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
	USER_NAME_EMPTY:              "用户名不能为空",
	USER_NAME_EXISTS:             "用户名【%s】已存在",
	USER_PASS_EMPTY:              "密码不能为空",
	USER_LOGIN_ERROR:             "账号或密码错误",
	USER_PASS_NOT_MATCH:          "原密码错误",
	USER_CONFIRM_PASS_NOT_MATCH:  "输入新密码不一致",
	USER_INVALID_PRIORITY:        "优先级格式不正确",
	USER_NOT_EXISTS:              "用户不存在",
	USER_TOKEN_CREATE_ERROR:      "登录失败，token生成有误",
	SERVER_IP_ADDRESS_EMPTY:      "服务器ip地址不能为空",
	SERVER_IP_ADDRESS_ERROR:      "服务器ip地址有误",
	SERVER_UDP_START_ERROR:       "udp服务器启动失败",
	SERVER_IP_ADDRESS_PING_ERROR: "服务器ip地址不存在",

	TERMINAL_ID_EMPTY:               "终端ID不能为空",
	TERMINAL_NAME_EMPTY:             "终端名称不能为空",
	TERMINAL_NAME_OVER_COUNT:        "终端名称过长",
	TERMINAL_NAME_EXISTS:            "终端名称【%s】已存在",
	TERMINAL_IP_ADDRESS_EMPTY:       "终端ip地址不能为空",
	TERMINAL_IP_ADDRESS_ERROR:       "终端ip地址有误",
	TERMINAL_UNIQUE_NAME_EMPTY:      "终端mac地址不能为空",
	TERMINAL_UNIQUE_NAME_ERROR:      "终端mac地址有误",
	TERMINAL_UNIQUE_NAME_EXISTS:     "终端mac地址【%s】已存在",
	TERMINAL_TYPE_ERROR:             "终端类型有误",
	TERMINAL_SUBNETMASK_EMPTY:       "终端子网掩码不能为空",
	TERMINAL_SUBNETMASK_ERROR:       "终端子网掩码有误",
	TERMINAL_CHANNEL_NUM_ERROR:      "终端设置通道不能为空",
	TERMINAL_CHANNEL_ID_ERROR:       "终端通道ID有误",
	TERMINAL_CHANNEL_NAME_ERROR:     "终端通道名称有误",
	TERMINAL_CHANNEL_PRIORITY_ERROR: "终端通道优先级有误",
	TERMINAL_CHANNEL_VOLUME_ERROR:   "终端通道音量有误",
	TERMINAL_CHANNEL_NAME_EXISTS:    "终端通道名称【%s】已存在",
	TERMINAL_OFFLINE:                "设备离线编辑保存失败",

	UDP_NEW_CONNECT_ERROR:  "新建UDP链接失败",
	UDP_SEND_CONTENT_ERROR: "发送UDP协议失败",

	SCENE_JOIN_TERMINAL_NOT_EXIST: "场景关联终端不能为空",
	SCENE_JOIN_MATRIX_NOT_EXIST:   "场景关联矩阵不能为空",
	SCENE_NAME_EMPTY:              "场景名称不能为空",
	SCENE_NAME_EXISTS:             "场景名称已存在",
	SCENE_DISABLE_ERROR:           "场景关闭失败",
	SCENE_APPLY_ERROR:             "场景应用失败",

	DB_SELECT_SCENE_MATRIX_ERROR:               "查询场景矩阵列表失败",
	DB_CREATE_SCENE_ERROR:                      "创建场景失败",
	DB_CREATE_SCENE_JOIN_TERMINAL_MATRIX_ERROR: "创建场景关联矩阵失败",
	DB_UPDATE_SCENE_ERROR:                      "更新场景失败",
	DB_UPDATE_SCENE_JOIN_TERMINAL_MATRIX_ERROR: "更新场景关联矩阵失败",
	DB_DELETE_SCENE_JOIN_TERMINAL_MATRIX_ERROR: "删除场景关联矩阵失败",

	MATRIX_CREATE_EXCEL_SHEET_ERROR:          "矩阵导出关联输出通道创建表格失败",
	MATRIX_CREATE_EXCEL_SAVE_FILE_ERROR:      "矩阵导出关联输出通道保存表格失败",
	MATRIX_SEND_CHANNEL_VOLUME_ERROR:         "修改矩阵通道中的音量失败",
	MATRIX_CREATE_INPUT_CHANNEL_ERROR:        "矩阵应用创建输入设备的通道失败",
	MATRIX_CREATE_OUTPUT_CHANNEL_ERROR:       "矩阵应用创建输出设备的通道失败",
	MATRIX_DROP_INPUT_CHANNEL_ERROR:          "矩阵应用删除输入设备的通道失败",
	MATRIX_DROP_OUTPUT_CHANNEL_ERROR:         "矩阵应用删除输出设备的通道失败",
	MATRIX_GET_INPUT_TERMINAL_ERROR:          "矩阵查询输入终端信息失败",
	MATRIX_GET_OUTPUT_TERMINAL_ERROR:         "矩阵查询输出终端信息失败",
	MATRIX_UPDATE_OUTPUT_CHANNEL_TOTAL_ERROR: "矩阵更新输出设备通道总数失败",
	MATRIX_GET_INPUT_CHANNEL_ERROR:           "矩阵查询应用输入通道数据有误",
	MATRIX_GET_OUTPUT_CHANNEL_ERROR:          "矩阵查询应用输出通道数据有误",
	MATRIX_CREATE_MATRIX_ERROR:               "矩阵应用矩阵创建失败",
	MATRIX_DROP_MATRIX_ERROR:                 "矩阵应用删除临时点失败",
	MATRIX_ADD_SPOT_TEMP_ERROR:               "矩阵临时点查询失败",
	MATRIX_SPOT_TEMP_EXISTS_ERROR:            "矩阵临时点已存在",
	MATRIX_GET_ONE_OUTPUT_CHANNEL_ERROR:      "矩阵查询应用一个输出通道数据有误",
	MATRIX_PRIORITY_MIN:                      "矩阵点优先级低，不可添加",
	MATRIX_ADD_SPOT_SELECT_CHANNEL_ERROR:     "矩阵点处理，查询通道有误",
	MATRIX_ADD_TERMINAL_JOIN_MATRIX_ERROR:    "添加设备小矩阵点失败",
}
