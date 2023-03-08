package jobs
// ----------------------------------------------------------------
// Automatically generated  "/jobs/index.go"
// ----------------------------------------------------------------
// Package              : jobs
// Object 			    : Index (index)
// Endpoint 	        : Index (IndexID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 04/03/2023 at 20:14:12
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	application "github.com/mt1976/ebEstimates/application"
	dm          "github.com/mt1976/ebEstimates/datamodel"
	logs        "github.com/mt1976/ebEstimates/logs"
		core        "github.com/mt1976/ebEstimates/core"

	cron        "github.com/robfig/cron/v3"
)

// IndexJob defines the job properties, name, period etc..
func Index_Job() dm.JobDefinition {
	// ----------------------------------------------------------------
	// Job Definition
	// Create a new job definition in the Index_Impl.go file called Index_JobImpl()
	// ----------------------------------------------------------------
	var j dm.JobDefinition
	j.ID = "Index"
	j.Name = "Index"
	j.Period = "10 1 * * *"
	j.Description = "Index processing"
	j.Type = core.General
	j = Index_Job_impl(j)
	return j
}

func Index_Register(c *cron.Cron) {
	application.Schedule_Register(Index_Job())
	c.AddFunc(Index_Job().Period, func() { Index_Run() })
}

// Index_Run initiates and runs the job as per the period.
func Index_Run() {
	logs.StartJob(Index_Job().Name)
	message := ""
	/// CONTENT STARTS
	// Create a func in the Index_Impl.go file called Index_Run_impl() that returns string,error	
	message,err := Index_Run_impl()
	if err != nil {
		logs.Warning(err.Error())
	}
	/// CONTENT ENDS
	application.Schedule_Update(Index_Job(), message)
	logs.EndJob(Index_Job().Name)
}