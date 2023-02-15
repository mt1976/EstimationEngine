package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/session.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Session (session)
// Endpoint 	        : Session (SessionID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 07/02/2023 at 18:52:39
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

var Session_SQLbase string
var Session_QualifiedName string
func init(){
	Session_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.Session_SQLTable)
	Session_SQLbase =  das.SELECTALL + das.FROM + Session_QualifiedName
}

// Session_GetList() returns a list of all Session records
func Session_GetList() (int, []dm.Session, error) {
	
	count, sessionList, err := Session_GetListFiltered("")
	
	return count, sessionList, err
}

// Session_GetListFiltered() returns a filtered list of all Session records
func Session_GetListFiltered(filter string) (int, []dm.Session, error) {
	
	tsql := Session_SQLbase
	if filter != "" {
		tsql = tsql + " " + das.WHERE + filter
	}
	count, sessionList, _, _ := session_Fetch(tsql)
	
	return count, sessionList, nil
}



// Session_GetByID() returns a single Session record
func Session_GetByID(id string) (int, dm.Session, error) {


	tsql := Session_SQLbase
	tsql = tsql + " " + das.WHERE + dm.Session_SQLSearchID + das.EQ + das.ID(id)
	_, _, sessionItem, _ := session_Fetch(tsql)

	// START
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, sessionItem, nil
}



// Session_DeleteByID() deletes a single Session record
func Session_Delete(id string) {


// Uses Hard Delete
	object_Table := Session_QualifiedName
	tsql := das.DELETE + das.FROM + object_Table
	tsql = tsql + " " + das.WHERE + dm.Session_SQLSearchID + das.EQ + das.ID(id)
	das.Execute(tsql)	

	
}














	


// Session_Store() saves/stores a Session record to the database
func Session_Store(r dm.Session,req *http.Request) error {

	r, err := Session_Validate(r)
	if err == nil {
		err = session_Save(r, Audit_User(req))
	} else {
		logs.Information("Session_Store()", err.Error())
	}

	return err
}

// Session_StoreSystem() saves/stores a Session record to the database
func Session_StoreSystem(r dm.Session) error {
	
	r, err := Session_Validate(r)
	if err == nil {
		err = session_Save(r, Audit_Host())
	} else {
		logs.Information("Session_Store()", err.Error())
	}

	return err
}

// Session_Validate() validates for saves/stores a Session record to the database
func Session_Validate(r dm.Session) (dm.Session, error) {
	var err error
	// START
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	

	return r,err
}
//

