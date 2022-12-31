package application
// ----------------------------------------------------------------
// Automatically generated  "/application/estimationstate.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : EstimationState (estimationstate)
// Endpoint 	        : EstimationState (EstimationStateID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 30/12/2022 at 12:18:20
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)

//EstimationState_Publish annouces the endpoints available for this object
//EstimationState_Publish - Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
func EstimationState_Publish(mux http.ServeMux) {
	// START
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// 
	mux.HandleFunc(dm.EstimationState_Path, EstimationState_Handler)
	mux.HandleFunc(dm.EstimationState_PathList, EstimationState_HandlerList)
	mux.HandleFunc(dm.EstimationState_PathView, EstimationState_HandlerView)
	mux.HandleFunc(dm.EstimationState_PathEdit, EstimationState_HandlerEdit)
	mux.HandleFunc(dm.EstimationState_PathNew, EstimationState_HandlerNew)
	mux.HandleFunc(dm.EstimationState_PathSave, EstimationState_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.EstimationState_Title)
    core.Catalog_Add(dm.EstimationState_Title, dm.EstimationState_Path, "", dm.EstimationState_QueryString, "Application")
	// 
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//EstimationState_HandlerList is the handler for the list page
//Allows Listing of EstimationState records
//EstimationState_HandlerList - Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
func EstimationState_HandlerList(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
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

	var returnList []dm.EstimationState
	noItems, returnList, _ := dao.EstimationState_GetList()

	pageDetail := dm.EstimationState_PageList{
		Title:            CardTitle(dm.EstimationState_Title, core.Action_List),
		PageTitle:        PageTitle(dm.EstimationState_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.EstimationState_TemplateList, w, r, pageDetail)
	// 
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

}


//EstimationState_HandlerView is the handler used to View a page
//Allows Viewing for an existing EstimationState record
//EstimationState_HandlerView - Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
func EstimationState_HandlerView(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
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

	searchID := core.GetURLparam(r, dm.EstimationState_QueryString)
	_, rD, _ := dao.EstimationState_GetByID(searchID)

	pageDetail := dm.EstimationState_Page{
		Title:       CardTitle(dm.EstimationState_Title, core.Action_View),
		PageTitle:   PageTitle(dm.EstimationState_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = estimationstate_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.EstimationState_TemplateView, w, r, pageDetail)
	// 
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//EstimationState_HandlerEdit is the handler used generate the Edit page
//Allows Editing for an existing EstimationState record and then allows the user to save the changes
//EstimationState_HandlerEdit - Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
func EstimationState_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
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

	searchID := core.GetURLparam(r, dm.EstimationState_QueryString)
	_, rD, _ := dao.EstimationState_GetByID(searchID)
	
	pageDetail := dm.EstimationState_Page{
		Title:       CardTitle(dm.EstimationState_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.EstimationState_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = estimationstate_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.EstimationState_TemplateEdit, w, r, pageDetail)
	// 
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//EstimationState_HandlerSave is the handler used process the saving of an EstimationState
//It is called from the Edit and New pages
//EstimationState_HandlerSave  - Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
func EstimationState_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// 
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("EstimationStateID"))

	item := estimationstate_DataFromRequest(r)
	
	dao.EstimationState_Store(item,r)	
	http.Redirect(w, r, dm.EstimationState_Redirect, http.StatusFound)
	// 
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//EstimationState_HandlerNew is the handler used process the creation of an EstimationState
//It will create a new EstimationState and then redirect to the Edit page
//EstimationState_HandlerNew  - Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
func EstimationState_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
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
	_, _, rD, _ := dao.EstimationState_New()

	pageDetail := dm.EstimationState_Page{
		Title:       CardTitle(dm.EstimationState_Title, core.Action_New),
		PageTitle:   PageTitle(dm.EstimationState_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = estimationstate_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.EstimationState_TemplateNew, w, r, pageDetail)
	// 
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}	



//estimationstate_PopulatePage Builds/Populates the EstimationState Page 
func estimationstate_PopulatePage(rD dm.EstimationState, pageDetail dm.EstimationState_Page) dm.EstimationState_Page {
	// START
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.EstimationStateID = rD.EstimationStateID
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
	// Automatically generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.EstimationStateID_props = rD.EstimationStateID_props
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
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	


//estimationstate_DataFromRequest is used process the content of an HTTP Request and return an instance of an EstimationState
func estimationstate_DataFromRequest(r *http.Request) dm.EstimationState {
	// START
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	var item dm.EstimationState
	// FIELD SET START
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.EstimationState_SYSId_scrn)
		item.EstimationStateID = r.FormValue(dm.EstimationState_EstimationStateID_scrn)
		item.Code = r.FormValue(dm.EstimationState_Code_scrn)
		item.Name = r.FormValue(dm.EstimationState_Name_scrn)
		item.SYSCreated = r.FormValue(dm.EstimationState_SYSCreated_scrn)
		item.SYSCreatedBy = r.FormValue(dm.EstimationState_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.EstimationState_SYSCreatedHost_scrn)
		item.SYSUpdated = r.FormValue(dm.EstimationState_SYSUpdated_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.EstimationState_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.EstimationState_SYSUpdatedHost_scrn)
		item.SYSDeleted = r.FormValue(dm.EstimationState_SYSDeleted_scrn)
		item.SYSDeletedBy = r.FormValue(dm.EstimationState_SYSDeletedBy_scrn)
		item.SYSDeletedHost = r.FormValue(dm.EstimationState_SYSDeletedHost_scrn)
		item.SYSDbVersion = r.FormValue(dm.EstimationState_SYSDbVersion_scrn)
	
	// 
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return item
}
