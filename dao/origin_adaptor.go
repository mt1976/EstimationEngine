package dao

import (
	"strconv"

	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/origin.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : Origin (origin)
// Endpoint 	        : Origin (OriginID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 03/01/2023 at 19:16:35
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in origin_impl.go

// If there are fields below, create the methods in adaptor\origin_impl.go
// START - GET API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
func Origin_StateID_OnStore_impl(fieldval string, rec dm.Origin, usr string) (string, error) {
	logs.Callout("Origin", dm.Origin_StateID_scrn, PUT, rec.OriginID)
	return fieldval, nil
}
func Origin_Rate_OnStore_impl(fieldval string, rec dm.Origin, usr string) (string, error) {
	logs.Callout("Origin", dm.Origin_Rate_scrn, PUT, rec.OriginID)
	return fieldval, nil
}
func Origin_NoActiveProjects_OnStore_impl(fieldval string, rec dm.Origin, usr string) (string, error) {
	logs.Callout("Origin", dm.Origin_NoActiveProjects_scrn, PUT, rec.OriginID)
	return fieldval, nil
}

// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - GET API/Callout
//
// START - PUT API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
func Origin_StateID_OnFetch_impl(rec dm.Origin) string {
	logs.Callout("Origin", dm.Origin_StateID_scrn, GET, rec.OriginID)
	return rec.StateID
}
func Origin_Rate_OnFetch_impl(rec dm.Origin) string {
	logs.Callout("Origin", dm.Origin_Rate_scrn, GET, rec.OriginID)
	return rec.Rate
}
func Origin_NoActiveProjects_OnFetch_impl(rec dm.Origin) string {
	logs.Callout("Origin", dm.Origin_NoActiveProjects_scrn, GET, rec.OriginID)
	noProj, _, _ := Project_ByOrigin_Active_GetList(rec.Code)
	strProj := strconv.Itoa(noProj)
	return strProj
}

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - PUT API/Callout

// START - Validation API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
//
// Origin_StateID_impl provides validation/actions for StateID
func Origin_StateID_impl(iAction string, iId string, iValue string, iRec dm.Origin, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Origin", dm.Origin_StateID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// Origin_Rate_impl provides validation/actions for Rate
func Origin_Rate_impl(iAction string, iId string, iValue string, iRec dm.Origin, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Origin", dm.Origin_Rate_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// Origin_NoActiveProjects_impl provides validation/actions for NoActiveProjects
func Origin_NoActiveProjects_impl(iAction string, iId string, iValue string, iRec dm.Origin, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Origin", dm.Origin_NoActiveProjects_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout

func Origin_ObjectValidation_impl(iAction string, iId string, iRec dm.Origin) (dm.Origin, string, error) {
	logs.Callout("Origin", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("Origin" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}
