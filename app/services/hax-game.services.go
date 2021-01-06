package services

import (
	"github.com/FarukKaradeniz/SpaceHax-server/app/dao"
	"github.com/FarukKaradeniz/SpaceHax-server/app/models"
	"github.com/gofiber/fiber/v2"
)

func SaveGame(ctx *fiber.Ctx) error {
	dto := new(models.GameStatsDTO)
	if err := ctx.BodyParser(dto); err != nil {
		return err
	}

	for _, playerId := range dto.Played {
		err := dao.UpdatePlayerStats(playerId, dto.Stats[playerId].GoalsCount, dto.Stats[playerId].AssistsCount, dto.Stats[playerId].Won).Error
		if err != nil {
			return fiber.NewError(fiber.StatusConflict, "error updating stats")
		}
	}
	return ctx.JSON(models.GameResponse{
		Message: "success",
	})
}

func ClearPlayerStats(ctx *fiber.Ctx) error {
	var dto struct {
		PlayerId uint `json:"playerId"`
	}
	if err := ctx.BodyParser(&dto); err != nil {
		return err
	}

	err := dao.ClearPlayerStats(dto.PlayerId).Error
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "error clearing stats")
	}

	return ctx.JSON(models.GameResponse{
		Message: "success",
	})
}

func GetStats(ctx *fiber.Ctx) error {
	var dto struct {
		PlayerId uint `json:"playerId"`
	}
	if err := ctx.BodyParser(&dto); err != nil {
		return err
	}

	playerStats, err := dao.GetPlayerStatsByID(dto.PlayerId)
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "error getting stats")
	}

	return ctx.JSON(playerStats)
}

func GetTop5PlayersByGoals(ctx *fiber.Ctx) error {
	players, err := dao.GetTop5PlayersByGoals()
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "error getting stats")
	}

	return ctx.JSON(players)
}

func GetTop5PlayersByAssists(ctx *fiber.Ctx) error {
	players, err := dao.GetTop5PlayersByAssists()
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "error getting stats")
	}

	return ctx.JSON(players)
}

func BanPlayer(ctx *fiber.Ctx) error {
	return ctx.SendString("this is from ban player")
}

func GetBanList(ctx *fiber.Ctx) error {
	return ctx.SendString("this is from get ban list")
}
