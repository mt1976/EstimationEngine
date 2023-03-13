package application
// ----------------------------------------------------------------
// Automatically generated  "/application/userrole.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : UserRole (userrole)
// Endpoint 	        : UserRole (Id)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 13/03/2023 at 14:22:32
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)

//UserRole_Publish annouces the endpoints available for this object
//UserRole_Publish - Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
func UserRole_Publish(mux http.ServeMux) {
	// START
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	mux.HandleFunc(dm.UserRole_Path, UserRole_Handler)
	mux.HandleFunc(dm.UserRole_PathList, UserRole_HandlerList)
	mux.HandleFunc(dm.UserRole_PathView, UserRole_HandlerView)
	mux.HandleFunc(dm.UserRole_PathEdit, UserRole_HandlerEdit)
	mux.HandleFunc(dm.UserRole_PathNew, UserRole_HandlerNew)
	mux.HandleFunc(dm.UserRole_PathSave, UserRole_HandlerSave)
	mux.HandleFunc(dm.UserRole_PathDelete, UserRole_HandlerDelete)
	logs.Publish("Application", dm.UserRole_Title)
    core.API = core.API.AddRoute(dm.UserRole_Title, dm.UserRole_Path, "", dm.UserRole_QueryString, "Application")
	// 
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//UserRole_HandlerList is the handler for the list page
//Allows Listing of UserRole records
//UserRole_HandlerList - Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
func UserRole_HandlerList(w http.ResponseWriter, r *http.Request) {
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

	var returnList []dm.UserRole

	objectName := dao.Translate("ObjectName", "UserRole")
	reqField := "Base"
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule)
	if filter == "" {
		logs.Warning("No filter found : " + reqField + " for Object: " + objectName)
	} 

	noItems, returnList, _ := dao.UserRole_GetListFiltered(filter)

	pageDetail := dm.UserRole_PageList{
		Title:            CardTitle(dm.UserRole_Title, core.Action_List),
		PageTitle:        PageTitle(dm.UserRole_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	nextTemplate :=  NextTemplate("UserRole", "List", dm.UserRole_TemplateList)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END

}


//UserRole_HandlerView is the handler used to View a page
//Allows Viewing for an existing UserRole record
//UserRole_HandlerView - Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func UserRole_HandlerView(w http.ResponseWriter, r *http.Request) {
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

	searchID := core.GetURLparam(r, dm.UserRole_QueryString)
	_, rD, _ := dao.UserRole_GetByID(searchID)

	pageDetail := dm.UserRole_Page{
		Title:       CardTitle(dm.UserRole_Title, core.Action_View),
		PageTitle:   PageTitle(dm.UserRole_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = userrole_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("UserRole", "View", dm.UserRole_TemplateView)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//UserRole_HandlerEdit is the handler used generate the Edit page
//Allows Editing for an existing UserRole record and then allows the user to save the changes
//UserRole_HandlerEdit - Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func UserRole_HandlerEdit(w http.ResponseWriter, r *http.Request) {
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

	searchID := core.GetURLparam(r, dm.UserRole_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.UserRole
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.UserRole)
	} else {
		_, rD, _ = dao.UserRole_GetByID(searchID)
	}

	
	pageDetail := dm.UserRole_Page{
		Title:       CardTitle(dm.UserRole_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.UserRole_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = userrole_PopulatePage(rD , pageDetail) 


	nextTemplate :=  NextTemplate("UserRole", "Edit", dm.UserRole_TemplateEdit)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//UserRole_HandlerSave is the handler used process the saving of an UserRole
//It is called from the Edit and New pages
//UserRole_HandlerSave  - Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func UserRole_HandlerSave(w http.ResponseWriter, r *http.Request) {
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
	itemID := r.FormValue("Id")
	logs.Servicing(r.URL.Path+itemID)

	item := userrole_DataFromRequest(r)
	
	item, errStore := dao.UserRole_Store(item,r)
	if errStore == nil {
		nextTemplate :=  NextTemplate("UserRole", "Save", dm.UserRole_Redirect)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.UserRole_Name, errStore.Error())
		//http.Redirect(w, r, r.Referer(), http.StatusFound)
		ExecuteRedirect(r.Referer(), w, r,dm.UserRole_QueryString,itemID,item)
	}
	// 
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//UserRole_HandlerNew is the handler used process the creation of an UserRole
//It will create a new UserRole and then redirect to the Edit page
//UserRole_HandlerNew  - Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func UserRole_HandlerNew(w http.ResponseWriter, r *http.Request) {
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

	searchID := core.GetURLparam(r, dm.UserRole_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.UserRole
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.UserRole)
	} else {
		_, _, rD, _ = dao.UserRole_New()
	}



	pageDetail := dm.UserRole_Page{
		Title:       CardTitle(dm.UserRole_Title, core.Action_New),
		PageTitle:   PageTitle(dm.UserRole_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = userrole_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("UserRole", "New", dm.UserRole_TemplateNew)
	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}	


//UserRole_HandlerDelete is the handler used process the deletion of an UserRole
// It will delete the UserRole and then redirect to the List page
//UserRole_HandlerDelete - Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func UserRole_HandlerDelete(w http.ResponseWriter, r *http.Request) {
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
	searchID := core.GetURLparam(r, dm.UserRole_QueryString)

	dao.UserRole_Delete(searchID)	

	nextTemplate :=  NextTemplate("UserRole", "Delete", dm.UserRole_Redirect)
	http.Redirect(w, r, nextTemplate, http.StatusFound)
	// 
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//userrole_PopulatePage Builds/Populates the UserRole Page 
//userrole_PopulatePage Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func userrole_PopulatePage(rD dm.UserRole, pageDetail dm.UserRole_Page) dm.UserRole_Page {
	// Real DB Fields
	pageDetail.SYSId = rD.SYSId
	pageDetail.Id = rD.Id
	pageDetail.Name = rD.Name
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSDeleted = rD.SYSDeleted
	pageDetail.SYSDeletedBy = rD.SYSDeletedBy
	pageDetail.SYSDeletedHost = rD.SYSDeletedHost
	pageDetail.SYSDbVersion = rD.SYSDbVersion
	// Add Pseudo/Extra Fields
	// Enrichment Fields 
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.Id_props = rD.Id_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSDeleted_props = rD.SYSDeleted_props
	pageDetail.SYSDeletedBy_props = rD.SYSDeletedBy_props
	pageDetail.SYSDeletedHost_props = rD.SYSDeletedHost_props
	pageDetail.SYSDbVersion_props = rD.SYSDbVersion_props
	return pageDetail
}
//userrole_DataFromRequest is used process the content of an HTTP Request and return an instance of an UserRole
//userrole_DataFromRequest Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func userrole_DataFromRequest(r *http.Request) dm.UserRole {

	var item dm.UserRole
		item.SYSId = r.FormValue(dm.UserRole_SYSId_scrn)
		item.Id = r.FormValue(dm.UserRole_Id_scrn)
		item.Name = r.FormValue(dm.UserRole_Name_scrn)
		item.SYSCreatedBy = r.FormValue(dm.UserRole_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.UserRole_SYSCreatedHost_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.UserRole_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.UserRole_SYSUpdatedHost_scrn)
		item.SYSUpdated = r.FormValue(dm.UserRole_SYSUpdated_scrn)
		item.SYSCreated = r.FormValue(dm.UserRole_SYSCreated_scrn)
		item.SYSDeleted = r.FormValue(dm.UserRole_SYSDeleted_scrn)
		item.SYSDeletedBy = r.FormValue(dm.UserRole_SYSDeletedBy_scrn)
		item.SYSDeletedHost = r.FormValue(dm.UserRole_SYSDeletedHost_scrn)
		item.SYSDbVersion = r.FormValue(dm.UserRole_SYSDbVersion_scrn)
	return item
}
