package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/doctype.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : DocType (doctype)
// Endpoint 	        : DocType (DocTypeID)
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

// DocType_GetByID() returns a single DocType record
func DocType_GetByCode(id string) (int, dm.DocType, error) {

	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.DocType_SQLTable)
	tsql = tsql + " WHERE " + dm.DocType_Code_sql + "='" + id + "'"
	_, _, doctypeItem, _ := doctype_Fetch(tsql)

	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return 1, doctypeItem, nil
}
