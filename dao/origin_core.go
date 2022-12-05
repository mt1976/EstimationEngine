package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/origin.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Origin (origin)
// Endpoint 	        : Origin (OriginID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 05/12/2022 at 21:52:44
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	adaptor "github.com/mt1976/ebEstimates/adaptor"
	core "github.com/mt1976/ebEstimates/core"
	das "github.com/mt1976/ebEstimates/das"

	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// Origin_GetList() returns a list of all Origin records
func Origin_GetList() (int, []dm.Origin, error) {

	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Origin_SQLTable)
	count, originList, _, _ := origin_Fetch(tsql)

	return count, originList, nil
}

// Origin_GetLookup() returns a lookup list of all Origin items in lookup format
func Origin_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, originList, _ := Origin_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: originList[i].Code, Name: originList[i].FullName})
	}
	return returnList
}

// Origin_GetByID() returns a single Origin record
func Origin_GetByID(id string) (int, dm.Origin, error) {

	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Origin_SQLTable)
	tsql = tsql + " WHERE " + dm.Origin_SQLSearchID + "='" + id + "'"
	_, _, originItem, _ := origin_Fetch(tsql)

	// START
	// Dynamically generated 05/12/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	originItem.NoActiveProjects, originItem.NoActiveProjects_props = adaptor.Origin_NoActiveProjects_impl(adaptor.GET, id, originItem.NoActiveProjects, originItem, originItem.NoActiveProjects_props)
	//
	// Dynamically generated 05/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return 1, originItem, nil
}

// Origin_DeleteByID() deletes a single Origin record
func Origin_Delete(id string) {

	// Uses Hard Delete
	object_Table := core.GetSQLSchema(core.ApplicationPropertiesDB) + "." + dm.Origin_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Origin_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}

// Origin_SoftDeleteByID() soft deletes a single Origin record
func Origin_SoftDelete(id string) {
	//Uses Soft Delete
	_, originItem, _ := Origin_GetByID(id)
	originItem.SYSDeletedBy = Audit_Update("", Audit_Host())
	originItem.SYSDeleted = Audit_Update("", Audit_TimeStamp())
	originItem.SYSDeletedHost = Audit_Update("", Audit_Host())
	Origin_StoreSystem(originItem)
}

// Origin_Store() saves/stores a Origin record to the database
func Origin_Store(r dm.Origin, req *http.Request) error {

	err, r := Origin_Validate(r)
	if err == nil {
		err = origin_Save(r, Audit_User(req))
	} else {
		logs.Information("Origin_Store()", err.Error())
	}

	return err
}

// Origin_StoreSystem() saves/stores a Origin record to the database
func Origin_StoreSystem(r dm.Origin) error {

	err, r := Origin_Validate(r)
	if err == nil {
		err = origin_Save(r, Audit_Host())
	} else {
		logs.Information("Origin_Store()", err.Error())
	}

	return err
}

// Origin_Validate() validates for saves/stores a Origin record to the database
func Origin_Validate(r dm.Origin) (error, dm.Origin) {
	var err error
	// START
	// Dynamically generated 05/12/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	r.NoActiveProjects, r.NoActiveProjects_props = adaptor.Origin_NoActiveProjects_impl(adaptor.PUT, r.OriginID, r.NoActiveProjects, r, r.NoActiveProjects_props)
	//
	// Dynamically generated 05/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//

	return err, r
}

//

