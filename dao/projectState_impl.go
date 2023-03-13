package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/projectstate.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : ProjectState (projectstate)
// Endpoint 	        : ProjectState (ProjectStateID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 27/11/2022 at 20:46:15
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	core "github.com/mt1976/ebEstimates/core"
	das "github.com/mt1976/ebEstimates/das"

	dm "github.com/mt1976/ebEstimates/datamodel"
)

// ProjectState_GetByID() returns a single ProjectState record
func ProjectState_GetByCode(id string) (int, dm.ProjectState, error) {

	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.ProjectState_SQLTable)
	tsql = tsql + " WHERE " + dm.ProjectState_Code_sql + "='" + id + "'"
	_, _, projectstateItem, _ := projectstate_Fetch(tsql)

	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return 1, projectstateItem, nil
}

func ProjectState_GetMigrateable() (int, []dm.ProjectState, error) {
	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.ProjectState_SQLTable)
	tsql = tsql + " WHERE " + dm.ProjectState_Migrate_sql + das.EQ + das.VALUE(das.TRUE)
	noItems, projectstateItems, _, _ := projectstate_Fetch(tsql)

	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return noItems, projectstateItems, nil
}
