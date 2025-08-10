package repository

import (
	"alpha-core/internal/model"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{
		database: database,
	}
}

func (repository *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User

	err := repository.database.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, err
}

func (repository *UserRepository) Create(user *model.User) error {
	return repository.database.Create(user).Error
}
