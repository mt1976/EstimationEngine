package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/credentials.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Credentials (credentials)
// Endpoint 	        : Credentials (Id)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/03/2023 at 19:24:47
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"fmt"
	"net/http"
	"errors"
	core "github.com/mt1976/ebEstimates/core"
	"github.com/google/uuid"
	das  "github.com/mt1976/ebEstimates/das"
	dm   "github.com/mt1976/ebEstimates/datamodel"
	logs   "github.com/mt1976/ebEstimates/logs"
)

var Credentials_SQLbase string
var Credentials_QualifiedName string
func init(){
	Credentials_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.Credentials_SQLTable)
	Credentials_SQLbase =  das.SELECTALL + das.FROM + Credentials_QualifiedName
}

// Credentials_GetList() returns a list of all Credentials records
func Credentials_GetList() (int, []dm.Credentials, error) {
	count, credentialsList, err := Credentials_GetListFiltered("")
	return count, credentialsList, err
}

// Credentials_GetListFiltered() returns a filtered list of all Credentials records
func Credentials_GetListFiltered(filter string) (int, []dm.Credentials, error) {
	
	tsql := Credentials_SQLbase
	if filter != "" {
		tsql = tsql + " " + das.WHERE + filter
	}
	count, credentialsList, _, _ := credentials_Fetch(tsql)
	
	return count, credentialsList, nil
}


// Credentials_GetLookup() returns a lookup list of all Credentials items in lookup format
func Credentials_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, credentialsList, _ := Credentials_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: credentialsList[i].Id, Name: credentialsList[i].Username})
	}
	return returnList
}

// Credentials_GetFilteredLookup() returns a lookup list of all Credentials items in lookup format
func Credentials_GetFilteredLookup(requestObject string,requestField string) []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	objectName := Translate("ObjectName", requestObject)
	reqField := requestField+"_Credentials_Filter"
	
	usage := "Defines a filter for a lookup list of Credentials records, when requested by "+requestField+"." + core.TEXTAREA_CR
	usage = usage + "Fields can be any of those in the underlying DB table." + core.TEXTAREA_CR
	usage = usage + "Examples Below:" + core.TEXTAREA_CR
	usage = usage + "* datalength(_deleted) = 0 or " + core.TEXTAREA_CR 
	usage = usage + "* class IN ('x','y','z')"

	filter,_ := Data_GetString(objectName, reqField, dm.Data_Category_FilterRule,usage)
	if filter == "" {
		logs.Warning("Credentials_GetFilteredLookup() - No filter found : " + reqField + " for Object: " + objectName)
	} 
	count, credentialsList, _ := Credentials_GetListFiltered(filter)
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: credentialsList[i].Id, Name: credentialsList[i].Username})
	}
	return returnList
}



// Credentials_GetByID() returns a single Credentials record
func Credentials_GetByID(id string) (int, dm.Credentials, error) {


	tsql := Credentials_SQLbase
	tsql = tsql + " " + das.WHERE + dm.Credentials_SQLSearchID + das.EQ + das.ID(id)
	_, _, credentialsItem, _ := credentials_Fetch(tsql)


	credentialsItem = Credentials_PostGet(credentialsItem,id)

	return 1, credentialsItem, nil
}

func Credentials_PostGet(credentialsItem dm.Credentials,id string) dm.Credentials {
	credentialsItem.State,credentialsItem.State_props = Credentials_State_validate_impl (GET,id,credentialsItem.State,credentialsItem,credentialsItem.State_props)

	return credentialsItem
}




// Credentials_DeleteByID() deletes a single Credentials record
func Credentials_Delete(id string) {
		Credentials_HardDelete(id)
	}

