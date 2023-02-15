package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/resource.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Resource (resource)
// Endpoint 	        : Resource (ResourceID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 07/02/2023 at 18:52:38
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

var Resource_SQLbase string
var Resource_QualifiedName string
func init(){
	Resource_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.Resource_SQLTable)
	Resource_SQLbase =  das.SELECTALL + das.FROM + Resource_QualifiedName
}

// Resource_GetList() returns a list of all Resource records
func Resource_GetList() (int, []dm.Resource, error) {
	
	count, resourceList, err := Resource_GetListFiltered("")
	
	return count, resourceList, err
}

// Resource_GetListFiltered() returns a filtered list of all Resource records
func Resource_GetListFiltered(filter string) (int, []dm.Resource, error) {
	
	tsql := Resource_SQLbase
	if filter != "" {
		tsql = tsql + " " + das.WHERE + filter
	}
	count, resourceList, _, _ := resource_Fetch(tsql)
	
	return count, resourceList, nil
}


// Resource_GetLookup() returns a lookup list of all Resource items in lookup format
func Resource_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, resourceList, _ := Resource_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: resourceList[i].Code, Name: resourceList[i].Name})
	}
	return returnList
}

// Resource_GetFilteredLookup() returns a lookup list of all Resource items in lookup format
func Resource_GetFilteredLookup(requestObject string,requestField string) []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	reqClass := "Resource"
	reqField := requestObject+"-"+requestField
	reqCategory := "Filter"
	filter,_ := Data_GetString(reqClass, reqField, reqCategory)
	if filter == "" {
		logs.Warning("Resource_GetFilteredLookup() - No filter found for " + reqClass + " " + reqField)
	} 
	count, resourceList, _ := Resource_GetListFiltered(filter)
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: resourceList[i].Code, Name: resourceList[i].Name})
	}
	return returnList
}



// Resource_GetByID() returns a single Resource record
func Resource_GetByID(id string) (int, dm.Resource, error) {


	tsql := Resource_SQLbase
	tsql = tsql + " " + das.WHERE + dm.Resource_SQLSearchID + das.EQ + das.ID(id)
	_, _, resourceItem, _ := resource_Fetch(tsql)

	// START
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, resourceItem, nil
}



// Resource_DeleteByID() deletes a single Resource record
func Resource_Delete(id string) {


// Uses Hard Delete
	object_Table := Resource_QualifiedName
	tsql := das.DELETE + das.FROM + object_Table
	tsql = tsql + " " + das.WHERE + dm.Resource_SQLSearchID + das.EQ + das.ID(id)
	das.Execute(tsql)	

	
}














	
	// Resource_SoftDeleteByID() soft deletes a single Resource record
func Resource_SoftDelete(id string) {
	//Uses Soft Delete
		_,resourceItem,_ := Resource_GetByID(id)
		resourceItem.SYSDeletedBy = Audit_Update("", Audit_Host())
		resourceItem.SYSDeleted = Audit_Update("", Audit_TimeStamp())
		resourceItem.SYSDeletedHost = Audit_Update("", Audit_Host())
		Resource_StoreSystem(resourceItem)
}
	


// Resource_Store() saves/stores a Resource record to the database
func Resource_Store(r dm.Resource,req *http.Request) error {

	r, err := Resource_Validate(r)
	if err == nil {
		err = resource_Save(r, Audit_User(req))
	} else {
		logs.Information("Resource_Store()", err.Error())
	}

	return err
}

// Resource_StoreSystem() saves/stores a Resource record to the database
func Resource_StoreSystem(r dm.Resource) error {
	
	r, err := Resource_Validate(r)
	if err == nil {
		err = resource_Save(r, Audit_Host())
	} else {
		logs.Information("Resource_Store()", err.Error())
	}

	return err
}

