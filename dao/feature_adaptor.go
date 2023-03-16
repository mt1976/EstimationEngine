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
// Date & Time		    : 05/03/2023 at 13:59:52
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------
// The following functions should be created in feature_impl.go
// ----------------------------------------------------------------

// ----------------------------------------------------------------

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
// END   Feature_FeatureID
// END   Feature_FeatureID
// END   Feature_FeatureID
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// BEGIN Feature_EstimationSessionID
// BEGIN Feature_EstimationSessionID
// BEGIN Feature_EstimationSessionID
// ----------------------------------------------------------------
// Feature_EstimationSessionID_OnStore_impl provides the implementation for the callout
func Feature_EstimationSessionID_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_EstimationSessionID_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_EstimationSessionID_OnFetch_impl provides the implementation for the callout
func Feature_EstimationSessionID_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_EstimationSessionID_scrn, GET, rec.FeatureID)
	return rec.EstimationSessionID
}

// ----------------------------------------------------------------
// END   Feature_EstimationSessionID
// END   Feature_EstimationSessionID
// END   Feature_EstimationSessionID
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_Name
// BEGIN Feature_Name
// BEGIN Feature_Name
// ----------------------------------------------------------------
// Feature_Name_OnStore_impl provides the implementation for the callout
func Feature_Name_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_Name_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_Name_OnFetch_impl provides the implementation for the callout
func Feature_Name_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_Name_scrn, GET, rec.FeatureID)
	return rec.Name
}

// ----------------------------------------------------------------
// END   Feature_Name
// END   Feature_Name
// END   Feature_Name
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_DevelopmentEstimate
// BEGIN Feature_DevelopmentEstimate
// BEGIN Feature_DevelopmentEstimate
// ----------------------------------------------------------------
// Feature_DevelopmentEstimate_OnStore_impl provides the implementation for the callout
func Feature_DevelopmentEstimate_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_DevelopmentEstimate_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_DevelopmentEstimate_OnFetch_impl provides the implementation for the callout
func Feature_DevelopmentEstimate_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_DevelopmentEstimate_scrn, GET, rec.FeatureID)
	return rec.DevelopmentEstimate
}

// ----------------------------------------------------------------
// END   Feature_DevelopmentEstimate
// END   Feature_DevelopmentEstimate
// END   Feature_DevelopmentEstimate
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_DevelopmentFactored
// BEGIN Feature_DevelopmentFactored
// BEGIN Feature_DevelopmentFactored
// ----------------------------------------------------------------
// Feature_DevelopmentFactored_OnStore_impl provides the implementation for the callout
func Feature_DevelopmentFactored_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_DevelopmentFactored_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_DevelopmentFactored_OnFetch_impl provides the implementation for the callout
func Feature_DevelopmentFactored_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_DevelopmentFactored_scrn, GET, rec.FeatureID)
	return rec.DevelopmentFactored
}

// ----------------------------------------------------------------
// END   Feature_DevelopmentFactored
// END   Feature_DevelopmentFactored
// END   Feature_DevelopmentFactored
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_Requirements
// BEGIN Feature_Requirements
// BEGIN Feature_Requirements
// ----------------------------------------------------------------
// Feature_Requirements_OnStore_impl provides the implementation for the callout
func Feature_Requirements_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_Requirements_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_Requirements_OnFetch_impl provides the implementation for the callout
func Feature_Requirements_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_Requirements_scrn, GET, rec.FeatureID)
	return rec.Requirements
}

// ----------------------------------------------------------------
// END   Feature_Requirements
// END   Feature_Requirements
// END   Feature_Requirements
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_AnalystTest
// BEGIN Feature_AnalystTest
// BEGIN Feature_AnalystTest
// ----------------------------------------------------------------
// Feature_AnalystTest_OnStore_impl provides the implementation for the callout
func Feature_AnalystTest_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_AnalystTest_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_AnalystTest_OnFetch_impl provides the implementation for the callout
func Feature_AnalystTest_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_AnalystTest_scrn, GET, rec.FeatureID)
	return rec.AnalystTest
}

