package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/userrole.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : UserRole (userrole)
// Endpoint 	        : UserRole (Id)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 10/03/2023 at 19:54:37
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

var UserRole_SQLbase string
var UserRole_QualifiedName string

func init() {
	UserRole_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.UserRole_SQLTable)
	UserRole_SQLbase = das.SELECTALL + das.FROM + UserRole_QualifiedName
}

// UserRole_GetList() returns a list of all UserRole records
func UserRole_GetList() (int, []dm.UserRole, error) {
	count, userroleList, err := UserRole_GetListFiltered("")
	return count, userroleList, err
}

// UserRole_GetListFiltered() returns a filtered list of all UserRole records
func UserRole_GetListFiltered(filter string) (int, []dm.UserRole, error) {

	tsql := UserRole_SQLbase
	if filter != "" {
		tsql = tsql + " " + das.WHERE + filter
	}
	count, userroleList, _, _ := userrole_Fetch(tsql)

	return count, userroleList, nil
}

// UserRole_GetLookup() returns a lookup list of all UserRole items in lookup format
func UserRole_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, userroleList, _ := UserRole_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: userroleList[i].Id, Name: userroleList[i].Name})
	}
	return returnList
}

// UserRole_GetFilteredLookup() returns a lookup list of all UserRole items in lookup format
func UserRole_GetFilteredLookup(requestObject string, requestField string) []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	objectName := Translate("ObjectName", requestObject)
	reqField := requestField + "_UserRole_Filter"
	filter, _ := Data_GetString(objectName, reqField, dm.Data_Category_FilterRule)
	if filter == "" {
		logs.Warning("UserRole_GetFilteredLookup() - No filter found : " + reqField + " for Object: " + objectName)
	}
	count, userroleList, _ := UserRole_GetListFiltered(filter)
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: userroleList[i].Id, Name: userroleList[i].Name})
	}
	return returnList
}

// UserRole_GetByID() returns a single UserRole record
func UserRole_GetByID(id string) (int, dm.UserRole, error) {

	tsql := UserRole_SQLbase
	tsql = tsql + " " + das.WHERE + dm.UserRole_SQLSearchID + das.EQ + das.ID(id)
	_, _, userroleItem, _ := userrole_Fetch(tsql)

	userroleItem = UserRole_PostGet(userroleItem, id)

	return 1, userroleItem, nil
}

func UserRole_PostGet(userroleItem dm.UserRole, id string) dm.UserRole {
	// START
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
	return userroleItem
}

// UserRole_DeleteByID() deletes a single UserRole record
func UserRole_Delete(id string) {
	UserRole_HardDelete(id)
}

// UserRole_HardDelete(id string) soft deletes a single UserRole record
func UserRole_HardDelete(id string) {
	// Uses Hard Delete
	object_Table := UserRole_QualifiedName
	tsql := das.DELETE + das.FROM + object_Table
	tsql = tsql + " " + das.WHERE + dm.UserRole_SQLSearchID + das.EQ + das.ID(id)
	das.Execute(tsql)
	//if err != nil {
	//	logs.Error("UserRole_SoftDelete()",err)
	//}
}

// UserRole_Store() saves/stores a UserRole record to the database
func UserRole_Store(r dm.UserRole, req *http.Request) (dm.UserRole, error) {

	r, err := UserRole_Validate(r)
	if err == nil {
		err = userrole_Save(r, Audit_User(req))
	} else {
		logs.Information("UserRole_Store()", err.Error())
	}

	return r, err
}

// UserRole_StoreSystem() saves/stores a UserRole record to the database
func UserRole_StoreSystem(r dm.UserRole) (dm.UserRole, error) {

	r, err := UserRole_Validate(r)
	if err == nil {
		err = userrole_Save(r, Audit_Host())
	} else {
		logs.Information("UserRole_Store()", err.Error())
	}

	return r, err
}

