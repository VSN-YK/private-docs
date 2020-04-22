package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type selectRecord struct {
	Id   string `gorm:"primary_key"`
	Name string
}

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("lang/:langId", func(ctx *gin.Context) {
			langId := ctx.Param("langId")
			db := ConnectDBWithGorm()
			rcd := selectRecord{}
			recordCounter := 0
			db.Table("LANGUAGE_TBL").Where("id = ?", langId).Find(&rcd).Count(&recordCounter)
			if recordCounter == 0 {
				rcd = selectRecord{}
			}
			ctx.JSON(200, gin.H{
				"LangInfo":          rcd,
				"ResultRecordCount": recordCounter,
			})

		})

		v1.GET("lang/", func(ctx *gin.Context) {
			db := ConnectDBWithGorm()
			rcds := []selectRecord{}
			db.Table("LANGUAGE_TBL").Find(&rcds)
			ctx.JSON(200, gin.H{
				"LangInfo": rcds,
			})
		})

	}
	r.Run(":8089")
}

func ConnectDBWithGorm() *gorm.DB {
	DBMS := "mysql"
	USER := "sub_user"
	PASS := "sub"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "sub"
	PARSETIME := "parseTime=true"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + PARSETIME
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic("Connection Error")
	}
	db.LogMode(true)
	return db
}
