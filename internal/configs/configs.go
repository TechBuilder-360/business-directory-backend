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

	SendGridAPIKey    string `env:"SENDGRID_API_KEY"`
	SendGridFromEmail string `env:"SEND_GRID_FROM_EMAIL"`

	PaperTailAppName string `env:"PAPER_TAIL_APP_NAME"`
	PaperTailPort    int    `env:"PAPER_TAIL_PORT"`

	EnvCloudName         string `env:"ENV_CLOUD_NAME"`
	EnvCloudAPIKey       string `env:"ENV_CLOUD_API_KEY"`
	EnvCloudAPISecret    string `env:"ENV_CLOUD_API_SECRET"`
	EnvCloudUploadFolder string `env:"ENV_CLOUD_UPLOAD_FOLDER"`
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
