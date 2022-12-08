package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/confidence.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Confidence (confidence)
// Endpoint 	        : Confidence (ConfidenceID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 08/12/2022 at 13:31:29
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

// Confidence_GetList() returns a list of all Confidence records
func Confidence_GetList() (int, []dm.Confidence, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Confidence_SQLTable)
	count, confidenceList, _, _ := confidence_Fetch(tsql)
	
	return count, confidenceList, nil
}


// Confidence_GetLookup() returns a lookup list of all Confidence items in lookup format
func Confidence_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, confidenceList, _ := Confidence_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: confidenceList[i].Code, Name: confidenceList[i].Name})
	}
	return returnList
}


// Confidence_GetByID() returns a single Confidence record
func Confidence_GetByID(id string) (int, dm.Confidence, error) {


	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Confidence_SQLTable)
	tsql = tsql + " WHERE " + dm.Confidence_SQLSearchID + "='" + id + "'"
	_, _, confidenceItem, _ := confidence_Fetch(tsql)

	// START
	// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, confidenceItem, nil
}



// Confidence_DeleteByID() deletes a single Confidence record
func Confidence_Delete(id string) {




// Uses Hard Delete
	object_Table := core.GetSQLSchema(core.ApplicationPropertiesDB) + "." + dm.Confidence_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Confidence_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
	
	

	
}














	
	// Confidence_SoftDeleteByID() soft deletes a single Confidence record
func Confidence_SoftDelete(id string) {
	//Uses Soft Delete
		_,confidenceItem,_ := Confidence_GetByID(id)
		confidenceItem.SYSDeletedBy = Audit_Update("", Audit_Host())
		confidenceItem.SYSDeleted = Audit_Update("", Audit_TimeStamp())
		confidenceItem.SYSDeletedHost = Audit_Update("", Audit_Host())
		Confidence_StoreSystem(confidenceItem)
}
	


// Confidence_Store() saves/stores a Confidence record to the database
func Confidence_Store(r dm.Confidence,req *http.Request) error {

	r, err := Confidence_Validate(r)
	if err == nil {
		err = confidence_Save(r, Audit_User(req))
	} else {
		logs.Information("Confidence_Store()", err.Error())
	}

	return err
}

// Confidence_StoreSystem() saves/stores a Confidence record to the database
func Confidence_StoreSystem(r dm.Confidence) error {
	
	r, err := Confidence_Validate(r)
	if err == nil {
		err = confidence_Save(r, Audit_Host())
	} else {
		logs.Information("Confidence_Store()", err.Error())
	}

	return err
}

// Confidence_Validate() validates for saves/stores a Confidence record to the database
func Confidence_Validate(r dm.Confidence) (dm.Confidence, error) {
	var err error
	// START
	// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	

	return r,err
}
//

// confidence_Save() saves/stores a Confidence record to the database
func confidence_Save(r dm.Confidence,usr string) error {

    var err error



	

	if len(r.ConfidenceID) == 0 {
		r.ConfidenceID = Confidence_NewID(r)
	}

// If there are fields below, create the methods in dao\confidence_impl.go

















	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	
logs.Storing("Confidence",fmt.Sprintf("%v", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.Confidence_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.Confidence_ConfidenceID_sql, r.ConfidenceID)
	ts = addData(ts, dm.Confidence_Code_sql, r.Code)
	ts = addData(ts, dm.Confidence_Name_sql, r.Name)
	ts = addData(ts, dm.Confidence_Perc_sql, r.Perc)
	ts = addData(ts, dm.Confidence_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.Confidence_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.Confidence_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.Confidence_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.Confidence_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.Confidence_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.Confidence_SYSDeleted_sql, r.SYSDeleted)
	ts = addData(ts, dm.Confidence_SYSDeletedBy_sql, r.SYSDeletedBy)
	ts = addData(ts, dm.Confidence_SYSDeletedHost_sql, r.SYSDeletedHost)
	ts = addData(ts, dm.Confidence_Notes_sql, r.Notes)
		
	// 
	// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := "INSERT INTO " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Confidence_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Confidence_Delete(r.ConfidenceID)
	das.Execute(tsql)



	return err

}



// confidence_Fetch read all Confidence's
func confidence_Fetch(tsql string) (int, []dm.Confidence, dm.Confidence, error) {

	var recItem dm.Confidence
	var recList []dm.Confidence

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.Confidence_SYSId_sql, "0")
	   recItem.ConfidenceID  = get_String(rec, dm.Confidence_ConfidenceID_sql, "")
	   recItem.Code  = get_String(rec, dm.Confidence_Code_sql, "")
	   recItem.Name  = get_String(rec, dm.Confidence_Name_sql, "")
	   recItem.Perc  = get_String(rec, dm.Confidence_Perc_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.Confidence_SYSCreated_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.Confidence_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.Confidence_SYSCreatedHost_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.Confidence_SYSUpdated_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.Confidence_SYSUpdatedBy_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.Confidence_SYSUpdatedHost_sql, "")
	   recItem.SYSDeleted  = get_String(rec, dm.Confidence_SYSDeleted_sql, "")
	   recItem.SYSDeletedBy  = get_String(rec, dm.Confidence_SYSDeletedBy_sql, "")
	   recItem.SYSDeletedHost  = get_String(rec, dm.Confidence_SYSDeletedHost_sql, "")
	   recItem.Notes  = get_String(rec, dm.Confidence_Notes_sql, "")
	
	// If there are fields below, create the methods in adaptor\Confidence_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func Confidence_NewID(r dm.Confidence) string {
	
			id := uuid.New().String()
	
	return id
}



// confidence_Fetch read all Confidence's
func Confidence_New() (int, []dm.Confidence, dm.Confidence, error) {

	var r = dm.Confidence{}
	var rList []dm.Confidence
	

	// START
	// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}