package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/Benny66/ginServer/log"
	"github.com/Benny66/ginServer/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	routers.R.AddMiddlewareSchema(&logger{})
}

type logger struct{}

func (m *logger) Name() string {
	return "logger"
}

func (m *logger) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		if raw != "" {
			path = path + "?" + raw
		}

		fmt.Fprintf(log.SystemLogger, "[GIN] %v | %3d | %13v | %15s | %-7s %s\n",
			end.Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			clientIP,
			method,
			path,
		)

		if gin.IsDebugging() {
			fmt.Fprintf(os.Stdout, "[GIN] %v | %3d | %13v | %15s | %-7s %s\n",
				end.Format("2006/01/02 - 15:04:05"),
				statusCode,
				latency,
				clientIP,
				method,
				path,
			)
		}
	}
}
