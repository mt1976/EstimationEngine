package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/catalog.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Catalog (catalog)
// Endpoint 	        : Catalog (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:44
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"github.com/google/uuid"
	adaptor "github.com/mt1976/ebEstimates/adaptor"
	dm "github.com/mt1976/ebEstimates/datamodel"
)

// Catalog_GetList() returns a list of all Catalog records
func Catalog_GetList_impl() (int, []dm.Catalog, error) {

	count, catalogList, _ := adaptor.Catalog_GetList_impl()

	return count, catalogList, nil
}

// Catalog_GetByID() returns a single Catalog record
func Catalog_GetByID_impl(id string) (int, dm.Catalog, error) {

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
func Catalog_Delete_impl(id string) {

	adaptor.Catalog_Delete_impl(id)

}

// catalog_Fetch is not required as GetByID, GetAll etc... have been diverted to _impl

func Catalog_NewID_impl(r dm.Catalog) string {

	id := uuid.New().String()

	return id
}
