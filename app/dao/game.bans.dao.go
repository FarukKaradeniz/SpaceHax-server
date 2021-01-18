package dao

import (
	"github.com/FarukKaradeniz/SpaceHax-server/config/database"
	"gorm.io/gorm"
	"time"
)

type BannedPlayer struct {
	ID          uint `gorm:"primarykey"`
	BannedUntil time.Time
	PlayerId    uint
	IsPerma     bool
	RoomId      string
	Type        string
}

func (BannedPlayer) TableName() string {
	return "banned_players"
}

func GetBanList(roomId string) ([]BannedPlayer, error) {
	var list []BannedPlayer
	tx := database.DB.Where("room_id = ?", roomId).Order("id asc").Find(&list)
	return list, tx.Error
}

func BanPlayer(playerId uint, isPerma bool, until time.Time, roomId string, banType string) *gorm.DB {
	bannedPlayer := &BannedPlayer{
		BannedUntil: until,
		PlayerId:    playerId,
		IsPerma:     isPerma,
		RoomId:      roomId,
		Type:        banType,
	}
	return database.DB.Create(bannedPlayer)
}

func ClearBan(playerId uint, roomId string) *gorm.DB {
	return database.DB.Where("player_id = ? and room_id = ?", playerId, roomId).Delete(&BannedPlayer{})
}

func ClearBans(roomId string) *gorm.DB {
	return database.DB.Where("room_id = ?", roomId).Delete(&BannedPlayer{})
}
