package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"webexam/config"
)

var L *zap.Logger
var S *zap.SugaredLogger

func Logger(c *gin.Context) {
	var err error
	config := config.GetConfig()

	switch config.GetString("environment") {
	case "development":
		level := zap.NewAtomicLevel()
		level.SetLevel(zapcore.DebugLevel)
		conf := zap.Config{
			Level:    level,
			Encoding: "console",
			EncoderConfig: zapcore.EncoderConfig{
				TimeKey:        "Time",
				LevelKey:       "Level",
				NameKey:        "Name",
				CallerKey:      "Caller",
				MessageKey:     "Msg",
				StacktraceKey:  "St",
				EncodeLevel:    zapcore.CapitalLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.StringDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			},
			OutputPaths:      []string{"stdout", "./logs/hoge.log"},
			ErrorOutputPaths: []string{"stderr"},
		}
		L, err = conf.Build()
		if err != nil {
			log.Fatal(err)
		}
		S = L.Sugar()
	case "production":
	}
}
