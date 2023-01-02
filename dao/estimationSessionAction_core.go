package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/estimationsessionaction.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : EstimationSessionAction (estimationsessionaction)
// Endpoint 	        : EstimationSessionAction (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 02/01/2023 at 15:17:13
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (

	"fmt"
	"net/http"
	core "github.com/mt1976/ebEstimates/core"
	"github.com/google/uuid"
	
	 adaptor   "github.com/mt1976/ebEstimates/adaptor"
	dm   "github.com/mt1976/ebEstimates/datamodel"
	logs   "github.com/mt1976/ebEstimates/logs"
)

var EstimationSessionAction_SQLbase string
var EstimationSessionAction_QualifiedName string
func init(){
	EstimationSessionAction_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.EstimationSessionAction_SQLTable)
	EstimationSessionAction_SQLbase =  core.DB_SELECT + " "+ core.DB_ALL + " " + core.DB_FROM + " " + EstimationSessionAction_QualifiedName
}

// EstimationSessionAction_GetList() returns a list of all EstimationSessionAction records
func EstimationSessionAction_GetList() (int, []dm.EstimationSessionAction, error) {
	
	count, estimationsessionactionList, _ := adaptor.EstimationSessionAction_GetList_impl()
	
	return count, estimationsessionactionList, nil
}



// EstimationSessionAction_GetByID() returns a single EstimationSessionAction record
func EstimationSessionAction_GetByID(id string) (int, dm.EstimationSessionAction, error) {


	 _, estimationsessionactionItem, _ := adaptor.EstimationSessionAction_GetByID_impl(id)
	
	// START
	// Dynamically generated 02/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 02/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, estimationsessionactionItem, nil
}



// EstimationSessionAction_DeleteByID() deletes a single EstimationSessionAction record
func EstimationSessionAction_Delete(id string) {


	adaptor.EstimationSessionAction_Delete_impl(id)
	
	
}














	


// EstimationSessionAction_Store() saves/stores a EstimationSessionAction record to the database
func EstimationSessionAction_Store(r dm.EstimationSessionAction,req *http.Request) error {

	r, err := EstimationSessionAction_Validate(r)
	if err == nil {
		err = estimationsessionaction_Save(r, Audit_User(req))
	} else {
		logs.Information("EstimationSessionAction_Store()", err.Error())
	}

	return err
}

// EstimationSessionAction_StoreSystem() saves/stores a EstimationSessionAction record to the database
func EstimationSessionAction_StoreSystem(r dm.EstimationSessionAction) error {
	
	r, err := EstimationSessionAction_Validate(r)
	if err == nil {
		err = estimationsessionaction_Save(r, Audit_Host())
	} else {
		logs.Information("EstimationSessionAction_Store()", err.Error())
	}

	return err
}

// EstimationSessionAction_Validate() validates for saves/stores a EstimationSessionAction record to the database
func EstimationSessionAction_Validate(r dm.EstimationSessionAction) (dm.EstimationSessionAction, error) {
	var err error
	// START
	// Dynamically generated 02/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 02/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	

	return r,err
}
//

// estimationsessionaction_Save() saves/stores a EstimationSessionAction record to the database
func estimationsessionaction_Save(r dm.EstimationSessionAction,usr string) error {

    var err error



	

	if len(r.ID) == 0 {
		r.ID = EstimationSessionAction_NewID(r)
	}

// If there are fields below, create the methods in dao\estimationsessionaction_impl.go






	
logs.Storing("EstimationSessionAction",fmt.Sprintf("%v", r))

// Please Create Functions Below in the adaptor/EstimationSessionAction_impl.go file
	err1 := adaptor.EstimationSessionAction_Delete_impl(r.ID)
	err2 := adaptor.EstimationSessionAction_Update_impl(r.ID,r,usr)
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
	// Dynamically generated 02/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 02/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}