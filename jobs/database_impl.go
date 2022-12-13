package jobs

// ----------------------------------------------------------------
// Automatically generated  "/jobs/origin.go"
// ----------------------------------------------------------------
// Package              : jobs
// Object 			    : Origin (origin)
// Endpoint 	        : Origin (OriginID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 11/12/2022 at 19:53:20
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	application "github.com/mt1976/ebEstimates/application"
	core "github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"

	cron "github.com/robfig/cron/v3"
)

// OriginJob defines the job properties, name, period etc..
func Database_Job() dm.JobDefinition {
	// ----------------------------------------------------------------
	// Job Definition
	// Create a new job definition in the Database_Impl.go file called Database_JobImpl()
	// ----------------------------------------------------------------
	var j dm.JobDefinition
	j.ID = "Database"
	j.Name = "Database"
	j.Period = "0 1/ * * *"
	j.Description = "Database Connection Check"
	j.Type = core.General
	j = Database_Job_impl(j)
	return j
}

func Database_Register(c *cron.Cron) {
	application.Schedule_Register(Database_Job())
	c.AddFunc(Database_Job().Period, func() { Database_Run() })
}

// Database_Run initiates and runs the job as per the period.
func Database_Run() {
	logs.StartJob(Database_Job().Name)
	message := ""
	/// CONTENT STARTS
	// Create a func in the Database_Impl.go file called Database_Run_impl() that returns string,error
	message, err := Database_Run_impl()
	if err != nil {
		logs.Warning(err.Error())
	}
	/// CONTENT ENDS
	application.Schedule_Update(Database_Job(), message)
	logs.EndJob(Database_Job().Name)
}

func Database_Job_impl(j dm.JobDefinition) dm.JobDefinition {
	j.Name = "Database"
	j.Period = "0 */1 * * *"
	j.Description = "Database Connection Check"
	j.Type = core.HouseKeeping
	j.ID = "Database"
	return j
}

func Database_Run_impl() (string, error) {

	message := ""
	/// CONTENT STARTS
	// Refreshes and reconnects to the database
	core.ApplicationDB = core.Database_Poke(core.ApplicationDB, core.ApplicationPropertiesDB)

	/// CONTENT ENDS

	return message, nil

}
