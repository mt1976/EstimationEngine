package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/resource.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Resource (resource)
// Endpoint 	        : Resource (ResourceID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 03/01/2023 at 19:18:11
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	core "github.com/mt1976/ebEstimates/core"

	dm "github.com/mt1976/ebEstimates/datamodel"
)

// Resource_GetByCode() returns a single Resource record
func Resource_GetByCode(id string) (int, dm.Resource, error) {

	tsql := Resource_SQLbase
	tsql = tsql + " " + core.DB_WHERE + " " + dm.Resource_Code_sql + core.DB_EQ + "'" + id + "'"
	_, _, resourceItem, _ := resource_Fetch(tsql)

	// START
	// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
	return 1, resourceItem, nil
}
