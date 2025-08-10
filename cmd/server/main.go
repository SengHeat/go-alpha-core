package main

import (
	"alpha-core/internal/config"
	"alpha-core/internal/database"
	"alpha-core/internal/handler"
	"alpha-core/internal/middleware"
	"alpha-core/internal/router"
	"alpha-core/internal/utils"
	"alpha-core/pkg/logger"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	//Step 1. load configuration before starting the server
	configure := config.Load()

	//Step 2. initialize logger
	log := logger.Init(configure)

	//Step 3. initialize database connection
	gormDB, err := database.InitializeDatabase(configure, log)

	if err != nil {
		log.ErrorLog(err.Error())
	}

	//Step 4. set up gin router
	if utils.IsProduction(configure, log) {
		gin.SetMode(gin.ReleaseMode)
	}

	//Step 5. initialize gin router
	routes := gin.New()
	routes.Use(gin.Logger())

	handlerInstand := handler.NewHandler(gormDB, configure, log)
	authMiddleware := middleware.NewJWTMiddleware(configure, log)
	apiRoute := router.NewRouter(routes, handlerInstand, authMiddleware)

	port := configure.AppPort
	address := fmt.Sprintf(":%s", port)

	if err := apiRoute.Run(address); err != nil {
		log.ErrorLog(err.Error())
	}
}
