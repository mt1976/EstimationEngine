package dao

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/resource.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : Resource (resource)
// Endpoint 	        : Resource (ResourceID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/02/2023 at 10:44:47
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------
// The following functions should be created in resource_impl.go
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Resource_ObjectValidation_impl provides Record/Object level validation for Resource
func Resource_ObjectValidation_impl(iAction string, iId string, iRec dm.Resource) (dm.Resource, string, error) {
	logs.Callout("Resource", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("Resource" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}

//
// ----------------------------------------------------------------
// These are the the default implementations, which do nothing

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------