// origin_Save() saves/stores a Origin record to the database
func origin_Save(r dm.Origin, usr string) error {

	var err error

	if len(r.OriginID) == 0 {
		r.OriginID = Origin_NewID(r)
	}

	// If there are fields below, create the methods in dao\origin_impl.go

	r.NoActiveProjects, err = adaptor.Origin_NoActiveProjects_OnStore_impl(r.NoActiveProjects, r, usr)

	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost, Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("", usr)
	r.SYSUpdatedHost = Audit_Update("", Audit_Host())

	logs.Storing("Origin", fmt.Sprintf("%s", r))

	//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 05/12/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	ts = addData(ts, dm.Origin_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.Origin_OriginID_sql, r.OriginID)
	ts = addData(ts, dm.Origin_StateID_sql, r.StateID)
	ts = addData(ts, dm.Origin_DocTypeID_sql, r.DocTypeID)
	ts = addData(ts, dm.Origin_Code_sql, r.Code)
	ts = addData(ts, dm.Origin_FullName_sql, r.FullName)
	ts = addData(ts, dm.Origin_Rate_sql, r.Rate)
	ts = addData(ts, dm.Origin_Notes_sql, r.Notes)
	ts = addData(ts, dm.Origin_StartDate_sql, r.StartDate)
	ts = addData(ts, dm.Origin_EndDate_sql, r.EndDate)
	ts = addData(ts, dm.Origin_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.Origin_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.Origin_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.Origin_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.Origin_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.Origin_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.Origin_SYSDeleted_sql, r.SYSDeleted)
	ts = addData(ts, dm.Origin_SYSDeletedBy_sql, r.SYSDeletedBy)
	ts = addData(ts, dm.Origin_SYSDeletedHost_sql, r.SYSDeletedHost)
	ts = addData(ts, dm.Origin_SYSActivity_sql, r.SYSActivity)
	ts = addData(ts, dm.Origin_Currency_sql, r.Currency)

	//
	// Dynamically generated 05/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	tsql := "INSERT INTO " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Origin_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Origin_Delete(r.OriginID)
	das.Execute(tsql)

	return err

}

// origin_Fetch read all Origin's
func origin_Fetch(tsql string) (int, []dm.Origin, dm.Origin, error) {

	var recItem dm.Origin
	var recList []dm.Origin

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(), err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
		// START
		// Dynamically generated 05/12/2022 by matttownsend (Matt Townsend) on silicon.local
		//
		recItem.SYSId = get_Int(rec, dm.Origin_SYSId_sql, "0")
		recItem.OriginID = get_String(rec, dm.Origin_OriginID_sql, "")
		recItem.StateID = get_String(rec, dm.Origin_StateID_sql, "")
		recItem.DocTypeID = get_String(rec, dm.Origin_DocTypeID_sql, "")
		recItem.Code = get_String(rec, dm.Origin_Code_sql, "")
		recItem.FullName = get_String(rec, dm.Origin_FullName_sql, "")
		recItem.Rate = get_String(rec, dm.Origin_Rate_sql, "")
		recItem.Notes = get_String(rec, dm.Origin_Notes_sql, "")
		recItem.StartDate = get_String(rec, dm.Origin_StartDate_sql, "")
		recItem.EndDate = get_String(rec, dm.Origin_EndDate_sql, "")
		recItem.SYSCreated = get_String(rec, dm.Origin_SYSCreated_sql, "")
		recItem.SYSCreatedBy = get_String(rec, dm.Origin_SYSCreatedBy_sql, "")
		recItem.SYSCreatedHost = get_String(rec, dm.Origin_SYSCreatedHost_sql, "")
		recItem.SYSUpdated = get_String(rec, dm.Origin_SYSUpdated_sql, "")
		recItem.SYSUpdatedBy = get_String(rec, dm.Origin_SYSUpdatedBy_sql, "")
		recItem.SYSUpdatedHost = get_String(rec, dm.Origin_SYSUpdatedHost_sql, "")
		recItem.SYSDeleted = get_String(rec, dm.Origin_SYSDeleted_sql, "")
		recItem.SYSDeletedBy = get_String(rec, dm.Origin_SYSDeletedBy_sql, "")
		recItem.SYSDeletedHost = get_String(rec, dm.Origin_SYSDeletedHost_sql, "")
		recItem.SYSActivity = get_String(rec, dm.Origin_SYSActivity_sql, "")
		recItem.Currency = get_String(rec, dm.Origin_Currency_sql, "")

		// If there are fields below, create the methods in adaptor\Origin_impl.go

		recItem.NoActiveProjects = adaptor.Origin_NoActiveProjects_OnFetch_impl(recItem)

		//
		// Dynamically generated 05/12/2022 by matttownsend (Matt Townsend) on silicon.local
		// END
		///
		//Add to the list
		//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}

func Origin_NewID(r dm.Origin) string {

	id := uuid.New().String()

	return id
}

// origin_Fetch read all Origin's
func Origin_New() (int, []dm.Origin, dm.Origin, error) {

	var r = dm.Origin{}
	var rList []dm.Origin

	// START
	// Dynamically generated 05/12/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	r.NoActiveProjects, r.NoActiveProjects_props = adaptor.Origin_NoActiveProjects_impl(adaptor.NEW, r.OriginID, r.NoActiveProjects, r, r.NoActiveProjects_props)
	//
	// Dynamically generated 05/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	rList = append(rList, r)

	return 1, rList, r, nil
}
