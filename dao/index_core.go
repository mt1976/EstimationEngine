package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/index.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Index (index)
// Endpoint 	        : Index (IndexID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 24/01/2023 at 13:18:09
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

var Index_SQLbase string
var Index_QualifiedName string
func init(){
	Index_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.Index_SQLTable)
	Index_SQLbase =  core.DB_SELECT + " "+ core.DB_ALL + " " + core.DB_FROM + " " + Index_QualifiedName
}

// Index_GetList() returns a list of all Index records
func Index_GetList() (int, []dm.Index, error) {
	
	count, indexList, err := Index_GetListFiltered("")
	
	return count, indexList, err
}

// Index_GetListFiltered() returns a filtered list of all Index records
func Index_GetListFiltered(filter string) (int, []dm.Index, error) {
	
	tsql := Index_SQLbase
	if filter != "" {
		tsql = tsql + " " + core.DB_WHERE + " " + filter
	}
	count, indexList, _, _ := index_Fetch(tsql)
	
	return count, indexList, nil
}



// Index_GetByID() returns a single Index record
func Index_GetByID(id string) (int, dm.Index, error) {


	tsql := Index_SQLbase
	tsql = tsql + " " + core.DB_WHERE + " " + dm.Index_SQLSearchID + core.DB_EQ + "'" + id + "'"
	_, _, indexItem, _ := index_Fetch(tsql)

	// START
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	indexItem.IndexID,indexItem.IndexID_props = Index_IndexID_impl (GET,id,indexItem.IndexID,indexItem,indexItem.IndexID_props)
	indexItem.Link,indexItem.Link_props = Index_Link_impl (GET,id,indexItem.Link,indexItem,indexItem.Link_props)
	indexItem.LinkView,indexItem.LinkView_props = Index_LinkView_impl (GET,id,indexItem.LinkView,indexItem,indexItem.LinkView_props)
	indexItem.LinkEdit,indexItem.LinkEdit_props = Index_LinkEdit_impl (GET,id,indexItem.LinkEdit,indexItem,indexItem.LinkEdit_props)
	// 
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, indexItem, nil
}



// Index_DeleteByID() deletes a single Index record
func Index_Delete(id string) {




// Uses Hard Delete
	object_Table := Index_QualifiedName
	tsql := core.DB_DELETE+" "+core.DB_FROM+" " + object_Table
	tsql = tsql + " " + core.DB_WHERE + " " + dm.Index_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
	
	

	
}














	


// Index_Store() saves/stores a Index record to the database
func Index_Store(r dm.Index,req *http.Request) error {

	r, err := Index_Validate(r)
	if err == nil {
		err = index_Save(r, Audit_User(req))
	} else {
		logs.Information("Index_Store()", err.Error())
	}

	return err
}

// Index_StoreSystem() saves/stores a Index record to the database
func Index_StoreSystem(r dm.Index) error {
	
	r, err := Index_Validate(r)
	if err == nil {
		err = index_Save(r, Audit_Host())
	} else {
		logs.Information("Index_Store()", err.Error())
	}

	return err
}

// Index_Validate() validates for saves/stores a Index record to the database
func Index_Validate(r dm.Index) (dm.Index, error) {
	var err error
	// START
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	r.IndexID,r.IndexID_props = Index_IndexID_impl (PUT,r.IndexID,r.IndexID,r,r.IndexID_props)
	r.Link,r.Link_props = Index_Link_impl (PUT,r.IndexID,r.Link,r,r.Link_props)
	r.LinkView,r.LinkView_props = Index_LinkView_impl (PUT,r.IndexID,r.LinkView,r,r.LinkView_props)
	r.LinkEdit,r.LinkEdit_props = Index_LinkEdit_impl (PUT,r.IndexID,r.LinkEdit,r,r.LinkEdit_props)
	// 
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	
	// START
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	r, _, err = Index_ObjectValidation_impl(PUT, r.IndexID, r)
	// 
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	

	return r,err
}
//

