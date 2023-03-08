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

// Feature_TrackerID_OnStore_impl provides the implementation for the callout
func Feature_TrackerID_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_TrackerID_scrn, PUT, rec.FeatureID)
	Indexer_Put("Feature", dm.Feature_TrackerID_scrn, rec.FeatureID, fieldval)
	return fieldval, nil
}

// Feature_AdoID_OnStore_impl provides the implementation for the callout
func Feature_AdoID_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	Indexer_Put("Feature", dm.Feature_AdoID_scrn, rec.FeatureID, fieldval)
	logs.Callout("Feature", dm.Feature_AdoID_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// Feature_FreshdeskID_OnStore_impl provides the implementation for the callout
func Feature_FreshdeskID_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_FreshdeskID_scrn, PUT, rec.FeatureID)
	Indexer_Put("Feature", dm.Feature_FreshdeskID_scrn, rec.FeatureID, fieldval)

	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_ExtRef_OnStore_impl provides the implementation for the callout
func Feature_ExtRef_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_ExtRef_scrn, PUT, rec.FeatureID)
	Indexer_Put("Feature", dm.Feature_ExtRef_scrn, rec.FeatureID, fieldval)

	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_ExtRef2_OnStore_impl provides the implementation for the callout
func Feature_ExtRef2_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_ExtRef2_scrn, PUT, rec.FeatureID)
	Indexer_Put("Feature", dm.Feature_ExtRef2_scrn, rec.FeatureID, fieldval)

	return fieldval, nil
}

// ----------------------------------------------------------------
// BEGIN Feature_OriginID
// BEGIN Feature_OriginID
// BEGIN Feature_OriginID
// ----------------------------------------------------------------
// Feature_OriginID_OnStore_impl provides the implementation for the callout
func Feature_OriginID_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_OriginID_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_OriginID_OnFetch_impl provides the implementation for the callout
func Feature_OriginID_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_OriginID_scrn, GET, rec.FeatureID)
	esID := rec.EstimationSessionID
	es := CACHE_Read(dm.EstimationSession_Name, esID).(dm.EstimationSession)
	projID := es.ProjectID
	proj := CACHE_Read(dm.Project_Name, projID).(dm.Project)
	oID := proj.OriginID
	origin := CACHE_Read(dm.Origin_Name, oID).(dm.Origin)
	return origin.OriginID
}

func Feature_ProjectID_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_ProjectID_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_ProjectID_OnFetch_impl provides the implementation for the callout
func Feature_ProjectID_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_ProjectID_scrn, GET, rec.FeatureID)
	esID := rec.EstimationSessionID
	es := CACHE_Read(dm.EstimationSession_Name, esID).(dm.EstimationSession)
	return es.ProjectID
}

// Feature_OriginCode_OnStore_impl provides the implementation for the callout
func Feature_OriginCode_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_OriginCode_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_OriginCode_OnFetch_impl provides the implementation for the callout
func Feature_OriginCode_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_OriginCode_scrn, GET, rec.FeatureID)
	esID := rec.EstimationSessionID
	es := CACHE_Read(dm.EstimationSession_Name, esID).(dm.EstimationSession)
	projID := es.ProjectID
	proj := CACHE_Read(dm.Project_Name, projID).(dm.Project)
	return proj.OriginID
}

// Feature_EstimationSessionName_OnStore_impl provides the implementation for the callout
func Feature_EstimationSessionName_OnStore_impl(fieldval string, rec dm.Feature, usr string) (string, error) {
	logs.Callout("Feature", dm.Feature_EstimationSessionName_scrn, PUT, rec.FeatureID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Feature_EstimationSessionName_OnFetch_impl provides the implementation for the callout
func Feature_EstimationSessionName_OnFetch_impl(rec dm.Feature) string {
	logs.Callout("Feature", dm.Feature_EstimationSessionName_scrn, GET, rec.FeatureID)
	esID := rec.EstimationSessionID
	es := CACHE_Read(dm.EstimationSession_Name, esID).(dm.EstimationSession)
	return es.Name
}
