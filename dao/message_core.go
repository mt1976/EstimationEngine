package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/message.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Message (message)
// Endpoint 	        : Message (Message)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:54
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

// Message_GetList() returns a list of all Message records
func Message_GetList() (int, []dm.Message, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.Message_SQLTable)
	count, messageList, _, _ := message_Fetch(tsql)

	return count, messageList, nil
}

// Message_GetByID() returns a single Message record
func Message_GetByID(id string) (int, dm.Message, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.Message_SQLTable)
	tsql = tsql + " WHERE " + dm.Message_SQLSearchID + "='" + id + "'"
	_, _, messageItem, _ := message_Fetch(tsql)

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return 1, messageItem, nil
}

// Message_DeleteByID() deletes a single Message record
func Message_Delete(id string) {

	object_Table := core.ApplicationSQLSchema() + "." + dm.Message_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Message_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}

// Message_Store() saves/stores a Message record to the database
func Message_Store(r dm.Message, req *http.Request) error {

	err, r := Message_Validate(r)
	if err == nil {
		err = message_Save(r, Audit_User(req))
	} else {
		logs.Information("Message_Store()", err.Error())
	}

	return err
}

// Message_StoreSystem() saves/stores a Message record to the database
func Message_StoreSystem(r dm.Message) error {

	err, r := Message_Validate(r)
	if err == nil {
		err = message_Save(r, Audit_Host())
	} else {
		logs.Information("Message_Store()", err.Error())
	}

	return err
}

// Message_Validate() validates for saves/stores a Message record to the database
func Message_Validate(r dm.Message) (error, dm.Message) {
	var err error
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//

	return err, r
}

//

// message_Save() saves/stores a Message record to the database
func message_Save(r dm.Message, usr string) error {

	var err error

	if len(r.Id) == 0 {
		r.Id = Message_NewID(r)
	}

	// If there are fields below, create the methods in dao\message_impl.go

	logs.Storing("Message", fmt.Sprintf("%v", r))

	//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	ts = addData(ts, dm.Message_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.Message_Id_sql, r.Id)
	ts = addData(ts, dm.Message_Message_sql, r.Message)
	ts = addData(ts, dm.Message_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.Message_SYSWho_sql, r.SYSWho)
	ts = addData(ts, dm.Message_SYSHost_sql, r.SYSHost)
	ts = addData(ts, dm.Message_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.Message_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.Message_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.Message_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.Message_SYSUpdatedHost_sql, r.SYSUpdatedHost)

	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	tsql := "INSERT INTO " + get_TableName(core.ApplicationSQLSchema(), dm.Message_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Message_Delete(r.Id)
	das.Execute(tsql)

	return err

}

// message_Fetch read all Message's
func message_Fetch(tsql string) (int, []dm.Message, dm.Message, error) {

	var recItem dm.Message
	var recList []dm.Message

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(), err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
		// START
		// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
		//
		recItem.SYSId = get_Int(rec, dm.Message_SYSId_sql, "0")
		recItem.Id = get_String(rec, dm.Message_Id_sql, "")
		recItem.Message = get_String(rec, dm.Message_Message_sql, "")
		recItem.SYSCreated = get_String(rec, dm.Message_SYSCreated_sql, "")
		recItem.SYSWho = get_String(rec, dm.Message_SYSWho_sql, "")
		recItem.SYSHost = get_String(rec, dm.Message_SYSHost_sql, "")
		recItem.SYSUpdated = get_String(rec, dm.Message_SYSUpdated_sql, "")
		recItem.SYSCreatedBy = get_String(rec, dm.Message_SYSCreatedBy_sql, "")
		recItem.SYSCreatedHost = get_String(rec, dm.Message_SYSCreatedHost_sql, "")
		recItem.SYSUpdatedBy = get_String(rec, dm.Message_SYSUpdatedBy_sql, "")
		recItem.SYSUpdatedHost = get_String(rec, dm.Message_SYSUpdatedHost_sql, "")

		// If there are fields below, create the methods in adaptor\Message_impl.go

		//
		// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
		// END
		///
		//Add to the list
		//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}

func Message_NewID(r dm.Message) string {

	id := uuid.New().String()

	return id
}

// message_Fetch read all Message's
func Message_New() (int, []dm.Message, dm.Message, error) {

	var r = dm.Message{}
	var rList []dm.Message

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	rList = append(rList, r)

	return 1, rList, r, nil
}
