package dao

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/confidence.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : Confidence (confidence)
// Endpoint 	        : Confidence (ConfidenceID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/02/2023 at 10:44:40
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------
// The following functions should be created in confidence_impl.go
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Confidence_ObjectValidation_impl provides Record/Object level validation for Confidence
func Confidence_ObjectValidation_impl(iAction string, iId string, iRec dm.Confidence) (dm.Confidence, string, error) {
	logs.Callout("Confidence", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("Confidence" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}

//
// ----------------------------------------------------------------
// These are the the default implementations, which do nothing

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------
