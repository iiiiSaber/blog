package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)
var Db *sql.DB
const(
	DbDriver   = "mysql"
	DbUserName = "myBlog"
	DbPasswd   = "439956461"
	DbHost     = "127.0.0.1"
	DbPort     = 3306
	DbName     = "myBlog"
)
func InitDb(){
	var err error
	Db, err = sql.Open(DbDriver, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		DbUserName, DbPasswd, DbHost, DbPort, DbName))
	if err != nil {
		panic(err)
	}
	defer Db.Close()
	err = Db.Ping()
	log.Println("success")
}