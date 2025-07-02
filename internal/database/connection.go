package database

import (
	"fmt"

	"github.com/Alfian57/belajar-golang/internal/config"
	"github.com/Alfian57/belajar-golang/internal/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Init(config config.DatabaseConfig) {
	databaseConnection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", config.Username, config.Password, config.Host, config.Port, config.Name)

	var err error
	DB, err = sqlx.Open("mysql", databaseConnection)
	if err != nil {
		messageLog := fmt.Sprintf("error opening database connection: %v", err)
		logger.Log.Errorln(messageLog)
	}

	if err = DB.Ping(); err != nil {
		messageLog := fmt.Sprintf("failed to ping database: %v", err)
		logger.Log.DPanicln(messageLog)
	} else {
		messageLog := "successfully connected to the database"
		logger.Log.Infoln(messageLog)
	}
}
