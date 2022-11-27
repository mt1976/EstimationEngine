package application
// ----------------------------------------------------------------
// Automatically generated  "/application/estimationsession.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : EstimationSession (estimationsession)
// Endpoint 	        : EstimationSession (EstimationSessionID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 27/11/2022 at 20:46:13
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)





//EstimationSession_Publish annouces the endpoints available for this object
func EstimationSession_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.EstimationSession_Path, EstimationSession_Handler)
	mux.HandleFunc(dm.EstimationSession_PathList, EstimationSession_HandlerList)
	mux.HandleFunc(dm.EstimationSession_PathView, EstimationSession_HandlerView)
	mux.HandleFunc(dm.EstimationSession_PathEdit, EstimationSession_HandlerEdit)
	mux.HandleFunc(dm.EstimationSession_PathNew, EstimationSession_HandlerNew)
	mux.HandleFunc(dm.EstimationSession_PathSave, EstimationSession_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.EstimationSession_Title)
    core.Catalog_Add(dm.EstimationSession_Title, dm.EstimationSession_Path, "", dm.EstimationSession_QueryString, "Application")
}


//EstimationSession_HandlerList is the handler for the list page
func EstimationSession_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.EstimationSession
	noItems, returnList, _ := dao.EstimationSession_GetList()

	pageDetail := dm.EstimationSession_PageList{
		Title:            CardTitle(dm.EstimationSession_Title, core.Action_List),
		PageTitle:        PageTitle(dm.EstimationSession_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.EstimationSession_TemplateList, w, r, pageDetail)

}


//EstimationSession_HandlerView is the handler used to View a page
func EstimationSession_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.EstimationSession_QueryString)
	_, rD, _ := dao.EstimationSession_GetByID(searchID)

	pageDetail := dm.EstimationSession_Page{
		Title:       CardTitle(dm.EstimationSession_Title, core.Action_View),
		PageTitle:   PageTitle(dm.EstimationSession_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = estimationsession_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.EstimationSession_TemplateView, w, r, pageDetail)

}


//EstimationSession_HandlerEdit is the handler used generate the Edit page
func EstimationSession_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.EstimationSession_QueryString)
	_, rD, _ := dao.EstimationSession_GetByID(searchID)
	
	pageDetail := dm.EstimationSession_Page{
		Title:       CardTitle(dm.EstimationSession_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.EstimationSession_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = estimationsession_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.EstimationSession_TemplateEdit, w, r, pageDetail)
}


//EstimationSession_HandlerSave is the handler used process the saving of an EstimationSession
func EstimationSession_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("EstimationSessionID"))

	var item dm.EstimationSession
	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.EstimationSession_SYSId_scrn)
		item.EstimationSessionID = r.FormValue(dm.EstimationSession_EstimationSessionID_scrn)
		item.ProjectID = r.FormValue(dm.EstimationSession_ProjectID_scrn)
		item.EstimationStateID = r.FormValue(dm.EstimationSession_EstimationStateID_scrn)
		item.Notes = r.FormValue(dm.EstimationSession_Notes_scrn)
		item.Releases = r.FormValue(dm.EstimationSession_Releases_scrn)
		item.Total = r.FormValue(dm.EstimationSession_Total_scrn)
		item.Contingency = r.FormValue(dm.EstimationSession_Contingency_scrn)
		item.ReqDays = r.FormValue(dm.EstimationSession_ReqDays_scrn)
		item.RegCost = r.FormValue(dm.EstimationSession_RegCost_scrn)
		item.ImpDays = r.FormValue(dm.EstimationSession_ImpDays_scrn)
		item.ImpCost = r.FormValue(dm.EstimationSession_ImpCost_scrn)
		item.UatDays = r.FormValue(dm.EstimationSession_UatDays_scrn)
		item.UatCost = r.FormValue(dm.EstimationSession_UatCost_scrn)
		item.MgtDays = r.FormValue(dm.EstimationSession_MgtDays_scrn)
		item.MgtCost = r.FormValue(dm.EstimationSession_MgtCost_scrn)
		item.RelDays = r.FormValue(dm.EstimationSession_RelDays_scrn)
		item.RelCost = r.FormValue(dm.EstimationSession_RelCost_scrn)
		item.SupportUplift = r.FormValue(dm.EstimationSession_SupportUplift_scrn)
		item.SYSCreated = r.FormValue(dm.EstimationSession_SYSCreated_scrn)
		item.SYSCreatedBy = r.FormValue(dm.EstimationSession_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.EstimationSession_SYSCreatedHost_scrn)
		item.SYSUpdated = r.FormValue(dm.EstimationSession_SYSUpdated_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.EstimationSession_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.EstimationSession_SYSUpdatedHost_scrn)
		item.SYSDeleted = r.FormValue(dm.EstimationSession_SYSDeleted_scrn)
		item.SYSDeletedBy = r.FormValue(dm.EstimationSession_SYSDeletedBy_scrn)
		item.SYSDeletedHost = r.FormValue(dm.EstimationSession_SYSDeletedHost_scrn)
		item.Name = r.FormValue(dm.EstimationSession_Name_scrn)
	
	// 
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.EstimationSession_Store(item,r)	
	http.Redirect(w, r, dm.EstimationSession_Redirect, http.StatusFound)
}


