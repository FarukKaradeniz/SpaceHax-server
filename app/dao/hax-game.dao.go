package dao

import (
	"errors"
	"github.com/FarukKaradeniz/SpaceHax-server/config/database"
	"gorm.io/gorm"
	"time"
)

// Sonradan OG count, fastest goal, pointswon eklenebilir

type PlayerStats struct {
	ID                uint `gorm:"primarykey"`
	UpdatedAt         time.Time
	PlayerId          uint
	TotalGoalsCount   uint
	TotalAssistsCount uint
	TotalGamesPlayed  uint
	TotalGamesWon     uint
	GoalsCount        uint
	AssistsCount      uint
	GamesPlayed       uint
	GamesWon          uint
	RoomId            string
}

func (PlayerStats) TableName() string {
	return "player_stats"
}

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

func createPlayerStats(stats *PlayerStats) *gorm.DB {
	return database.DB.Create(stats)
}

func UpdatePlayerStats(playerId, newGoals, newAssists, won uint, roomId string) *gorm.DB {
	updates := map[string]interface{}{
		"total_goals_count":   gorm.Expr("total_goals_count + ?", newGoals),
		"total_assists_count": gorm.Expr("total_assists_count + ?", newAssists),
		"total_games_played":  gorm.Expr("total_games_played + ?", 1),
		"total_games_won":     gorm.Expr("total_games_won + ?", won),
		"goals_count":         gorm.Expr("goals_count + ?", newGoals),
		"assists_count":       gorm.Expr("assists_count + ?", newAssists),
		"games_played":        gorm.Expr("games_played + ?", 1),
		"games_won":           gorm.Expr("games_won + ?", won),
	}
	return database.DB.Model(&PlayerStats{}).Where("player_id = ? and room_id = ?", playerId, roomId).Updates(updates)
}

func ClearPlayerStats(playerId uint, roomId string) *gorm.DB {
	updates := map[string]interface{}{
		"goals_count":   0,
		"assists_count": 0,
		"games_played":  0,
		"games_won":     0,
	}
	return database.DB.Model(&PlayerStats{}).Where("player_id = ? and room_id = ?", playerId, roomId).Updates(updates)
}

func GetPlayerStatsByID(playerId uint, roomId string) (PlayerStats, error) {
	stats := PlayerStats{}
	tx := database.DB.Where("player_id = ? and room_id = ?", playerId, roomId).First(&stats)
	if tx.Error != nil {
		return stats, tx.Error
	}
	if stats.PlayerId == playerId {
		return stats, nil
	}
	return stats, errors.New("error getting stats")
}

func GetPlayers(limit int, sortBy, roomId string) ([]PlayerStats, error) {
	sort := map[string]string{
		"goals":   "total_goals_count",
		"assists": "total_assists_count",
	}
	var stats []PlayerStats
	tx := database.DB.Where("room_id = ?", roomId).Order(sort[sortBy] + " desc").Limit(limit).Find(&stats)
	return stats, tx.Error
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

func RemoveStats(playerId uint, roomId string) *gorm.DB {
	return database.DB.Where("player_id = ? and room_id = ?", playerId, roomId).Delete(&PlayerStats{})
}
