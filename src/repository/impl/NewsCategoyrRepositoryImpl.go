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

func (repository *newsCategoryRepositoryImpl) FindAll(ctx context.Context, req model.NewsCategorySearchRequest) []entity.NewsCategory {
	// Lấy tất cả categories và posts trong một query
	var allCategories []entity.NewsCategory
	var allPosts []entity.NewPost

	// Query tất cả categories
	err := repository.db.WithContext(ctx).
		Order("sort_order, name").
		Find(&allCategories).Error
	if err != nil {
		return nil
	}

	// Query tất cả posts
	err = repository.db.WithContext(ctx).
		Find(&allPosts).Error
	if err != nil {
		return nil
	}

	// Tạo map để truy cập nhanh
	categoryMap := make(map[uint]*entity.NewsCategory)
	postMap := make(map[uint][]entity.NewPost)

	// Map posts theo category_id
	for _, post := range allPosts {
		postMap[post.CategoryID] = append(postMap[post.CategoryID], post)
	}

	// Map categories và gán posts
	for i := range allCategories {
		category := &allCategories[i]
		categoryMap[category.ID] = category
		
		// Gán posts cho category
		if posts, exists := postMap[category.ID]; exists {
			category.Posts = posts
		}
	}

	// Xây dựng cấu trúc phân cấp
	var rootCategories []entity.NewsCategory
	for i := range allCategories {
		category := &allCategories[i]
		if category.ParentID == nil {
			// Đây là root category
			repository.buildCategoryTreeFromMap(category, categoryMap)
			rootCategories = append(rootCategories, *category)
		}
	}

	return rootCategories
}

// buildCategoryTreeFromMap xây dựng cây phân cấp từ map đã có sẵn
func (repository *newsCategoryRepositoryImpl) buildCategoryTreeFromMap(category *entity.NewsCategory, categoryMap map[uint]*entity.NewsCategory) {
	var children []entity.NewsCategory
	
	// Tìm tất cả children của category hiện tại
	for _, cat := range categoryMap {
		if cat.ParentID != nil && *cat.ParentID == category.ID {
			// Đệ quy xây dựng cây cho child
			repository.buildCategoryTreeFromMap(cat, categoryMap)
			children = append(children, *cat)
		}
	}
	
	category.Children = children
}

// FindAllFlat trả về tất cả categories dưới dạng flat list (không phân cấp)
func (repository *newsCategoryRepositoryImpl) FindAllFlat(ctx context.Context, req model.NewsCategorySearchRequest) []entity.NewsCategory {
	var categories []entity.NewsCategory

	query := repository.db.WithContext(ctx).Preload("Posts")

	// Áp dụng filter nếu có
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}
	if req.ParentID != nil {
		query = query.Where("parent_id = ?", *req.ParentID)
	} else if req.Slug == "" {
		// Nếu không có parent_id và slug, chỉ lấy root categories
		query = query.Where("parent_id IS NULL")
	}
	if req.Slug != "" {
		query = query.Where("slug = ?", req.Slug)
	}

	err := query.Order("sort_order, name").Find(&categories).Error
	if err != nil {
		return nil
	}

	return categories
}

// FindCategoryWithFullTree trả về một category cụ thể với toàn bộ cây phân cấp
func (repository *newsCategoryRepositoryImpl) FindCategoryWithFullTree(ctx context.Context, id uint) (entity.NewsCategory, error) {
	// Lấy tất cả categories và posts
	var allCategories []entity.NewsCategory
	var allPosts []entity.NewPost

	// Query tất cả categories
	err := repository.db.WithContext(ctx).
		Order("sort_order, name").
		Find(&allCategories).Error
	if err != nil {
		return entity.NewsCategory{}, err
	}

	// Query tất cả posts
	err = repository.db.WithContext(ctx).
		Find(&allPosts).Error
	if err != nil {
		return entity.NewsCategory{}, err
	}

	// Tạo map để truy cập nhanh
	categoryMap := make(map[uint]*entity.NewsCategory)
	postMap := make(map[uint][]entity.NewPost)

	// Map posts theo category_id
	for _, post := range allPosts {
		postMap[post.CategoryID] = append(postMap[post.CategoryID], post)
	}

	// Map categories và gán posts
	for i := range allCategories {
		category := &allCategories[i]
		categoryMap[category.ID] = category
		
		// Gán posts cho category
		if posts, exists := postMap[category.ID]; exists {
			category.Posts = posts
		}
	}

	// Tìm category cần thiết
	targetCategory, exists := categoryMap[id]
	if !exists {
		return entity.NewsCategory{}, errors.New("category not found")
	}

	// Xây dựng cây phân cấp cho category này
	repository.buildCategoryTreeFromMap(targetCategory, categoryMap)
	
	return *targetCategory, nil
}

