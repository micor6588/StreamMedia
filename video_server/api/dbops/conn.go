package dbops

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

// 数据库的初始化
func init() {
	dbConn, err = sql.Open("mysql", "root:371871@tcp(127.0.0.1:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
