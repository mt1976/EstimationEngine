package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/inbox.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Inbox (inbox)
// Endpoint 	        : Inbox (Message)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 04/03/2023 at 20:14:12
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

var Inbox_SQLbase string
var Inbox_QualifiedName string

func init() {
	Inbox_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.Inbox_SQLTable)
	Inbox_SQLbase = das.SELECTALL + das.FROM + Inbox_QualifiedName
}

// Inbox_GetList() returns a list of all Inbox records
func Inbox_GetList() (int, []dm.Inbox, error) {
	count, inboxList, err := Inbox_GetListFiltered("")
	return count, inboxList, err
}

// Inbox_GetListFiltered() returns a filtered list of all Inbox records
func Inbox_GetListFiltered(filter string) (int, []dm.Inbox, error) {

	tsql := Inbox_SQLbase
	if filter != "" {
		tsql = tsql + " " + das.WHERE + filter
	}
	count, inboxList, _, _ := inbox_Fetch(tsql)

	return count, inboxList, nil
}

// Inbox_GetByID() returns a single Inbox record
func Inbox_GetByID(id string) (int, dm.Inbox, error) {

	tsql := Inbox_SQLbase
	tsql = tsql + " " + das.WHERE + dm.Inbox_SQLSearchID + das.EQ + das.ID(id)
	_, _, inboxItem, _ := inbox_Fetch(tsql)

	inboxItem = Inbox_PostGet(inboxItem, id)

	return 1, inboxItem, nil
}

func Inbox_PostGet(inboxItem dm.Inbox, id string) dm.Inbox {
	// START
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
	return inboxItem
}

// Inbox_DeleteByID() deletes a single Inbox record
func Inbox_Delete(id string) {
	Inbox_HardDelete(id)
}

// Inbox_HardDelete(id string) soft deletes a single Inbox record
func Inbox_HardDelete(id string) {
	// Uses Hard Delete
	object_Table := Inbox_QualifiedName
	tsql := das.DELETE + das.FROM + object_Table
	tsql = tsql + " " + das.WHERE + dm.Inbox_SQLSearchID + das.EQ + das.ID(id)
	das.Execute(tsql)
	//if err != nil {
	//	logs.Error("Inbox_SoftDelete()",err)
	//}
}

// Inbox_Store() saves/stores a Inbox record to the database
func Inbox_Store(r dm.Inbox, req *http.Request) (dm.Inbox, error) {

	r, err := Inbox_Validate(r)
	if err == nil {
		err = inbox_Save(r, Audit_User(req))
	} else {
		logs.Information("Inbox_Store()", err.Error())
	}

	return r, err
}

// Inbox_StoreSystem() saves/stores a Inbox record to the database
func Inbox_StoreSystem(r dm.Inbox) (dm.Inbox, error) {

	r, err := Inbox_Validate(r)
	if err == nil {
		err = inbox_Save(r, Audit_Host())
	} else {
		logs.Information("Inbox_Store()", err.Error())
	}

	return r, err
}

// Inbox_StoreProcess() saves/stores a Inbox record to the database
func Inbox_StoreProcess(r dm.Inbox, operator string) (dm.Inbox, error) {

	r, err := Inbox_Validate(r)
	if err == nil {
		err = inbox_Save(r, operator)
	} else {
		logs.Information("Inbox_StoreProcess()", err.Error())
	}

	return r, err
}

