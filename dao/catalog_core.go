package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/catalog.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Catalog (catalog)
// Endpoint 	        : Catalog (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 07/01/2023 at 23:01:28
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

var Catalog_SQLbase string
var Catalog_QualifiedName string
func init(){
	Catalog_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.Catalog_SQLTable)
	Catalog_SQLbase =  core.DB_SELECT + " "+ core.DB_ALL + " " + core.DB_FROM + " " + Catalog_QualifiedName
}

// Catalog_GetList() returns a list of all Catalog records
func Catalog_GetList() (int, []dm.Catalog, error) {
	
	tsql := Catalog_SQLbase
	count, catalogList, _, _ := catalog_Fetch(tsql)
	
	return count, catalogList, nil
}



// Catalog_GetByID() returns a single Catalog record
func Catalog_GetByID(id string) (int, dm.Catalog, error) {


	tsql := Catalog_SQLbase
	tsql = tsql + " " + core.DB_WHERE + " " + dm.Catalog_SQLSearchID + core.DB_EQ + "'" + id + "'"
	_, _, catalogItem, _ := catalog_Fetch(tsql)

	// START
	// Dynamically generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, catalogItem, nil
}



// Catalog_DeleteByID() deletes a single Catalog record
func Catalog_Delete(id string) {


	Catalog_Delete_impl(id)
	
	
}














	


// Catalog_Store() saves/stores a Catalog record to the database
func Catalog_Store(r dm.Catalog,req *http.Request) error {

	r, err := Catalog_Validate(r)
	if err == nil {
		err = catalog_Save(r, Audit_User(req))
	} else {
		logs.Information("Catalog_Store()", err.Error())
	}

	return err
}

// Catalog_StoreSystem() saves/stores a Catalog record to the database
func Catalog_StoreSystem(r dm.Catalog) error {
	
	r, err := Catalog_Validate(r)
	if err == nil {
		err = catalog_Save(r, Audit_Host())
	} else {
		logs.Information("Catalog_Store()", err.Error())
	}

	return err
}

// Catalog_Validate() validates for saves/stores a Catalog record to the database
func Catalog_Validate(r dm.Catalog) (dm.Catalog, error) {
	var err error
	// START
	// Dynamically generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	

	return r,err
}
//

// catalog_Save() saves/stores a Catalog record to the database
func catalog_Save(r dm.Catalog,usr string) error {

    var err error



	

	if len(r.ID) == 0 {
		r.ID = Catalog_NewID(r)
	}

// If there are fields below, create the methods in dao\catalog_impl.go







	
logs.Storing("Catalog",fmt.Sprintf("%v", r))

// Please Create Functions Below in the adaptor/Catalog_impl.go file
	err1 := Catalog_Delete_impl(r.ID)
	err2 := Catalog_Update_impl(r.ID,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// catalog_Fetch read all Catalog's
func catalog_Fetch(tsql string) (int, []dm.Catalog, dm.Catalog, error) {

	var recItem dm.Catalog
	var recList []dm.Catalog

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.ID  = get_String(rec, dm.Catalog_ID_sql, "")
	   recItem.Endpoint  = get_String(rec, dm.Catalog_Endpoint_sql, "")
	   recItem.Descr  = get_String(rec, dm.Catalog_Descr_sql, "")
	   recItem.Query  = get_String(rec, dm.Catalog_Query_sql, "")
	   recItem.Source  = get_String(rec, dm.Catalog_Source_sql, "")
	
	// If there are fields below, create the methods in adaptor\Catalog_impl.go
	
	
	
	
	
	
	// 
	// Dynamically generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func Catalog_NewID(r dm.Catalog) string {
	
			id := uuid.New().String()
	
	return id
}



// catalog_Fetch read all Catalog's
func Catalog_New() (int, []dm.Catalog, dm.Catalog, error) {

	var r = dm.Catalog{}
	var rList []dm.Catalog
	

	// START
	// Dynamically generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}