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
	application "github.com/mt1976/ebEstimates/application"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
	cron "github.com/robfig/cron/v3"
)

// ExternalMessageJob defines the job properties, name, period etc..
func ExternalMessage_Job() dm.JobDefinition {
	// ----------------------------------------------------------------
	// Job Definition
	// Create a new job definition in the ExternalMessage_Impl.go file called ExternalMessage_JobImpl()
	// ----------------------------------------------------------------
	// j.ID = "ExternalMessage"
	// j.Name = "ExternalMessage"
	// j.Period = "10 1 * * *"
	// j.Description = "ExternalMessage processing"
	// j.Type = core.HouseKeeping
	j := ExternalMessage_JobImpl()
	return j
}

func ExternalMessage_Register(c *cron.Cron) {
	application.Schedule_Register(ExternalMessage_Job())
	c.AddFunc(ExternalMessage_Job().Period, func() { ExternalMessage_Run() })
}

// ExternalMessage_Run initiates and runs the job as per the period.
func ExternalMessage_Run() {
	logs.StartJob(ExternalMessage_Job().Name)
	message := ""
	/// CONTENT STARTS
	// Create a func in the ExternalMessage_Impl.go file called ExternalMessage_RunImpl() that returns string,error
	message, err := ExternalMessage_RunImpl()
	if err != nil {
		logs.Warning(err.Error())
	}
	/// CONTENT ENDS
	application.Schedule_Update(ExternalMessage_Job(), message)
	logs.EndJob(ExternalMessage_Job().Name)
}
