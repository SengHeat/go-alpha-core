package handler

import (
	"alpha-core/internal/config"
	"alpha-core/internal/repository"
	"alpha-core/internal/service"
	"alpha-core/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	authService  *service.AuthService
	oauthService *service.OAuthService
	configure    *config.Config
	repository   *repository.UserRepository
	logger       *logger.Logger
}

func NewHandler(database *gorm.DB, configure *config.Config, log *logger.Logger) *AuthHandler {
	userRepository := repository.NewUserRepository(database)
	authService := service.NewAuthService(userRepository, configure)
	oauthService := service.NewOAuthService(database, configure, userRepository)

	return &AuthHandler{
		authService:  authService,
		oauthService: oauthService,
		configure:    configure,
		repository:   userRepository,
		logger:       log,
	}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func (handler *AuthHandler) Register(context *gin.Context) {
	var registerRequest RegisterRequest

	if err := context.ShouldBindJSON(&registerRequest); err != nil {
		handler.logger.ErrorLog("Failed to bind JSON", "error", err)
		context.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	user, err := handler.authService.Register(registerRequest.Username, registerRequest.Password, registerRequest.Email)

	if err != nil {
		handler.logger.ErrorLog("Failed to register user", "error", err)
		context.JSON(400, gin.H{"error": "Registration failed"})
		return
	}

	token, err := handler.authService.GenerateJWT(user)

	context.JSON(http.StatusCreated, gin.H{"user": user, "token": token})
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (handler *AuthHandler) Login(context *gin.Context) {
	var loginRequest LoginRequest

	if err := context.ShouldBindJSON(&loginRequest); err != nil {
		handler.logger.ErrorLog("Failed to bind JSON", "error", err)
		context.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	user, err := handler.authService.Authenticate(loginRequest.Username, loginRequest.Password)

	if err != nil {
		handler.logger.ErrorLog("Failed to authenticate user", "error", err)
		context.JSON(401, gin.H{"error": "Authentication failed"})
		return

	}

	token, err := handler.authService.GenerateJWT(user)
	if err != nil {
		handler.logger.ErrorLog("Failed to generate JWT", "error", err)
		context.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"user": user, "token": token})

}
