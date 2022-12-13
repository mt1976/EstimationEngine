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
// Date & Time		    : 11/12/2022 at 19:24:01
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (

	"fmt"
	"net/http"
	core "github.com/mt1976/ebEstimates/core"
	"github.com/google/uuid"
	das  "github.com/mt1976/ebEstimates/das"
	
		 
		// Does Lookup
		adaptor   "github.com/mt1976/ebEstimates/adaptor"
	
	dm   "github.com/mt1976/ebEstimates/datamodel"
	logs   "github.com/mt1976/ebEstimates/logs"
)

var Project_SQLbase string
var Project_QualifiedName string
func init(){
	Project_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.Project_SQLTable)
	Project_SQLbase =  core.DB_SELECT + " "+ core.DB_ALL + " " + core.DB_FROM + " " + Project_QualifiedName
}

// Project_GetList() returns a list of all Project records
func Project_GetList() (int, []dm.Project, error) {
	
	tsql := Project_SQLbase
	count, projectList, _, _ := project_Fetch(tsql)
	
	return count, projectList, nil
}


// Project_GetLookup() returns a lookup list of all Project items in lookup format
func Project_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, projectList, _ := Project_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: projectList[i].ProjectID, Name: projectList[i].Name})
	}
	return returnList
}


// Project_GetByID() returns a single Project record
func Project_GetByID(id string) (int, dm.Project, error) {


	tsql := Project_SQLbase
	tsql = tsql + " " + core.DB_WHERE + " " + dm.Project_SQLSearchID + core.DB_EQ + "'" + id + "'"
	_, _, projectItem, _ := project_Fetch(tsql)

	// START
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	projectItem.NoEstimationSessions,projectItem.NoEstimationSessions_props = adaptor.Project_NoEstimationSessions_impl (adaptor.GET,id,projectItem.NoEstimationSessions,projectItem,projectItem.NoEstimationSessions_props)
	// 
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, projectItem, nil
}



// Project_DeleteByID() deletes a single Project record
func Project_Delete(id string) {




// Uses Hard Delete
	object_Table := Project_QualifiedName
	tsql := core.DB_DELETE+" "+core.DB_FROM+" " + object_Table
	tsql = tsql + " " + core.DB_WHERE + " " + dm.Project_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
	
	

	
}














	
	// Project_SoftDeleteByID() soft deletes a single Project record
func Project_SoftDelete(id string) {
	//Uses Soft Delete
		_,projectItem,_ := Project_GetByID(id)
		projectItem.SYSDeletedBy = Audit_Update("", Audit_Host())
		projectItem.SYSDeleted = Audit_Update("", Audit_TimeStamp())
		projectItem.SYSDeletedHost = Audit_Update("", Audit_Host())
		Project_StoreSystem(projectItem)
}
	


// Project_Store() saves/stores a Project record to the database
func Project_Store(r dm.Project,req *http.Request) error {

	r, err := Project_Validate(r)
	if err == nil {
		err = project_Save(r, Audit_User(req))
	} else {
		logs.Information("Project_Store()", err.Error())
	}

	return err
}

// Project_StoreSystem() saves/stores a Project record to the database
func Project_StoreSystem(r dm.Project) error {
	
	r, err := Project_Validate(r)
	if err == nil {
		err = project_Save(r, Audit_Host())
	} else {
		logs.Information("Project_Store()", err.Error())
	}

	return err
}

// Project_Validate() validates for saves/stores a Project record to the database
func Project_Validate(r dm.Project) (dm.Project, error) {
	var err error
	// START
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	r.NoEstimationSessions,r.NoEstimationSessions_props = adaptor.Project_NoEstimationSessions_impl (adaptor.PUT,r.ProjectID,r.NoEstimationSessions,r,r.NoEstimationSessions_props)
	// 
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	

	return r,err
}
//

