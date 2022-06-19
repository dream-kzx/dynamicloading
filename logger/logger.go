package logger

import (
	"log"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type DefaultLogger struct {
}

func (d DefaultLogger) Debug(msg string) {
	log.Printf("Debug | %s", msg)
}

func (d DefaultLogger) Info(msg string) {
	log.Printf("Info  | %s", msg)
}
func (d DefaultLogger) Warn(msg string) {
	log.Printf("Warn  | %s", msg)
}
func (d DefaultLogger) Error(msg string) {
	log.Printf("Error | %s", msg)
}
