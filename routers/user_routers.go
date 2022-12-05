package routers

import (
	"github.com/aZ4ziL/go-sosmed/handlers"
	"github.com/gin-gonic/gin"
)

func UserRouterV1(group *gin.RouterGroup) {
	userHandler := handlers.NewUserHandler()

	group.POST("/auth/user/sign-in", userHandler.SignIn)
	group.POST("/auth/user/sign-up", userHandler.SignUp)
	group.GET("/auth/user/check-token", userHandler.CheckToken)
}
