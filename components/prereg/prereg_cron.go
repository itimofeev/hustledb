package prereg

import "github.com/itimofeev/hustledb/components/util"

func GetPreregCron() func() {
	return func() {
		util.CronLog.Debug("Run components.PreregService.ParsePreregInfo()")
		GetPreregService().ParsePreregInfo()
		util.CronLog.Debug("Run components.PreregService.ParsePreregInfo() completed")
	}
}
