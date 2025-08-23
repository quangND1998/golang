package impl

import (
	"context"
	"fmt"
	"nextlend-api-web-frontend/src/model"
	"nextlend-api-web-frontend/src/service"
)

// ExampleUsage minh họa cách sử dụng các function mới với cách tiếp cận hiệu quả
func ExampleUsage() {
	ctx := context.Background()
	
	// Khởi tạo repository và service
	repo := NewsCategoryRepositoryImpl()
	newsCategoryService := service.NewCategoryServiceImpl(repo)

	// 1. Lấy tất cả categories với cấu trúc phân cấp đầy đủ (chỉ 2 queries)
	fmt.Println("=== Lấy tất cả categories với cấu trúc phân cấp ===")
	categories := newsCategoryService.FindAll(ctx, model.NewsCategorySearchRequest{})
	
	// Format dữ liệu theo cấu trúc mong muốn
	formattedData := newsCategoryService.GetFormattedCategoryData(ctx)
	fmt.Printf("Số lượng root categories: %d\n", len(formattedData))

	// 2. Lấy categories dưới dạng flat list
	fmt.Println("\n=== Lấy categories dưới dạng flat list ===")
	flatCategories := newsCategoryService.FindAllFlat(ctx, model.NewsCategorySearchRequest{})
	fmt.Printf("Số lượng categories: %d\n", len(flatCategories))

	// 3. Lấy một category cụ thể với toàn bộ cây phân cấp
	fmt.Println("\n=== Lấy category cụ thể với cây phân cấp ===")
	if len(categories) > 0 {
		categoryWithTree, err := newsCategoryService.FindCategoryWithFullTree(ctx, categories[0].ID)
		if err == nil {
			fmt.Printf("Category: %s, Số children: %d\n", categoryWithTree.Name, len(categoryWithTree.Children))
		}
	}

	// 4. Lấy category tree với độ sâu giới hạn
	fmt.Println("\n=== Lấy category tree với độ sâu giới hạn ===")
	categoriesWithDepth := newsCategoryService.GetCategoryTreeWithDepth(ctx, 2) // Chỉ lấy 2 cấp
	fmt.Printf("Số lượng root categories với depth=2: %d\n", len(categoriesWithDepth))
}

// GetFormattedCategoryData trả về dữ liệu đã được format theo cấu trúc mong muốn
func GetFormattedCategoryData(ctx context.Context) []map[string]interface{} {
	repo := NewsCategoryRepositoryImpl()
	newsCategoryService := service.NewCategoryServiceImpl(repo)
	
	// Lấy tất cả categories với cấu trúc phân cấp đầy đủ và format
	return newsCategoryService.GetFormattedCategoryData(ctx)
}

// GetCategoryTreeWithCustomDepth trả về category tree với độ sâu tùy chỉnh
func GetCategoryTreeWithCustomDepth(ctx context.Context, maxDepth int) []map[string]interface{} {
	repo := NewsCategoryRepositoryImpl()
	newsCategoryService := service.NewCategoryServiceImpl(repo)
	
	// Lấy categories với độ sâu giới hạn và format
	return newsCategoryService.GetCategoryTreeWithCustomDepth(ctx, maxDepth)
}

// PerformanceComparison so sánh hiệu suất giữa cách cũ và cách mới
func PerformanceComparison() {
	fmt.Println("=== So sánh hiệu suất ===")
	fmt.Println("Cách cũ (N+1 queries):")
	fmt.Println("- Query 1: Lấy root categories")
	fmt.Println("- Query 2: Lấy children của category 1")
	fmt.Println("- Query 3: Lấy children của category 2")
	fmt.Println("- Query 4: Lấy children của category 3")
	fmt.Println("- ... và cứ thế cho mỗi category")
	fmt.Println("- Tổng: 1 + N queries (N = số lượng categories)")
	
	fmt.Println("\nCách mới (2 queries):")
	fmt.Println("- Query 1: Lấy tất cả categories")
	fmt.Println("- Query 2: Lấy tất cả posts")
	fmt.Println("- Map dữ liệu trong memory để tạo cấu trúc phân cấp")
	fmt.Println("- Tổng: 2 queries cố định")
	
	fmt.Println("\nƯu điểm của cách mới:")
	fmt.Println("1. Hiệu suất cao hơn nhiều lần")
	fmt.Println("2. Không bị N+1 query problem")
	fmt.Println("3. Dữ liệu được load một lần và tái sử dụng")
	fmt.Println("4. Có thể xử lý cấu trúc phân cấp phức tạp")
}
