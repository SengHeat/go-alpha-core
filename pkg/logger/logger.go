package logger

import (
	"alpha-core/internal/config"
	"fmt"
	"log"
	"os"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

type Logger struct{}

func Init(configure *config.Config) *Logger {
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	_ = configure
	return &Logger{}
}

func (logger *Logger) InfoLog(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Println(string(Green) + "[INFO] : " + msg + string(Reset))
}

func (logger *Logger) WarningLog(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Println(string(Yellow) + "[WARN] : " + msg + string(Reset))
}

func (logger *Logger) ErrorLog(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Println(string(Red) + "[FAIL] : " + msg + string(Reset))
	os.Exit(1)
}
