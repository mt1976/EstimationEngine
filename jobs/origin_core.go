package jobs
// ----------------------------------------------------------------
// Automatically generated  "/jobs/origin.go"
// ----------------------------------------------------------------
// Package              : jobs
// Object 			    : Origin (origin)
// Endpoint 	        : Origin (OriginID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 07/02/2023 at 18:52:37
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	application "github.com/mt1976/ebEstimates/application"
	dm          "github.com/mt1976/ebEstimates/datamodel"
	logs        "github.com/mt1976/ebEstimates/logs"
		core        "github.com/mt1976/ebEstimates/core"

	cron        "github.com/robfig/cron/v3"
)

// OriginJob defines the job properties, name, period etc..
func Origin_Job() dm.JobDefinition {
	// ----------------------------------------------------------------
	// Job Definition
	// Create a new job definition in the Origin_Impl.go file called Origin_JobImpl()
	// ----------------------------------------------------------------
	var j dm.JobDefinition
	j.ID = "Origin"
	j.Name = "Origin"
	j.Period = "10 1 * * *"
	j.Description = "Origin processing"
	j.Type = core.General
	j = Origin_Job_impl(j)
	return j
}

func Origin_Register(c *cron.Cron) {
	application.Schedule_Register(Origin_Job())
	c.AddFunc(Origin_Job().Period, func() { Origin_Run() })
}

// Origin_Run initiates and runs the job as per the period.
func Origin_Run() {
	logs.StartJob(Origin_Job().Name)
	message := ""
	/// CONTENT STARTS
	// Create a func in the Origin_Impl.go file called Origin_Run_impl() that returns string,error	
	message,err := Origin_Run_impl()
	if err != nil {
		logs.Warning(err.Error())
	}
	/// CONTENT ENDS
	application.Schedule_Update(Origin_Job(), message)
	logs.EndJob(Origin_Job().Name)
}