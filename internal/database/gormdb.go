package database

import (
	"alpha-core/internal/config"
	"alpha-core/internal/model"
	"alpha-core/pkg/logger"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase(configure *config.Config, log *logger.Logger) (*gorm.DB, error) {

	dataSourceName := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		configure.DBHost,
		configure.DBPort,
		configure.DBUser,
		configure.DBPass,
		configure.DBName,
	)

	database, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err = database.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Permission{},
		&model.RolePermission{},
		&model.UserRole{},
		&model.OAuthClient{},
	); err != nil {
		return nil, fmt.Errorf("failed to auto migrate: %w", err)
	}

	log.InfoLog("Database connection established successfully")

	return database, nil
}
