package repository

import (
	"context"
	"nextlend-api-web-frontend/src/entity"
	"nextlend-api-web-frontend/src/model"
)

type NewsCategoryRepository interface {
	Insert(ctx context.Context, news_category entity.NewsCategory) entity.NewsCategory
	Update(ctx context.Context, news_category entity.NewsCategory) entity.NewsCategory
	Delete(ctx context.Context, news_category entity.NewsCategory)
	FindById(ctx context.Context, id string) (entity.NewsCategory, error)
	FindAll(ctx context.Context, newsCategorySearchRequest model.NewsCategorySearchRequest) ([]entity.NewsCategory)
}
