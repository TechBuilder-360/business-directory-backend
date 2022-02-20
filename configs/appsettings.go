package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	AppName          string `yaml:"AppName"`
	Host             string `yaml:"Host"`
	Port             string `yaml:"Port"`
	DEBUG            bool   `yaml:"DEBUG"`
	MongoURI         string `yaml:"MongoURI"`
	AesKey           string `yaml:"AesKey"`
	MongoDBName      string `yaml:"MongoDBName"`
	ClientCollection string `yaml:"ClientCollection"`
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
	return conf
}
