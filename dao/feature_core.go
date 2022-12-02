package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/feature.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Feature (feature)
// Endpoint 	        : Feature (FeatureID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 01/12/2022 at 09:40:01
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

// Feature_GetList() returns a list of all Feature records
func Feature_GetList() (int, []dm.Feature, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Feature_SQLTable)
	count, featureList, _, _ := feature_Fetch(tsql)
	
	return count, featureList, nil
}


// Feature_GetLookup() returns a lookup list of all Feature items in lookup format
func Feature_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, featureList, _ := Feature_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: featureList[i].FeatureID, Name: featureList[i].Name})
	}
	return returnList
}


// Feature_GetByID() returns a single Feature record
func Feature_GetByID(id string) (int, dm.Feature, error) {


	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Feature_SQLTable)
	tsql = tsql + " WHERE " + dm.Feature_SQLSearchID + "='" + id + "'"
	_, _, featureItem, _ := feature_Fetch(tsql)

	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, featureItem, nil
}



// Feature_DeleteByID() deletes a single Feature record
func Feature_Delete(id string) {




// Uses Hard Delete
	object_Table := core.GetSQLSchema(core.ApplicationPropertiesDB) + "." + dm.Feature_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Feature_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
	
	

	
}














	
	// Feature_SoftDeleteByID() soft deletes a single Feature record
func Feature_SoftDelete(id string) {
	//Uses Soft Delete
		_,featureItem,_ := Feature_GetByID(id)
		featureItem.SYSDeletedBy = Audit_Update("", Audit_Host())
		featureItem.SYSDeleted = Audit_Update("", Audit_TimeStamp())
		featureItem.SYSDeletedHost = Audit_Update("", Audit_Host())
		Feature_StoreSystem(featureItem)
}
	


// Feature_Store() saves/stores a Feature record to the database
func Feature_Store(r dm.Feature,req *http.Request) error {

	err, r := Feature_Validate(r)
	if err == nil {
		err = feature_Save(r, Audit_User(req))
	} else {
		logs.Information("Feature_Store()", err.Error())
	}

	return err
}

// Feature_StoreSystem() saves/stores a Feature record to the database
func Feature_StoreSystem(r dm.Feature) error {
	
	err, r := Feature_Validate(r)
	if err == nil {
		err = feature_Save(r, Audit_Host())
	} else {
		logs.Information("Feature_Store()", err.Error())
	}

	return err
}

// Feature_Validate() validates for saves/stores a Feature record to the database
func Feature_Validate(r dm.Feature) (error,dm.Feature) {
	var err error
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	

	return err,r
}
//

