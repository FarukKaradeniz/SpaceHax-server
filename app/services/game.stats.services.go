package services

import (
	"github.com/FarukKaradeniz/SpaceHax-server/app/dao"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type GameStatsDTO struct {
	Played []uint         `json:"played"`
	Stats  map[uint]Stats `json:"stats"`
	RoomId string         `json:"room"`
}

type Stats struct {
	GoalsCount   uint `json:"goalsCount"`
	AssistsCount uint `json:"assistsCount"`
	Won          uint `json:"won"`
}

func SaveGame(ctx *fiber.Ctx) error {
	dto := new(GameStatsDTO)
	if err := ctx.BodyParser(dto); err != nil {
		return err
	}

	for _, playerId := range dto.Played {
		err := dao.UpdatePlayerStats(playerId, dto.Stats[playerId].GoalsCount, dto.Stats[playerId].AssistsCount, dto.Stats[playerId].Won, dto.RoomId).Error
		if err != nil {
			return fiber.NewError(fiber.StatusConflict, "error updating stats")
		}
	}
	return ctx.SendStatus(fiber.StatusOK)
}

func ClearPlayerStats(ctx *fiber.Ctx) error {
	playerId, err := strconv.Atoi(ctx.Params("playerId"))
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "invalid payload")
	}

	err = dao.ClearPlayerStats(uint(playerId), ctx.Query("room")).Error
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "error clearing stats")
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func GetStats(ctx *fiber.Ctx) error {
	playerId := ctx.Params("playerName")
	playerStats, err := dao.GetPlayerStatsByID(playerId, ctx.Query("room"))

	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "error getting stats")
	}

	return ctx.JSON(playerStats)
}
