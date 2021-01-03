package utils

import (
	"github.com/gofiber/fiber/v2"
)

type httpError struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
}

// ErrorHandler is used to catch error thrown inside the routes by ctx.Next(err)
func ErrorHandler(c *fiber.Ctx, err error) error {
	// StatusCode defaults to 500
	code := fiber.StatusInternalServerError

	// Check if it's an fiber.Error type
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(&httpError{
		StatusCode: code,
		Error:      err.Error(),
	})
}
