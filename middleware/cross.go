package middleware

import (
	"strings"

	"github.com/Benny66/ginServer/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	routers.R.AddMiddlewareSchema(&cross{})
}

type cross struct{}

func (m *cross) Name() string {
	return "cross"
}
func (m *cross) Handler() gin.HandlerFunc {
	return func(context *gin.Context) {
		if !strings.HasPrefix(context.Request.URL.Path, "/docs") {
			context.Header("Access-Control-Allow-Origin", "*")
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			context.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			context.Header("Content-Type", "application/json; charset=utf-8")
		}
		context.Next()
	}
}
