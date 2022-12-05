package models

import (
	"database/sql"
	"time"

	"github.com/aZ4ziL/go-sosmed/auth"
)

type User struct {
	ID                   uint               `gorm:"primaryKey" json:"id"`
	FirstName            string             `gorm:"size:30" json:"first_name"`
	LastName             string             `gorm:"size:30" json:"last_name"`
	Email                string             `gorm:"size:30;unique;index" json:"email"`
	Password             string             `gorm:"size:100" json:"-"`
	IsAdmin              bool               `gorm:"default:0" json:"is_admin"`
	IsActive             bool               `gorm:"default:0" json:"is_active"`
	LastLogin            sql.NullTime       `gorm:"null" json:"last_login"`
	UpdatedAt            time.Time          `gorm:"autoUpdateTime:nano" json:"updated_at"`
	DateJoined           time.Time          `gorm:"autoCreateTime" json:"date_joined"`
	UserPosts            []UserPost         `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_posts"`
	UserPostLikes        []*UserPost        `gorm:"many2many:user_post_likes" json:"user_post_likes"`
	UserPostComments     []UserPostComment  `gorm:"foreignKey:UserID" json:"user_post_comments"`
	UserPostCommentLikes []*UserPostComment `gorm:"many2many:user_post_comment_likes" json:"user_post_comment_likes"`
	UserPostPhotoLikes   []*UserPostPhoto   `gorm:"many2many:user_post_photo_likes" json:"user_post_photo_likes"`
}

type userModel struct{}

func NewUserModel() userModel {
	return userModel{}
}

// CreateNewUser
// create a new user
func (u userModel) CreateNewUser(user *User) error {
	passwordEncrypt, err := auth.EncryptionPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = passwordEncrypt
	err = db.Create(user).Error
	return err
}

// GetUserByEmail
// get user by passing the `email`
func (u userModel) GetUserByEmail(email string) (User, error) {
	var user User
	err := db.Model(&User{}).Where("email = ?", email).
		Preload("UserPosts").
		Preload("UserPostLikes").
		Preload("UserPostComments").
		Preload("UserPostCommentLikes").
		Preload("UserPostPhotoLikes").
		First(&user).Error
	return user, err
}

// GetUserByID
// get user by passing the `id`
func (u userModel) GetUserByID(id uint) (User, error) {
	var user User
	err := db.Model(&User{}).Where("id = ?", id).
		Preload("UserPosts").
		Preload("UserPostLikes").
		Preload("UserPostComments").
		Preload("UserPostCommentLikes").
		Preload("UserPostPhotoLikes").
		First(&user).Error
	return user, err
}
