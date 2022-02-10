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
	MongoDBName      string `yaml:"MongoDBName"`
	ClientCollection string `yaml:"ClientCollection"`
}

func Configuration() *Config {
	var conf *Config
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")     // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	return conf
}
