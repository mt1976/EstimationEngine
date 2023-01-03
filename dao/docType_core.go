package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/doctype.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : DocType (doctype)
// Endpoint 	        : DocType (DocTypeID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 03/01/2023 at 19:18:08
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

var DocType_SQLbase string
var DocType_QualifiedName string
func init(){
	DocType_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.DocType_SQLTable)
	DocType_SQLbase =  core.DB_SELECT + " "+ core.DB_ALL + " " + core.DB_FROM + " " + DocType_QualifiedName
}

// DocType_GetList() returns a list of all DocType records
func DocType_GetList() (int, []dm.DocType, error) {
	
	tsql := DocType_SQLbase
	count, doctypeList, _, _ := doctype_Fetch(tsql)
	
	return count, doctypeList, nil
}


// DocType_GetLookup() returns a lookup list of all DocType items in lookup format
func DocType_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, doctypeList, _ := DocType_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: doctypeList[i].Code, Name: doctypeList[i].Name})
	}
	return returnList
}


// DocType_GetByID() returns a single DocType record
func DocType_GetByID(id string) (int, dm.DocType, error) {


	tsql := DocType_SQLbase
	tsql = tsql + " " + core.DB_WHERE + " " + dm.DocType_SQLSearchID + core.DB_EQ + "'" + id + "'"
	_, _, doctypeItem, _ := doctype_Fetch(tsql)

	// START
	// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, doctypeItem, nil
}



// DocType_DeleteByID() deletes a single DocType record
func DocType_Delete(id string) {




// Uses Hard Delete
	object_Table := DocType_QualifiedName
	tsql := core.DB_DELETE+" "+core.DB_FROM+" " + object_Table
	tsql = tsql + " " + core.DB_WHERE + " " + dm.DocType_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
	
	

	
}














	
	// DocType_SoftDeleteByID() soft deletes a single DocType record
func DocType_SoftDelete(id string) {
	//Uses Soft Delete
		_,doctypeItem,_ := DocType_GetByID(id)
		doctypeItem.SYSDeletedBy = Audit_Update("", Audit_Host())
		doctypeItem.SYSDeleted = Audit_Update("", Audit_TimeStamp())
		doctypeItem.SYSDeletedHost = Audit_Update("", Audit_Host())
		DocType_StoreSystem(doctypeItem)
}
	


// DocType_Store() saves/stores a DocType record to the database
func DocType_Store(r dm.DocType,req *http.Request) error {

	r, err := DocType_Validate(r)
	if err == nil {
		err = doctype_Save(r, Audit_User(req))
	} else {
		logs.Information("DocType_Store()", err.Error())
	}

	return err
}

// DocType_StoreSystem() saves/stores a DocType record to the database
func DocType_StoreSystem(r dm.DocType) error {
	
	r, err := DocType_Validate(r)
	if err == nil {
		err = doctype_Save(r, Audit_Host())
	} else {
		logs.Information("DocType_Store()", err.Error())
	}

	return err
}

// DocType_Validate() validates for saves/stores a DocType record to the database
func DocType_Validate(r dm.DocType) (dm.DocType, error) {
	var err error
	// START
	// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	

	return r,err
}
//

// doctype_Save() saves/stores a DocType record to the database
func doctype_Save(r dm.DocType,usr string) error {

    var err error



	

	if len(r.DocTypeID) == 0 {
		r.DocTypeID = DocType_NewID(r)
	}

// If there are fields below, create the methods in dao\doctype_impl.go

















	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	r.SYSDbVersion = core.DB_Version()
	
logs.Storing("DocType",fmt.Sprintf("%v", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.DocType_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.DocType_DocTypeID_sql, r.DocTypeID)
	ts = addData(ts, dm.DocType_Code_sql, r.Code)
	ts = addData(ts, dm.DocType_Name_sql, r.Name)
	ts = addData(ts, dm.DocType_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.DocType_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.DocType_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.DocType_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.DocType_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.DocType_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.DocType_SYSDeleted_sql, r.SYSDeleted)
	ts = addData(ts, dm.DocType_SYSDeletedBy_sql, r.SYSDeletedBy)
	ts = addData(ts, dm.DocType_SYSDeletedHost_sql, r.SYSDeletedHost)
	ts = addData(ts, dm.DocType_SYSDbVersion_sql, r.SYSDbVersion)
	ts = addData(ts, dm.DocType_Comments_sql, r.Comments)
		
	// 
	// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := core.DB_INSERT + " " + core.DB_INTO + " " + DocType_QualifiedName
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " "+core.DB_VALUES +" (" + values(ts) + ")"

	DocType_Delete(r.DocTypeID)
	das.Execute(tsql)



	return err

}



// doctype_Fetch read all DocType's
func doctype_Fetch(tsql string) (int, []dm.DocType, dm.DocType, error) {

	var recItem dm.DocType
	var recList []dm.DocType

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.DocType_SYSId_sql, "0")
	   recItem.DocTypeID  = get_String(rec, dm.DocType_DocTypeID_sql, "")
	   recItem.Code  = get_String(rec, dm.DocType_Code_sql, "")
	   recItem.Name  = get_String(rec, dm.DocType_Name_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.DocType_SYSCreated_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.DocType_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.DocType_SYSCreatedHost_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.DocType_SYSUpdated_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.DocType_SYSUpdatedBy_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.DocType_SYSUpdatedHost_sql, "")
	   recItem.SYSDeleted  = get_String(rec, dm.DocType_SYSDeleted_sql, "")
	   recItem.SYSDeletedBy  = get_String(rec, dm.DocType_SYSDeletedBy_sql, "")
	   recItem.SYSDeletedHost  = get_String(rec, dm.DocType_SYSDeletedHost_sql, "")
	   recItem.SYSDbVersion  = get_String(rec, dm.DocType_SYSDbVersion_sql, "")
	   recItem.Comments  = get_String(rec, dm.DocType_Comments_sql, "")
	
	// If there are fields below, create the methods in adaptor\DocType_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func DocType_NewID(r dm.DocType) string {
	
			id := uuid.New().String()
	
	return id
}



// doctype_Fetch read all DocType's
func DocType_New() (int, []dm.DocType, dm.DocType, error) {

	var r = dm.DocType{}
	var rList []dm.DocType
	

	// START
	// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}