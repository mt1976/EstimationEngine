package application
// ----------------------------------------------------------------
// Automatically generated  "/application/projectstate.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : ProjectState (projectstate)
// Endpoint 	        : ProjectState (ProjectStateID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 13/03/2023 at 14:22:30
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)

//ProjectState_Publish annouces the endpoints available for this object
//ProjectState_Publish - Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
func ProjectState_Publish(mux http.ServeMux) {
	// START
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	mux.HandleFunc(dm.ProjectState_Path, ProjectState_Handler)
	mux.HandleFunc(dm.ProjectState_PathList, ProjectState_HandlerList)
	mux.HandleFunc(dm.ProjectState_PathView, ProjectState_HandlerView)
	mux.HandleFunc(dm.ProjectState_PathEdit, ProjectState_HandlerEdit)
	mux.HandleFunc(dm.ProjectState_PathNew, ProjectState_HandlerNew)
	mux.HandleFunc(dm.ProjectState_PathSave, ProjectState_HandlerSave)
	mux.HandleFunc(dm.ProjectState_PathDelete, ProjectState_HandlerDelete)
	logs.Publish("Application", dm.ProjectState_Title)
    core.API = core.API.AddRoute(dm.ProjectState_Title, dm.ProjectState_Path, "", dm.ProjectState_QueryString, "Application")
	// 
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//ProjectState_HandlerList is the handler for the list page
//Allows Listing of ProjectState records
//ProjectState_HandlerList - Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
func ProjectState_HandlerList(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
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

	var returnList []dm.ProjectState

	objectName := dao.Translate("ObjectName", "ProjectState")
	reqField := "Base"
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule)
	if filter == "" {
		logs.Warning("No filter found : " + reqField + " for Object: " + objectName)
	} 

	noItems, returnList, _ := dao.ProjectState_GetListFiltered(filter)

	pageDetail := dm.ProjectState_PageList{
		Title:            CardTitle(dm.ProjectState_Title, core.Action_List),
		PageTitle:        PageTitle(dm.ProjectState_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	nextTemplate :=  NextTemplate("ProjectState", "List", dm.ProjectState_TemplateList)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END

}


//ProjectState_HandlerView is the handler used to View a page
//Allows Viewing for an existing ProjectState record
//ProjectState_HandlerView - Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func ProjectState_HandlerView(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
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

	searchID := core.GetURLparam(r, dm.ProjectState_QueryString)
	_, rD, _ := dao.ProjectState_GetByID(searchID)

	pageDetail := dm.ProjectState_Page{
		Title:       CardTitle(dm.ProjectState_Title, core.Action_View),
		PageTitle:   PageTitle(dm.ProjectState_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = projectstate_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("ProjectState", "View", dm.ProjectState_TemplateView)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//ProjectState_HandlerEdit is the handler used generate the Edit page
//Allows Editing for an existing ProjectState record and then allows the user to save the changes
//ProjectState_HandlerEdit - Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func ProjectState_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
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

	searchID := core.GetURLparam(r, dm.ProjectState_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.ProjectState
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.ProjectState)
	} else {
		_, rD, _ = dao.ProjectState_GetByID(searchID)
	}

	
	pageDetail := dm.ProjectState_Page{
		Title:       CardTitle(dm.ProjectState_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.ProjectState_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = projectstate_PopulatePage(rD , pageDetail) 


	nextTemplate :=  NextTemplate("ProjectState", "Edit", dm.ProjectState_TemplateEdit)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//ProjectState_HandlerSave is the handler used process the saving of an ProjectState
//It is called from the Edit and New pages
//ProjectState_HandlerSave  - Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func ProjectState_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("ProjectStateID")
	logs.Servicing(r.URL.Path+itemID)

	item := projectstate_DataFromRequest(r)
	
	item, errStore := dao.ProjectState_Store(item,r)
	if errStore == nil {
		nextTemplate :=  NextTemplate("ProjectState", "Save", dm.ProjectState_Redirect)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.ProjectState_Name, errStore.Error())
		//http.Redirect(w, r, r.Referer(), http.StatusFound)
		ExecuteRedirect(r.Referer(), w, r,dm.ProjectState_QueryString,itemID,item)
	}
	// 
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//ProjectState_HandlerNew is the handler used process the creation of an ProjectState
//It will create a new ProjectState and then redirect to the Edit page
//ProjectState_HandlerNew  - Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func ProjectState_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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

	searchID := core.GetURLparam(r, dm.ProjectState_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.ProjectState
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.ProjectState)
	} else {
		_, _, rD, _ = dao.ProjectState_New()
	}



	pageDetail := dm.ProjectState_Page{
		Title:       CardTitle(dm.ProjectState_Title, core.Action_New),
		PageTitle:   PageTitle(dm.ProjectState_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = projectstate_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("ProjectState", "New", dm.ProjectState_TemplateNew)
	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}	


//ProjectState_HandlerDelete is the handler used process the deletion of an ProjectState
// It will delete the ProjectState and then redirect to the List page
//ProjectState_HandlerDelete - Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func ProjectState_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	//
	// Code Continues Below
	//
	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.ProjectState_QueryString)

	dao.ProjectState_Delete(searchID)	

	nextTemplate :=  NextTemplate("ProjectState", "Delete", dm.ProjectState_Redirect)
	http.Redirect(w, r, nextTemplate, http.StatusFound)
	// 
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//projectstate_PopulatePage Builds/Populates the ProjectState Page 
//projectstate_PopulatePage Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func projectstate_PopulatePage(rD dm.ProjectState, pageDetail dm.ProjectState_Page) dm.ProjectState_Page {
	// Real DB Fields
	pageDetail.SYSId = rD.SYSId
	pageDetail.ProjectStateID = rD.ProjectStateID
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
	pageDetail.IsLocked = rD.IsLocked
	pageDetail.Notify = rD.Notify
	pageDetail.Migrate = rD.Migrate
	// Add Pseudo/Extra Fields
	// Enrichment Fields 
	
	pageDetail.IsLocked_lookup = dao.StubLists_Get("tf")
	
	pageDetail.Notify_lookup = dao.StubLists_Get("tf")
	
	pageDetail.Migrate_lookup = dao.StubLists_Get("tf")
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.ProjectStateID_props = rD.ProjectStateID_props
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
	pageDetail.IsLocked_props = rD.IsLocked_props
	pageDetail.Notify_props = rD.Notify_props
	pageDetail.Migrate_props = rD.Migrate_props
	return pageDetail
}
//projectstate_DataFromRequest is used process the content of an HTTP Request and return an instance of an ProjectState
//projectstate_DataFromRequest Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func projectstate_DataFromRequest(r *http.Request) dm.ProjectState {

	var item dm.ProjectState
		item.SYSId = r.FormValue(dm.ProjectState_SYSId_scrn)
		item.ProjectStateID = r.FormValue(dm.ProjectState_ProjectStateID_scrn)
		item.Code = r.FormValue(dm.ProjectState_Code_scrn)
		item.Name = r.FormValue(dm.ProjectState_Name_scrn)
		item.SYSCreated = r.FormValue(dm.ProjectState_SYSCreated_scrn)
		item.SYSCreatedBy = r.FormValue(dm.ProjectState_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.ProjectState_SYSCreatedHost_scrn)
		item.SYSUpdated = r.FormValue(dm.ProjectState_SYSUpdated_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.ProjectState_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.ProjectState_SYSUpdatedHost_scrn)
		item.SYSDeleted = r.FormValue(dm.ProjectState_SYSDeleted_scrn)
		item.SYSDeletedBy = r.FormValue(dm.ProjectState_SYSDeletedBy_scrn)
		item.SYSDeletedHost = r.FormValue(dm.ProjectState_SYSDeletedHost_scrn)
		item.SYSDbVersion = r.FormValue(dm.ProjectState_SYSDbVersion_scrn)
		item.IsLocked = r.FormValue(dm.ProjectState_IsLocked_scrn)
		item.Notify = r.FormValue(dm.ProjectState_Notify_scrn)
		item.Migrate = r.FormValue(dm.ProjectState_Migrate_scrn)
	return item
}