// ----------------------------------------------------------------
// END   Feature_AnalystTest
// END   Feature_AnalystTest
// END   Feature_AnalystTest
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_Documentation
// BEGIN Feature_Documentation
// BEGIN Feature_Documentation
// ----------------------------------------------------------------
// Feature_Documentation_OnStore_impl provides the implementation for the callout
func Feature_Documentation_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_Documentation_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_Documentation_OnFetch_impl provides the implementation for the callout
func Feature_Documentation_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_Documentation_scrn, GET, rec.FeatureID)
	return rec.Documentation
}

// ----------------------------------------------------------------
// END   Feature_Documentation
// END   Feature_Documentation
// END   Feature_Documentation
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_Management
// BEGIN Feature_Management
// BEGIN Feature_Management
// ----------------------------------------------------------------
// Feature_Management_OnStore_impl provides the implementation for the callout
func Feature_Management_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_Management_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_Management_OnFetch_impl provides the implementation for the callout
func Feature_Management_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_Management_scrn, GET, rec.FeatureID)
	return rec.Management
}

// ----------------------------------------------------------------
// END   Feature_Management
// END   Feature_Management
// END   Feature_Management
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_Marketing
// BEGIN Feature_Marketing
// BEGIN Feature_Marketing
// ----------------------------------------------------------------
// Feature_Marketing_OnStore_impl provides the implementation for the callout
func Feature_Marketing_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_Marketing_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_Marketing_OnFetch_impl provides the implementation for the callout
func Feature_Marketing_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_Marketing_scrn, GET, rec.FeatureID)
	return rec.Marketing
}

// ----------------------------------------------------------------
// END   Feature_Marketing
// END   Feature_Marketing
// END   Feature_Marketing
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_Contingency
// BEGIN Feature_Contingency
// BEGIN Feature_Contingency
// ----------------------------------------------------------------
// Feature_Contingency_OnStore_impl provides the implementation for the callout
func Feature_Contingency_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_Contingency_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_Contingency_OnFetch_impl provides the implementation for the callout
func Feature_Contingency_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_Contingency_scrn, GET, rec.FeatureID)
	return rec.Contingency
}

// ----------------------------------------------------------------
// END   Feature_Contingency
// END   Feature_Contingency
// END   Feature_Contingency
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_TrackerID
// BEGIN Feature_TrackerID
// BEGIN Feature_TrackerID
// ----------------------------------------------------------------
// Feature_TrackerID_OnStore_impl provides the implementation for the callout
//
//	func Feature_TrackerID_OnStore_impl (fieldval string,rec dm.Feature,usr string) (string,error)  {
//		logs.Callout("Feature", dm.Feature_TrackerID_scrn, PUT, rec.FeatureID)
//		return fieldval,nil
//	}
//
// ----------------------------------------------------------------
// Feature_TrackerID_OnFetch_impl provides the implementation for the callout
func Feature_TrackerID_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_TrackerID_scrn, GET, rec.FeatureID)
	return rec.TrackerID
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
//
//	func Feature_AdoID_OnStore_impl (fieldval string,rec dm.Feature,usr string) (string,error)  {
//		logs.Callout("Feature", dm.Feature_AdoID_scrn, PUT, rec.FeatureID)
//		return fieldval,nil
//	}
//
// ----------------------------------------------------------------
// Feature_AdoID_OnFetch_impl provides the implementation for the callout
func Feature_AdoID_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_AdoID_scrn, GET, rec.FeatureID)
	return rec.AdoID
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
//
//	func Feature_FreshdeskID_OnStore_impl (fieldval string,rec dm.Feature,usr string) (string,error)  {
//		logs.Callout("Feature", dm.Feature_FreshdeskID_scrn, PUT, rec.FeatureID)
//		return fieldval,nil
//	}
//
// ----------------------------------------------------------------
// Feature_FreshdeskID_OnFetch_impl provides the implementation for the callout
func Feature_FreshdeskID_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_FreshdeskID_scrn, GET, rec.FeatureID)
	return rec.FreshdeskID
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
//
//	func Feature_ExtRef_OnStore_impl (fieldval string,rec dm.Feature,usr string) (string,error)  {
//		logs.Callout("Feature", dm.Feature_ExtRef_scrn, PUT, rec.FeatureID)
//		return fieldval,nil
//	}
//
// ----------------------------------------------------------------
// Feature_ExtRef_OnFetch_impl provides the implementation for the callout
func Feature_ExtRef_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_ExtRef_scrn, GET, rec.FeatureID)
	return rec.ExtRef
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
//
//	func Feature_ExtRef2_OnStore_impl (fieldval string,rec dm.Feature,usr string) (string,error)  {
//		logs.Callout("Feature", dm.Feature_ExtRef2_scrn, PUT, rec.FeatureID)
//		return fieldval,nil
//	}
//
// ----------------------------------------------------------------
// Feature_ExtRef2_OnFetch_impl provides the implementation for the callout
func Feature_ExtRef2_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_ExtRef2_scrn, GET, rec.FeatureID)
	return rec.ExtRef2
}

