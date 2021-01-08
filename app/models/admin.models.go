package models

type RoomConfig struct {
	ID         uint   `gorm:"primarykey" json:"id"`
	Map        string `json:"map"`
	RoomName   string `json:"roomName"`
	ScoreLimit int8   `json:"scoreLimit"`
	TimeLimit  int8   `json:"timeLimit"`
	Alias      string `json:"alias"` // spacebouncebrakesv3, spacebouncev4 gibi
}

func (RoomConfig) TableName() string {
	return "room_config"
}