// FormatCategoryTree format lại dữ liệu category tree theo cấu trúc mong muốn
func (repository *newsCategoryRepositoryImpl) FormatCategoryTree(categories []entity.NewsCategory) []map[string]interface{} {
	var result []map[string]interface{}
	
	for _, category := range categories {
		formatted := repository.formatCategory(category)
		result = append(result, formatted)
	}
	
	return result
}

// formatCategory format một category và các children của nó
func (repository *newsCategoryRepositoryImpl) formatCategory(category entity.NewsCategory) map[string]interface{} {
	formatted := map[string]interface{}{
		"id":         category.ID,
		"name":       category.Name,
		"slug":       category.Slug,
		"parent_id":  category.ParentID,
		"sort_order": category.SortOrder,
		"status":     category.Status,
		"created_at": category.CreatedAt,
	}

	// Format posts
	if len(category.Posts) > 0 {
		var posts []map[string]interface{}
		for _, post := range category.Posts {
			postMap := map[string]interface{}{
				"id":           post.ID,
				"slug":         post.Slug,
				"title":        post.Title,
				"summary":      post.Summary,
				"content":      post.Content,
				"cover_image":  post.CoverImage,
				"status":       post.Status,
				"is_featured":  post.IsFeatured,
				"published_at": post.PublishedAt,
				"created_at":   post.CreatedAt,
				"updated_at":   post.UpdatedAt,
				"category_id":  post.CategoryID,
				"category":     nil, // Tránh circular reference
			}
			posts = append(posts, postMap)
		}
		formatted["posts"] = posts
	}

	// Format children recursively
	if len(category.Children) > 0 {
		var children []map[string]interface{}
		for _, child := range category.Children {
			childMap := repository.formatCategory(child)
			children = append(children, childMap)
		}
		formatted["children"] = children
	}

	return formatted
}

// GetCategoryTreeWithDepth trả về category tree với độ sâu tối đa được chỉ định
func (repository *newsCategoryRepositoryImpl) GetCategoryTreeWithDepth(ctx context.Context, maxDepth int) []entity.NewsCategory {
	// Lấy tất cả categories và posts
	var allCategories []entity.NewsCategory
	var allPosts []entity.NewPost

	// Query tất cả categories
	err := repository.db.WithContext(ctx).
		Order("sort_order, name").
		Find(&allCategories).Error
	if err != nil {
		return nil
	}

	// Query tất cả posts
	err = repository.db.WithContext(ctx).
		Find(&allPosts).Error
	if err != nil {
		return nil
	}

	// Tạo map để truy cập nhanh
	categoryMap := make(map[uint]*entity.NewsCategory)
	postMap := make(map[uint][]entity.NewPost)

	// Map posts theo category_id
	for _, post := range allPosts {
		postMap[post.CategoryID] = append(postMap[post.CategoryID], post)
	}

	// Map categories và gán posts
	for i := range allCategories {
		category := &allCategories[i]
		categoryMap[category.ID] = category
		
		// Gán posts cho category
		if posts, exists := postMap[category.ID]; exists {
			category.Posts = posts
		}
	}

	// Xây dựng cấu trúc phân cấp với độ sâu giới hạn
	var rootCategories []entity.NewsCategory
	for i := range allCategories {
		category := &allCategories[i]
		if category.ParentID == nil {
			// Đây là root category
			repository.buildCategoryTreeWithDepthFromMap(category, categoryMap, 1, maxDepth)
			rootCategories = append(rootCategories, *category)
		}
	}

	return rootCategories
}

// buildCategoryTreeWithDepthFromMap xây dựng cây phân cấp với độ sâu giới hạn từ map
func (repository *newsCategoryRepositoryImpl) buildCategoryTreeWithDepthFromMap(category *entity.NewsCategory, categoryMap map[uint]*entity.NewsCategory, currentDepth, maxDepth int) {
	if currentDepth >= maxDepth {
		return
	}

	var children []entity.NewsCategory
	
	// Tìm tất cả children của category hiện tại
	for _, cat := range categoryMap {
		if cat.ParentID != nil && *cat.ParentID == category.ID {
			// Đệ quy xây dựng cây cho child với độ sâu tăng dần
			repository.buildCategoryTreeWithDepthFromMap(cat, categoryMap, currentDepth+1, maxDepth)
			children = append(children, *cat)
		}
	}
	
	category.Children = children
}
