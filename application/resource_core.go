package application
// ----------------------------------------------------------------
// Automatically generated  "/application/resource.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Resource (resource)
// Endpoint 	        : Resource (ResourceID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 03/03/2023 at 17:00:59
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)

//Resource_Publish annouces the endpoints available for this object
//Resource_Publish - Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local
func Resource_Publish(mux http.ServeMux) {
	// START
	// Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	mux.HandleFunc(dm.Resource_Path, Resource_Handler)
	mux.HandleFunc(dm.Resource_PathList, Resource_HandlerList)
	mux.HandleFunc(dm.Resource_PathView, Resource_HandlerView)
	mux.HandleFunc(dm.Resource_PathEdit, Resource_HandlerEdit)
	mux.HandleFunc(dm.Resource_PathNew, Resource_HandlerNew)
	mux.HandleFunc(dm.Resource_PathSave, Resource_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Resource_Title)
    core.Catalog_Add(dm.Resource_Title, dm.Resource_Path, "", dm.Resource_QueryString, "Application")
	// 
	// Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//Resource_HandlerList is the handler for the list page
//Allows Listing of Resource records
//Resource_HandlerList - Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local
func Resource_HandlerList(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local
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

	var returnList []dm.Resource
	noItems, returnList, _ := dao.Resource_GetList()

	pageDetail := dm.Resource_PageList{
		Title:            CardTitle(dm.Resource_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Resource_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Resource_TemplateList, w, r, pageDetail)
	// 
	// Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END

}


//Resource_HandlerView is the handler used to View a page
//Allows Viewing for an existing Resource record
//Resource_HandlerView - Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func Resource_HandlerView(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local
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

	searchID := core.GetURLparam(r, dm.Resource_QueryString)
	_, rD, _ := dao.Resource_GetByID(searchID)

	pageDetail := dm.Resource_Page{
		Title:       CardTitle(dm.Resource_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Resource_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = resource_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Resource_TemplateView, w, r, pageDetail)
	// 
	// Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//Resource_HandlerEdit is the handler used generate the Edit page
//Allows Editing for an existing Resource record and then allows the user to save the changes
//Resource_HandlerEdit - Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func Resource_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local
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

	searchID := core.GetURLparam(r, dm.Resource_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.Resource
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.Resource)
	} else {
		_, rD, _ = dao.Resource_GetByID(searchID)
	}

	
	pageDetail := dm.Resource_Page{
		Title:       CardTitle(dm.Resource_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Resource_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = resource_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Resource_TemplateEdit, w, r, pageDetail)
	// 
	// Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//Resource_HandlerSave is the handler used process the saving of an Resource
//It is called from the Edit and New pages
//Resource_HandlerSave  - Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func Resource_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("ResourceID")
	logs.Servicing(r.URL.Path+itemID)

	item := resource_DataFromRequest(r)
	
	item, errStore := dao.Resource_Store(item,r)
	if errStore == nil {	
		http.Redirect(w, r, dm.Resource_Redirect, http.StatusFound)
	} else {
		logs.Information(dm.Resource_Name, errStore.Error())
		//http.Redirect(w, r, r.Referer(), http.StatusFound)
		ExecuteRedirect(r.Referer(), w, r,dm.Resource_QueryString,itemID,item)
	}
	// 
	// Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//Resource_HandlerNew is the handler used process the creation of an Resource
//It will create a new Resource and then redirect to the Edit page
//Resource_HandlerNew  - Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func Resource_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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

	searchID := core.GetURLparam(r, dm.Resource_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.Resource
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.Resource)
	} else {
		_, _, rD, _ = dao.Resource_New()
	}



	pageDetail := dm.Resource_Page{
		Title:       CardTitle(dm.Resource_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Resource_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = resource_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Resource_TemplateNew, w, r, pageDetail)
	// 
	// Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}	



//resource_PopulatePage Builds/Populates the Resource Page 
//resource_PopulatePage Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func resource_PopulatePage(rD dm.Resource, pageDetail dm.Resource_Page) dm.Resource_Page {
	// Real DB Fields
	pageDetail.SYSId = rD.SYSId
	pageDetail.ResourceID = rD.ResourceID
	pageDetail.Code = rD.Code
	pageDetail.Name = rD.Name
	pageDetail.Email = rD.Email
	pageDetail.Class = rD.Class
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.SYSDeleted = rD.SYSDeleted
	pageDetail.SYSDeletedBy = rD.SYSDeletedBy
	pageDetail.SYSDeletedHost = rD.SYSDeletedHost
	pageDetail.UserActive = rD.UserActive
	pageDetail.SYSActivity = rD.SYSActivity
	pageDetail.Notes = rD.Notes
	pageDetail.SYSDbVersion = rD.SYSDbVersion
	pageDetail.Comments = rD.Comments
	// Add Pseudo/Extra Fields
	// Enrichment Fields 
	
	pageDetail.Class_lookup = dao.StubLists_Get("resourcetype")
	
	pageDetail.UserActive_lookup = dao.StubLists_Get("tf")
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.ResourceID_props = rD.ResourceID_props
	pageDetail.Code_props = rD.Code_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.Email_props = rD.Email_props
	pageDetail.Class_props = rD.Class_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.SYSDeleted_props = rD.SYSDeleted_props
	pageDetail.SYSDeletedBy_props = rD.SYSDeletedBy_props
	pageDetail.SYSDeletedHost_props = rD.SYSDeletedHost_props
	pageDetail.UserActive_props = rD.UserActive_props
	pageDetail.SYSActivity_props = rD.SYSActivity_props
	pageDetail.Notes_props = rD.Notes_props
	pageDetail.SYSDbVersion_props = rD.SYSDbVersion_props
	pageDetail.Comments_props = rD.Comments_props
	return pageDetail
}
//resource_DataFromRequest is used process the content of an HTTP Request and return an instance of an Resource
//resource_DataFromRequest Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func resource_DataFromRequest(r *http.Request) dm.Resource {

	var item dm.Resource
		item.SYSId = r.FormValue(dm.Resource_SYSId_scrn)
		item.ResourceID = r.FormValue(dm.Resource_ResourceID_scrn)
		item.Code = r.FormValue(dm.Resource_Code_scrn)
		item.Name = r.FormValue(dm.Resource_Name_scrn)
		item.Email = r.FormValue(dm.Resource_Email_scrn)
		item.Class = r.FormValue(dm.Resource_Class_scrn)
		item.SYSCreated = r.FormValue(dm.Resource_SYSCreated_scrn)
		item.SYSCreatedBy = r.FormValue(dm.Resource_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.Resource_SYSCreatedHost_scrn)
		item.SYSUpdated = r.FormValue(dm.Resource_SYSUpdated_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.Resource_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.Resource_SYSUpdatedHost_scrn)
		item.SYSDeleted = r.FormValue(dm.Resource_SYSDeleted_scrn)
		item.SYSDeletedBy = r.FormValue(dm.Resource_SYSDeletedBy_scrn)
		item.SYSDeletedHost = r.FormValue(dm.Resource_SYSDeletedHost_scrn)
		item.UserActive = r.FormValue(dm.Resource_UserActive_scrn)
		item.SYSActivity = r.FormValue(dm.Resource_SYSActivity_scrn)
		item.Notes = r.FormValue(dm.Resource_Notes_scrn)
		item.SYSDbVersion = r.FormValue(dm.Resource_SYSDbVersion_scrn)
		item.Comments = r.FormValue(dm.Resource_Comments_scrn)
	return item
}
