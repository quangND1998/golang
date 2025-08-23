package service

import (
	"context"
	"nextlend-api-web-frontend/src/entity"
	"nextlend-api-web-frontend/src/model"
)

type NewCategoryService interface {
	Create(ctx context.Context, category entity.NewsCategory) entity.NewsCategory
	Update(ctx context.Context, category entity.NewsCategory, id string) entity.NewsCategory
	Delete(ctx context.Context, id string)
	FindById(ctx context.Context, id string) (entity.NewsCategory, error)
	FindAll(ctx context.Context, req model.NewsCategorySearchRequest) []entity.NewsCategory
	FindAllFlat(ctx context.Context, req model.NewsCategorySearchRequest) []entity.NewsCategory
	FindCategoryWithFullTree(ctx context.Context, id uint) (entity.NewsCategory, error)
	GetCategoryTreeWithDepth(ctx context.Context, maxDepth int) []entity.NewsCategory
	GetFormattedCategoryData(ctx context.Context) []map[string]interface{}
	GetCategoryTreeWithCustomDepth(ctx context.Context, maxDepth int) []map[string]interface{}
}
