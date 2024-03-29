package dao

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/externalmessage.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : ExternalMessage (externalmessage)
// Endpoint 	        : ExternalMessage (Message)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/02/2023 at 10:44:43
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------
// The following functions should be created in externalmessage_impl.go
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// ExternalMessage_ObjectValidation_impl provides Record/Object level validation for ExternalMessage
func ExternalMessage_ObjectValidation_impl(iAction string, iId string, iRec dm.ExternalMessage) (dm.ExternalMessage, string, error) {
	logs.Callout("ExternalMessage", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("ExternalMessage" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}

//
// ----------------------------------------------------------------
// These are the the default implementations, which do nothing

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------
