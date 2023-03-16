package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/estimationstate.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : EstimationState (estimationstate)
// Endpoint 	        : EstimationState (EstimationStateID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/03/2023 at 19:24:48
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

var EstimationState_SQLbase string
var EstimationState_QualifiedName string

func init() {
	EstimationState_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.EstimationState_SQLTable)
	EstimationState_SQLbase = das.SELECTALL + das.FROM + EstimationState_QualifiedName
}

// EstimationState_GetList() returns a list of all EstimationState records
func EstimationState_GetList() (int, []dm.EstimationState, error) {
	count, estimationstateList, err := EstimationState_GetListFiltered("")
	return count, estimationstateList, err
}

// EstimationState_GetListFiltered() returns a filtered list of all EstimationState records
func EstimationState_GetListFiltered(filter string) (int, []dm.EstimationState, error) {

	tsql := EstimationState_SQLbase
	if filter != "" {
		tsql = tsql + " " + das.WHERE + filter
	}
	count, estimationstateList, _, _ := estimationstate_Fetch(tsql)

	return count, estimationstateList, nil
}

// EstimationState_GetLookup() returns a lookup list of all EstimationState items in lookup format
func EstimationState_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, estimationstateList, _ := EstimationState_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: estimationstateList[i].Code, Name: estimationstateList[i].Name})
	}
	return returnList
}

// EstimationState_GetFilteredLookup() returns a lookup list of all EstimationState items in lookup format
func EstimationState_GetFilteredLookup(requestObject string, requestField string) []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	objectName := Translate("ObjectName", requestObject)
	reqField := requestField + "_EstimationState_Filter"

	usage := "Defines a filter for a lookup list of EstimationState records, when requested by " + requestField + "." + core.TEXTAREA_CR
	usage = usage + "Fields can be any of those in the underlying DB table." + core.TEXTAREA_CR
	usage = usage + "Examples Below:" + core.TEXTAREA_CR
	usage = usage + "* datalength(_deleted) = 0 or " + core.TEXTAREA_CR
	usage = usage + "* class IN ('x','y','z')"

	filter, _ := Data_GetString(objectName, reqField, dm.Data_Category_FilterRule, usage)
	if filter == "" {
		logs.Warning("EstimationState_GetFilteredLookup() - No filter found : " + reqField + " for Object: " + objectName)
	}
	count, estimationstateList, _ := EstimationState_GetListFiltered(filter)
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: estimationstateList[i].Code, Name: estimationstateList[i].Name})
	}
	return returnList
}

// EstimationState_GetByID() returns a single EstimationState record
func EstimationState_GetByID(id string) (int, dm.EstimationState, error) {

	tsql := EstimationState_SQLbase
	tsql = tsql + " " + das.WHERE + dm.EstimationState_SQLSearchID + das.EQ + das.ID(id)
	_, _, estimationstateItem, _ := estimationstate_Fetch(tsql)

	estimationstateItem = EstimationState_PostGet(estimationstateItem, id)

	return 1, estimationstateItem, nil
}

func EstimationState_PostGet(estimationstateItem dm.EstimationState, id string) dm.EstimationState {

	return estimationstateItem
}

// EstimationState_DeleteByID() deletes a single EstimationState record
func EstimationState_Delete(id string) {
	EstimationState_SoftDelete(id)
}

// EstimationState_SoftDelete(id string) soft deletes a single EstimationState record
func EstimationState_SoftDelete(id string) {
	//Uses Soft Delete
	_, estimationstateItem, _ := EstimationState_GetByID(id)
	estimationstateItem.SYSDeletedBy = Audit_Update("", Audit_Host())
	estimationstateItem.SYSDeleted = Audit_Update("", Audit_TimeStamp())
	estimationstateItem.SYSDeletedHost = Audit_Update("", Audit_Host())
	_, err := EstimationState_StoreSystem(estimationstateItem)
	if err != nil {
		logs.Error("EstimationState_SoftDelete()", err)
	}
}

// EstimationState_HardDelete(id string) soft deletes a single EstimationState record
func EstimationState_HardDelete(id string) {
	// Uses Hard Delete
	object_Table := EstimationState_QualifiedName
	tsql := das.DELETE + das.FROM + object_Table
	tsql = tsql + " " + das.WHERE + dm.EstimationState_SQLSearchID + das.EQ + das.ID(id)
	das.Execute(tsql)
	//if err != nil {
	//	logs.Error("EstimationState_SoftDelete()",err)
	//}
}

// EstimationState_Store() saves/stores a EstimationState record to the database
func EstimationState_Store(r dm.EstimationState, req *http.Request) (dm.EstimationState, error) {

	r, err := EstimationState_Validate(r)
	if err == nil {
		err = estimationstate_Save(r, Audit_User(req))
	} else {
		logs.Information("EstimationState_Store()", err.Error())
	}

	return r, err
}

// EstimationState_StoreSystem() saves/stores a EstimationState record to the database
func EstimationState_StoreSystem(r dm.EstimationState) (dm.EstimationState, error) {

	r, err := EstimationState_Validate(r)
	if err == nil {
		err = estimationstate_Save(r, Audit_Host())
	} else {
		logs.Information("EstimationState_Store()", err.Error())
	}

	return r, err
}

