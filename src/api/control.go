package api

import (
	"fmt"
	"time"

	"github.com/theamrendrasingh/boolipi/auth"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/theamrendrasingh/boolipi/db"
)

// Request struct to denote the request JSON
type Request struct {
	Value bool
	Key   string `json:"key" binding:"required"`
}

//Tokener : return a new generated token. Using the current timestamp as the username and a generated uuid as the password
func Tokener(c *gin.Context) {
	pswd := uuid.NewV4().String()
	user := time.Now().String()
	token, err := auth.GenerateAccessToken(user, pswd)
	if err != nil {
		c.Writer.WriteHeader(500)
		return
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}

// Posting to handle POST requests for creating a new boolean
func Posting(c *gin.Context) {

	var request Request
	err := c.BindJSON(&request)

	if err != nil {
		fmt.Println(err)
		c.Writer.WriteHeader(400)
		return
	}

	u2 := uuid.NewV4().String()
	e, err := db.GetRepo().Create(u2, request.Value, request.Key)
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
	e, err := db.GetRepo().Fetch(id)
	if err != nil && err.Error() == "record not found" {
		c.Writer.WriteHeader(404)
		return
	}
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

// Patching : function to handle patch request, to modify an existing booelan identified by id
func Patching(c *gin.Context) {

	var request Request
	err := c.BindJSON(&request)

	if err != nil {
		fmt.Println(err)
		c.Writer.WriteHeader(400)
		return
	}

	e, err := db.GetRepo().Patch(c.Param("id"), request.Value, request.Key)
	if err != nil && err.Error() == "record not found" {
		c.Writer.WriteHeader(404)
		return
	}
	if err != nil {
		c.Writer.WriteHeader(500)
		return
	}

	c.JSON(200, gin.H{
		"id":    e.Uuid,
		"value": e.Value,
		"key":   e.Key,
	})
	return
}

// Deleting : function to handle delete requests
func Deleting(c *gin.Context) {
	id := c.Param("id")
	err := db.GetRepo().Delete(id)

	if err != nil && err.Error() == "record not found" {
		c.Writer.WriteHeader(404)
		return
	}

	if err != nil {
		c.Writer.WriteHeader(500)
		return
	}
	c.Writer.WriteHeader(204)
	return
}

//NoRoute : Handle requests which do not match any defined route
func NoRoute(c *gin.Context) {
	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}
