package database

import (
	"github.com/jitendraks2/auth-react-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	// connection, err := gorm.Open(mysql.Open(dsn: ""), &gorm.Config{})
	// _, err := gorm.Open(mysql.Open(dsn: ""), &gorm.Config{})
	dsn := "Jitendra:jitendraaa8990@tcp(127.0.0.1:3306)/learnauth?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to db")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
