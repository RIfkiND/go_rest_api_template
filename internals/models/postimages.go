package models

import "time"

type PostImage struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    PostID    uint      `json:"post_id"`
    ImageURL  string    `json:"image_url"`
    CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
