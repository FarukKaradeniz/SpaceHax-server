package middleware

import (
	"github.com/FarukKaradeniz/SpaceHax-server/config"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// Protected protect routes
func Protected(roles []string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(config.SECRET),
		SuccessHandler: func(ctx *fiber.Ctx) error {
			return role(ctx, roles)
		},
	})
}

func role(ctx *fiber.Ctx, roles []string) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	role := claims["role"].(string)
	isExist := false
	for _, ro := range roles {
		if ro == role {
			isExist = true
			break
		}
	}
	if !isExist {
		return ctx.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
	}

	return ctx.Next()
}
