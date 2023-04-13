package models

import (
	database "github.com/Benny66/ginServer/db"

	"gorm.io/gorm"
)

var UserDao *userDao

func init() {
	UserDao = NewUserDao()
}

func NewUserDao() *userDao {
	return &userDao{
		gm: database.Orm.DB(),
	}
}

type userDao struct {
	gm *gorm.DB
}

func (dao *userDao) Create(tx *gorm.DB, data *UserModel) (rowsAffected int64, err error) {
	db := tx.Create(data)
	if err = db.Error; db.Error != nil {
		return
	}
	rowsAffected = db.RowsAffected
	return
}

func (dao *userDao) Update(tx *gorm.DB, id uint, data map[string]interface{}) (rowsAffected int64, err error) {
	db := tx.Model(&UserModel{}).Where("id = ?", id).Updates(data)
	if err = db.Error; db.Error != nil {
		return
	}
	rowsAffected = db.RowsAffected
	return
}

func (dao *userDao) Delete(tx *gorm.DB, data []int) (rowsAffected int64, err error) {
	db := tx.Where("id in (?)", data).Delete(&UserModel{})
	if err = db.Error; db.Error != nil {
		return
	}
	rowsAffected = db.RowsAffected
	return
}

func (dao *userDao) FindAll() (list []UserFindModel, err error) {
	db := dao.gm.Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *userDao) FindAllWhere(query interface{}, args ...interface{}) (list []UserFindModel, err error) {
	db := dao.gm.Where(query, args...).Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *userDao) FindOneWhere(query interface{}, args ...interface{}) (record UserModel, err error) {
	db := dao.gm.Where(query, args...).First(&record)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *userDao) FindCountWhere(query interface{}, args ...interface{}) (count int64, err error) {
	db := dao.gm.Model(&UserModel{}).Where(query, args...).Count(&count)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *userDao) FindCount() (count int64, err error) {
	db := dao.gm.Model(&UserModel{}).Count(&count)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *userDao) Raw(sqlStr string, params ...interface{}) (list []UserFindModel, err error) {
	db := dao.gm.Debug().Raw(sqlStr, params...).Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}
func (dao *userDao) WhereQuery(query interface{}, args ...interface{}) *userDao {
	return &userDao{
		dao.gm.Where(query, args...),
	}

}

func (dao *userDao) WhereUserNameLike(username string) *userDao {
	return &userDao{
		dao.gm.Where("username like ?", "%"+username+"%"),
	}
}

func (dao *userDao) WhereDisabled(isDisabled int) *userDao {
	return &userDao{
		dao.gm.Where("is_disabled = ?", isDisabled),
	}
}

func (dao *userDao) Paginate(offset, limit int) (count int64, list []UserFindModel, err error) {
	db := dao.gm.Model(&UserFindModel{}).Count(&count).Offset(offset).Limit(limit).Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *userDao) Debug() *userDao {
	return &userDao{
		dao.gm.Debug(),
	}
}

func (dao *userDao) Offset(offset int) *userDao {
	return &userDao{
		dao.gm.Offset(offset),
	}
}

func (dao *userDao) Limit(limit int) *userDao {
	return &userDao{
		dao.gm.Limit(limit),
	}
}

func (dao *userDao) OrderBy(sortFlag, sortOrder string) *userDao {
	return &userDao{
		dao.gm.Order(sortFlag + " " + sortOrder),
	}
}

func (dao *userDao) Joins(query string, args ...interface{}) *userDao {
	return &userDao{
		dao.gm.Joins(query, args),
	}
}

func (dao *userDao) Preloads(query string) *userDao {
	return &userDao{
		dao.gm.Preload(query),
	}
}
