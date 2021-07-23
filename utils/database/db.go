package database

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-07 09:20:21
 * @LastEditors: shahao
 * @LastEditTime: 2021-07-23 15:17:36
 */

import (
	"database/sql"
	"ginServer/utils/function"
	log2 "ginServer/utils/log"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Orm *orm

func init() {
	Orm = NewOrm()
}

func NewOrm() *orm {
	o := new(orm)
	o.loadDBConfig()
	return o
}

type orm struct {
	engine *gorm.DB
}

func (db *orm) loadDBConfig() {
	if db.engine != nil {
		return
	}

	var err error
	// dbPath := filepath.Join(filepath.Dir(strings.TrimRight(function.GetCurrentAbsPath(), string(os.PathSeparator))), config.Config.GetDBPath())
	// log2.SystemLog("数据库文件路径" + dbPath)
	// if !function.IsFileExists(dbPath) {
	// 	log2.SystemLog("数据库文件路径不存在")
	// 	return
	// }
	db.engine, err = gorm.Open(sqlite.Open(function.GetAbsPath("db/ginServer.db")), &gorm.Config{})
	if err != nil {
		log2.SystemLog(err)
		log.Fatal("open database error: " + err.Error())
	}
	sqlDB, _ := db.engine.DB()
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(1000)                 //必须小于等于数据库max_connections参数值
	sqlDB.SetConnMaxLifetime(110 * time.Second) //必须小于数据库wait_timeout参数值
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
func (db *orm) DB() *gorm.DB {
	return db.engine
}

func (db *orm) SqlDb() *sql.DB {
	e, _ := db.engine.DB()
	return e
}
