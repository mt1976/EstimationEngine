package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/confidence.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Confidence (confidence)
// Endpoint 	        : Confidence (ConfidenceID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 27/11/2022 at 20:46:12
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	core "github.com/mt1976/ebEstimates/core"
	das "github.com/mt1976/ebEstimates/das"
	dm "github.com/mt1976/ebEstimates/datamodel"
)

// Confidence_GetByID() returns a single Confidence record
func Confidence_GetByCode(id string) (int, dm.Confidence, error) {

	tsql := das.SELECTALL + " " + das.FROM + " " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Confidence_SQLTable)
	tsql = tsql + " WHERE " + dm.Confidence_Code_sql + "='" + id + "'"
	_, _, confidenceItem, _ := confidence_Fetch(tsql)

	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return 1, confidenceItem, nil
}
