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
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 03/01/2023 at 19:18:10
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in projectaction_impl.go

// If there are fields below, create the methods in adaptor\projectaction_impl.go
// START - GET API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
func ProjectAction_NoEstimationSessions_OnStore_impl(fieldval string, rec dm.ProjectAction, usr string) (string, error) {
	logs.Callout("ProjectAction", dm.ProjectAction_NoEstimationSessions_scrn, PUT, rec.ProjectID)
	return fieldval, nil
}

// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - GET API/Callout
//
// START - PUT API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
func ProjectAction_NoEstimationSessions_OnFetch_impl(rec dm.ProjectAction) string {
	logs.Callout("ProjectAction", dm.ProjectAction_NoEstimationSessions_scrn, GET, rec.ProjectID)
	return rec.NoEstimationSessions
}

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - PUT API/Callout

// START - Validation API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
//
// ProjectAction_NoEstimationSessions_impl provides validation/actions for NoEstimationSessions
func ProjectAction_NoEstimationSessions_impl(iAction string, iId string, iValue string, iRec dm.ProjectAction, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("ProjectAction", dm.ProjectAction_NoEstimationSessions_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout

// ----------------------------------------------------------------
// BEGIN ProjectAction_OriginName
// BEGIN ProjectAction_OriginName
// BEGIN ProjectAction_OriginName
// ----------------------------------------------------------------
// ProjectAction_OriginName_OnStore_impl provides the implementation for the callout
func ProjectAction_OriginName_OnStore_impl(fieldval string, rec dm.ProjectAction, usr string) (string, error) {
	logs.Callout("ProjectAction", dm.ProjectAction_OriginName_scrn, PUT, rec.ProjectID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// ProjectAction_OriginName_OnFetch_impl provides the implementation for the callout
func ProjectAction_OriginName_OnFetch_impl(rec dm.ProjectAction) string {
	logs.Callout("ProjectAction", dm.ProjectAction_OriginName_scrn, GET, rec.ProjectID)
	//fmt.Printf("rec: %v\n", rec)
	orginID := rec.OriginID
	if orginID == "" {
		return ""
	}
	origin := CacheRead(dm.Origin_Name, orginID).(dm.Origin)

	return origin.FullName
}

// ----------------------------------------------------------------
// END   ProjectAction_OriginName
// END   ProjectAction_OriginName
// END   ProjectAction_OriginName
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// BEGIN ProjectAction_OriginKey
// BEGIN ProjectAction_OriginKey
// BEGIN ProjectAction_OriginKey
// ----------------------------------------------------------------
// ProjectAction_OriginKey_OnStore_impl provides the implementation for the callout
func ProjectAction_OriginKey_OnStore_impl(fieldval string, rec dm.ProjectAction, usr string) (string, error) {
	logs.Callout("ProjectAction", dm.ProjectAction_OriginKey_scrn, PUT, rec.ProjectID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// ProjectAction_OriginKey_OnFetch_impl provides the implementation for the callout
func ProjectAction_OriginKey_OnFetch_impl(rec dm.ProjectAction) string {
	logs.Callout("ProjectAction", dm.ProjectAction_OriginKey_scrn, GET, rec.ProjectID)
	originID := rec.OriginID
	origin := CacheRead(dm.Origin_Name, originID).(dm.Origin)

	return origin.OriginID
}

// ----------------------------------------------------------------
// END   ProjectAction_OriginKey
// END   ProjectAction_OriginKey
// END   ProjectAction_OriginKey
// ----------------------------------------------------------------
