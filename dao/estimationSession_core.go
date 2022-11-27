package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/estimationsession.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : EstimationSession (estimationsession)
// Endpoint 	        : EstimationSession (EstimationSessionID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 27/11/2022 at 20:46:13
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

// EstimationSession_GetList() returns a list of all EstimationSession records
func EstimationSession_GetList() (int, []dm.EstimationSession, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.EstimationSession_SQLTable)
	count, estimationsessionList, _, _ := estimationsession_Fetch(tsql)
	
	return count, estimationsessionList, nil
}


// EstimationSession_GetLookup() returns a lookup list of all EstimationSession items in lookup format
func EstimationSession_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, estimationsessionList, _ := EstimationSession_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: estimationsessionList[i].EstimationSessionID, Name: estimationsessionList[i].Name})
	}
	return returnList
}


// EstimationSession_GetByID() returns a single EstimationSession record
func EstimationSession_GetByID(id string) (int, dm.EstimationSession, error) {


	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.EstimationSession_SQLTable)
	tsql = tsql + " WHERE " + dm.EstimationSession_SQLSearchID + "='" + id + "'"
	_, _, estimationsessionItem, _ := estimationsession_Fetch(tsql)

	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, estimationsessionItem, nil
}



// EstimationSession_DeleteByID() deletes a single EstimationSession record
func EstimationSession_Delete(id string) {




// Uses Hard Delete
	object_Table := core.GetSQLSchema(core.ApplicationPropertiesDB) + "." + dm.EstimationSession_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.EstimationSession_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
	
	

	
}














	
	// EstimationSession_SoftDeleteByID() soft deletes a single EstimationSession record
func EstimationSession_SoftDelete(id string) {
	//Uses Soft Delete
		_,estimationsessionItem,_ := EstimationSession_GetByID(id)
		estimationsessionItem.SYSDeletedBy = Audit_Update("", Audit_Host())
		estimationsessionItem.SYSDeleted = Audit_Update("", Audit_TimeStamp())
		estimationsessionItem.SYSDeletedHost = Audit_Update("", Audit_Host())
		EstimationSession_StoreSystem(estimationsessionItem)
}
	


// EstimationSession_Store() saves/stores a EstimationSession record to the database
func EstimationSession_Store(r dm.EstimationSession,req *http.Request) error {

	err, r := EstimationSession_Validate(r)
	if err == nil {
		err = estimationsession_Save(r, Audit_User(req))
	} else {
		logs.Information("EstimationSession_Store()", err.Error())
	}

	return err
}

// EstimationSession_StoreSystem() saves/stores a EstimationSession record to the database
func EstimationSession_StoreSystem(r dm.EstimationSession) error {
	
	err, r := EstimationSession_Validate(r)
	if err == nil {
		err = estimationsession_Save(r, Audit_Host())
	} else {
		logs.Information("EstimationSession_Store()", err.Error())
	}

	return err
}

// EstimationSession_Validate() validates for saves/stores a EstimationSession record to the database
func EstimationSession_Validate(r dm.EstimationSession) (error,dm.EstimationSession) {
	var err error
	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	

	return err,r
}
//

