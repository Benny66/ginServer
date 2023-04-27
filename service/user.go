package service

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-07 09:20:20
 * @LastEditors: shahao
 * @LastEditTime: 2021-07-23 17:21:16
 */

import (
	"time"

	database "github.com/Benny66/ginServer/db"
	"github.com/Benny66/ginServer/log"
	"github.com/Benny66/ginServer/middleware"
	"github.com/Benny66/ginServer/models"
	"github.com/Benny66/ginServer/schemas"
	"github.com/Benny66/ginServer/utils/function"
	"github.com/Benny66/ginServer/utils/language"

	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserInterface interface {
	Login(req *schemas.UserLoginApiReq) (data interface{}, err IServiceError)
	Refresh(req schemas.UserInfo) (data interface{}, err IServiceError)
	UpdatePassword(req *schemas.UserUpdatePasswordApiReq, user schemas.UserInfo) (data interface{}, err IServiceError)
	User(context *gin.Context) (user schemas.UserInfo)
}

var UserService UserInterface = &userService{}

type userService struct{}

func (svr *userService) Login(req *schemas.UserLoginApiReq) (data interface{}, err IServiceError) {
	if req.UserName == "" {
		err = NewServiceError(language.USER_NAME_EMPTY)
		return
	}
	if req.Password == "" {
		err = NewServiceError(language.USER_PASS_EMPTY)
		return
	}
	//md5加盐
	password := function.CreateMD5(req.Password + "audio2021")
	userInfo, dbErr := models.UserDao.FindOneWhere("username = ? and password = ?", req.UserName, password)
	if dbErr != nil {
		log.SystemLog(dbErr)
		err = NewServiceError(language.USER_LOGIN_ERROR)
		return
	}
	var tokenErr error
	claims := jwt2.MapClaims{
		"user_id":  userInfo.ID,
		"username": userInfo.UserName,
		"exp":      time.Now().Add(2 * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
		"iss":      "GP",
		"platform": "web",
		"sub":      "all",
		"typ":      "JWT",
	}

	data, tokenErr = middleware.CreateToken(claims)
	if tokenErr != nil {
		log.SystemLog(tokenErr)
		err = NewServiceError(language.USER_TOKEN_CREATE_ERROR)
		return
	}
	return
}
func (svr *userService) Refresh(req schemas.UserInfo) (data interface{}, err IServiceError) {
	var tokenErr error
	claims := jwt2.MapClaims{
		"user_id":  req.UserId,
		"username": req.UserName,
		"exp":      time.Now().Add(2 * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
		"iss":      "GP",
		"platform": "web",
		"sub":      "all",
		"typ":      "JWT",
	}

	data, tokenErr = middleware.CreateToken(claims)
	if tokenErr != nil {
		log.SystemLog(tokenErr)
		err = NewServiceError(language.USER_TOKEN_CREATE_ERROR)
		return
	}
	return
}

func (svr *userService) UpdatePassword(req *schemas.UserUpdatePasswordApiReq, user schemas.UserInfo) (data interface{}, err IServiceError) {
	if req.OldPassword == "" || req.NewPassword == "" || req.ConfirmPassword == "" {
		err = NewServiceError(language.INVALID_PARMAS)
		return
	}
	if req.NewPassword != req.ConfirmPassword {
		err = NewServiceError(language.USER_CONFIRM_PASS_NOT_MATCH)
		return
	}
	//md5加盐
	password := function.CreateMD5(req.OldPassword + "audio2021")

	userInfo, dbErr := models.UserDao.FindOneWhere("id = ?", user.UserId)
	if dbErr != nil {
		log.SystemLog(dbErr)
		err = NewServiceError(language.USER_NOT_EXISTS)
		return
	}
	if userInfo.Password != password {
		err = NewServiceError(language.USER_PASS_NOT_MATCH)
		return
	}
	password = function.CreateMD5(req.NewPassword + "audio2021")
	data, dbErr = models.UserDao.Update(database.Orm.DB(), userInfo.ID, map[string]interface{}{
		"password": password,
	})
	if dbErr != nil {
		log.SystemLog(dbErr)
		err = NewServiceError(language.DB_ERROR)
		return
	}
	return
}

func (svr *userService) User(context *gin.Context) (user schemas.UserInfo) {
	u, isOK := context.Get("user")
	if isOK {
		user = u.(schemas.UserInfo)
	}
	return
}
