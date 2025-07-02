package main

import (
	"fmt"

	"github.com/Alfian57/belajar-golang/internal/config"
	"github.com/Alfian57/belajar-golang/internal/router"
)

func main() {
	config.LoadEnv()
	config.LoadValidator()
	config.LoadLogger()
	config.LoadDB()

	router := router.NewRouter()
	appUrl := config.GetEnv("APP_URL", "localhost:8000")
	config.Logger.Infoln(fmt.Sprintf("App running on %s", appUrl))
	router.Run(appUrl)
}
