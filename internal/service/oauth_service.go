package service

import (
	"alpha-core/internal/config"
	"alpha-core/internal/repository"

	"gorm.io/gorm"
)

type OAuthService struct {
	database       *gorm.DB
	configure      *config.Config
	userRepository *repository.UserRepository
}

func NewOAuthService(database *gorm.DB, configure *config.Config, userRepository *repository.UserRepository) *OAuthService {
	return &OAuthService{
		database:       database,
		configure:      configure,
		userRepository: userRepository,
	}
}
