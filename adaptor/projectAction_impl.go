package adaptor

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/projectaction.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : ProjectAction (projectaction)
// Endpoint 	        : ProjectAction (ProjectID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 02/01/2023 at 12:08:02
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in projectaction_impl.go

// If there are fields below, create the methods in adaptor\projectaction_impl.go
// START - GET API/Callout
// Dynamically generated 02/01/2023 by matttownsend (Matt Townsend) on silicon.local
func ProjectAction_NoEstimationSessions_OnStore_impl(fieldval string, rec dm.ProjectAction, usr string) (string, error) {
	logs.Callout("ProjectAction", "NoEstimationSessions", PUT, rec.ProjectID)
	return fieldval, nil
}

// Dynamically generated 02/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - GET API/Callout
//
// START - PUT API/Callout
// Dynamically generated 02/01/2023 by matttownsend (Matt Townsend) on silicon.local
func ProjectAction_NoEstimationSessions_OnFetch_impl(rec dm.ProjectAction) string {
	logs.Callout("ProjectAction", "NoEstimationSessions", GET, rec.ProjectID)
	return rec.NoEstimationSessions
}

//
// Dynamically generated 02/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - PUT API/Callout

// START - Validation API/Callout
// Dynamically generated 02/01/2023 by matttownsend (Matt Townsend) on silicon.local
//
// ProjectAction_NoEstimationSessions_impl provides validation/actions for NoEstimationSessions
func ProjectAction_NoEstimationSessions_impl(iAction string, iId string, iValue string, iRec dm.ProjectAction, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("ProjectAction", "NoEstimationSessions", VAL+"-"+iAction, iId)
	return iValue, fP
}

//
// Dynamically generated 02/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout

func ProjectAction_ObjectValidation_impl(iAction string, iId string, iRec dm.ProjectAction) (dm.ProjectAction, string, error) {
	logs.Callout("ProjectAction", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("Tmpl_TDate_impl" + " - Invalid Action")
	}
	return iRec, "", nil
}
