package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/wybbb1/SiMo/internal/config"
)

var Logger *zap.Logger

func SetupLogger() {
	var cfg zap.Config
	if config.Config.Logging.Mode == "development" {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(config.Config.Logging.TimeLayout) // 格式化时间
	cfg.OutputPaths = config.Config.Logging.Path

	var err error
	Logger, err = cfg.Build()
	if err != nil {
		panic(err)
	}
}

func Sync() {
	if Logger != nil {
		Logger.Sync()
	}
}