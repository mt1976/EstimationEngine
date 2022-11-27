package jobs

// ----------------------------------------------------------------
// Automatically generated  "/application/template.go"
// ----------------------------------------------------------------
// Package            : jobs
// Object 			  : Template
// Endpoint Root 	  : Template
// Search QueryString : TemplateID
// From   			  :
// ----------------------------------------------------------------
// Template Generator : RussetAlbatross [r0-21.11.01]
// ----------------------------------------------------------------
// Date & Time		  : 17/11/2021 at 22:28:20
// Who & Where		  : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	application "github.com/mt1976/ebEstimates/application"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

func Tmpl_Job_impl(j dm.JobDefinition) dm.JobDefinition {
	return j
}

// RunJobRollover is a Rollover function
func Tmpl_Run_impl() (string, error) {
	logs.StartJob(Tmpl_Job().Name)
	message := ""
	/// CONTENT STARTS

	/// CONTENT ENDS
	application.Schedule_Update(Tmpl_Job(), message)
	logs.EndJob(Tmpl_Job().Name)
	return message, nil
}
