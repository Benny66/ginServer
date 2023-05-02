package routers

import "github.com/gin-gonic/gin"

type Endpoint interface {
	Group() string
	Auth() string
	Method() string
	URL() string
	Handler() gin.HandlerFunc
}

type MiddlewareSchema interface {
	Name() string
	Handler() gin.HandlerFunc
}
