package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/credentialspassword.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : CredentialsPassword (credentialspassword)
// Endpoint 	        : CredentialsPassword (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 08/03/2023 at 18:42:23
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

var CredentialsPassword_SQLbase string
var CredentialsPassword_QualifiedName string
func init(){
	CredentialsPassword_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.CredentialsPassword_SQLTable)
	CredentialsPassword_SQLbase =  das.SELECTALL + das.FROM + CredentialsPassword_QualifiedName
}

// CredentialsPassword_GetList() returns a list of all CredentialsPassword records
func CredentialsPassword_GetList() (int, []dm.CredentialsPassword, error) {
	count, credentialspasswordList, err := CredentialsPassword_GetListFiltered("")
	return count, credentialspasswordList, err
}

// CredentialsPassword_GetListFiltered() returns a filtered list of all CredentialsPassword records
func CredentialsPassword_GetListFiltered(filter string) (int, []dm.CredentialsPassword, error) {
	
	count, credentialspasswordList, _ := CredentialsPassword_GetList_impl(filter)
	
	return count, credentialspasswordList, nil
}



// CredentialsPassword_GetByID() returns a single CredentialsPassword record
func CredentialsPassword_GetByID(id string) (int, dm.CredentialsPassword, error) {


	 _, credentialspasswordItem, _ := CredentialsPassword_GetByID_impl(id)
	

	credentialspasswordItem = CredentialsPassword_PostGet(credentialspasswordItem,id)

	return 1, credentialspasswordItem, nil
}

func CredentialsPassword_PostGet(credentialspasswordItem dm.CredentialsPassword,id string) dm.CredentialsPassword {
	// START
	// Dynamically generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	credentialspasswordItem.UserName,credentialspasswordItem.UserName_props = CredentialsPassword_UserName_validate_impl (GET,id,credentialspasswordItem.UserName,credentialspasswordItem,credentialspasswordItem.UserName_props)
	credentialspasswordItem.PasswordOld,credentialspasswordItem.PasswordOld_props = CredentialsPassword_PasswordOld_validate_impl (GET,id,credentialspasswordItem.PasswordOld,credentialspasswordItem,credentialspasswordItem.PasswordOld_props)
	credentialspasswordItem.PasswordNew,credentialspasswordItem.PasswordNew_props = CredentialsPassword_PasswordNew_validate_impl (GET,id,credentialspasswordItem.PasswordNew,credentialspasswordItem,credentialspasswordItem.PasswordNew_props)
	credentialspasswordItem.PasswordConfirm,credentialspasswordItem.PasswordConfirm_props = CredentialsPassword_PasswordConfirm_validate_impl (GET,id,credentialspasswordItem.PasswordConfirm,credentialspasswordItem,credentialspasswordItem.PasswordConfirm_props)
	// 
	// Dynamically generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return credentialspasswordItem
}




// CredentialsPassword_DeleteByID() deletes a single CredentialsPassword record
func CredentialsPassword_Delete(id string) {
	CredentialsPassword_Delete_impl(id)
}

// CredentialsPassword_HardDelete(id string) soft deletes a single CredentialsPassword record
func CredentialsPassword_HardDelete(id string) {
			// Uses Hard Delete
		object_Table := CredentialsPassword_QualifiedName
		tsql := das.DELETE + das.FROM + object_Table
		tsql = tsql + " " + das.WHERE + dm.CredentialsPassword_SQLSearchID + das.EQ + das.ID(id)
		das.Execute(tsql)
		//if err != nil {
		//	logs.Error("CredentialsPassword_SoftDelete()",err)
		//}
}


// CredentialsPassword_Store() saves/stores a CredentialsPassword record to the database
func CredentialsPassword_Store(r dm.CredentialsPassword,req *http.Request) (dm.CredentialsPassword,error) {

	r, err := CredentialsPassword_Validate(r)
	if err == nil {
		err = credentialspassword_Save(r, Audit_User(req))
	} else {
		logs.Information("CredentialsPassword_Store()", err.Error())
	}

	return r, err
}

// CredentialsPassword_StoreSystem() saves/stores a CredentialsPassword record to the database
func CredentialsPassword_StoreSystem(r dm.CredentialsPassword) (dm.CredentialsPassword,error) {
	
	r, err := CredentialsPassword_Validate(r)
	if err == nil {
		err = credentialspassword_Save(r, Audit_Host())
	} else {
		logs.Information("CredentialsPassword_Store()", err.Error())
	}

	return r, err
}

