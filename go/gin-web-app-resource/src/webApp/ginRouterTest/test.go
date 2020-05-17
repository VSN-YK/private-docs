package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	fmt.Printf("%v", r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ping!",
		})
	})
	r.Run(":8089") //listen and serv on 127.0.0.1:8089
}
