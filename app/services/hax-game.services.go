package services

import (
	"github.com/FarukKaradeniz/SpaceHax-server/app/dao"
	"github.com/FarukKaradeniz/SpaceHax-server/app/models"
	"github.com/gofiber/fiber/v2"
	"time"
)

func SaveGame(ctx *fiber.Ctx) error {
	dto := new(models.GameStatsDTO)
	if err := ctx.BodyParser(dto); err != nil {
		return err
	}

	for _, playerId := range dto.Played {
		err := dao.UpdatePlayerStats(playerId, dto.Stats[playerId].GoalsCount, dto.Stats[playerId].AssistsCount, dto.Stats[playerId].Won, dto.RoomId).Error
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
		PlayerId uint   `json:"playerId"`
		RoomId   string `json:"room"`
	}
	if err := ctx.BodyParser(&dto); err != nil {
		return err
	}

	err := dao.ClearPlayerStats(dto.PlayerId, dto.RoomId).Error
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "error clearing stats")
	}

	return ctx.JSON(models.GameResponse{
		Message: "success",
	})
}

func GetStats(ctx *fiber.Ctx) error {
	var dto struct {
		PlayerId uint   `json:"playerId"`
		RoomId   string `json:"room"`
	}
	if err := ctx.BodyParser(&dto); err != nil {
		return err
	}

	playerStats, err := dao.GetPlayerStatsByID(dto.PlayerId, dto.RoomId)
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "error getting stats")
	}

	return ctx.JSON(playerStats)
}

func GetTop5PlayersByGoals(ctx *fiber.Ctx) error {
	players, err := dao.GetTop5PlayersByGoals(ctx.Params("room"))
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "error getting stats")
	}

	return ctx.JSON(players)
}

func GetTop5PlayersByAssists(ctx *fiber.Ctx) error {
	players, err := dao.GetTop5PlayersByAssists(ctx.Params("room"))
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "error getting stats")
	}

	return ctx.JSON(players)
}

func BanPlayer(ctx *fiber.Ctx) error {
	var dto struct {
		PlayerId uint      `json:"playerId"`
		Until    time.Time `json:"until"`
		IsPerma  bool      `json:"is_perma"`
		RoomId   string    `json:"room"`
		BanType  string    `json:"ban_type"`
	}
	if err := ctx.BodyParser(&dto); err != nil {
		return err
	}

	if err := dao.BanPlayer(dto.PlayerId, dto.IsPerma, dto.Until, dto.RoomId, dto.BanType).Error; err != nil {
		return fiber.NewError(fiber.StatusConflict, "error banning player")
	}

	return ctx.JSON(models.GameResponse{
		Message: "success",
	})
}

func GetBanList(ctx *fiber.Ctx) error {
	banList, err := dao.GetBanList(ctx.Params("room"))
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "error getting banned players list")
	}

	return ctx.JSON(banList)
}

func ClearBan(ctx *fiber.Ctx) error {
	var dto struct {
		PlayerId uint   `json:"playerId"`
		RoomId   string `json:"room"`
	}
	if err := ctx.BodyParser(&dto); err != nil {
		return err
	}

	if err := dao.ClearBan(dto.PlayerId, dto.RoomId).Error; err != nil {
		return fiber.NewError(fiber.StatusConflict, "error clearing ban")
	}

	return ctx.JSON(models.GameResponse{
		Message: "success",
	})
}
