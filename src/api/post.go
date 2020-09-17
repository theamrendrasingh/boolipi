package api

import "github.com/gin-gonic/gin"

func Posting(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Posting",
	})
}

// func Getting()
