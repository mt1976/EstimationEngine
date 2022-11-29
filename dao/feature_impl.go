package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/Feature.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Feature (Feature)
// Endpoint 	        : Feature (FeatureID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 27/11/2022 at 20:46:13
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	core "github.com/mt1976/ebEstimates/core"

	dm "github.com/mt1976/ebEstimates/datamodel"
)

// Feature_GetList() returns a list of all Feature records
func Feature_ByEstimationSession_GetList() (int, []dm.Feature, error) {

	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Feature_SQLTable)
	count, FeatureList, _, _ := feature_Fetch(tsql)

	return count, FeatureList, nil
}

// Feature_GetList() returns a list of all Feature records
func Feature_Active_ByEstimationSession_GetList(id string) (int, []dm.Feature, error) {

	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Feature_SQLTable)
	tsql = tsql + " WHERE " + dm.Feature_EstimationSessionID_sql + "='" + id + "'"
	tsql = tsql + " AND datalength(" + dm.Feature_SYSDeleted_sql + ") = 0"
	count, FeatureList, _, _ := feature_Fetch(tsql)

	return count, FeatureList, nil
}
