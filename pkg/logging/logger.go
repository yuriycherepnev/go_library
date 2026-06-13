package logger

import (
	"log"
	"os"
)

type Logger struct {
	Info  *log.Logger
	Error *log.Logger
}

var instance *Logger

func GetLogger() *Logger {
	if instance == nil {
		instance = &Logger{
			Info:  log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime),
			Error: log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
		}
	}
	return instance
}
