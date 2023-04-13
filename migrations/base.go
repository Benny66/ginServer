package migrations

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-07 09:20:35
 * @LastEditors: shahao
 * @LastEditTime: 2021-07-23 14:50:51
 */

import (
	"database/sql"
	"fmt"
	"os"

	database "github.com/Benny66/ginServer/db"
	"github.com/Benny66/ginServer/utils/function"
	"github.com/Benny66/ginServer/utils/migrate"
)

var Migration *migration

func init() {
	Migration = NewMigration(database.Orm.SqlDb(), function.GetAbsPath("install.lock"))
}

func NewMigration(sqlDB *sql.DB, lockFilePath string) *migration {
	m := &migration{
		sqlDB:        sqlDB,
		lockFilePath: lockFilePath,
	}
	fmt.Println(m.migrateTable())
	return m
}

type migration struct {
	sqlDB        *sql.DB
	lockFilePath string
}

func (m *migration) isExistsLockFile() bool {
	return function.IsFileExists(m.lockFilePath)
}

func (m *migration) migrateTable() error {
	if !m.isExistsLockFile() {
		return nil
	}

	obj, err := migrate.NewGoMigration(m.sqlDB)
	if err != nil {
		return err
	}

	//第一版数据库迁移
	err = obj.AddMigration(m.migrateV1()...)
	if err != nil {
		return err
	}

	_, err = obj.Migrate()
	if err != nil {
		return err
	}

	err = m.migrateData()
	if err != nil {
		return err
	}

	m.deleteLockFile()
	return nil
}

func (m *migration) deleteLockFile() error {
	return os.Remove(m.lockFilePath)
}

// 填充默认数据
func (m *migration) migrateData() error {
	// to do

	return nil
}
