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
var CronLog *log.Logger

func InitLogs(c Config) {
	var logLevel = log.DebugLevel.String()
	logDirPath := c.App().LogDirPath
	if len(logDirPath) == 0 {
		var lg = log.New()
		lg.Out = os.Stdout
		lg.Level = log.DebugLevel

		GinLog = lg
		RecLog = lg
		AnyLog = lg
		CompLog = lg
		CronLog = lg
	} else {
		GinLog = newFileLog(logDirPath, logLevel, "gin.log")
		RecLog = newFileLog(logDirPath, logLevel, "rec.log")
		AnyLog = newFileLog(logDirPath, logLevel, "any.log")
		CompLog = newFileLog(logDirPath, logLevel, "comp.log")
		CronLog = newFileLog(logDirPath, logLevel, "cron.log")
	}

	AnyLog.Debug("Lets start fun with hustledb :)")
}

func newFileLog(logDir, logLevel, logName string) *log.Logger {
	fileLog := &lumberjack.Logger{
		Filename:   logDir + "/" + logName,
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
