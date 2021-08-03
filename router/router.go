package router

import (
	"github.com/airoasis/auth/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("users", handler.CreateUser)
	r.GET("users/:id", handler.GetUserByID)
	r.DELETE("users/:id", handler.DeleteUser)
	r.POST("users/oauth", handler.GetUserByUsernameAndPassword)

	return r
}