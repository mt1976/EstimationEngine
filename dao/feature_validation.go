package dao

import (
	"fmt"

	"github.com/mt1976/ebEstimates/core"
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
// Date & Time		    : 05/03/2023 at 13:59:52
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

		if iRec.ProfileSelectedOld != iRec.ProfileSelected {
			iRec = Feature_CalcDefaults(iRec, false, "")
		}

		offProfile := core.FALSE
		offProfile, iRec.Activity = feature_ValidateEstimate(iRec.DevelopmentFactored, iRec.DevelopmentFactoredDefault, offProfile, iRec.Activity, "Dev + Uplift")
		offProfile, iRec.Activity = feature_ValidateEstimate(iRec.Requirements, iRec.RequirementsDefault, offProfile, iRec.Activity, "Requirements")
		if iRec.AnalystEstimate == "" {
			offProfile, iRec.Activity = feature_ValidateEstimate(iRec.AnalystTest, iRec.AnalystTestDefault, offProfile, iRec.Activity, "Analyst Test")
		}
		offProfile, iRec.Activity = feature_ValidateEstimate(iRec.Documentation, iRec.DocumentationDefault, offProfile, iRec.Activity, "Documentation")
		offProfile, iRec.Activity = feature_ValidateEstimate(iRec.Management, iRec.ManagementDefault, offProfile, iRec.Activity, "Management")

		offProfile, iRec.Activity = feature_ValidateEstimate(iRec.UatSupport, iRec.UatSupportDefault, offProfile, iRec.Activity, "UAT Support")
		offProfile, iRec.Activity = feature_ValidateEstimate(iRec.Training, iRec.TrainingDefault, offProfile, iRec.Activity, "Training")
		offProfile, iRec.Activity = feature_ValidateEstimate(iRec.Marketing, iRec.MarketingDefault, offProfile, iRec.Activity, "Marketing")
		offProfile, iRec.Activity = feature_ValidateEstimate(iRec.Contingency, iRec.ContingencyDefault, offProfile, iRec.Activity, "Contingency")

		//iRec.Total = feature_CalculateTotal(iRec)

		if offProfile == core.TRUE {
			iRec.OffProfile = core.TRUE
			mTXT_1 := Translate("AuditMessage", "ESTIMATE OFF PROFILE")
			iRec.Activity = core.AddActivity(iRec.Activity, mTXT_1)
		}

		if iRec.OffProfileJustification != "" {
			mTXT_2 := Translate("AuditMessage", "JUSTIFICATION [%s]")
			mTXT_2 = fmt.Sprintf(mTXT_2, iRec.OffProfileJustification)
			iRec.Activity = core.AddActivity(iRec.Activity, mTXT_2)
		}
		if iRec.ApproverResource != "" {
			mTXT_3 := Translate("AuditMessage", "APPROVED BY [%s]")
			mTXT_3 = fmt.Sprintf(mTXT_3, iRec.ApproverResource)
			iRec.Activity = core.AddActivity(iRec.Activity, mTXT_3)
			mTXT_4 := Translate("AuditMessage", "Feature %s has been approved by %s.")
			mTXT_4 = fmt.Sprintf(mTXT_4, iRec.Name, iRec.ApproverResource)
			mSBJ := Translate("AuditMessage", "Feature Approved - %s")
			mSBJ = fmt.Sprintf(mSBJ, iRec.Name)
			Resource_SendMail(iRec.ApproverResource, mSBJ, mTXT_4)
			Resource_SendMail(iRec.ProjectManagerResource, mSBJ, mTXT_4)
		}

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
// Feature_FeatureID_validate_impl provides validation/actions for FeatureID
func Feature_FeatureID_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_FeatureID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_EstimationSessionID_validate_impl provides validation/actions for EstimationSessionID
func Feature_EstimationSessionID_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_EstimationSessionID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_Name_validate_impl provides validation/actions for Name
func Feature_Name_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_Name_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_DevelopmentEstimate_validate_impl provides validation/actions for DevelopmentEstimate
func Feature_DevelopmentEstimate_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_DevelopmentEstimate_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_DevelopmentFactored_validate_impl provides validation/actions for DevelopmentFactored
func Feature_DevelopmentFactored_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_DevelopmentFactored_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_Requirements_validate_impl provides validation/actions for Requirements
func Feature_Requirements_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_Requirements_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_AnalystTest_validate_impl provides validation/actions for AnalystTest
func Feature_AnalystTest_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_AnalystTest_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_Documentation_validate_impl provides validation/actions for Documentation
func Feature_Documentation_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_Documentation_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_Management_validate_impl provides validation/actions for Management
func Feature_Management_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_Management_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_Marketing_validate_impl provides validation/actions for Marketing
func Feature_Marketing_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_Marketing_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_Contingency_validate_impl provides validation/actions for Contingency
func Feature_Contingency_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_Contingency_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_TrackerID_validate_impl provides validation/actions for TrackerID
func Feature_TrackerID_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_TrackerID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_AdoID_validate_impl provides validation/actions for AdoID
func Feature_AdoID_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_AdoID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_FreshdeskID_validate_impl provides validation/actions for FreshdeskID
func Feature_FreshdeskID_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_FreshdeskID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_ExtRef_validate_impl provides validation/actions for ExtRef
func Feature_ExtRef_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_ExtRef_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_ExtRef2_validate_impl provides validation/actions for ExtRef2
func Feature_ExtRef2_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_ExtRef2_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_DeveloperResource_validate_impl provides validation/actions for DeveloperResource
func Feature_DeveloperResource_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_DeveloperResource_scrn, VAL+"-"+iAction, iId)
	if iAction == PUT {
		if iValue == "" {
			fP := SetFieldError(fP, "developer is required")
			return iValue, fP
		}
		fP = IsValidResource(iValue, fP)
	}
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_ApproverResource_validate_impl provides validation/actions for ApproverResource
func Feature_ApproverResource_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_ApproverResource_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_OffProfileJustification_validate_impl provides validation/actions for OffProfileJustification
func Feature_OffProfileJustification_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_OffProfileJustification_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_Total_validate_impl provides validation/actions for Total
func Feature_Total_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_Total_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_AnalystResource_validate_impl provides validation/actions for AnalystResource
func Feature_AnalystResource_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_AnalystResource_scrn, VAL+"-"+iAction, iId)
	if iAction == PUT {
		if iValue == "" {
			fP := SetFieldError(fP, "analyst is required")
			return iValue, fP
		}
		fP = IsValidResource(iValue, fP)
	}
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_Training_validate_impl provides validation/actions for Training
func Feature_Training_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_Training_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_EstimateEffort_validate_impl provides validation/actions for EstimateEffort
func Feature_EstimateEffort_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_EstimateEffort_scrn, VAL+"-"+iAction, iId)
	if iAction == PUT {
		if iValue == "" {
			//fP = SetFieldError(fP, "estimate effort is required")
			iValue = "0"
		}
	}
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_AnalystEstimate_validate_impl provides validation/actions for AnalystEstimate
func Feature_AnalystEstimate_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_AnalystEstimate_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_OriginName_validate_impl provides validation/actions for OriginName
func Feature_OriginName_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_OriginName_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_ProjectName_validate_impl provides validation/actions for ProjectName
func Feature_ProjectName_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_ProjectName_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// Feature_ProfileSelectedOld_validate_impl provides validation/actions for ProfileSelectedOld
func Feature_ProfileSelectedOld_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_ProfileSelectedOld_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_CCY_validate_impl provides validation/actions for CCY
func Feature_CCY_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_CCY_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// // ----------------------------------------------------------------
// // Feature_RoundTo_validate_impl provides validation/actions for RoundTo
// func Feature_RoundTo_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
// 	logs.Callout("Feature", dm.Feature_RoundTo_scrn, VAL+"-"+iAction, iId)
// 	return iValue, fP
// }

