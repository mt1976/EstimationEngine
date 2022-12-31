package application

// ----------------------------------------------------------------
// Automatically generated  "/application/estimationsessionaction.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : EstimationSessionAction (estimationsessionaction)
// Endpoint 	        : EstimationSessionAction (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 29/11/2022 at 13:02:22
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"net/http"
	"strconv"

	core "github.com/mt1976/ebEstimates/core"
	dao "github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// EstimationSessionAction_Publish annouces the endpoints available for this object
func EstimationSessionAction_Publish_Impl(mux http.ServeMux) {
	mux.HandleFunc(dm.EstimationSessionAction_PathAction, EstimationSessionAction_HandlerAction)
	mux.HandleFunc(dm.EstimationSessionAction_PathStore, EstimationSessionAction_HandlerStore)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.EstimationSessionAction_Title)
	core.Catalog_Add(dm.EstimationSessionAction_Title, dm.EstimationSessionAction_Path, "", dm.EstimationSessionAction_QueryString, "Application")
}

// EstimationSessionAction_HandlerEdit is the handler used generate the Edit page
func EstimationSessionAction_HandlerAction(w http.ResponseWriter, r *http.Request) {
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

	sessionID := core.GetURLparam(r, "EstimationSessionID")
	stateID := core.GetURLparam(r, "EstimationStateID")

	pageDetail := dm.EstimationSessionAction_Page{
		Title:     CardTitle(dm.EstimationSessionAction_Title, core.Action_Edit),
		PageTitle: PageTitle(dm.EstimationSessionAction_Title, core.Action_Edit),
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = estimationsessionaction_PopulatePage(rD, pageDetail)

	pageDetail.EstimationSession = sessionID
	pageDetail.Action = stateID
	// Get Current Date

	ExecuteTemplate(dm.EstimationSessionAction_TemplateEdit, w, r, pageDetail)
}

// EstimationSessionAction_HandlerSave is the handler used process the saving of an EstimationSessionAction
func EstimationSessionAction_HandlerStore(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path + r.FormValue("ID"))

	item := estimationsessionaction_DataFromRequest(r)

	logs.Information("EstimationSessionID", item.EstimationSession)
	noRet, es, _ := dao.EstimationSession_GetByID(item.EstimationSession)
	logs.Information("NoFound", strconv.Itoa(noRet))
	//spew.Dump(es)

	msgTXT := es.EstimationStateID + "->" + item.Action
	item.Notes = addActivity_System(item.Notes, msgTXT)

	es.EstimationStateID = item.Action

	dao.EstimationSession_Store(es, r)

	// dao.EstimationSessionAction_Store(item, r)
	REDR := dm.EstimationSession_Formatted_PathView + "?" + dm.EstimationSession_QueryString + "=" + item.EstimationSession

	http.Redirect(w, r, REDR, http.StatusFound)
}
