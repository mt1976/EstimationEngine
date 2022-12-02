package application
// ----------------------------------------------------------------
// Automatically generated  "/application/resource.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Resource (resource)
// Endpoint 	        : Resource (ResourceID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 01/12/2022 at 09:40:03
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
func Resource_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Resource_Path, Resource_Handler)
	mux.HandleFunc(dm.Resource_PathList, Resource_HandlerList)
	mux.HandleFunc(dm.Resource_PathView, Resource_HandlerView)
	mux.HandleFunc(dm.Resource_PathEdit, Resource_HandlerEdit)
	mux.HandleFunc(dm.Resource_PathNew, Resource_HandlerNew)
	mux.HandleFunc(dm.Resource_PathSave, Resource_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Resource_Title)
    core.Catalog_Add(dm.Resource_Title, dm.Resource_Path, "", dm.Resource_QueryString, "Application")
}


//Resource_HandlerList is the handler for the list page
func Resource_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

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

}


//Resource_HandlerView is the handler used to View a page
func Resource_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
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

}


//Resource_HandlerEdit is the handler used generate the Edit page
func Resource_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
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
		Title:       CardTitle(dm.Resource_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Resource_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = resource_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Resource_TemplateEdit, w, r, pageDetail)
}


//Resource_HandlerSave is the handler used process the saving of an Resource
func Resource_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("ResourceID"))

	var item dm.Resource
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
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
	
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Resource_Store(item,r)	
	http.Redirect(w, r, dm.Resource_Redirect, http.StatusFound)
}


//Resource_HandlerNew is the handler used process the creation of an Resource
func Resource_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.Resource_New()

	pageDetail := dm.Resource_Page{
		Title:       CardTitle(dm.Resource_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Resource_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = resource_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Resource_TemplateNew, w, r, pageDetail)

}	



// Builds/Popuplates the Resource Page 
func resource_PopulatePage(rD dm.Resource, pageDetail dm.Resource_Page) dm.Resource_Page {
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
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
	
	
	//
	// Automatically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
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
	
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	