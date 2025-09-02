package models

import "time"

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"uniqueIndex" validate:"required,min=3,max=100,unique"`
	Content   string    `json:"content" validate:"required,min=10,max=500"`
	UserID    uint      `json:"user_id"`
	CategoryID uint     `json:"category_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
	