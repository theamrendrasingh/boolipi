package api

import "github.com/gin-gonic/gin"

func Getting(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "getting",
		"id":      id,
	})
}

// func Getting()
