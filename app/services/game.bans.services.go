package services

import (
	"github.com/FarukKaradeniz/SpaceHax-server/app/dao"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func BanPlayer(ctx *fiber.Ctx) error {
	var dto struct {
		PlayerId uint      `json:"playerId"`
		Until    time.Time `json:"until"`
		IsPerma  bool      `json:"isPerma"`
		RoomId   string    `json:"room"`
		BanType  string    `json:"banType"`
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
