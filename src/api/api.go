package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/theamrendrasingh/boolipi/db"
)

// Request struct to denote the request JSON
type Request struct {
	Value bool
	Key   string `json:"key" binding:"required"`
}

// Posting to handle POST requests for creating a new boolean
func Posting(c *gin.Context) {

	var request Request
	err := c.BindJSON(&request)

	if err != nil {
		fmt.Println(err)
	}

	e, err := db.AddValue(request.Value, request.Key)

	if err != nil {
		c.Writer.WriteHeader(500)
		return
	}

	c.JSON(200, gin.H{
		"id":    e.Uuid,
		"key":   e.Key,
		"value": e.Value,
	})
	return
}

// Getting : To handle GET requests for retieving
func Getting(c *gin.Context) {
	id := c.Param("id")
	e, err := db.Fetch(id)

	if err != nil {
		c.Writer.WriteHeader(500)
		return
	}

	c.JSON(200, gin.H{
		"id":    e.Uuid,
		"value": e.Value,
		"key":   e.Key,
	})
}

// Patching function to handle patch request, to modify an existing booelan identified by id
func Patching(c *gin.Context) {

	var request Request
	err := c.BindJSON(&request)

	if err != nil {
		fmt.Println(err)
	}

	e, err := db.Patch(c.Param("id"), request.Value, request.Key)

	if err != nil {
		c.Writer.WriteHeader(500)
		return
	}

	c.JSON(200, gin.H{
		"id":    e.Uuid,
		"value": e.Value,
		"name":  e.Key,
	})
	return
}

// Deleting function to delete an entry
func Deleting(c *gin.Context) {
	id := c.Param("id")
	err := db.Delete(id)

	if err != nil {
		c.Writer.WriteHeader(500)
		return
	}
	c.Writer.WriteHeader(204)
	return
}
