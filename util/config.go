package util

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDriver          string
	DBSource          string
	MigrationURL      string
	HTTPServerAddress string
}

func LoadConfig(path string) (*Config, error) {

	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	return &Config{
		DBDriver:          os.Getenv("DB_DRIVER"),
		DBSource:          os.Getenv("DB_SOURCE"),
		HTTPServerAddress: os.Getenv("HTTP_SERVER_ADDRESS"),
		MigrationURL:      os.Getenv("MIGRATION_URL"),
	}, nil

}
