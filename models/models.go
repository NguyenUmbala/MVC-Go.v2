package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type Person struct {
	ID       uint   `json:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type Account struct {
	username string `form:"username" json:"username" binding:"required"`
	password string `form:"password" json:"password" binding:"required"`
}

func SignUp(person Person) {
	db.Save(&person)
}

func Login(acc Person) bool {
	var person Person

	if err := db.Where("username = ?", acc.Username).First(&person).Error; err != nil {
		fmt.Println(err)
		return false
	}
	if person.Password != acc.Password {
		return false
	}

	return true
}

// Get database
func GetDB() *gorm.DB {
	return db
}

// Open connection to database
func OpenConnection() {
	db, err = gorm.Open("sqlite3", "./gorm.db")

	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&Person{})
}

// Close connection to database
func CloseConnection() {
	db.Close()
}
