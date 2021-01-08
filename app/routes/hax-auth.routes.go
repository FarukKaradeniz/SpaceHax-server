package routes

import (
	"github.com/FarukKaradeniz/SpaceHax-server/app/services"
	"github.com/gofiber/fiber/v2"
)

func HaxAuthRoutes(app fiber.Router) {
	r := app.Group("/auth")

	r.Post("/signup", services.SignUp)
	r.Post("/login", services.Login)
	r.Post("/changePassword", services.ChangePassword)
	// TODO şifre unutmalar için passwordRequest gibi bişey oluştur
}