// EstimationState_StoreProcess() saves/stores a EstimationState record to the database
func EstimationState_StoreProcess(r dm.EstimationState, operator string) (dm.EstimationState, error) {

	r, err := EstimationState_Validate(r)
	if err == nil {
		err = estimationstate_Save(r, operator)
	} else {
		logs.Information("EstimationState_StoreProcess()", err.Error())
	}

	return r, err
}

// EstimationState_Validate() validates for saves/stores a EstimationState record to the database
func EstimationState_Validate(r dm.EstimationState) (dm.EstimationState, error) {
	var err error

	// Cross Validation
	var errVal error
	r, _, errVal = EstimationState_ObjectValidation_impl(PUT, r.EstimationStateID, r)
	if errVal != nil {
		err = errVal
	}
	return r, err
}

//

// estimationstate_Save() saves/stores a EstimationState record to the database
func estimationstate_Save(r dm.EstimationState, usr string) error {

	var err error

	if len(r.EstimationStateID) == 0 {
		r.EstimationStateID = EstimationState_NewID(r)
	}

	// If there are fields below, create the methods in dao\estimationstate_impl.go

	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost, Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("", usr)
	r.SYSUpdatedHost = Audit_Update("", Audit_Host())
	r.SYSDbVersion = core.DB_Version()

	logs.Storing("EstimationState", fmt.Sprintf("%v", r))

	//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	ts = addData(ts, dm.EstimationState_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.EstimationState_EstimationStateID_sql, r.EstimationStateID)
	ts = addData(ts, dm.EstimationState_Code_sql, r.Code)
	ts = addData(ts, dm.EstimationState_Name_sql, r.Name)
	ts = addData(ts, dm.EstimationState_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.EstimationState_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.EstimationState_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.EstimationState_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.EstimationState_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.EstimationState_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.EstimationState_SYSDeleted_sql, r.SYSDeleted)
	ts = addData(ts, dm.EstimationState_SYSDeletedBy_sql, r.SYSDeletedBy)
	ts = addData(ts, dm.EstimationState_SYSDeletedHost_sql, r.SYSDeletedHost)
	ts = addData(ts, dm.EstimationState_SYSDbVersion_sql, r.SYSDbVersion)
	ts = addData(ts, dm.EstimationState_IsLocked_sql, r.IsLocked)
	ts = addData(ts, dm.EstimationState_Notify_sql, r.Notify)
	ts = addData(ts, dm.EstimationState_Migrate_sql, r.Migrate)

	//
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END

	tsql := das.INSERT + das.INTO + EstimationState_QualifiedName
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " " + das.VALUES + "(" + values(ts) + ")"

	EstimationState_HardDelete(r.EstimationStateID)
	das.Execute(tsql)

	return err

}

// estimationstate_Fetch read all EstimationState's
func estimationstate_Fetch(tsql string) (int, []dm.EstimationState, dm.EstimationState, error) {

	var recItem dm.EstimationState
	var recList []dm.EstimationState

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(), err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
		// START
		// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local
		//
		recItem.SYSId = get_Int(rec, dm.EstimationState_SYSId_sql, "0")
		recItem.EstimationStateID = get_String(rec, dm.EstimationState_EstimationStateID_sql, "")
		recItem.Code = get_String(rec, dm.EstimationState_Code_sql, "")
		recItem.Name = get_String(rec, dm.EstimationState_Name_sql, "")
		recItem.SYSCreated = get_String(rec, dm.EstimationState_SYSCreated_sql, "")
		recItem.SYSCreatedBy = get_String(rec, dm.EstimationState_SYSCreatedBy_sql, "")
		recItem.SYSCreatedHost = get_String(rec, dm.EstimationState_SYSCreatedHost_sql, "")
		recItem.SYSUpdated = get_String(rec, dm.EstimationState_SYSUpdated_sql, "")
		recItem.SYSUpdatedBy = get_String(rec, dm.EstimationState_SYSUpdatedBy_sql, "")
		recItem.SYSUpdatedHost = get_String(rec, dm.EstimationState_SYSUpdatedHost_sql, "")
		recItem.SYSDeleted = get_String(rec, dm.EstimationState_SYSDeleted_sql, "")
		recItem.SYSDeletedBy = get_String(rec, dm.EstimationState_SYSDeletedBy_sql, "")
		recItem.SYSDeletedHost = get_String(rec, dm.EstimationState_SYSDeletedHost_sql, "")
		recItem.SYSDbVersion = get_String(rec, dm.EstimationState_SYSDbVersion_sql, "")
		recItem.IsLocked = get_String(rec, dm.EstimationState_IsLocked_sql, "")
		recItem.Notify = get_String(rec, dm.EstimationState_Notify_sql, "")
		recItem.Migrate = get_String(rec, dm.EstimationState_Migrate_sql, "")

		// If there are fields below, create the methods in dao\EstimationState_adaptor.go
		//
		// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local
		// END
		///
		//Add to the list
		//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}

func EstimationState_NewID(r dm.EstimationState) string {

	id := uuid.New().String()

	return id
}

// estimationstate_Fetch read all EstimationState's
func EstimationState_New() (int, []dm.EstimationState, dm.EstimationState, error) {

	var r = dm.EstimationState{}
	var rList []dm.EstimationState

	// START
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//

	//
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
	rList = append(rList, r)
	return 1, rList, r, nil
}
