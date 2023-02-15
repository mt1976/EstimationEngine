package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/projectaction.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : ProjectAction (projectaction)
// Endpoint 	        : ProjectAction (ProjectID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 07/02/2023 at 18:52:38
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

var ProjectAction_SQLbase string
var ProjectAction_QualifiedName string
func init(){
	ProjectAction_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.ProjectAction_SQLTable)
	ProjectAction_SQLbase =  das.SELECTALL + das.FROM + ProjectAction_QualifiedName
}

// ProjectAction_GetList() returns a list of all ProjectAction records
func ProjectAction_GetList() (int, []dm.ProjectAction, error) {
	
	count, projectactionList, err := ProjectAction_GetListFiltered("")
	
	return count, projectactionList, err
}

// ProjectAction_GetListFiltered() returns a filtered list of all ProjectAction records
func ProjectAction_GetListFiltered(filter string) (int, []dm.ProjectAction, error) {
	
	tsql := ProjectAction_SQLbase
	if filter != "" {
		tsql = tsql + " " + das.WHERE + filter
	}
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

// ProjectAction_GetFilteredLookup() returns a lookup list of all ProjectAction items in lookup format
func ProjectAction_GetFilteredLookup(requestObject string,requestField string) []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	reqClass := "ProjectAction"
	reqField := requestObject+"-"+requestField
	reqCategory := "Filter"
	filter,_ := Data_GetString(reqClass, reqField, reqCategory)
	if filter == "" {
		logs.Warning("ProjectAction_GetFilteredLookup() - No filter found for " + reqClass + " " + reqField)
	} 
	count, projectactionList, _ := ProjectAction_GetListFiltered(filter)
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: projectactionList[i].ProjectID, Name: projectactionList[i].Name})
	}
	return returnList
}



// ProjectAction_GetByID() returns a single ProjectAction record
func ProjectAction_GetByID(id string) (int, dm.ProjectAction, error) {


	tsql := ProjectAction_SQLbase
	tsql = tsql + " " + das.WHERE + dm.ProjectAction_SQLSearchID + das.EQ + das.ID(id)
	_, _, projectactionItem, _ := projectaction_Fetch(tsql)

	// START
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	projectactionItem.NoEstimationSessions,projectactionItem.NoEstimationSessions_props = ProjectAction_NoEstimationSessions_impl (GET,id,projectactionItem.NoEstimationSessions,projectactionItem,projectactionItem.NoEstimationSessions_props)
	// 
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, projectactionItem, nil
}



// ProjectAction_DeleteByID() deletes a single ProjectAction record
func ProjectAction_Delete(id string) {


// Uses Hard Delete
	object_Table := ProjectAction_QualifiedName
	tsql := das.DELETE + das.FROM + object_Table
	tsql = tsql + " " + das.WHERE + dm.ProjectAction_SQLSearchID + das.EQ + das.ID(id)
	das.Execute(tsql)	

	
}














	


// ProjectAction_Store() saves/stores a ProjectAction record to the database
func ProjectAction_Store(r dm.ProjectAction,req *http.Request) error {

	r, err := ProjectAction_Validate(r)
	if err == nil {
		err = projectaction_Save(r, Audit_User(req))
	} else {
		logs.Information("ProjectAction_Store()", err.Error())
	}

	return err
}

// ProjectAction_StoreSystem() saves/stores a ProjectAction record to the database
func ProjectAction_StoreSystem(r dm.ProjectAction) error {
	
	r, err := ProjectAction_Validate(r)
	if err == nil {
		err = projectaction_Save(r, Audit_Host())
	} else {
		logs.Information("ProjectAction_Store()", err.Error())
	}

	return err
}

// ProjectAction_Validate() validates for saves/stores a ProjectAction record to the database
func ProjectAction_Validate(r dm.ProjectAction) (dm.ProjectAction, error) {
	var err error
	// START
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	r.NoEstimationSessions,r.NoEstimationSessions_props = ProjectAction_NoEstimationSessions_impl (PUT,r.ProjectID,r.NoEstimationSessions,r,r.NoEstimationSessions_props)
	// 
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	

	return r,err
}
//

// projectaction_Save() saves/stores a ProjectAction record to the database
func projectaction_Save(r dm.ProjectAction,usr string) error {

    var err error



	

	if len(r.ProjectID) == 0 {
		r.ProjectID = ProjectAction_NewID(r)
	}

// If there are fields below, create the methods in dao\projectaction_impl.go




























  r.NoEstimationSessions,err = ProjectAction_NoEstimationSessions_OnStore_impl (r.NoEstimationSessions,r,usr)


	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	r.SYSDbVersion = core.DB_Version()
	
logs.Storing("ProjectAction",fmt.Sprintf("%v", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	ts = addData(ts, dm.ProjectAction_SYSDbVersion_sql, r.SYSDbVersion)
	ts = addData(ts, dm.ProjectAction_Comments_sql, r.Comments)
	ts = addData(ts, dm.ProjectAction_ProjectRate_sql, r.ProjectRate)
	ts = addData(ts, dm.ProjectAction_DefaultRate_sql, r.DefaultRate)
	ts = addData(ts, dm.ProjectAction_ProjectAnalyst_sql, r.ProjectAnalyst)
	ts = addData(ts, dm.ProjectAction_ProjectEngineer_sql, r.ProjectEngineer)
	ts = addData(ts, dm.ProjectAction_ProjectManager_sql, r.ProjectManager)
	ts = addData(ts, dm.ProjectAction_Releases_sql, r.Releases)
	ts = addData(ts, dm.ProjectAction_Notes_sql, r.Notes)
	
		
	// 
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := das.INSERT + das.INTO + ProjectAction_QualifiedName
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " "+das.VALUES +"(" + values(ts) + ")"

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
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	   recItem.SYSDbVersion  = get_String(rec, dm.ProjectAction_SYSDbVersion_sql, "")
	   recItem.Comments  = get_String(rec, dm.ProjectAction_Comments_sql, "")
	   recItem.ProjectRate  = get_String(rec, dm.ProjectAction_ProjectRate_sql, "")
	   recItem.DefaultRate  = get_String(rec, dm.ProjectAction_DefaultRate_sql, "")
	   recItem.ProjectAnalyst  = get_String(rec, dm.ProjectAction_ProjectAnalyst_sql, "")
	   recItem.ProjectEngineer  = get_String(rec, dm.ProjectAction_ProjectEngineer_sql, "")
	   recItem.ProjectManager  = get_String(rec, dm.ProjectAction_ProjectManager_sql, "")
	   recItem.Releases  = get_String(rec, dm.ProjectAction_Releases_sql, "")
	   recItem.Notes  = get_String(rec, dm.ProjectAction_Notes_sql, "")
	
	
	// If there are fields below, create the methods in adaptor\ProjectAction_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	   recItem.NoEstimationSessions  = ProjectAction_NoEstimationSessions_OnFetch_impl (recItem)
	
	// 
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	r.NoEstimationSessions,r.NoEstimationSessions_props = ProjectAction_NoEstimationSessions_impl (NEW,r.ProjectID,r.NoEstimationSessions,r,r.NoEstimationSessions_props)
	// 
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}