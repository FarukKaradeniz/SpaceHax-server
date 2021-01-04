package database

import (
	"fmt"
	"github.com/FarukKaradeniz/SpaceHax-server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dbUri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s",
		config.DB_HOST, config.DB_USER, config.DB_PASS, config.DB_NAME)
	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		fmt.Println("database connection error")
		panic(err)
	}

	DB = db
	fmt.Println("database connected")
}
