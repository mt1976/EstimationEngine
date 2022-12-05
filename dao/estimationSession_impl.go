package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/estimationsession.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : EstimationSession (estimationsession)
// Endpoint 	        : EstimationSession (EstimationSessionID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 27/11/2022 at 20:46:13
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"errors"

	core "github.com/mt1976/ebEstimates/core"

	dm "github.com/mt1976/ebEstimates/datamodel"
)

// EstimationSession_GetList() returns a list of all EstimationSession records
func EstimationSession_ByProject_GetList() (int, []dm.EstimationSession, error) {

	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.EstimationSession_SQLTable)
	count, estimationsessionList, _, _ := estimationsession_Fetch(tsql)

	return count, estimationsessionList, nil
}

// EstimationSession_GetList() returns a list of all EstimationSession records
func EstimationSession_Active_ByProject_GetList(id string) (int, []dm.EstimationSession, error) {

	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.EstimationSession_SQLTable)
	tsql = tsql + " WHERE " + dm.EstimationSession_ProjectID_sql + "='" + id + "'"
	tsql = tsql + " AND datalength(" + dm.EstimationSession_SYSDeleted_sql + ") = 0"
	count, estimationsessionList, _, _ := estimationsession_Fetch(tsql)

	return count, estimationsessionList, nil
}

// EstimationSession_GetList() returns a list of all EstimationSession records
func EstimationSession_ByADO_GetList(id string) (int, []dm.EstimationSession, error) {

	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.EstimationSession_SQLTable)
	tsql = tsql + " WHERE " + dm.EstimationSession_AdoID_sql + "='" + id + "'"
	count, estimationsessionList, _, _ := estimationsession_Fetch(tsql)
	if count > 1 {
		return count, estimationsessionList, errors.New("duplicate records found")
	}
	return count, estimationsessionList, nil
}

// EstimationSession_GetList() returns a list of all EstimationSession records
func EstimationSession_ByFreshDesk_GetList(id string) (int, []dm.EstimationSession, error) {

	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.EstimationSession_SQLTable)
	tsql = tsql + " WHERE " + dm.EstimationSession_FreshdeskID_sql + "='" + id + "'"
	count, estimationsessionList, _, _ := estimationsession_Fetch(tsql)
	if count > 1 {
		return count, estimationsessionList, errors.New("duplicate records found")
	}
	return count, estimationsessionList, nil
}
