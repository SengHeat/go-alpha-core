package router

import (
	"alpha-core/internal/handler"

	"github.com/gin-gonic/gin"
)

type APIUsers struct {
	router *gin.RouterGroup
}

func NewAPIUsers(router *gin.RouterGroup, authHandler *handler.AuthHandler) *APIUsers {
	usersGroup := router.Group("/users")

	// Initialize the APIUsers struct
	apiUsers := &APIUsers{router: usersGroup}

	// Register endpoints
	apiUsers.router.GET("/", apiUsers.ListUsers)
	apiUsers.router.POST("/", apiUsers.CreateUser)
	apiUsers.router.GET("/:id", apiUsers.GetUser)
	apiUsers.router.PUT("/:id", apiUsers.UpdateUser)
	apiUsers.router.DELETE("/:id", apiUsers.DeleteUser)
	apiUsers.router.POST("/register", authHandler.Register)
	apiUsers.router.POST("/login", authHandler.Login)

	return apiUsers
}

func (a *APIUsers) ListUsers(c *gin.Context)  { c.JSON(200, gin.H{"message": "list users"}) }
func (a *APIUsers) CreateUser(c *gin.Context) { c.JSON(200, gin.H{"message": "create user"}) }
func (a *APIUsers) GetUser(c *gin.Context)    { c.JSON(200, gin.H{"message": "get user"}) }
func (a *APIUsers) UpdateUser(c *gin.Context) { c.JSON(200, gin.H{"message": "update user"}) }
func (a *APIUsers) DeleteUser(c *gin.Context) { c.JSON(200, gin.H{"message": "delete user"}) }