// ----------------------------------------------------------------
// END   Feature_ExtRef2
// END   Feature_ExtRef2
// END   Feature_ExtRef2
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_DeveloperResource
// BEGIN Feature_DeveloperResource
// BEGIN Feature_DeveloperResource
// ----------------------------------------------------------------
// Feature_DeveloperResource_OnStore_impl provides the implementation for the callout
func Feature_DeveloperResource_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_DeveloperResource_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_DeveloperResource_OnFetch_impl provides the implementation for the callout
func Feature_DeveloperResource_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_DeveloperResource_scrn, GET, rec.FeatureID)
	return rec.DeveloperResource
}

// ----------------------------------------------------------------
// END   Feature_DeveloperResource
// END   Feature_DeveloperResource
// END   Feature_DeveloperResource
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_ApproverResource
// BEGIN Feature_ApproverResource
// BEGIN Feature_ApproverResource
// ----------------------------------------------------------------
// Feature_ApproverResource_OnStore_impl provides the implementation for the callout
func Feature_ApproverResource_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_ApproverResource_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_ApproverResource_OnFetch_impl provides the implementation for the callout
func Feature_ApproverResource_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_ApproverResource_scrn, GET, rec.FeatureID)
	return rec.ApproverResource
}

// ----------------------------------------------------------------
// END   Feature_ApproverResource
// END   Feature_ApproverResource
// END   Feature_ApproverResource
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_OffProfileJustification
// BEGIN Feature_OffProfileJustification
// BEGIN Feature_OffProfileJustification
// ----------------------------------------------------------------
// Feature_OffProfileJustification_OnStore_impl provides the implementation for the callout
func Feature_OffProfileJustification_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_OffProfileJustification_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_OffProfileJustification_OnFetch_impl provides the implementation for the callout
func Feature_OffProfileJustification_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_OffProfileJustification_scrn, GET, rec.FeatureID)
	return rec.OffProfileJustification
}

// ----------------------------------------------------------------
// END   Feature_OffProfileJustification
// END   Feature_OffProfileJustification
// END   Feature_OffProfileJustification
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_Total
// BEGIN Feature_Total
// BEGIN Feature_Total
// ----------------------------------------------------------------
// Feature_Total_OnStore_impl provides the implementation for the callout
func Feature_Total_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_Total_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_Total_OnFetch_impl provides the implementation for the callout
func Feature_Total_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_Total_scrn, GET, rec.FeatureID)
	return rec.Total
}

// ----------------------------------------------------------------
// END   Feature_Total
// END   Feature_Total
// END   Feature_Total
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_AnalystResource
// BEGIN Feature_AnalystResource
// BEGIN Feature_AnalystResource
// ----------------------------------------------------------------
// Feature_AnalystResource_OnStore_impl provides the implementation for the callout
func Feature_AnalystResource_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_AnalystResource_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_AnalystResource_OnFetch_impl provides the implementation for the callout
func Feature_AnalystResource_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_AnalystResource_scrn, GET, rec.FeatureID)
	return rec.AnalystResource
}

