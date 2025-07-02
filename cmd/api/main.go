package main

import (
	"fmt"

	"github.com/Alfian57/belajar-golang/internal/config"
	"github.com/Alfian57/belajar-golang/internal/database"
	"github.com/Alfian57/belajar-golang/internal/logger"
	"github.com/Alfian57/belajar-golang/internal/router"
	"github.com/Alfian57/belajar-golang/internal/validation"
)

func main() {
	config.LoadEnv()

	config, err := config.Load()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	database.Init(config.Database)
	logger.Init()
	validation.Init()

	router := router.NewRouter()
	appUrl := config.Server.Url
	logger.Log.Infoln(fmt.Sprintf("App running on %s", appUrl))
	router.Run(appUrl)
}
