package jobs
// ----------------------------------------------------------------
// Automatically generated  "/jobs/project.go"
// ----------------------------------------------------------------
// Package              : jobs
// Object 			    : Project (project)
// Endpoint 	        : Project (ProjectID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 07/01/2023 at 23:01:30
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	application "github.com/mt1976/ebEstimates/application"
	dm          "github.com/mt1976/ebEstimates/datamodel"
	logs        "github.com/mt1976/ebEstimates/logs"
		core        "github.com/mt1976/ebEstimates/core"

	cron        "github.com/robfig/cron/v3"
)

// ProjectJob defines the job properties, name, period etc..
func Project_Job() dm.JobDefinition {
	// ----------------------------------------------------------------
	// Job Definition
	// Create a new job definition in the Project_Impl.go file called Project_JobImpl()
	// ----------------------------------------------------------------
	var j dm.JobDefinition
	j.ID = "Project"
	j.Name = "Project"
	j.Period = "10 1 * * *"
	j.Description = "Project processing"
	j.Type = core.General
	j = Project_Job_impl(j)
	return j
}

func Project_Register(c *cron.Cron) {
	application.Schedule_Register(Project_Job())
	c.AddFunc(Project_Job().Period, func() { Project_Run() })
}

// Project_Run initiates and runs the job as per the period.
func Project_Run() {
	logs.StartJob(Project_Job().Name)
	message := ""
	/// CONTENT STARTS
	// Create a func in the Project_Impl.go file called Project_Run_impl() that returns string,error	
	message,err := Project_Run_impl()
	if err != nil {
		logs.Warning(err.Error())
	}
	/// CONTENT ENDS
	application.Schedule_Update(Project_Job(), message)
	logs.EndJob(Project_Job().Name)
}