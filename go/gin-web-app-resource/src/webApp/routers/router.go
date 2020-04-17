package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/test-send", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Ping Test",
			"Status":  true,
		})
	})

	return r
}
