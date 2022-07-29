package configs

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"go.deanishe.net/env"
)

const (
	PRODUCTION types.ENVIRONMENT = "Production"
	SANDBOX    types.ENVIRONMENT = "SandBox"
)

var Instance *Config

type Config struct {
	AppName       string            `env:"APP_NAME"`
	Namespace     string            `env:"NAMESPACE"`
	Host          string            `env:"BASE_URL"`
	Issuer        string            `env:"PORT"`
	Environment   types.ENVIRONMENT `env:"ENVIRONMENT"`
	Secret        string            `env:"SECRET"`
	URLPrefix     string            `env:"ISSUER"`
	TOKENLIFESPAN uint              `env:"TOKEN_LIFE_SPAN"`

	DbName string `env:"DB_NAME"`
	DbUser string `env:"DB_USER"`
	DbPass string `env:"DB_PASS"`
	DbHost string `env:"DB_HOST"`
	DbPort uint   `env:"DB_PORT"`

	RedisURL      string `env:"REDIS_URL"`
	RedisPassword string `env:"REDIS_PASSWORD"`
}

func Load() {
	c := &Config{}
	if err := env.Bind(c); err != nil {
		panic(err.Error())
	}
	Instance = c
	return
}

func GetEnv() types.ENVIRONMENT {
	return Instance.Environment
}
