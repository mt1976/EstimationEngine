package dao

import (
	core "github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/catalog.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : Catalog (catalog)
// Endpoint 	        : Catalog (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 03/01/2023 at 19:18:07
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in catalog_impl.go

func Catalog_NewID_impl(rec dm.Catalog) string { return rec.ID }
func Catalog_Delete_impl(id string) error      { return nil }

func Catalog_Update_impl(id string, rec dm.Catalog, usr string) error { return nil }

// If there are fields below, create the methods in adaptor\catalog_impl.go
// START - GET API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - GET API/Callout
//
// START - PUT API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - PUT API/Callout

// START - Validation API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout

// Catalog_GetList() returns a list of all API records
func Catalog_GetList_impl() (int, []dm.Catalog, error) {

	//	count, apiList, _, _ := adaptor.Catalog_GetList_impl()
	var apiList []dm.Catalog
	count := len(core.Catalog)
	for _, v := range core.Catalog {
		apiList = append(apiList, dm.Catalog{ID: v.ID, Descr: v.Descr, Endpoint: v.Path, Query: v.Query, Source: v.Source})
	}
	return count, apiList, nil
}

// Catalog_GetByID() returns a single API record
func Catalog_GetByID_impl(id string) (int, dm.Catalog, error) {

	//_, _, apiItem, _ := adaptor.Catalog_GetByID_impl(id)
	var apiItem dm.Catalog

	for _, v := range core.Catalog {
		if v.ID == id {
			apiItem = dm.Catalog{ID: v.ID, Descr: v.Descr, Endpoint: v.Path, Query: v.Query, Source: v.Source}
		}
	}
	//fmt.Printf("apiItem: %v\n", apiItem)
	return 1, apiItem, nil
}
