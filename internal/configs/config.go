// Package configs provides configuration structures for the application.
package configs

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBURL string `envconfig:"DB_URL" required:"true"`
	Port  string `envconfig:"SERVER_PORT" required:"true"`
}

var Env Config

func StartConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	if err := envconfig.Process("", &Env); err != nil {
		return err
	}

	return nil
}
