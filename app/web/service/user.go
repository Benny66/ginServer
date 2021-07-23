package service

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-07 09:20:20
 * @LastEditors: shahao
 * @LastEditTime: 2021-07-23 15:20:02
 */

import (
	"ginServer/app/dao"
	"ginServer/app/web/define"
	"ginServer/utils/database"
	"ginServer/utils/function"
	"ginServer/utils/jwt"
	"ginServer/utils/language"
	"ginServer/utils/log"
	"regexp"
	"time"

	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var UserService *userService

func init() {
	UserService = NewUserService()
}

func NewUserService() *userService {
	return &userService{}
}

type userService struct {
}

func (svr *userService) Login(req *define.UserLoginApiReq) (data interface{}, err IServiceError) {

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
	userInfo, dbErr := dao.UserDao.FindOneWhere("username = ? and password = ?", req.UserName, password)
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

	data, tokenErr = jwt.CreateToken(claims)
	if tokenErr != nil {
		log.SystemLog(tokenErr)
		err = NewServiceError(language.USER_TOKEN_CREATE_ERROR)
		return
	}
	return
}
func (svr *userService) Refresh(req define.UserInfo) (data interface{}, err IServiceError) {
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

	data, tokenErr = jwt.CreateToken(claims)
	if tokenErr != nil {
		log.SystemLog(tokenErr)
		err = NewServiceError(language.USER_TOKEN_CREATE_ERROR)
		return
	}
	return
}

func (svr *userService) UpdatePassword(req *define.UserUpdatePasswordApiReq, user define.UserInfo) (data interface{}, err IServiceError) {
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

	userInfo, dbErr := dao.UserDao.FindOneWhere("id = ?", user.UserId)
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
	data, dbErr = dao.UserDao.Update(database.Orm.DB(), userInfo.ID, map[string]interface{}{
		"password": password,
	})
	if dbErr != nil {
		log.SystemLog(dbErr)
		err = NewServiceError(language.DB_ERROR)
		return
	}
	return
}

func (svr *userService) User(context *gin.Context) (user define.UserInfo) {
	u, isOK := context.Get("user")
	if isOK {
		user = u.(define.UserInfo)
	}
	return
}

func (svr *userService) checkIpAddress(ipAddress string) (err IServiceError) {
	if ipAddress == "" {
		err = NewServiceError(language.SERVER_IP_ADDRESS_EMPTY)
		return
	}
	ok, _ := regexp.MatchString("((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})(\\.((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})){3}", ipAddress)
	if !ok {
		err = NewServiceError(language.SERVER_IP_ADDRESS_ERROR, ipAddress)
		return
	}
	if !function.IsPing(ipAddress) {
		err = NewServiceError(language.SERVER_IP_ADDRESS_PING_ERROR, ipAddress)
		return
	}
	return
}
