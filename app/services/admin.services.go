package services

import (
	"github.com/FarukKaradeniz/SpaceHax-server/app/dao"
	"github.com/FarukKaradeniz/SpaceHax-server/app/models"
	"github.com/gofiber/fiber/v2"
)

func AddRoomConfig(ctx *fiber.Ctx) error {
	dto := new(models.RoomConfig)
	if err := ctx.BodyParser(dto); err != nil {
		return err
	}

	err := dao.AddRoomConfig(dto.Alias, dto.RoomName, dto.Map, dto.ScoreLimit, dto.TimeLimit).Error
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "error creating room config")
	}

	return ctx.JSON(models.GameResponse{
		Message: "success",
	})
}

func GetRoomConfig(ctx *fiber.Ctx) error {
	var dto struct {
		Alias string `json:"alias"`
	}
	if err := ctx.BodyParser(&dto); err != nil {
		return err
	}

	config, err := dao.GetRoomConfig(dto.Alias)
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "error getting room config")
	}

	return ctx.JSON(config)
}

func GetAllRoomConfigs(ctx *fiber.Ctx) error {
	configs, err := dao.GetAllRoomConfigs()
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "error creating room configs")
	}

	return ctx.JSON(configs)
}

func UpdateConfig(ctx *fiber.Ctx) error {
	dto := new(models.RoomConfig)
	if err := ctx.BodyParser(dto); err != nil {
		return err
	}

	if err := dao.UpdateConfig(dto).Error; err != nil {
		return fiber.NewError(fiber.StatusConflict, "error updating room config")
	}

	return ctx.JSON(dto)
}

func RemoveConfig(ctx *fiber.Ctx) error {
	var dto struct {
		Id int `json:"id"`
	}
	if err := ctx.BodyParser(&dto); err != nil {
		return err
	}

	if err := dao.RemoveConfig(dto.Id); err != nil {
		return fiber.NewError(fiber.StatusConflict, "error removing room config")
	}

	return ctx.JSON(models.GameResponse{
		Message: "success",
	})
}
