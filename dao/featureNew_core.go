package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/featurenew.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : FeatureNew (featurenew)
// Endpoint 	        : FeatureNew (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 03/03/2023 at 13:56:51
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

var FeatureNew_SQLbase string
var FeatureNew_QualifiedName string
func init(){
	FeatureNew_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.FeatureNew_SQLTable)
	FeatureNew_SQLbase =  das.SELECTALL + das.FROM + FeatureNew_QualifiedName
}

// FeatureNew_GetList() returns a list of all FeatureNew records
func FeatureNew_GetList() (int, []dm.FeatureNew, error) {
	count, featurenewList, err := FeatureNew_GetListFiltered("")
	return count, featurenewList, err
}

// FeatureNew_GetListFiltered() returns a filtered list of all FeatureNew records
func FeatureNew_GetListFiltered(filter string) (int, []dm.FeatureNew, error) {
	
	count, featurenewList, _ := FeatureNew_GetList_impl(filter)
	
	return count, featurenewList, nil
}



// FeatureNew_GetByID() returns a single FeatureNew record
func FeatureNew_GetByID(id string) (int, dm.FeatureNew, error) {


	 _, featurenewItem, _ := FeatureNew_GetByID_impl(id)
	

	featurenewItem = FeatureNew_PostGet(featurenewItem,id)

	return 1, featurenewItem, nil
}

func FeatureNew_PostGet(featurenewItem dm.FeatureNew,id string) dm.FeatureNew {
	// START
	// Dynamically generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return featurenewItem
}




// FeatureNew_DeleteByID() deletes a single FeatureNew record
func FeatureNew_Delete(id string) {


	FeatureNew_Delete_impl(id)
	
	
}

	


// FeatureNew_Store() saves/stores a FeatureNew record to the database
func FeatureNew_Store(r dm.FeatureNew,req *http.Request) (dm.FeatureNew,error) {

	r, err := FeatureNew_Validate(r)
	if err == nil {
		err = featurenew_Save(r, Audit_User(req))
	} else {
		logs.Information("FeatureNew_Store()", err.Error())
	}

	return r, err
}

// FeatureNew_StoreSystem() saves/stores a FeatureNew record to the database
func FeatureNew_StoreSystem(r dm.FeatureNew) (dm.FeatureNew,error) {
	
	r, err := FeatureNew_Validate(r)
	if err == nil {
		err = featurenew_Save(r, Audit_Host())
	} else {
		logs.Information("FeatureNew_Store()", err.Error())
	}

	return r, err
}

// FeatureNew_Validate() validates for saves/stores a FeatureNew record to the database
func FeatureNew_Validate(r dm.FeatureNew) (dm.FeatureNew, error) {
	var err error
	// START
	// Dynamically generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	

	return r,err
}
//

// featurenew_Save() saves/stores a FeatureNew record to the database
func featurenew_Save(r dm.FeatureNew,usr string) error {

    var err error

	if len(r.ID) == 0 {
		r.ID = FeatureNew_NewID(r)
	}

// If there are fields below, create the methods in dao\featurenew_impl.go

logs.Storing("FeatureNew",fmt.Sprintf("%v", r))

// Please Create Functions Below in the adaptor/FeatureNew_impl.go file
	err1 := FeatureNew_Delete_impl(r.ID)
	err2 := FeatureNew_Update_impl(r.ID,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



	// featurenew_Fetch is not required as GetByID, GetAll etc... have been diverted to _impl
	


func FeatureNew_NewID(r dm.FeatureNew) string {
	
	id := uuid.New().String()
	
	return id
}



// featurenew_Fetch read all FeatureNew's
func FeatureNew_New() (int, []dm.FeatureNew, dm.FeatureNew, error) {

	var r = dm.FeatureNew{}
	var rList []dm.FeatureNew
	

	// START
	// Dynamically generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	
	// 
	// Dynamically generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	rList = append(rList, r)
	return 1, rList, r, nil
}