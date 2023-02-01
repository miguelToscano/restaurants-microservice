package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	LOCALHOST = "localhost"
	PORT      = 8080
)

func Start() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	fmt.Printf("Listening on %s:%d", LOCALHOST, PORT)

	r.Run()
}
