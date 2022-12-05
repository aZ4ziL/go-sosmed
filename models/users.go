package models

import "time"

type User struct {
	ID                   uint               `gorm:"primaryKey" json:"id"`
	FirstName            string             `gorm:"size:30" json:"first_name" form:"first_name" validate:"required"`
	LastName             string             `gorm:"size:30" json:"last_name" form:"last_name" validate:"required"`
	Email                string             `gorm:"size:30;unique;index" json:"email" form:"email" validate:"required,email"`
	Password             string             `gorm:"size:100" json:"-" form:"password" validate:"required"`
	IsAdmin              bool               `gorm:"default:0" json:"is_admin"`
	IsActive             bool               `gorm:"default:0" json:"is_active"`
	LastLogin            time.Time          `gorm:"null" json:"last_login"`
	UpdatedAt            time.Time          `gorm:"autoUpdateTime:nano" json:"updated_at"`
	DateJoined           time.Time          `gorm:"autoCreateTime" json:"date_joined"`
	UserPosts            []UserPost         `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_posts"`
	UserPostLikes        []*UserPost        `gorm:"many2many:user_post_likes" json:"user_post_likes"`
	UserPostComments     []UserPostComment  `gorm:"foreignKey:UserID" json:"user_post_comments"`
	UserPostCommentLikes []*UserPostComment `gorm:"many2many:user_post_comment_likes" json:"user_post_comment_likes"`
	UserPostPhotoLikes   []*UserPostPhoto   `gorm:"many2many:user_post_photo_likes" json:"user_post_photo_likes"`
}
