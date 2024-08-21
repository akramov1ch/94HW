package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func ConnectPostgres() (*sqlx.DB, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_NAME"),
	)

	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	return db, nil
}
