package jobs
// ----------------------------------------------------------------
// Automatically generated  "/jobs/ssloader.go"
// ----------------------------------------------------------------
// Package              : jobs
// Object 			    : ssloader (ssloader)
// Endpoint 	        : Ssloader (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 08/03/2023 at 18:42:26
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	application "github.com/mt1976/ebEstimates/application"
	dm          "github.com/mt1976/ebEstimates/datamodel"
	logs        "github.com/mt1976/ebEstimates/logs"
		core        "github.com/mt1976/ebEstimates/core"

	cron        "github.com/robfig/cron/v3"
)

// ssloaderJob defines the job properties, name, period etc..
func ssloader_Job() dm.JobDefinition {
	// ----------------------------------------------------------------
	// Job Definition
	// Create a new job definition in the ssloader_Impl.go file called ssloader_JobImpl()
	// ----------------------------------------------------------------
	var j dm.JobDefinition
	j.ID = "ssloader"
	j.Name = "ssloader"
	j.Period = "10 1 * * *"
	j.Description = "ssloader processing"
	j.Type = core.General
	j = ssloader_Job_impl(j)
	return j
}

func ssloader_Register(c *cron.Cron) {
	application.Schedule_Register(ssloader_Job())
	c.AddFunc(ssloader_Job().Period, func() { ssloader_Run() })
}

// ssloader_Run initiates and runs the job as per the period.
func ssloader_Run() {
	logs.StartJob(ssloader_Job().Name)
	message := ""
	/// CONTENT STARTS
	// Create a func in the ssloader_Impl.go file called ssloader_Run_impl() that returns string,error	
	message,err := ssloader_Run_impl()
	if err != nil {
		logs.Warning(err.Error())
	}
	/// CONTENT ENDS
	application.Schedule_Update(ssloader_Job(), message)
	logs.EndJob(ssloader_Job().Name)
}