package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/projectstate.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : ProjectState (projectstate)
// Endpoint 	        : ProjectState (ProjectStateID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 04/03/2023 at 20:14:14
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	core "github.com/mt1976/ebEstimates/core"
	das "github.com/mt1976/ebEstimates/das"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

var ProjectState_SQLbase string
var ProjectState_QualifiedName string

func init() {
	ProjectState_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.ProjectState_SQLTable)
	ProjectState_SQLbase = das.SELECTALL + das.FROM + ProjectState_QualifiedName
}

// ProjectState_GetList() returns a list of all ProjectState records
func ProjectState_GetList() (int, []dm.ProjectState, error) {
	count, projectstateList, err := ProjectState_GetListFiltered("")
	return count, projectstateList, err
}

// ProjectState_GetListFiltered() returns a filtered list of all ProjectState records
func ProjectState_GetListFiltered(filter string) (int, []dm.ProjectState, error) {

	tsql := ProjectState_SQLbase
	if filter != "" {
		tsql = tsql + " " + das.WHERE + filter
	}
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

// ProjectState_GetFilteredLookup() returns a lookup list of all ProjectState items in lookup format
func ProjectState_GetFilteredLookup(requestObject string, requestField string) []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	objectName := Translate("ObjectName", requestObject)
	reqField := requestField + "_ProjectState_Filter"
	filter, _ := Data_GetString(objectName, reqField, dm.Data_Category_FilterRule)
	if filter == "" {
		logs.Warning("ProjectState_GetFilteredLookup() - No filter found : " + reqField + " for Object: " + objectName)
	}
	count, projectstateList, _ := ProjectState_GetListFiltered(filter)
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: projectstateList[i].Code, Name: projectstateList[i].Name})
	}
	return returnList
}

// ProjectState_GetByID() returns a single ProjectState record
func ProjectState_GetByID(id string) (int, dm.ProjectState, error) {

	tsql := ProjectState_SQLbase
	tsql = tsql + " " + das.WHERE + dm.ProjectState_SQLSearchID + das.EQ + das.ID(id)
	_, _, projectstateItem, _ := projectstate_Fetch(tsql)

	projectstateItem = ProjectState_PostGet(projectstateItem, id)

	return 1, projectstateItem, nil
}

func ProjectState_PostGet(projectstateItem dm.ProjectState, id string) dm.ProjectState {
	// START
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
	return projectstateItem
}

// ProjectState_DeleteByID() deletes a single ProjectState record
func ProjectState_Delete(id string) {
	ProjectState_SoftDelete(id)
}

// ProjectState_SoftDelete(id string) soft deletes a single ProjectState record
func ProjectState_SoftDelete(id string) {
	//Uses Soft Delete
	_, projectstateItem, _ := ProjectState_GetByID(id)
	projectstateItem.SYSDeletedBy = Audit_Update("", Audit_Host())
	projectstateItem.SYSDeleted = Audit_Update("", Audit_TimeStamp())
	projectstateItem.SYSDeletedHost = Audit_Update("", Audit_Host())
	_, err := ProjectState_StoreSystem(projectstateItem)
	if err != nil {
		logs.Error("ProjectState_SoftDelete()", err)
	}
}

// ProjectState_HardDelete(id string) soft deletes a single ProjectState record
func ProjectState_HardDelete(id string) {
	// Uses Hard Delete
	object_Table := ProjectState_QualifiedName
	tsql := das.DELETE + das.FROM + object_Table
	tsql = tsql + " " + das.WHERE + dm.ProjectState_SQLSearchID + das.EQ + das.ID(id)
	das.Execute(tsql)
	//if err != nil {
	//	logs.Error("ProjectState_SoftDelete()",err)
	//}
}

// ProjectState_Store() saves/stores a ProjectState record to the database
func ProjectState_Store(r dm.ProjectState, req *http.Request) (dm.ProjectState, error) {

	r, err := ProjectState_Validate(r)
	if err == nil {
		err = projectstate_Save(r, Audit_User(req))
	} else {
		logs.Information("ProjectState_Store()", err.Error())
	}

	return r, err
}

// ProjectState_StoreSystem() saves/stores a ProjectState record to the database
func ProjectState_StoreSystem(r dm.ProjectState) (dm.ProjectState, error) {

	r, err := ProjectState_Validate(r)
	if err == nil {
		err = projectstate_Save(r, Audit_Host())
	} else {
		logs.Information("ProjectState_Store()", err.Error())
	}

	return r, err
}

