package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/theamrendrasingh/boolipi/api"
	"github.com/theamrendrasingh/boolipi/db"
)

func main() {

	db.InitDB()

	// someid := "20ac7665-9451-4346-830f-396e92cc238a"

	// id, err := db.AddValue(true, "meowBool")

	// if err != nil {
	// 	fmt.Println("Something wrong")
	// }

	// fmt.Println(id)
	// db.Fetch(id)

	// fmt.Println(db.Delete(someid))

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	r.GET("/:id", api.Getting)
	r.POST("/", api.Posting)
	r.PATCH("/:id", api.Patching)
	r.DELETE("/:id", api.Deleting)

	r.Run()

}
