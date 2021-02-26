package routes

import (
	"github.com/FarukKaradeniz/SpaceHax-server/app/services"
	"github.com/FarukKaradeniz/SpaceHax-server/middleware"
	"github.com/gofiber/fiber/v2"
)

func HaxAuthRoutes(app fiber.Router) {
	r := app.Group("/auth")

	r.Post("/signup", middleware.Protected([]string{"room"}), services.SignUp)
	r.Post("/login", middleware.Protected([]string{"room"}), services.Login)
	r.Post("/changePassword", middleware.Protected([]string{"admin", "room"}), services.ChangePassword)
	// TODO şifre unutmalar için passwordRequest gibi bişey oluştur
}
