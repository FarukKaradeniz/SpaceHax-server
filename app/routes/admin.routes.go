package routes

import (
	"github.com/FarukKaradeniz/SpaceHax-server/app/services"
	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(app fiber.Router) {
	r := app.Group("/admin")

	config := r.Group("/configs")
	config.Post("", services.AddRoomConfig)
	config.Put("/:alias", services.UpdateConfig)
	config.Delete("/:alias", services.RemoveConfig) // TODO will be tested
	config.Get("/:alias?", services.GetRoomConfig)
}
