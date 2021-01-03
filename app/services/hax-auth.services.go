package services

import "github.com/gofiber/fiber/v2"

func Login(ctx *fiber.Ctx) error {
	return ctx.SendString("this is from login")
}

func SignUp(ctx *fiber.Ctx) error {
	return ctx.SendString("this is from sign up")
}