// feature_Save() saves/stores a Feature record to the database
func feature_Save(r dm.Feature,usr string) error {

    var err error



	

	if len(r.FeatureID) == 0 {
		r.FeatureID = Feature_NewID(r)
	}

// If there are fields below, create the methods in dao\feature_impl.go




































	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	
logs.Storing("Feature",fmt.Sprintf("%s", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.Feature_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.Feature_FeatureID_sql, r.FeatureID)
	ts = addData(ts, dm.Feature_EstimationSessionID_sql, r.EstimationSessionID)
	ts = addData(ts, dm.Feature_ConfidenceID_sql, r.ConfidenceID)
	ts = addData(ts, dm.Feature_Name_sql, r.Name)
	ts = addData(ts, dm.Feature_DevEstimate_sql, r.DevEstimate)
	ts = addData(ts, dm.Feature_DevUplift_sql, r.DevUplift)
	ts = addData(ts, dm.Feature_Reqs_sql, r.Reqs)
	ts = addData(ts, dm.Feature_AnalystTest_sql, r.AnalystTest)
	ts = addData(ts, dm.Feature_Docs_sql, r.Docs)
	ts = addData(ts, dm.Feature_Mgt_sql, r.Mgt)
	ts = addData(ts, dm.Feature_UatSupport_sql, r.UatSupport)
	ts = addData(ts, dm.Feature_Marketing_sql, r.Marketing)
	ts = addData(ts, dm.Feature_Contingency_sql, r.Contingency)
	ts = addData(ts, dm.Feature_TrackerID_sql, r.TrackerID)
	ts = addData(ts, dm.Feature_AdoID_sql, r.AdoID)
	ts = addData(ts, dm.Feature_FreshdeskID_sql, r.FreshdeskID)
	ts = addData(ts, dm.Feature_ExtRef_sql, r.ExtRef)
	ts = addData(ts, dm.Feature_ExtRef2_sql, r.ExtRef2)
	ts = addData(ts, dm.Feature_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.Feature_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.Feature_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.Feature_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.Feature_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.Feature_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.Feature_SYSDeleted_sql, r.SYSDeleted)
	ts = addData(ts, dm.Feature_SYSDeletedBy_sql, r.SYSDeletedBy)
	ts = addData(ts, dm.Feature_SYSDeletedHost_sql, r.SYSDeletedHost)
	ts = addData(ts, dm.Feature_Developer_sql, r.Developer)
	ts = addData(ts, dm.Feature_Approver_sql, r.Approver)
	ts = addData(ts, dm.Feature_Notes_sql, r.Notes)
	ts = addData(ts, dm.Feature_OffProfile_sql, r.OffProfile)
	ts = addData(ts, dm.Feature_OffProfileJustification_sql, r.OffProfileJustification)
	ts = addData(ts, dm.Feature_SYSActivity_sql, r.SYSActivity)
		
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := "INSERT INTO " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Feature_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Feature_Delete(r.FeatureID)
	das.Execute(tsql)



	return err

}



// feature_Fetch read all Feature's
func feature_Fetch(tsql string) (int, []dm.Feature, dm.Feature, error) {

	var recItem dm.Feature
	var recList []dm.Feature

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.Feature_SYSId_sql, "0")
	   recItem.FeatureID  = get_String(rec, dm.Feature_FeatureID_sql, "")
	   recItem.EstimationSessionID  = get_String(rec, dm.Feature_EstimationSessionID_sql, "")
	   recItem.ConfidenceID  = get_String(rec, dm.Feature_ConfidenceID_sql, "")
	   recItem.Name  = get_String(rec, dm.Feature_Name_sql, "")
	   recItem.DevEstimate  = get_String(rec, dm.Feature_DevEstimate_sql, "")
	   recItem.DevUplift  = get_String(rec, dm.Feature_DevUplift_sql, "")
	   recItem.Reqs  = get_String(rec, dm.Feature_Reqs_sql, "")
	   recItem.AnalystTest  = get_String(rec, dm.Feature_AnalystTest_sql, "")
	   recItem.Docs  = get_String(rec, dm.Feature_Docs_sql, "")
	   recItem.Mgt  = get_String(rec, dm.Feature_Mgt_sql, "")
	   recItem.UatSupport  = get_String(rec, dm.Feature_UatSupport_sql, "")
	   recItem.Marketing  = get_String(rec, dm.Feature_Marketing_sql, "")
	   recItem.Contingency  = get_String(rec, dm.Feature_Contingency_sql, "")
	   recItem.TrackerID  = get_String(rec, dm.Feature_TrackerID_sql, "")
	   recItem.AdoID  = get_String(rec, dm.Feature_AdoID_sql, "")
	   recItem.FreshdeskID  = get_String(rec, dm.Feature_FreshdeskID_sql, "")
	   recItem.ExtRef  = get_String(rec, dm.Feature_ExtRef_sql, "")
	   recItem.ExtRef2  = get_String(rec, dm.Feature_ExtRef2_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.Feature_SYSCreated_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.Feature_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.Feature_SYSCreatedHost_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.Feature_SYSUpdated_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.Feature_SYSUpdatedBy_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.Feature_SYSUpdatedHost_sql, "")
	   recItem.SYSDeleted  = get_String(rec, dm.Feature_SYSDeleted_sql, "")
	   recItem.SYSDeletedBy  = get_String(rec, dm.Feature_SYSDeletedBy_sql, "")
	   recItem.SYSDeletedHost  = get_String(rec, dm.Feature_SYSDeletedHost_sql, "")
	   recItem.Developer  = get_String(rec, dm.Feature_Developer_sql, "")
	   recItem.Approver  = get_String(rec, dm.Feature_Approver_sql, "")
	   recItem.Notes  = get_String(rec, dm.Feature_Notes_sql, "")
	   recItem.OffProfile  = get_String(rec, dm.Feature_OffProfile_sql, "")
	   recItem.OffProfileJustification  = get_String(rec, dm.Feature_OffProfileJustification_sql, "")
	   recItem.SYSActivity  = get_String(rec, dm.Feature_SYSActivity_sql, "")
	
	// If there are fields below, create the methods in adaptor\Feature_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func Feature_NewID(r dm.Feature) string {
	
			id := uuid.New().String()
	
	return id
}



// feature_Fetch read all Feature's
func Feature_New() (int, []dm.Feature, dm.Feature, error) {

	var r = dm.Feature{}
	var rList []dm.Feature
	

	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}