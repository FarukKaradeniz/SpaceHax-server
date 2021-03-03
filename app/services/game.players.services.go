package services

import (
	"errors"
	"github.com/FarukKaradeniz/SpaceHax-server/app/dao"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LoginDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	RoomId   string `json:"room"`
}

type SignUpDTO struct {
	LoginDTO
	Connection string `json:"conn"`
}

func Login(ctx *fiber.Ctx) error {
	dto := new(LoginDTO)
	if err := ctx.BodyParser(dto); err != nil {
		return err
	}

	player := &dao.Player{}
	err := dao.GetPlayerByNameAndPassword(player, dto.Name, dto.Password, dto.RoomId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid name or password")
	}
	return ctx.JSON(struct {
		IsAdmin      *bool `json:"isAdmin,omitempty"`
		IsSuperAdmin *bool `json:"isSuperAdmin,omitempty"`
		PlayerId     uint  `json:"playerId,omitempty"`
	}{
		IsAdmin:      &player.IsAdmin,
		IsSuperAdmin: &player.IsSuperAdmin,
		PlayerId:     player.ID,
	})
}

func SignUp(ctx *fiber.Ctx) error {
	dto := new(SignUpDTO)
	if err := ctx.BodyParser(dto); err != nil {
		return err
	}

	player := &dao.Player{}
	err := dao.GetPlayerByName(player, dto.Name, dto.RoomId).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusConflict, "you already have an account")
	}

	player.Name = dto.Name
	player.Password = dto.Password
	player.Connection = dto.Connection
	player.RoomId = dto.RoomId

	if err := dao.CreatePlayer(player); err.Error != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error.Error())
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func ChangePassword(ctx *fiber.Ctx) error {
	var dto struct {
		PlayerId uint   `json:"playerId"`
		Password string `json:"password"`
		RoomId   string `json:"room"`
	}
	if err := ctx.BodyParser(&dto); err != nil {
		return err
	}

	if err := dao.ChangePassword(dto.PlayerId, dto.Password, dto.RoomId).Error; err != nil {
		return fiber.NewError(fiber.StatusConflict, "error changing password")
	}

	return ctx.SendStatus(fiber.StatusOK)
}
