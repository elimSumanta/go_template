package logger

import (
	"fmt"
	"os"

	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
)

type LoggerStr struct {
	Log *rkentry.LoggerEntry
}

var log *LoggerStr

func LoggerInit() {
	logRK := rkentry.GlobalAppCtx.GetLoggerEntry("my-logger")

	log = &LoggerStr{
		Log: logRK,
	}
}

func Infof(message string, args ...interface{}) {
	msg := fmt.Sprintf(message, args...)
	log.Log.Info(msg)
}

func Debugf(message string, args ...interface{}) {
	msg := fmt.Sprintf(message, args...)
	log.Log.Debug(msg)
}

func Errorf(message string, args ...interface{}) {
	msg := fmt.Sprintf(message, args...)
	log.Log.Error(msg)
}

func Fatal(message string, args ...interface{}) {
	msg := fmt.Sprintf(message, args...)
	log.Log.Error(msg)
	os.Exit(1)
}
