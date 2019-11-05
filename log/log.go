package log

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"os"
)

var logger Logger = defaultLogger{}

// Logger describes what a Logger should support
type Logger interface {
	Info(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
}

func init() {
	// TODO: allow to configure the logger
	lg, err := zap.NewDevelopment()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to instatiate the zap logger: %s, will use default logger\n", err)
	} else {
		logger = lg.WithOptions(zap.AddCallerSkip(1))
	}
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

type defaultLogger struct{}

func (d defaultLogger) fmt(level string, msg string, fields ...zap.Field) string {
	return fmt.Sprintf("%s: %s: %v", level, msg, fields)
}

func (d defaultLogger) Fatal(msg string, fields ...zap.Field) {
	log.Fatalf(d.fmt("FATAL", msg, fields...))
}

func (d defaultLogger) Info(msg string, fields ...zap.Field) {
	log.Printf(d.fmt("INFO", msg, fields...))
}

func (d defaultLogger) Debug(msg string, fields ...zap.Field) {
	log.Printf(d.fmt("DEBUG", msg, fields...))
}

func (d defaultLogger) Error(msg string, fields ...zap.Field) {
	log.Printf(d.fmt("ERROR", msg, fields...))
}

var _ Logger = defaultLogger{}
