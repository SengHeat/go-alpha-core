package router

import (
	"alpha-core/internal/handler"
	"alpha-core/internal/middleware"

	"github.com/gin-gonic/gin"
)

type APIUsers struct {
	router *gin.RouterGroup
}

func NewAPIUsers(router *gin.RouterGroup, authHandler *handler.AuthHandler, authMiddleWare *middleware.JWTMiddleware) *APIUsers {
	usersGroup := router.Group("/users")

	// Initialize the APIUsers struct
	apiUsers := &APIUsers{router: usersGroup}

	// Register endpoints
	// login and register without middleware
	apiUsers.router.POST("/register", authHandler.Register)
	apiUsers.router.POST("/login", authHandler.Login)
	apiUsers.router.Use(middleware.Middleware([]string{"auth", "rate-limited"}, authMiddleWare)...).GET("/profile", authHandler.Profile)

	return apiUsers
}
