package response

import (
	"nextlend-api-web-frontend/src/model"

	"github.com/gofiber/fiber/v2"
)

// ResponseHelper chứa các helper function để trả về response
type ResponseHelper struct{}

// NewResponseHelper tạo instance mới của ResponseHelper
func NewResponseHelper() *ResponseHelper {
	return &ResponseHelper{}
}

// Success trả về response thành công với status 200
func (h *ResponseHelper) Success(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    data,
	})
}

// SuccessWithMessage trả về response thành công với message tùy chỉnh
func (h *ResponseHelper) SuccessWithMessage(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: message,
		Data:    data,
	})
}

// Created trả về response thành công với status 201 (Created)
func (h *ResponseHelper) Created(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    201,
		Message: "Created Successfully",
		Data:    data,
	})
}

// Error trả về response lỗi
func (h *ResponseHelper) Error(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(model.GeneralResponse{
		Code:    statusCode,
		Message: message,
		Data:    data,
	})
}

// BadRequest trả về response lỗi 400
func (h *ResponseHelper) BadRequest(c *fiber.Ctx, message string, data interface{}) error {
	return h.Error(c, fiber.StatusBadRequest, message, data)
}

// Unauthorized trả về response lỗi 401
func (h *ResponseHelper) Unauthorized(c *fiber.Ctx, message string, data interface{}) error {
	return h.Error(c, fiber.StatusUnauthorized, message, data)
}

// NotFound trả về response lỗi 404
func (h *ResponseHelper) NotFound(c *fiber.Ctx, message string, data interface{}) error {
	return h.Error(c, fiber.StatusNotFound, message, data)
}

// InternalServerError trả về response lỗi 500
func (h *ResponseHelper) InternalServerError(c *fiber.Ctx, message string, data interface{}) error {
	return h.Error(c, fiber.StatusInternalServerError, message, data)
}

// Paginated trả về response có phân trang
func (h *ResponseHelper) Paginated(c *fiber.Ctx, data interface{}, meta model.PaginationMeta) error {
	return c.Status(fiber.StatusOK).JSON(model.PaginatedResponse{
		Code:    200,
		Message: "Success",
		Data:    data,
		Meta:    meta,
	})
}

// ValidationError trả về response lỗi validation
func (h *ResponseHelper) ValidationError(c *fiber.Ctx, errors []model.ErrorDetail) error {
	return c.Status(fiber.StatusBadRequest).JSON(model.ValidationErrorResponse{
		Code:    400,
		Message: "Validation Error",
		Errors:  errors,
	})
}

// Global Response Helper Instance
// ==============================

var responseHelper = NewResponseHelper()

// Global Functions để sử dụng trực tiếp
// =====================================

// Success trả về response thành công với status 200
func Success(c *fiber.Ctx, data interface{}) error {
	return responseHelper.Success(c, data)
}

// SuccessWithMessage trả về response thành công với message tùy chỉnh
func SuccessWithMessage(c *fiber.Ctx, message string, data interface{}) error {
	return responseHelper.SuccessWithMessage(c, message, data)
}

// Created trả về response thành công với status 201
func Created(c *fiber.Ctx, data interface{}) error {
	return responseHelper.Created(c, data)
}

// Error trả về response lỗi
func Error(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return responseHelper.Error(c, statusCode, message, data)
}

// BadRequest trả về response lỗi 400
func BadRequest(c *fiber.Ctx, message string, data interface{}) error {
	return responseHelper.BadRequest(c, message, data)
}

// Unauthorized trả về response lỗi 401
func Unauthorized(c *fiber.Ctx, message string, data interface{}) error {
	return responseHelper.Unauthorized(c, message, data)
}

// NotFound trả về response lỗi 404
func NotFound(c *fiber.Ctx, message string, data interface{}) error {
	return responseHelper.NotFound(c, message, data)
}

// InternalServerError trả về response lỗi 500
func InternalServerError(c *fiber.Ctx, message string, data interface{}) error {
	return responseHelper.InternalServerError(c, message, data)
}

// Paginated trả về response có phân trang
func Paginated(c *fiber.Ctx, data interface{}, meta model.PaginationMeta) error {
	return responseHelper.Paginated(c, data, meta)
}

// ValidationError trả về response lỗi validation
func ValidationError(c *fiber.Ctx, errors []model.ErrorDetail) error {
	return responseHelper.ValidationError(c, errors)
}
