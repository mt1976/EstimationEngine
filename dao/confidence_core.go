package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/confidence.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Confidence (confidence)
// Endpoint 	        : Confidence (ConfidenceID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 13/03/2023 at 14:22:25
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

var Confidence_SQLbase string
var Confidence_QualifiedName string

func init() {
	Confidence_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.Confidence_SQLTable)
	Confidence_SQLbase = das.SELECTALL + das.FROM + Confidence_QualifiedName
}

// Confidence_GetList() returns a list of all Confidence records
func Confidence_GetList() (int, []dm.Confidence, error) {
	count, confidenceList, err := Confidence_GetListFiltered("")
	return count, confidenceList, err
}

// Confidence_GetListFiltered() returns a filtered list of all Confidence records
func Confidence_GetListFiltered(filter string) (int, []dm.Confidence, error) {

	tsql := Confidence_SQLbase
	if filter != "" {
		tsql = tsql + " " + das.WHERE + filter
	}
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

// Confidence_GetFilteredLookup() returns a lookup list of all Confidence items in lookup format
func Confidence_GetFilteredLookup(requestObject string, requestField string) []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	objectName := Translate("ObjectName", requestObject)
	reqField := requestField + "_Confidence_Filter"
	filter, _ := Data_GetString(objectName, reqField, dm.Data_Category_FilterRule)
	if filter == "" {
		logs.Warning("Confidence_GetFilteredLookup() - No filter found : " + reqField + " for Object: " + objectName)
	}
	count, confidenceList, _ := Confidence_GetListFiltered(filter)
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: confidenceList[i].Code, Name: confidenceList[i].Name})
	}
	return returnList
}

// Confidence_GetByID() returns a single Confidence record
func Confidence_GetByID(id string) (int, dm.Confidence, error) {

	tsql := Confidence_SQLbase
	tsql = tsql + " " + das.WHERE + dm.Confidence_SQLSearchID + das.EQ + das.ID(id)
	_, _, confidenceItem, _ := confidence_Fetch(tsql)

	confidenceItem = Confidence_PostGet(confidenceItem, id)

	return 1, confidenceItem, nil
}

func Confidence_PostGet(confidenceItem dm.Confidence, id string) dm.Confidence {
	// START
	// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
	return confidenceItem
}

// Confidence_DeleteByID() deletes a single Confidence record
func Confidence_Delete(id string) {
	Confidence_SoftDelete(id)
}

// Confidence_SoftDelete(id string) soft deletes a single Confidence record
func Confidence_SoftDelete(id string) {
	//Uses Soft Delete
	_, confidenceItem, _ := Confidence_GetByID(id)
	confidenceItem.SYSDeletedBy = Audit_Update("", Audit_Host())
	confidenceItem.SYSDeleted = Audit_Update("", Audit_TimeStamp())
	confidenceItem.SYSDeletedHost = Audit_Update("", Audit_Host())
	_, err := Confidence_StoreSystem(confidenceItem)
	if err != nil {
		logs.Error("Confidence_SoftDelete()", err)
	}
}

// Confidence_HardDelete(id string) soft deletes a single Confidence record
func Confidence_HardDelete(id string) {
	// Uses Hard Delete
	object_Table := Confidence_QualifiedName
	tsql := das.DELETE + das.FROM + object_Table
	tsql = tsql + " " + das.WHERE + dm.Confidence_SQLSearchID + das.EQ + das.ID(id)
	das.Execute(tsql)
	//if err != nil {
	//	logs.Error("Confidence_SoftDelete()",err)
	//}
}

// Confidence_Store() saves/stores a Confidence record to the database
func Confidence_Store(r dm.Confidence, req *http.Request) (dm.Confidence, error) {

	r, err := Confidence_Validate(r)
	if err == nil {
		err = confidence_Save(r, Audit_User(req))
	} else {
		logs.Information("Confidence_Store()", err.Error())
	}

	return r, err
}

