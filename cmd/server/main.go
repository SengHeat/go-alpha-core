package main

import (
	"alpha-core/internal/config"
	"alpha-core/internal/database"
	"alpha-core/internal/utils"
	"alpha-core/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	//Step 1. load configuration before starting the server
	configure := config.Load()

	//Step 2. initialize logger
	log := logger.Init(configure)

	//Step 3. initialize database connection
	_, err := database.InitializeDatabase(configure, log)

	if err != nil {
		log.ErrorLog(err.Error())
	}

	//Step 4. set up gin router
	if utils.IsProduction(configure, log) {
		log.InfoLog("Running in production mode")
		gin.SetMode(gin.ReleaseMode)
	}

	//Step 5. initialize gin router
	_ = gin.New()

	//handler := Handle.N

}
