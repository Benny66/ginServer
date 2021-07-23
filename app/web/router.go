package web

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-07 09:20:20
 * @LastEditors: shahao
 * @LastEditTime: 2021-07-23 14:53:59
 */

import (
	"ginServer/app/web/api"
	"ginServer/utils/middleware"

	"github.com/gin-gonic/gin"
)

// @title ginServerweb端API接口文档
// @version 1.0
// @description API接口文档
// @host 127.0.0.1:8090
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /web/
var WebRouter *webRouter

func init() {
	WebRouter = NewWebRouter()
}

func NewWebRouter() *webRouter {
	return &webRouter{}
}

type webRouter struct{}

func (router *webRouter) Init(group *gin.RouterGroup) {
	router.routerV1(group.Group("/v1"))
}

// v1版本接口
func (router *webRouter) routerV1(group *gin.RouterGroup) {
	router.routerNotNeedLogin(group)
	group.Use(middleware.JWTMiddleware())
	{
		router.routerUser(group)
	}

}

func (router *webRouter) routerNotNeedLogin(group *gin.RouterGroup) {
	group.POST("/user/login", api.UserApi.Login)
}

func (router *webRouter) routerUser(group *gin.RouterGroup) {
	group.GET("/user/logout", api.UserApi.Logout)
	group.GET("/user/refresh", api.UserApi.Refresh)
	group.PUT("/user/update", api.UserApi.UpdatePassword)

}
