package application
// ----------------------------------------------------------------
// Automatically generated  "/application/catalog.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Catalog (catalog)
// Endpoint 	        : Catalog (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 08/03/2023 at 18:42:22
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)

//Catalog_Publish annouces the endpoints available for this object
//Catalog_Publish - Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local
func Catalog_Publish(mux http.ServeMux) {
	// START
	// Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	mux.HandleFunc(dm.Catalog_Path, Catalog_Handler)
	mux.HandleFunc(dm.Catalog_PathList, Catalog_HandlerList)
	mux.HandleFunc(dm.Catalog_PathView, Catalog_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Catalog_Title)
    core.API.AddRoute(dm.Catalog_Title, dm.Catalog_Path, "", dm.Catalog_QueryString, "Application")
	// 
	// Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//Catalog_HandlerList is the handler for the list page
//Allows Listing of Catalog records
//Catalog_HandlerList - Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local
func Catalog_HandlerList(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Catalog

	objectName := dao.Translate("ObjectName", "Catalog")
	reqField := "Base"
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule)
	if filter == "" {
		logs.Warning("No filter found : " + reqField + " for Object: " + objectName)
	} 

	noItems, returnList, _ := dao.Catalog_GetListFiltered(filter)

	pageDetail := dm.Catalog_PageList{
		Title:            CardTitle(dm.Catalog_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Catalog_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	nextTemplate :=  NextTemplate("Catalog", "List", dm.Catalog_TemplateList)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END

}


//Catalog_HandlerView is the handler used to View a page
//Allows Viewing for an existing Catalog record
//Catalog_HandlerView - Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func Catalog_HandlerView(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Catalog_QueryString)
	_, rD, _ := dao.Catalog_GetByID(searchID)

	pageDetail := dm.Catalog_Page{
		Title:       CardTitle(dm.Catalog_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Catalog_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = catalog_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Catalog", "View", dm.Catalog_TemplateView)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}






//catalog_PopulatePage Builds/Populates the Catalog Page 
//catalog_PopulatePage Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func catalog_PopulatePage(rD dm.Catalog, pageDetail dm.Catalog_Page) dm.Catalog_Page {
	// Real DB Fields
	pageDetail.ID = rD.ID
	pageDetail.Endpoint = rD.Endpoint
	pageDetail.Descr = rD.Descr
	pageDetail.Query = rD.Query
	pageDetail.Source = rD.Source
	// Add Pseudo/Extra Fields
	// Enrichment Fields 
	pageDetail.ID_props = rD.ID_props
	pageDetail.Endpoint_props = rD.Endpoint_props
	pageDetail.Descr_props = rD.Descr_props
	pageDetail.Query_props = rD.Query_props
	pageDetail.Source_props = rD.Source_props
	return pageDetail
}