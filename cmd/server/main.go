package main

import (
	"alpha-core/internal/config"
	"alpha-core/pkg/logger"
	"fmt"
)

func main() {
	configure := config.Load()
	log := logger.Init(configure)

	fmt.Println(configure)

	log.InfoLog("failed to connect db: %v", configure)
	log.WarningLog("failed to connect db: %v", configure)
	log.ErrorLog("failed to connect db: %v", configure)

}
