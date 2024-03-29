package dao

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/projectaction.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : ProjectAction (projectaction)
// Endpoint 	        : ProjectAction (ProjectID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/02/2023 at 10:44:46
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------
// The following functions should be created in projectaction_impl.go
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// ProjectAction_ObjectValidation_impl provides Record/Object level validation for ProjectAction
func ProjectAction_ObjectValidation_impl(iAction string, iId string, iRec dm.ProjectAction) (dm.ProjectAction, string, error) {
	logs.Callout("ProjectAction", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("ProjectAction" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}

//
// ----------------------------------------------------------------
// These are the the default implementations, which do nothing

// ----------------------------------------------------------------
// BEGIN ProjectAction_NoEstimationSessions
// BEGIN ProjectAction_NoEstimationSessions
// BEGIN ProjectAction_NoEstimationSessions
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// ProjectAction_NoEstimationSessions_validate_impl provides validation/actions for NoEstimationSessions
func ProjectAction_NoEstimationSessions_validate_impl(iAction string, iId string, iValue string, iRec dm.ProjectAction, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("ProjectAction", dm.ProjectAction_NoEstimationSessions_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   ProjectAction_NoEstimationSessions
// END   ProjectAction_NoEstimationSessions
// END   ProjectAction_NoEstimationSessions
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// ProjectAction_OriginName_validate_impl provides validation/actions for OriginName
func ProjectAction_OriginName_validate_impl(iAction string, iId string, iValue string, iRec dm.ProjectAction, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("ProjectAction", dm.ProjectAction_OriginName_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// ProjectAction_OriginKey_validate_impl provides validation/actions for OriginKey
func ProjectAction_OriginKey_validate_impl(iAction string, iId string, iValue string, iRec dm.ProjectAction, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("ProjectAction", dm.ProjectAction_OriginKey_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ProjectAction_ProjectStateID_validate_impl provides validation/actions for ProjectStateID
func ProjectAction_ProjectStateID_validate_impl(iAction string, iId string, iValue string, iRec dm.ProjectAction, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("ProjectAction", dm.ProjectAction_ProjectStateID_scrn, VAL+"-"+iAction, iId)
	if iValue == "" {
		fP = SetFieldError(fP, "status is required")
	}
	return iValue, fP
}

// ----------------------------------------------------------------
// ProjectAction_ProfileID_validate_impl provides validation/actions for ProfileID
func ProjectAction_ProfileID_validate_impl(iAction string, iId string, iValue string, iRec dm.ProjectAction, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("ProjectAction", dm.ProjectAction_ProfileID_scrn, VAL+"-"+iAction, iId)
	if iValue == "" {
		fP = SetFieldError(fP, "activity profile is required")
	}
	return iValue, fP
}
