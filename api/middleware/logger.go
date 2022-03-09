package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"webxam/config"
)

var l *zap.Logger

func Logger(c *gin.Context) {
	var err error
	config := config.GetConfig()

	switch config.GetString("environment") {
	case "development":
		// level := zap.NewAtomicLevel()
		// level.SetLevel(zapcore.DebugLevel)
		// conf := zap.Config{
		// 	Level:            level,
		// 	Encoding:         "json",
		// 	OutputPaths:      []string{"stdout", "./logs/hoge.log"},
		// 	ErrorOutputPaths: []string{"stderr"},
		// }
		l, err = zap.NewDevelopment()
		if err != nil {

		}
	case "production":

	}
}

func GetLogger() *zap.SugaredLogger {
	return l.Sugar()
}
