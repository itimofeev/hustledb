package comp

import "github.com/itimofeev/hustledb/components/util"

func GetCompCron() func() {
	return func() {
		util.CronLog.Debug("Run GetFCompService().ParseCompetitions()")
		GetFCompService().ParseCompetitions()
		util.CronLog.Debug("Run GetFCompService().ParseCompetitions() completed")
	}
}
