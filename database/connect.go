package database

import (
	"fmt"
	"log"
	"os"

	"github.com/ariefro/go-note-rest-api/config"
	"github.com/ariefro/go-note-rest-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s TimeZone=Asia/Shanghai", config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("DB_USER"), config.Config("DB_NAME"), config.Config("DB_PASSWORD"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database. \n", err)
		os.Exit(2)
	}

	fmt.Println("Connection Opened to Database")

	DB.AutoMigrate(&model.Note{}, &model.User{})
	fmt.Println("Database Migrated")
}
