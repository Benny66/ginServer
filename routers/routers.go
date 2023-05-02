// routers.go

package routers

import (
	"fmt"

	"github.com/Benny66/ginServer/config"
	"github.com/Benny66/ginServer/log"
	"github.com/gin-gonic/gin"
)

var R *RouterRegister

func init() {
	r := gin.Default()
	log.Init(config.Config.Mode, config.Config.LogExpire)
	R = NewRouterRegister(r)
}

type RouterRegister struct {
	engine             *gin.Engine
	endpoints          []Endpoint
	endpointMiddleware map[string]MiddlewareSchema
}

func NewRouterRegister(engine *gin.Engine) *RouterRegister {
	return &RouterRegister{
		engine:             engine,
		endpointMiddleware: make(map[string]MiddlewareSchema),
	}
}

func (r *RouterRegister) Engine() *gin.Engine {
	return r.engine
}

func (r *RouterRegister) AddEndpointSchema(endpoint Endpoint) {
	r.endpoints = append(r.endpoints, endpoint)
}

func (r *RouterRegister) AddMiddlewareSchema(schema MiddlewareSchema) {
	middleware := r.endpointMiddleware
	if _, ok := middleware[schema.Name()]; !ok {
		middleware[schema.Name()] = schema
		r.endpointMiddleware = middleware
	}
}

func (r *RouterRegister) Register() {
	//加载中间件
	for _, middlewareSchema := range r.endpointMiddleware {
		//鉴权在下方接口添加
		if middlewareSchema.Name() == "jwt" {
			continue
		}
		r.engine.Use(middlewareSchema.Handler())
	}
	for _, endpoint := range r.endpoints {
		url := endpoint.URL()
		//组
		if endpoint.Group() != "" {
			url = fmt.Sprintf("%s/%s", endpoint.Group(), url)
		}
		//鉴权中间件
		if endpoint.Auth() != "" {
			if _, ok := r.endpointMiddleware[endpoint.Auth()]; ok {
				r.engine.Handle(endpoint.Method(), url, endpoint.Handler()).Use(r.endpointMiddleware[endpoint.Auth()].Handler())
			}
		} else {
			r.engine.Handle(endpoint.Method(), url, endpoint.Handler())
		}
	}

}
