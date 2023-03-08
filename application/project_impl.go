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
// Date & Time		    : 27/11/2022 at 20:46:14
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"net/http"

	core "github.com/mt1976/ebEstimates/core"
	dao "github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// Project_Publish annouces the endpoints available for this object
func Project_Publish_Impl(mux http.ServeMux) {
	mux.HandleFunc(dm.Project_Origin_PathList, Project_ByOrigin_HandlerList)
	mux.HandleFunc(dm.Project_PathRemove, Project_HandlerRemove)
	mux.HandleFunc(dm.Project_PathUpdate, Project_HandlerUpdate)
	logs.Publish("Implementation", dm.Project_Title)

}

// Project_HandlerList is the handler for the list page
func Project_ByOrigin_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	searchID := core.GetURLparam(r, dm.Project_Origin_QueryString)

	_, Origin, _ := dao.Origin_GetByID(searchID)

	var returnList []dm.Project
	noItems, returnList, _ := dao.Project_ByOrigin_Active_GetList(Origin.Code)

	pageDetail := dm.Project_Origin_List{
		Title:       CardTitle(dm.Project_Title, core.Action_List),
		PageTitle:   PageTitle(dm.Project_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:    returnList,
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.Origin = Origin

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Project_Origin_TemplateList, w, r, pageDetail)

}

func Project_HandlerRemove(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	projectID := core.GetURLparam(r, dm.Project_QueryString)

	_, projectREC, _ := dao.ProjectAction_GetByID(projectID)

	dao.Project_SoftDelete(projectID)

	err := EstimationSession_SoftDeleteByProjectID(projectID)
	if err != nil {
		logs.Panic("Cannot SoftDelete Estimates {"+projectID+"}", err)
	}

	_, originREC, _ := dao.Origin_GetByCode(projectREC.OriginID)

	REDR := dm.Project_Origin_PathList + "?" + dm.Project_Origin_QueryString + "=" + originREC.OriginID
	http.Redirect(w, r, REDR, http.StatusFound)
}

func Project_SoftDeleteByOriginCode(originCODE string) error {
	var returnList []dm.Project
	noItems, returnList, _ := dao.Project_ByOrigin_Active_GetList(originCODE)
	if noItems > 0 {
		for _, projectREC := range returnList {
			dao.Project_SoftDelete(projectREC.ProjectID)
			err := EstimationSession_SoftDeleteByProjectID(projectREC.ProjectID)
			if err != nil {
				logs.Panic("Cannot SoftDelete Estimates {"+projectREC.ProjectID+"}", err)
			}
		}
	}
	return nil
}

// Project_HandlerUpdate is the handler used process the saving of an Project
// It is called from the Edit and New pages
// Project_HandlerUpdate  - Manually generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
func Project_HandlerUpdate(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path + r.FormValue("ProjectID"))

	item := project_DataFromRequest(r)

	dao.Project_Store(item, r)

	projectID := item.ProjectID
	_, esSessions, _ := dao.EstimationSession_Active_ByProject_GetList(projectID)
	if len(esSessions) != 0 {
		for _, esSession := range esSessions {
			trigger := "Project " + item.Name + " was updated (" + item.ProjectID + ")"
			go dao.Estimationsession_Calculate(esSession.EstimationSessionID, trigger)
		}
	}

	_, origin, _ := dao.Origin_GetByCode(item.OriginID)

	REDR := dm.Project_Origin_PathList + "?" + dm.Project_Origin_QueryString + "=" + origin.OriginID
	http.Redirect(w, r, REDR, http.StatusFound)
	//
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}
