package services

import (
	"github.com/FarukKaradeniz/SpaceHax-server/app/dao"
	"github.com/gofiber/fiber/v2"
)

func AddRoomConfig(ctx *fiber.Ctx) error {
	dto := new(dao.RoomConfig)
	if err := ctx.BodyParser(dto); err != nil {
		return err
	}

	err := dao.AddRoomConfig(dto.Alias, dto.RoomName, dto.Map, dto.ScoreLimit, dto.TimeLimit, dto.MaxPlayer).Error
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "error creating room config")
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func GetRoomConfig(ctx *fiber.Ctx) error {
	alias := ctx.Params("alias")
	var config interface{}
	var err error

	if alias != "" {
		config, err = dao.GetRoomConfig(alias)
	} else {
		config, err = dao.GetAllRoomConfigs()
	}

	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "error getting room config")
	}

	return ctx.JSON(config)
}

func UpdateConfig(ctx *fiber.Ctx) error {
	dto := new(dao.RoomConfig)
	if err := ctx.BodyParser(dto); err != nil {
		return err
	}

	if dto.Alias != ctx.Params("alias") {
		return fiber.NewError(fiber.StatusConflict, "invalid payload")
	}

	if err := dao.UpdateConfig(dto).Error; err != nil {
		return fiber.NewError(fiber.StatusConflict, "error updating room config")
	}

	return ctx.JSON(dto)
}

func RemoveConfig(ctx *fiber.Ctx) error {
	alias := ctx.Params("alias")
	if err := dao.RemoveConfig(alias); err != nil {
		return fiber.NewError(fiber.StatusConflict, "error removing room config")
	}

	return ctx.SendStatus(fiber.StatusOK)
}
