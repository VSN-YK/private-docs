package main

import (
	"fmt"
	"github.com/comail/colog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

// Set Up in Colog Condition
func setLogCondition() {
	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()
}

// difinate db recode type
type ResultRecord struct {
	Id   string `gorm:"primary_key" json:"id"`
	Name string
}

func mysqlConnectWithGorm() *gorm.DB {
	DBMS := "mysql"
	USER := "sub_user"
	PASS := "sub"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "sub"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	fmt.Println(CONNECT)
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	db.LogMode(true)
	fmt.Println("Connected!!")
	return db
}

func InsertCheck(db *gorm.DB, id string) bool {
	rcd := []ResultRecord{}
	db.Table("LANGUAGE_TBL").Find(&rcd)
	for _, r := range rcd {
		if id == r.Id {
			fmt.Println(r)
			return false
		}
	}
	return true
}

func main() {
	db := mysqlConnectWithGorm()
	defer db.Close()
	setLogCondition()
	log.Printf("debug: Debug")
	if InsertCheck(db, "L003") {
		fmt.Println("*********Insert Data********")
		insert_rcd := ResultRecord{}
		insert_rcd.Id = "L003"
		insert_rcd.Name = "Python"
		db.Table("LANGUAGE_TBL").Create(&insert_rcd)
		fmt.Println(insert_rcd)
	} else {
		fmt.Println("ALL Ready Exist")
		os.Exit(1)
	}

	fmt.Println("******Select*********")
	rcd := []ResultRecord{}
	db.Table("LANGUAGE_TBL").Find(&rcd)
	fmt.Println(rcd)

	fmt.Println("******Update Record*********")
	updRecordBefore := ResultRecord{}
	updRecordBefore.Id = "L003"
	db.Table("LANGUAGE_TBL").First(&updRecordBefore)
	fmt.Println(updRecordBefore)
	updRecordAfter := updRecordBefore
	updRecordAfter.Name = "C++"
	db.Table("LANGUAGE_TBL").Model(&updRecordBefore).Update(&updRecordAfter)

	fmt.Println("******DeleteRecord*********")
	delRecord := ResultRecord{}
	delRecord.Id = "L003"
	db.Table("LANGUAGE_TBL").First(&delRecord)
	fmt.Println(delRecord)
	db.Table("LANGUAGE_TBL").Delete(&delRecord)
	log.Printf("debug:%v", delRecord)
}
