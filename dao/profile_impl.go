package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/profile.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Profile (profile)
// Endpoint 	        : Profile (ProfileID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 27/11/2022 at 20:46:14
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	core "github.com/mt1976/ebEstimates/core"

	dm "github.com/mt1976/ebEstimates/datamodel"
)

// Profile_GetByID() returns a single Profile record
func Profile_GetByCode(id string) (int, dm.Profile, error) {

	tsql := core.DB_SELECT + " " + core.DB_ALL + " " + core.DB_FROM + " " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Profile_SQLTable)
	tsql = tsql + " " + core.DB_WHERE + " " + dm.Profile_Code_sql + "='" + id + "'"
	_, _, profileItem, _ := profile_Fetch(tsql)

	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return 1, profileItem, nil
}
