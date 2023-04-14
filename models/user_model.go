package models

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
