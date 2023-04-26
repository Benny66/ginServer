package models

import (
	database "github.com/Benny66/ginServer/db"
	"gorm.io/gorm"
)

type UserInterface interface {
	Create(tx *gorm.DB, data *UserModel) (rowsAffected int64, err error)
	Update(tx *gorm.DB, id uint, data map[string]interface{}) (rowsAffected int64, err error)
	Delete(tx *gorm.DB, data []int) (rowsAffected int64, err error)
	FindOneWhere(query interface{}, args ...interface{}) (record UserModel, err error)
	FindAll() (list []UserFindModel, err error)
	FindAllWhere(query interface{}, args ...interface{}) (list []UserFindModel, err error)
	FindCountWhere(query interface{}, args ...interface{}) (count int64, err error)
	FindCount() (count int64, err error)
	Raw(sqlStr string, params ...interface{}) (list []UserFindModel, err error)
}

var UserDao UserInterface = &userDao{
	gm: database.Orm.DB(),
}

type userDao struct {
	gm *gorm.DB
}

type UserModel struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"id"`
	UserName  string    `gorm:"column:username;unique;not null" json:"username"`
	Password  string    `gorm:"column:password;not null" json:"password"`
	CreatedAt ModelTime `gorm:"column:created_at" json:"created_at"`
	UpdatedAt ModelTime `gorm:"column:updated_at" json:"updated_at"`
}

func (um UserModel) TableName() string {
	return "users"
}

// 对外查询使用用户模型
type UserFindModel struct {
	UserModel
	Password string `gorm:"column:password;not null" json:"-"`
}
