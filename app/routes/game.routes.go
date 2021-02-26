package routes

import (
	"github.com/FarukKaradeniz/SpaceHax-server/app/services"
	"github.com/FarukKaradeniz/SpaceHax-server/middleware"
	"github.com/gofiber/fiber/v2"
)

func HaxGameRoutes(app fiber.Router) {
	r := app.Group("/game")

	stats := r.Group("/stats")
	stats.Post("", middleware.Protected([]string{"room"}), services.SaveGame)
	stats.Delete("/:playerId", middleware.Protected([]string{"admin", "room"}), services.ClearPlayerStats)
	stats.Get("/:playerName", middleware.Protected([]string{"admin", "room"}), services.GetStats)
	// TODO topGoals, topAssists, topPoints

	bans := r.Group("/bans")
	bans.Post("", middleware.Protected([]string{"admin", "room"}), services.BanPlayer)
	bans.Get("/:room", middleware.Protected([]string{"admin", "room"}), services.GetBanList)
	bans.Delete("/:playerId", middleware.Protected([]string{"admin", "room"}), services.ClearBan)
}
