package dao

import (
	"github.com/FarukKaradeniz/SpaceHax-server/app/models"
	"github.com/FarukKaradeniz/SpaceHax-server/config/database"
	"gorm.io/gorm"
)

func AddRoomConfig(alias, name, mapType string, scoreLimit, timeLimit int8) *gorm.DB {
	config := &models.RoomConfig{
		Alias:      alias,
		Map:        mapType,
		RoomName:   name,
		ScoreLimit: scoreLimit,
		TimeLimit:  timeLimit,
	}
	return database.DB.Create(config)
}

func GetRoomConfig(alias string) (models.RoomConfig, error) {
	var config models.RoomConfig
	tx := database.DB.Where("alias = ?", alias).First(&config)
	return config, tx.Error
}

func GetAllRoomConfigs() ([]models.RoomConfig, error) {
	var configs []models.RoomConfig
	tx := database.DB.Find(&configs)
	return configs, tx.Error
}

func UpdateConfig(config *models.RoomConfig) *gorm.DB {
	return database.DB.Updates(config)
}

func RemoveConfig(id int) *gorm.DB {
	return database.DB.Where("id = ?", id).Delete(&models.RoomConfig{})
}
