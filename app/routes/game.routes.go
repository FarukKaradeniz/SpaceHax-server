package routes

import (
	"github.com/FarukKaradeniz/SpaceHax-server/app/services"
	"github.com/gofiber/fiber/v2"
)

func HaxGameRoutes(app fiber.Router) {
	r := app.Group("/game")

	stats := r.Group("/stats")
	stats.Post("", services.SaveGame)
	stats.Delete("/:playerId", services.ClearPlayerStats)
	stats.Get("/:playerId?", services.GetStats)
	// TODO topGoals, topAssists, topPoints

	bans := r.Group("/bans")
	bans.Post("", services.BanPlayer)
	bans.Get("/:room", services.GetBanList)
	bans.Delete("/:playerId", services.ClearBan)
}