// estimationsession_Save() saves/stores a EstimationSession record to the database
func estimationsession_Save(r dm.EstimationSession,usr string) error {

    var err error



	

	if len(r.EstimationSessionID) == 0 {
		r.EstimationSessionID = EstimationSession_NewID(r)
	}

// If there are fields below, create the methods in dao\estimationsession_impl.go































	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	
logs.Storing("EstimationSession",fmt.Sprintf("%s", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.EstimationSession_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.EstimationSession_EstimationSessionID_sql, r.EstimationSessionID)
	ts = addData(ts, dm.EstimationSession_ProjectID_sql, r.ProjectID)
	ts = addData(ts, dm.EstimationSession_EstimationStateID_sql, r.EstimationStateID)
	ts = addData(ts, dm.EstimationSession_Notes_sql, r.Notes)
	ts = addData(ts, dm.EstimationSession_Releases_sql, r.Releases)
	ts = addData(ts, dm.EstimationSession_Total_sql, r.Total)
	ts = addData(ts, dm.EstimationSession_Contingency_sql, r.Contingency)
	ts = addData(ts, dm.EstimationSession_ReqDays_sql, r.ReqDays)
	ts = addData(ts, dm.EstimationSession_RegCost_sql, r.RegCost)
	ts = addData(ts, dm.EstimationSession_ImpDays_sql, r.ImpDays)
	ts = addData(ts, dm.EstimationSession_ImpCost_sql, r.ImpCost)
	ts = addData(ts, dm.EstimationSession_UatDays_sql, r.UatDays)
	ts = addData(ts, dm.EstimationSession_UatCost_sql, r.UatCost)
	ts = addData(ts, dm.EstimationSession_MgtDays_sql, r.MgtDays)
	ts = addData(ts, dm.EstimationSession_MgtCost_sql, r.MgtCost)
	ts = addData(ts, dm.EstimationSession_RelDays_sql, r.RelDays)
	ts = addData(ts, dm.EstimationSession_RelCost_sql, r.RelCost)
	ts = addData(ts, dm.EstimationSession_SupportUplift_sql, r.SupportUplift)
	ts = addData(ts, dm.EstimationSession_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.EstimationSession_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.EstimationSession_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.EstimationSession_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.EstimationSession_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.EstimationSession_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.EstimationSession_SYSDeleted_sql, r.SYSDeleted)
	ts = addData(ts, dm.EstimationSession_SYSDeletedBy_sql, r.SYSDeletedBy)
	ts = addData(ts, dm.EstimationSession_SYSDeletedHost_sql, r.SYSDeletedHost)
	ts = addData(ts, dm.EstimationSession_Name_sql, r.Name)
		
	// 
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := "INSERT INTO " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.EstimationSession_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	EstimationSession_Delete(r.EstimationSessionID)
	das.Execute(tsql)



	return err

}



// estimationsession_Fetch read all EstimationSession's
func estimationsession_Fetch(tsql string) (int, []dm.EstimationSession, dm.EstimationSession, error) {

	var recItem dm.EstimationSession
	var recList []dm.EstimationSession

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.EstimationSession_SYSId_sql, "0")
	   recItem.EstimationSessionID  = get_String(rec, dm.EstimationSession_EstimationSessionID_sql, "")
	   recItem.ProjectID  = get_String(rec, dm.EstimationSession_ProjectID_sql, "")
	   recItem.EstimationStateID  = get_String(rec, dm.EstimationSession_EstimationStateID_sql, "")
	   recItem.Notes  = get_String(rec, dm.EstimationSession_Notes_sql, "")
	   recItem.Releases  = get_String(rec, dm.EstimationSession_Releases_sql, "")
	   recItem.Total  = get_String(rec, dm.EstimationSession_Total_sql, "")
	   recItem.Contingency  = get_String(rec, dm.EstimationSession_Contingency_sql, "")
	   recItem.ReqDays  = get_String(rec, dm.EstimationSession_ReqDays_sql, "")
	   recItem.RegCost  = get_String(rec, dm.EstimationSession_RegCost_sql, "")
	   recItem.ImpDays  = get_String(rec, dm.EstimationSession_ImpDays_sql, "")
	   recItem.ImpCost  = get_String(rec, dm.EstimationSession_ImpCost_sql, "")
	   recItem.UatDays  = get_String(rec, dm.EstimationSession_UatDays_sql, "")
	   recItem.UatCost  = get_String(rec, dm.EstimationSession_UatCost_sql, "")
	   recItem.MgtDays  = get_String(rec, dm.EstimationSession_MgtDays_sql, "")
	   recItem.MgtCost  = get_String(rec, dm.EstimationSession_MgtCost_sql, "")
	   recItem.RelDays  = get_String(rec, dm.EstimationSession_RelDays_sql, "")
	   recItem.RelCost  = get_String(rec, dm.EstimationSession_RelCost_sql, "")
	   recItem.SupportUplift  = get_String(rec, dm.EstimationSession_SupportUplift_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.EstimationSession_SYSCreated_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.EstimationSession_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.EstimationSession_SYSCreatedHost_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.EstimationSession_SYSUpdated_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.EstimationSession_SYSUpdatedBy_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.EstimationSession_SYSUpdatedHost_sql, "")
	   recItem.SYSDeleted  = get_String(rec, dm.EstimationSession_SYSDeleted_sql, "")
	   recItem.SYSDeletedBy  = get_String(rec, dm.EstimationSession_SYSDeletedBy_sql, "")
	   recItem.SYSDeletedHost  = get_String(rec, dm.EstimationSession_SYSDeletedHost_sql, "")
	   recItem.Name  = get_String(rec, dm.EstimationSession_Name_sql, "")
	
	// If there are fields below, create the methods in adaptor\EstimationSession_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func EstimationSession_NewID(r dm.EstimationSession) string {
	
			id := uuid.New().String()
	
	return id
}



// estimationsession_Fetch read all EstimationSession's
func EstimationSession_New() (int, []dm.EstimationSession, dm.EstimationSession, error) {

	var r = dm.EstimationSession{}
	var rList []dm.EstimationSession
	

	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}