// session_Save() saves/stores a Session record to the database
func session_Save(r dm.Session,usr string) error {

    var err error



	

	if len(r.Id) == 0 {
		r.Id = Session_NewID(r)
	}

// If there are fields below, create the methods in dao\session_impl.go































	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	r.SYSDbVersion = core.DB_Version()
	
logs.Storing("Session",fmt.Sprintf("%v", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.Session_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.Session_Id_sql, r.Id)
	ts = addData(ts, dm.Session_Apptoken_sql, r.Apptoken)
	ts = addData(ts, dm.Session_Createdate_sql, r.Createdate)
	ts = addData(ts, dm.Session_Createtime_sql, r.Createtime)
	ts = addData(ts, dm.Session_Uniqueid_sql, r.Uniqueid)
	ts = addData(ts, dm.Session_Sessiontoken_sql, r.Sessiontoken)
	ts = addData(ts, dm.Session_Username_sql, r.Username)
	ts = addData(ts, dm.Session_Password_sql, r.Password)
	ts = addData(ts, dm.Session_Userip_sql, r.Userip)
	ts = addData(ts, dm.Session_Userhost_sql, r.Userhost)
	ts = addData(ts, dm.Session_Appip_sql, r.Appip)
	ts = addData(ts, dm.Session_Apphost_sql, r.Apphost)
	ts = addData(ts, dm.Session_Issued_sql, r.Issued)
	ts = addData(ts, dm.Session_Expiry_sql, r.Expiry)
	ts = addData(ts, dm.Session_Expiryraw_sql, r.Expiryraw)
	ts = addData(ts, dm.Session_Expires_sql, r.Expires)
	ts = addData(ts, dm.Session_Brand_sql, r.Brand)
	ts = addData(ts, dm.Session_SessionRole_sql, r.SessionRole)
	ts = addData(ts, dm.Session_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.Session_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.Session_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.Session_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.Session_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.Session_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.Session_SYSDeleted_sql, r.SYSDeleted)
	ts = addData(ts, dm.Session_SYSDeletedBy_sql, r.SYSDeletedBy)
	ts = addData(ts, dm.Session_SYSDeletedHost_sql, r.SYSDeletedHost)
	ts = addData(ts, dm.Session_SYSDbVersion_sql, r.SYSDbVersion)
		
	// 
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := das.INSERT + das.INTO + Session_QualifiedName
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " "+das.VALUES +"(" + values(ts) + ")"

	Session_Delete(r.Id)
	das.Execute(tsql)



	return err

}



// session_Fetch read all Session's
func session_Fetch(tsql string) (int, []dm.Session, dm.Session, error) {

	var recItem dm.Session
	var recList []dm.Session

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.Session_SYSId_sql, "0")
	   recItem.Id  = get_String(rec, dm.Session_Id_sql, "")
	   recItem.Apptoken  = get_String(rec, dm.Session_Apptoken_sql, "")
	   recItem.Createdate  = get_String(rec, dm.Session_Createdate_sql, "")
	   recItem.Createtime  = get_String(rec, dm.Session_Createtime_sql, "")
	   recItem.Uniqueid  = get_String(rec, dm.Session_Uniqueid_sql, "")
	   recItem.Sessiontoken  = get_String(rec, dm.Session_Sessiontoken_sql, "")
	   recItem.Username  = get_String(rec, dm.Session_Username_sql, "")
	   recItem.Password  = get_String(rec, dm.Session_Password_sql, "")
	   recItem.Userip  = get_String(rec, dm.Session_Userip_sql, "")
	   recItem.Userhost  = get_String(rec, dm.Session_Userhost_sql, "")
	   recItem.Appip  = get_String(rec, dm.Session_Appip_sql, "")
	   recItem.Apphost  = get_String(rec, dm.Session_Apphost_sql, "")
	   recItem.Issued  = get_String(rec, dm.Session_Issued_sql, "")
	   recItem.Expiry  = get_String(rec, dm.Session_Expiry_sql, "")
	   recItem.Expiryraw  = get_String(rec, dm.Session_Expiryraw_sql, "")
	   recItem.Expires  = get_Time(rec, dm.Session_Expires_sql, "")
	   recItem.Brand  = get_String(rec, dm.Session_Brand_sql, "")
	   recItem.SessionRole  = get_String(rec, dm.Session_SessionRole_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.Session_SYSCreated_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.Session_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.Session_SYSCreatedHost_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.Session_SYSUpdated_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.Session_SYSUpdatedBy_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.Session_SYSUpdatedHost_sql, "")
	   recItem.SYSDeleted  = get_String(rec, dm.Session_SYSDeleted_sql, "")
	   recItem.SYSDeletedBy  = get_String(rec, dm.Session_SYSDeletedBy_sql, "")
	   recItem.SYSDeletedHost  = get_String(rec, dm.Session_SYSDeletedHost_sql, "")
	   recItem.SYSDbVersion  = get_String(rec, dm.Session_SYSDbVersion_sql, "")
	
	// If there are fields below, create the methods in adaptor\Session_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
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
	


func Session_NewID(r dm.Session) string {
	
			id := uuid.New().String()
	
	return id
}



// session_Fetch read all Session's
func Session_New() (int, []dm.Session, dm.Session, error) {

	var r = dm.Session{}
	var rList []dm.Session
	

	// START
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}