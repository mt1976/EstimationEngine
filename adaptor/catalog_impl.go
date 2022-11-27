package adaptor

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	"github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
	"github.com/mt1976/ebEstimates/logs"
)

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

// Catalog_DeleteByID() deletes a single API record
func Catalog_Delete_impl(id string) error {
	logs.Information("Catalog_Delete_impl()", "NO ACTION")
	//adaptor.Catalog_Delete_impl(id)
	return nil
}

// Catalog_DeleteByID() deletes a single API record
func Catalog_Update_impl(id string, item dm.Catalog, usr string) error {
	logs.Information("Catalog_Delete_impl()", "NO ACTION")
	//adaptor.Catalog_Delete_impl(id)
	return nil
}

// Catalog_DeleteByID() deletes a single API record
func Catalog_NewID_impl(r dm.Catalog) string {
	logs.Information("Catalog_Delete_impl()", "NO ACTION")
	//adaptor.Catalog_Delete_impl(id)
	return r.ID
}
