package dao

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/feature.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : Feature (feature)
// Endpoint 	        : Feature (FeatureID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 03/01/2023 at 19:18:09
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in feature_impl.go

// If there are fields below, create the methods in adaptor\feature_impl.go
// START - GET API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - GET API/Callout
//
// START - PUT API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - PUT API/Callout

// START - Validation API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout

func Feature_ObjectValidation_impl(iAction string, iId string, iRec dm.Feature) (dm.Feature, string, error) {
	logs.Callout("Feature", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("Feature" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}

// ----------------------------------------------------------------
// BEGIN Feature_FeatureID
// BEGIN Feature_FeatureID
// BEGIN Feature_FeatureID
// ----------------------------------------------------------------
// Feature_FeatureID_OnStore_impl provides the implementation for the callout
func Feature_FeatureID_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_FeatureID_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_FeatureID_OnFetch_impl provides the implementation for the callout
func Feature_FeatureID_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_FeatureID_scrn, GET, rec.FeatureID)
	return rec.FeatureID
}

// ----------------------------------------------------------------
// Feature_FeatureID_impl provides validation/actions for FeatureID
func Feature_FeatureID_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_FeatureID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   Feature_FeatureID
// END   Feature_FeatureID
// END   Feature_FeatureID
// ----------------------------------------------------------------
