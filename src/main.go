package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/theamrendrasingh/boolipi/api"
	"github.com/theamrendrasingh/boolipi/auth"
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

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	var r *gin.Engine

	if v, f := os.LookupEnv("USE_AUTH"); f && v == "true" {
		fmt.Println("Using Auth")
		r = authRouter()
	} else {
		r = router()
	}

	r.Run()

}
func router() *gin.Engine {
	r := gin.Default()
	r.GET("/:id", api.Getting)
	r.POST("/", api.Posting)
	r.PATCH("/:id", api.Patching)
	r.DELETE("/:id", api.Deleting)
	r.NoRoute(api.NoRoute)

	return r
}

func authRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/token", api.Tokener)
	authGroup := r.Group("/")
	authGroup.Use(auth.AuthMiddleware())
	{
		authGroup.GET("/:id", api.Getting)
		authGroup.POST("/", api.Posting)
		authGroup.PATCH("/:id", api.Patching)
		authGroup.DELETE("/:id", api.Deleting)
	}
	return r
}
