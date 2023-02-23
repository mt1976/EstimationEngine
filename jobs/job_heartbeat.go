package jobs

import (
	"fmt"

	application "github.com/mt1976/ebEstimates/application"
	core "github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	"github.com/mt1976/ebEstimates/logs"
	"github.com/robfig/cron/v3"
)

func HeartBeat_Job() dm.JobDefinition {
	var j dm.JobDefinition
	j.ID = "Heartbeat"
	j.Name = "Heartbeat"
	j.Period = "*/15 * * * *"
	j.Description = "System Heartbeat, Uptime, etc"
	j.Type = core.HouseKeeping
	return j
}

func HeartBeat_Register(c *cron.Cron) {
	application.Schedule_Register(HeartBeat_Job())
	c.AddFunc(HeartBeat_Job().Period, func() { HeartBeat_Run() })
}

// RunJobHeartBeat is a HeartBeat function
func HeartBeat_Run() {

	logs.StartJob(HeartBeat_Job().Name)
	core.Log_uptime()

	MSG_TXT := dao.Translate("JobMessage", "Uptime = %v")
	message := fmt.Sprintf(MSG_TXT, core.ApplicationUpTime())

	application.Schedule_Update(HeartBeat_Job(), message)
	logs.EndJob(HeartBeat_Job().Name)

	logs.Break()
	logs.URI("http://localhost:" + core.ApplicationHTTPPort())
	logs.Break()

}
