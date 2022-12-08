package adaptor

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/project.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Project (project)
// Endpoint 	        : Project (ProjectID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 08/12/2022 at 13:18:35
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in project_impl.go

// If there are fields below, create the methods in adaptor\project_impl.go
// START - GET API/Callout
// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local
func Project_NoEstimationSessions_OnStore_impl(fieldval string, rec dm.Project, usr string) (string, error) {
	logs.Callout("Project", "NoEstimationSessions", PUT, rec.ProjectID)
	return fieldval, nil
}

// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local
// END - GET API/Callout
//
// START - PUT API/Callout
// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local
func Project_NoEstimationSessions_OnFetch_impl(rec dm.Project) string {
	logs.Callout("Project", "NoEstimationSessions", GET, rec.ProjectID)
	return rec.NoEstimationSessions
}

//
// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local
// END - PUT API/Callout

// START - Validation API/Callout
// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local
//
// Project_NoEstimationSessions_impl provides validation/actions for NoEstimationSessions
func Project_NoEstimationSessions_impl(iAction string, iId string, iValue string, iRec dm.Project, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Project", "NoEstimationSessions", VAL+"-"+iAction, iId)
	return iValue, fP
}

//
// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout

func Project_ObjectValidation_impl(iAction string, iId string, iRec dm.Project) (dm.Project, error, string) {
	logs.Callout("Project", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("Tmpl_TDate_impl" + " - Invalid Action")
	}
	return iRec, nil, ""
}
