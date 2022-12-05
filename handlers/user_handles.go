package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aZ4ziL/go-sosmed/auth"
	"github.com/aZ4ziL/go-sosmed/models"
	"github.com/aZ4ziL/go-sosmed/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type userHandler struct{}

func NewUserHandler() userHandler {
	return userHandler{}
}

// SignUp
// register a new user
func (u userHandler) SignUp(ctx *gin.Context) {
	var userRequest UserRequest

	err := ctx.ShouldBindWith(&userRequest, binding.FormPost)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	validate = validator.New()
	err = validate.Struct(&userRequest)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err.Error())
			return
		}
		errorMessages := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field: %s with error type: %s", err.Field(), err.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": errorMessages,
		})
		return
	}

	// Save the user
	user := models.User{
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Email:     userRequest.Email,
		Password:  userRequest.Password,
	}
	err = models.NewUserModel().CreateNewUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

// SignIn
// this function to add login method
func (u userHandler) SignIn(ctx *gin.Context) {
	// Check if user is already sign in
	session := sessions.Default(ctx)
	if user := session.Get("user"); user != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "already_joined",
			"message": "You has already joined.",
		})
		return
	}
	var creds auth.Credential

	err := ctx.ShouldBindWith(&creds, binding.FormPost)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	validate = validator.New()
	err = validate.Struct(&creds)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err.Error())
			return
		}
		errorMessages := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field: %s, with error type: %s", err.Field(), err.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": errorMessages,
		})
		return
	}

	user, err := models.NewUserModel().GetUserByEmail(creds.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "login_failed",
			"message": "Email or password is incorrect.",
		})
		return
	}

	// Declare the expiration time of the token
	// here, we have kept it as 24 Hour
	expirationTime := time.Now().Add(24 * time.Hour)
	// Create the JWT claims, which includes the email and expiry time
	claims := &auth.Claims{
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		IsAdmin:    user.IsAdmin,
		IsActive:   user.IsActive,
		LastLogin:  user.LastLogin.Time,
		DateJoined: user.DateJoined,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT String
	tokenString, err := token.SignedString(auth.JWTKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	authSession := map[string]interface{}{
		"token": tokenString,
	}

	session.Set("user", authSession)
	session.Options(sessions.Options{
		MaxAge: 3600 * 24,
	})
	err = session.Save()
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

// CheckToken
// this is for check the token auth
func (u userHandler) CheckToken(ctx *gin.Context) {
	session := sessions.Default(ctx)

	user := session.Get("user")
	if user == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Please login using token.",
		})
		return
	}

	// Get token string from the session
	tokenString := utils.GetInterfaceValue(user)["token"].(string)
	// Initialize a new instance of `Claims`
	claims := &auth.Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return auth.JWTKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(http.StatusUnauthorized, nil)
			return
		}
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}

	// Check if token is valid or not
	if !tkn.Valid {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}

	ctx.JSON(http.StatusOK, claims)
}
