package dao

import (
	"github.com/FarukKaradeniz/SpaceHax-server/config/database"
	"gorm.io/gorm"
)

type RoomConfig struct {
	Map        string `json:"map"`
	RoomName   string `json:"roomName"`
	ScoreLimit int8   `json:"scoreLimit"`
	TimeLimit  int8   `json:"timeLimit"`
	Alias      string `json:"alias"` // spacebouncebrakesv3, spacebouncev4 gibi
}

func (RoomConfig) TableName() string {
	return "room_config"
}

func AddRoomConfig(alias, name, mapType string, scoreLimit, timeLimit int8) *gorm.DB {
	config := &RoomConfig{
		Alias:      alias,
		Map:        mapType,
		RoomName:   name,
		ScoreLimit: scoreLimit,
		TimeLimit:  timeLimit,
	}
	return database.DB.Create(config)
}

func GetRoomConfig(alias string) (RoomConfig, error) {
	var config RoomConfig
	tx := database.DB.Where("alias = ?", alias).First(&config)
	return config, tx.Error
}

func GetAllRoomConfigs() ([]RoomConfig, error) {
	var configs []RoomConfig
	tx := database.DB.Find(&configs)
	return configs, tx.Error
}

func UpdateConfig(config *RoomConfig) *gorm.DB {
	return database.DB.Where("alias = ?", config.Alias).Updates(config)
}

func RemoveConfig(alias string) *gorm.DB {
	return database.DB.Where("alias = ?", alias).Delete(&RoomConfig{Alias: alias})
}

func (config *RoomConfig) BeforeDelete(tx *gorm.DB) (err error) {
	return RemovePlayers(config.Alias).Error
}
