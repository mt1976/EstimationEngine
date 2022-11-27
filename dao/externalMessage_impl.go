package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/externalmessage.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : ExternalMessage (externalmessage)
// Endpoint 	        : ExternalMessage (MessageID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/12/2021 at 09:31:50
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	core "github.com/mt1976/ebEstimates/core"

	dm "github.com/mt1976/ebEstimates/datamodel"
)

// ExternalMessage_GetByID() returns a single ExternalMessage record
func ExternalMessage_GetActive() (int, []dm.ExternalMessage, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.ExternalMessage_SQLTable)
	tsql = tsql + " WHERE " + dm.ExternalMessage_ResponseDate_sql + "= ''"
	tsql = tsql + " AND " + dm.ExternalMessage_MessageACKNAK_sql + "= ''"
	noItems, externalmessageList, _, _ := externalmessage_Fetch(tsql)

	return noItems, externalmessageList, nil
}
