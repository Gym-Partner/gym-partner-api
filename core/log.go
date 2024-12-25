package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

type Log struct {
	FilePath string
	*zap.Logger
}

func NewLog(file string, isTest bool) *Log {
	if !isTest {
		return &Log{
			FilePath: fmt.Sprintf("%s/logs", file),
		}
	} else {
		return &Log{
			FilePath: fmt.Sprintf("%s/logs-test", file),
		}
	}
}

func (l *Log) ChargeLog() {
	gin.DisableConsoleColor()
	os.Mkdir(l.FilePath, 0755)

	logDir := l.FilePath + "/" + time.Now().Format("20060102")
	logFilePath := logDir + "-session.log"

	// Gin's log creation
	f, _ := os.Create(fmt.Sprintf("%s/gin.log", l.FilePath))
	gin.DefaultWriter = io.MultiWriter(f)

	// API's log creation
	if fileExist(logFilePath) {
		os.Remove(logFilePath)
	}

	l.Logger = createLogger(logFilePath)
	defer l.Logger.Sync()
}

func createLogger(filePath string) *zap.Logger {
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.StacktraceKey = ""

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			filePath,
		},
		ErrorOutputPaths: []string{
			filePath,
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getegid(),
		},
	}

	return zap.Must(config.Build())
}

func fileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func (l *Log) Debug(msg string) {
	l.Logger.Debug(msg)
}

func (l *Log) Info(msg string) {
	l.Logger.Info(msg)
}

func (l *Log) Warn(msg string) {
	l.Logger.Warn(msg)
}

func (l *Log) Error(msg string) {
	l.Logger.Error(msg)
}
