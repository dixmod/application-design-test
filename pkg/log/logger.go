package log

import (
	"fmt"
	"log"
)

func NewLogger() LocalLogger {
	return &AppLogger{
		log.Default(),
	}
}

type LocalLogger interface {
	LogErrorf(format string, v ...any)
	LogInfo(format string, v ...any)
}

type AppLogger struct {
	logger *log.Logger
}

func (l *AppLogger) LogErrorf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	l.logger.Printf("[Error]: %s\n", msg)
}

func (l *AppLogger) LogInfo(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	l.logger.Printf("[Info]: %s\n", msg)
}
