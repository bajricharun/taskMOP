package database

import (
	"github.com/bajricharun/taskMOP/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:harun123@/127.0.0.1:3306"), &gorm.Config{})
	if err != nil {
		panic("could not connect to db")
	}

	DB = connection
	connection.AutoMigrate(&models.User{})
}
