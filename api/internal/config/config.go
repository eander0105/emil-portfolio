package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/eander0105/emil-portfolio/api/internal/types"
)

type Config struct {
	DB types.DBConfig
}

func LoadConfig() (Config, error) {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return Config{}, err
	}

	// Ensure values are set
	host := os.Getenv("DB_HOST")
	if host == "" {
		return Config{}, fmt.Errorf("DB_HOST is not set")
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		return Config{}, fmt.Errorf("DB_USER is not set")
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		return Config{}, fmt.Errorf("DB_PASSWORD is not set")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return Config{}, fmt.Errorf("DB_NAME is not set")
	}

	return Config{
		DB: types.DBConfig{
			Host:     host,
			User:     user,
			Password: password,
			DBName:   dbName,
			Port:     port,
		},
	}, nil
}
