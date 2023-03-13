package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/originstate.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : OriginState (originstate)
// Endpoint 	        : OriginState (OriginStateID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 27/11/2022 at 20:46:14
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	core "github.com/mt1976/ebEstimates/core"
	das "github.com/mt1976/ebEstimates/das"

	dm "github.com/mt1976/ebEstimates/datamodel"
)

// OriginState_GetByID() returns a single OriginState record
func OriginState_GetByCode(id string) (int, dm.OriginState, error) {

	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.OriginState_SQLTable)
	tsql = tsql + " WHERE " + dm.OriginState_Code_sql + "='" + id + "'"
	_, _, originstateItem, _ := originstate_Fetch(tsql)

	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return 1, originstateItem, nil
}

func OriginState_GetMigrateable() (int, []dm.OriginState, error) {
	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.OriginState_SQLTable)
	tsql = tsql + " WHERE " + dm.OriginState_Migrate_sql + das.EQ + das.VALUE(das.TRUE)
	noItems, originstateItems, _, _ := originstate_Fetch(tsql)

	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return noItems, originstateItems, nil
}
