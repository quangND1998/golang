package service

import (
	"context"

	"nextlend-api-web-frontend/src/model"
)

type NewCategoryService interface {
	Create(ctx context.Context, model model.ProductCreateOrUpdateModel) model.ProductCreateOrUpdateModel
	Update(ctx context.Context, productModel model.ProductCreateOrUpdateModel, id string) model.ProductCreateOrUpdateModel
	Delete(ctx context.Context, id string)
	FindById(ctx context.Context, id string) model.ProductModel
	FindAll(ctx context.Context, req model.NewsCategorySearchRequest) ([]model.ProductModel)
}