// UserRole_StoreProcess() saves/stores a UserRole record to the database
func UserRole_StoreProcess(r dm.UserRole, operator string) (dm.UserRole, error) {

	r, err := UserRole_Validate(r)
	if err == nil {
		err = userrole_Save(r, operator)
	} else {
		logs.Information("UserRole_StoreProcess()", err.Error())
	}

	return r, err
}

// UserRole_Validate() validates for saves/stores a UserRole record to the database
func UserRole_Validate(r dm.UserRole) (dm.UserRole, error) {
	var err error
	// START
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	//

	return r, err
}

//

// userrole_Save() saves/stores a UserRole record to the database
func userrole_Save(r dm.UserRole, usr string) error {

	var err error

	if len(r.Id) == 0 {
		r.Id = UserRole_NewID(r)
	}

	// If there are fields below, create the methods in dao\userrole_impl.go

	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost, Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("", usr)
	r.SYSUpdatedHost = Audit_Update("", Audit_Host())
	r.SYSDbVersion = core.DB_Version()

	logs.Storing("UserRole", fmt.Sprintf("%v", r))

	//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	ts = addData(ts, dm.UserRole_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.UserRole_Id_sql, r.Id)
	ts = addData(ts, dm.UserRole_Name_sql, r.Name)
	ts = addData(ts, dm.UserRole_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.UserRole_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.UserRole_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.UserRole_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.UserRole_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.UserRole_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.UserRole_SYSDeleted_sql, r.SYSDeleted)
	ts = addData(ts, dm.UserRole_SYSDeletedBy_sql, r.SYSDeletedBy)
	ts = addData(ts, dm.UserRole_SYSDeletedHost_sql, r.SYSDeletedHost)
	ts = addData(ts, dm.UserRole_SYSDbVersion_sql, r.SYSDbVersion)

	//
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END

	tsql := das.INSERT + das.INTO + UserRole_QualifiedName
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " " + das.VALUES + "(" + values(ts) + ")"

	UserRole_HardDelete(r.Id)
	das.Execute(tsql)

	return err

}

// userrole_Fetch read all UserRole's
func userrole_Fetch(tsql string) (int, []dm.UserRole, dm.UserRole, error) {

	var recItem dm.UserRole
	var recList []dm.UserRole

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(), err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
		// START
		// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
		//
		recItem.SYSId = get_Int(rec, dm.UserRole_SYSId_sql, "0")
		recItem.Id = get_String(rec, dm.UserRole_Id_sql, "")
		recItem.Name = get_String(rec, dm.UserRole_Name_sql, "")
		recItem.SYSCreatedBy = get_String(rec, dm.UserRole_SYSCreatedBy_sql, "")
		recItem.SYSCreatedHost = get_String(rec, dm.UserRole_SYSCreatedHost_sql, "")
		recItem.SYSUpdatedBy = get_String(rec, dm.UserRole_SYSUpdatedBy_sql, "")
		recItem.SYSUpdatedHost = get_String(rec, dm.UserRole_SYSUpdatedHost_sql, "")
		recItem.SYSUpdated = get_String(rec, dm.UserRole_SYSUpdated_sql, "")
		recItem.SYSCreated = get_String(rec, dm.UserRole_SYSCreated_sql, "")
		recItem.SYSDeleted = get_String(rec, dm.UserRole_SYSDeleted_sql, "")
		recItem.SYSDeletedBy = get_String(rec, dm.UserRole_SYSDeletedBy_sql, "")
		recItem.SYSDeletedHost = get_String(rec, dm.UserRole_SYSDeletedHost_sql, "")
		recItem.SYSDbVersion = get_String(rec, dm.UserRole_SYSDbVersion_sql, "")

		// If there are fields below, create the methods in dao\UserRole_adaptor.go
		//
		// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
		// END
		///
		//Add to the list
		//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}

func UserRole_NewID(r dm.UserRole) string {

	id := uuid.New().String()

	return id
}

// userrole_Fetch read all UserRole's
func UserRole_New() (int, []dm.UserRole, dm.UserRole, error) {

	var r = dm.UserRole{}
	var rList []dm.UserRole

	// START
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//

	//
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
	rList = append(rList, r)
	return 1, rList, r, nil
}
