package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/translation.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Translation (translation)
// Endpoint 	        : Translation (Message)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 10/03/2023 at 19:54:36
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

var Translation_SQLbase string
var Translation_QualifiedName string

func init() {
	Translation_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.Translation_SQLTable)
	Translation_SQLbase = das.SELECTALL + das.FROM + Translation_QualifiedName
}

// Translation_GetList() returns a list of all Translation records
func Translation_GetList() (int, []dm.Translation, error) {
	count, translationList, err := Translation_GetListFiltered("")
	return count, translationList, err
}

// Translation_GetListFiltered() returns a filtered list of all Translation records
func Translation_GetListFiltered(filter string) (int, []dm.Translation, error) {

	tsql := Translation_SQLbase
	if filter != "" {
		tsql = tsql + " " + das.WHERE + filter
	}
	count, translationList, _, _ := translation_Fetch(tsql)

	return count, translationList, nil
}

// Translation_GetByID() returns a single Translation record
func Translation_GetByID(id string) (int, dm.Translation, error) {

	tsql := Translation_SQLbase
	tsql = tsql + " " + das.WHERE + dm.Translation_SQLSearchID + das.EQ + das.ID(id)
	_, _, translationItem, _ := translation_Fetch(tsql)

	translationItem = Translation_PostGet(translationItem, id)

	return 1, translationItem, nil
}

func Translation_PostGet(translationItem dm.Translation, id string) dm.Translation {
	// START
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
	return translationItem
}

// Translation_DeleteByID() deletes a single Translation record
func Translation_Delete(id string) {
	Translation_HardDelete(id)
}

// Translation_HardDelete(id string) soft deletes a single Translation record
func Translation_HardDelete(id string) {
	// Uses Hard Delete
	object_Table := Translation_QualifiedName
	tsql := das.DELETE + das.FROM + object_Table
	tsql = tsql + " " + das.WHERE + dm.Translation_SQLSearchID + das.EQ + das.ID(id)
	das.Execute(tsql)
	//if err != nil {
	//	logs.Error("Translation_SoftDelete()",err)
	//}
}

// Translation_Store() saves/stores a Translation record to the database
func Translation_Store(r dm.Translation, req *http.Request) (dm.Translation, error) {

	r, err := Translation_Validate(r)
	if err == nil {
		err = translation_Save(r, Audit_User(req))
	} else {
		logs.Information("Translation_Store()", err.Error())
	}

	return r, err
}

// Translation_StoreSystem() saves/stores a Translation record to the database
func Translation_StoreSystem(r dm.Translation) (dm.Translation, error) {

	r, err := Translation_Validate(r)
	if err == nil {
		err = translation_Save(r, Audit_Host())
	} else {
		logs.Information("Translation_Store()", err.Error())
	}

	return r, err
}

// Translation_StoreProcess() saves/stores a Translation record to the database
func Translation_StoreProcess(r dm.Translation, operator string) (dm.Translation, error) {

	r, err := Translation_Validate(r)
	if err == nil {
		err = translation_Save(r, operator)
	} else {
		logs.Information("Translation_StoreProcess()", err.Error())
	}

	return r, err
}

// Translation_Validate() validates for saves/stores a Translation record to the database
func Translation_Validate(r dm.Translation) (dm.Translation, error) {
	var err error
	// START
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	//

	return r, err
}

//

// translation_Save() saves/stores a Translation record to the database
func translation_Save(r dm.Translation, usr string) error {

	var err error

	if len(r.Id) == 0 {
		r.Id = Translation_NewID(r)
	}

	// If there are fields below, create the methods in dao\translation_impl.go

	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost, Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("", usr)
	r.SYSUpdatedHost = Audit_Update("", Audit_Host())
	r.SYSDbVersion = core.DB_Version()

	logs.Storing("Translation", fmt.Sprintf("%v", r))

	//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	ts = addData(ts, dm.Translation_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.Translation_Id_sql, r.Id)
	ts = addData(ts, dm.Translation_Class_sql, r.Class)
	ts = addData(ts, dm.Translation_Message_sql, r.Message)
	ts = addData(ts, dm.Translation_Translation_sql, r.Translation)
	ts = addData(ts, dm.Translation_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.Translation_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.Translation_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.Translation_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.Translation_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.Translation_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.Translation_SYSDeleted_sql, r.SYSDeleted)
	ts = addData(ts, dm.Translation_SYSDeletedBy_sql, r.SYSDeletedBy)
	ts = addData(ts, dm.Translation_SYSDeletedHost_sql, r.SYSDeletedHost)
	ts = addData(ts, dm.Translation_SYSDbVersion_sql, r.SYSDbVersion)

	//
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END

	tsql := das.INSERT + das.INTO + Translation_QualifiedName
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " " + das.VALUES + "(" + values(ts) + ")"

	Translation_HardDelete(r.Id)
	das.Execute(tsql)

	return err

}

// translation_Fetch read all Translation's
func translation_Fetch(tsql string) (int, []dm.Translation, dm.Translation, error) {

	var recItem dm.Translation
	var recList []dm.Translation

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(), err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
		// START
		// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
		//
		recItem.SYSId = get_Int(rec, dm.Translation_SYSId_sql, "0")
		recItem.Id = get_String(rec, dm.Translation_Id_sql, "")
		recItem.Class = get_String(rec, dm.Translation_Class_sql, "")
		recItem.Message = get_String(rec, dm.Translation_Message_sql, "")
		recItem.Translation = get_String(rec, dm.Translation_Translation_sql, "")
		recItem.SYSCreated = get_String(rec, dm.Translation_SYSCreated_sql, "")
		recItem.SYSUpdated = get_String(rec, dm.Translation_SYSUpdated_sql, "")
		recItem.SYSCreatedBy = get_String(rec, dm.Translation_SYSCreatedBy_sql, "")
		recItem.SYSCreatedHost = get_String(rec, dm.Translation_SYSCreatedHost_sql, "")
		recItem.SYSUpdatedBy = get_String(rec, dm.Translation_SYSUpdatedBy_sql, "")
		recItem.SYSUpdatedHost = get_String(rec, dm.Translation_SYSUpdatedHost_sql, "")
		recItem.SYSDeleted = get_String(rec, dm.Translation_SYSDeleted_sql, "")
		recItem.SYSDeletedBy = get_String(rec, dm.Translation_SYSDeletedBy_sql, "")
		recItem.SYSDeletedHost = get_String(rec, dm.Translation_SYSDeletedHost_sql, "")
		recItem.SYSDbVersion = get_String(rec, dm.Translation_SYSDbVersion_sql, "")

		// If there are fields below, create the methods in dao\Translation_adaptor.go
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

func Translation_NewID(r dm.Translation) string {

	id := uuid.New().String()

	return id
}

// translation_Fetch read all Translation's
func Translation_New() (int, []dm.Translation, dm.Translation, error) {

	var r = dm.Translation{}
	var rList []dm.Translation

	// START
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//

	//
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
	rList = append(rList, r)
	return 1, rList, r, nil
}
