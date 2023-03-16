package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/credentialsaction.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : CredentialsAction (credentialsaction)
// Endpoint 	        : CredentialsAction (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/03/2023 at 19:24:47
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

var CredentialsAction_SQLbase string
var CredentialsAction_QualifiedName string

func init() {
	CredentialsAction_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.CredentialsAction_SQLTable)
	CredentialsAction_SQLbase = das.SELECTALL + das.FROM + CredentialsAction_QualifiedName
}

// CredentialsAction_GetList() returns a list of all CredentialsAction records
func CredentialsAction_GetList() (int, []dm.CredentialsAction, error) {
	count, credentialsactionList, err := CredentialsAction_GetListFiltered("")
	return count, credentialsactionList, err
}

// CredentialsAction_GetListFiltered() returns a filtered list of all CredentialsAction records
func CredentialsAction_GetListFiltered(filter string) (int, []dm.CredentialsAction, error) {

	count, credentialsactionList, _ := CredentialsAction_GetList_impl(filter)

	return count, credentialsactionList, nil
}

// CredentialsAction_GetByID() returns a single CredentialsAction record
func CredentialsAction_GetByID(id string) (int, dm.CredentialsAction, error) {

	_, credentialsactionItem, _ := CredentialsAction_GetByID_impl(id)

	credentialsactionItem = CredentialsAction_PostGet(credentialsactionItem, id)

	return 1, credentialsactionItem, nil
}

func CredentialsAction_PostGet(credentialsactionItem dm.CredentialsAction, id string) dm.CredentialsAction {

	return credentialsactionItem
}

// CredentialsAction_DeleteByID() deletes a single CredentialsAction record
func CredentialsAction_Delete(id string) {
	CredentialsAction_Delete_impl(id)
}

// CredentialsAction_HardDelete(id string) soft deletes a single CredentialsAction record
func CredentialsAction_HardDelete(id string) {
	// Uses Hard Delete
	object_Table := CredentialsAction_QualifiedName
	tsql := das.DELETE + das.FROM + object_Table
	tsql = tsql + " " + das.WHERE + dm.CredentialsAction_SQLSearchID + das.EQ + das.ID(id)
	das.Execute(tsql)
	//if err != nil {
	//	logs.Error("CredentialsAction_SoftDelete()",err)
	//}
}

// CredentialsAction_Store() saves/stores a CredentialsAction record to the database
func CredentialsAction_Store(r dm.CredentialsAction, req *http.Request) (dm.CredentialsAction, error) {

	r, err := CredentialsAction_Validate(r)
	if err == nil {
		err = credentialsaction_Save(r, Audit_User(req))
	} else {
		logs.Information("CredentialsAction_Store()", err.Error())
	}

	return r, err
}

// CredentialsAction_StoreSystem() saves/stores a CredentialsAction record to the database
func CredentialsAction_StoreSystem(r dm.CredentialsAction) (dm.CredentialsAction, error) {

	r, err := CredentialsAction_Validate(r)
	if err == nil {
		err = credentialsaction_Save(r, Audit_Host())
	} else {
		logs.Information("CredentialsAction_Store()", err.Error())
	}

	return r, err
}

// CredentialsAction_StoreProcess() saves/stores a CredentialsAction record to the database
func CredentialsAction_StoreProcess(r dm.CredentialsAction, operator string) (dm.CredentialsAction, error) {

	r, err := CredentialsAction_Validate(r)
	if err == nil {
		err = credentialsaction_Save(r, operator)
	} else {
		logs.Information("CredentialsAction_StoreProcess()", err.Error())
	}

	return r, err
}

// CredentialsAction_Validate() validates for saves/stores a CredentialsAction record to the database
func CredentialsAction_Validate(r dm.CredentialsAction) (dm.CredentialsAction, error) {
	var err error

	// Cross Validation
	var errVal error
	r, _, errVal = CredentialsAction_ObjectValidation_impl(PUT, r.ID, r)
	if errVal != nil {
		err = errVal
	}
	return r, err
}

//

// credentialsaction_Save() saves/stores a CredentialsAction record to the database
func credentialsaction_Save(r dm.CredentialsAction, usr string) error {

	var err error

	if len(r.ID) == 0 {
		r.ID = CredentialsAction_NewID(r)
	}

	// If there are fields below, create the methods in dao\credentialsaction_impl.go

	logs.Storing("CredentialsAction", fmt.Sprintf("%v", r))

	// Please Create Functions Below in the adaptor/CredentialsAction_impl.go file
	err1 := CredentialsAction_Delete_impl(r.ID)
	err2 := CredentialsAction_Update_impl(r.ID, r, usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}

	return err

}

// credentialsaction_Fetch is not required as GetByID, GetAll etc... have been diverted to _impl

func CredentialsAction_NewID(r dm.CredentialsAction) string {

	id := uuid.New().String()

	return id
}

// credentialsaction_Fetch read all CredentialsAction's
func CredentialsAction_New() (int, []dm.CredentialsAction, dm.CredentialsAction, error) {

	var r = dm.CredentialsAction{}
	var rList []dm.CredentialsAction

	// START
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//

	//
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
	rList = append(rList, r)
	return 1, rList, r, nil
}
