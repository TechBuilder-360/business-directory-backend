package utility

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var once sync.Once

var (
	CLIENTERROR        = "CLI001"
	AUTHERROR004       = "AUTHERR004"
	SYSTEM001          = "SUC001"
	SECURITYDECRYPTERR = "SECURITYDECRYPTERR"
	Authorization = "AuthERR001"
	VALIDATIONERR = "VALIDATIONERR"
	UNAUTHORISE        = "UNAUTHORISE"
	BAD_REQUEST = "BAD_REQUEST"
	SERVER_ERROR = "SERVER_ERROR"
	SMMERROR004 = "SMMERR"
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
