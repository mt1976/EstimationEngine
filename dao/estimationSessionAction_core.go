package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/estimationsessionaction.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : EstimationSessionAction (estimationsessionaction)
// Endpoint 	        : EstimationSessionAction (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 04/03/2023 at 20:14:11
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

var EstimationSessionAction_SQLbase string
var EstimationSessionAction_QualifiedName string

func init() {
	EstimationSessionAction_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.EstimationSessionAction_SQLTable)
	EstimationSessionAction_SQLbase = das.SELECTALL + das.FROM + EstimationSessionAction_QualifiedName
}

// EstimationSessionAction_GetList() returns a list of all EstimationSessionAction records
func EstimationSessionAction_GetList() (int, []dm.EstimationSessionAction, error) {
	count, estimationsessionactionList, err := EstimationSessionAction_GetListFiltered("")
	return count, estimationsessionactionList, err
}

// EstimationSessionAction_GetListFiltered() returns a filtered list of all EstimationSessionAction records
func EstimationSessionAction_GetListFiltered(filter string) (int, []dm.EstimationSessionAction, error) {

	count, estimationsessionactionList, _ := EstimationSessionAction_GetList_impl(filter)

	return count, estimationsessionactionList, nil
}

// EstimationSessionAction_GetByID() returns a single EstimationSessionAction record
func EstimationSessionAction_GetByID(id string) (int, dm.EstimationSessionAction, error) {

	_, estimationsessionactionItem, _ := EstimationSessionAction_GetByID_impl(id)

	estimationsessionactionItem = EstimationSessionAction_PostGet(estimationsessionactionItem, id)

	return 1, estimationsessionactionItem, nil
}

func EstimationSessionAction_PostGet(estimationsessionactionItem dm.EstimationSessionAction, id string) dm.EstimationSessionAction {
	// START
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
	return estimationsessionactionItem
}

// EstimationSessionAction_DeleteByID() deletes a single EstimationSessionAction record
func EstimationSessionAction_Delete(id string) {
	EstimationSessionAction_Delete_impl(id)
}

// EstimationSessionAction_HardDelete(id string) soft deletes a single EstimationSessionAction record
func EstimationSessionAction_HardDelete(id string) {
	// Uses Hard Delete
	object_Table := EstimationSessionAction_QualifiedName
	tsql := das.DELETE + das.FROM + object_Table
	tsql = tsql + " " + das.WHERE + dm.EstimationSessionAction_SQLSearchID + das.EQ + das.ID(id)
	das.Execute(tsql)
	//if err != nil {
	//	logs.Error("EstimationSessionAction_SoftDelete()",err)
	//}
}

// EstimationSessionAction_Store() saves/stores a EstimationSessionAction record to the database
func EstimationSessionAction_Store(r dm.EstimationSessionAction, req *http.Request) (dm.EstimationSessionAction, error) {

	r, err := EstimationSessionAction_Validate(r)
	if err == nil {
		err = estimationsessionaction_Save(r, Audit_User(req))
	} else {
		logs.Information("EstimationSessionAction_Store()", err.Error())
	}

	return r, err
}

// EstimationSessionAction_StoreSystem() saves/stores a EstimationSessionAction record to the database
func EstimationSessionAction_StoreSystem(r dm.EstimationSessionAction) (dm.EstimationSessionAction, error) {

	r, err := EstimationSessionAction_Validate(r)
	if err == nil {
		err = estimationsessionaction_Save(r, Audit_Host())
	} else {
		logs.Information("EstimationSessionAction_Store()", err.Error())
	}

	return r, err
}

// EstimationSessionAction_StoreProcess() saves/stores a EstimationSessionAction record to the database
func EstimationSessionAction_StoreProcess(r dm.EstimationSessionAction, operator string) (dm.EstimationSessionAction, error) {

	r, err := EstimationSessionAction_Validate(r)
	if err == nil {
		err = estimationsessionaction_Save(r, operator)
	} else {
		logs.Information("EstimationSessionAction_StoreProcess()", err.Error())
	}

	return r, err
}

// EstimationSessionAction_Validate() validates for saves/stores a EstimationSessionAction record to the database
func EstimationSessionAction_Validate(r dm.EstimationSessionAction) (dm.EstimationSessionAction, error) {
	var err error
	// START
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
	//

	return r, err
}

//

// estimationsessionaction_Save() saves/stores a EstimationSessionAction record to the database
func estimationsessionaction_Save(r dm.EstimationSessionAction, usr string) error {

	var err error

	if len(r.ID) == 0 {
		r.ID = EstimationSessionAction_NewID(r)
	}

	// If there are fields below, create the methods in dao\estimationsessionaction_impl.go

	logs.Storing("EstimationSessionAction", fmt.Sprintf("%v", r))

	// Please Create Functions Below in the adaptor/EstimationSessionAction_impl.go file
	err1 := EstimationSessionAction_Delete_impl(r.ID)
	err2 := EstimationSessionAction_Update_impl(r.ID, r, usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}

	return err

}

// estimationsessionaction_Fetch is not required as GetByID, GetAll etc... have been diverted to _impl

func EstimationSessionAction_NewID(r dm.EstimationSessionAction) string {

	id := uuid.New().String()

	return id
}

// estimationsessionaction_Fetch read all EstimationSessionAction's
func EstimationSessionAction_New() (int, []dm.EstimationSessionAction, dm.EstimationSessionAction, error) {

	var r = dm.EstimationSessionAction{}
	var rList []dm.EstimationSessionAction

	// START
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//

	//
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
	rList = append(rList, r)
	return 1, rList, r, nil
}
