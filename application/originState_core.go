package application
// ----------------------------------------------------------------
// Automatically generated  "/application/originstate.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : OriginState (originstate)
// Endpoint 	        : OriginState (OriginStateID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 11/12/2022 at 19:24:01
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)

//OriginState_Publish annouces the endpoints available for this object
//OriginState_Publish - Auto generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local
func OriginState_Publish(mux http.ServeMux) {
	// START
	// Auto generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// 
	mux.HandleFunc(dm.OriginState_Path, OriginState_Handler)
	mux.HandleFunc(dm.OriginState_PathList, OriginState_HandlerList)
	mux.HandleFunc(dm.OriginState_PathView, OriginState_HandlerView)
	mux.HandleFunc(dm.OriginState_PathEdit, OriginState_HandlerEdit)
	mux.HandleFunc(dm.OriginState_PathNew, OriginState_HandlerNew)
	mux.HandleFunc(dm.OriginState_PathSave, OriginState_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.OriginState_Title)
    core.Catalog_Add(dm.OriginState_Title, dm.OriginState_Path, "", dm.OriginState_QueryString, "Application")
	// 
	// Auto generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//OriginState_HandlerList is the handler for the list page
//Allows Listing of OriginState records
//OriginState_HandlerList - Auto generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local
func OriginState_HandlerList(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local
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

	var returnList []dm.OriginState
	noItems, returnList, _ := dao.OriginState_GetList()

	pageDetail := dm.OriginState_PageList{
		Title:            CardTitle(dm.OriginState_Title, core.Action_List),
		PageTitle:        PageTitle(dm.OriginState_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.OriginState_TemplateList, w, r, pageDetail)
	// 
	// Auto generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

}


//OriginState_HandlerView is the handler used to View a page
//Allows Viewing for an existing OriginState record
//OriginState_HandlerView - Auto generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
func OriginState_HandlerView(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local
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

	searchID := core.GetURLparam(r, dm.OriginState_QueryString)
	_, rD, _ := dao.OriginState_GetByID(searchID)

	pageDetail := dm.OriginState_Page{
		Title:       CardTitle(dm.OriginState_Title, core.Action_View),
		PageTitle:   PageTitle(dm.OriginState_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = originstate_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.OriginState_TemplateView, w, r, pageDetail)
	// 
	// Auto generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//OriginState_HandlerEdit is the handler used generate the Edit page
//Allows Editing for an existing OriginState record and then allows the user to save the changes
//OriginState_HandlerEdit - Auto generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
func OriginState_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.OriginState_QueryString)
	_, rD, _ := dao.OriginState_GetByID(searchID)
	
	pageDetail := dm.OriginState_Page{
		Title:       CardTitle(dm.OriginState_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.OriginState_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = originstate_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.OriginState_TemplateEdit, w, r, pageDetail)
	// 
	// Auto generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//OriginState_HandlerSave is the handler used process the saving of an OriginState
//It is called from the Edit and New pages
//OriginState_HandlerSave  - Auto generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
func OriginState_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// 
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("OriginStateID"))

	item := originstate_DataFromRequest(r)
	
	dao.OriginState_Store(item,r)	
	http.Redirect(w, r, dm.OriginState_Redirect, http.StatusFound)
	// 
	// Auto generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//OriginState_HandlerNew is the handler used process the creation of an OriginState
//It will create a new OriginState and then redirect to the Edit page
//OriginState_HandlerNew  - Auto generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
func OriginState_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
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
	_, _, rD, _ := dao.OriginState_New()

	pageDetail := dm.OriginState_Page{
		Title:       CardTitle(dm.OriginState_Title, core.Action_New),
		PageTitle:   PageTitle(dm.OriginState_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = originstate_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.OriginState_TemplateNew, w, r, pageDetail)
	// 
	// Auto generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}	



//originstate_PopulatePage Builds/Populates the OriginState Page 
func originstate_PopulatePage(rD dm.OriginState, pageDetail dm.OriginState_Page) dm.OriginState_Page {
	// START
	// Auto generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.OriginStateID = rD.OriginStateID
	pageDetail.Code = rD.Code
	pageDetail.Name = rD.Name
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.SYSDeleted = rD.SYSDeleted
	pageDetail.SYSDeletedBy = rD.SYSDeletedBy
	pageDetail.SYSDeletedHost = rD.SYSDeletedHost
	pageDetail.SYSDbVersion = rD.SYSDbVersion
	
	
	//
	// Automatically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.OriginStateID_props = rD.OriginStateID_props
	pageDetail.Code_props = rD.Code_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.SYSDeleted_props = rD.SYSDeleted_props
	pageDetail.SYSDeletedBy_props = rD.SYSDeletedBy_props
	pageDetail.SYSDeletedHost_props = rD.SYSDeletedHost_props
	pageDetail.SYSDbVersion_props = rD.SYSDbVersion_props
	
	// 
	// Auto generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	


//originstate_DataFromRequest is used process the content of an HTTP Request and return an instance of an OriginState
func originstate_DataFromRequest(r *http.Request) dm.OriginState {
	// START
	// Auto generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	var item dm.OriginState
	// FIELD SET START
	// Auto generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.OriginState_SYSId_scrn)
		item.OriginStateID = r.FormValue(dm.OriginState_OriginStateID_scrn)
		item.Code = r.FormValue(dm.OriginState_Code_scrn)
		item.Name = r.FormValue(dm.OriginState_Name_scrn)
		item.SYSCreated = r.FormValue(dm.OriginState_SYSCreated_scrn)
		item.SYSCreatedBy = r.FormValue(dm.OriginState_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.OriginState_SYSCreatedHost_scrn)
		item.SYSUpdated = r.FormValue(dm.OriginState_SYSUpdated_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.OriginState_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.OriginState_SYSUpdatedHost_scrn)
		item.SYSDeleted = r.FormValue(dm.OriginState_SYSDeleted_scrn)
		item.SYSDeletedBy = r.FormValue(dm.OriginState_SYSDeletedBy_scrn)
		item.SYSDeletedHost = r.FormValue(dm.OriginState_SYSDeletedHost_scrn)
		item.SYSDbVersion = r.FormValue(dm.OriginState_SYSDbVersion_scrn)
	
	// 
	// Auto generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return item
}
