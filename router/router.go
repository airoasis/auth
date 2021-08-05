package router

import (
	"github.com/airoasis/user/handler"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/heptiolabs/healthcheck"
	"github.com/rs/zerolog"
	"os"
	"time"
)

func SetupRouter() *gin.Engine {
	health := healthcheck.NewHandler()

	r := gin.Default()
	r.Use(logger.SetLogger(logger.WithWriter(zerolog.ConsoleWriter{Out:os.Stderr,TimeFormat: time.RFC3339})))

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