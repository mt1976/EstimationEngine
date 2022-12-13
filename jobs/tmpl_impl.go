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
	dm "github.com/mt1976/ebEstimates/datamodel"
)

func Tmpl_Job_impl(j dm.JobDefinition) dm.JobDefinition {
	return j
}

// RunJobRollover is a Rollover function
func Tmpl_Run_impl() (string, error) {
	message := ""
	/// CONTENT STARTS

	/// CONTENT ENDS
	return message, nil
}
