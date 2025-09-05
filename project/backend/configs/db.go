package configs

import (
	"fmt"

	"G13.com/backend/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("sa.db?cache=shared"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	fmt.Println("connected database")
	db = database
}
func SetupDatabase() {
	db.AutoMigrate(
		&entity.Categories{},
		&entity.User{},
		&entity.Game{},
		&entity.KeyGame{},
		&entity.Market{},
		&entity.MinimumSpec{},
		&entity.Request{},
		&entity.Role{})

	//Categories
	db.Model(&entity.Categories{}).Create(&entity.Categories{
		Title: "FPS",
	})

	db.Model(&entity.Categories{}).Create(&entity.Categories{
		Title: "Horror",
	})

	db.Model(&entity.Categories{}).Create(&entity.Categories{
		Title: "TPS",
	})

	//Role
	db.Model(&entity.Role{}).Create(&entity.Role{
		Title:       "User",
		Description: "bet",
	})

	db.Model(&entity.Role{}).Create(&entity.Role{
		Title:       "Admin",
		Description: "abby",
	})
}
