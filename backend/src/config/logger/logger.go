package logger

/**
 * Package to record information and errors in a more readable way in log files of the day
 */

import (
	"fmt"
	"os"
	"strings"
	"time"

	_log "log"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger

	LOG_OUTPUT = "LOG_OUTPUT"
	LOG_LEVEL  = "LOG_LEVEL"
)

func Init() {
	logConfig := zap.Config{
		OutputPaths: getOutputLogs(),
		Level:       zap.NewAtomicLevelAt(getLevelLogs()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ = logConfig.Build()
}

func Info(message string, tags ...zap.Field) {
	if log != nil {
		log.Info(message, tags...)
		log.Sync()
	}
}

func Error(message string, err error, tags ...zap.Field) {
	if log != nil {
		tags = append(tags, zap.NamedError("error", err))
		log.Info(message, tags...)
		log.Sync()
	}
}

func getOutputLogs() []string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dir := os.Getenv("LOG_OUTPUT")

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			_log.Println("Error creating output directory: ", err)
		}
	}

	if dir == "" {
		return []string{"stdout"}
	}

	output := fmt.Sprintf("%s/BACK-%s.log", dir, time.Now().Format("2006-01-02"))
	return []string{"stdout", output}
}

func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv("LOG_LEVEL"))) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}