// ----------------------------------------------------------------
// END   Feature_AnalystResource
// END   Feature_AnalystResource
// END   Feature_AnalystResource
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_Training
// BEGIN Feature_Training
// BEGIN Feature_Training
// ----------------------------------------------------------------
// Feature_Training_OnStore_impl provides the implementation for the callout
func Feature_Training_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_Training_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_Training_OnFetch_impl provides the implementation for the callout
func Feature_Training_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_Training_scrn, GET, rec.FeatureID)
	return rec.Training
}

// ----------------------------------------------------------------
// END   Feature_Training
// END   Feature_Training
// END   Feature_Training
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_EstimateEffort
// BEGIN Feature_EstimateEffort
// BEGIN Feature_EstimateEffort
// ----------------------------------------------------------------
// Feature_EstimateEffort_OnStore_impl provides the implementation for the callout
func Feature_EstimateEffort_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_EstimateEffort_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_EstimateEffort_OnFetch_impl provides the implementation for the callout
func Feature_EstimateEffort_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_EstimateEffort_scrn, GET, rec.FeatureID)
	return rec.EstimateEffort
}

// ----------------------------------------------------------------
// END   Feature_EstimateEffort
// END   Feature_EstimateEffort
// END   Feature_EstimateEffort
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_AnalystEstimate
// BEGIN Feature_AnalystEstimate
// BEGIN Feature_AnalystEstimate
// ----------------------------------------------------------------
// Feature_AnalystEstimate_OnStore_impl provides the implementation for the callout
func Feature_AnalystEstimate_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_AnalystEstimate_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_AnalystEstimate_OnFetch_impl provides the implementation for the callout
func Feature_AnalystEstimate_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_AnalystEstimate_scrn, GET, rec.FeatureID)
	return rec.AnalystEstimate
}

// ----------------------------------------------------------------
// END   Feature_AnalystEstimate
// END   Feature_AnalystEstimate
// END   Feature_AnalystEstimate
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// BEGIN Feature_OriginName
// BEGIN Feature_OriginName
// BEGIN Feature_OriginName
// ----------------------------------------------------------------
// Feature_OriginName_OnStore_impl provides the implementation for the callout
func Feature_OriginName_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_OriginName_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_OriginName_OnFetch_impl provides the implementation for the callout
func Feature_OriginName_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_OriginName_scrn, GET, rec.FeatureID)

	estID := rec.EstimationSessionID
	est := CacheRead(dm.EstimationSession_Name, estID).(dm.EstimationSession)
	// _, est, err := EstimationSession_GetByID(estID)
	// if err != nil {
	// 	return ""
	// }
	projID := est.ProjectID
	proj := CacheRead(dm.Project_Name, projID).(dm.Project)
	// _, proj, err := Project_GetByID(projID)
	// if err != nil {
	// 	return ""
	// }

	originID := proj.OriginID
	origin := CacheRead(dm.Origin_Name, originID).(dm.Origin)
	// _, origin, err := Origin_GetByCode(originID)
	// if err != nil {
	// 	return ""
	// }
	retVal := "(" + origin.Code + ") " + origin.FullName
	return retVal
}

// ----------------------------------------------------------------
// END   Feature_OriginName
// END   Feature_OriginName
// END   Feature_OriginName
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_ProjectName
// BEGIN Feature_ProjectName
// BEGIN Feature_ProjectName
// ----------------------------------------------------------------
// Feature_ProjectName_OnStore_impl provides the implementation for the callout
func Feature_ProjectName_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_ProjectName_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_ProjectName_OnFetch_impl provides the implementation for the callout
func Feature_ProjectName_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_ProjectName_scrn, GET, rec.FeatureID)

	// Get the project from the estimationsession
	estID := rec.EstimationSessionID
	est := CacheRead(dm.EstimationSession_Name, estID).(dm.EstimationSession)
	// _, est, err := EstimationSession_GetByID(estID)
	// if err != nil {
	// 	return ""
	// }
	projID := est.ProjectID
	proj := CacheRead(dm.Project_Name, projID).(dm.Project)
	// _, proj, err := Project_GetByID(projID)
	// if err != nil {
	// 	return ""
	// }

	return proj.Name
}

