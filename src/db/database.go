package db

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	USER   = "amru"
	PASS   = "q"
	HOST   = "127.0.0.1"
	PORT   = "3306"
	DBNAME = "boolipi"
)

var database *gorm.DB

type (
	Entry struct {
		// gorm.Model
		Uuid  string `gorm:"primaryKey;autoIncrement:false"`
		Value bool   `gorm:"not null;"`
		Key   string
	}

	User struct {
		Token string
	}
)

func InitDB() error {
	createDBDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", USER, PASS, HOST, PORT)
	db, err := gorm.Open(mysql.Open(createDBDsn), &gorm.Config{})

	_ = db.Exec("CREATE DATABASE IF NOT EXISTS " + DBNAME + ";")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	database.AutoMigrate(&Entry{})
	database.AutoMigrate(&User{})

	return err
}

func AddValue(val bool, key string) (Entry, error) {

	u2 := uuid.NewV4().String()
	e := Entry{Uuid: u2, Value: val, Key: key}
	result := database.Model(&Entry{}).Create(&e)

	if result.Error != nil {
		return Entry{}, result.Error
	}

	return Fetch(u2)
}

func Fetch(id string) (Entry, error) {

	e := Entry{}
	result := database.Where("uuid = ?", id).First(&e)
	return e, result.Error
}

func Delete(id string) error {

	result := database.Where("uuid = ? ", id).Delete(&Entry{})
	// fmt.Println(result.Error)
	return result.Error
}

func Patch(id string, val bool, key string) (Entry, error) {

	result := database.Model(&Entry{}).Where("uuid = ?", id).Updates(Entry{Uuid: id, Value: val, Key: key})
	if result.Error != nil {
		return Entry{}, result.Error
	}
	return Fetch(id)
}
