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

func NewCategoryServiceImpl(newsCategoryRepository repository.NewsCategoryRepository) service.NewCategoryService {
	return &newCategoryServiceImpl{NewsCategoryRepository: newsCategoryRepository}
}

type newCategoryServiceImpl struct {
	repository.NewsCategoryRepository
}

func (service *newCategoryServiceImpl) Create(ctx context.Context, category entity.NewsCategory) entity.NewsCategory {
	// common.Validate(category)
	result := service.NewsCategoryRepository.Insert(ctx, category)
	return result
}

func (service *newCategoryServiceImpl) Update(ctx context.Context, category entity.NewsCategory, id string) entity.NewsCategory {
	// common.Validate(category)
	// Set ID từ parameter
	category.ID = uint(0) // Sẽ được set từ FindById
	existingCategory, err := service.NewsCategoryRepository.FindById(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	category.ID = existingCategory.ID
	result := service.NewsCategoryRepository.Update(ctx, category)
	return result
}

func (service *newCategoryServiceImpl) Delete(ctx context.Context, id string) {
	category, err := service.NewsCategoryRepository.FindById(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	service.NewsCategoryRepository.Delete(ctx, category)
}

func (service *newCategoryServiceImpl) FindById(ctx context.Context, id string) (entity.NewsCategory, error) {
	return service.NewsCategoryRepository.FindById(ctx, id)
}

func (service *newCategoryServiceImpl) FindAll(ctx context.Context, req model.NewsCategorySearchRequest) []entity.NewsCategory {
	categories := service.NewsCategoryRepository.FindAll(ctx, req)
	logger.Info("Find all categories", categories)
	return categories
}

func (service *newCategoryServiceImpl) FindAllFlat(ctx context.Context, req model.NewsCategorySearchRequest) []entity.NewsCategory {
	return service.NewsCategoryRepository.FindAllFlat(ctx, req)
}

func (service *newCategoryServiceImpl) FindCategoryWithFullTree(ctx context.Context, id uint) (entity.NewsCategory, error) {
	return service.NewsCategoryRepository.FindCategoryWithFullTree(ctx, id)
}

func (service *newCategoryServiceImpl) GetCategoryTreeWithDepth(ctx context.Context, maxDepth int) []entity.NewsCategory {
	return service.NewsCategoryRepository.GetCategoryTreeWithDepth(ctx, maxDepth)
}

func (service *newCategoryServiceImpl) GetFormattedCategoryData(ctx context.Context) []map[string]interface{} {
	categories := service.NewsCategoryRepository.FindAll(ctx, model.NewsCategorySearchRequest{})
	return service.NewsCategoryRepository.FormatCategoryTree(categories)
}

func (service *newCategoryServiceImpl) GetCategoryTreeWithCustomDepth(ctx context.Context, maxDepth int) []map[string]interface{} {
	categories := service.NewsCategoryRepository.GetCategoryTreeWithDepth(ctx, maxDepth)
	return service.NewsCategoryRepository.FormatCategoryTree(categories)
}