// ----------------------------------------------------------------
// END   Feature_ProjectName
// END   Feature_ProjectName
// END   Feature_ProjectName
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// BEGIN Feature_ProfileSelectedOld
// BEGIN Feature_ProfileSelectedOld
// BEGIN Feature_ProfileSelectedOld
// ----------------------------------------------------------------
// Feature_ProfileSelectedOld_OnStore_impl provides the implementation for the callout
func Feature_ProfileSelectedOld_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_ProfileSelectedOld_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_ProfileSelectedOld_OnFetch_impl provides the implementation for the callout
func Feature_ProfileSelectedOld_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_ProfileSelectedOld_scrn, GET, rec.FeatureID)
	rec.ProfileSelectedOld = rec.ProfileSelected
	return rec.ProfileSelectedOld
}

// ----------------------------------------------------------------
// END   Feature_ProfileSelectedOld
// END   Feature_ProfileSelectedOld
// END   Feature_ProfileSelectedOld
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// BEGIN Feature_CCY
// BEGIN Feature_CCY
// BEGIN Feature_CCY
// ----------------------------------------------------------------
// Feature_CCY_OnStore_impl provides the implementation for the callout
func Feature_CCY_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_CCY_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_CCY_OnFetch_impl provides the implementation for the callout
func Feature_CCY_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_CCY_scrn, GET, rec.FeatureID)
	//	rec.CCY = "TST"
	//	return ""

	esID := rec.EstimationSessionID
	//logs.Information("Cache Read : EstimationSession", esID)
	es := CacheRead(dm.EstimationSession_Name, esID).(dm.EstimationSession)
	// _, es, err := EstimationSession_GetByID(esID)
	// if err != nil {
	// 	return "GBP"
	// }
	projID := es.ProjectID
	//logs.Information("Cache Read : Project", projID)
	proj := CacheRead(dm.Project_Name, projID).(dm.Project)
	// _, proj, err := Project_GetByID(projID)
	// if err != nil {
	// 	return "GBP"
	// }
	origID := proj.OriginID
	//logs.Information("Cache Read : Origin", origID)
	orig := CacheRead(dm.Origin_Name, origID).(dm.Origin)
	// _, orig, err := Origin_GetByCode(origID)
	// if err != nil {
	// 	return "GBP"
	// }
	//rec.CCY = orig.Currency

	return orig.Currency
}

// ----------------------------------------------------------------
// END   Feature_CCY
// END   Feature_CCY
// END   Feature_CCY
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Feature_RoundTo
// BEGIN Feature_RoundTo
// BEGIN Feature_RoundTo
// ----------------------------------------------------------------
// // Feature_RoundTo_OnStore_impl provides the implementation for the callout
// func Feature_RoundTo_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
// 	logs.Callout("Feature", dm.Feature_RoundedTo_scrn, PUT, rec.FeatureID)
// 	return fieldval, nil
// }

// // ----------------------------------------------------------------
// // Feature_RoundTo_OnFetch_impl provides the implementation for the callout
// func Feature_RoundTo_OnFetch_impl(rec dm.Feature) string {
// 	logs.Callout("Feature", dm.Feature_RoundedTo_scrn, GET, rec.FeatureID)
// 	//core.GetApplicationProperty("roundhoursto")
// 	return core.GetApplicationProperty("roundhoursto")
// }

func feature_PostPutAction_impl(id string, rec dm.Feature, usr string) error {
	logs.Callout("Feature", "PostPutAction", PUT, rec.FeatureID)
	trigger := "Feature Store " + rec.Name + " (" + rec.FeatureID + ")"
	go Estimationsession_Calculate(rec.EstimationSessionID, trigger)
	return nil
}
