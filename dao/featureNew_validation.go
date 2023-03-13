package dao

import (
	core "github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/featurenew.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : FeatureNew (featurenew)
// Endpoint 	        : FeatureNew (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 05/03/2023 at 13:03:29
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------
// The following functions should be created in featurenew_impl.go
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// FeatureNew_ObjectValidation_impl provides Record/Object level validation for FeatureNew
func FeatureNew_ObjectValidation_impl(iAction string, iId string, iRec dm.FeatureNew) (dm.FeatureNew, string, error) {
	logs.Warning("FeatureNew-ObjectValidation " + VAL + "-" + iAction + "-" + iId)
	logs.Warning("FeatureNew-ObjectValidation " + VAL + "-" + iAction + "-" + iId)
	logs.Warning("FeatureNew-ObjectValidation " + VAL + "-" + iAction + "-" + iId)
	logs.Warning("FeatureNew-ObjectValidation " + VAL + "-" + iAction + "-" + iId)
	logs.Warning("FeatureNew-ObjectValidation " + VAL + "-" + iAction + "-" + iId)

	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("FeatureNew" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}

//
// ----------------------------------------------------------------
// These are the the default implementations, which do nothing

// ----------------------------------------------------------------
// FeatureNew_DeveloperEstimate_validate_impl provides validation/actions for DeveloperEstimate
func FeatureNew_DeveloperEstimate_validate_impl(iAction string, iId string, iValue string, iRec dm.FeatureNew, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("FeatureNew", dm.FeatureNew_DeveloperEstimate_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// FeatureNew_DeveloperResource_validate_impl provides validation/actions for DeveloperResource
func FeatureNew_DeveloperResource_validate_impl(iAction string, iId string, iValue string, iRec dm.FeatureNew, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("FeatureNew", dm.FeatureNew_DeveloperResource_scrn, VAL+"-"+iAction, iId)
	if iValue == "" {
		fP = SetFieldError(fP, "engineer is required")
	}
	//fP = IsValidResource(iValue, fP)

	return iValue, fP
}

// ----------------------------------------------------------------
// FeatureNew_AnalystResource_validate_impl provides validation/actions for AnalystResource
func FeatureNew_AnalystResource_validate_impl(iAction string, iId string, iValue string, iRec dm.FeatureNew, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("FeatureNew", dm.FeatureNew_AnalystResource_scrn, VAL+"-"+iAction, iId)
	//fP = IsValidResource(iValue, fP)
	return iValue, fP
}

// ----------------------------------------------------------------
// FeatureNew_EstimateEffort_validate_impl provides validation/actions for EstimateEffort
func FeatureNew_EstimateEffort_validate_impl(iAction string, iId string, iValue string, iRec dm.FeatureNew, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("FeatureNew", dm.FeatureNew_EstimateEffort_scrn, VAL+"-"+iAction, iId)
	if iAction == VAL {
		if iValue != "" {

			miniumum, errData := Data_GetFloat("Feature", "Estimation_Effort_Minimum", dm.Data_Category_Setting)
			if errData != nil {
				logs.Warning("Unable to get minimum effort from data - assuming 0")
				return iValue, fP
			}
			// If it's not blank, then it must be a valid number
			effort, err := core.Numeric(iValue)
			if err != nil {
				fP = SetFieldError(fP, "must be a valid number")
				return iValue, fP
			}
			if effort < miniumum {
				fP = SetFieldError(fP, "must be greater than "+core.FloatToString(miniumum))
				return iValue, fP
			}
		}
	}
	return iValue, fP
}

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------
// FeatureNew_EstimationSessionName_validate_impl provides validation/actions for EstimationSessionName
func FeatureNew_EstimationSessionName_validate_impl(iAction string, iId string, iValue string, iRec dm.FeatureNew, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("FeatureNew", dm.FeatureNew_EstimationSessionName_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// FeatureNew_ProjectID_validate_impl provides validation/actions for ProjectID
func FeatureNew_ProjectID_validate_impl(iAction string, iId string, iValue string, iRec dm.FeatureNew, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("FeatureNew", dm.FeatureNew_ProjectID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// FeatureNew_OriginCode_validate_impl provides validation/actions for OriginCode
func FeatureNew_OriginCode_validate_impl(iAction string, iId string, iValue string, iRec dm.FeatureNew, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("FeatureNew", dm.FeatureNew_OriginCode_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// FeatureNew_OriginID_validate_impl provides validation/actions for OriginID
func FeatureNew_OriginID_validate_impl(iAction string, iId string, iValue string, iRec dm.FeatureNew, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("FeatureNew", dm.FeatureNew_OriginID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// FeatureNew_ProjectName_validate_impl provides validation/actions for ProjectName
func FeatureNew_ProjectName_validate_impl(iAction string, iId string, iValue string, iRec dm.FeatureNew, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("FeatureNew", dm.FeatureNew_ProjectName_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// FeatureNew_ConfidenceCODE_validate_impl provides validation/actions for ConfidenceCODE
func FeatureNew_ConfidenceCODE_validate_impl(iAction string, iId string, iValue string, iRec dm.FeatureNew, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("FeatureNew", dm.FeatureNew_ConfidenceCODE_scrn, VAL+"-"+iAction, iId)
	if iValue == "" {
		fP = SetFieldError(fP, "confidence is required")
	}
	return iValue, fP
}
