package dao

import (
	"github.com/FarukKaradeniz/SpaceHax-server/config/database"
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Name         string `gorm:"column:player_name"`
	Password     string
	Connection   string
	IsAdmin      bool
	IsSuperAdmin bool
}

func CreatePlayer(player *Player) *gorm.DB {
	return database.DB.Create(player)
}

func GetPlayerByNameAndPassword(player *Player, name, password string) *gorm.DB {
	return database.DB.Where("player_name = ? and password = ?", name, password).First(player)
}
