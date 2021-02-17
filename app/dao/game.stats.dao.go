package dao

import (
	"errors"
	"github.com/FarukKaradeniz/SpaceHax-server/config/database"
	"gorm.io/gorm"
	"net/url"
	"strings"
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
	PlayerName        string `gorm:"->"`
}

func (PlayerStats) TableName() string {
	return "player_stats"
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

func GetPlayerStatsByID(playerName, roomId string) (PlayerStats, error) {
	stats := PlayerStats{}
	playerName, err := url.QueryUnescape(playerName)
	if err != nil || strings.ReplaceAll(playerName, " ", "") == "" {
		return stats, errors.New("problem decoding")
	}
	tx := database.DB.Model(&PlayerStats{}).Select("player_stats.id, player_stats.player_id, player_stats.total_goals_count,"+
		"player_stats.total_assists_count, player_stats.total_games_won, player_stats.goals_count, player_stats.assists_count,"+
		"player_stats.games_played, player_stats.games_won, player_stats.room_id, players.player_name").Joins("inner join players on "+
		"player_stats.player_id = players.id").Where("players.player_name = ? and player_stats.room_id = ?", playerName, roomId).Scan(&stats)

	if tx.Error != nil {
		return stats, tx.Error
	}
	if stats.ID == 0 {
		return stats, errors.New("stats not found")
	}

	return stats, nil
}

// TODO needs fixing and another route
func GetPlayers(limit int, sortBy, roomId string) ([]PlayerStats, error) {
	sort := map[string]string{
		"goals":   "total_goals_count",
		"assists": "total_assists_count",
	}
	var stats []PlayerStats
	tx := database.DB.Where("room_id = ?", roomId).Order(sort[sortBy] + " desc").Limit(limit).Find(&stats)
	return stats, tx.Error
}

func RemoveStats(roomId string) *gorm.DB {
	return database.DB.Where("room_id = ?", roomId).Delete(&PlayerStats{})
}
