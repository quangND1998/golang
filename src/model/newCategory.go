package model


type NewsCategorySearchRequest struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Slug     string `json:"slug,omitempty"`
	Status   *int   `json:"status,omitempty"`
	ParentID *uint  `json:"parent_id,omitempty"`
}
