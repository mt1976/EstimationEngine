package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/data.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Data (data)
// Endpoint 	        : Data (DataID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 05/01/2023 at 20:43:15
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

var Data_SQLbase string
var Data_QualifiedName string
func init(){
	Data_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.Data_SQLTable)
	Data_SQLbase =  core.DB_SELECT + " "+ core.DB_ALL + " " + core.DB_FROM + " " + Data_QualifiedName
}

// Data_GetList() returns a list of all Data records
func Data_GetList() (int, []dm.Data, error) {
	
	tsql := Data_SQLbase
	count, dataList, _, _ := data_Fetch(tsql)
	
	return count, dataList, nil
}



// Data_GetByID() returns a single Data record
func Data_GetByID(id string) (int, dm.Data, error) {


	tsql := Data_SQLbase
	tsql = tsql + " " + core.DB_WHERE + " " + dm.Data_SQLSearchID + core.DB_EQ + "'" + id + "'"
	_, _, dataItem, _ := data_Fetch(tsql)

	// START
	// Dynamically generated 05/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	dataItem.DataID,dataItem.DataID_props = Data_DataID_impl (GET,id,dataItem.DataID,dataItem,dataItem.DataID_props)
	// 
	// Dynamically generated 05/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, dataItem, nil
}



// Data_DeleteByID() deletes a single Data record
func Data_Delete(id string) {




// Uses Hard Delete
	object_Table := Data_QualifiedName
	tsql := core.DB_DELETE+" "+core.DB_FROM+" " + object_Table
	tsql = tsql + " " + core.DB_WHERE + " " + dm.Data_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
	
	

	
}














	
	// Data_SoftDeleteByID() soft deletes a single Data record
func Data_SoftDelete(id string) {
	//Uses Soft Delete
		_,dataItem,_ := Data_GetByID(id)
		dataItem.SYSDeletedBy = Audit_Update("", Audit_Host())
		dataItem.SYSDeleted = Audit_Update("", Audit_TimeStamp())
		dataItem.SYSDeletedHost = Audit_Update("", Audit_Host())
		Data_StoreSystem(dataItem)
}
	


// Data_Store() saves/stores a Data record to the database
func Data_Store(r dm.Data,req *http.Request) error {

	r, err := Data_Validate(r)
	if err == nil {
		err = data_Save(r, Audit_User(req))
	} else {
		logs.Information("Data_Store()", err.Error())
	}

	return err
}

// Data_StoreSystem() saves/stores a Data record to the database
func Data_StoreSystem(r dm.Data) error {
	
	r, err := Data_Validate(r)
	if err == nil {
		err = data_Save(r, Audit_Host())
	} else {
		logs.Information("Data_Store()", err.Error())
	}

	return err
}

// Data_Validate() validates for saves/stores a Data record to the database
func Data_Validate(r dm.Data) (dm.Data, error) {
	var err error
	// START
	// Dynamically generated 05/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	r.DataID,r.DataID_props = Data_DataID_impl (PUT,r.DataID,r.DataID,r,r.DataID_props)
	// 
	// Dynamically generated 05/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	
	// START
	// Dynamically generated 05/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	r, _, err = Data_ObjectValidation_impl(PUT, r.DataID, r)
	// 
	// Dynamically generated 05/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	

	return r,err
}
//

// data_Save() saves/stores a Data record to the database
func data_Save(r dm.Data,usr string) error {

    var err error



	

	if len(r.DataID) == 0 {
		r.DataID = Data_NewID(r)
	}

// If there are fields below, create the methods in dao\data_impl.go

  r.DataID,err = Data_DataID_OnStore_impl (r.DataID,r,usr)















	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	r.SYSDbVersion = core.DB_Version()
	
logs.Storing("Data",fmt.Sprintf("%v", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 05/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.Data_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.Data_DataID_sql, r.DataID)
	ts = addData(ts, dm.Data_Class_sql, r.Class)
	ts = addData(ts, dm.Data_Field_sql, r.Field)
	ts = addData(ts, dm.Data_Value_sql, r.Value)
	ts = addData(ts, dm.Data_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.Data_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.Data_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.Data_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.Data_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.Data_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.Data_SYSDeleted_sql, r.SYSDeleted)
	ts = addData(ts, dm.Data_SYSDeletedBy_sql, r.SYSDeletedBy)
	ts = addData(ts, dm.Data_SYSDeletedHost_sql, r.SYSDeletedHost)
	ts = addData(ts, dm.Data_SYSDbVersion_sql, r.SYSDbVersion)
		
	// 
	// Dynamically generated 05/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := core.DB_INSERT + " " + core.DB_INTO + " " + Data_QualifiedName
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " "+core.DB_VALUES +" (" + values(ts) + ")"

	Data_Delete(r.DataID)
	das.Execute(tsql)



	return err

}



// data_Fetch read all Data's
func data_Fetch(tsql string) (int, []dm.Data, dm.Data, error) {

	var recItem dm.Data
	var recList []dm.Data

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 05/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.Data_SYSId_sql, "0")
	   recItem.DataID  = get_String(rec, dm.Data_DataID_sql, "")
	   recItem.Class  = get_String(rec, dm.Data_Class_sql, "")
	   recItem.Field  = get_String(rec, dm.Data_Field_sql, "")
	   recItem.Value  = get_String(rec, dm.Data_Value_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.Data_SYSCreated_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.Data_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.Data_SYSCreatedHost_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.Data_SYSUpdated_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.Data_SYSUpdatedBy_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.Data_SYSUpdatedHost_sql, "")
	   recItem.SYSDeleted  = get_String(rec, dm.Data_SYSDeleted_sql, "")
	   recItem.SYSDeletedBy  = get_String(rec, dm.Data_SYSDeletedBy_sql, "")
	   recItem.SYSDeletedHost  = get_String(rec, dm.Data_SYSDeletedHost_sql, "")
	   recItem.SYSDbVersion  = get_String(rec, dm.Data_SYSDbVersion_sql, "")
	
	// If there are fields below, create the methods in adaptor\Data_impl.go
	
	   recItem.DataID  = Data_DataID_OnFetch_impl (recItem)
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 05/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func Data_NewID(r dm.Data) string {
	
			id := uuid.New().String()
	
	return id
}



// data_Fetch read all Data's
func Data_New() (int, []dm.Data, dm.Data, error) {

	var r = dm.Data{}
	var rList []dm.Data
	

	// START
	// Dynamically generated 05/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	r.DataID,r.DataID_props = Data_DataID_impl (NEW,r.DataID,r.DataID,r,r.DataID_props)
	// 
	// Dynamically generated 05/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}