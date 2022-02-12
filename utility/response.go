package utility

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var once sync.Once

var (
	CLIENTERROR = "CLI001"
)

func GetCodeMsg(code string) string {
	once.Do(func() {
		viper.SetConfigName("response")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %w \n", err))
		}
	})

	return viper.GetString(code)
}
