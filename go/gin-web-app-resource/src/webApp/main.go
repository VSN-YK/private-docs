package main

import (
	"./consoleLogger"
	"./db"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type ReservedSessionRecord struct {
	Session_id string
	Saved_date time.Time
}

func main() {

	consoleLogger.CustomConsoleLog()
	log.Printf("info: This Is Debug")

	title := "Gin App"
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	//template配下のhtmlを適応する
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/assets", "./assets")
	//リクエストハンドラーの定義
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{
			"title": title,
		})
	})

	router.POST("/wellcome", LoginHandler)
	router.POST("/logout", LogoutHandler)
	router.POST("/registerAccount", accountViewerHundler)
	router.POST("/registerAccountOperation", createAccountHandler)
	router.Run(":8089")
}

func accountViewerHundler(ctx *gin.Context) {
	ctx.HTML(200, "registerAccount.html", gin.H{"title": "Register Application Account"})
}

func LoginHandler(ctx *gin.Context) {
	userId := ctx.PostForm("loginId")
	fmt.Println(userId)
	message, loginStatus := loginCheck(ctx, userId)
	log.Printf("info: [Resolved This Error By Function]%s", message)
	ctx.HTML(200, "wellcome.html", gin.H{
		"userId":      userId,
		"message":     message,
		"loginStatus": loginStatus,
	})
}

func deleteUserRecord(loginId string) {
	log.Printf("debug: %s", loginId)
	db := db.ConnectDBWithGorm()
	defer db.Close()
	deleteLoginRecord := ReservedSessionRecord{}
	db.Table("MANAGE_SESSIONS").Where("session_id =?", loginId).Delete(&deleteLoginRecord)
}

func LogoutHandler(ctx *gin.Context) {
	fmt.Printf("debug: Logout Handler Called!")
	session := sessions.Default(ctx)
	loginUserId := session.Get("LoginId").(string)
	log.Printf("debug: [UserInfo]%s", loginUserId)
	session.Clear()
	session.Save()
	deleteUserRecord(loginUserId)
	ctx.Redirect(http.StatusMovedPermanently, "/")
}

func isExistAccount(loginId string) bool {
	db := db.ConnectDBWithGorm()
	defer db.Close()
	existRecordCounter := 0
	existRecord := ReservedSessionRecord{}
	existRecord.Session_id = loginId
	log.Printf("debug: [AccountStatus] %v", existRecord)
	//db.Table("MANAGE_SESSIONS").Find(&existRecord).Count(&existRecordCounter)
	db.Table("MANAGE_SESSIONS").Where("session_id = ?", loginId).Find(&existRecord).Count(&existRecordCounter)
	log.Printf("debug: %d", existRecordCounter)
	if existRecordCounter == 0 {
		return false
	} else {
		return true
	}
}

func createAccountHandler(ctx *gin.Context) {
	loginId := ctx.PostForm("loginId")
	session := sessions.Default(ctx)

	if !isExistAccount(loginId) {
		db := db.ConnectDBWithGorm()
		defer db.Close()
		session.Set("LoginId", loginId)
		session.Save()

		insertLoginRecord := ReservedSessionRecord{}
		insertLoginRecord.Session_id = loginId
		insertLoginRecord.Saved_date = time.Now()
		db.Table("MANAGE_SESSIONS").Create(&insertLoginRecord)
		ctx.Redirect(http.StatusMovedPermanently, "/")
	} else {
		session.Set("LoginId", loginId)
		session.Save()
		ctx.Redirect(http.StatusMovedPermanently, "/")
	}
}

func loginCheck(ctx *gin.Context, loginId string) (string, bool) {
	log.Printf("debug:%s", loginId)
	db := db.ConnectDBWithGorm()
	defer db.Close()
	record := ReservedSessionRecord{}
	record.Session_id = loginId
	log.Printf("debug:[LoginRecord]%v", record)
	resultRecordCount := 0
	db.Table("MANAGE_SESSIONS").Find(&record).Count(&resultRecordCount)
	log.Printf("debug[RecordStatus]:%s\n%v", record.Session_id, record)

	session := sessions.Default(ctx)
	log.Printf("info: [Getting Current Session Info]%v", session)
	if resultRecordCount != 0 {
		log.Printf("debug:[Session Id]%v", session.Get("LoginId"))
		if session.Get("LoginId") == loginId {
			return loginId + "is Member of This Application", true
		} else {
			return "Not a Memeber of This Application", false
		}
	}
	return "First Time", false
}
