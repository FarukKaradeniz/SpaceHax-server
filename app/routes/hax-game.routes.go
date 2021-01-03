package routes

import (
	"github.com/FarukKaradeniz/SpaceHax-server/app/services"
	"github.com/gofiber/fiber/v2"
)

func HaxGameRoutes(app fiber.Router) {
	r := app.Group("/game")

	r.Post("/save", services.SaveGame)
	r.Get("/stats/:username", services.GetStats)
	r.Get("/banList", services.GetBanList)
	r.Post("/banPlayer", services.BanPlayer)
}