// index_Save() saves/stores a Index record to the database
func index_Save(r dm.Index,usr string) error {

    var err error



	

	if len(r.IndexID) == 0 {
		r.IndexID = Index_NewID(r)
	}

// If there are fields below, create the methods in dao\index_impl.go

  r.IndexID,err = Index_IndexID_OnStore_impl (r.IndexID,r,usr)



  r.Link,err = Index_Link_OnStore_impl (r.Link,r,usr)
  r.LinkView,err = Index_LinkView_OnStore_impl (r.LinkView,r,usr)
  r.LinkEdit,err = Index_LinkEdit_OnStore_impl (r.LinkEdit,r,usr)













	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	r.SYSDbVersion = core.DB_Version()
	
logs.Storing("Index",fmt.Sprintf("%v", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.Index_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.Index_IndexID_sql, r.IndexID)
	ts = addData(ts, dm.Index_KeyClass_sql, r.KeyClass)
	ts = addData(ts, dm.Index_KeyName_sql, r.KeyName)
	ts = addData(ts, dm.Index_KeyID_sql, r.KeyID)
	ts = addData(ts, dm.Index_Link_sql, r.Link)
	ts = addData(ts, dm.Index_LinkView_sql, r.LinkView)
	ts = addData(ts, dm.Index_LinkEdit_sql, r.LinkEdit)
	ts = addData(ts, dm.Index_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.Index_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.Index_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.Index_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.Index_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.Index_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.Index_SYSDeleted_sql, r.SYSDeleted)
	ts = addData(ts, dm.Index_SYSDeletedBy_sql, r.SYSDeletedBy)
	ts = addData(ts, dm.Index_SYSDeletedHost_sql, r.SYSDeletedHost)
	ts = addData(ts, dm.Index_SYSDbVersion_sql, r.SYSDbVersion)
	ts = addData(ts, dm.Index_KeyValue_sql, r.KeyValue)
		
	// 
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := core.DB_INSERT + " " + core.DB_INTO + " " + Index_QualifiedName
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " "+core.DB_VALUES +" (" + values(ts) + ")"

	Index_Delete(r.IndexID)
	das.Execute(tsql)



	return err

}



// index_Fetch read all Index's
func index_Fetch(tsql string) (int, []dm.Index, dm.Index, error) {

	var recItem dm.Index
	var recList []dm.Index

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.Index_SYSId_sql, "0")
	   recItem.IndexID  = get_String(rec, dm.Index_IndexID_sql, "")
	   recItem.KeyClass  = get_String(rec, dm.Index_KeyClass_sql, "")
	   recItem.KeyName  = get_String(rec, dm.Index_KeyName_sql, "")
	   recItem.KeyID  = get_String(rec, dm.Index_KeyID_sql, "")
	   recItem.Link  = get_String(rec, dm.Index_Link_sql, "")
	   recItem.LinkView  = get_String(rec, dm.Index_LinkView_sql, "")
	   recItem.LinkEdit  = get_String(rec, dm.Index_LinkEdit_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.Index_SYSCreated_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.Index_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.Index_SYSCreatedHost_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.Index_SYSUpdated_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.Index_SYSUpdatedBy_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.Index_SYSUpdatedHost_sql, "")
	   recItem.SYSDeleted  = get_String(rec, dm.Index_SYSDeleted_sql, "")
	   recItem.SYSDeletedBy  = get_String(rec, dm.Index_SYSDeletedBy_sql, "")
	   recItem.SYSDeletedHost  = get_String(rec, dm.Index_SYSDeletedHost_sql, "")
	   recItem.SYSDbVersion  = get_String(rec, dm.Index_SYSDbVersion_sql, "")
	   recItem.KeyValue  = get_String(rec, dm.Index_KeyValue_sql, "")
	
	// If there are fields below, create the methods in adaptor\Index_impl.go
	
	   recItem.IndexID  = Index_IndexID_OnFetch_impl (recItem)
	
	
	
	   recItem.Link  = Index_Link_OnFetch_impl (recItem)
	   recItem.LinkView  = Index_LinkView_OnFetch_impl (recItem)
	   recItem.LinkEdit  = Index_LinkEdit_OnFetch_impl (recItem)
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func Index_NewID(r dm.Index) string {
	
			id := uuid.New().String()
	
	return id
}



// index_Fetch read all Index's
func Index_New() (int, []dm.Index, dm.Index, error) {

	var r = dm.Index{}
	var rList []dm.Index
	

	// START
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	r.IndexID,r.IndexID_props = Index_IndexID_impl (NEW,r.IndexID,r.IndexID,r,r.IndexID_props)
	r.Link,r.Link_props = Index_Link_impl (NEW,r.IndexID,r.Link,r,r.Link_props)
	r.LinkView,r.LinkView_props = Index_LinkView_impl (NEW,r.IndexID,r.LinkView,r,r.LinkView_props)
	r.LinkEdit,r.LinkEdit_props = Index_LinkEdit_impl (NEW,r.IndexID,r.LinkEdit,r,r.LinkEdit_props)
	// 
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}