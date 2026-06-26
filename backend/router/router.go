package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	authorized := r.Group("/")
	authorized.GET("/api/me", func(c *gin.Context) {
		userID, _ := c.Get("UserID")

		c.JSON(200, gin.H{
			"user_id": userID,
		})
	})

	return r
}