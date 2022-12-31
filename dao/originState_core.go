package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/originstate.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : OriginState (originstate)
// Endpoint 	        : OriginState (OriginStateID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 30/12/2022 at 12:18:21
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

var OriginState_SQLbase string
var OriginState_QualifiedName string
func init(){
	OriginState_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.OriginState_SQLTable)
	OriginState_SQLbase =  core.DB_SELECT + " "+ core.DB_ALL + " " + core.DB_FROM + " " + OriginState_QualifiedName
}

// OriginState_GetList() returns a list of all OriginState records
func OriginState_GetList() (int, []dm.OriginState, error) {
	
	tsql := OriginState_SQLbase
	count, originstateList, _, _ := originstate_Fetch(tsql)
	
	return count, originstateList, nil
}


// OriginState_GetLookup() returns a lookup list of all OriginState items in lookup format
func OriginState_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, originstateList, _ := OriginState_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: originstateList[i].Code, Name: originstateList[i].Name})
	}
	return returnList
}


// OriginState_GetByID() returns a single OriginState record
func OriginState_GetByID(id string) (int, dm.OriginState, error) {


	tsql := OriginState_SQLbase
	tsql = tsql + " " + core.DB_WHERE + " " + dm.OriginState_SQLSearchID + core.DB_EQ + "'" + id + "'"
	_, _, originstateItem, _ := originstate_Fetch(tsql)

	// START
	// Dynamically generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, originstateItem, nil
}



// OriginState_DeleteByID() deletes a single OriginState record
func OriginState_Delete(id string) {




// Uses Hard Delete
	object_Table := OriginState_QualifiedName
	tsql := core.DB_DELETE+" "+core.DB_FROM+" " + object_Table
	tsql = tsql + " " + core.DB_WHERE + " " + dm.OriginState_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
	
	

	
}














	
	// OriginState_SoftDeleteByID() soft deletes a single OriginState record
func OriginState_SoftDelete(id string) {
	//Uses Soft Delete
		_,originstateItem,_ := OriginState_GetByID(id)
		originstateItem.SYSDeletedBy = Audit_Update("", Audit_Host())
		originstateItem.SYSDeleted = Audit_Update("", Audit_TimeStamp())
		originstateItem.SYSDeletedHost = Audit_Update("", Audit_Host())
		OriginState_StoreSystem(originstateItem)
}
	


// OriginState_Store() saves/stores a OriginState record to the database
func OriginState_Store(r dm.OriginState,req *http.Request) error {

	r, err := OriginState_Validate(r)
	if err == nil {
		err = originstate_Save(r, Audit_User(req))
	} else {
		logs.Information("OriginState_Store()", err.Error())
	}

	return err
}

// OriginState_StoreSystem() saves/stores a OriginState record to the database
func OriginState_StoreSystem(r dm.OriginState) error {
	
	r, err := OriginState_Validate(r)
	if err == nil {
		err = originstate_Save(r, Audit_Host())
	} else {
		logs.Information("OriginState_Store()", err.Error())
	}

	return err
}

// OriginState_Validate() validates for saves/stores a OriginState record to the database
func OriginState_Validate(r dm.OriginState) (dm.OriginState, error) {
	var err error
	// START
	// Dynamically generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	

	return r,err
}
//

// originstate_Save() saves/stores a OriginState record to the database
func originstate_Save(r dm.OriginState,usr string) error {

    var err error



	

	if len(r.OriginStateID) == 0 {
		r.OriginStateID = OriginState_NewID(r)
	}

// If there are fields below, create the methods in dao\originstate_impl.go
















	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	r.SYSDbVersion = core.DB_Version()
	
logs.Storing("OriginState",fmt.Sprintf("%v", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.OriginState_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.OriginState_OriginStateID_sql, r.OriginStateID)
	ts = addData(ts, dm.OriginState_Code_sql, r.Code)
	ts = addData(ts, dm.OriginState_Name_sql, r.Name)
	ts = addData(ts, dm.OriginState_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.OriginState_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.OriginState_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.OriginState_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.OriginState_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.OriginState_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.OriginState_SYSDeleted_sql, r.SYSDeleted)
	ts = addData(ts, dm.OriginState_SYSDeletedBy_sql, r.SYSDeletedBy)
	ts = addData(ts, dm.OriginState_SYSDeletedHost_sql, r.SYSDeletedHost)
	ts = addData(ts, dm.OriginState_SYSDbVersion_sql, r.SYSDbVersion)
		
	// 
	// Dynamically generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := core.DB_INSERT + " " + core.DB_INTO + " " + OriginState_QualifiedName
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " "+core.DB_VALUES +" (" + values(ts) + ")"

	OriginState_Delete(r.OriginStateID)
	das.Execute(tsql)



	return err

}



// originstate_Fetch read all OriginState's
func originstate_Fetch(tsql string) (int, []dm.OriginState, dm.OriginState, error) {

	var recItem dm.OriginState
	var recList []dm.OriginState

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.OriginState_SYSId_sql, "0")
	   recItem.OriginStateID  = get_String(rec, dm.OriginState_OriginStateID_sql, "")
	   recItem.Code  = get_String(rec, dm.OriginState_Code_sql, "")
	   recItem.Name  = get_String(rec, dm.OriginState_Name_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.OriginState_SYSCreated_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.OriginState_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.OriginState_SYSCreatedHost_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.OriginState_SYSUpdated_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.OriginState_SYSUpdatedBy_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.OriginState_SYSUpdatedHost_sql, "")
	   recItem.SYSDeleted  = get_String(rec, dm.OriginState_SYSDeleted_sql, "")
	   recItem.SYSDeletedBy  = get_String(rec, dm.OriginState_SYSDeletedBy_sql, "")
	   recItem.SYSDeletedHost  = get_String(rec, dm.OriginState_SYSDeletedHost_sql, "")
	   recItem.SYSDbVersion  = get_String(rec, dm.OriginState_SYSDbVersion_sql, "")
	
	// If there are fields below, create the methods in adaptor\OriginState_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func OriginState_NewID(r dm.OriginState) string {
	
			id := uuid.New().String()
	
	return id
}



// originstate_Fetch read all OriginState's
func OriginState_New() (int, []dm.OriginState, dm.OriginState, error) {

	var r = dm.OriginState{}
	var rList []dm.OriginState
	

	// START
	// Dynamically generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}