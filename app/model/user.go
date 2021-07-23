package model

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-14 09:56:53
 * @LastEditors: shahao
 * @LastEditTime: 2021-07-23 15:56:33
 */

type User struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"id"`
	UserName  string    `gorm:"column:username;unique;not null" json:"username"`
	Password  string    `gorm:"column:password;not null" json:"password"`
	CreatedAt ModelTime `gorm:"column:created_at" json:"created_at"`
	UpdatedAt ModelTime `gorm:"column:updated_at" json:"updated_at"`
}

func (um User) TableName() string {
	return "users"
}

//对外查询使用用户模型
type UserFind struct {
	User
	Password string `gorm:"column:password;not null" json:"-"`
}
