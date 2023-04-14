package api

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-07 09:20:20
 * @LastEditors: shahao
 * @LastEditTime: 2021-07-20 16:26:34
 */

import (
	"github.com/Benny66/ginServer/schemas"
	"github.com/Benny66/ginServer/service"
	"github.com/Benny66/ginServer/utils/format"
	"github.com/Benny66/ginServer/utils/language"

	"github.com/gin-gonic/gin"
)

var UserApi *userApi

func init() {
	UserApi = NewUserApi()
}

func NewUserApi() *userApi {
	return &userApi{}
}

type userApi struct {
}

// @Summary 登录
// @Description 登录
// @Tags 用户
// @accept json
// @Produce  json
// @Param create_data body define.UserLoginApiReq true "创建数据模型"
// @Success 200 {string} string {"company":"BL","device_name":"Audio Matrix","result":"0","result_message":"成功","version":"1.0", "db_version":"202103101750","language":"zh-cn","data":""}
// @Router /v1/user/login [post]
func (api *userApi) Login(context *gin.Context) {
	var req schemas.UserLoginApiReq

	if err := context.BindJSON(&req); err != nil {
		format.NewResponseJson(context).Error(language.INVALID_PARMAS)
		return
	}
	token, err := service.UserService.Login(&req)
	if err != nil {
		format.NewResponseJson(context).Error(err.GetErrorCode(), err.GetParams()...)
		return
	}
	var data = map[string]interface{}{
		"token": token,
	}
	//service.SystemLogService.CreateLog(context, language.SYS_LOG_LOGIN_LOGIN, req.UserName)
	format.NewResponseJson(context).Success(data)
}

// @Summary 刷新登录token
// @Description 刷新登录token
// @Tags 用户
// @Security ApiKeyAuth
// @accept x-www-form-urlencoded
// @Produce  json
// @Success 200 {string} string {"company":"BL","device_name":"Audio Matrix","result":"0","result_message":"成功","version":"1.0", "db_version":"202103101750","language":"zh-cn","data":""}
// @Router /v1/user/refresh [get]
func (api *userApi) Refresh(context *gin.Context) {
	userInfo := service.UserService.User(context)

	token, err := service.UserService.Refresh(userInfo)
	if err != nil {
		format.NewResponseJson(context).Error(err.GetErrorCode(), err.GetParams()...)
		return
	}
	var data = map[string]interface{}{
		"token": token,
	}
	//service.SystemLogService.CreateLog(context, language.SYS_LOG_LOGIN_LOGIN, req.UserName)
	format.NewResponseJson(context).Success(data)
}

// @Summary 退出登录
// @Description 退出登录
// @Tags 用户
// @Security ApiKeyAuth
// @accept x-www-form-urlencoded
// @Produce  json
// @Success 200 {string} string {"company":"BL","device_name":"Audio Matrix","result":"0","result_message":"成功","version":"1.0", "db_version":"202103101750","language":"zh-cn","data":""}
// @Router /v1/user/logout [get]
func (api *userApi) Logout(context *gin.Context) {
	format.NewResponseJson(context).Success("")
}

// @Summary 修改密码
// @Description 修改密码
// @Tags 用户
// @Security ApiKeyAuth
// @accept json
// @Produce  json
// @Param create_data body define.UserUpdatePasswordApiReq true "创建数据模型"
// @Success 200 {string} string {"company":"BL","device_name":"Audio Matrix","result":"0","result_message":"成功","version":"1.0", "db_version":"202103101750","language":"zh-cn","data":""}
// @Router /v1/user/update [put]
func (api *userApi) UpdatePassword(context *gin.Context) {
	var req schemas.UserUpdatePasswordApiReq

	if err := context.BindJSON(&req); err != nil {
		format.NewResponseJson(context).Error(language.INVALID_PARMAS)
		return
	}
	userInfo := service.UserService.User(context)
	data, err := service.UserService.UpdatePassword(&req, userInfo)
	if err != nil {
		format.NewResponseJson(context).Error(err.GetErrorCode(), err.GetParams()...)
		return
	}
	format.NewResponseJson(context).Success(data)
}
