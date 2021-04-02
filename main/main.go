package main

import (
	postgres "github.com/woon-acornsoft/turbo-lamp/db"

	"github.com/gin-gonic/gin"
)

func main() {
	postgres.ExampleDB_Model()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
