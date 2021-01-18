package models

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