// Confidence_StoreSystem() saves/stores a Confidence record to the database
func Confidence_StoreSystem(r dm.Confidence) (dm.Confidence, error) {

	r, err := Confidence_Validate(r)
	if err == nil {
		err = confidence_Save(r, Audit_Host())
	} else {
		logs.Information("Confidence_Store()", err.Error())
	}

	return r, err
}

// Confidence_StoreProcess() saves/stores a Confidence record to the database
func Confidence_StoreProcess(r dm.Confidence, operator string) (dm.Confidence, error) {

	r, err := Confidence_Validate(r)
	if err == nil {
		err = confidence_Save(r, operator)
	} else {
		logs.Information("Confidence_StoreProcess()", err.Error())
	}

	return r, err
}

// Confidence_Validate() validates for saves/stores a Confidence record to the database
func Confidence_Validate(r dm.Confidence) (dm.Confidence, error) {
	var err error
	// START
	// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	//

	return r, err
}

//

// confidence_Save() saves/stores a Confidence record to the database
func confidence_Save(r dm.Confidence, usr string) error {

	var err error

	if len(r.ConfidenceID) == 0 {
		r.ConfidenceID = Confidence_NewID(r)
	}

	// If there are fields below, create the methods in dao\confidence_impl.go

	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost, Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("", usr)
	r.SYSUpdatedHost = Audit_Update("", Audit_Host())
	r.SYSDbVersion = core.DB_Version()

	logs.Storing("Confidence", fmt.Sprintf("%v", r))

	//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
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
	ts = addData(ts, dm.Confidence_SYSDbVersion_sql, r.SYSDbVersion)
	ts = addData(ts, dm.Confidence_Comments_sql, r.Comments)

	//
	// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END

	tsql := das.INSERT + das.INTO + Confidence_QualifiedName
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " " + das.VALUES + "(" + values(ts) + ")"

	Confidence_HardDelete(r.ConfidenceID)
	das.Execute(tsql)

	return err

}

// confidence_Fetch read all Confidence's
func confidence_Fetch(tsql string) (int, []dm.Confidence, dm.Confidence, error) {

	var recItem dm.Confidence
	var recList []dm.Confidence

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(), err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
		// START
		// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
		//
		recItem.SYSId = get_Int(rec, dm.Confidence_SYSId_sql, "0")
		recItem.ConfidenceID = get_String(rec, dm.Confidence_ConfidenceID_sql, "")
		recItem.Code = get_String(rec, dm.Confidence_Code_sql, "")
		recItem.Name = get_String(rec, dm.Confidence_Name_sql, "")
		recItem.Perc = get_String(rec, dm.Confidence_Perc_sql, "")
		recItem.SYSCreated = get_String(rec, dm.Confidence_SYSCreated_sql, "")
		recItem.SYSCreatedBy = get_String(rec, dm.Confidence_SYSCreatedBy_sql, "")
		recItem.SYSCreatedHost = get_String(rec, dm.Confidence_SYSCreatedHost_sql, "")
		recItem.SYSUpdated = get_String(rec, dm.Confidence_SYSUpdated_sql, "")
		recItem.SYSUpdatedBy = get_String(rec, dm.Confidence_SYSUpdatedBy_sql, "")
		recItem.SYSUpdatedHost = get_String(rec, dm.Confidence_SYSUpdatedHost_sql, "")
		recItem.SYSDeleted = get_String(rec, dm.Confidence_SYSDeleted_sql, "")
		recItem.SYSDeletedBy = get_String(rec, dm.Confidence_SYSDeletedBy_sql, "")
		recItem.SYSDeletedHost = get_String(rec, dm.Confidence_SYSDeletedHost_sql, "")
		recItem.Notes = get_String(rec, dm.Confidence_Notes_sql, "")
		recItem.SYSDbVersion = get_String(rec, dm.Confidence_SYSDbVersion_sql, "")
		recItem.Comments = get_String(rec, dm.Confidence_Comments_sql, "")

		// If there are fields below, create the methods in dao\Confidence_adaptor.go
		//
		// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
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
	// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//

	//
	// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
	rList = append(rList, r)
	return 1, rList, r, nil
}
