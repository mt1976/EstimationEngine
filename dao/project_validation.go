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
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/02/2023 at 10:44:46
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------
// The following functions should be created in project_impl.go
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Project_ObjectValidation_impl provides Record/Object level validation for Project
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

//
// ----------------------------------------------------------------
// These are the the default implementations, which do nothing

// ----------------------------------------------------------------
// BEGIN Project_ProjectID
// BEGIN Project_ProjectID
// BEGIN Project_ProjectID
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Project_ProjectID_validate_impl provides validation/actions for ProjectID
func Project_ProjectID_validate_impl(iAction string, iId string, iValue string, iRec dm.Project, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Project", dm.Project_ProjectID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   Project_ProjectID
// END   Project_ProjectID
// END   Project_ProjectID
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Project_Description
// BEGIN Project_Description
// BEGIN Project_Description
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Project_Description_validate_impl provides validation/actions for Description
func Project_Description_validate_impl(iAction string, iId string, iValue string, iRec dm.Project, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Project", dm.Project_Description_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   Project_Description
// END   Project_Description
// END   Project_Description
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Project_NoEstimationSessions
// BEGIN Project_NoEstimationSessions
// BEGIN Project_NoEstimationSessions
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Project_NoEstimationSessions_validate_impl provides validation/actions for NoEstimationSessions
func Project_NoEstimationSessions_validate_impl(iAction string, iId string, iValue string, iRec dm.Project, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Project", dm.Project_NoEstimationSessions_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   Project_NoEstimationSessions
// END   Project_NoEstimationSessions
// END   Project_NoEstimationSessions
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------
