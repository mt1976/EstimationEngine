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
	logs "github.com/mt1976/ebEstimates/logs"

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

	tsql := das.SELECTALL + das.FROM + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Origin_SQLTable)
	tsql = tsql + " " + das.WHERE + das.ISNULL(dm.Project_SYSDeleted_sql)
	tsql = tsql + " " + das.SORTBY + dm.Origin_Code_sql
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

// ----------------------------------------------------------------
// BEGIN Origin_Code
// BEGIN Origin_Code
// BEGIN Origin_Code
// ----------------------------------------------------------------
// Origin_Code_OnStore_impl provides the implementation for the callout
func Origin_Code_OnStore_impl(fieldval string, rec dm.Origin, usr string) (string, error) {
	logs.Callout("Origin", dm.Origin_Code_scrn, PUT, rec.OriginID)
	Indexer_Put("Origin", dm.Origin_Code_scrn, rec.OriginID, fieldval)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Origin_Code_OnFetch_impl provides the implementation for the callout
func Origin_Code_OnFetch_impl(rec dm.Origin) string {
	logs.Callout("Origin", dm.Origin_Code_scrn, GET, rec.OriginID)

	return rec.Code
}

// ----------------------------------------------------------------
// Origin_Code_impl provides validation/actions for Code
func Origin_Code_impl(iAction string, iId string, iValue string, iRec dm.Origin, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Origin", dm.Origin_Code_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   Origin_Code
// END   Origin_Code
// END   Origin_Code
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Origin_FullName
// BEGIN Origin_FullName
// BEGIN Origin_FullName
// ----------------------------------------------------------------
// Origin_FullName_OnStore_impl provides the implementation for the callout
func Origin_FullName_OnStore_impl(fieldval string, rec dm.Origin, usr string) (string, error) {
	logs.Callout("Origin", dm.Origin_FullName_scrn, PUT, rec.OriginID)
	Indexer_Put("Origin", dm.Origin_FullName_scrn, rec.OriginID, fieldval)

	return fieldval, nil
}

// ----------------------------------------------------------------
// Origin_FullName_OnFetch_impl provides the implementation for the callout
func Origin_FullName_OnFetch_impl(rec dm.Origin) string {
	logs.Callout("Origin", dm.Origin_FullName_scrn, GET, rec.OriginID)
	return rec.FullName
}

// ----------------------------------------------------------------
// Origin_FullName_impl provides validation/actions for FullName
func Origin_FullName_impl(iAction string, iId string, iValue string, iRec dm.Origin, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Origin", dm.Origin_FullName_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   Origin_FullName
// END   Origin_FullName
// END   Origin_FullName
// ----------------------------------------------------------------