//EstimationSession_HandlerNew is the handler used process the creation of an EstimationSession
func EstimationSession_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.EstimationSession_New()

	pageDetail := dm.EstimationSession_Page{
		Title:       CardTitle(dm.EstimationSession_Title, core.Action_New),
		PageTitle:   PageTitle(dm.EstimationSession_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = estimationsession_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.EstimationSession_TemplateNew, w, r, pageDetail)

}	



// Builds/Popuplates the EstimationSession Page 
func estimationsession_PopulatePage(rD dm.EstimationSession, pageDetail dm.EstimationSession_Page) dm.EstimationSession_Page {
	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.EstimationSessionID = rD.EstimationSessionID
	pageDetail.ProjectID = rD.ProjectID
	pageDetail.EstimationStateID = rD.EstimationStateID
	pageDetail.Notes = rD.Notes
	pageDetail.Releases = rD.Releases
	pageDetail.Total = rD.Total
	pageDetail.Contingency = rD.Contingency
	pageDetail.ReqDays = rD.ReqDays
	pageDetail.RegCost = rD.RegCost
	pageDetail.ImpDays = rD.ImpDays
	pageDetail.ImpCost = rD.ImpCost
	pageDetail.UatDays = rD.UatDays
	pageDetail.UatCost = rD.UatCost
	pageDetail.MgtDays = rD.MgtDays
	pageDetail.MgtCost = rD.MgtCost
	pageDetail.RelDays = rD.RelDays
	pageDetail.RelCost = rD.RelCost
	pageDetail.SupportUplift = rD.SupportUplift
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.SYSDeleted = rD.SYSDeleted
	pageDetail.SYSDeletedBy = rD.SYSDeletedBy
	pageDetail.SYSDeletedHost = rD.SYSDeletedHost
	pageDetail.Name = rD.Name
	
	
	//
	// Automatically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	pageDetail.ProjectID_lookup = dao.Project_GetLookup()
	
	
	
	pageDetail.EstimationStateID_lookup = dao.EstimationState_GetLookup()
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.EstimationSessionID_props = rD.EstimationSessionID_props
	pageDetail.ProjectID_props = rD.ProjectID_props
	pageDetail.EstimationStateID_props = rD.EstimationStateID_props
	pageDetail.Notes_props = rD.Notes_props
	pageDetail.Releases_props = rD.Releases_props
	pageDetail.Total_props = rD.Total_props
	pageDetail.Contingency_props = rD.Contingency_props
	pageDetail.ReqDays_props = rD.ReqDays_props
	pageDetail.RegCost_props = rD.RegCost_props
	pageDetail.ImpDays_props = rD.ImpDays_props
	pageDetail.ImpCost_props = rD.ImpCost_props
	pageDetail.UatDays_props = rD.UatDays_props
	pageDetail.UatCost_props = rD.UatCost_props
	pageDetail.MgtDays_props = rD.MgtDays_props
	pageDetail.MgtCost_props = rD.MgtCost_props
	pageDetail.RelDays_props = rD.RelDays_props
	pageDetail.RelCost_props = rD.RelCost_props
	pageDetail.SupportUplift_props = rD.SupportUplift_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.SYSDeleted_props = rD.SYSDeleted_props
	pageDetail.SYSDeletedBy_props = rD.SYSDeletedBy_props
	pageDetail.SYSDeletedHost_props = rD.SYSDeletedHost_props
	pageDetail.Name_props = rD.Name_props
	
	// 
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	