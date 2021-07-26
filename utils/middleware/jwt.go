package middleware

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-07 09:20:20
 * @LastEditors: shahao
 * @LastEditTime: 2021-07-26 17:20:29
 */

import (
	"ginServer/app/web/define"
	"ginServer/utils/format"
	"ginServer/utils/jwt"
	"ginServer/utils/language"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if token == "" {
			format.NewResponseJson(context).Error(language.TOKEN_EMPTY)
			return
		}

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
