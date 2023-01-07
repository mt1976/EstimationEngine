package jobs
// ----------------------------------------------------------------
// Automatically generated  "/jobs/estimationsession.go"
// ----------------------------------------------------------------
// Package              : jobs
// Object 			    : EstimationSession (estimationsession)
// Endpoint 	        : EstimationSession (EstimationSessionID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 07/01/2023 at 23:01:29
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	application "github.com/mt1976/ebEstimates/application"
	dm          "github.com/mt1976/ebEstimates/datamodel"
	logs        "github.com/mt1976/ebEstimates/logs"
		core        "github.com/mt1976/ebEstimates/core"

	cron        "github.com/robfig/cron/v3"
)

// EstimationSessionJob defines the job properties, name, period etc..
func EstimationSession_Job() dm.JobDefinition {
	// ----------------------------------------------------------------
	// Job Definition
	// Create a new job definition in the EstimationSession_Impl.go file called EstimationSession_JobImpl()
	// ----------------------------------------------------------------
	var j dm.JobDefinition
	j.ID = "EstimationSession"
	j.Name = "EstimationSession"
	j.Period = "10 1 * * *"
	j.Description = "EstimationSession processing"
	j.Type = core.General
	j = EstimationSession_Job_impl(j)
	return j
}

func EstimationSession_Register(c *cron.Cron) {
	application.Schedule_Register(EstimationSession_Job())
	c.AddFunc(EstimationSession_Job().Period, func() { EstimationSession_Run() })
}

// EstimationSession_Run initiates and runs the job as per the period.
func EstimationSession_Run() {
	logs.StartJob(EstimationSession_Job().Name)
	message := ""
	/// CONTENT STARTS
	// Create a func in the EstimationSession_Impl.go file called EstimationSession_Run_impl() that returns string,error	
	message,err := EstimationSession_Run_impl()
	if err != nil {
		logs.Warning(err.Error())
	}
	/// CONTENT ENDS
	application.Schedule_Update(EstimationSession_Job(), message)
	logs.EndJob(EstimationSession_Job().Name)
}