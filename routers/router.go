package routers

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-14 09:56:53
 * @LastEditors: shahao
 * @LastEditTime: 2021-07-26 11:43:58
 */

import (
	"ginServer/app/web"
	"ginServer/app/web/api"
	"ginServer/config"
	"ginServer/utils/format"
	"ginServer/utils/function"
	"ginServer/utils/language"
	"ginServer/utils/middleware"
	"ginServer/utils/websocket"
	"net/http"

	"github.com/gin-contrib/pprof"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

var Router *router

func init() {
	Router = NewRouter()
}

func NewRouter() *router {
	return &router{}
}

type router struct{}

func (router *router) Init() *gin.Engine {
	r := gin.New()
	gin.SetMode(config.Config.Mode)
	if gin.IsDebugging() {
		pprof.Register(r, "/debug")
	}
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.Recover())
	r.Use(middleware.CrossMiddleware())
	go websocket.WebsocketManager.Start()
	r.GET("/ws", api.WsClient)

	r.NoRoute(routeNotFound)
	r.NoMethod(methodNotFound)
	r.StaticFS("/public", http.Dir(function.GetAbsPath("public")))
	if gin.IsDebugging() {
		r.GET("/docs/web/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/public/web_doc/swagger.json")))
		r.GET("/docs/client/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/public/client_doc/swagger.json")))
	}

	web.WebRouter.Init(r.Group("/web"))
	return r
}

func methodNotFound(context *gin.Context) {
	format.NewResponseJson(context).Error(language.METHOD_NOT_FOUND)
}

func routeNotFound(context *gin.Context) {
	format.NewResponseJson(context).Error(language.METHOD_NOT_FOUND)
	//context.HTML(http.StatusOK, "index.html", nil)
}
