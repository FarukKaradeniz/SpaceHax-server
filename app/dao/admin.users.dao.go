package dao

import "github.com/FarukKaradeniz/SpaceHax-server/config/database"

type User struct {
	Name     string
	Password string
	Role     string
}

func (User) TableName() string {
	return "users"
}

func GetUser(username, password string) (User, error) {
	var user User
	tx := database.DB.Where("name = ? and password = ?", username, password).First(&user)
	return user, tx.Error
}
