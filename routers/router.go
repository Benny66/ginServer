package routers

import (
	"net/http"

	"github.com/Benny66/ginServer/config"
	"github.com/Benny66/ginServer/log"
	"github.com/Benny66/ginServer/middleware"
	"github.com/Benny66/ginServer/utils/format"
	"github.com/Benny66/ginServer/utils/language"

	"github.com/gin-contrib/pprof"

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
	//初始化日志
	log.Init(config.Config.Mode, config.Config.LogExpire)
	gin.SetMode(config.Config.Mode)
	if gin.IsDebugging() {
		pprof.Register(r, "/debug/pprof")
	}
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.Recover())
	r.Use(middleware.CrossMiddleware())

	r.NoRoute(routeNotFound)
	r.NoMethod(methodNotFound)
	r.StaticFS("/public", http.Dir("public"))

	routerV1(r.Group("/api"))
	return r
}

func methodNotFound(context *gin.Context) {
	format.NewResponseJson(context).Error(language.METHOD_NOT_FOUND)
}

func routeNotFound(context *gin.Context) {
	format.NewResponseJson(context).Error(language.METHOD_NOT_FOUND)
}
