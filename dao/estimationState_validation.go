package dao

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/estimationstate.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : EstimationState (estimationstate)
// Endpoint 	        : EstimationState (EstimationStateID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/02/2023 at 10:44:43
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------
// The following functions should be created in estimationstate_impl.go
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// EstimationState_ObjectValidation_impl provides Record/Object level validation for EstimationState
func EstimationState_ObjectValidation_impl(iAction string, iId string, iRec dm.EstimationState) (dm.EstimationState, string, error) {
	logs.Callout("EstimationState", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("EstimationState" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}

//
// ----------------------------------------------------------------
// These are the the default implementations, which do nothing

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------
