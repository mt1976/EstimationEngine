package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/project.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Project (project)
// Endpoint 	        : Project (ProjectID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 04/03/2023 at 20:14:13
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
	"github.com/pkg/errors"
)

var Project_SQLbase string
var Project_QualifiedName string
func init(){
	Project_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.Project_SQLTable)
	Project_SQLbase =  das.SELECTALL + das.FROM + Project_QualifiedName
}

// Project_GetList() returns a list of all Project records
func Project_GetList() (int, []dm.Project, error) {
	count, projectList, err := Project_GetListFiltered("")
	return count, projectList, err
}

// Project_GetListFiltered() returns a filtered list of all Project records
func Project_GetListFiltered(filter string) (int, []dm.Project, error) {
	
	tsql := Project_SQLbase
	if filter != "" {
		tsql = tsql + " " + das.WHERE + filter
	}
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

// Project_GetFilteredLookup() returns a lookup list of all Project items in lookup format
func Project_GetFilteredLookup(requestObject string,requestField string) []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	objectName := Translate("ObjectName", requestObject)
	reqField := requestField+"_Project_Filter"
	filter,_ := Data_GetString(objectName, reqField, dm.Data_Category_FilterRule)
	if filter == "" {
		logs.Warning("Project_GetFilteredLookup() - No filter found : " + reqField + " for Object: " + objectName)
	} 
	count, projectList, _ := Project_GetListFiltered(filter)
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: projectList[i].ProjectID, Name: projectList[i].Name})
	}
	return returnList
}



// Project_GetByID() returns a single Project record
func Project_GetByID(id string) (int, dm.Project, error) {


	tsql := Project_SQLbase
	tsql = tsql + " " + das.WHERE + dm.Project_SQLSearchID + das.EQ + das.ID(id)
	_, _, projectItem, _ := project_Fetch(tsql)


	projectItem = Project_PostGet(projectItem,id)

	return 1, projectItem, nil
}

func Project_PostGet(projectItem dm.Project,id string) dm.Project {
	// START
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	projectItem.ProjectID,projectItem.ProjectID_props = Project_ProjectID_validate_impl (GET,id,projectItem.ProjectID,projectItem,projectItem.ProjectID_props)
	projectItem.Description,projectItem.Description_props = Project_Description_validate_impl (GET,id,projectItem.Description,projectItem,projectItem.Description_props)
	projectItem.NoEstimationSessions,projectItem.NoEstimationSessions_props = Project_NoEstimationSessions_validate_impl (GET,id,projectItem.NoEstimationSessions,projectItem,projectItem.NoEstimationSessions_props)
	// 
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return projectItem
}




// Project_DeleteByID() deletes a single Project record
func Project_Delete(id string) {
		Project_SoftDelete(id)
	}
// Project_SoftDelete(id string) soft deletes a single Project record
func Project_SoftDelete(id string) {
	//Uses Soft Delete
		_,projectItem,_ := Project_GetByID(id)
		projectItem.SYSDeletedBy = Audit_Update("", Audit_Host())
		projectItem.SYSDeleted = Audit_Update("", Audit_TimeStamp())
		projectItem.SYSDeletedHost = Audit_Update("", Audit_Host())
		_,err := Project_StoreSystem(projectItem)
		if err != nil {
			logs.Error("Project_SoftDelete()",err)
		}
}
	

// Project_HardDelete(id string) soft deletes a single Project record
func Project_HardDelete(id string) {
			// Uses Hard Delete
		object_Table := Project_QualifiedName
		tsql := das.DELETE + das.FROM + object_Table
		tsql = tsql + " " + das.WHERE + dm.Project_SQLSearchID + das.EQ + das.ID(id)
		das.Execute(tsql)
		//if err != nil {
		//	logs.Error("Project_SoftDelete()",err)
		//}
}


// Project_Store() saves/stores a Project record to the database
func Project_Store(r dm.Project,req *http.Request) (dm.Project,error) {

	r, err := Project_Validate(r)
	if err == nil {
		err = project_Save(r, Audit_User(req))
	} else {
		logs.Information("Project_Store()", err.Error())
	}

	return r, err
}

// Project_StoreSystem() saves/stores a Project record to the database
func Project_StoreSystem(r dm.Project) (dm.Project,error) {
	
	r, err := Project_Validate(r)
	if err == nil {
		err = project_Save(r, Audit_Host())
	} else {
		logs.Information("Project_Store()", err.Error())
	}

	return r, err
}

