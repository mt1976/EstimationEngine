package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/featurenew.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : FeatureNew (featurenew)
// Endpoint 	        : FeatureNew (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 10/12/2022 at 21:40:38
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	adaptor "github.com/mt1976/ebEstimates/adaptor"
	core "github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

var FeatureNew_SQLbase string

func init() {
	FeatureNew_SQLbase = core.DB_SELECT + " " + core.DB_ALL + " " + core.DB_FROM + " " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.FeatureNew_SQLTable)
}

// FeatureNew_GetList() returns a list of all FeatureNew records
func FeatureNew_GetList() (int, []dm.FeatureNew, error) {

	count, featurenewList, _ := adaptor.FeatureNew_GetList_impl()

	return count, featurenewList, nil
}

// FeatureNew_GetByID() returns a single FeatureNew record
func FeatureNew_GetByID(id string) (int, dm.FeatureNew, error) {

	_, featurenewItem, _ := adaptor.FeatureNew_GetByID_impl(id)

	// START
	// Dynamically generated 10/12/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 10/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return 1, featurenewItem, nil
}

// FeatureNew_DeleteByID() deletes a single FeatureNew record
func FeatureNew_Delete(id string) {

	adaptor.FeatureNew_Delete_impl(id)

}

// FeatureNew_Store() saves/stores a FeatureNew record to the database
func FeatureNew_Store(r dm.FeatureNew, req *http.Request) error {

	r, err := FeatureNew_Validate(r)
	if err == nil {
		err = featurenew_Save(r, Audit_User(req))
	} else {
		logs.Information("FeatureNew_Store()", err.Error())
	}

	return err
}

// FeatureNew_StoreSystem() saves/stores a FeatureNew record to the database
func FeatureNew_StoreSystem(r dm.FeatureNew) error {

	r, err := FeatureNew_Validate(r)
	if err == nil {
		err = featurenew_Save(r, Audit_Host())
	} else {
		logs.Information("FeatureNew_Store()", err.Error())
	}

	return err
}

// FeatureNew_Validate() validates for saves/stores a FeatureNew record to the database
func FeatureNew_Validate(r dm.FeatureNew) (dm.FeatureNew, error) {
	var err error
	// START
	// Dynamically generated 10/12/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 10/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//

	return r, err
}

//

// featurenew_Save() saves/stores a FeatureNew record to the database
func featurenew_Save(r dm.FeatureNew, usr string) error {

	var err error

	if len(r.ID) == 0 {
		r.ID = FeatureNew_NewID(r)
	}

	// If there are fields below, create the methods in dao\featurenew_impl.go

	logs.Storing("FeatureNew", fmt.Sprintf("%v", r))

	// Please Create Functions Below in the adaptor/FeatureNew_impl.go file
	err1 := adaptor.FeatureNew_Delete_impl(r.ID)
	err2 := adaptor.FeatureNew_Update_impl(r.ID, r, usr)
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
	// Dynamically generated 10/12/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 10/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	rList = append(rList, r)

	return 1, rList, r, nil
}
