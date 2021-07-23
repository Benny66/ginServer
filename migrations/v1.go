package migrations

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-07 09:20:34
 * @LastEditors: shahao
 * @LastEditTime: 2021-07-23 15:12:17
 */

import (
	"database/sql"
	"ginServer/utils/migrate"
)

func (m *migration) migrateV1() []*migrate.Migration {
	return []*migrate.Migration{
		&migrate.Migration{
			Name:        "20210430144502_create_users_table",
			Description: "创建账号表",
			Up: func(db *sql.Tx) error {
				_, err := db.Exec(`
				CREATE TABLE users
						--账号表
						(
							"id" integer PRIMARY KEY AUTOINCREMENT,
							"username" varchar(20,2),--ip地址
						    "password" varchar(50,2),--mac地址
							"created_at" timestamp NULL DEFAULT NULL,   --创建时间
						    "updated_at" timestamp NULL DEFAULT NULL  --更新时间
						)
				`)
				if err != nil {
					return err
				}
				//密码md5(123456audio2021)
				_, err = db.Exec(`
					insert into users(username, password) values('admin', '1f6959b7c79fd0b227d617c743154fa4')
				`)
				if err != nil {
					return err
				}
				return nil
			},
			Down: func(db *sql.Tx) error {
				_, err := db.Exec(`
				DROP TABLE IF EXISTS users;
				`)
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
}
