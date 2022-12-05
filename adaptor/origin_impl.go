package adaptor

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/origin.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Origin (origin)
// Endpoint 	        : Origin (OriginID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 05/12/2022 at 21:52:44
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in origin_impl.go

// If there are fields below, create the methods in adaptor\origin_impl.go
// START - GET API/Callout
// Dynamically generated 05/12/2022 by matttownsend (Matt Townsend) on silicon.local
func Origin_NoActiveProjects_OnStore_impl(fieldval string, rec dm.Origin, usr string) (string, error) {
	logs.Callout("Origin", "NoActiveProjects", PUT, rec.OriginID)
	return fieldval, nil
}

// Dynamically generated 05/12/2022 by matttownsend (Matt Townsend) on silicon.local
// END - GET API/Callout
//
// START - PUT API/Callout
// Dynamically generated 05/12/2022 by matttownsend (Matt Townsend) on silicon.local
func Origin_NoActiveProjects_OnFetch_impl(rec dm.Origin) string {
	logs.Callout("Origin", "NoActiveProjects", GET, rec.OriginID)
	return rec.NoActiveProjects
}

//
// Dynamically generated 05/12/2022 by matttownsend (Matt Townsend) on silicon.local
// END - PUT API/Callout

// START - Validation API/Callout
// Dynamically generated 05/12/2022 by matttownsend (Matt Townsend) on silicon.local
//
// Origin_NoActiveProjects_impl provides validation/actions for NoActiveProjects
func Origin_NoActiveProjects_impl(iAction string, iId string, iValue string, iRec dm.Origin, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Origin", "NoActiveProjects", VAL+"-"+iAction, iId)
	return iValue, fP
}

//
// Dynamically generated 05/12/2022 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout

func Origin_ObjectValidation_impl(iAction string, iId string, iRec dm.Origin) (dm.Origin, error, string) {
	logs.Callout("Origin", "ObjectValidation", VAL+"-"+iAction, iId)
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
