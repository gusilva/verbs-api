package pkg

import "log"

type ILogger interface {
	Info(v ...interface{})
	Error(v ...interface{})
	Debug(v ...interface{})
}

type Logger struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	DebugLogger *log.Logger
	IsDevMode   bool
}

func (l *Logger) Info(v ...interface{}) {
	l.InfoLogger.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.ErrorLogger.Println(v...)
}

func (l *Logger) Debug(v ...interface{}) {
	if l.IsDevMode {
		l.DebugLogger.Println(v...)
	}
}
