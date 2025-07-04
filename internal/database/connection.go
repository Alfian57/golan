package database

import (
	"fmt"
	"time"

	"github.com/Alfian57/belajar-golang/internal/config"
	"github.com/Alfian57/belajar-golang/internal/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Init(config config.DatabaseConfig) {
	// Initialize the database connection string
	databaseConnection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", config.Username, config.Password, config.Host, config.Port, config.Name)

	// Open a new database connection
	var err error
	DB, err = sqlx.Open("mysql", databaseConnection)
	if err != nil {
		logger.Log.Fatalf("error opening database connection: %v", err)
	}

	// Configure connection pool
	DB.SetMaxOpenConns(25)                 // Maximum number of open connections
	DB.SetMaxIdleConns(25)                 // Maximum number of idle connections
	DB.SetConnMaxLifetime(5 * time.Minute) // Maximum connection lifetime

	// Ping the database to ensure the connection is established
	if err = DB.Ping(); err != nil {
		logger.Log.Fatalf("failed to ping database: %v", err)
	}

	logger.Log.Infoln("successfully connected to the database")
}
