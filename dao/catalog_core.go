package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/catalog.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Catalog (catalog)
// Endpoint 	        : Catalog (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:44
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"

	adaptor "github.com/mt1976/ebEstimates/adaptor"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// Catalog_GetList() returns a list of all Catalog records
func Catalog_GetList() (int, []dm.Catalog, error) {

	count, catalogList, _ := adaptor.Catalog_GetList_impl()

	return count, catalogList, nil
}

// Catalog_GetByID() returns a single Catalog record
func Catalog_GetByID(id string) (int, dm.Catalog, error) {

	_, catalogItem, _ := adaptor.Catalog_GetByID_impl(id)

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return 1, catalogItem, nil
}

// Catalog_DeleteByID() deletes a single Catalog record
func Catalog_Delete(id string) {

	adaptor.Catalog_Delete_impl(id)

}

// Catalog_Store() saves/stores a Catalog record to the database
func Catalog_Store(r dm.Catalog, req *http.Request) error {

	err, r := Catalog_Validate(r)
	if err == nil {
		err = catalog_Save(r, Audit_User(req))
	} else {
		logs.Information("Catalog_Store()", err.Error())
	}

	return err
}

// Catalog_StoreSystem() saves/stores a Catalog record to the database
func Catalog_StoreSystem(r dm.Catalog) error {

	err, r := Catalog_Validate(r)
	if err == nil {
		err = catalog_Save(r, Audit_Host())
	} else {
		logs.Information("Catalog_Store()", err.Error())
	}

	return err
}

// Catalog_Validate() validates for saves/stores a Catalog record to the database
func Catalog_Validate(r dm.Catalog) (error, dm.Catalog) {
	var err error
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//

	return err, r
}

//

// catalog_Save() saves/stores a Catalog record to the database
func catalog_Save(r dm.Catalog, usr string) error {

	var err error

	if len(r.ID) == 0 {
		r.ID = Catalog_NewID(r)
	}

	// If there are fields below, create the methods in dao\catalog_impl.go

	logs.Storing("Catalog", fmt.Sprintf("%s", r))

	// Please Create Functions Below in the adaptor/Catalog_impl.go file
	err1 := adaptor.Catalog_Delete_impl(r.ID)
	err2 := adaptor.Catalog_Update_impl(r.ID, r, usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}

	return err

}

// catalog_Fetch is not required as GetByID, GetAll etc... have been diverted to _impl

func Catalog_NewID(r dm.Catalog) string {

	id := uuid.New().String()

	return id
}

// catalog_Fetch read all Catalog's
func Catalog_New() (int, []dm.Catalog, dm.Catalog, error) {

	var r = dm.Catalog{}
	var rList []dm.Catalog

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	rList = append(rList, r)

	return 1, rList, r, nil
}