// Credentials_HardDelete(id string) soft deletes a single Credentials record
func Credentials_HardDelete(id string) {
			// Uses Hard Delete
		object_Table := Credentials_QualifiedName
		tsql := das.DELETE + das.FROM + object_Table
		tsql = tsql + " " + das.WHERE + dm.Credentials_SQLSearchID + das.EQ + das.ID(id)
		das.Execute(tsql)
		//if err != nil {
		//	logs.Error("Credentials_SoftDelete()",err)
		//}
}


// Credentials_Store() saves/stores a Credentials record to the database
func Credentials_Store(r dm.Credentials,req *http.Request) (dm.Credentials,error) {

	r, err := Credentials_Validate(r)
	if err == nil {
		err = credentials_Save(r, Audit_User(req))
	} else {
		logs.Information("Credentials_Store()", err.Error())
	}

	return r, err
}

// Credentials_StoreSystem() saves/stores a Credentials record to the database
func Credentials_StoreSystem(r dm.Credentials) (dm.Credentials,error) {
	
	r, err := Credentials_Validate(r)
	if err == nil {
		err = credentials_Save(r, Audit_Host())
	} else {
		logs.Information("Credentials_Store()", err.Error())
	}

	return r, err
}

// Credentials_StoreProcess() saves/stores a Credentials record to the database
func Credentials_StoreProcess(r dm.Credentials, operator string) (dm.Credentials,error) {
	
	r, err := Credentials_Validate(r)
	if err == nil {
		err = credentials_Save(r, operator)
	} else {
		logs.Information("Credentials_StoreProcess()", err.Error())
	}

	return r, err
}

// Credentials_Validate() validates for saves/stores a Credentials record to the database
func Credentials_Validate(r dm.Credentials) (dm.Credentials, error) {
	var err error
	r.State,r.State_props = Credentials_State_validate_impl (PUT,r.Id,r.State,r,r.State_props)
	if r.State_props.MsgMessage != "" {
		err = errors.New(r.State_props.MsgMessage)
	}

	// Cross Validation
	var errVal error
	r, _, errVal = Credentials_ObjectValidation_impl(PUT, r.Id, r)
	if errVal != nil {
		err = errVal
	}
	return r,err
}
//

