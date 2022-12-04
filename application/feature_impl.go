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
func Feature_Publish_Impl(mux http.ServeMux) {
	mux.HandleFunc(dm.Feature_ByEstimationSession_PathList, Feature_ByEstimationSession_HandlerList)
	mux.HandleFunc(dm.Feature_SoftDelete_Path, Feature_SoftDelete_Handler)
	logs.Publish("Implementation", dm.Feature_Title)

}

// Project_HandlerList is the handler for the list page
func Feature_ByEstimationSession_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	searchID := core.GetURLparam(r, dm.Feature_ByEstimationSession_QueryString)
	_, EstimationSession, _ := dao.EstimationSession_GetByID(searchID)
	_, Project, _ := dao.Project_GetByID(EstimationSession.ProjectID)
	_, Origin, _ := dao.Origin_GetByCode(Project.OriginID)
	var returnList []dm.Feature
	noItems, returnList, _ := dao.Feature_Active_ByEstimationSession_GetList(searchID)

	pageDetail := dm.Feature_ByEstimationSession_List{
		Title:       CardTitle(dm.Feature_Title, core.Action_List),
		PageTitle:   PageTitle(dm.Feature_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:    returnList,
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.EstimationSession = EstimationSession
	pageDetail.Origin = Origin
	pageDetail.Project = Project

	pageDetail.CCY = "GBP"
	pageDetail.CCYCode = "£"

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Feature_ByEstimationSession_TemplateList, w, r, pageDetail)

}

// Project_HandlerList is the handler for the list page
func Feature_SoftDelete_Handler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	featureID := core.GetURLparam(r, dm.Feature_SoftDelete_QueryString)

	_, featureREC, _ := dao.Feature_GetByID(featureID)

	dao.Feature_SoftDelete(featureID)

	REDR := dm.Feature_ByEstimationSession_PathList + "?" + dm.Feature_ByEstimationSession_QueryString + "=" + featureREC.EstimationSessionID
	http.Redirect(w, r, REDR, http.StatusFound)
	//ExecuteTemplate(dm.Feature_ByEstimationSession_TemplateList, w, r, pageDetail)

}

func Feature_SoftDeleteByEstimationSessionID(esID string) error {
	_, featureList, err := dao.Feature_Active_ByEstimationSession_GetList(esID)
	for _, feature := range featureList {
		dao.Feature_SoftDelete(feature.FeatureID)
	}
	return err
}
