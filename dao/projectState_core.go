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
// Date & Time		    : 01/12/2022 at 09:40:03
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (

	"fmt"
	"net/http"
core "github.com/mt1976/ebEstimates/core"
"github.com/google/uuid"
das  "github.com/mt1976/ebEstimates/das"
	
	
	
	dm   "github.com/mt1976/ebEstimates/datamodel"
	logs   "github.com/mt1976/ebEstimates/logs"
)

// ProjectState_GetList() returns a list of all ProjectState records
func ProjectState_GetList() (int, []dm.ProjectState, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.ProjectState_SQLTable)
	count, projectstateList, _, _ := projectstate_Fetch(tsql)
	
	return count, projectstateList, nil
}


// ProjectState_GetLookup() returns a lookup list of all ProjectState items in lookup format
func ProjectState_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, projectstateList, _ := ProjectState_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: projectstateList[i].Code, Name: projectstateList[i].Name})
	}
	return returnList
}


// ProjectState_GetByID() returns a single ProjectState record
func ProjectState_GetByID(id string) (int, dm.ProjectState, error) {


	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.ProjectState_SQLTable)
	tsql = tsql + " WHERE " + dm.ProjectState_SQLSearchID + "='" + id + "'"
	_, _, projectstateItem, _ := projectstate_Fetch(tsql)

	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, projectstateItem, nil
}



// ProjectState_DeleteByID() deletes a single ProjectState record
func ProjectState_Delete(id string) {




// Uses Hard Delete
	object_Table := core.GetSQLSchema(core.ApplicationPropertiesDB) + "." + dm.ProjectState_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.ProjectState_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
	
	

	
}














	
	// ProjectState_SoftDeleteByID() soft deletes a single ProjectState record
func ProjectState_SoftDelete(id string) {
	//Uses Soft Delete
		_,projectstateItem,_ := ProjectState_GetByID(id)
		projectstateItem.SYSDeletedBy = Audit_Update("", Audit_Host())
		projectstateItem.SYSDeleted = Audit_Update("", Audit_TimeStamp())
		projectstateItem.SYSDeletedHost = Audit_Update("", Audit_Host())
		ProjectState_StoreSystem(projectstateItem)
}
	


// ProjectState_Store() saves/stores a ProjectState record to the database
func ProjectState_Store(r dm.ProjectState,req *http.Request) error {

	err, r := ProjectState_Validate(r)
	if err == nil {
		err = projectstate_Save(r, Audit_User(req))
	} else {
		logs.Information("ProjectState_Store()", err.Error())
	}

	return err
}

// ProjectState_StoreSystem() saves/stores a ProjectState record to the database
func ProjectState_StoreSystem(r dm.ProjectState) error {
	
	err, r := ProjectState_Validate(r)
	if err == nil {
		err = projectstate_Save(r, Audit_Host())
	} else {
		logs.Information("ProjectState_Store()", err.Error())
	}

	return err
}

// ProjectState_Validate() validates for saves/stores a ProjectState record to the database
func ProjectState_Validate(r dm.ProjectState) (error,dm.ProjectState) {
	var err error
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	

	return err,r
}
//

// projectstate_Save() saves/stores a ProjectState record to the database
func projectstate_Save(r dm.ProjectState,usr string) error {

    var err error



	

	if len(r.ProjectStateID) == 0 {
		r.ProjectStateID = ProjectState_NewID(r)
	}

// If there are fields below, create the methods in dao\projectstate_impl.go















	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	
logs.Storing("ProjectState",fmt.Sprintf("%s", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.ProjectState_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.ProjectState_ProjectStateID_sql, r.ProjectStateID)
	ts = addData(ts, dm.ProjectState_Code_sql, r.Code)
	ts = addData(ts, dm.ProjectState_Name_sql, r.Name)
	ts = addData(ts, dm.ProjectState_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.ProjectState_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.ProjectState_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.ProjectState_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.ProjectState_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.ProjectState_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.ProjectState_SYSDeleted_sql, r.SYSDeleted)
	ts = addData(ts, dm.ProjectState_SYSDeletedBy_sql, r.SYSDeletedBy)
	ts = addData(ts, dm.ProjectState_SYSDeletedHost_sql, r.SYSDeletedHost)
		
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := "INSERT INTO " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.ProjectState_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	ProjectState_Delete(r.ProjectStateID)
	das.Execute(tsql)



	return err

}



// projectstate_Fetch read all ProjectState's
func projectstate_Fetch(tsql string) (int, []dm.ProjectState, dm.ProjectState, error) {

	var recItem dm.ProjectState
	var recList []dm.ProjectState

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.ProjectState_SYSId_sql, "0")
	   recItem.ProjectStateID  = get_String(rec, dm.ProjectState_ProjectStateID_sql, "")
	   recItem.Code  = get_String(rec, dm.ProjectState_Code_sql, "")
	   recItem.Name  = get_String(rec, dm.ProjectState_Name_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.ProjectState_SYSCreated_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.ProjectState_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.ProjectState_SYSCreatedHost_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.ProjectState_SYSUpdated_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.ProjectState_SYSUpdatedBy_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.ProjectState_SYSUpdatedHost_sql, "")
	   recItem.SYSDeleted  = get_String(rec, dm.ProjectState_SYSDeleted_sql, "")
	   recItem.SYSDeletedBy  = get_String(rec, dm.ProjectState_SYSDeletedBy_sql, "")
	   recItem.SYSDeletedHost  = get_String(rec, dm.ProjectState_SYSDeletedHost_sql, "")
	
	// If there are fields below, create the methods in adaptor\ProjectState_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func ProjectState_NewID(r dm.ProjectState) string {
	
			id := uuid.New().String()
	
	return id
}



// projectstate_Fetch read all ProjectState's
func ProjectState_New() (int, []dm.ProjectState, dm.ProjectState, error) {

	var r = dm.ProjectState{}
	var rList []dm.ProjectState
	

	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}