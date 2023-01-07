package application
// ----------------------------------------------------------------
// Automatically generated  "/application/origin.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Origin (origin)
// Endpoint 	        : Origin (OriginID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 07/01/2023 at 23:01:29
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)

//Origin_Publish annouces the endpoints available for this object
//Origin_Publish - Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
func Origin_Publish(mux http.ServeMux) {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	mux.HandleFunc(dm.Origin_Path, Origin_Handler)
	mux.HandleFunc(dm.Origin_PathList, Origin_HandlerList)
	mux.HandleFunc(dm.Origin_PathView, Origin_HandlerView)
	mux.HandleFunc(dm.Origin_PathEdit, Origin_HandlerEdit)
	mux.HandleFunc(dm.Origin_PathNew, Origin_HandlerNew)
	mux.HandleFunc(dm.Origin_PathSave, Origin_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Origin_Title)
    core.Catalog_Add(dm.Origin_Title, dm.Origin_Path, "", dm.Origin_QueryString, "Application")
	// 
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//Origin_HandlerList is the handler for the list page
//Allows Listing of Origin records
//Origin_HandlerList - Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
func Origin_HandlerList(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
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

	var returnList []dm.Origin
	noItems, returnList, _ := dao.Origin_GetList()

	pageDetail := dm.Origin_PageList{
		Title:            CardTitle(dm.Origin_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Origin_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Origin_TemplateList, w, r, pageDetail)
	// 
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END

}


//Origin_HandlerView is the handler used to View a page
//Allows Viewing for an existing Origin record
//Origin_HandlerView - Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
func Origin_HandlerView(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
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

	searchID := core.GetURLparam(r, dm.Origin_QueryString)
	_, rD, _ := dao.Origin_GetByID(searchID)

	pageDetail := dm.Origin_Page{
		Title:       CardTitle(dm.Origin_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Origin_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = origin_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Origin_TemplateView, w, r, pageDetail)
	// 
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//Origin_HandlerEdit is the handler used generate the Edit page
//Allows Editing for an existing Origin record and then allows the user to save the changes
//Origin_HandlerEdit - Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
func Origin_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
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

	searchID := core.GetURLparam(r, dm.Origin_QueryString)
	_, rD, _ := dao.Origin_GetByID(searchID)
	
	pageDetail := dm.Origin_Page{
		Title:       CardTitle(dm.Origin_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Origin_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = origin_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Origin_TemplateEdit, w, r, pageDetail)
	// 
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//Origin_HandlerSave is the handler used process the saving of an Origin
//It is called from the Edit and New pages
//Origin_HandlerSave  - Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
func Origin_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("OriginID"))

	item := origin_DataFromRequest(r)
	
	dao.Origin_Store(item,r)	
	http.Redirect(w, r, dm.Origin_Redirect, http.StatusFound)
	// 
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//Origin_HandlerNew is the handler used process the creation of an Origin
//It will create a new Origin and then redirect to the Edit page
//Origin_HandlerNew  - Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
func Origin_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	_, _, rD, _ := dao.Origin_New()

	pageDetail := dm.Origin_Page{
		Title:       CardTitle(dm.Origin_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Origin_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = origin_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Origin_TemplateNew, w, r, pageDetail)
	// 
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}	



//origin_PopulatePage Builds/Populates the Origin Page 
func origin_PopulatePage(rD dm.Origin, pageDetail dm.Origin_Page) dm.Origin_Page {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.OriginID = rD.OriginID
	pageDetail.StateID = rD.StateID
	pageDetail.DocTypeID = rD.DocTypeID
	pageDetail.Code = rD.Code
	pageDetail.FullName = rD.FullName
	pageDetail.Rate = rD.Rate
	pageDetail.Notes = rD.Notes
	pageDetail.StartDate = rD.StartDate
	pageDetail.EndDate = rD.EndDate
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.SYSDeleted = rD.SYSDeleted
	pageDetail.SYSDeletedBy = rD.SYSDeletedBy
	pageDetail.SYSDeletedHost = rD.SYSDeletedHost
	pageDetail.SYSActivity = rD.SYSActivity
	pageDetail.Currency = rD.Currency
	pageDetail.SYSDbVersion = rD.SYSDbVersion
	pageDetail.Comments = rD.Comments
	pageDetail.ProjectManager = rD.ProjectManager
	pageDetail.AccountManager = rD.AccountManager
	
	pageDetail.NoActiveProjects = rD.NoActiveProjects
	
	//
	// Automatically generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	pageDetail.StateID_lookup = dao.OriginState_GetLookup()
	
	
	
	pageDetail.DocTypeID_lookup = dao.DocType_GetLookup()
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.Currency_lookup = dao.StubLists_Get("ccy")
	
	
	
	
	
	
	pageDetail.ProjectManager_lookup = dao.Resource_GetLookup()
	
	
	
	pageDetail.AccountManager_lookup = dao.Resource_GetLookup()
	
	
	
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.OriginID_props = rD.OriginID_props
	pageDetail.StateID_props = rD.StateID_props
	pageDetail.DocTypeID_props = rD.DocTypeID_props
	pageDetail.Code_props = rD.Code_props
	pageDetail.FullName_props = rD.FullName_props
	pageDetail.Rate_props = rD.Rate_props
	pageDetail.Notes_props = rD.Notes_props
	pageDetail.StartDate_props = rD.StartDate_props
	pageDetail.EndDate_props = rD.EndDate_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.SYSDeleted_props = rD.SYSDeleted_props
	pageDetail.SYSDeletedBy_props = rD.SYSDeletedBy_props
	pageDetail.SYSDeletedHost_props = rD.SYSDeletedHost_props
	pageDetail.SYSActivity_props = rD.SYSActivity_props
	pageDetail.Currency_props = rD.Currency_props
	pageDetail.SYSDbVersion_props = rD.SYSDbVersion_props
	pageDetail.Comments_props = rD.Comments_props
	pageDetail.ProjectManager_props = rD.ProjectManager_props
	pageDetail.AccountManager_props = rD.AccountManager_props
	pageDetail.NoActiveProjects_props = rD.NoActiveProjects_props
	
	// 
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	


//origin_DataFromRequest is used process the content of an HTTP Request and return an instance of an Origin
func origin_DataFromRequest(r *http.Request) dm.Origin {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	var item dm.Origin
	// FIELD SET START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.Origin_SYSId_scrn)
		item.OriginID = r.FormValue(dm.Origin_OriginID_scrn)
		item.StateID = r.FormValue(dm.Origin_StateID_scrn)
		item.DocTypeID = r.FormValue(dm.Origin_DocTypeID_scrn)
		item.Code = r.FormValue(dm.Origin_Code_scrn)
		item.FullName = r.FormValue(dm.Origin_FullName_scrn)
		item.Rate = r.FormValue(dm.Origin_Rate_scrn)
		item.Notes = r.FormValue(dm.Origin_Notes_scrn)
		item.StartDate = r.FormValue(dm.Origin_StartDate_scrn)
		item.EndDate = r.FormValue(dm.Origin_EndDate_scrn)
		item.SYSCreated = r.FormValue(dm.Origin_SYSCreated_scrn)
		item.SYSCreatedBy = r.FormValue(dm.Origin_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.Origin_SYSCreatedHost_scrn)
		item.SYSUpdated = r.FormValue(dm.Origin_SYSUpdated_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.Origin_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.Origin_SYSUpdatedHost_scrn)
		item.SYSDeleted = r.FormValue(dm.Origin_SYSDeleted_scrn)
		item.SYSDeletedBy = r.FormValue(dm.Origin_SYSDeletedBy_scrn)
		item.SYSDeletedHost = r.FormValue(dm.Origin_SYSDeletedHost_scrn)
		item.SYSActivity = r.FormValue(dm.Origin_SYSActivity_scrn)
		item.Currency = r.FormValue(dm.Origin_Currency_scrn)
		item.SYSDbVersion = r.FormValue(dm.Origin_SYSDbVersion_scrn)
		item.Comments = r.FormValue(dm.Origin_Comments_scrn)
		item.ProjectManager = r.FormValue(dm.Origin_ProjectManager_scrn)
		item.AccountManager = r.FormValue(dm.Origin_AccountManager_scrn)
		item.NoActiveProjects = r.FormValue(dm.Origin_NoActiveProjects_scrn)
	
	// 
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return item
}
