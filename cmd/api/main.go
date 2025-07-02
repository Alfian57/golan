package main

import (
	"fmt"

	"github.com/Alfian57/belajar-golang/internal/config"
	"github.com/Alfian57/belajar-golang/internal/database"
	"github.com/Alfian57/belajar-golang/internal/logger"
	"github.com/Alfian57/belajar-golang/internal/middleware"
	"github.com/Alfian57/belajar-golang/internal/router"
	"github.com/Alfian57/belajar-golang/internal/validation"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	config, err := config.Load()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	logger.Init()
	database.Init(config.Database)
	validation.Init()

	router := router.NewRouter()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.ErrorMiddleware())
	router.SetTrustedProxies(config.Server.TrustedProxies)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     config.Cors.AllowOrigins,
		AllowMethods:     config.Cors.AllowMethods,
		AllowCredentials: config.Cors.AllowCredentials,
	}))

	appUrl := config.Server.Url
	router.Run(appUrl)
}
