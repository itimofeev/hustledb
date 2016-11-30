package components

import (
	"github.com/itimofeev/hustledb/components/forum/comp"
	"github.com/itimofeev/hustledb/components/prereg"
	"github.com/robfig/cron"
)

func InitCronTasks() {
	c := cron.New()
	c.AddFunc("@every 1h", prereg.GetPreregCron())
	c.AddFunc("@every 1h", comp.GetCompCron())
	c.Start()
}
