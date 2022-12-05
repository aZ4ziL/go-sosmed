package main

import (
	"encoding/gob"
	"time"

	"github.com/aZ4ziL/go-sosmed/auth"
	"github.com/aZ4ziL/go-sosmed/models"
	"github.com/aZ4ziL/go-sosmed/routers"
	"github.com/gin-contrib/sessions"
	gormsession "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
)

func init() {
	gob.Register(time.Time{})
	gob.Register(map[string]interface{}{})
}

func main() {
	r := gin.Default()
	store := gormsession.NewStore(models.GetDB(), true, auth.JWTKey)
	r.Use(sessions.Sessions("goSosMedID", store))

	v1 := r.Group("/v1")

	routers.UserRouterV1(v1)

	r.Run(":8000")
}
