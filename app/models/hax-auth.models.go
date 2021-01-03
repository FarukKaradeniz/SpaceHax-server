package models

type LoginDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SignUpDTO struct {
	LoginDTO
	Connection string `json:"conn"`
}

type AuthResponse struct {
	Message string `json:"message"`
}

// TODO admin/süper admin durumunu döndürebilirsin
