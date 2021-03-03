package routes

import (
	"github.com/FarukKaradeniz/SpaceHax-server/app/services"
	"github.com/FarukKaradeniz/SpaceHax-server/middleware"
	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(app fiber.Router) {
	r := app.Group("/admin")

	r.Post("/login", services.UserLogin)

	config := r.Group("/configs")
	config.Post("", middleware.Protected([]string{"admin"}), services.AddRoomConfig)
	config.Put("/:alias", middleware.Protected([]string{"admin", "room"}), services.UpdateConfig)
	config.Delete("/:alias", middleware.Protected([]string{"admin"}), services.RemoveConfig)
	config.Get("/:alias?", middleware.Protected([]string{"room", "admin"}), services.GetRoomConfig)
}
