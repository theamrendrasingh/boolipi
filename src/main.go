package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/theamrendrasingh/boolipi/api"
	"github.com/theamrendrasingh/boolipi/db"
)

func main() {

	d, err := db.InitDB()
	if err != nil {
		log.Fatal("Unable to initialize the database setup")
	}
	db.Database = db.Db{DB: d}
	repoImpl := db.RepoImpl{}
	db.SetRepo(&repoImpl)

	// err = db.InitDB()

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	r.GET("/:id", api.Getting)
	r.POST("/", api.Posting)
	r.PATCH("/:id", api.Patching)
	r.DELETE("/:id", api.Deleting)
	r.NoRoute(api.NoRoute)

	r.Run()

}
