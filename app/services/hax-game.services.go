package services

import "github.com/gofiber/fiber/v2"

func SaveGame(ctx *fiber.Ctx) error {
	return ctx.SendString("this is from save game")
}

func GetStats(ctx *fiber.Ctx) error {
	return ctx.SendString("this is from get stats")
}

func BanPlayer(ctx *fiber.Ctx) error {
	return ctx.SendString("this is from ban player")
}

func GetBanList(ctx *fiber.Ctx) error {
	return ctx.SendString("this is from get ban list")
}
