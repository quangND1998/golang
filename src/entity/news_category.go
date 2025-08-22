package entity

import "time"

type NewsCategory struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Slug      string    `json:"slug" gorm:"uniqueIndex;size:120;not null"`
	Name      string    `json:"name" gorm:"size:150;not null"`
	ParentID  *uint     `json:"parent_id" gorm:"index"`
	SortOrder int       `json:"sort_order" gorm:"default:0"`
	Status    int       `json:"status" gorm:"type:tinyint(1);default:1"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	// Relations
	Parent   *NewsCategory  `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Children []NewsCategory `json:"children,omitempty" gorm:"foreignKey:ParentID"`
	Posts    []NewPost      `json:"posts,omitempty" gorm:"foreignKey:CategoryID"`
}

func (NewsCategory) TableName() string {
	return "news_category"
}
