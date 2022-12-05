package models

import (
	"time"

	"gorm.io/gorm"
)

type UserPost struct {
	ID        uint              `gorm:"primaryKey" json:"id"`
	UserID    uint              `json:"user_id" form:"user_id" validate:"required,number"`
	Text      string            `gorm:"type:longtext" json:"text"`
	Likes     []*User           `gorm:"many2many:user_post_likes" json:"likes"`
	UpdatedAt time.Time         `gorm:"autoUpdateTime:nano" json:"updated_at"`
	CreatedAt time.Time         `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt gorm.DeletedAt    `gorm:"index" json:"deleted_at"`
	Comments  []UserPostComment `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comments"`
	Photos    []UserPostPhoto   `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photos"`
}
