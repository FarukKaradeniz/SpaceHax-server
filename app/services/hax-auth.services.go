package services

import (
	"errors"
	"github.com/FarukKaradeniz/SpaceHax-server/app/dao"
	"github.com/FarukKaradeniz/SpaceHax-server/app/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Login(ctx *fiber.Ctx) error {
	dto := new(models.LoginDTO)
	if err := ctx.BodyParser(dto); err != nil {
		return err
	}

	player := &dao.Player{}
	err := dao.GetPlayerByNameAndPassword(player, dto.Name, dto.Password).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid name or password")
	}

	return ctx.JSON(models.AuthResponse{
		Message:      "success",
		IsAdmin:      &player.IsAdmin,
		IsSuperAdmin: &player.IsSuperAdmin,
		PlayerId:     player.ID,
	})
}

func SignUp(ctx *fiber.Ctx) error {
	dto := new(models.SignUpDTO)
	if err := ctx.BodyParser(dto); err != nil {
		return err
	}

	player := &dao.Player{}
	err := dao.GetPlayerByNameAndPassword(player, dto.Name, dto.Password).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusConflict, "you already have an account")
	}

	player.Name = dto.Name
	player.Password = dto.Password
	player.Connection = dto.Connection

	if err := dao.CreatePlayer(player); err.Error != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error.Error())
	}

	return ctx.JSON(models.AuthResponse{
		Message: "success",
	})
}

func ChangePassword(ctx *fiber.Ctx) error {
	var dto struct {
		PlayerId uint   `json:"playerId"`
		Password string `json:"password"`
	}
	if err := ctx.BodyParser(&dto); err != nil {
		return err
	}

	if err := dao.ChangePassword(dto.PlayerId, dto.Password).Error; err != nil {
		return fiber.NewError(fiber.StatusConflict, "error changing password")
	}

	return ctx.JSON(models.AuthResponse{
		Message: "success",
	})
}
