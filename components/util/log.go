package util

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var GinLog *log.Logger
var RecLog *log.Logger
var AnyLog *log.Logger
var CompLog *log.Logger

func InitLogs(c Config) {
	const logLevel = "debug"
	if len(c.App().LogDirPath) == 0 {
		var lg = log.New()
		lg.Out = os.Stdout
		formatter := new(log.TextFormatter)
		formatter.ForceColors = true
		lg.Formatter = formatter

		GinLog = lg
		RecLog = lg
		AnyLog = lg
		CompLog = lg
	} else {
		GinLog = newFileLog(c.App().LogDirPath, logLevel, "gin.log")
		RecLog = newFileLog(c.App().LogDirPath, logLevel, "rec.log")
		AnyLog = newFileLog(c.App().LogDirPath, logLevel, "any.log")
		CompLog = newFileLog(c.App().LogDirPath, logLevel, "comp.log")
	}
}

func newFileLog(logDir, logLevel, logName string) *log.Logger {
	fileLog := &lumberjack.Logger{
		Filename:   logDir + logName,
		MaxSize:    5, // megabytes
		MaxBackups: 10,
		MaxAge:     28, //days
	}

	var lg = log.New()
	lg.Out = fileLog
	level, err := log.ParseLevel(logLevel)
	if err == nil {
		lg.Level = level
	}

	return lg
}
