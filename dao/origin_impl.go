package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/origin.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Origin (origin)
// Endpoint 	        : Origin (OriginID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 27/11/2022 at 20:46:14
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"github.com/bojanz/currency"
	core "github.com/mt1976/ebEstimates/core"

	dm "github.com/mt1976/ebEstimates/datamodel"
)

// Origin_GetByCode() returns a single Origin record
func Origin_GetByCode(id string) (int, dm.Origin, error) {

	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Origin_SQLTable)
	tsql = tsql + " WHERE " + dm.Origin_Code_sql + "='" + id + "'"
	_, _, originItem, _ := origin_Fetch(tsql)

	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return 1, originItem, nil
}

// Origin_GetList() returns a list of all Origin records
func Origin_GetActiveList() (int, []dm.Origin, error) {

	tsql := core.DB_SELECT + " " + core.DB_ALL + " " + core.DB_FROM + " " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Origin_SQLTable)
	tsql = tsql + " " + core.DB_WHERE + " datalength(" + dm.Project_SYSDeleted_sql + ") = 0"
	tsql = tsql + " " + core.DB_SORTBY + " " + dm.Origin_Code_sql
	count, originList, _, _ := origin_Fetch(tsql)

	for i := 0; i < count; i++ {
		// START
		// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
		//
		symbol := "£"
		if originList[i].Currency != "" {
			locale := currency.NewLocale("en")
			symbol, _ = currency.GetSymbol(originList[i].Currency, locale)
		}
		originList[i].Currency = symbol
		//
		// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
		// END
	}

	return count, originList, nil
}
