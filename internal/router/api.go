package router

import (
	"alpha-core/internal/handler"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
}

func NewRouter(router *gin.Engine, authHandler *handler.AuthHandler) *gin.Engine {

	api := router.Group("/api")
	api.PUT("/update", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "This is a PUT request!"})
	})

	NewAPIUsers(api, authHandler)

	return router
}
