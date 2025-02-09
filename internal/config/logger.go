package config

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

const (
	loggerLevelEnvKey = "LOGGER_LEVEL"
)

type LoggerConfig interface {
	GetCore() zapcore.Core
}

type loggerConfig struct {
	level zapcore.Level
}

func NewLoggerConfig() (LoggerConfig, error) {
	loglevel := os.Getenv(loggerLevelEnvKey)
	if len(loglevel) == 0 {
		loglevel = "info"
	}

	var level zapcore.Level
	if err := level.Set(loglevel); err != nil {
		return nil, err
	}

	return &loggerConfig{
		level: level,
	}, nil
}

func getAtomicLevel(loglevel string) (zap.AtomicLevel, error) {
	var level zapcore.Level
	if err := level.Set(loglevel); err != nil {
		return zap.NewAtomicLevel(), err
	}

	return zap.NewAtomicLevelAt(level), nil
}

func (l *loggerConfig) GetCore() zapcore.Core {
	stdout := zapcore.AddSync(os.Stdout)

	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10, //Mb
		MaxBackups: 3,
		MaxAge:     7, //days
	})

	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.TimeKey = "timestamp"
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
	fileEncoder := zapcore.NewJSONEncoder(productionCfg)

	return zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, l.level),
		zapcore.NewCore(fileEncoder, file, l.level),
	)
}
