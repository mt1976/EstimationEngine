package jobs

import (
	"fmt"
	"time"

	application "github.com/mt1976/ebEstimates/application"
	core "github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	"github.com/mt1976/ebEstimates/logs"
	cron "github.com/robfig/cron/v3"
)

func SessionHouseKeeping_Job() dm.JobDefinition {
	var j dm.JobDefinition
	j.ID = "SESSION"
	j.Name = "Session"
	j.Period = "*/5 * * * *"
	j.Description = "Session Management"
	j.Type = core.HouseKeeping
	return j
}

func SessionHouseKeeping_Register(c *cron.Cron) {
	application.Schedule_Register(SessionHouseKeeping_Job())
	c.AddFunc(SessionHouseKeeping_Job().Period, func() { SessionHouseKeeping_Run() })
}

// SessionHouseKeeping_Run is a Rollover function
func SessionHouseKeeping_Run() {
	logs.StartJob(SessionHouseKeeping_Job().Name)
	/// CONTENT STARTS

	message, err := Session_HouseKeeping()
	if err != nil {
		logs.Error("SessionHouseKeeping_Run", err)
	}

	/// CONTENT ENDS
	application.Schedule_Update(SessionHouseKeeping_Job(), message)
	logs.EndJob(SessionHouseKeeping_Job().Name)
}

func Session_HouseKeeping() (string, error) {
	// Delete all expired sessions
	//
	// Get All Sessions
	noSessions, sessions, err := dao.Session_GetList()
	if err != nil {
		logs.Error("Session_HouseKeeping", err)
		return "", err
	}
	// For each session
	noProcess := 0
	for _, session := range sessions {
		// If expired
		if session.Expiry != "" {
			expiry, timeerr := time.Parse(core.DATETIMEFORMATUSER, session.Expiry)
			if timeerr != nil {
				logs.Error("Session_HouseKeeping", timeerr)
				return "", timeerr
			}
			if expiry.Before(time.Now()) {
				// Delete the session
				dao.Session_Delete(session.Id)
				noProcess++
			}
		}
	}
	noActive := noSessions - noProcess
	// TODO: PLACE A CALL TO DOA.CACHE.HOUSEKEEP - This should resolve the issue.
	// TODO: CREATE A CACHE HOUSKEEPING JOB
	core.ApplicationCache = core.ApplicationCache.Housekeep()
	fmt.Printf("core.ApplicationCache.Stat(): %v\n", core.ApplicationCache.Stat())

	MSG := dao.Translate("SessionExpiries", "Expired %d of %d sessions, %d active")
	MSG = fmt.Sprintf(MSG, noProcess, noSessions, noActive)
	return MSG, nil
}
