package api

import "github.com/gin-gonic/gin"

func Deleting(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "deleting",
		"id":      id,
	})
}

// func Getting()
