package middleware


import (
	"strings"

	"github.com/gin-gonic/gin"
)

func CrossMiddleware() gin.HandlerFunc {
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
