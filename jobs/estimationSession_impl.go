package jobs

// ----------------------------------------------------------------
// Automatically generated  "/jobs/EstimationSession.go"
// ----------------------------------------------------------------
// Package              : jobs
// Object 			    : EstimationSession (EstimationSession)
// Endpoint 	        : EstimationSession (MessageID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/12/2021 at 09:31:50
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	core "github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
)

// EstimationSessionJob defines the job properties, name, period etc..
func EstimationSession_Job_impl(j dm.JobDefinition) dm.JobDefinition {
	// ----------------------------------------------------------------
	// Job Definition
	// ----------------------------------------------------------------
	j.ID = "EstimationSession"
	j.Name = "EstimationSession"
	j.Period = "10 00 * * *"
	j.Description = "EstimationSession processing"
	j.Type = core.General
	return j
}

// EstimationSession_Run initiates and runs the job as per the period.
func EstimationSession_Run_impl() (string, error) {

	//GEt a list of unprocessed messages from EstimationSession

	return "", nil
}
