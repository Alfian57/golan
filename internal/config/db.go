package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func LoadDB() {
	dbUsername := GetEnv("DB_USERNAME", "root")
	dbPassword := GetEnv("DB_PASSWORD", "")
	dbHost := GetEnv("DB_HOST", "127.0.0.1")
	dbPort := GetEnv("DB_PORT", "3306")
	dbName := GetEnv("DB_NAME", "golang")

	databaseConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)

	var err error
	DB, err = sqlx.Open("mysql", databaseConnection)
	if err != nil {
		Logger.DPanicln(fmt.Sprintf("failed to connect to database: %v", err))
	}
}