// credentials_Save() saves/stores a Credentials record to the database
func credentials_Save(r dm.Credentials,usr string) error {

    var err error

	if len(r.Id) == 0 {
		r.Id = Credentials_NewID(r)
	}

// If there are fields below, create the methods in dao\credentials_impl.go
    r.State,err = Credentials_State_OnStore_impl (r.State,r,usr)

	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	r.SYSDbVersion = core.DB_Version()
	
logs.Storing("Credentials",fmt.Sprintf("%v", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.Credentials_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.Credentials_Id_sql, r.Id)
	ts = addData(ts, dm.Credentials_Username_sql, r.Username)
	ts = addData(ts, dm.Credentials_Password_sql, r.Password)
	ts = addData(ts, dm.Credentials_Firstname_sql, r.Firstname)
	ts = addData(ts, dm.Credentials_Lastname_sql, r.Lastname)
	ts = addData(ts, dm.Credentials_Knownas_sql, r.Knownas)
	ts = addData(ts, dm.Credentials_Email_sql, r.Email)
	ts = addData(ts, dm.Credentials_Issued_sql, r.Issued)
	ts = addData(ts, dm.Credentials_Expiry_sql, r.Expiry)
	ts = addData(ts, dm.Credentials_RoleType_sql, r.RoleType)
	ts = addData(ts, dm.Credentials_Brand_sql, r.Brand)
	ts = addData(ts, dm.Credentials_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.Credentials_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.Credentials_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.Credentials_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.Credentials_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.Credentials_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.Credentials_State_sql, r.State)
	ts = addData(ts, dm.Credentials_Notes_sql, r.Notes)
	ts = addData(ts, dm.Credentials_SYSDeleted_sql, r.SYSDeleted)
	ts = addData(ts, dm.Credentials_SYSDeletedBy_sql, r.SYSDeletedBy)
	ts = addData(ts, dm.Credentials_SYSDeletedHost_sql, r.SYSDeletedHost)
	ts = addData(ts, dm.Credentials_SYSActivity_sql, r.SYSActivity)
	ts = addData(ts, dm.Credentials_SYSDbVersion_sql, r.SYSDbVersion)
	ts = addData(ts, dm.Credentials_EmailNotifications_sql, r.EmailNotifications)
	ts = addData(ts, dm.Credentials_PasswordExpiry_sql, r.PasswordExpiry)
		
	// 
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := das.INSERT + das.INTO + Credentials_QualifiedName
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " "+das.VALUES +"(" + values(ts) + ")"

	Credentials_HardDelete(r.Id)
	das.Execute(tsql)

	



	return err

}



// credentials_Fetch read all Credentials's
func credentials_Fetch(tsql string) (int, []dm.Credentials, dm.Credentials, error) {

	var recItem dm.Credentials
	var recList []dm.Credentials

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.Credentials_SYSId_sql, "0")
	   recItem.Id  = get_String(rec, dm.Credentials_Id_sql, "")
	   recItem.Username  = get_String(rec, dm.Credentials_Username_sql, "")
	   recItem.Password  = get_String(rec, dm.Credentials_Password_sql, "")
	   recItem.Firstname  = get_String(rec, dm.Credentials_Firstname_sql, "")
	   recItem.Lastname  = get_String(rec, dm.Credentials_Lastname_sql, "")
	   recItem.Knownas  = get_String(rec, dm.Credentials_Knownas_sql, "")
	   recItem.Email  = get_String(rec, dm.Credentials_Email_sql, "")
	   recItem.Issued  = get_String(rec, dm.Credentials_Issued_sql, "")
	   recItem.Expiry  = get_String(rec, dm.Credentials_Expiry_sql, "")
	   recItem.RoleType  = get_String(rec, dm.Credentials_RoleType_sql, "")
	   recItem.Brand  = get_String(rec, dm.Credentials_Brand_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.Credentials_SYSCreated_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.Credentials_SYSUpdated_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.Credentials_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.Credentials_SYSCreatedHost_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.Credentials_SYSUpdatedBy_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.Credentials_SYSUpdatedHost_sql, "")
	   recItem.State  = get_String(rec, dm.Credentials_State_sql, "")
	   recItem.Notes  = get_String(rec, dm.Credentials_Notes_sql, "")
	   recItem.SYSDeleted  = get_String(rec, dm.Credentials_SYSDeleted_sql, "")
	   recItem.SYSDeletedBy  = get_String(rec, dm.Credentials_SYSDeletedBy_sql, "")
	   recItem.SYSDeletedHost  = get_String(rec, dm.Credentials_SYSDeletedHost_sql, "")
	   recItem.SYSActivity  = get_String(rec, dm.Credentials_SYSActivity_sql, "")
	   recItem.SYSDbVersion  = get_String(rec, dm.Credentials_SYSDbVersion_sql, "")
	   recItem.EmailNotifications  = get_String(rec, dm.Credentials_EmailNotifications_sql, "")
	   recItem.PasswordExpiry  = get_String(rec, dm.Credentials_PasswordExpiry_sql, "")
	
	// If there are fields below, create the methods in dao\Credentials_adaptor.go
	   recItem.State  = Credentials_State_OnFetch_impl (recItem)
	// 
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func Credentials_NewID(r dm.Credentials) string {
	
	id := uuid.New().String()
	
	return id
}



// credentials_Fetch read all Credentials's
func Credentials_New() (int, []dm.Credentials, dm.Credentials, error) {

	var r = dm.Credentials{}
	var rList []dm.Credentials
	

	// START
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	r.State,r.State_props = Credentials_State_validate_impl (NEW,r.Id,r.State,r,r.State_props)
	
	// 
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	rList = append(rList, r)
	return 1, rList, r, nil
}