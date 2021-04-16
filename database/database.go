package database

import (
	"fmt"

	"github.com/diogosonsim/crud-go-test/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Create database access and migrate the tables
func Connect() {
	database, err := gorm.Open(mysql.Open("root:password@/go-test"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database!")
	}

	DB = database

	database.AutoMigrate(&models.User{})

	fmt.Println("Database connected!")
}
