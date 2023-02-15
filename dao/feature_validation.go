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
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/02/2023 at 10:41:15
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------
// The following functions should be created in feature_impl.go
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Feature_ObjectValidation_impl provides Record/Object level validation for Feature
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

//
// ----------------------------------------------------------------
// These are the the default implementations, which do nothing

// ----------------------------------------------------------------
// BEGIN Feature_FeatureID
// BEGIN Feature_FeatureID
// BEGIN Feature_FeatureID
// ----------------------------------------------------------------
// Feature_FeatureID_OnStore_impl provides the implementation for the callout

// ----------------------------------------------------------------
// Feature_FeatureID_validate_impl provides validation/actions for FeatureID
func Feature_FeatureID_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
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

// ----------------------------------------------------------------
// Feature_TrackerID_validate_impl provides validation/actions for TrackerID
func Feature_TrackerID_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
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

// ----------------------------------------------------------------
// Feature_AdoID_validate_impl provides validation/actions for AdoID
func Feature_AdoID_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
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

// ----------------------------------------------------------------
// Feature_FreshdeskID_validate_impl provides validation/actions for FreshdeskID
func Feature_FreshdeskID_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
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

// ----------------------------------------------------------------
// Feature_ExtRef_validate_impl provides validation/actions for ExtRef
func Feature_ExtRef_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
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

// ----------------------------------------------------------------
// Feature_ExtRef2_validate_impl provides validation/actions for ExtRef2
func Feature_ExtRef2_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_ExtRef2_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   Feature_ExtRef2
// END   Feature_ExtRef2
// END   Feature_ExtRef2
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------
