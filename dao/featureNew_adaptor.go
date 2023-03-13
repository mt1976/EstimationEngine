package dao

import (
	"github.com/mt1976/ebEstimates/core"
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
// FeatureNew_GetList_impl provides the implementation to get a list of records
func FeatureNew_GetList_impl(filter string) (int, []dm.FeatureNew, error) { return 0, nil, nil }

// ----------------------------------------------------------------
// FeatureNew_GetByID_impl provides the implementation to get a record by ID
func FeatureNew_GetByID_impl(id string) (int, dm.FeatureNew, error) { return 0, dm.FeatureNew{}, nil }

//

// ----------------------------------------------------------------

// ----------------------------------------------------------------
// FeatureNew_NewID_impl provides the implementation to generate a new ID for a new record
func FeatureNew_NewID_impl(rec dm.FeatureNew) string { return rec.ID }

// ----------------------------------------------------------------
// FeatureNew_New_impl provides the implementation to delete a new record
func FeatureNew_Delete_impl(id string) error { return nil }

// ----------------------------------------------------------------

// ----------------------------------------------------------------
// FeatureNew_Store_impl provides the implementation for the store adaptor
func FeatureNew_Update_impl(id string, rec dm.FeatureNew, usr string) error {

	var item dm.Feature
	// START
	// Dynamically generated 29/11/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	newID := Feature_NewID(item)

	logs.Processing("New ID: " + newID)

	item.FeatureID = newID
	item.EstimationSessionID = rec.EstimationSessionID
	item.Name = rec.Name
	item.DevelopmentEstimate = rec.DeveloperEstimate
	item.ConfidenceCODE = rec.ConfidenceCODE
	item.DeveloperResource = rec.DeveloperResource

	item.EstimationSessionID = rec.EstimationSessionID
	item.Name = rec.Name
	item.DevelopmentEstimate = rec.DeveloperEstimate
	item.ConfidenceCODE = rec.ConfidenceCODE
	item.DeveloperResource = rec.DeveloperResource

	item.AdoID = rec.DevOpsID
	item.FreshdeskID = rec.FreshDeskID
	item.TrackerID = rec.RSCID
	item.ExtRef = rec.OtherID
	item.ExtRef2 = rec.OtherID2
	item.ProfileSelected = rec.ProfileSelected
	item.ProfileDefault = rec.ProfileDefault

	item = Feature_CalcDefaults(item, false)
	//
	// Dynamically generated 29/11/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	//spew.Dump(item)

	//tm := time.Now().Format("02/01/2006 ")
	//item.Notes = tm + Session_GetUserName(r) + ": Created"

	msgTXT := Translate("Message", "New Feature Created")
	item.Activity = core.AddActivity_ForUser(item.Activity, msgTXT, usr)

	//Feature_StoreSystem(item)
	Feature_StoreProcess(item, usr)

	//Estimationsession_Calculate(item.EstimationSessionID)

	return nil
}

// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN FeatureNew_DeveloperEstimate
// BEGIN FeatureNew_DeveloperEstimate
// BEGIN FeatureNew_DeveloperEstimate
// ----------------------------------------------------------------
// FeatureNew_DeveloperEstimate_OnStore_impl provides the implementation for the callout
func FeatureNew_DeveloperEstimate_OnStore_impl(fieldval string, rec dm.FeatureNew, usr string) (string, error) {
	logs.Callout("FeatureNew", dm.FeatureNew_DeveloperEstimate_scrn, PUT, rec.ID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// FeatureNew_DeveloperEstimate_OnFetch_impl provides the implementation for the callout
func FeatureNew_DeveloperEstimate_OnFetch_impl(rec dm.FeatureNew) string {
	logs.Callout("FeatureNew", dm.FeatureNew_DeveloperEstimate_scrn, GET, rec.ID)
	return rec.DeveloperEstimate
}

// ----------------------------------------------------------------
// END   FeatureNew_DeveloperEstimate
// END   FeatureNew_DeveloperEstimate
// END   FeatureNew_DeveloperEstimate
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN FeatureNew_DeveloperResource
// BEGIN FeatureNew_DeveloperResource
// BEGIN FeatureNew_DeveloperResource
// ----------------------------------------------------------------
// FeatureNew_DeveloperResource_OnStore_impl provides the implementation for the callout
func FeatureNew_DeveloperResource_OnStore_impl(fieldval string, rec dm.FeatureNew, usr string) (string, error) {
	logs.Callout("FeatureNew", dm.FeatureNew_DeveloperResource_scrn, PUT, rec.ID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// FeatureNew_DeveloperResource_OnFetch_impl provides the implementation for the callout
func FeatureNew_DeveloperResource_OnFetch_impl(rec dm.FeatureNew) string {
	logs.Callout("FeatureNew", dm.FeatureNew_DeveloperResource_scrn, GET, rec.ID)
	return rec.DeveloperResource
}

// ----------------------------------------------------------------
// END   FeatureNew_DeveloperResource
// END   FeatureNew_DeveloperResource
// END   FeatureNew_DeveloperResource
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN FeatureNew_AnalystResource
// BEGIN FeatureNew_AnalystResource
// BEGIN FeatureNew_AnalystResource
// ----------------------------------------------------------------
// FeatureNew_AnalystResource_OnStore_impl provides the implementation for the callout
func FeatureNew_AnalystResource_OnStore_impl(fieldval string, rec dm.FeatureNew, usr string) (string, error) {
	logs.Callout("FeatureNew", dm.FeatureNew_AnalystResource_scrn, PUT, rec.ID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// FeatureNew_AnalystResource_OnFetch_impl provides the implementation for the callout
func FeatureNew_AnalystResource_OnFetch_impl(rec dm.FeatureNew) string {
	logs.Callout("FeatureNew", dm.FeatureNew_AnalystResource_scrn, GET, rec.ID)
	return rec.AnalystResource
}

// ----------------------------------------------------------------
// END   FeatureNew_AnalystResource
// END   FeatureNew_AnalystResource
// END   FeatureNew_AnalystResource
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN FeatureNew_EstimateEffort
// BEGIN FeatureNew_EstimateEffort
// BEGIN FeatureNew_EstimateEffort
// ----------------------------------------------------------------
// FeatureNew_EstimateEffort_OnStore_impl provides the implementation for the callout
func FeatureNew_EstimateEffort_OnStore_impl(fieldval string, rec dm.FeatureNew, usr string) (string, error) {
	logs.Callout("FeatureNew", dm.FeatureNew_EstimateEffort_scrn, PUT, rec.ID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// FeatureNew_EstimateEffort_OnFetch_impl provides the implementation for the callout
func FeatureNew_EstimateEffort_OnFetch_impl(rec dm.FeatureNew) string {
	logs.Callout("FeatureNew", dm.FeatureNew_EstimateEffort_scrn, GET, rec.ID)
	return rec.EstimateEffort
}

// ----------------------------------------------------------------
// END   FeatureNew_EstimateEffort
// END   FeatureNew_EstimateEffort
// END   FeatureNew_EstimateEffort
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// BEGIN FeatureNew_EstimationSessionName
// BEGIN FeatureNew_EstimationSessionName
// BEGIN FeatureNew_EstimationSessionName
// ----------------------------------------------------------------
// FeatureNew_EstimationSessionName_OnStore_impl provides the implementation for the callout
func FeatureNew_EstimationSessionName_OnStore_impl(fieldval string, rec dm.FeatureNew, usr string) (string, error) {
	logs.Callout("FeatureNew", dm.FeatureNew_EstimationSessionName_scrn, PUT, rec.ID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// FeatureNew_EstimationSessionName_OnFetch_impl provides the implementation for the callout
func FeatureNew_EstimationSessionName_OnFetch_impl(rec dm.FeatureNew) string {
	logs.Callout("FeatureNew", dm.FeatureNew_EstimationSessionName_scrn, GET, rec.ID)
	return rec.EstimationSessionName
}

// ----------------------------------------------------------------
// END   FeatureNew_EstimationSessionName
// END   FeatureNew_EstimationSessionName
// END   FeatureNew_EstimationSessionName
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN FeatureNew_ProjectID
// BEGIN FeatureNew_ProjectID
// BEGIN FeatureNew_ProjectID
// ----------------------------------------------------------------
// FeatureNew_ProjectID_OnStore_impl provides the implementation for the callout
func FeatureNew_ProjectID_OnStore_impl(fieldval string, rec dm.FeatureNew, usr string) (string, error) {
	logs.Callout("FeatureNew", dm.FeatureNew_ProjectID_scrn, PUT, rec.ID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// FeatureNew_ProjectID_OnFetch_impl provides the implementation for the callout
func FeatureNew_ProjectID_OnFetch_impl(rec dm.FeatureNew) string {
	logs.Callout("FeatureNew", dm.FeatureNew_ProjectID_scrn, GET, rec.ID)
	return rec.ProjectID
}

// ----------------------------------------------------------------
// END   FeatureNew_ProjectID
// END   FeatureNew_ProjectID
// END   FeatureNew_ProjectID
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN FeatureNew_OriginCode
// BEGIN FeatureNew_OriginCode
// BEGIN FeatureNew_OriginCode
// ----------------------------------------------------------------
// FeatureNew_OriginCode_OnStore_impl provides the implementation for the callout
func FeatureNew_OriginCode_OnStore_impl(fieldval string, rec dm.FeatureNew, usr string) (string, error) {
	logs.Callout("FeatureNew", dm.FeatureNew_OriginCode_scrn, PUT, rec.ID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// FeatureNew_OriginCode_OnFetch_impl provides the implementation for the callout
func FeatureNew_OriginCode_OnFetch_impl(rec dm.FeatureNew) string {
	logs.Callout("FeatureNew", dm.FeatureNew_OriginCode_scrn, GET, rec.ID)
	return rec.OriginCode
}

// ----------------------------------------------------------------
// END   FeatureNew_OriginCode
// END   FeatureNew_OriginCode
// END   FeatureNew_OriginCode
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN FeatureNew_OriginID
// BEGIN FeatureNew_OriginID
// BEGIN FeatureNew_OriginID
// ----------------------------------------------------------------
// FeatureNew_OriginID_OnStore_impl provides the implementation for the callout
func FeatureNew_OriginID_OnStore_impl(fieldval string, rec dm.FeatureNew, usr string) (string, error) {
	logs.Callout("FeatureNew", dm.FeatureNew_OriginID_scrn, PUT, rec.ID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// FeatureNew_OriginID_OnFetch_impl provides the implementation for the callout
func FeatureNew_OriginID_OnFetch_impl(rec dm.FeatureNew) string {
	logs.Callout("FeatureNew", dm.FeatureNew_OriginID_scrn, GET, rec.ID)
	return rec.OriginID
}

// ----------------------------------------------------------------
// END   FeatureNew_OriginID
// END   FeatureNew_OriginID
// END   FeatureNew_OriginID
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN FeatureNew_ProjectName
// BEGIN FeatureNew_ProjectName
// BEGIN FeatureNew_ProjectName
// ----------------------------------------------------------------
// FeatureNew_ProjectName_OnStore_impl provides the implementation for the callout
func FeatureNew_ProjectName_OnStore_impl(fieldval string, rec dm.FeatureNew, usr string) (string, error) {
	logs.Callout("FeatureNew", dm.FeatureNew_ProjectName_scrn, PUT, rec.ID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// FeatureNew_ProjectName_OnFetch_impl provides the implementation for the callout
func FeatureNew_ProjectName_OnFetch_impl(rec dm.FeatureNew) string {
	logs.Callout("FeatureNew", dm.FeatureNew_ProjectName_scrn, GET, rec.ID)
	return rec.ProjectName
}

// ----------------------------------------------------------------
// END   FeatureNew_ProjectName
// END   FeatureNew_ProjectName
// END   FeatureNew_ProjectName
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// FeatureNew_ConfidenceCODE_OnStore_impl provides the implementation for the callout
func FeatureNew_ConfidenceCODE_OnStore_impl(fieldval string, rec dm.FeatureNew, usr string) (string, error) {
	logs.Callout("FeatureNew", dm.FeatureNew_ConfidenceCODE_scrn, PUT, rec.ID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// FeatureNew_ConfidenceCODE_OnFetch_impl provides the implementation for the callout
func FeatureNew_ConfidenceCODE_OnFetch_impl(rec dm.FeatureNew) string {
	logs.Callout("FeatureNew", dm.FeatureNew_ConfidenceCODE_scrn, GET, rec.ID)
	return rec.ConfidenceCODE
}