// Resource_Validate() validates for saves/stores a Resource record to the database
func Resource_Validate(r dm.Resource) (dm.Resource, error) {
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

// resource_Save() saves/stores a Resource record to the database
func resource_Save(r dm.Resource,usr string) error {

    var err error



	

	if len(r.ResourceID) == 0 {
		r.ResourceID = Resource_NewID(r)
	}

// If there are fields below, create the methods in dao\resource_impl.go






















	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	r.SYSDbVersion = core.DB_Version()
	
logs.Storing("Resource",fmt.Sprintf("%v", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.Resource_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.Resource_ResourceID_sql, r.ResourceID)
	ts = addData(ts, dm.Resource_Code_sql, r.Code)
	ts = addData(ts, dm.Resource_Name_sql, r.Name)
	ts = addData(ts, dm.Resource_Email_sql, r.Email)
	ts = addData(ts, dm.Resource_Class_sql, r.Class)
	ts = addData(ts, dm.Resource_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.Resource_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.Resource_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.Resource_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.Resource_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.Resource_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.Resource_SYSDeleted_sql, r.SYSDeleted)
	ts = addData(ts, dm.Resource_SYSDeletedBy_sql, r.SYSDeletedBy)
	ts = addData(ts, dm.Resource_SYSDeletedHost_sql, r.SYSDeletedHost)
	ts = addData(ts, dm.Resource_UserActive_sql, r.UserActive)
	ts = addData(ts, dm.Resource_SYSActivity_sql, r.SYSActivity)
	ts = addData(ts, dm.Resource_Notes_sql, r.Notes)
	ts = addData(ts, dm.Resource_SYSDbVersion_sql, r.SYSDbVersion)
	ts = addData(ts, dm.Resource_Comments_sql, r.Comments)
		
	// 
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := das.INSERT + das.INTO + Resource_QualifiedName
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " "+das.VALUES +"(" + values(ts) + ")"

	Resource_Delete(r.ResourceID)
	das.Execute(tsql)



	return err

}



// resource_Fetch read all Resource's
func resource_Fetch(tsql string) (int, []dm.Resource, dm.Resource, error) {

	var recItem dm.Resource
	var recList []dm.Resource

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.Resource_SYSId_sql, "0")
	   recItem.ResourceID  = get_String(rec, dm.Resource_ResourceID_sql, "")
	   recItem.Code  = get_String(rec, dm.Resource_Code_sql, "")
	   recItem.Name  = get_String(rec, dm.Resource_Name_sql, "")
	   recItem.Email  = get_String(rec, dm.Resource_Email_sql, "")
	   recItem.Class  = get_String(rec, dm.Resource_Class_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.Resource_SYSCreated_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.Resource_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.Resource_SYSCreatedHost_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.Resource_SYSUpdated_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.Resource_SYSUpdatedBy_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.Resource_SYSUpdatedHost_sql, "")
	   recItem.SYSDeleted  = get_String(rec, dm.Resource_SYSDeleted_sql, "")
	   recItem.SYSDeletedBy  = get_String(rec, dm.Resource_SYSDeletedBy_sql, "")
	   recItem.SYSDeletedHost  = get_String(rec, dm.Resource_SYSDeletedHost_sql, "")
	   recItem.UserActive  = get_String(rec, dm.Resource_UserActive_sql, "")
	   recItem.SYSActivity  = get_String(rec, dm.Resource_SYSActivity_sql, "")
	   recItem.Notes  = get_String(rec, dm.Resource_Notes_sql, "")
	   recItem.SYSDbVersion  = get_String(rec, dm.Resource_SYSDbVersion_sql, "")
	   recItem.Comments  = get_String(rec, dm.Resource_Comments_sql, "")
	
	// If there are fields below, create the methods in adaptor\Resource_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
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
	


func Resource_NewID(r dm.Resource) string {
	
			id := uuid.New().String()
	
	return id
}



// resource_Fetch read all Resource's
func Resource_New() (int, []dm.Resource, dm.Resource, error) {

	var r = dm.Resource{}
	var rList []dm.Resource
	

	// START
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}