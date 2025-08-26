package controller

import (
	"nextlend-api-web-frontend/src/common/response"
	"nextlend-api-web-frontend/src/model"
	"nextlend-api-web-frontend/src/service"
	"strconv"

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
	app.Post("/api/lending-portal/new-category-flat", controller.FindAllFlat)
	app.Post("/api/lending-portal/new-category-tree", controller.GetFormattedCategoryData)
	app.Post("/api/lending-portal/new-category-tree-depth", controller.GetCategoryTreeWithCustomDepth)
	app.Get("/api/lending-portal/new-category/:id", controller.FindById)
	app.Get("/api/lending-portal/new-category/:id/tree", controller.FindCategoryWithFullTree)
}

func (controller *NewCategoryController) FindAll(c *fiber.Ctx) error {
	var req model.NewsCategorySearchRequest
	// parse query params vào struct
	if err := c.QueryParser(&req); err != nil {
		return response.BadRequest(c, "Invalid query parameters", err.Error())
	}
	result := controller.NewCategoryService.FindAll(c.Context(), req)
	return response.Success(c, result)
}

func (controller *NewCategoryController) FindAllFlat(c *fiber.Ctx) error {
	var req model.NewsCategorySearchRequest

	// parse query params vào struct
	if err := c.QueryParser(&req); err != nil {
		return response.BadRequest(c, "Invalid query parameters", err.Error())
	}
	result := controller.NewCategoryService.FindAllFlat(c.Context(), req)
	return response.Success(c, result)
}

func (controller *NewCategoryController) GetFormattedCategoryData(c *fiber.Ctx) error {
	result := controller.NewCategoryService.GetFormattedCategoryData(c.Context())
	return response.Success(c, result)
}

func (controller *NewCategoryController) GetCategoryTreeWithCustomDepth(c *fiber.Ctx) error {
	var req struct {
		MaxDepth int `json:"max_depth" query:"max_depth"`
	}

	// parse query params vào struct
	if err := c.QueryParser(&req); err != nil {
		return response.BadRequest(c, "Invalid query parameters", err.Error())
	}

	// Default depth nếu không được cung cấp
	if req.MaxDepth <= 0 {
		req.MaxDepth = 3
	}

	result := controller.NewCategoryService.GetCategoryTreeWithCustomDepth(c.Context(), req.MaxDepth)
	return response.Success(c, result)
}

func (controller *NewCategoryController) FindById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return response.BadRequest(c, "ID is required", nil)
	}

	result, err := controller.NewCategoryService.FindById(c.Context(), id)
	if err != nil {
		return response.NotFound(c, "Category not found", err.Error())
	}
	return response.Success(c, result)
}

func (controller *NewCategoryController) FindCategoryWithFullTree(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return response.BadRequest(c, "ID is required", nil)
	}

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return response.BadRequest(c, "Invalid ID format", err.Error())
	}

	result, err := controller.NewCategoryService.FindCategoryWithFullTree(c.Context(), uint(idUint))
	if err != nil {
		return response.NotFound(c, "Category not found", err.Error())
	}
	return response.Success(c, result)
}
