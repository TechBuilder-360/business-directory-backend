package utility

import (
	"net/http"

	"github.com/Toflex/oris_log/logger"
)

type RequestBody struct{
Code string
Response string
Request string
Url string
Method string
}

type ExternalConsole struct{
	Logger logger.Logger
}


var l *ExternalConsole
func Get(url string) *http.Response{

	resp,err:=http.Get(url)
if err!=nil{
	l.Logger.Error(err.Error())
}

	return resp
}