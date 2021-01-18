package services

import (
	"github.com/FarukKaradeniz/SpaceHax-server/app/dao"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
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
	playerId, err := strconv.Atoi(ctx.Params("playerId"))
	var playerStats interface{}
	if playerId != 0 {
		playerStats, err = dao.GetPlayerStatsByID(uint(playerId), ctx.Query("room"))
	} else {
		limit, _ := strconv.Atoi(ctx.Query("limit", "5"))
		room := ctx.Query("room")
		sortBy := ctx.Query("sortBy", "goals")
		playerStats, err = dao.GetPlayers(limit, sortBy, room)
	}

	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "error getting stats")
	}

	return ctx.JSON(playerStats)
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

	return ctx.SendStatus(fiber.StatusOK)
}

func GetBanList(ctx *fiber.Ctx) error {
	banList, err := dao.GetBanList(ctx.Params("room"))
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "error getting banned players list")
	}

	return ctx.JSON(banList)
}

func ClearBan(ctx *fiber.Ctx) error {
	playerId, err := strconv.Atoi(ctx.Params("playerId"))
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "invalid payload")
	}

	if err := dao.ClearBan(uint(playerId), ctx.Query("room")).Error; err != nil {
		return fiber.NewError(fiber.StatusConflict, "error clearing ban")
	}

	return ctx.SendStatus(fiber.StatusOK)
}
