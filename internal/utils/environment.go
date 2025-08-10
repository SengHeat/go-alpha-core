package utils

import (
	"alpha-core/internal/config"
	"alpha-core/pkg/logger"
)

const (
	Development = "development"
	Production  = "production"
)

func GetEnvironment(configure *config.Config, log *logger.Logger) string {
	return configure.DBName
}

func IsProduction(configure *config.Config, log *logger.Logger) bool {
	log.InfoLog("Running in " + configure.AppEnv + " mode")
	return configure.AppEnv == Production
}

func IsDevelopment(configure *config.Config, log *logger.Logger) bool {
	log.InfoLog("Running in " + configure.AppEnv + " mode")
	return configure.AppEnv == Development
}
