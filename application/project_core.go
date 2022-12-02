package application
// ----------------------------------------------------------------
// Automatically generated  "/application/project.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Project (project)
// Endpoint 	        : Project (ProjectID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 01/12/2022 at 09:40:02
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)





//Project_Publish annouces the endpoints available for this object
func Project_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Project_Path, Project_Handler)
	mux.HandleFunc(dm.Project_PathList, Project_HandlerList)
	mux.HandleFunc(dm.Project_PathView, Project_HandlerView)
	mux.HandleFunc(dm.Project_PathEdit, Project_HandlerEdit)
	mux.HandleFunc(dm.Project_PathNew, Project_HandlerNew)
	mux.HandleFunc(dm.Project_PathSave, Project_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Project_Title)
    core.Catalog_Add(dm.Project_Title, dm.Project_Path, "", dm.Project_QueryString, "Application")
}


//Project_HandlerList is the handler for the list page
func Project_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Project
	noItems, returnList, _ := dao.Project_GetList()

	pageDetail := dm.Project_PageList{
		Title:            CardTitle(dm.Project_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Project_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Project_TemplateList, w, r, pageDetail)

}


//Project_HandlerView is the handler used to View a page
func Project_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Project_QueryString)
	_, rD, _ := dao.Project_GetByID(searchID)

	pageDetail := dm.Project_Page{
		Title:       CardTitle(dm.Project_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Project_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = project_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Project_TemplateView, w, r, pageDetail)

}


//Project_HandlerEdit is the handler used generate the Edit page
func Project_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Project_QueryString)
	_, rD, _ := dao.Project_GetByID(searchID)
	
	pageDetail := dm.Project_Page{
		Title:       CardTitle(dm.Project_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Project_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = project_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Project_TemplateEdit, w, r, pageDetail)
}


//Project_HandlerSave is the handler used process the saving of an Project
func Project_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("ProjectID"))

	var item dm.Project
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.Project_SYSId_scrn)
		item.ProjectID = r.FormValue(dm.Project_ProjectID_scrn)
		item.OriginID = r.FormValue(dm.Project_OriginID_scrn)
		item.ProjectStateID = r.FormValue(dm.Project_ProjectStateID_scrn)
		item.ProfileID = r.FormValue(dm.Project_ProfileID_scrn)
		item.Name = r.FormValue(dm.Project_Name_scrn)
		item.Description = r.FormValue(dm.Project_Description_scrn)
		item.StartDate = r.FormValue(dm.Project_StartDate_scrn)
		item.EndDate = r.FormValue(dm.Project_EndDate_scrn)
		item.SYSCreated = r.FormValue(dm.Project_SYSCreated_scrn)
		item.SYSCreatedBy = r.FormValue(dm.Project_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.Project_SYSCreatedHost_scrn)
		item.SYSUpdated = r.FormValue(dm.Project_SYSUpdated_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.Project_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.Project_SYSUpdatedHost_scrn)
		item.SYSDeleted = r.FormValue(dm.Project_SYSDeleted_scrn)
		item.SYSDeletedBy = r.FormValue(dm.Project_SYSDeletedBy_scrn)
		item.SYSDeletedHost = r.FormValue(dm.Project_SYSDeletedHost_scrn)
		item.SYSActivity = r.FormValue(dm.Project_SYSActivity_scrn)
	
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Project_Store(item,r)	
	http.Redirect(w, r, dm.Project_Redirect, http.StatusFound)
}


//Project_HandlerNew is the handler used process the creation of an Project
func Project_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.Project_New()

	pageDetail := dm.Project_Page{
		Title:       CardTitle(dm.Project_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Project_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = project_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Project_TemplateNew, w, r, pageDetail)

}	



// Builds/Popuplates the Project Page 
func project_PopulatePage(rD dm.Project, pageDetail dm.Project_Page) dm.Project_Page {
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.ProjectID = rD.ProjectID
	pageDetail.OriginID = rD.OriginID
	pageDetail.ProjectStateID = rD.ProjectStateID
	pageDetail.ProfileID = rD.ProfileID
	pageDetail.Name = rD.Name
	pageDetail.Description = rD.Description
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
	
	
	//
	// Automatically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	pageDetail.OriginID_lookup = dao.Origin_GetLookup()
	
	
	
	pageDetail.ProjectStateID_lookup = dao.ProjectState_GetLookup()
	
	
	
	pageDetail.ProfileID_lookup = dao.Profile_GetLookup()
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.ProjectID_props = rD.ProjectID_props
	pageDetail.OriginID_props = rD.OriginID_props
	pageDetail.ProjectStateID_props = rD.ProjectStateID_props
	pageDetail.ProfileID_props = rD.ProfileID_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.Description_props = rD.Description_props
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
	
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	