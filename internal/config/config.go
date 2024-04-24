package config

import (
	"github.com/akram620/alif/internal/logger"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type values struct {
	APIPort     string `envconfig:"API_PORT" required:"true"`
	DatabaseURL string `envconfig:"DB_URL" required:"true"`
}

var Values values

func LoadFromFile(fpath string) error {
	godotenv.Load(fpath)

	err := envconfig.Process("", &Values)
	if err != nil {
		logger.Errorf("envconfig.Process(): %v", err.Error())
	}

	return err
}
