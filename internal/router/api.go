package router

import (
	"alpha-core/internal/handler"
	"alpha-core/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
}

func NewRouter(router *gin.Engine, authHandler *handler.AuthHandler, authMiddleware *middleware.JWTMiddleware) *gin.Engine {

	router.NoMethod(func(c *gin.Context) {
		c.JSON(405, gin.H{"error": "Method not allowed"})
	})

	api := router.Group("/api")
	NewAPIUsers(api, authHandler, authMiddleware)

	return router
}
