package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JWTKey = []byte("mySecretKey")

type Credential struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type Claims struct {
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	IsAdmin    bool      `json:"is_admin"`
	IsActive   bool      `json:"is_active"`
	LastLogin  time.Time `json:"last_login"`
	DateJoined time.Time `json:"date_joined"`
	jwt.RegisteredClaims
}
