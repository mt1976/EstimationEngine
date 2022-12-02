package adaptor

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/featurenew.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : FeatureNew (featurenew)
// Endpoint 	        : FeatureNew (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 29/11/2022 at 19:40:01
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in featurenew_impl.go

func FeatureNew_GetList_impl() (int, []dm.FeatureNew, error)        { return 0, nil, nil }
func FeatureNew_GetByID_impl(id string) (int, dm.FeatureNew, error) { return 0, dm.FeatureNew{}, nil }

func FeatureNew_NewID_impl(rec dm.FeatureNew) string { return rec.ID }
func FeatureNew_Delete_impl(id string) error         { return nil }

func FeatureNew_Update_impl(id string, rec dm.FeatureNew, usr string) error { return nil }

// If there are fields below, create the methods in adaptor\featurenew_impl.go
// START - GET API/Callout
// Dynamically generated 29/11/2022 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 29/11/2022 by matttownsend (Matt Townsend) on silicon.local
// END - GET API/Callout
//
// START - PUT API/Callout
// Dynamically generated 29/11/2022 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 29/11/2022 by matttownsend (Matt Townsend) on silicon.local
// END - PUT API/Callout

// START - Validation API/Callout
// Dynamically generated 29/11/2022 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 29/11/2022 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout

func FeatureNew_ObjectValidation_impl(iAction string, iId string, iRec dm.FeatureNew) (dm.FeatureNew, error, string) {
	logs.Callout("FeatureNew", "ObjectValidation", VAL+"-"+iAction, iId)
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
