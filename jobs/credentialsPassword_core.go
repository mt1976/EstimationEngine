package jobs
// ----------------------------------------------------------------
// Automatically generated  "/jobs/credentialspassword.go"
// ----------------------------------------------------------------
// Package              : jobs
// Object 			    : CredentialsPassword (credentialspassword)
// Endpoint 	        : CredentialsPassword (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/03/2023 at 19:24:47
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	application "github.com/mt1976/ebEstimates/application"
	dm          "github.com/mt1976/ebEstimates/datamodel"
	logs        "github.com/mt1976/ebEstimates/logs"
		core        "github.com/mt1976/ebEstimates/core"

	cron        "github.com/robfig/cron/v3"
)

// CredentialsPasswordJob defines the job properties, name, period etc..
func CredentialsPassword_Job() dm.JobDefinition {
	// ----------------------------------------------------------------
	// Job Definition
	// Create a new job definition in the CredentialsPassword_Impl.go file called CredentialsPassword_JobImpl()
	// ----------------------------------------------------------------
	var j dm.JobDefinition
	j.ID = "CredentialsPassword"
	j.Name = "CredentialsPassword"
	j.Period = "10 1 * * *"
	j.Description = "CredentialsPassword processing"
	j.Type = core.General
	j = CredentialsPassword_Job_impl(j)
	return j
}

func CredentialsPassword_Register(c *cron.Cron) {
	application.Schedule_Register(CredentialsPassword_Job())
	c.AddFunc(CredentialsPassword_Job().Period, func() { CredentialsPassword_Run() })
}

// CredentialsPassword_Run initiates and runs the job as per the period.
func CredentialsPassword_Run() {
	logs.StartJob(CredentialsPassword_Job().Name)
	message := ""
	/// CONTENT STARTS
	// Create a func in the CredentialsPassword_Impl.go file called CredentialsPassword_Run_impl() that returns string,error	
	message,err := CredentialsPassword_Run_impl()
	if err != nil {
		logs.Warning(err.Error())
	}
	/// CONTENT ENDS
	application.Schedule_Update(CredentialsPassword_Job(), message)
	logs.EndJob(CredentialsPassword_Job().Name)
}