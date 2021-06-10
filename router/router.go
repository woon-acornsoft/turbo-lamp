package router

import (
	"github.com/gin-gonic/gin"
	"github.com/woon-acornsoft/turbo-lamp/service"
)

func Router() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/users", func(c *gin.Context) {
		users := service.Users()
		c.JSON(200, users)
	})
	r.GET("/user/:userId", func(c *gin.Context) {
		userId := c.Param("userId")
		user := service.User(userId)
		c.JSON(200, user)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
