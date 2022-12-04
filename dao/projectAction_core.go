package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/projectaction.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : ProjectAction (projectaction)
// Endpoint 	        : ProjectAction (ProjectID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 01/12/2022 at 09:40:02
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

// ProjectAction_GetList() returns a list of all ProjectAction records
func ProjectAction_GetList() (int, []dm.ProjectAction, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.ProjectAction_SQLTable)
	count, projectactionList, _, _ := projectaction_Fetch(tsql)
	
	return count, projectactionList, nil
}


// ProjectAction_GetLookup() returns a lookup list of all ProjectAction items in lookup format
func ProjectAction_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, projectactionList, _ := ProjectAction_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: projectactionList[i].ProjectID, Name: projectactionList[i].Name})
	}
	return returnList
}


// ProjectAction_GetByID() returns a single ProjectAction record
func ProjectAction_GetByID(id string) (int, dm.ProjectAction, error) {


	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.ProjectAction_SQLTable)
	tsql = tsql + " WHERE " + dm.ProjectAction_SQLSearchID + "='" + id + "'"
	_, _, projectactionItem, _ := projectaction_Fetch(tsql)

	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, projectactionItem, nil
}



// ProjectAction_DeleteByID() deletes a single ProjectAction record
func ProjectAction_Delete(id string) {




// Uses Hard Delete
	object_Table := core.GetSQLSchema(core.ApplicationPropertiesDB) + "." + dm.ProjectAction_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.ProjectAction_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
	
	

	
}














	


// ProjectAction_Store() saves/stores a ProjectAction record to the database
func ProjectAction_Store(r dm.ProjectAction,req *http.Request) error {

	err, r := ProjectAction_Validate(r)
	if err == nil {
		err = projectaction_Save(r, Audit_User(req))
	} else {
		logs.Information("ProjectAction_Store()", err.Error())
	}

	return err
}

// ProjectAction_StoreSystem() saves/stores a ProjectAction record to the database
func ProjectAction_StoreSystem(r dm.ProjectAction) error {
	
	err, r := ProjectAction_Validate(r)
	if err == nil {
		err = projectaction_Save(r, Audit_Host())
	} else {
		logs.Information("ProjectAction_Store()", err.Error())
	}

	return err
}

// ProjectAction_Validate() validates for saves/stores a ProjectAction record to the database
func ProjectAction_Validate(r dm.ProjectAction) (error,dm.ProjectAction) {
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

// projectaction_Save() saves/stores a ProjectAction record to the database
func projectaction_Save(r dm.ProjectAction,usr string) error {

    var err error



	

	if len(r.ProjectID) == 0 {
		r.ProjectID = ProjectAction_NewID(r)
	}

// If there are fields below, create the methods in dao\projectaction_impl.go





















	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	
logs.Storing("ProjectAction",fmt.Sprintf("%s", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.ProjectAction_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.ProjectAction_ProjectID_sql, r.ProjectID)
	ts = addData(ts, dm.ProjectAction_OriginID_sql, r.OriginID)
	ts = addData(ts, dm.ProjectAction_ProjectStateID_sql, r.ProjectStateID)
	ts = addData(ts, dm.ProjectAction_ProfileID_sql, r.ProfileID)
	ts = addData(ts, dm.ProjectAction_Name_sql, r.Name)
	ts = addData(ts, dm.ProjectAction_Description_sql, r.Description)
	ts = addData(ts, dm.ProjectAction_StartDate_sql, r.StartDate)
	ts = addData(ts, dm.ProjectAction_EndDate_sql, r.EndDate)
	ts = addData(ts, dm.ProjectAction_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.ProjectAction_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.ProjectAction_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.ProjectAction_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.ProjectAction_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.ProjectAction_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.ProjectAction_SYSDeleted_sql, r.SYSDeleted)
	ts = addData(ts, dm.ProjectAction_SYSDeletedBy_sql, r.SYSDeletedBy)
	ts = addData(ts, dm.ProjectAction_SYSDeletedHost_sql, r.SYSDeletedHost)
	ts = addData(ts, dm.ProjectAction_SYSActivity_sql, r.SYSActivity)
		
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := "INSERT INTO " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.ProjectAction_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	ProjectAction_Delete(r.ProjectID)
	das.Execute(tsql)



	return err

}



// projectaction_Fetch read all ProjectAction's
func projectaction_Fetch(tsql string) (int, []dm.ProjectAction, dm.ProjectAction, error) {

	var recItem dm.ProjectAction
	var recList []dm.ProjectAction

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.ProjectAction_SYSId_sql, "0")
	   recItem.ProjectID  = get_String(rec, dm.ProjectAction_ProjectID_sql, "")
	   recItem.OriginID  = get_String(rec, dm.ProjectAction_OriginID_sql, "")
	   recItem.ProjectStateID  = get_String(rec, dm.ProjectAction_ProjectStateID_sql, "")
	   recItem.ProfileID  = get_String(rec, dm.ProjectAction_ProfileID_sql, "")
	   recItem.Name  = get_String(rec, dm.ProjectAction_Name_sql, "")
	   recItem.Description  = get_String(rec, dm.ProjectAction_Description_sql, "")
	   recItem.StartDate  = get_String(rec, dm.ProjectAction_StartDate_sql, "")
	   recItem.EndDate  = get_String(rec, dm.ProjectAction_EndDate_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.ProjectAction_SYSCreated_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.ProjectAction_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.ProjectAction_SYSCreatedHost_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.ProjectAction_SYSUpdated_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.ProjectAction_SYSUpdatedBy_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.ProjectAction_SYSUpdatedHost_sql, "")
	   recItem.SYSDeleted  = get_String(rec, dm.ProjectAction_SYSDeleted_sql, "")
	   recItem.SYSDeletedBy  = get_String(rec, dm.ProjectAction_SYSDeletedBy_sql, "")
	   recItem.SYSDeletedHost  = get_String(rec, dm.ProjectAction_SYSDeletedHost_sql, "")
	   recItem.SYSActivity  = get_String(rec, dm.ProjectAction_SYSActivity_sql, "")
	
	// If there are fields below, create the methods in adaptor\ProjectAction_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
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
	


func ProjectAction_NewID(r dm.ProjectAction) string {
	
			id := uuid.New().String()
	
	return id
}



// projectaction_Fetch read all ProjectAction's
func ProjectAction_New() (int, []dm.ProjectAction, dm.ProjectAction, error) {

	var r = dm.ProjectAction{}
	var rList []dm.ProjectAction
	

	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}