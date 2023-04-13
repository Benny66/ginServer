项目地址：[https://github.com/Benny66/github.com/Benny66/ginServer](https://github.com/Benny66/github.com/Benny66/ginServer)

此框架目前有跨域、jwt、日志、异常捕获等中间件，使用很简单，参考gin中间件使用方法。
```
r := gin.New()
r.Use(middleware.LoggerMiddleware())
r.Use(middleware.Recover())
r.Use(middleware.CrossMiddleware())
r.Use(middleware.JWTMiddleware())
```
所谓中间件，在正式处理http请求之前，做一层逻辑控制，可以校验token合法性，可以打印好请求数据的日志，可以捕获错误的异常等等

按作用域可以分为存在全局中间件和局部中间件。


### jwt
主要进行用户合法性令牌（token）校验,合法的则解析数据后进行下一步逻辑处理，不合法直接返回错误提醒。

```
graph LR
获取token --> jwt解析校验token 
jwt解析校验token--> context.Set设置数据
context.Set设置数据-->context.Next逻辑处理
```
代码展示，userInfo数据可自定义
```
func JWTMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if token == "" {
			format.NewResponseJson(context).Error(language.TOKEN_EMPTY)
			return
		}
        //jwt解析校验token ，使用github.com/dgrijalva/jwt-go
		claims, err := jwt.ParseToken(token)
		if jwt.IsTokenExpireError(err) {
			format.NewResponseJson(context).Error(language.TOKEN_EXPIRE)
			return
		}
		if jwt.IsTokenInvalidError(err) {
			format.NewResponseJson(context).Error(language.TOKEN_INVALID)
			return
		}

		userInfo := define.UserInfo{
			UserId:   uint(claims["user_id"].(float64)),
			UserName: claims["username"].(string),
		}
		context.Set("user", userInfo)
		context.Set("token", token)
		context.Next()
	}
}

```
### 日志

主要将http请求的各项参数（请求头、请求参数、请求地址等）及请求时间打印出来，可保存到文件内。

```
func LoggerMiddleware() gin.HandlerFunc {
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
```

### 跨域
```
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

```

### 异常捕获
```

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

```