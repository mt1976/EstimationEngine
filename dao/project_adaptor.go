package dao

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/project.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : Project (project)
// Endpoint 	        : Project (ProjectID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 03/01/2023 at 19:18:10
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in project_impl.go

// If there are fields below, create the methods in adaptor\project_impl.go
// START - GET API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
func Project_NoEstimationSessions_OnStore_impl(fieldval string, rec dm.Project, usr string) (string, error) {
	logs.Callout("Project", dm.Project_NoEstimationSessions_scrn, PUT, rec.ProjectID)
	return fieldval, nil
}

// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - GET API/Callout
//
// START - PUT API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
func Project_NoEstimationSessions_OnFetch_impl(rec dm.Project) string {
	logs.Callout("Project", dm.Project_NoEstimationSessions_scrn, GET, rec.ProjectID)
	return rec.NoEstimationSessions
}

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - PUT API/Callout

// START - Validation API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
//
// Project_NoEstimationSessions_impl provides validation/actions for NoEstimationSessions
func Project_NoEstimationSessions_impl(iAction string, iId string, iValue string, iRec dm.Project, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Project", dm.Project_NoEstimationSessions_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout

func Project_ObjectValidation_impl(iAction string, iId string, iRec dm.Project) (dm.Project, string, error) {
	logs.Callout("Project", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("Project" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}