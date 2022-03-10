package database

import (
	"fmt"

	"github.com/ariefro/notes-server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s TimeZone=Asia/Shanghai", config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("DB_USER"), config.Config("DB_NAME"), config.Config("DB_PASSWORD"))

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
}