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
// Date & Time		    : 04/03/2023 at 20:14:10
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	application "github.com/mt1976/ebEstimates/application"
	dm          "github.com/mt1976/ebEstimates/datamodel"
	logs        "github.com/mt1976/ebEstimates/logs"
		core        "github.com/mt1976/ebEstimates/core"

	cron        "github.com/robfig/cron/v3"
)

// DataJob defines the job properties, name, period etc..
func Data_Job() dm.JobDefinition {
	// ----------------------------------------------------------------
	// Job Definition
	// Create a new job definition in the Data_Impl.go file called Data_JobImpl()
	// ----------------------------------------------------------------
	var j dm.JobDefinition
	j.ID = "Data"
	j.Name = "Data"
	j.Period = "10 1 * * *"
	j.Description = "Data processing"
	j.Type = core.General
	j = Data_Job_impl(j)
	return j
}

func Data_Register(c *cron.Cron) {
	application.Schedule_Register(Data_Job())
	c.AddFunc(Data_Job().Period, func() { Data_Run() })
}

// Data_Run initiates and runs the job as per the period.
func Data_Run() {
	logs.StartJob(Data_Job().Name)
	message := ""
	/// CONTENT STARTS
	// Create a func in the Data_Impl.go file called Data_Run_impl() that returns string,error	
	message,err := Data_Run_impl()
	if err != nil {
		logs.Warning(err.Error())
	}
	/// CONTENT ENDS
	application.Schedule_Update(Data_Job(), message)
	logs.EndJob(Data_Job().Name)
}