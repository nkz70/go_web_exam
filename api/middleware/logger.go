package middleware

import (
	"go.uber.org/zap"
	"github.com/gin-gonic/gin"

	"webxam/config"
)

var l *zap.Logger

func Logger(c *gin.Context) {
	var err error
	config := config.GetConfig()

	switch config.GetString("environment") {
	case "development":
		l, err = zap.NewDevelopment()
		if err != nil {

		}
	case "production":

	}
}

func GetLogger() *zap.SugaredLogger {
	return l.Sugar()
}
