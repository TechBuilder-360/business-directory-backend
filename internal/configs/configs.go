package configs

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"go.deanishe.net/env"
	"strings"
)

const (
	PRODUCTION  types.ENVIRONMENT = "PRODUCTION"
	DEVELOPMENT types.ENVIRONMENT = "DEVELOPMENT"
	SANDBOX     types.ENVIRONMENT = "SANDBOX"
)

var Instance *Config

type Config struct {
	AppName       string            `env:"APP_NAME"`
	Namespace     string            `env:"NAMESPACE"`
	BASEURL       string            `env:"BASE_URL"`
	Port          string            `env:"PORT"`
	Environment   types.ENVIRONMENT `env:"ENVIRONMENT"`
	Secret        string            `env:"SECRET"`
	TOKENLIFESPAN uint              `env:"TOKEN_LIFE_SPAN"`

	DbName string `env:"DB_NAME"`
	DbUser string `env:"DB_USER"`
	DbPass string `env:"DB_PASS"`
	DbHost string `env:"DB_HOST"`
	DbPort uint   `env:"DB_PORT"`

	RedisURL          string `env:"REDIS_URL"`
	RedisDB           int    `env:"REDIS_DB"`
	RedisPassword     string `env:"REDIS_PASSWORD"`
	RedisCacheRefresh string `env:"REDIS_CACHE_REFRESH"`

	SendGridAPIKey    *string `env:"SENDGRID_API_KEY"`
	SendGridFromEmail *string `env:"SEND_GRID_FROM_EMAIL"`

	PaperTailAppName *string `env:"PAPER_TAIL_APP_NAME"`
	PaperTailPort    *string `env:"PAPER_TAIL_PORT"`

	CloudinaryName   *string `env:"ENV_CLOUD_NAME"`
	CloudinaryAPIKey *string `env:"ENV_CLOUD_API_KEY"`
	CloudinarySecret *string `env:"ENV_CLOUD_API_SECRET"`

	AuthServerBaseURL   *string `env:"AUTH_SERVER_BASE_URL"`
	AuthServerSecretKey *string `env:"AUTH_SERVER_SECRET_KEY"`
}

func Load() {
	c := &Config{}
	if err := env.Bind(c); err != nil {
		panic(err.Error())
	}
	Instance = c
	return
}

func (c *Config) GetEnv() types.ENVIRONMENT {
	return types.ENVIRONMENT(strings.ToUpper(string(Instance.Environment)))
}

func (c *Config) IsSandbox() bool {
	return types.ENVIRONMENT(strings.ToUpper(string(c.Environment))) == SANDBOX
}

func (c *Config) IsProduction() bool {
	return types.ENVIRONMENT(strings.ToUpper(string(c.Environment))) == PRODUCTION
}
