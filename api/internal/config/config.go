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
	if host := os.Getenv("DB_HOST"); host == "" {
		return Config{}, fmt.Errorf("DB_HOST is not set")
	}

	if user := os.Getenv("DB_USER"); user == "" {
		return Config{}, fmt.Errorf("DB_USER is not set")
	}

	if password := os.Getenv("DB_PASSWORD"); password == "" {
		return Config{}, fmt.Errorf("DB_PASSWORD is not set")
	}

	if dbName := os.Getenv("DB_NAME"); dbName == "" {
		return Config{}, fmt.Errorf("DB_NAME is not set")
	}

	return Config{
		DB: types.DBConfig{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			Port:     port,
		},
	}, nil
}
