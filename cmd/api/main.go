package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	logger.Init()
	database.Init(cfg.Database)
	validation.Init()

	router := router.NewRouter()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.ErrorMiddleware())

	if err := router.SetTrustedProxies(cfg.Server.TrustedProxies); err != nil {
		panic(fmt.Sprintf("Failed to set trusted proxies: %v", err))
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.Cors.AllowOrigins,
		AllowMethods:     cfg.Cors.AllowMethods,
		AllowCredentials: cfg.Cors.AllowCredentials,
	}))

	server := &http.Server{
		Addr:    cfg.Server.Url,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		logger.Log.Infof("Server starting on %s", cfg.Server.Url)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Log.Info("Server shutting down...")

	// Give the server 5 seconds to shutdown gracefully
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Log.Fatalf("Server forced to shutdown: %v", err)
	}

	logger.Log.Info("Server exited")
}