// project_Save() saves/stores a Project record to the database
func project_Save(r dm.Project,usr string) error {

    var err error



	

	if len(r.ProjectID) == 0 {
		r.ProjectID = Project_NewID(r)
	}

// If there are fields below, create the methods in dao\project_impl.go





















  r.NoEstimationSessions,err = adaptor.Project_NoEstimationSessions_OnStore_impl (r.NoEstimationSessions,r,usr)


	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	r.SYSDbVersion = core.DB_Version()
	
logs.Storing("Project",fmt.Sprintf("%v", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.Project_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.Project_ProjectID_sql, r.ProjectID)
	ts = addData(ts, dm.Project_OriginID_sql, r.OriginID)
	ts = addData(ts, dm.Project_ProjectStateID_sql, r.ProjectStateID)
	ts = addData(ts, dm.Project_ProfileID_sql, r.ProfileID)
	ts = addData(ts, dm.Project_Name_sql, r.Name)
	ts = addData(ts, dm.Project_Description_sql, r.Description)
	ts = addData(ts, dm.Project_StartDate_sql, r.StartDate)
	ts = addData(ts, dm.Project_EndDate_sql, r.EndDate)
	ts = addData(ts, dm.Project_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.Project_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.Project_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.Project_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.Project_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.Project_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.Project_SYSDeleted_sql, r.SYSDeleted)
	ts = addData(ts, dm.Project_SYSDeletedBy_sql, r.SYSDeletedBy)
	ts = addData(ts, dm.Project_SYSDeletedHost_sql, r.SYSDeletedHost)
	ts = addData(ts, dm.Project_SYSActivity_sql, r.SYSActivity)
	ts = addData(ts, dm.Project_SYSDbVersion_sql, r.SYSDbVersion)
	ts = addData(ts, dm.Project_Comments_sql, r.Comments)
	
		
	// 
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := core.DB_INSERT + " " + core.DB_INTO + " " + Project_QualifiedName
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " "+core.DB_VALUES +" (" + values(ts) + ")"

	Project_Delete(r.ProjectID)
	das.Execute(tsql)



	return err

}



// project_Fetch read all Project's
func project_Fetch(tsql string) (int, []dm.Project, dm.Project, error) {

	var recItem dm.Project
	var recList []dm.Project

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.Project_SYSId_sql, "0")
	   recItem.ProjectID  = get_String(rec, dm.Project_ProjectID_sql, "")
	   recItem.OriginID  = get_String(rec, dm.Project_OriginID_sql, "")
	   recItem.ProjectStateID  = get_String(rec, dm.Project_ProjectStateID_sql, "")
	   recItem.ProfileID  = get_String(rec, dm.Project_ProfileID_sql, "")
	   recItem.Name  = get_String(rec, dm.Project_Name_sql, "")
	   recItem.Description  = get_String(rec, dm.Project_Description_sql, "")
	   recItem.StartDate  = get_String(rec, dm.Project_StartDate_sql, "")
	   recItem.EndDate  = get_String(rec, dm.Project_EndDate_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.Project_SYSCreated_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.Project_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.Project_SYSCreatedHost_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.Project_SYSUpdated_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.Project_SYSUpdatedBy_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.Project_SYSUpdatedHost_sql, "")
	   recItem.SYSDeleted  = get_String(rec, dm.Project_SYSDeleted_sql, "")
	   recItem.SYSDeletedBy  = get_String(rec, dm.Project_SYSDeletedBy_sql, "")
	   recItem.SYSDeletedHost  = get_String(rec, dm.Project_SYSDeletedHost_sql, "")
	   recItem.SYSActivity  = get_String(rec, dm.Project_SYSActivity_sql, "")
	   recItem.SYSDbVersion  = get_String(rec, dm.Project_SYSDbVersion_sql, "")
	   recItem.Comments  = get_String(rec, dm.Project_Comments_sql, "")
	
	
	// If there are fields below, create the methods in adaptor\Project_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	   recItem.NoEstimationSessions  = adaptor.Project_NoEstimationSessions_OnFetch_impl (recItem)
	
	// 
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func Project_NewID(r dm.Project) string {
	
			id := uuid.New().String()
	
	return id
}



// project_Fetch read all Project's
func Project_New() (int, []dm.Project, dm.Project, error) {

	var r = dm.Project{}
	var rList []dm.Project
	

	// START
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	r.NoEstimationSessions,r.NoEstimationSessions_props = adaptor.Project_NoEstimationSessions_impl (adaptor.NEW,r.ProjectID,r.NoEstimationSessions,r,r.NoEstimationSessions_props)
	// 
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}