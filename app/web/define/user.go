package define

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-07 09:20:20
 * @LastEditors: shahao
 * @LastEditTime: 2021-07-23 15:19:53
 */

type UserLoginApiReq struct {
	UserName string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type UserUpdatePasswordApiReq struct {
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type UserInfo struct {
	UserId   uint   `json:"user_id"`
	UserName string `json:"username"`
}
