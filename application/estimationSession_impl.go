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
	mux.HandleFunc(dm.EstimationSession_PathSetup, EstimationSession_HandlerSetup)
	mux.HandleFunc(dm.EstimationSession_PathCreate, EstimationSession_HandlerCreate)
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

// EstimationSession_HandlerSave is the handler used process the saving of an EstimationSession
func EstimationSession_HandlerCreate(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path + r.FormValue("EstimationSessionID"))

	var item dm.EstimationSession
	// START
	// Dynamically generated 28/11/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	item.SYSId = r.FormValue(dm.EstimationSession_SYSId_scrn)
	item.EstimationSessionID = r.FormValue(dm.EstimationSession_EstimationSessionID_scrn)
	item.ProjectID = r.FormValue(dm.EstimationSession_ProjectID_scrn)
	item.EstimationStateID = r.FormValue(dm.EstimationSession_EstimationStateID_scrn)
	item.Notes = r.FormValue(dm.EstimationSession_Notes_scrn)
	// item.Releases = r.FormValue(dm.EstimationSession_Releases_scrn)
	// item.Total = r.FormValue(dm.EstimationSession_Total_scrn)
	// item.Contingency = r.FormValue(dm.EstimationSession_Contingency_scrn)
	// item.ReqDays = r.FormValue(dm.EstimationSession_ReqDays_scrn)
	// item.RegCost = r.FormValue(dm.EstimationSession_RegCost_scrn)
	// item.ImpDays = r.FormValue(dm.EstimationSession_ImpDays_scrn)
	// item.ImpCost = r.FormValue(dm.EstimationSession_ImpCost_scrn)
	// item.UatDays = r.FormValue(dm.EstimationSession_UatDays_scrn)
	// item.UatCost = r.FormValue(dm.EstimationSession_UatCost_scrn)
	// item.MgtDays = r.FormValue(dm.EstimationSession_MgtDays_scrn)
	// item.MgtCost = r.FormValue(dm.EstimationSession_MgtCost_scrn)
	// item.RelDays = r.FormValue(dm.EstimationSession_RelDays_scrn)
	// item.RelCost = r.FormValue(dm.EstimationSession_RelCost_scrn)
	//item.SupportUplift = r.FormValue(dm.EstimationSession_SupportUplift_scrn)
	// item.SYSCreated = r.FormValue(dm.EstimationSession_SYSCreated_scrn)
	// item.SYSCreatedBy = r.FormValue(dm.EstimationSession_SYSCreatedBy_scrn)
	// item.SYSCreatedHost = r.FormValue(dm.EstimationSession_SYSCreatedHost_scrn)
	// item.SYSUpdated = r.FormValue(dm.EstimationSession_SYSUpdated_scrn)
	// item.SYSUpdatedBy = r.FormValue(dm.EstimationSession_SYSUpdatedBy_scrn)
	// item.SYSUpdatedHost = r.FormValue(dm.EstimationSession_SYSUpdatedHost_scrn)
	// item.SYSDeleted = r.FormValue(dm.EstimationSession_SYSDeleted_scrn)
	// item.SYSDeletedBy = r.FormValue(dm.EstimationSession_SYSDeletedBy_scrn)
	// item.SYSDeletedHost = r.FormValue(dm.EstimationSession_SYSDeletedHost_scrn)
	item.Name = r.FormValue(dm.EstimationSession_Name_scrn)
	item.AdoID = r.FormValue(dm.EstimationSession_AdoID_scrn)
	item.FreshdeskID = r.FormValue(dm.EstimationSession_FreshdeskID_scrn)
	item.TrackerID = r.FormValue(dm.EstimationSession_TrackerID_scrn)
	item.EstRef = r.FormValue(dm.EstimationSession_EstRef_scrn)
	item.ExtRef = r.FormValue(dm.EstimationSession_ExtRef_scrn)
	// item.Origin = r.FormValue(dm.EstimationSession_Origin_scrn)
	// item.OriginStateID = r.FormValue(dm.EstimationSession_OriginStateID_scrn)
	// item.OriginState = r.FormValue(dm.EstimationSession_OriginState_scrn)
	// item.OriginDocTypeID = r.FormValue(dm.EstimationSession_OriginDocTypeID_scrn)
	// item.OriginDocType = r.FormValue(dm.EstimationSession_OriginDocType_scrn)
	// item.OriginCode = r.FormValue(dm.EstimationSession_OriginCode_scrn)
	// item.OriginName = r.FormValue(dm.EstimationSession_OriginName_scrn)
	// item.OriginRate = r.FormValue(dm.EstimationSession_OriginRate_scrn)
	// item.ProjectProfileID = r.FormValue(dm.EstimationSession_ProjectProfileID_scrn)
	// item.ProjectProfile = r.FormValue(dm.EstimationSession_ProjectProfile_scrn)
	// item.ProjectDefaultReleases = r.FormValue(dm.EstimationSession_ProjectDefaultReleases_scrn)
	// item.ProjectDefaultReleaseHours = r.FormValue(dm.EstimationSession_ProjectDefaultReleaseHours_scrn)
	// item.ProjectBlendedRate = r.FormValue(dm.EstimationSession_ProjectBlendedRate_scrn)
	// item.ProjectStateID = r.FormValue(dm.EstimationSession_ProjectStateID_scrn)
	// item.ProjectState = r.FormValue(dm.EstimationSession_ProjectState_scrn)
	// item.ProjectName = r.FormValue(dm.EstimationSession_ProjectName_scrn)
	// item.ProjectStartDate = r.FormValue(dm.EstimationSession_ProjectStartDate_scrn)
	// item.ProjectEndDate = r.FormValue(dm.EstimationSession_ProjectEndDate_scrn)
	// item.ProfileSupportUpliftPerc = r.FormValue(dm.EstimationSession_ProfileSupportUpliftPerc_scrn)
	// item.CCY = r.FormValue(dm.EstimationSession_CCY_scrn)
	// item.CCYCode = r.FormValue(dm.EstimationSession_CCYCode_scrn)
	// item.EffortTotal = r.FormValue(dm.EstimationSession_EffortTotal_scrn)
	// item.FreshDeskURI = r.FormValue(dm.EstimationSession_FreshDeskURI_scrn)
	// item.ADOURI = r.FormValue(dm.EstimationSession_ADOURI_scrn)

	//
	// Dynamically generated 28/11/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	item.Notes = addActivity(item.Notes, "CREATED -> "+item.EstimationStateID, r)

	dao.EstimationSession_Store(item, r)
	REDR := dm.EstimationSession_ByProject_PathList + "/?" + dm.Project_QueryString + "=" + item.ProjectID
	http.Redirect(w, r, REDR, http.StatusFound)
}

// EstimationSession_HandlerNew is the handler used process the creation of an EstimationSession
func EstimationSession_HandlerSetup(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.EstimationSession_New()

	pageDetail := dm.EstimationSession_Page{
		Title:     CardTitle(dm.EstimationSession_Title, core.Action_New),
		PageTitle: PageTitle(dm.EstimationSession_Title, core.Action_New),
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	projectID := core.GetURLparam(r, dm.Project_QueryString)
	if projectID == "" {
		ExecuteTemplate("html/Impl_Oops.html", w, r, pageDetail)
		return
	}

	pageDetail = estimationsession_PopulatePage(rD, pageDetail)
	pageDetail.EstimationStateID = "NEWE"
	pageDetail.ProjectID = projectID

	_, proj, _ := dao.Project_GetByID(projectID)

	pageDetail.Name = proj.Name

	ExecuteTemplate(dm.EstimationSession_TemplateSetup, w, r, pageDetail)
}
