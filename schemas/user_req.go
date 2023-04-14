package schemas

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
