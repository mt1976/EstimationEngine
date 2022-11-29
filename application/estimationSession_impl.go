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
	"strconv"

	core "github.com/mt1976/ebEstimates/core"
	dao "github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// Project_Publish annouces the endpoints available for this object
func EstimationSession_Publish_Impl(mux http.ServeMux) {
	mux.HandleFunc(dm.EstimationSession_ByProject_PathList, EstimationSession_ByProject_HandlerList)
	mux.HandleFunc(dm.EstimationSession_Formatted_PathView, EstimationSession_HandlerFormatted)
	logs.Publish("Implementation", dm.EstimationSession_Title)

}

// Project_HandlerList is the handler for the list page
func EstimationSession_ByProject_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	searchID := core.GetURLparam(r, dm.EstimationSession_ByProject_QueryString)

	_, Project, _ := dao.Project_GetByID(searchID)
	_, Origin, _ := dao.Origin_GetByCode(Project.OriginID)
	var returnList []dm.EstimationSession
	noItems, returnList, _ := dao.EstimationSession_Active_ByProject_GetList(searchID)

	pageDetail := dm.EstimationSession_ByProject_List{
		Title:       CardTitle(dm.EstimationSession_Title, core.Action_List),
		PageTitle:   PageTitle(dm.EstimationSession_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:    returnList,
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.Origin = Origin
	pageDetail.Project = Project

	pageDetail.CCY = "GBP"
	pageDetail.CCYCode = "£"

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.EstimationSession_ByProject_TemplateList, w, r, pageDetail)

}

// EstimationSession_HandlerView is the handler used to View a page
func EstimationSession_HandlerFormatted(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.EstimationSession_QueryString)
	_, rD, _ := dao.EstimationSession_GetByID(searchID)

	pageDetail := dm.EstimationSession_Page{
		Title:     CardTitle(dm.EstimationSession_Title, core.Action_View),
		PageTitle: PageTitle(dm.EstimationSession_Title, core.Action_View),
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = estimationsession_PopulatePage(rD, pageDetail)
	// Populate Page Level Data
	projectID := rD.ProjectID
	_, thisProject, _ := dao.Project_GetByID(projectID)
	_, thisOrigin, _ := dao.Origin_GetByCode(thisProject.OriginID)
	_, thisOriginState, _ := dao.OriginState_GetByCode(thisOrigin.StateID)
	_, thisOriginDocType, _ := dao.DocType_GetByCode(thisOrigin.DocTypeID)
	_, thisProjectProfile, _ := dao.Profile_GetByCode(thisProject.ProfileID)
	_, thisProjectState, _ := dao.ProjectState_GetByCode(thisProject.ProjectStateID)
	pageDetail.Origin = thisOrigin.OriginID
	pageDetail.OriginCode = thisOrigin.Code
	pageDetail.OriginName = thisOrigin.FullName
	pageDetail.OriginStateID = thisOriginState.OriginStateID
	pageDetail.OriginState = thisOriginState.Name
	pageDetail.OriginDocTypeID = thisOriginDocType.DocTypeID
	pageDetail.OriginDocType = thisOriginDocType.Name
	pageDetail.OriginRate = thisOrigin.Rate
	pageDetail.ProjectProfileID = thisProject.ProfileID
	pageDetail.ProjectProfile = thisProjectProfile.Name
	pageDetail.ProjectDefaultReleases = thisProjectProfile.DefaultReleases
	pageDetail.ProjectDefaultReleaseHours = thisProjectProfile.DefaultReleaseHours
	pageDetail.ProjectBlendedRate = thisProjectProfile.BlendedRate
	pageDetail.ProjectStateID = thisProject.ProjectStateID
	pageDetail.ProjectState = thisProjectState.Name
	pageDetail.ProjectName = thisProject.Name
	pageDetail.ProjectStartDate = thisProject.StartDate
	pageDetail.ProjectEndDate = thisProject.EndDate
	pageDetail.ProfileSupportUpliftPerc = thisProjectProfile.SupportUplift
	pageDetail.CCY = "£"
	pageDetail.CCYCode = "GBP"
	pageDetail.EffortTotal = "00.00"

	var totEffort float64

	reqDays, _ := strconv.ParseFloat(rD.ReqDays, 64)
	impDays, _ := strconv.ParseFloat(rD.ImpDays, 64)

	uatDays, _ := strconv.ParseFloat(rD.UatDays, 64)
	pmDays, _ := strconv.ParseFloat(rD.MgtDays, 64)
	relDays, _ := strconv.ParseFloat(rD.RelDays, 64)

	totEffort = 0.00 + uatDays + pmDays + relDays + reqDays + impDays

	pageDetail.EffortTotal = strconv.FormatFloat(totEffort, 'f', 2, 64)

	// Add up the effort
	ExecuteTemplate(dm.EstimationSession_Formatted_TemplateView, w, r, pageDetail)

}
