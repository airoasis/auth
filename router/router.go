package router

import (
	"github.com/airoasis/auth/handler"
	"github.com/gin-gonic/gin"
	"github.com/heptiolabs/healthcheck"
)

func SetupRouter() *gin.Engine {
	health := healthcheck.NewHandler()
	r := gin.Default()

	ug := r.Group("/users")
	{
		ug.POST("", handler.CreateUser)
		ug.GET("/:id", handler.GetUserByID)
		ug.DELETE("/:id", handler.DeleteUser)
		ug.POST("/oauth", handler.GetUserByUsernameAndPassword)
	}

	hg := r.Group("/health")
	{
		hg.GET("/live", gin.WrapF(health.LiveEndpoint))
		hg.GET("/ready", gin.WrapF(health.ReadyEndpoint))
	}

	return r
}