package exception

import (
	"encoding/json"
	"nextlend-api-web-frontend/src/common/logger"
	"nextlend-api-web-frontend/src/common/response"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	logger.Error("ErrorHandler: ", err.Error())
	_, validationError := err.(ValidationError)
	if validationError {
		data := err.Error()
		var messages []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &messages)
		PanicLogging(errJson)
		// return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
		// 	Code:    400,
		// 	Message: "Bad Request",
		// 	Data:    messages,
		// })
		return response.BadRequest(ctx, "Validation Error", messages)
	}
	_, notFoundError := err.(NotFoundError)
	if notFoundError {
		return response.NotFound(ctx, "Resource Not Found", err.Error())
	}

	_, unauthorizedError := err.(UnauthorizedError)
	if unauthorizedError {
		// return ctx.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
		// 	Code:    401,
		// 	Message: "Unauthorized",
		// 	Data:    err.Error(),
		// })
		return response.Unauthorized(ctx, "Unauthorized", err.Error())
	}

	// return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
	// 	Code:    500,
	// 	Message: "General Error",
	// 	Data:    err.Error(),
	// })
	return response.InternalServerError(ctx, "Internal Server Error", err.Error())
}
