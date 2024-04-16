package db

import (
	"johanbx/go-web-server-2/internal/catalog"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dataSourceName string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&catalog.CatalogItem{})
}
