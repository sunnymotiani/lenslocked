package models

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

//OPEN WILL OPEN A SQL CONNECTION WITH PROVIDED SQL DATABSE
//CALLERS OF OPEN NEED TO ENSURE THAT CONNECTION IS EVENTUALLY CLOSED VIA
//DB.close() METHOD

func Open(config PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", config.string())
	if err != nil {
		return nil, fmt.Errorf("Open: %w", err)
	}

	return db, nil
}

func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "localhost",
		Port:     "5433",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) string() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}
