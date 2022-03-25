package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	AppName          string `yaml:"AppName"`
	Host             string `yaml:"Host"`
	Port             string `yaml:"Port"`
	DEBUG            bool   `yaml:"DEBUG"`
	MongoURI         string `yaml:"MongoURI"`
	Secret			 string  `yaml:"Secret"`
	URLPrefix        string  `yaml:"URLPrefix"`
	AesKey           string `yaml:"AesKey"`
	MongoDBName      string `yaml:"MongoDBName"`
	AllowedOrigin    []string `yaml:"AllowedOrigin"`
	TrustedProxies   []string `yaml:"TrustedProxies"`
}

func Configuration() *Config {
	var conf *Config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	conf.MongoURI = os.Getenv("MongoURI")
	conf.Secret = os.Getenv("Secret")

	envKey := []string{"Secret", "MongoURI"}
	for _, k := range envKey {
		if os.Getenv(k) == "" {
			panic(fmt.Sprintf("Environment variable '%s' not set.", k))
		}
	}

	return conf
}
