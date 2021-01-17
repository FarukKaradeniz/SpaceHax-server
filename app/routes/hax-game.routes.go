package routes

import (
	"github.com/FarukKaradeniz/SpaceHax-server/app/services"
	"github.com/gofiber/fiber/v2"
)

func HaxGameRoutes(app fiber.Router) {
	r := app.Group("/game")

	r.Post("/clear", services.ClearPlayerStats)
	r.Post("/save", services.SaveGame)
	r.Get("/stats/:username", services.GetStats)
	r.Get("/top5byGoals/:room", services.GetTop5PlayersByGoals)
	r.Get("/top5byAssists/:room", services.GetTop5PlayersByAssists)
	r.Get("/banList/:room", services.GetBanList)
	r.Post("/banPlayer", services.BanPlayer)
	r.Post("/clearBan", services.ClearBan)
}
