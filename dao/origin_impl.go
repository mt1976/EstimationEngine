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
	"github.com/mt1976/ebEstimates/das"

	dm "github.com/mt1976/ebEstimates/datamodel"
)

// Origin_GetByCode() returns a single Origin record
func Origin_GetByCode(id string) (int, dm.Origin, error) {
	//logs.Warning("Origin_GetByCode " + id + " " + core.GetSQLSchema(core.ApplicationPropertiesDB))
	tsql := das.SELECTALL + das.FROM + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Origin_SQLTable)
	tsql = tsql + " WHERE " + dm.Origin_Code_sql + "='" + id + "'"
	//logs.Warning("Origin GetByCode " + tsql)
	no, _, originItem, err := origin_Fetch(tsql)
	//logs.Warning("Query Done")
	if no == 0 {
		return no, dm.Origin{}, nil
	}
	if no > 1 {
		return no, dm.Origin{}, nil
	}
	//logs.Warning("Origin_GetByCode " + id + " " + originItem.FullName)
	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return no, originItem, err
}

// Origin_GetList() returns a list of all Origin records
func Origin_GetActiveList() (int, []dm.Origin, error) {

	tsql := das.SELECTALL + das.FROM + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Origin_SQLTable)
	tsql = tsql + " " + das.WHERE + das.ISNULL(dm.Project_SYSDeleted_sql)
	tsql = tsql + " " + das.SORTBY + dm.Origin_Code_sql
	count, originList, _, _ := origin_Fetch(tsql)

	for i := 0; i < count; i++ {
		// START
		// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
		//
		symbol := "£"
		ok := false
		if originList[i].Currency != "" {
			locale := currency.NewLocale("en")
			symbol, ok = currency.GetSymbol(originList[i].Currency, locale)
			if !ok {
				symbol = originList[i].Currency
			}
		}
		originList[i].Currency = symbol
		//
		// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
		// END
	}

	return count, originList, nil
}
