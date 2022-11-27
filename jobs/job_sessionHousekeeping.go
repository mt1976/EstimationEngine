package jobs

import (
	application "github.com/mt1976/ebEstimates/application"
	core "github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
	"github.com/mt1976/ebEstimates/logs"
	cron "github.com/robfig/cron/v3"
)

func SessionHouseKeeping_Job() dm.JobDefinition {
	var j dm.JobDefinition
	j.ID = "SESSION"
	j.Name = "SESSION"
	j.Period = "*/15 * * * *"
	j.Description = "Session Management"
	j.Type = core.HouseKeeping
	return j
}

func SessionHouseKeeping_Register(c *cron.Cron) {
	application.Schedule_Register(SessionHouseKeeping_Job())
	c.AddFunc(SessionHouseKeeping_Job().Period, func() { SessionHouseKeeping_Run() })
}

// RunJobRollover is a Rollover function
func SessionHouseKeeping_Run() {
	logs.StartJob(SessionHouseKeeping_Job().Name)
	/// CONTENT STARTS

	// _, err := application.HousekeepSessionStore()
	// // handle the error if there is one
	// if err != nil {
	// 	log.Println(err)
	// 	panic(err)
	// }
	// message := ""
	// if err != nil {
	// 	message = err.Error()
	// }

	/// CONTENT ENDS
	application.Schedule_Update(SessionHouseKeeping_Job(), "")
	logs.EndJob(SessionHouseKeeping_Job().Name)
}
