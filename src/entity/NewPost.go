package entity

import "time"

type NewPost struct {
	ID          uint          `json:"id" gorm:"primaryKey;autoIncrement"`
	Slug        string        `json:"slug" gorm:"uniqueIndex;size:160;not null"`
	Title       string        `json:"title" gorm:"size:255;not null"`
	Summary     *string       `json:"summary" gorm:"type:text"`
	Content     string        `json:"content" gorm:"type:mediumtext;not null"`
	CoverImage  *string       `json:"cover_image" gorm:"size:255"`
	Status      string        `json:"status" gorm:"type:enum('draft','scheduled','published','archived');default:'draft'"`
	IsFeatured  int           `json:"is_featured" gorm:"type:tinyint(1);default:0"`
	PublishedAt *time.Time    `json:"published_at"`
	CreatedAt   time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	CategoryID  uint          `json:"category_id"`
	Category    *NewsCategory `json:"category" gorm:"foreignKey:CategoryID"`
}

func (NewPost) TableName() string {
	return "news_post"
}
