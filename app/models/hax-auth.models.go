package models

type LoginDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	RoomId   string `json:"room"`
}

type SignUpDTO struct {
	LoginDTO
	Connection string `json:"conn"`
}

type AuthResponse struct {
	Message      string `json:"message"`
	IsAdmin      *bool  `json:"is_admin,omitempty"`
	IsSuperAdmin *bool  `json:"is_super_admin,omitempty"`
	PlayerId     uint   `json:"player_id,omitempty"`
}
