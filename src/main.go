package main

import (
	"github.com/gin-gonic/gin"
	"github.com/theamrendrasingh/boolipi/api"
)

func main() {
	r := gin.Default()

	r.GET("/:id", api.Getting)
	r.POST("/", api.Posting)
	r.PATCH("/:id", api.Patching)
	r.DELETE("/:id", api.Deleting)

	r.Run()
}
