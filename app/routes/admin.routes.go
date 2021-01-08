package routes

import (
	"github.com/FarukKaradeniz/SpaceHax-server/app/services"
	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(app fiber.Router) {
	r := app.Group("/admin")

	r.Post("/addConfig", services.AddRoomConfig)
	r.Post("/updateConfig", services.UpdateConfig)
	r.Post("/removeConfig", services.RemoveConfig)
	r.Post("/getRoomConfig", services.GetRoomConfig)
	r.Get("/getRoomConfigs", services.GetAllRoomConfigs)
}
