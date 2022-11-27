package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/inbox.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Inbox (inbox)
// Endpoint 	        : Inbox (Message)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/06/2022 at 17:00:15
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	core "github.com/mt1976/ebEstimates/core"

	dm "github.com/mt1976/ebEstimates/datamodel"
)

// Inbox_GetList() returns a list of all Inbox records
func Inbox_GetListByUser(inUser string) (int, []dm.Inbox, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.Inbox_SQLTable)
	tsql = tsql + " WHERE " + dm.Inbox_MailTo_sql + "='" + inUser + "'"

	count, inboxList, _, _ := inbox_Fetch(tsql)

	return count, inboxList, nil
}

// Inbox_GetList() returns a list of all Inbox records
func Inbox_GetUnreadByUser(inUser string) (int, []dm.Inbox, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.Inbox_SQLTable)
	tsql = tsql + " WHERE " + dm.Inbox_MailTo_sql + "='" + inUser + "'"
	tsql = tsql + " AND " + dm.Inbox_MailUnread_sql + "='" + "true" + "'"

	count, inboxList, _, _ := inbox_Fetch(tsql)

	return count, inboxList, nil
}