// CredentialsPassword_StoreProcess() saves/stores a CredentialsPassword record to the database
func CredentialsPassword_StoreProcess(r dm.CredentialsPassword, operator string) (dm.CredentialsPassword,error) {
	
	r, err := CredentialsPassword_Validate(r)
	if err == nil {
		err = credentialspassword_Save(r, operator)
	} else {
		logs.Information("CredentialsPassword_StoreProcess()", err.Error())
	}

	return r, err
}

// CredentialsPassword_Validate() validates for saves/stores a CredentialsPassword record to the database
func CredentialsPassword_Validate(r dm.CredentialsPassword) (dm.CredentialsPassword, error) {
	var err error
	// START
	// Dynamically generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	r.UserName,r.UserName_props = CredentialsPassword_UserName_validate_impl (PUT,r.ID,r.UserName,r,r.UserName_props)
	if r.UserName_props.MsgMessage != "" {
		err = errors.New(r.UserName_props.MsgMessage)
	}
	r.PasswordOld,r.PasswordOld_props = CredentialsPassword_PasswordOld_validate_impl (PUT,r.ID,r.PasswordOld,r,r.PasswordOld_props)
	if r.PasswordOld_props.MsgMessage != "" {
		err = errors.New(r.PasswordOld_props.MsgMessage)
	}
	r.PasswordNew,r.PasswordNew_props = CredentialsPassword_PasswordNew_validate_impl (PUT,r.ID,r.PasswordNew,r,r.PasswordNew_props)
	if r.PasswordNew_props.MsgMessage != "" {
		err = errors.New(r.PasswordNew_props.MsgMessage)
	}
	r.PasswordConfirm,r.PasswordConfirm_props = CredentialsPassword_PasswordConfirm_validate_impl (PUT,r.ID,r.PasswordConfirm,r,r.PasswordConfirm_props)
	if r.PasswordConfirm_props.MsgMessage != "" {
		err = errors.New(r.PasswordConfirm_props.MsgMessage)
	}
	// 

	

	return r,err
}
//

// credentialspassword_Save() saves/stores a CredentialsPassword record to the database
func credentialspassword_Save(r dm.CredentialsPassword,usr string) error {

    var err error

	if len(r.ID) == 0 {
		r.ID = CredentialsPassword_NewID(r)
	}

// If there are fields below, create the methods in dao\credentialspassword_impl.go
    r.UserName,err = CredentialsPassword_UserName_OnStore_impl (r.UserName,r,usr)
    r.PasswordOld,err = CredentialsPassword_PasswordOld_OnStore_impl (r.PasswordOld,r,usr)
    r.PasswordNew,err = CredentialsPassword_PasswordNew_OnStore_impl (r.PasswordNew,r,usr)
    r.PasswordConfirm,err = CredentialsPassword_PasswordConfirm_OnStore_impl (r.PasswordConfirm,r,usr)

logs.Storing("CredentialsPassword",fmt.Sprintf("%v", r))

// Please Create Functions Below in the adaptor/CredentialsPassword_impl.go file
	err1 := CredentialsPassword_Delete_impl(r.ID)
	err2 := CredentialsPassword_Update_impl(r.ID,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



	// credentialspassword_Fetch is not required as GetByID, GetAll etc... have been diverted to _impl
	


func CredentialsPassword_NewID(r dm.CredentialsPassword) string {
	
	id := uuid.New().String()
	
	return id
}



// credentialspassword_Fetch read all CredentialsPassword's
func CredentialsPassword_New() (int, []dm.CredentialsPassword, dm.CredentialsPassword, error) {

	var r = dm.CredentialsPassword{}
	var rList []dm.CredentialsPassword
	

	// START
	// Dynamically generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	r.UserName,r.UserName_props = CredentialsPassword_UserName_validate_impl (NEW,r.ID,r.UserName,r,r.UserName_props)
	r.PasswordOld,r.PasswordOld_props = CredentialsPassword_PasswordOld_validate_impl (NEW,r.ID,r.PasswordOld,r,r.PasswordOld_props)
	r.PasswordNew,r.PasswordNew_props = CredentialsPassword_PasswordNew_validate_impl (NEW,r.ID,r.PasswordNew,r,r.PasswordNew_props)
	r.PasswordConfirm,r.PasswordConfirm_props = CredentialsPassword_PasswordConfirm_validate_impl (NEW,r.ID,r.PasswordConfirm,r,r.PasswordConfirm_props)
	
	// 
	// Dynamically generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	rList = append(rList, r)
	return 1, rList, r, nil
}