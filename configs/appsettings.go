package configs

import (
	"fmt"
	"github.com/TechBuilder-360/business-directory-backend/common/types"
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/spf13/viper"
)

const (
	PRODUCTION types.ENVIRONMENT = "Production"
	SANDBOX    types.ENVIRONMENT = "SandBox"
)

var Instance *Config

type Config struct {
	AppName        string             `yaml:"AppName"`
	Host           string             `yaml:"Host"`
	Port           string             `yaml:"Port"`
	Issuer         string             `yaml:"Issuer"`
	Environment    *types.ENVIRONMENT `yaml:"Environment"`
	Secret         string             `yaml:"Secret"`
	URLPrefix      string             `yaml:"URLPrefix"`
	AllowedOrigin  []string           `yaml:"AllowedOrigin"`
	TrustedProxies []string           `yaml:"TrustedProxies"`
	TOKENLIFESPAN  uint               `yaml:"TOKENLIFESPAN"`
	DbName         string             `yaml:"DBNAME"`
	DbUser         string             `yaml:"DBUSER"`
	DbPass         string             `yaml:"DBPASS"`
	DbHost         string             `yaml:"DBHOST"`
	DbPort         uint               `yaml:"DBPORT"`
}

func Configuration() {
	var conf *Config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../..")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	conf.Secret = os.Getenv("SecretKey")

	envKey := []string{"SecretKey"}
	for _, k := range envKey {
		if os.Getenv(k) == "" {
			panic(fmt.Sprintf("Environment variable '%s' not set.", k))
		}
	}

	// Check environment mode
	if conf.Environment == nil {
		log.Error("Environment Mode not set")
	}

	Instance = conf
}
