package configuration

import (
	// "nextlend-api-web-frontend/src/exception"

	"github.com/gofiber/fiber/v2"
)

func NewFiberConfiguration() fiber.Config {
	return fiber.Config{
		// ErrorHandler: exception.ErrorHandler,
	}
}