// Project_StoreProcess() saves/stores a Project record to the database
func Project_StoreProcess(r dm.Project, operator string) (dm.Project,error) {
	
	r, err := Project_Validate(r)
	if err == nil {
		err = project_Save(r, operator)
	} else {
		logs.Information("Project_StoreProcess()", err.Error())
	}

	return r, err
}

// Project_Validate() validates for saves/stores a Project record to the database
func Project_Validate(r dm.Project) (dm.Project, error) {
	var err error
	// START
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	r.ProjectID,r.ProjectID_props = Project_ProjectID_validate_impl (PUT,r.ProjectID,r.ProjectID,r,r.ProjectID_props)
	if r.ProjectID_props.MsgMessage != "" {
		err = errors.New(r.ProjectID_props.MsgMessage)
	}
	r.Description,r.Description_props = Project_Description_validate_impl (PUT,r.ProjectID,r.Description,r,r.Description_props)
	if r.Description_props.MsgMessage != "" {
		err = errors.New(r.Description_props.MsgMessage)
	}
	r.NoEstimationSessions,r.NoEstimationSessions_props = Project_NoEstimationSessions_validate_impl (PUT,r.ProjectID,r.NoEstimationSessions,r,r.NoEstimationSessions_props)
	if r.NoEstimationSessions_props.MsgMessage != "" {
		err = errors.New(r.NoEstimationSessions_props.MsgMessage)
	}
	// 
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
    r.ProjectID,err = Project_ProjectID_OnStore_impl (r.ProjectID,r,usr)
    r.Description,err = Project_Description_OnStore_impl (r.Description,r,usr)
    r.NoEstimationSessions,err = Project_NoEstimationSessions_OnStore_impl (r.NoEstimationSessions,r,usr)

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
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	ts = addData(ts, dm.Project_ProjectRate_sql, r.ProjectRate)
	ts = addData(ts, dm.Project_DefaultRate_sql, r.DefaultRate)
	ts = addData(ts, dm.Project_ProjectAnalyst_sql, r.ProjectAnalyst)
	ts = addData(ts, dm.Project_ProjectEngineer_sql, r.ProjectEngineer)
	ts = addData(ts, dm.Project_ProjectManager_sql, r.ProjectManager)
	ts = addData(ts, dm.Project_Releases_sql, r.Releases)
	ts = addData(ts, dm.Project_Notes_sql, r.Notes)
	
		
	// 
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := das.INSERT + das.INTO + Project_QualifiedName
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " "+das.VALUES +"(" + values(ts) + ")"

	Project_HardDelete(r.ProjectID)
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
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	   recItem.ProjectRate  = get_String(rec, dm.Project_ProjectRate_sql, "")
	   recItem.DefaultRate  = get_String(rec, dm.Project_DefaultRate_sql, "")
	   recItem.ProjectAnalyst  = get_String(rec, dm.Project_ProjectAnalyst_sql, "")
	   recItem.ProjectEngineer  = get_String(rec, dm.Project_ProjectEngineer_sql, "")
	   recItem.ProjectManager  = get_String(rec, dm.Project_ProjectManager_sql, "")
	   recItem.Releases  = get_String(rec, dm.Project_Releases_sql, "")
	   recItem.Notes  = get_String(rec, dm.Project_Notes_sql, "")
	
	
	// If there are fields below, create the methods in dao\Project_adaptor.go
	   recItem.ProjectID  = Project_ProjectID_OnFetch_impl (recItem)
	   recItem.Description  = Project_Description_OnFetch_impl (recItem)
	   recItem.NoEstimationSessions  = Project_NoEstimationSessions_OnFetch_impl (recItem)
	// 
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	r.ProjectID,r.ProjectID_props = Project_ProjectID_validate_impl (NEW,r.ProjectID,r.ProjectID,r,r.ProjectID_props)
	r.Description,r.Description_props = Project_Description_validate_impl (NEW,r.ProjectID,r.Description,r,r.Description_props)
	r.NoEstimationSessions,r.NoEstimationSessions_props = Project_NoEstimationSessions_validate_impl (NEW,r.ProjectID,r.NoEstimationSessions,r,r.NoEstimationSessions_props)
	
	// 
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	rList = append(rList, r)
	return 1, rList, r, nil
}