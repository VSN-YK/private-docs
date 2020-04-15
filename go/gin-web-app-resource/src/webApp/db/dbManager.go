package db

import (
	"../consoleLogger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func ConnectDBWithGorm() *gorm.DB {
	consoleLogger.CustomConsoleLog()
	DBMS := "mysql"
	USER := "sub_user"
	PASS := "sub"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "sub"
	PARSETIME := "parseTime=true"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + PARSETIME
	log.Printf("debug:%s", CONNECT)
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic("Connection Error")
	}
	db.LogMode(true)
	return db
}
