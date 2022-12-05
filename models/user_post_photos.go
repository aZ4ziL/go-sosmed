package models

import (
	"time"

	"gorm.io/gorm"
)

type UserPostPhoto struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	PostID    uint           `json:"post_id" form:"post_id" validate:"required,number"`
	File      string         `gorm:"size:255;null" json:"file" form:"file" validate:"required"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:nano" json:"updated_at"`
	Likes     []*User        `gorm:"many2many:user_post_photo_likes" json:"likes"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
