package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	v := os.Getenv("VAR1")

	msg := "Server is running successfully"
	if v != "" {
		msg = msg + " with environment variable " + v
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": msg,
		})
	})
	r.Run()
	// listen and serve on 0.0.0.0:8080
}
