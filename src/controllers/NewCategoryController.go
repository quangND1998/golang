package controller

import (
	"nextlend-api-web-frontend/src/common/response"
	"nextlend-api-web-frontend/src/model"
	"nextlend-api-web-frontend/src/service"

	"github.com/gofiber/fiber/v2"
)

type NewCategoryController struct {
	service.NewCategoryService
}

func InitNewCategoryController(newCategoryService *service.NewCategoryService) *NewCategoryController {
	return &NewCategoryController{NewCategoryService: *newCategoryService}
}

func (controller *NewCategoryController) Route(app *fiber.App) {
	app.Post("/api/lending-portal/new-category-list", controller.FindAll)
}

func (controller *NewCategoryController) FindAll(c *fiber.Ctx) error {
	// return response.BadRequest(c, "This endpoint is not implemented yet. Please check back later.", nil)
	var req model.NewsCategorySearchRequest

	// parse query params vào struct
	if err := c.QueryParser(&req); err != nil {
		return response.BadRequest(c, "Invalid query parameters", err.Error())
	}
	result := controller.NewCategoryService.FindAll(c.Context(), req)
	return response.Success(c, result)
}
