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

// ----------------------------------------------------------------
// BEGIN Feature_TrackerID
// BEGIN Feature_TrackerID
// BEGIN Feature_TrackerID
// ----------------------------------------------------------------
// Feature_TrackerID_OnStore_impl provides the implementation for the callout
func Feature_TrackerID_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_TrackerID_scrn, PUT, rec.FeatureID)
	Indexer_Put("Feature", dm.Feature_TrackerID_scrn, rec.FeatureID, fieldval)

	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_TrackerID_OnFetch_impl provides the implementation for the callout
func Feature_TrackerID_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_TrackerID_scrn, GET, rec.FeatureID)
	return rec.TrackerID
}

// ----------------------------------------------------------------
// Feature_TrackerID_impl provides validation/actions for TrackerID
func Feature_TrackerID_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_TrackerID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   Feature_TrackerID
// END   Feature_TrackerID
// END   Feature_TrackerID
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_AdoID
// BEGIN Feature_AdoID
// BEGIN Feature_AdoID
// ----------------------------------------------------------------
// Feature_AdoID_OnStore_impl provides the implementation for the callout
func Feature_AdoID_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	Indexer_Put("Feature", dm.Feature_AdoID_scrn, rec.FeatureID, fieldval)

	logs.Callout("Feature", dm.Feature_AdoID_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_AdoID_OnFetch_impl provides the implementation for the callout
func Feature_AdoID_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_AdoID_scrn, GET, rec.FeatureID)
	return rec.AdoID
}

// ----------------------------------------------------------------
// Feature_AdoID_impl provides validation/actions for AdoID
func Feature_AdoID_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_AdoID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   Feature_AdoID
// END   Feature_AdoID
// END   Feature_AdoID
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_FreshdeskID
// BEGIN Feature_FreshdeskID
// BEGIN Feature_FreshdeskID
// ----------------------------------------------------------------
// Feature_FreshdeskID_OnStore_impl provides the implementation for the callout
func Feature_FreshdeskID_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_FreshdeskID_scrn, PUT, rec.FeatureID)
	Indexer_Put("Feature", dm.Feature_FreshdeskID_scrn, rec.FeatureID, fieldval)

	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_FreshdeskID_OnFetch_impl provides the implementation for the callout
func Feature_FreshdeskID_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_FreshdeskID_scrn, GET, rec.FeatureID)
	return rec.FreshdeskID
}

// ----------------------------------------------------------------
// Feature_FreshdeskID_impl provides validation/actions for FreshdeskID
func Feature_FreshdeskID_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_FreshdeskID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   Feature_FreshdeskID
// END   Feature_FreshdeskID
// END   Feature_FreshdeskID
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_ExtRef
// BEGIN Feature_ExtRef
// BEGIN Feature_ExtRef
// ----------------------------------------------------------------
// Feature_ExtRef_OnStore_impl provides the implementation for the callout
func Feature_ExtRef_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_ExtRef_scrn, PUT, rec.FeatureID)
	Indexer_Put("Feature", dm.Feature_ExtRef_scrn, rec.FeatureID, fieldval)

	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_ExtRef_OnFetch_impl provides the implementation for the callout
func Feature_ExtRef_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_ExtRef_scrn, GET, rec.FeatureID)
	return rec.ExtRef
}

// ----------------------------------------------------------------
// Feature_ExtRef_impl provides validation/actions for ExtRef
func Feature_ExtRef_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_ExtRef_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   Feature_ExtRef
// END   Feature_ExtRef
// END   Feature_ExtRef
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_ExtRef2
// BEGIN Feature_ExtRef2
// BEGIN Feature_ExtRef2
// ----------------------------------------------------------------
// Feature_ExtRef2_OnStore_impl provides the implementation for the callout
func Feature_ExtRef2_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_ExtRef2_scrn, PUT, rec.FeatureID)
	Indexer_Put("Feature", dm.Feature_ExtRef2_scrn, rec.FeatureID, fieldval)

	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_ExtRef2_OnFetch_impl provides the implementation for the callout
func Feature_ExtRef2_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_ExtRef2_scrn, GET, rec.FeatureID)
	return rec.ExtRef2
}

// ----------------------------------------------------------------
// Feature_ExtRef2_impl provides validation/actions for ExtRef2
func Feature_ExtRef2_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_ExtRef2_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   Feature_ExtRef2
// END   Feature_ExtRef2
// END   Feature_ExtRef2
// ----------------------------------------------------------------
