package jwt

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-07 09:52:53
 * @LastEditors: shahao
 * @LastEditTime: 2021-04-07 09:58:55
 */

import (
	"errors"
	"fmt"
	"ginServer/config"

	"github.com/dgrijalva/jwt-go"
)

var (
	TokenInvalidErr = errors.New("token invalid")
	TokenExpireErr  = errors.New("token expire")
)

func CreateToken(claims jwt.MapClaims) (token string, err error) {
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenObj.SignedString([]byte(config.Config.GetTokenSecret()))
	return
}

func ParseToken(token string) (claims jwt.MapClaims, err error) {
	tokenObj, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Config.GetTokenSecret()), nil
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

func IsTokenInvalidError(err error) bool {
	return err == TokenInvalidErr
}

func IsTokenExpireError(err error) bool {
	return err == TokenExpireErr
}
