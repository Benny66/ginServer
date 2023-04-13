package routers

import (
	"github.com/Benny66/ginServer/api"
	"github.com/Benny66/ginServer/middleware"
	"github.com/gin-gonic/gin"
)

// v1版本接口
func routerV1(group *gin.RouterGroup) {
	group.POST("/user/login", api.UserApi.Login)
	group.POST("/redis/test", api.RedisApi.Test)
	group.Use(middleware.JWTMiddleware())
	{
		group.GET("/user/logout", api.UserApi.Logout)
		group.GET("/user/refresh", api.UserApi.Refresh)
		group.PUT("/user/update", api.UserApi.UpdatePassword)
	}
}
