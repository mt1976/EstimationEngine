package jobs
// ----------------------------------------------------------------
// Automatically generated  "/jobs/credentials.go"
// ----------------------------------------------------------------
// Package              : jobs
// Object 			    : Credentials (credentials)
// Endpoint 	        : Credentials (Id)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 13/03/2023 at 14:22:25
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	application "github.com/mt1976/ebEstimates/application"
	dm          "github.com/mt1976/ebEstimates/datamodel"
	logs        "github.com/mt1976/ebEstimates/logs"
		core        "github.com/mt1976/ebEstimates/core"

	cron        "github.com/robfig/cron/v3"
)

// CredentialsJob defines the job properties, name, period etc..
func Credentials_Job() dm.JobDefinition {
	// ----------------------------------------------------------------
	// Job Definition
	// Create a new job definition in the Credentials_Impl.go file called Credentials_JobImpl()
	// ----------------------------------------------------------------
	var j dm.JobDefinition
	j.ID = "Credentials"
	j.Name = "Credentials"
	j.Period = "10 1 * * *"
	j.Description = "Credentials processing"
	j.Type = core.General
	j = Credentials_Job_impl(j)
	return j
}

func Credentials_Register(c *cron.Cron) {
	application.Schedule_Register(Credentials_Job())
	c.AddFunc(Credentials_Job().Period, func() { Credentials_Run() })
}

// Credentials_Run initiates and runs the job as per the period.
func Credentials_Run() {
	logs.StartJob(Credentials_Job().Name)
	message := ""
	/// CONTENT STARTS
	// Create a func in the Credentials_Impl.go file called Credentials_Run_impl() that returns string,error	
	message,err := Credentials_Run_impl()
	if err != nil {
		logs.Warning(err.Error())
	}
	/// CONTENT ENDS
	application.Schedule_Update(Credentials_Job(), message)
	logs.EndJob(Credentials_Job().Name)
}