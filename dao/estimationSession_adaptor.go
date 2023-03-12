package dao

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/estimationsession.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : EstimationSession (estimationsession)
// Endpoint 	        : EstimationSession (EstimationSessionID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 03/01/2023 at 19:57:49
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in estimationsession_impl.go

// EstimationSession_OriginKey_OnStore_impl provides the implementation for the callout
func EstimationSession_OriginKey_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginKey_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// EstimationSession_OriginKey_OnFetch_impl provides the implementation for the callout
func EstimationSession_OriginKey_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginKey_scrn, GET, rec.EstimationSessionID)
	return rec.OriginKey
}
