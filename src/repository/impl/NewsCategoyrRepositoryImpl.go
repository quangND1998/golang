package impl

import (
	"context"
	"errors"
	"nextlend-api-web-frontend/src/database"
	"nextlend-api-web-frontend/src/entity"
	"nextlend-api-web-frontend/src/exception"
	"nextlend-api-web-frontend/src/model"
	"nextlend-api-web-frontend/src/repository"

	"gorm.io/gorm"
)

func NewsCategoryRepositoryImpl() repository.NewsCategoryRepository {
	return &newsCategoryRepositoryImpl{db: database.Get("main")}
}

type newsCategoryRepositoryImpl struct {
	db *gorm.DB
}

func (repository *newsCategoryRepositoryImpl) Insert(ctx context.Context, product entity.NewsCategory) entity.NewsCategory {
	err := repository.db.WithContext(ctx).Create(&product).Error
	exception.PanicLogging(err)
	return product
}

func (repository *newsCategoryRepositoryImpl) Update(ctx context.Context, product entity.NewsCategory) entity.NewsCategory {
	err := repository.db.WithContext(ctx).Where("product_id = ?", product.ID).Updates(&product).Error
	exception.PanicLogging(err)
	return product
}

func (repository *newsCategoryRepositoryImpl) Delete(ctx context.Context, product entity.NewsCategory) {
	err := repository.db.WithContext(ctx).Delete(&product).Error
	exception.PanicLogging(err)
}

func (repository *newsCategoryRepositoryImpl) FindById(ctx context.Context, id string) (entity.NewsCategory, error) {
	var product entity.NewsCategory
	result := repository.db.WithContext(ctx).Unscoped().Where("id = ?", id).First(&product)
	if result.RowsAffected == 0 {
		return entity.NewsCategory{}, errors.New("product Not Found")
	}
	return product, nil
}

func (repository *newsCategoryRepositoryImpl) FindAll(ctx context.Context, req model.NewsCategorySearchRequest) ([]entity.NewsCategory) {
	var products []entity.NewsCategory

	// khởi tạo query
	query := repository.db.WithContext(ctx).Model(&entity.NewsCategory{})



	// lấy dữ liệu có phân trang
	err := query.
		Preload("Parent").Preload("Children.Posts").Preload("Posts").
		Order("parent_id, sort_order, name").
		Find(&products).Error

	if err != nil {
		return nil
	}

	return products
}
