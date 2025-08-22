package model

// Base Response Models
// ===================

// GeneralResponse là model response chuẩn cho tất cả API
// Sử dụng để trả về dữ liệu với format thống nhất
type GeneralResponse struct {
	Code    int         `json:"code"`    // Mã HTTP status code
	Message string      `json:"message"` // Thông báo kết quả
	Data    interface{} `json:"data"`    // Dữ liệu trả về (có thể là bất kỳ kiểu dữ liệu nào)
}

// Base Request Models
// ==================

// PaginationRequest là model base cho các request có phân trang
type PaginationRequest struct {
	Page     int `json:"page" query:"page"`           // Trang hiện tại (bắt đầu từ 1)
	PageSize int `json:"page_size" query:"page_size"` // Số lượng item trên mỗi trang
}

// SearchRequest là model base cho các request có tìm kiếm
type SearchRequest struct {
	Keyword string `json:"keyword" query:"keyword"` // Từ khóa tìm kiếm
}

// Base Response Models cho Pagination
// ==================================

// PaginationMeta chứa thông tin phân trang
type PaginationMeta struct {
	CurrentPage int   `json:"current_page"` // Trang hiện tại
	PageSize    int   `json:"page_size"`    // Số item trên mỗi trang
	TotalItems  int64 `json:"total_items"`  // Tổng số item
	TotalPages  int   `json:"total_pages"`  // Tổng số trang
	HasNext     bool  `json:"has_next"`     // Có trang tiếp theo không
	HasPrevious bool  `json:"has_previous"` // Có trang trước không
}

// PaginatedResponse là response chuẩn cho dữ liệu có phân trang
type PaginatedResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    interface{}    `json:"data"` // Danh sách dữ liệu
	Meta    PaginationMeta `json:"meta"` // Thông tin phân trang
}

// Error Response Models
// ====================

// ErrorDetail chứa chi tiết lỗi validation
type ErrorDetail struct {
	Field   string `json:"field"`   // Tên field bị lỗi
	Message string `json:"message"` // Thông báo lỗi
}

// ValidationErrorResponse là response cho lỗi validation
type ValidationErrorResponse struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Errors  []ErrorDetail `json:"errors"` // Danh sách các lỗi validation
}

// Success Response Models
// ======================

// SuccessResponse là response đơn giản cho các thao tác thành công
type SuccessResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// DataResponse là response chỉ chứa dữ liệu
type DataResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// Helper Functions
// ================

// NewSuccessResponse tạo response thành công
func NewSuccessResponse(data interface{}) GeneralResponse {
	return GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    data,
	}
}

// NewErrorResponse tạo response lỗi
func NewErrorResponse(code int, message string, data interface{}) GeneralResponse {
	return GeneralResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// NewValidationErrorResponse tạo response lỗi validation
func NewValidationErrorResponse(errors []ErrorDetail) ValidationErrorResponse {
	return ValidationErrorResponse{
		Code:    400,
		Message: "Validation Error",
		Errors:  errors,
	}
}

// NewPaginatedResponse tạo response có phân trang
func NewPaginatedResponse(data interface{}, meta PaginationMeta) PaginatedResponse {
	return PaginatedResponse{
		Code:    200,
		Message: "Success",
		Data:    data,
		Meta:    meta,
	}
}
