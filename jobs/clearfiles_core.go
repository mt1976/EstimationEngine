package jobs

// ----------------------------------------------------------------
// Automatically generated  "/jobs/data.go"
// ----------------------------------------------------------------
// Package              : jobs
// Object 			    : Data (data)
// Endpoint 	        : Data (DataID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 13/03/2023 at 14:22:26
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	application "github.com/mt1976/ebEstimates/application"
	core "github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"

	cron "github.com/robfig/cron/v3"
)

// DataJob defines the job properties, name, period etc..
func ClearFiles_Job() dm.JobDefinition {
	// ----------------------------------------------------------------
	// Job Definition
	// Create a new job definition in the ClearFiles_Impl.go file called ClearFiles_JobImpl()
	// ----------------------------------------------------------------
	var j dm.JobDefinition
	j.ID = "Data"
	j.Name = "Data"
	j.Period = "10 1 * * *"
	j.Description = "Data processing"
	j.Type = core.General
	j = ClearFiles_Job_impl(j)
	return j
}

func ClearFiles_Register(c *cron.Cron) {
	application.Schedule_Register(ClearFiles_Job())
	c.AddFunc(ClearFiles_Job().Period, func() { ClearFiles_Run() })
}

// ClearFiles_Run initiates and runs the job as per the period.
func ClearFiles_Run() {
	logs.StartJob(ClearFiles_Job().Name)
	message := ""
	/// CONTENT STARTS
	// Create a func in the ClearFiles_Impl.go file called ClearFiles_Run_impl() that returns string,error
	message, err := ClearFiles_Run_impl()
	if err != nil {
		logs.Warning(err.Error())
	}
	/// CONTENT ENDS
	application.Schedule_Update(ClearFiles_Job(), message)
	logs.EndJob(ClearFiles_Job().Name)
}
