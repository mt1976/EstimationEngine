package application
// ----------------------------------------------------------------
// Automatically generated  "/application/estimationstate.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : EstimationState (estimationstate)
// Endpoint 	        : EstimationState (EstimationStateID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 10/03/2023 at 19:54:33
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
//EstimationState_Publish - Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
func EstimationState_Publish(mux http.ServeMux) {
	// START
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	mux.HandleFunc(dm.EstimationState_Path, EstimationState_Handler)
	mux.HandleFunc(dm.EstimationState_PathList, EstimationState_HandlerList)
	mux.HandleFunc(dm.EstimationState_PathView, EstimationState_HandlerView)
	mux.HandleFunc(dm.EstimationState_PathEdit, EstimationState_HandlerEdit)
	mux.HandleFunc(dm.EstimationState_PathNew, EstimationState_HandlerNew)
	mux.HandleFunc(dm.EstimationState_PathSave, EstimationState_HandlerSave)
	mux.HandleFunc(dm.EstimationState_PathDelete, EstimationState_HandlerDelete)
	logs.Publish("Application", dm.EstimationState_Title)
    core.API = core.API.AddRoute(dm.EstimationState_Title, dm.EstimationState_Path, "", dm.EstimationState_QueryString, "Application")
	// 
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//EstimationState_HandlerList is the handler for the list page
//Allows Listing of EstimationState records
//EstimationState_HandlerList - Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
func EstimationState_HandlerList(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
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

	objectName := dao.Translate("ObjectName", "EstimationState")
	reqField := "Base"
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule)
	if filter == "" {
		logs.Warning("No filter found : " + reqField + " for Object: " + objectName)
	} 

	noItems, returnList, _ := dao.EstimationState_GetListFiltered(filter)

	pageDetail := dm.EstimationState_PageList{
		Title:            CardTitle(dm.EstimationState_Title, core.Action_List),
		PageTitle:        PageTitle(dm.EstimationState_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	nextTemplate :=  NextTemplate("EstimationState", "List", dm.EstimationState_TemplateList)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END

}


//EstimationState_HandlerView is the handler used to View a page
//Allows Viewing for an existing EstimationState record
//EstimationState_HandlerView - Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func EstimationState_HandlerView(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
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

	nextTemplate :=  NextTemplate("EstimationState", "View", dm.EstimationState_TemplateView)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//EstimationState_HandlerEdit is the handler used generate the Edit page
//Allows Editing for an existing EstimationState record and then allows the user to save the changes
//EstimationState_HandlerEdit - Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func EstimationState_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
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
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.EstimationState
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.EstimationState)
	} else {
		_, rD, _ = dao.EstimationState_GetByID(searchID)
	}

	
	pageDetail := dm.EstimationState_Page{
		Title:       CardTitle(dm.EstimationState_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.EstimationState_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = estimationstate_PopulatePage(rD , pageDetail) 


	nextTemplate :=  NextTemplate("EstimationState", "Edit", dm.EstimationState_TemplateEdit)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//EstimationState_HandlerSave is the handler used process the saving of an EstimationState
//It is called from the Edit and New pages
//EstimationState_HandlerSave  - Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func EstimationState_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("EstimationStateID")
	logs.Servicing(r.URL.Path+itemID)

	item := estimationstate_DataFromRequest(r)
	
	item, errStore := dao.EstimationState_Store(item,r)
	if errStore == nil {
		nextTemplate :=  NextTemplate("EstimationState", "Save", dm.EstimationState_Redirect)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.EstimationState_Name, errStore.Error())
		//http.Redirect(w, r, r.Referer(), http.StatusFound)
		ExecuteRedirect(r.Referer(), w, r,dm.EstimationState_QueryString,itemID,item)
	}
	// 
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//EstimationState_HandlerNew is the handler used process the creation of an EstimationState
//It will create a new EstimationState and then redirect to the Edit page
//EstimationState_HandlerNew  - Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func EstimationState_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.EstimationState
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.EstimationState)
	} else {
		_, _, rD, _ = dao.EstimationState_New()
	}



	pageDetail := dm.EstimationState_Page{
		Title:       CardTitle(dm.EstimationState_Title, core.Action_New),
		PageTitle:   PageTitle(dm.EstimationState_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = estimationstate_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("EstimationState", "New", dm.EstimationState_TemplateNew)
	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}	


//EstimationState_HandlerDelete is the handler used process the deletion of an EstimationState
// It will delete the EstimationState and then redirect to the List page
//EstimationState_HandlerDelete - Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func EstimationState_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	searchID := core.GetURLparam(r, dm.EstimationState_QueryString)

	dao.EstimationState_Delete(searchID)	

	nextTemplate :=  NextTemplate("EstimationState", "Delete", dm.EstimationState_Redirect)
	http.Redirect(w, r, nextTemplate, http.StatusFound)
	// 
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//estimationstate_PopulatePage Builds/Populates the EstimationState Page 
//estimationstate_PopulatePage Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func estimationstate_PopulatePage(rD dm.EstimationState, pageDetail dm.EstimationState_Page) dm.EstimationState_Page {
	// Real DB Fields
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
	pageDetail.IsLocked = rD.IsLocked
	pageDetail.Notify = rD.Notify
	// Add Pseudo/Extra Fields
	// Enrichment Fields 
	
	pageDetail.IsLocked_lookup = dao.StubLists_Get("tf")
	
	pageDetail.Notify_lookup = dao.StubLists_Get("tf")
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
	pageDetail.IsLocked_props = rD.IsLocked_props
	pageDetail.Notify_props = rD.Notify_props
	return pageDetail
}
//estimationstate_DataFromRequest is used process the content of an HTTP Request and return an instance of an EstimationState
//estimationstate_DataFromRequest Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func estimationstate_DataFromRequest(r *http.Request) dm.EstimationState {

	var item dm.EstimationState
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
		item.IsLocked = r.FormValue(dm.EstimationState_IsLocked_scrn)
		item.Notify = r.FormValue(dm.EstimationState_Notify_scrn)
	return item
}
