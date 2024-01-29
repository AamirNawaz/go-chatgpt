package config

import (
	"go-chatgpt-app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=aamirnawaz password=aamirnawaz dbname=go-chatgpt-app-db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("connection failed to the database with postgres")
	}

	DB = db

	// CreateTableSchema(db)
}

func CreateTableSchema(connection *gorm.DB) {
	// connection.Debug().Migrator().DropTable(&models.User{}, &models.SearchHistory{})
	connection.Debug().AutoMigrate(&models.User{}, &models.SearchHistory{})
}
