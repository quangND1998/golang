package configuration

import (
	// "nextlend-api-web-frontend/src/exception"
	// "nextlend-api-web-frontend/src/exception"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func NewFiberConfiguration() fiber.Config {
	return fiber.Config{
		// ErrorHandler: exception.ErrorHandler,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	}
}
