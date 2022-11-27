package jobs

// ----------------------------------------------------------------
// Automatically generated  "/jobs/externalmessage.go"
// ----------------------------------------------------------------
// Package              : jobs
// Object 			    : ExternalMessage (externalmessage)
// Endpoint 	        : ExternalMessage (MessageID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/12/2021 at 09:31:50
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"strconv"
	"time"

	"github.com/mt1976/ebEstimates/application"
	core "github.com/mt1976/ebEstimates/core"
	dao "github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
)

// ExternalMessageJob defines the job properties, name, period etc..
func ExternalMessage_JobImpl() dm.JobDefinition {
	// ----------------------------------------------------------------
	// Job Definition
	// ----------------------------------------------------------------
	var j dm.JobDefinition
	j.ID = "ExternalMessage"
	j.Name = "ExternalMessage"
	j.Period = "*/2 * * * *"
	j.Description = "ExternalMessage processing"
	j.Type = core.HouseKeeping
	return j
}

// ExternalMessage_Run initiates and runs the job as per the period.
func ExternalMessage_RunImpl() (string, error) {

	//GEt a list of unprocessed messages from ExternalMessage
	noMessages, externalmessageList, _ := dao.ExternalMessage_GetActive()
	outmsg := "Active Messages = " + strconv.Itoa(noMessages)

	message := "Messages Outstanding = " + strconv.Itoa(noMessages)
	core.Notification_Normal(ExternalMessage_Job().Name, message)

	for messagePos := 0; messagePos < noMessages; messagePos++ {
		msg := externalmessageList[messagePos]

		now := time.Now()
		expires, _ := time.Parse(core.DATETIME, msg.MessageTimeout)
		if now.After(expires) {
			application.ExternalMessage_Notify(externalmessageList[messagePos])
			application.ExternalMessage_Expired(msg)
		}

	}

	return outmsg, nil
}
