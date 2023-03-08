package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/doctype.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : DocType (doctype)
// Endpoint 	        : DocType (DocTypeID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 04/03/2023 at 20:14:10
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
	"github.com/pkg/errors"
)

var DocType_SQLbase string
var DocType_QualifiedName string
func init(){
	DocType_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.DocType_SQLTable)
	DocType_SQLbase =  das.SELECTALL + das.FROM + DocType_QualifiedName
}

// DocType_GetList() returns a list of all DocType records
func DocType_GetList() (int, []dm.DocType, error) {
	count, doctypeList, err := DocType_GetListFiltered("")
	return count, doctypeList, err
}

// DocType_GetListFiltered() returns a filtered list of all DocType records
func DocType_GetListFiltered(filter string) (int, []dm.DocType, error) {
	
	tsql := DocType_SQLbase
	if filter != "" {
		tsql = tsql + " " + das.WHERE + filter
	}
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

// DocType_GetFilteredLookup() returns a lookup list of all DocType items in lookup format
func DocType_GetFilteredLookup(requestObject string,requestField string) []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	objectName := Translate("ObjectName", requestObject)
	reqField := requestField+"_DocType_Filter"
	filter,_ := Data_GetString(objectName, reqField, dm.Data_Category_FilterRule)
	if filter == "" {
		logs.Warning("DocType_GetFilteredLookup() - No filter found : " + reqField + " for Object: " + objectName)
	} 
	count, doctypeList, _ := DocType_GetListFiltered(filter)
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: doctypeList[i].Code, Name: doctypeList[i].Name})
	}
	return returnList
}



// DocType_GetByID() returns a single DocType record
func DocType_GetByID(id string) (int, dm.DocType, error) {


	tsql := DocType_SQLbase
	tsql = tsql + " " + das.WHERE + dm.DocType_SQLSearchID + das.EQ + das.ID(id)
	_, _, doctypeItem, _ := doctype_Fetch(tsql)


	doctypeItem = DocType_PostGet(doctypeItem,id)

	return 1, doctypeItem, nil
}

func DocType_PostGet(doctypeItem dm.DocType,id string) dm.DocType {
	// START
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	doctypeItem.Code,doctypeItem.Code_props = DocType_Code_validate_impl (GET,id,doctypeItem.Code,doctypeItem,doctypeItem.Code_props)
	// 
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return doctypeItem
}




// DocType_DeleteByID() deletes a single DocType record
func DocType_Delete(id string) {
		DocType_SoftDelete(id)
	}
// DocType_SoftDelete(id string) soft deletes a single DocType record
func DocType_SoftDelete(id string) {
	//Uses Soft Delete
		_,doctypeItem,_ := DocType_GetByID(id)
		doctypeItem.SYSDeletedBy = Audit_Update("", Audit_Host())
		doctypeItem.SYSDeleted = Audit_Update("", Audit_TimeStamp())
		doctypeItem.SYSDeletedHost = Audit_Update("", Audit_Host())
		_,err := DocType_StoreSystem(doctypeItem)
		if err != nil {
			logs.Error("DocType_SoftDelete()",err)
		}
}
	

// DocType_HardDelete(id string) soft deletes a single DocType record
func DocType_HardDelete(id string) {
			// Uses Hard Delete
		object_Table := DocType_QualifiedName
		tsql := das.DELETE + das.FROM + object_Table
		tsql = tsql + " " + das.WHERE + dm.DocType_SQLSearchID + das.EQ + das.ID(id)
		das.Execute(tsql)
		//if err != nil {
		//	logs.Error("DocType_SoftDelete()",err)
		//}
}


// DocType_Store() saves/stores a DocType record to the database
func DocType_Store(r dm.DocType,req *http.Request) (dm.DocType,error) {

	r, err := DocType_Validate(r)
	if err == nil {
		err = doctype_Save(r, Audit_User(req))
	} else {
		logs.Information("DocType_Store()", err.Error())
	}

	return r, err
}

// DocType_StoreSystem() saves/stores a DocType record to the database
func DocType_StoreSystem(r dm.DocType) (dm.DocType,error) {
	
	r, err := DocType_Validate(r)
	if err == nil {
		err = doctype_Save(r, Audit_Host())
	} else {
		logs.Information("DocType_Store()", err.Error())
	}

	return r, err
}

// DocType_StoreProcess() saves/stores a DocType record to the database
func DocType_StoreProcess(r dm.DocType, operator string) (dm.DocType,error) {
	
	r, err := DocType_Validate(r)
	if err == nil {
		err = doctype_Save(r, operator)
	} else {
		logs.Information("DocType_StoreProcess()", err.Error())
	}

	return r, err
}

// DocType_Validate() validates for saves/stores a DocType record to the database
func DocType_Validate(r dm.DocType) (dm.DocType, error) {
	var err error
	// START
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	r.Code,r.Code_props = DocType_Code_validate_impl (PUT,r.DocTypeID,r.Code,r,r.Code_props)
	if r.Code_props.MsgMessage != "" {
		err = errors.New(r.Code_props.MsgMessage)
	}
	// 
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
    r.Code,err = DocType_Code_OnStore_impl (r.Code,r,usr)

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
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := das.INSERT + das.INTO + DocType_QualifiedName
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " "+das.VALUES +"(" + values(ts) + ")"

	DocType_HardDelete(r.DocTypeID)
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
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	
	// If there are fields below, create the methods in dao\DocType_adaptor.go
	   recItem.Code  = DocType_Code_OnFetch_impl (recItem)
	// 
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	r.Code,r.Code_props = DocType_Code_validate_impl (NEW,r.DocTypeID,r.Code,r,r.Code_props)
	
	// 
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	rList = append(rList, r)
	return 1, rList, r, nil
}