package impl

import (
	"context"
	"nextlend-api-web-frontend/src/common/logger"
	"nextlend-api-web-frontend/src/entity"
	"nextlend-api-web-frontend/src/exception"
	"nextlend-api-web-frontend/src/model"
	"nextlend-api-web-frontend/src/repository"
	"nextlend-api-web-frontend/src/service"
)

func NewCategoryServiceImpl(newsCategoryRepository *repository.NewsCategoryRepository) service.NewCategoryService {
	return &newCategoryServiceImpl{NewsCategoryRepository: *newsCategoryRepository}
}

type newCategoryServiceImpl struct {
	repository.NewsCategoryRepository
}

func (service *newCategoryServiceImpl) Create(ctx context.Context, productModel model.ProductCreateOrUpdateModel) model.ProductCreateOrUpdateModel {
	// common.Validate(productModel)
	product := entity.NewsCategory{
		Name: productModel.Name,
	}
	service.NewsCategoryRepository.Insert(ctx, product)
	return productModel
}

func (service *newCategoryServiceImpl) Update(ctx context.Context, productModel model.ProductCreateOrUpdateModel, id string) model.ProductCreateOrUpdateModel {
	// common.Validate(productModel)
	product := entity.NewsCategory{

		Name: productModel.Name,
	}
	service.NewsCategoryRepository.Update(ctx, product)
	return productModel
}

func (service *newCategoryServiceImpl) Delete(ctx context.Context, id string) {
	product, err := service.NewsCategoryRepository.FindById(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	service.NewsCategoryRepository.Delete(ctx, product)
}

func (service *newCategoryServiceImpl) FindById(ctx context.Context, id string) model.ProductModel {
	productCache, err := service.NewsCategoryRepository.FindById(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	return model.ProductModel{
		// Id:       productCache.Id.String(),
		Name: productCache.Name,
		// Price:    productCache.Price,
		// Quantity: productCache.Quantity,
	}
}

func (s *newCategoryServiceImpl) FindAll(ctx context.Context, req model.NewsCategorySearchRequest) []model.ProductModel {
	products := s.NewsCategoryRepository.FindAll(ctx, req)
	logger.Info("Find all products", products)
	// Trường hợp không có dữ liệu thì trả luôn empty slice
	if len(products) == 0 {
		return []model.ProductModel{}
	}

	// Map entity -> response model
	responses := make([]model.ProductModel, 0, len(products))
	for _, product := range products {
		responses = append(responses, model.ProductModel{
			Id:        product.ID,
			Name:      product.Name,
			Slug:      product.Slug,
			ParentID:  product.ParentID,
			SortOrder: product.SortOrder,
			Status:    product.Status,

			Parent:   product.Parent,
			Children: product.Children,
			Posts:    product.Posts,
			// Id:       product.Id.String(),
			// Price:    product.Price,
			// Quantity: product.Quantity,
		})
	}

	return responses
}
