package db

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	//USER : database username
	USER = "amru"
	//PASS database password
	PASS = "q"
	//HOST : database host
	HOST = "127.0.0.1"
	// PORT : database port
	PORT = "3306"
	//DBNAME : database name
	DBNAME = "boolipi"
)

// Database : var storing current instance of the datbase
var Database Db

//Db : struct to hold an instance of the database pointer
type Db struct {
	DB *gorm.DB
}

type (
	// Entry struct denoting a value, identified by a uuid string, having a name denoted by key and a boolean Value
	Entry struct {
		Uuid  string `gorm:"primaryKey;autoIncrement:false"`
		Value bool   `gorm:"not null;"`
		Key   string
	}
)

func createDBDsn() string {
	if v, f := os.LookupEnv("DOCKER_MODE"); f && v == "true" {
		fmt.Println("DOCKER MODE FOUND")
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	} else {
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/", USER, PASS, HOST, PORT)
	}

}

func createDsn() string {
	if v, f := os.LookupEnv("DOCKER_MODE"); f && v == "true" {
		fmt.Println("DOCKER MODE FOUND")
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), DBNAME)
	} else {
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	}

}

//InitDB to connect to the database server, create the database if not exists, create the required tables
func InitDB() (*gorm.DB, error) {

	DBDsn := createDBDsn()
	db, err := gorm.Open("mysql", DBDsn)
	if err != nil {
		fmt.Println("Unable to connect database : ", err)
		return nil, err
	}

	db = db.Exec("CREATE DATABASE IF NOT EXISTS " + DBNAME + ";")
	if db.Error != nil {
		fmt.Println("Unable to initialise database : ", err)
		return nil, err
	}

	dsn := createDsn()
	db, err = gorm.Open("mysql", dsn)

	if db.AutoMigrate(&Entry{}).Error != nil {
		fmt.Println("Unable to create/modify tables")
		return nil, err
	}

	return db, nil
}
