package middleware

import (
	"fmt"
	"ginServer/utils/format"
	"ginServer/utils/language"
	"ginServer/utils/log"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.SystemLog(fmt.Sprintf("%s", err))
				if gin.IsDebugging() {
					debug.PrintStack()
				}
				format.NewResponseJson(context).Error(language.SERVER_PANIC)
			}
		}()
		context.Next()
	}
}
