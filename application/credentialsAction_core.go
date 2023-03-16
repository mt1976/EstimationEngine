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
// Date & Time		    : 15/03/2023 at 19:24:47
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"
	"strings"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)

//CredentialsAction_Publish annouces the endpoints available for this object
func CredentialsAction_Publish(mux http.ServeMux) {
	//No API
	//Cannot List via GUI
	mux.HandleFunc(dm.CredentialsAction_PathView, CredentialsAction_HandlerView)
	mux.HandleFunc(dm.CredentialsAction_PathEdit, CredentialsAction_HandlerEdit)
	mux.HandleFunc(dm.CredentialsAction_PathNew, CredentialsAction_HandlerNew)
	mux.HandleFunc(dm.CredentialsAction_PathSave, CredentialsAction_HandlerSave)
	//Cannot Delete via GUI
    //No API
	logs.Publish("Application", dm.CredentialsAction_Title)
}

//CredentialsAction_HandlerView is the handler used to View a CredentialsAction database record
func CredentialsAction_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
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

	nextTemplate :=  NextTemplate("CredentialsAction", "View", dm.CredentialsAction_TemplateView)
	nextTemplate = credentialsaction_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//CredentialsAction_HandlerEdit is the handler used to Edit of an existing CredentialsAction database record
func CredentialsAction_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
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

	nextTemplate :=  NextTemplate("CredentialsAction", "Edit", dm.CredentialsAction_TemplateEdit)
	nextTemplate = credentialsaction_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//CredentialsAction_HandlerSave is the handler used process the saving of an CredentialsAction database record, either new or existing, referenced by Edit & New Handlers.
func CredentialsAction_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
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
		nextTemplate :=  NextTemplate("CredentialsAction", "Save", dm.CredentialsAction_Redirect)
		nextTemplate = credentialsaction_URIQueryData(nextTemplate,item,itemID)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.CredentialsAction_Name, errStore.Error())
		ExecuteRedirect(r.Referer(), w, r,dm.CredentialsAction_QueryString,itemID,item)
	}
}

//CredentialsAction_HandlerNew is the handler used process the creation of an new CredentialsAction database record, then redirect to Edit
func CredentialsAction_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
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

	nextTemplate :=  NextTemplate("CredentialsAction", "New", dm.CredentialsAction_TemplateNew)
	nextTemplate = credentialsaction_URIQueryData(nextTemplate,dm.CredentialsAction{},searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}	
//credentialsaction_PopulatePage Builds/Populates the CredentialsAction Page from an instance of CredentialsAction from the Data Model
func credentialsaction_PopulatePage(rD dm.CredentialsAction, pageDetail dm.CredentialsAction_Page) dm.CredentialsAction_Page {
	// Real DB Fields
	pageDetail.ID = rD.ID
	pageDetail.User = rD.User
	pageDetail.Action = rD.Action
	pageDetail.Notes = rD.Notes
	// Add Pseudo/Extra Fields, fields that are not in the DB but are used in the UI
	// Enrichment content, content used provide lookups,lists etc
	pageDetail.User_lookup = dao.Credentials_GetLookup()
	pageDetail.Action_lookup = dao.StubLists_Get("credentialStates")
	// Add the Properties for the Fields
	pageDetail.ID_props = rD.ID_props
	pageDetail.User_props = rD.User_props
	pageDetail.Action_props = rD.Action_props
	pageDetail.Notes_props = rD.Notes_props
	return pageDetail
}
//credentialsaction_DataFromRequest is used process the content of an HTTP Request and return an instance of an CredentialsAction
func credentialsaction_DataFromRequest(r *http.Request) dm.CredentialsAction {
	var item dm.CredentialsAction
	item.ID = r.FormValue(dm.CredentialsAction_ID_scrn)
	item.User = r.FormValue(dm.CredentialsAction_User_scrn)
	item.Action = r.FormValue(dm.CredentialsAction_Action_scrn)
	item.Notes = r.FormValue(dm.CredentialsAction_Notes_scrn)
	return item
}
//credentialsaction_URIQueryData is used to replace the wildcards in the URI Query Path with the values from the CredentialsAction Data Model
func credentialsaction_URIQueryData(queryPath string,item dm.CredentialsAction,itemID string) string {
	if queryPath == "" {
		return ""
	}
	queryPath = core.ReplaceWildcard(queryPath, strings.ToUpper("ID"), itemID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.CredentialsAction_ID_scrn), item.ID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.CredentialsAction_User_scrn), item.User)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.CredentialsAction_Action_scrn), item.Action)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.CredentialsAction_Notes_scrn), item.Notes)
	return queryPath
}