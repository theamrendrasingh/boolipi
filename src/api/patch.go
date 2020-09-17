package api

import "github.com/gin-gonic/gin"

func Patching(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "patching",
		"id":      id,
	})
}

// func Getting()
