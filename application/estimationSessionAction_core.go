package application
// ----------------------------------------------------------------
// Automatically generated  "/application/estimationsessionaction.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : EstimationSessionAction (estimationsessionaction)
// Endpoint 	        : EstimationSessionAction (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/03/2023 at 19:24:48
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

//EstimationSessionAction_Publish annouces the endpoints available for this object
func EstimationSessionAction_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.EstimationSessionAction_Path, EstimationSessionAction_Handler)
	//Cannot List via GUI
	mux.HandleFunc(dm.EstimationSessionAction_PathView, EstimationSessionAction_HandlerView)
	mux.HandleFunc(dm.EstimationSessionAction_PathEdit, EstimationSessionAction_HandlerEdit)
	mux.HandleFunc(dm.EstimationSessionAction_PathNew, EstimationSessionAction_HandlerNew)
	mux.HandleFunc(dm.EstimationSessionAction_PathSave, EstimationSessionAction_HandlerSave)
	//Cannot Delete via GUI
    core.API = core.API.AddRoute(dm.EstimationSessionAction_Title, dm.EstimationSessionAction_Path, "", dm.EstimationSessionAction_QueryString, "Application")
	logs.Publish("Application", dm.EstimationSessionAction_Title)
}

//EstimationSessionAction_HandlerView is the handler used to View a EstimationSessionAction database record
func EstimationSessionAction_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.EstimationSessionAction_QueryString)
	_, rD, _ := dao.EstimationSessionAction_GetByID(searchID)

	pageDetail := dm.EstimationSessionAction_Page{
		Title:       CardTitle(dm.EstimationSessionAction_Title, core.Action_View),
		PageTitle:   PageTitle(dm.EstimationSessionAction_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = estimationsessionaction_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("EstimationSessionAction", "View", dm.EstimationSessionAction_TemplateView)
	nextTemplate = estimationsessionaction_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//EstimationSessionAction_HandlerEdit is the handler used to Edit of an existing EstimationSessionAction database record
func EstimationSessionAction_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.EstimationSessionAction_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.EstimationSessionAction
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.EstimationSessionAction)
	} else {
		_, rD, _ = dao.EstimationSessionAction_GetByID(searchID)
	}
	
	pageDetail := dm.EstimationSessionAction_Page{
		Title:       CardTitle(dm.EstimationSessionAction_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.EstimationSessionAction_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = estimationsessionaction_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("EstimationSessionAction", "Edit", dm.EstimationSessionAction_TemplateEdit)
	nextTemplate = estimationsessionaction_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//EstimationSessionAction_HandlerSave is the handler used process the saving of an EstimationSessionAction database record, either new or existing, referenced by Edit & New Handlers.
func EstimationSessionAction_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("ID")
	logs.Servicing(r.URL.Path+itemID)

	item := estimationsessionaction_DataFromRequest(r)
	
	item, errStore := dao.EstimationSessionAction_Store(item,r)
	if errStore == nil {
		nextTemplate :=  NextTemplate("EstimationSessionAction", "Save", dm.EstimationSessionAction_Redirect)
		nextTemplate = estimationsessionaction_URIQueryData(nextTemplate,item,itemID)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.EstimationSessionAction_Name, errStore.Error())
		ExecuteRedirect(r.Referer(), w, r,dm.EstimationSessionAction_QueryString,itemID,item)
	}
}

//EstimationSessionAction_HandlerNew is the handler used process the creation of an new EstimationSessionAction database record, then redirect to Edit
func EstimationSessionAction_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.EstimationSessionAction_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.EstimationSessionAction
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.EstimationSessionAction)
	} else {
		_, _, rD, _ = dao.EstimationSessionAction_New()
	}

	pageDetail := dm.EstimationSessionAction_Page{
		Title:       CardTitle(dm.EstimationSessionAction_Title, core.Action_New),
		PageTitle:   PageTitle(dm.EstimationSessionAction_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = estimationsessionaction_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("EstimationSessionAction", "New", dm.EstimationSessionAction_TemplateNew)
	nextTemplate = estimationsessionaction_URIQueryData(nextTemplate,dm.EstimationSessionAction{},searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}	
//estimationsessionaction_PopulatePage Builds/Populates the EstimationSessionAction Page from an instance of EstimationSessionAction from the Data Model
func estimationsessionaction_PopulatePage(rD dm.EstimationSessionAction, pageDetail dm.EstimationSessionAction_Page) dm.EstimationSessionAction_Page {
	// Real DB Fields
	pageDetail.ID = rD.ID
	pageDetail.EstimationSession = rD.EstimationSession
	pageDetail.Action = rD.Action
	pageDetail.Notes = rD.Notes
	// Add Pseudo/Extra Fields, fields that are not in the DB but are used in the UI
	// Enrichment content, content used provide lookups,lists etc
	pageDetail.EstimationSession_lookup = dao.EstimationSession_GetLookup()
	pageDetail.Action_lookup = dao.EstimationState_GetLookup()
	// Add the Properties for the Fields
	pageDetail.ID_props = rD.ID_props
	pageDetail.EstimationSession_props = rD.EstimationSession_props
	pageDetail.Action_props = rD.Action_props
	pageDetail.Notes_props = rD.Notes_props
	return pageDetail
}
//estimationsessionaction_DataFromRequest is used process the content of an HTTP Request and return an instance of an EstimationSessionAction
func estimationsessionaction_DataFromRequest(r *http.Request) dm.EstimationSessionAction {
	var item dm.EstimationSessionAction
	item.ID = r.FormValue(dm.EstimationSessionAction_ID_scrn)
	item.EstimationSession = r.FormValue(dm.EstimationSessionAction_EstimationSession_scrn)
	item.Action = r.FormValue(dm.EstimationSessionAction_Action_scrn)
	item.Notes = r.FormValue(dm.EstimationSessionAction_Notes_scrn)
	return item
}
//estimationsessionaction_URIQueryData is used to replace the wildcards in the URI Query Path with the values from the EstimationSessionAction Data Model
func estimationsessionaction_URIQueryData(queryPath string,item dm.EstimationSessionAction,itemID string) string {
	if queryPath == "" {
		return ""
	}
	queryPath = core.ReplaceWildcard(queryPath, strings.ToUpper("ID"), itemID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSessionAction_ID_scrn), item.ID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSessionAction_EstimationSession_scrn), item.EstimationSession)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSessionAction_Action_scrn), item.Action)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSessionAction_Notes_scrn), item.Notes)
	return queryPath
}