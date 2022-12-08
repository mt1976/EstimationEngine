package adaptor

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/feature.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Feature (feature)
// Endpoint 	        : Feature (FeatureID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 08/12/2022 at 13:18:34
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in feature_impl.go

// If there are fields below, create the methods in adaptor\feature_impl.go
// START - GET API/Callout
// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local
// END - GET API/Callout
//
// START - PUT API/Callout
// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local
// END - PUT API/Callout

// START - Validation API/Callout
// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout

func Feature_ObjectValidation_impl(iAction string, iId string, iRec dm.Feature) (dm.Feature, string, error) {
	logs.Callout("Feature", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("Tmpl_TDate_impl" + " - Invalid Action")
	}
	return iRec, "", nil
}
