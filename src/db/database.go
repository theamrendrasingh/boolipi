package db

import (
	"fmt"
	"os"

	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// constants for database connection
const (
	//USER : database username
	USER = "amru"
	//PASS database password
	PASS = "q"
	HOST = "127.0.0.1"
	PORT = "3306"

	DBNAME = "boolipi"
)

var database *gorm.DB

type (
	// Entry struct denoting a value, identified by a uuid string, having a name denoted by key and a boolean Value
	Entry struct {
		Uuid  string `gorm:"primaryKey;autoIncrement:false"`
		Value bool   `gorm:"not null;"`
		Key   string
	}

	// User struct denoting a user, identified by a token
	User struct {
		Token string
	}
)

//InitDB to connect to the database, create the database if not exists, create the required tables and initialize the databse connection
func InitDB() error {

	var createDBDsn string

	if _, f := os.LookupEnv("DOCKER_MODE"); f {
		fmt.Println("DOCKER MODE FOUND")
		createDBDsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	} else {
		createDBDsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/", USER, PASS, HOST, PORT)
	}

	fmt.Println(createDBDsn)

	db, err := gorm.Open(mysql.Open(createDBDsn), &gorm.Config{})

	_ = db.Exec("CREATE DATABASE IF NOT EXISTS " + DBNAME + ";")

	var dsn string

	if _, f := os.LookupEnv("DOCKER_MODE"); f {
		fmt.Println("DOCKER MODE FOUND")
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), DBNAME)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	}

	database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	database.AutoMigrate(&Entry{})
	database.AutoMigrate(&User{})

	return err
}

//AddValue adds a new entry to the database
func AddValue(val bool, key string) (Entry, error) {

	u2 := uuid.NewV4().String()
	e := Entry{Uuid: u2, Value: val, Key: key}
	result := database.Model(&Entry{}).Create(&e)

	if result.Error != nil {
		return Entry{}, result.Error
	}

	return Fetch(u2)
}

// Fetch is used to retrieve an entry from the database
func Fetch(id string) (Entry, error) {

	e := Entry{}
	result := database.Where("uuid = ?", id).First(&e)
	return e, result.Error
}

// Delete fucntion deletes the entry from the database with macthing id
func Delete(id string) error {

	result := database.Where("uuid = ? ", id).Delete(&Entry{})
	return result.Error
}

// Patch Function to update the entry in the database with given new values by matching id
func Patch(id string, newval bool, newkey string) (Entry, error) {

	fmt.Println("Patch ::::::::")
	fmt.Println(id)
	fmt.Println(newval)
	fmt.Println(newkey)

	// This update statement does not work properly. It cannot update a value to false as it is '0' (empty), and gorm does not update empty values
	// result := database.Model(&Entry{}).Where("uuid = ?", id).Updates(Entry{Uuid: id, Value: newval, Key: newkey})

	// Hence we need to update it after creating a string map
	result := database.Model(&Entry{}).Where("uuid = ?", id).Updates(map[string]interface{}{"Uuid": id, "Value": newval, "Key": newkey})

	if result.Error != nil {
		return Entry{}, result.Error
	}

	return Fetch(id)
}
