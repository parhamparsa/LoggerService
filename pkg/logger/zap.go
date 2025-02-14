package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitDefaultLogger(lvl zapcore.Level) {
	cfg := zap.NewProductionConfig()
	cfg.Level.SetLevel(lvl)
	cfg.DisableStacktrace = true
	cfg.DisableCaller = true
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := cfg.Build(zap.AddCallerSkip(1))
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
}
