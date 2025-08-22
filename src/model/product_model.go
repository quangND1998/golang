package model

import "nextlend-api-web-frontend/src/entity"

type ProductModel struct {
	Id        uint                  `json:"id"`
	Name      string                `json:"name"`
	Slug      string                `json:"slug"`
	ParentID  *uint                 `json:"parent_id"`
	SortOrder int                   `json:"sort_order"`
	Status    int                   `json:"status"`
	Parent    *entity.NewsCategory  `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Children  []entity.NewsCategory `json:"children,omitempty" gorm:"foreignKey:ParentID"`
	Posts     []entity.NewPost      `json:"post,omitempty" gorm:"foreignKey:ProductID"`
}

type ProductCreateOrUpdateModel struct {
	Name     string `json:"name" validate:"required"`
	Price    int64  `json:"price" validate:"required"`
	Quantity int32  `json:"quantity" validate:"required"`
}
