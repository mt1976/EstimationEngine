package jobs
// ----------------------------------------------------------------
// Automatically generated  "/jobs/ssloader.go"
// ----------------------------------------------------------------
// Package              : jobs
// Object 			    : Ssloader (ssloader)
// Endpoint 	        : Ssloader (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 07/02/2023 at 18:52:39
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	application "github.com/mt1976/ebEstimates/application"
	dm          "github.com/mt1976/ebEstimates/datamodel"
	logs        "github.com/mt1976/ebEstimates/logs"
		core        "github.com/mt1976/ebEstimates/core"

	cron        "github.com/robfig/cron/v3"
)

// SsloaderJob defines the job properties, name, period etc..
func Ssloader_Job() dm.JobDefinition {
	// ----------------------------------------------------------------
	// Job Definition
	// Create a new job definition in the Ssloader_Impl.go file called Ssloader_JobImpl()
	// ----------------------------------------------------------------
	var j dm.JobDefinition
	j.ID = "Ssloader"
	j.Name = "Ssloader"
	j.Period = "10 1 * * *"
	j.Description = "Ssloader processing"
	j.Type = core.General
	j = Ssloader_Job_impl(j)
	return j
}

func Ssloader_Register(c *cron.Cron) {
	application.Schedule_Register(Ssloader_Job())
	c.AddFunc(Ssloader_Job().Period, func() { Ssloader_Run() })
}

// Ssloader_Run initiates and runs the job as per the period.
func Ssloader_Run() {
	logs.StartJob(Ssloader_Job().Name)
	message := ""
	/// CONTENT STARTS
	// Create a func in the Ssloader_Impl.go file called Ssloader_Run_impl() that returns string,error	
	message,err := Ssloader_Run_impl()
	if err != nil {
		logs.Warning(err.Error())
	}
	/// CONTENT ENDS
	application.Schedule_Update(Ssloader_Job(), message)
	logs.EndJob(Ssloader_Job().Name)
}