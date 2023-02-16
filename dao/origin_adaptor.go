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

// ----------------------------------------------------------------
// BEGIN Origin_Code
// BEGIN Origin_Code
// BEGIN Origin_Code
// ----------------------------------------------------------------
// Origin_Code_OnStore_impl provides the implementation for the callout
func Origin_Code_OnStore_impl(fieldval string, rec dm.Origin, usr string) (string, error) {
	logs.Callout("Origin", dm.Origin_Code_scrn, PUT, rec.OriginID)
	Indexer_Put("Origin", dm.Origin_Code_scrn, rec.OriginID, fieldval)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Origin_Code_OnFetch_impl provides the implementation for the callout
func Origin_Code_OnFetch_impl(rec dm.Origin) string {
	logs.Callout("Origin", dm.Origin_Code_scrn, GET, rec.OriginID)

	return rec.Code
}

// ----------------------------------------------------------------
// Origin_Code_impl provides validation/actions for Code
func Origin_Code_impl(iAction string, iId string, iValue string, iRec dm.Origin, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Origin", dm.Origin_Code_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   Origin_Code
// END   Origin_Code
// END   Origin_Code
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Origin_FullName
// BEGIN Origin_FullName
// BEGIN Origin_FullName
// ----------------------------------------------------------------
// Origin_FullName_OnStore_impl provides the implementation for the callout
func Origin_FullName_OnStore_impl(fieldval string, rec dm.Origin, usr string) (string, error) {
	logs.Callout("Origin", dm.Origin_FullName_scrn, PUT, rec.OriginID)
	Indexer_Put("Origin", dm.Origin_FullName_scrn, rec.OriginID, fieldval)

	return fieldval, nil
}

// ----------------------------------------------------------------
// Origin_FullName_OnFetch_impl provides the implementation for the callout
func Origin_FullName_OnFetch_impl(rec dm.Origin) string {
	logs.Callout("Origin", dm.Origin_FullName_scrn, GET, rec.OriginID)
	return rec.FullName
}

// ----------------------------------------------------------------
// Origin_FullName_impl provides validation/actions for FullName
func Origin_FullName_impl(iAction string, iId string, iValue string, iRec dm.Origin, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Origin", dm.Origin_FullName_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   Origin_FullName
// END   Origin_FullName
// END   Origin_FullName
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Origin_StartDate
// BEGIN Origin_StartDate
// BEGIN Origin_StartDate
// ----------------------------------------------------------------
// Origin_StartDate_OnStore_impl provides the implementation for the callout
func Origin_StartDate_OnStore_impl(fieldval string, rec dm.Origin, usr string) (string, error) {
	logs.Callout("Origin", dm.Origin_StartDate_scrn, PUT, rec.OriginID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Origin_StartDate_OnFetch_impl provides the implementation for the callout
func Origin_StartDate_OnFetch_impl(rec dm.Origin) string {
	logs.Callout("Origin", dm.Origin_StartDate_scrn, GET, rec.OriginID)
	return rec.StartDate
}

// ----------------------------------------------------------------
// END   Origin_StartDate
// END   Origin_StartDate
// END   Origin_StartDate
// ----------------------------------------------------------------
