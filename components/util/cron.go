package util

import (
	"github.com/robfig/cron"
)

func InitCronTasks() {
	c := cron.New()
	//c.AddFunc("@every 2h30m", func() {
	c.AddFunc("@every 10s", func() {
		CronLog.Debug("Hello, there :)")
	})
	c.Start()
}
