package services

import (
	"github.com/FarukKaradeniz/SpaceHax-server/app/dao"
	"github.com/FarukKaradeniz/SpaceHax-server/config"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"time"
)

func UserLogin(ctx *fiber.Ctx) error {
	var dto struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.BodyParser(&dto); err != nil {
		return err
	}

	user, err := dao.GetUser(dto.Username, dto.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid credentials")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = user.Name
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	t, err := token.SignedString([]byte(config.SECRET))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(fiber.Map{"token": t})
}
