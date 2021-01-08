package dao

import (
	"errors"
	"github.com/FarukKaradeniz/SpaceHax-server/app/models"
	"github.com/FarukKaradeniz/SpaceHax-server/config/database"
	"gorm.io/gorm"
	"time"
)

func createPlayerStats(stats *models.PlayerStats) *gorm.DB {
	return database.DB.Create(stats)
}

func UpdatePlayerStats(playerId, newGoals, newAssists, won uint) *gorm.DB {
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
	return database.DB.Model(&models.PlayerStats{}).Where("player_id = ?", playerId).Updates(updates)
}

func ClearPlayerStats(playerId uint) *gorm.DB {
	updates := map[string]interface{}{
		"goals_count":   0,
		"assists_count": 0,
		"games_played":  0,
		"games_won":     0,
	}
	return database.DB.Model(&models.PlayerStats{}).Where("player_id = ?", playerId).Updates(updates)
}

func GetPlayerStatsByID(playerId uint) (models.PlayerStats, error) {
	stats := models.PlayerStats{ID: playerId}
	tx := database.DB.First(&stats)
	if tx.Error != nil {
		return stats, tx.Error
	}
	if stats.ID == playerId {
		return stats, nil
	}
	return stats, errors.New("error getting stats")
}

func GetTop5PlayersByGoals() ([]models.PlayerStats, error) {
	var stats []models.PlayerStats
	tx := database.DB.Order("total_goals_count desc").Find(&stats).Limit(5)
	return stats, tx.Error
}

func GetTop5PlayersByAssists() ([]models.PlayerStats, error) {
	var stats []models.PlayerStats
	tx := database.DB.Order("total_assists_count desc").Find(&stats).Limit(5)
	return stats, tx.Error
}

func GetBanList() ([]models.BannedPlayer, error) {
	var list []models.BannedPlayer
	tx := database.DB.Order("id asc").Find(&list)
	return list, tx.Error
}

func BanPlayer(playerId uint, isPerma bool, until time.Time) *gorm.DB {
	bannedPlayer := &models.BannedPlayer{
		BannedUntil: until,
		PlayerId:    playerId,
		IsPerma:     isPerma,
	}
	return database.DB.Create(bannedPlayer)
}

func ClearBan(playerId uint) *gorm.DB {
	return database.DB.Where("player_id = ?", playerId).Delete(&models.BannedPlayer{})
}
