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
	RoomId       string
}

func (Player) TableName() string {
	return "players"
}

func CreatePlayer(player *Player) *gorm.DB {
	return database.DB.Create(player)
}

func GetPlayerByNameAndPassword(player *Player, name, password string, roomId string) *gorm.DB {
	return database.DB.Where("player_name = ? and password = ? and room_id = ?", name, password, roomId).First(player)
}

func ChangePassword(playerId uint, newPassword string, roomId string) *gorm.DB {
	return database.DB.Model(&Player{}).Where("id = ? and room_id = ?", playerId, roomId).Update("password", newPassword)
}

func RemovePlayers(roomId string) *gorm.DB {
	return database.DB.Where("room_id = ?", roomId).Delete(&Player{})
}

func (player *Player) AfterSave(tx *gorm.DB) (err error) {
	stats := PlayerStats{PlayerId: player.ID, RoomId: player.RoomId}
	if err := createPlayerStats(&stats); err.Error != nil {
		return err.Error
	}
	return nil
}

func (player *Player) BeforeDelete(tx *gorm.DB) (err error) {
	err = ClearBan(player.ID, player.RoomId).Error
	if err != nil {
		return err
	}

	return RemoveStats(player.ID, player.RoomId).Error
}