// Inbox_Validate() validates for saves/stores a Inbox record to the database
func Inbox_Validate(r dm.Inbox) (dm.Inbox, error) {
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

// inbox_Save() saves/stores a Inbox record to the database
func inbox_Save(r dm.Inbox, usr string) error {

	var err error

	if len(r.MailId) == 0 {
		r.MailId = Inbox_NewID(r)
	}

	// If there are fields below, create the methods in dao\inbox_impl.go

	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost, Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("", usr)
	r.SYSUpdatedHost = Audit_Update("", Audit_Host())
	r.SYSDbVersion = core.DB_Version()

	logs.Storing("Inbox", fmt.Sprintf("%v", r))

	//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	ts = addData(ts, dm.Inbox_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.Inbox_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.Inbox_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.Inbox_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.Inbox_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.Inbox_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.Inbox_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.Inbox_MailId_sql, r.MailId)
	ts = addData(ts, dm.Inbox_MailTo_sql, r.MailTo)
	ts = addData(ts, dm.Inbox_MailFrom_sql, r.MailFrom)
	ts = addData(ts, dm.Inbox_MailSource_sql, r.MailSource)
	ts = addData(ts, dm.Inbox_MailSent_sql, r.MailSent)
	ts = addData(ts, dm.Inbox_MailUnread_sql, r.MailUnread)
	ts = addData(ts, dm.Inbox_MailSubject_sql, r.MailSubject)
	ts = addData(ts, dm.Inbox_MailContent_sql, r.MailContent)
	ts = addData(ts, dm.Inbox_MailAcknowledged_sql, r.MailAcknowledged)
	ts = addData(ts, dm.Inbox_SYSDeleted_sql, r.SYSDeleted)
	ts = addData(ts, dm.Inbox_SYSDeletedBy_sql, r.SYSDeletedBy)
	ts = addData(ts, dm.Inbox_SYSDeletedHost_sql, r.SYSDeletedHost)
	ts = addData(ts, dm.Inbox_SYSDbVersion_sql, r.SYSDbVersion)

	//
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END

	tsql := das.INSERT + das.INTO + Inbox_QualifiedName
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " " + das.VALUES + "(" + values(ts) + ")"

	Inbox_HardDelete(r.MailId)
	das.Execute(tsql)

	return err

}

// inbox_Fetch read all Inbox's
func inbox_Fetch(tsql string) (int, []dm.Inbox, dm.Inbox, error) {

	var recItem dm.Inbox
	var recList []dm.Inbox

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(), err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
		// START
		// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
		//
		recItem.SYSId = get_Int(rec, dm.Inbox_SYSId_sql, "0")
		recItem.SYSCreated = get_String(rec, dm.Inbox_SYSCreated_sql, "")
		recItem.SYSUpdated = get_String(rec, dm.Inbox_SYSUpdated_sql, "")
		recItem.SYSCreatedBy = get_String(rec, dm.Inbox_SYSCreatedBy_sql, "")
		recItem.SYSCreatedHost = get_String(rec, dm.Inbox_SYSCreatedHost_sql, "")
		recItem.SYSUpdatedBy = get_String(rec, dm.Inbox_SYSUpdatedBy_sql, "")
		recItem.SYSUpdatedHost = get_String(rec, dm.Inbox_SYSUpdatedHost_sql, "")
		recItem.MailId = get_String(rec, dm.Inbox_MailId_sql, "")
		recItem.MailTo = get_String(rec, dm.Inbox_MailTo_sql, "")
		recItem.MailFrom = get_String(rec, dm.Inbox_MailFrom_sql, "")
		recItem.MailSource = get_String(rec, dm.Inbox_MailSource_sql, "")
		recItem.MailSent = get_String(rec, dm.Inbox_MailSent_sql, "")
		recItem.MailUnread = get_String(rec, dm.Inbox_MailUnread_sql, "")
		recItem.MailSubject = get_String(rec, dm.Inbox_MailSubject_sql, "")
		recItem.MailContent = get_String(rec, dm.Inbox_MailContent_sql, "")
		recItem.MailAcknowledged = get_String(rec, dm.Inbox_MailAcknowledged_sql, "")
		recItem.SYSDeleted = get_String(rec, dm.Inbox_SYSDeleted_sql, "")
		recItem.SYSDeletedBy = get_String(rec, dm.Inbox_SYSDeletedBy_sql, "")
		recItem.SYSDeletedHost = get_String(rec, dm.Inbox_SYSDeletedHost_sql, "")
		recItem.SYSDbVersion = get_String(rec, dm.Inbox_SYSDbVersion_sql, "")

		// If there are fields below, create the methods in dao\Inbox_adaptor.go
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

func Inbox_NewID(r dm.Inbox) string {

	id := uuid.New().String()

	return id
}

// inbox_Fetch read all Inbox's
func Inbox_New() (int, []dm.Inbox, dm.Inbox, error) {

	var r = dm.Inbox{}
	var rList []dm.Inbox

	// START
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//

	//
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
	rList = append(rList, r)
	return 1, rList, r, nil
}
