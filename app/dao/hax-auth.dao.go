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

func (Player) TableName() string {
	return "players"
}

func CreatePlayer(player *Player) *gorm.DB {
	return database.DB.Create(player)
}

func GetPlayerByNameAndPassword(player *Player, name, password string) *gorm.DB {
	return database.DB.Where("player_name = ? and password = ?", name, password).First(player)
}

func ChangePassword(playerId uint, newPassword string) *gorm.DB {
	return database.DB.Model(&Player{}).Where("id = ?", playerId).Update("password", newPassword)
}

func (player *Player) AfterSave(tx *gorm.DB) (err error) {
	stats := PlayerStats{PlayerId: player.ID}
	if err := createPlayerStats(&stats); err.Error != nil {
		return err.Error
	}
	return nil
}
