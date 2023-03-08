package application
// ----------------------------------------------------------------
// Automatically generated  "/application/credentialsaction.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CredentialsAction (credentialsaction)
// Endpoint 	        : CredentialsAction (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 04/03/2023 at 20:14:10
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)

//CredentialsAction_Publish annouces the endpoints available for this object
//CredentialsAction_Publish - Auto generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
func CredentialsAction_Publish(mux http.ServeMux) {
	// START
	// Auto generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	//No API
	//Cannot List via GUI
	mux.HandleFunc(dm.CredentialsAction_PathView, CredentialsAction_HandlerView)
	mux.HandleFunc(dm.CredentialsAction_PathEdit, CredentialsAction_HandlerEdit)
	mux.HandleFunc(dm.CredentialsAction_PathNew, CredentialsAction_HandlerNew)
	mux.HandleFunc(dm.CredentialsAction_PathSave, CredentialsAction_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.CredentialsAction_Title)
    //No API
	// 
	// Auto generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}



//CredentialsAction_HandlerView is the handler used to View a page
//Allows Viewing for an existing CredentialsAction record
//CredentialsAction_HandlerView - Auto generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func CredentialsAction_HandlerView(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
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

	searchID := core.GetURLparam(r, dm.CredentialsAction_QueryString)
	_, rD, _ := dao.CredentialsAction_GetByID(searchID)

	pageDetail := dm.CredentialsAction_Page{
		Title:       CardTitle(dm.CredentialsAction_Title, core.Action_View),
		PageTitle:   PageTitle(dm.CredentialsAction_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = credentialsaction_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.CredentialsAction_TemplateView, w, r, pageDetail)
	// 
	// Auto generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//CredentialsAction_HandlerEdit is the handler used generate the Edit page
//Allows Editing for an existing CredentialsAction record and then allows the user to save the changes
//CredentialsAction_HandlerEdit - Auto generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func CredentialsAction_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
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

	searchID := core.GetURLparam(r, dm.CredentialsAction_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.CredentialsAction
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.CredentialsAction)
	} else {
		_, rD, _ = dao.CredentialsAction_GetByID(searchID)
	}

	
	pageDetail := dm.CredentialsAction_Page{
		Title:       CardTitle(dm.CredentialsAction_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.CredentialsAction_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = credentialsaction_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.CredentialsAction_TemplateEdit, w, r, pageDetail)
	// 
	// Auto generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//CredentialsAction_HandlerSave is the handler used process the saving of an CredentialsAction
//It is called from the Edit and New pages
//CredentialsAction_HandlerSave  - Auto generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func CredentialsAction_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("ID")
	logs.Servicing(r.URL.Path+itemID)

	item := credentialsaction_DataFromRequest(r)
	
	item, errStore := dao.CredentialsAction_Store(item,r)
	if errStore == nil {	
		http.Redirect(w, r, dm.CredentialsAction_Redirect, http.StatusFound)
	} else {
		logs.Information(dm.CredentialsAction_Name, errStore.Error())
		//http.Redirect(w, r, r.Referer(), http.StatusFound)
		ExecuteRedirect(r.Referer(), w, r,dm.CredentialsAction_QueryString,itemID,item)
	}
	// 
	// Auto generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//CredentialsAction_HandlerNew is the handler used process the creation of an CredentialsAction
//It will create a new CredentialsAction and then redirect to the Edit page
//CredentialsAction_HandlerNew  - Auto generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func CredentialsAction_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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

	searchID := core.GetURLparam(r, dm.CredentialsAction_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.CredentialsAction
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.CredentialsAction)
	} else {
		_, _, rD, _ = dao.CredentialsAction_New()
	}



	pageDetail := dm.CredentialsAction_Page{
		Title:       CardTitle(dm.CredentialsAction_Title, core.Action_New),
		PageTitle:   PageTitle(dm.CredentialsAction_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = credentialsaction_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.CredentialsAction_TemplateNew, w, r, pageDetail)
	// 
	// Auto generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}	



//credentialsaction_PopulatePage Builds/Populates the CredentialsAction Page 
//credentialsaction_PopulatePage Auto generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func credentialsaction_PopulatePage(rD dm.CredentialsAction, pageDetail dm.CredentialsAction_Page) dm.CredentialsAction_Page {
	// Real DB Fields
	pageDetail.ID = rD.ID
	pageDetail.User = rD.User
	pageDetail.Action = rD.Action
	pageDetail.Notes = rD.Notes
	// Add Pseudo/Extra Fields
	// Enrichment Fields 
	 
	pageDetail.User_lookup = dao.Credentials_GetLookup()
	
	pageDetail.Action_lookup = dao.StubLists_Get("credentialStates")
	pageDetail.ID_props = rD.ID_props
	pageDetail.User_props = rD.User_props
	pageDetail.Action_props = rD.Action_props
	pageDetail.Notes_props = rD.Notes_props
	return pageDetail
}
//credentialsaction_DataFromRequest is used process the content of an HTTP Request and return an instance of an CredentialsAction
//credentialsaction_DataFromRequest Auto generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func credentialsaction_DataFromRequest(r *http.Request) dm.CredentialsAction {

	var item dm.CredentialsAction
		item.ID = r.FormValue(dm.CredentialsAction_ID_scrn)
		item.User = r.FormValue(dm.CredentialsAction_User_scrn)
		item.Action = r.FormValue(dm.CredentialsAction_Action_scrn)
		item.Notes = r.FormValue(dm.CredentialsAction_Notes_scrn)
	return item
}
