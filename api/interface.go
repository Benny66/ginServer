package api

import "github.com/gin-gonic/gin"

type UserInterface interface {
	Login(context *gin.Context)
	Refresh(context *gin.Context)
	Logout(context *gin.Context)
	UpdatePassword(context *gin.Context)
}

var UserApi UserInterface = &userApi{}

type userApi struct{}

type RedisInterface interface {
	Test(context *gin.Context)
}

var RedisApi RedisInterface = &redisApi{}

type redisApi struct{}

type WsInterface interface {
	WsClient(context *gin.Context)
}

var WsApi WsInterface = &wsApi{}

type wsApi struct{}
