package application
// ----------------------------------------------------------------
// Automatically generated  "/application/credentialspassword.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CredentialsPassword (credentialspassword)
// Endpoint 	        : CredentialsPassword (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 13/03/2023 at 14:22:26
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)

//CredentialsPassword_Publish annouces the endpoints available for this object
//CredentialsPassword_Publish - Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
func CredentialsPassword_Publish(mux http.ServeMux) {
	// START
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	//No API
	//Cannot List via GUI
	mux.HandleFunc(dm.CredentialsPassword_PathView, CredentialsPassword_HandlerView)
	mux.HandleFunc(dm.CredentialsPassword_PathEdit, CredentialsPassword_HandlerEdit)
	mux.HandleFunc(dm.CredentialsPassword_PathNew, CredentialsPassword_HandlerNew)
	mux.HandleFunc(dm.CredentialsPassword_PathSave, CredentialsPassword_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.CredentialsPassword_Title)
    //No API
	// 
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}



//CredentialsPassword_HandlerView is the handler used to View a page
//Allows Viewing for an existing CredentialsPassword record
//CredentialsPassword_HandlerView - Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func CredentialsPassword_HandlerView(w http.ResponseWriter, r *http.Request) {
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

	searchID := core.GetURLparam(r, dm.CredentialsPassword_QueryString)
	_, rD, _ := dao.CredentialsPassword_GetByID(searchID)

	pageDetail := dm.CredentialsPassword_Page{
		Title:       CardTitle(dm.CredentialsPassword_Title, core.Action_View),
		PageTitle:   PageTitle(dm.CredentialsPassword_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = credentialspassword_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("CredentialsPassword", "View", dm.CredentialsPassword_TemplateView)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//CredentialsPassword_HandlerEdit is the handler used generate the Edit page
//Allows Editing for an existing CredentialsPassword record and then allows the user to save the changes
//CredentialsPassword_HandlerEdit - Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func CredentialsPassword_HandlerEdit(w http.ResponseWriter, r *http.Request) {
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

	searchID := core.GetURLparam(r, dm.CredentialsPassword_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.CredentialsPassword
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.CredentialsPassword)
	} else {
		_, rD, _ = dao.CredentialsPassword_GetByID(searchID)
	}

	
	pageDetail := dm.CredentialsPassword_Page{
		Title:       CardTitle(dm.CredentialsPassword_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.CredentialsPassword_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = credentialspassword_PopulatePage(rD , pageDetail) 


	nextTemplate :=  NextTemplate("CredentialsPassword", "Edit", dm.CredentialsPassword_TemplateEdit)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//CredentialsPassword_HandlerSave is the handler used process the saving of an CredentialsPassword
//It is called from the Edit and New pages
//CredentialsPassword_HandlerSave  - Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func CredentialsPassword_HandlerSave(w http.ResponseWriter, r *http.Request) {
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
	itemID := r.FormValue("ID")
	logs.Servicing(r.URL.Path+itemID)

	item := credentialspassword_DataFromRequest(r)
	
	item, errStore := dao.CredentialsPassword_Store(item,r)
	if errStore == nil {
		nextTemplate :=  NextTemplate("CredentialsPassword", "Save", dm.CredentialsPassword_Redirect)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.CredentialsPassword_Name, errStore.Error())
		//http.Redirect(w, r, r.Referer(), http.StatusFound)
		ExecuteRedirect(r.Referer(), w, r,dm.CredentialsPassword_QueryString,itemID,item)
	}
	// 
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//CredentialsPassword_HandlerNew is the handler used process the creation of an CredentialsPassword
//It will create a new CredentialsPassword and then redirect to the Edit page
//CredentialsPassword_HandlerNew  - Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func CredentialsPassword_HandlerNew(w http.ResponseWriter, r *http.Request) {
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

	searchID := core.GetURLparam(r, dm.CredentialsPassword_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.CredentialsPassword
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.CredentialsPassword)
	} else {
		_, _, rD, _ = dao.CredentialsPassword_New()
	}



	pageDetail := dm.CredentialsPassword_Page{
		Title:       CardTitle(dm.CredentialsPassword_Title, core.Action_New),
		PageTitle:   PageTitle(dm.CredentialsPassword_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = credentialspassword_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("CredentialsPassword", "New", dm.CredentialsPassword_TemplateNew)
	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}	



//credentialspassword_PopulatePage Builds/Populates the CredentialsPassword Page 
//credentialspassword_PopulatePage Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func credentialspassword_PopulatePage(rD dm.CredentialsPassword, pageDetail dm.CredentialsPassword_Page) dm.CredentialsPassword_Page {
	// Real DB Fields
	pageDetail.ID = rD.ID
	pageDetail.UserName = rD.UserName
	pageDetail.PasswordOld = rD.PasswordOld
	pageDetail.PasswordNew = rD.PasswordNew
	pageDetail.PasswordConfirm = rD.PasswordConfirm
	// Add Pseudo/Extra Fields
	// Enrichment Fields 
	pageDetail.ID_props = rD.ID_props
	pageDetail.UserName_props = rD.UserName_props
	pageDetail.PasswordOld_props = rD.PasswordOld_props
	pageDetail.PasswordNew_props = rD.PasswordNew_props
	pageDetail.PasswordConfirm_props = rD.PasswordConfirm_props
	return pageDetail
}
//credentialspassword_DataFromRequest is used process the content of an HTTP Request and return an instance of an CredentialsPassword
//credentialspassword_DataFromRequest Auto generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func credentialspassword_DataFromRequest(r *http.Request) dm.CredentialsPassword {

	var item dm.CredentialsPassword
		item.ID = r.FormValue(dm.CredentialsPassword_ID_scrn)
		item.UserName = r.FormValue(dm.CredentialsPassword_UserName_scrn)
		item.PasswordOld = r.FormValue(dm.CredentialsPassword_PasswordOld_scrn)
		item.PasswordNew = r.FormValue(dm.CredentialsPassword_PasswordNew_scrn)
		item.PasswordConfirm = r.FormValue(dm.CredentialsPassword_PasswordConfirm_scrn)
	return item
}
