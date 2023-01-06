package logger

import (
	"log"
	"os"
)

type Logger struct {
	Info *log.Logger
	Err  *log.Logger
}

func NewLogger() *Logger {
	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}
	return &Logger{
		Info: log.New(file, "INFO:  ", log.Ldate|log.Ltime|log.Lshortfile),
		Err:  log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
