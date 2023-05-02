package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Benny66/ginServer/config"

	"gorm.io/driver/mysql"
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
	dbName := config.Config.DBName
	dbUser := config.Config.DBUsername
	dbPassword := config.Config.DBPassword
	dbHost := config.Config.DBHost
	dbPort := config.Config.DBPort
	// Prefix := "m_"
	db.engine, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)), &gorm.Config{})
	if err != nil {
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
