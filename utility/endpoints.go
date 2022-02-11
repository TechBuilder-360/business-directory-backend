package utility

import (
	"github.com/TechBuilder-360/business-directory-backend.git/configs"
)

type EndPoint struct {
	Method string
	BaseURL string
	Path string
}

func GetEndpoint(urlName string, config *configs.Config) EndPoint {

	switch urlName {

	}

	return EndPoint{}
}
