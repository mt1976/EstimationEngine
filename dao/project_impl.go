package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/project.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Project (project)
// Endpoint 	        : Project (ProjectID)
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

// Project_GetList() returns a list of all Project records
func Project_ByOrigin_Active_GetList(id string) (int, []dm.Project, error) {

	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Project_SQLTable)
	tsql = tsql + " WHERE " + dm.Project_OriginID_sql + "='" + id + "'"
	tsql = tsql + " AND datalength(" + dm.Project_SYSDeleted_sql + ") = 0"

	count, projectList, _, _ := project_Fetch(tsql)

	return count, projectList, nil
}
