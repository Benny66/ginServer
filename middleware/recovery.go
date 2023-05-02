package middleware

import (
	"fmt"
	"runtime/debug"

	"github.com/Benny66/ginServer/log"
	"github.com/Benny66/ginServer/routers"
	"github.com/Benny66/ginServer/utils/format"
	"github.com/Benny66/ginServer/utils/language"

	"github.com/gin-gonic/gin"
)

func init() {
	routers.R.AddMiddlewareSchema(&Recover{})
}

type Recover struct{}

func (m *Recover) Name() string {
	return "recover"
}

func (m *Recover) Handler() gin.HandlerFunc {
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
