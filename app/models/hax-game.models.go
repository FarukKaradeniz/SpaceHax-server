package models

import "time"

type GameStatsDTO struct {
	Played []uint         `json:"played"`
	Stats  map[uint]Stats `json:"stats"`
	RoomId string         `json:"room"`
}

type Stats struct {
	GoalsCount   uint `json:"goalsCount"`
	AssistsCount uint `json:"assistsCount"`
	Won          uint `json:"won"`
}

type GameResponse struct {
	Message string `json:"message"`
}

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

// Sonradan OG count, fastest goal, pointswon eklenebilir

type BannedPlayer struct {
	ID          uint `gorm:"primarykey"`
	BannedUntil time.Time
	PlayerId    uint
	IsPerma     bool
	RoomId      string
	Type        string
}

func (PlayerStats) TableName() string {
	return "player_stats"
}

func (BannedPlayer) TableName() string {
	return "banned_players"
}
