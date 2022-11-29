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