// ----------------------------------------------------------------
// Feature_OriginID_validate_impl provides validation/actions for OriginID
func Feature_OriginID_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_OriginID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_ProjectID_validate_impl provides validation/actions for ProjectID
func Feature_ProjectID_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_ProjectID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_OriginCode_validate_impl provides validation/actions for OriginCode
func Feature_OriginCode_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_OriginCode_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_EstimationSessionName_validate_impl provides validation/actions for EstimationSessionName
func Feature_EstimationSessionName_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_EstimationSessionName_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_RequirementsPerc_validate_impl provides validation/actions for RequirementsPerc
func Feature_RequirementsPerc_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_RequirementsPerc_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_AnalystTestPerc_validate_impl provides validation/actions for AnalystTestPerc
func Feature_AnalystTestPerc_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_AnalystTestPerc_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_DocumentationPerc_validate_impl provides validation/actions for DocumentationPerc
func Feature_DocumentationPerc_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_DocumentationPerc_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_ManagementPerc_validate_impl provides validation/actions for ManagementPerc
func Feature_ManagementPerc_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_ManagementPerc_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_UatSupportPerc_validate_impl provides validation/actions for UatSupportPerc
func Feature_UatSupportPerc_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_UatSupportPerc_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_MarketingPerc_validate_impl provides validation/actions for MarketingPerc
func Feature_MarketingPerc_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_MarketingPerc_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_ContingencyPerc_validate_impl provides validation/actions for ContingencyPerc
func Feature_ContingencyPerc_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_ContingencyPerc_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_TrainingPerc_validate_impl provides validation/actions for TrainingPerc
func Feature_TrainingPerc_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_TrainingPerc_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// Feature_RoundedTo_validate_impl provides validation/actions for RoundedTo
func Feature_RoundedTo_validate_impl(iAction string, iId string, iValue string, iRec dm.Feature, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Feature", dm.Feature_RoundedTo_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}
