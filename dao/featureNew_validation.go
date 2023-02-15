package dao

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/featurenew.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : FeatureNew (featurenew)
// Endpoint 	        : FeatureNew (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/02/2023 at 10:44:44
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------
// The following functions should be created in featurenew_impl.go
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// FeatureNew_ObjectValidation_impl provides Record/Object level validation for FeatureNew
func FeatureNew_ObjectValidation_impl(iAction string, iId string, iRec dm.FeatureNew) (dm.FeatureNew, string, error) {
	logs.Callout("FeatureNew", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("FeatureNew" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}

//
// ----------------------------------------------------------------
// These are the the default implementations, which do nothing

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------
