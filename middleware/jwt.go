package middleware

import (
	"errors"
	"fmt"

	"github.com/Benny66/ginServer/config"
	"github.com/Benny66/ginServer/schemas"
	"github.com/Benny66/ginServer/utils/format"
	"github.com/Benny66/ginServer/utils/language"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if token == "" {
			format.NewResponseJson(context).Error(language.TOKEN_EMPTY)
			return
		}

		claims, err := ParseToken(token)
		if isTokenExpireError(err) {
			format.NewResponseJson(context).Error(language.TOKEN_EXPIRE)
			return
		}
		if isTokenInvalidError(err) {
			format.NewResponseJson(context).Error(language.TOKEN_INVALID)
			return
		}

		userInfo := schemas.UserInfo{
			UserId:   uint(claims["user_id"].(float64)),
			UserName: claims["username"].(string),
		}
		context.Set("user", userInfo)
		context.Set("token", token)
		context.Next()
	}
}

var (
	TokenInvalidErr = errors.New("token invalid")
	TokenExpireErr  = errors.New("token expire")
)

func CreateToken(claims jwt.MapClaims) (token string, err error) {
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenObj.SignedString([]byte(config.Config.TokenSecret))
	return
}

func ParseToken(token string) (claims jwt.MapClaims, err error) {
	tokenObj, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Config.TokenSecret), nil
	})

	var isOK bool
	if err != nil {
		ve, ok := err.(*jwt.ValidationError)
		if ok && (ve.Errors&(jwt.ValidationErrorExpired) != 0) {
			err = TokenExpireErr
			return
		}
		err = TokenInvalidErr
		return
	}
	claims, isOK = tokenObj.Claims.(jwt.MapClaims)
	if !isOK {
		err = TokenInvalidErr
		return
	}
	return
}

func isTokenInvalidError(err error) bool {
	return err == TokenInvalidErr
}

func isTokenExpireError(err error) bool {
	return err == TokenExpireErr
}
