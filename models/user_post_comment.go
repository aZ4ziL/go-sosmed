package models

import (
	"time"

	"gorm.io/gorm"
)

type UserPostComment struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `json:"user_id" form:"user_id" validate:"required,number"`
	PostID    uint           `json:"post_id" form:"post_id" validate:"required,number"`
	Text      string         `gorm:"type:longtext" json:"text" form:"text" validate:"required"`
	Likes     []*User        `gorm:"many2many:user_post_comment_likes" json:"likes"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:nano" json:"updated_at"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
