package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/WillyWilsen/Deall-Jobs_Technical-Test.git/database"
	"github.com/WillyWilsen/Deall-Jobs_Technical-Test.git/utility"
	"github.com/WillyWilsen/Deall-Jobs_Technical-Test.git/repository"
	"github.com/WillyWilsen/Deall-Jobs_Technical-Test.git/handler"
)

func main() {
	utility.PrintConsole("API started", "info")
	utility.PrintConsole("Loading application configuration", "info")
	configuration, errConfig := utility.LoadApplicationConfiguration("")
	if errConfig != nil {
		log.WithFields(log.Fields{"error": errConfig}).Fatal("Failed to load app configuration")
	}
	utility.PrintConsole("Application configuration loaded successfully", "info")

	db, gormDB, err := database.Open(configuration)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Fatal("Failed to open database")
	}
	defer db.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))


	// Repository
	authRepository := repository.NewAuthRepository(gormDB)

	// Handler
	authHandler := handler.NewAuthHandler(authRepository)

	// Router
	api := e.Group("/api")
	api.POST("/register", authHandler.Register)
    api.POST("/login", authHandler.Login)

	e.Logger.Fatal(e.Start(":" + configuration.Http.HttpPort))
}