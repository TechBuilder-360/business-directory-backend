package apps

import (
	"github.com/Toflex/oris_log/logger"
	"github.com/gin-gonic/gin"
)

type App struct {
	srv interface{}
	repo interface{}
	Router *gin.Engine
	Logger logger.Logger
}