// ProjectState_StoreProcess() saves/stores a ProjectState record to the database
func ProjectState_StoreProcess(r dm.ProjectState, operator string) (dm.ProjectState, error) {

	r, err := ProjectState_Validate(r)
	if err == nil {
		err = projectstate_Save(r, operator)
	} else {
		logs.Information("ProjectState_StoreProcess()", err.Error())
	}

	return r, err
}

// ProjectState_Validate() validates for saves/stores a ProjectState record to the database
func ProjectState_Validate(r dm.ProjectState) (dm.ProjectState, error) {
	var err error
	// START
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
	//

	return r, err
}

//

// projectstate_Save() saves/stores a ProjectState record to the database
func projectstate_Save(r dm.ProjectState, usr string) error {

	var err error

	if len(r.ProjectStateID) == 0 {
		r.ProjectStateID = ProjectState_NewID(r)
	}

	// If there are fields below, create the methods in dao\projectstate_impl.go

	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost, Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("", usr)
	r.SYSUpdatedHost = Audit_Update("", Audit_Host())
	r.SYSDbVersion = core.DB_Version()

	logs.Storing("ProjectState", fmt.Sprintf("%v", r))

	//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
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
	ts = addData(ts, dm.ProjectState_SYSDbVersion_sql, r.SYSDbVersion)
	ts = addData(ts, dm.ProjectState_IsLocked_sql, r.IsLocked)
	ts = addData(ts, dm.ProjectState_Notify_sql, r.Notify)

	//
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END

	tsql := das.INSERT + das.INTO + ProjectState_QualifiedName
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " " + das.VALUES + "(" + values(ts) + ")"

	ProjectState_HardDelete(r.ProjectStateID)
	das.Execute(tsql)

	return err

}

// projectstate_Fetch read all ProjectState's
func projectstate_Fetch(tsql string) (int, []dm.ProjectState, dm.ProjectState, error) {

	var recItem dm.ProjectState
	var recList []dm.ProjectState

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(), err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
		// START
		// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
		//
		recItem.SYSId = get_Int(rec, dm.ProjectState_SYSId_sql, "0")
		recItem.ProjectStateID = get_String(rec, dm.ProjectState_ProjectStateID_sql, "")
		recItem.Code = get_String(rec, dm.ProjectState_Code_sql, "")
		recItem.Name = get_String(rec, dm.ProjectState_Name_sql, "")
		recItem.SYSCreated = get_String(rec, dm.ProjectState_SYSCreated_sql, "")
		recItem.SYSCreatedBy = get_String(rec, dm.ProjectState_SYSCreatedBy_sql, "")
		recItem.SYSCreatedHost = get_String(rec, dm.ProjectState_SYSCreatedHost_sql, "")
		recItem.SYSUpdated = get_String(rec, dm.ProjectState_SYSUpdated_sql, "")
		recItem.SYSUpdatedBy = get_String(rec, dm.ProjectState_SYSUpdatedBy_sql, "")
		recItem.SYSUpdatedHost = get_String(rec, dm.ProjectState_SYSUpdatedHost_sql, "")
		recItem.SYSDeleted = get_String(rec, dm.ProjectState_SYSDeleted_sql, "")
		recItem.SYSDeletedBy = get_String(rec, dm.ProjectState_SYSDeletedBy_sql, "")
		recItem.SYSDeletedHost = get_String(rec, dm.ProjectState_SYSDeletedHost_sql, "")
		recItem.SYSDbVersion = get_String(rec, dm.ProjectState_SYSDbVersion_sql, "")
		recItem.IsLocked = get_String(rec, dm.ProjectState_IsLocked_sql, "")
		recItem.Notify = get_String(rec, dm.ProjectState_Notify_sql, "")

		// If there are fields below, create the methods in dao\ProjectState_adaptor.go
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

func ProjectState_NewID(r dm.ProjectState) string {

	id := uuid.New().String()

	return id
}

// projectstate_Fetch read all ProjectState's
func ProjectState_New() (int, []dm.ProjectState, dm.ProjectState, error) {

	var r = dm.ProjectState{}
	var rList []dm.ProjectState

	// START
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//

	//
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
	rList = append(rList, r)
	return 1, rList, r, nil
}
