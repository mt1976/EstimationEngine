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
	"github.com/mt1976/ebEstimates/logs"

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

// Project_GetList() returns a list of all Project records
func Project_GetList_ByCustomerAndName(originID string, name string) (int, dm.Project, error) {

	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Project_SQLTable)
	tsql = tsql + " WHERE " + dm.Project_OriginID_sql + "='" + originID + "'"
	tsql = tsql + " AND " + dm.Project_Name_sql + "='" + name + "'"

	count, projectList, _, _ := project_Fetch(tsql)
	if count == 0 {
		return count, dm.Project{}, nil
	}
	return count, projectList[0], nil
}

func Project_UpdateRate(project dm.Project, rate float64) error {

	//Get project status
	stateID := project.ProjectStateID
	_, state, err := ProjectState_GetByCode(stateID)
	if state.IsLocked != core.TRUE {

		newRate := core.FloatToString(rate)
		project.Notes = core.AddActivity(project.Notes, "Rate updated to "+newRate+" from "+project.DefaultRate)

		//Update the rate
		project.DefaultRate = newRate
		project.ProjectRate = newRate

		_, err := Project_StoreSystem(project)
		if err != nil {
			logs.Warning("Project_UpdateRate " + err.Error())
		}
		if state.Notify == core.TRUE {
			// TODO: Send email to Account Manager
			// TODO: Send email to Project Manager
		}
	}
	return err
}
