package dao

import (
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
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/02/2023 at 10:44:45
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------
// The following functions should be created in origin_impl.go
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Origin_ObjectValidation_impl provides Record/Object level validation for Origin
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

//
// ----------------------------------------------------------------
// These are the the default implementations, which do nothing

// ----------------------------------------------------------------
// BEGIN Origin_StateID
// BEGIN Origin_StateID
// BEGIN Origin_StateID
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Origin_StateID_validate_impl provides validation/actions for StateID
func Origin_StateID_validate_impl(iAction string, iId string, iValue string, iRec dm.Origin, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Origin", dm.Origin_StateID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   Origin_StateID
// END   Origin_StateID
// END   Origin_StateID
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Origin_Code
// BEGIN Origin_Code
// BEGIN Origin_Code
// ----------------------------------------------------------------
// Origin_Code_OnStore_impl provides the implementation for the callout

// ----------------------------------------------------------------
// Origin_Code_validate_impl provides validation/actions for Code
func Origin_Code_validate_impl(iAction string, iId string, iValue string, iRec dm.Origin, fP dm.FieldProperties) (string, dm.FieldProperties) {
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

// ----------------------------------------------------------------
// Origin_FullName_validate_impl provides validation/actions for FullName
func Origin_FullName_validate_impl(iAction string, iId string, iValue string, iRec dm.Origin, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Origin", dm.Origin_FullName_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   Origin_FullName
// END   Origin_FullName
// END   Origin_FullName
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Origin_Rate
// BEGIN Origin_Rate
// BEGIN Origin_Rate
// ----------------------------------------------------------------
// Origin_Rate_OnStore_impl provides the implementation for the callout

// ----------------------------------------------------------------
// Origin_Rate_validate_impl provides validation/actions for Rate
func Origin_Rate_validate_impl(iAction string, iId string, iValue string, iRec dm.Origin, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Origin", dm.Origin_Rate_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   Origin_Rate
// END   Origin_Rate
// END   Origin_Rate
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Origin_NoActiveProjects
// BEGIN Origin_NoActiveProjects
// BEGIN Origin_NoActiveProjects
// ----------------------------------------------------------------
// Origin_NoActiveProjects_OnStore_impl provides the implementation for the callout

// ----------------------------------------------------------------
// Origin_NoActiveProjects_OnFetch_impl provides the implementation for the callout

// ----------------------------------------------------------------
// Origin_NoActiveProjects_validate_impl provides validation/actions for NoActiveProjects
func Origin_NoActiveProjects_validate_impl(iAction string, iId string, iValue string, iRec dm.Origin, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Origin", dm.Origin_NoActiveProjects_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   Origin_NoActiveProjects
// END   Origin_NoActiveProjects
// END   Origin_NoActiveProjects
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------
