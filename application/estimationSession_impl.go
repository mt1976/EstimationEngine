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
	"fmt"
	"math"
	"net/http"
	"strconv"

	currency "github.com/bojanz/currency"
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
	mux.HandleFunc(dm.EstimationSession_PathRemove, EstimationSession_HandlerRemove)
	mux.HandleFunc(dm.EstimationSession_PathClone, EstimationSession_HandlerClone)
	mux.HandleFunc(dm.EstimationSession_PathStore, EstimationSession_HandlerStore)
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

	symbol := "£"
	if Origin.Currency != "" {
		locale := currency.NewLocale("en")
		symbol, _ = currency.GetSymbol(Origin.Currency, locale)
	}

	pageDetail.CCY = symbol
	pageDetail.CCYCode = symbol

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
	//_, rD, _ := dao.EstimationSession_GetByID(searchID)

	rD := Estimationsession_Calculate(searchID)

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

	symbol := "£"
	if thisOrigin.Currency != "" {
		locale := currency.NewLocale("en")
		symbol, _ = currency.GetSymbol(thisOrigin.Currency, locale)
	}
	pageDetail.CCY = symbol
	pageDetail.CCYCode = symbol
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

func Estimationsession_Calculate(searchID string) dm.EstimationSession {

	logs.Processing("estimationsession_ReCalculate-->" + searchID)

	_, esRecord, _ := dao.EstimationSession_GetByID(searchID)

	// Get the list of all Features for this EstimationSession
	foundFeatures, featureList, _ := dao.Feature_Active_ByEstimationSession_GetList(searchID)
	logs.Information("Found Features: ", strconv.Itoa(foundFeatures))
	// Get the current projiect id
	_, proj, _ := dao.Project_GetByID(esRecord.ProjectID)

	//Get the current profile
	profileID := proj.ProfileID
	originID := proj.OriginID
	logs.Information("ProjectUI: ", esRecord.ProjectID)
	logs.Information("ProfileID: ", profileID)
	logs.Information("OriginID: ", originID)

	_, profile, _ := dao.Profile_GetByCode(profileID)
	//Get the current rate from the origin
	_, origin, _ := dao.Origin_GetByCode(originID)

	RoundingFactor := stf(profile.Rounding)

	HoursInDay := stf(profile.HoursPerDay)
	Rate := stf(origin.Rate)

	logs.Information("HoursInDay: ", fts(HoursInDay))
	logs.Information("Rate: ", fts(Rate))

	Total_Reqs := 0.00
	Total_AnaTest := 0.00
	Total_Docs := 0.00
	Total_Mgt := 0.00
	Total_UAT := 0.00
	Total_MKT := 0.00
	Total_Contingency := 0.00
	Total_DevEstimate := 0.00
	Total_DevUplift := 0.00
	// Range through the list of Features
	for thisFeature, feature := range featureList {
		// Sum up the estimates in this featre
		// convert feature.Reqs to INT
		logs.Information("Feature: ", strconv.Itoa(thisFeature))
		logs.Information("Feature.Reqs: ", feature.Reqs)

		Total_Reqs += stf(feature.Reqs)
		Total_AnaTest += stf(feature.AnalystTest)
		Total_Docs += stf(feature.Docs)
		Total_Mgt += stf(feature.Mgt)
		Total_UAT += stf(feature.UatSupport)
		Total_MKT += stf(feature.Marketing)
		Total_Contingency += stf(feature.Contingency)
		Total_DevUplift += stf(feature.DevUplift)
		Total_DevEstimate += stf(feature.DevEstimate)
	}

	logs.Information("Total_Reqs: ", fmt.Sprintf("%.2f", Total_Reqs))

	ReleaseHoursInDay := stf(profile.DefaultReleaseHours)
	RelHours := stf(profile.DefaultReleases) * ReleaseHoursInDay

	// Convert Hours to Days

	Total_Reqs_Days, _ := calculate(Total_Reqs, HoursInDay, Rate)
	Total_AnaTest_Days, _ := calculate(Total_AnaTest, HoursInDay, Rate)
	Total_Docs_Days, _ := calculate(Total_Docs, HoursInDay, Rate)
	Total_Mgt_Days, _ := calculate(Total_Mgt, HoursInDay, Rate)
	Total_UAT_Days, _ := calculate(Total_UAT, HoursInDay, Rate)
	Total_MKT_Days, _ := calculate(Total_MKT, HoursInDay, Rate)
	//Total_Contingency_Days, Total_Contingency_Cost := calculate(Total_Contingency, HoursInDay, Rate)
	//Total_DevEstimate_Days, Total_DevEstimate_Cost := calculate(Total_DevEstimate, HoursInDay, Rate)
	Total_DevUplift_Days, _ := calculate(Total_DevUplift, HoursInDay, Rate)
	Total_Rel_Days, _ := calculate(RelHours, ReleaseHoursInDay, Rate)

	TReqs, _ := rtn(Total_Reqs_Days+Total_AnaTest_Days+Total_Docs_Days+Total_MKT_Days, RoundingFactor)

	// Requirement Days
	esRecord.ReqDays = fts(TReqs)
	esRecord.RegCost = fts((TReqs * HoursInDay) * Rate)
	Timps, _ := rtn(Total_DevUplift_Days, RoundingFactor)
	// Implementaion Days
	esRecord.ImpDays = fts(Timps)
	esRecord.ImpCost = fts((Timps * HoursInDay) * Rate)
	// UAT Days
	Tuat, _ := rtn(Total_UAT_Days, RoundingFactor)
	esRecord.UatDays = fts(Tuat)
	esRecord.UatCost = fts((Tuat * HoursInDay) * Rate)
	// Management Days
	Tman, _ := rtn(Total_Mgt_Days, RoundingFactor)
	esRecord.MgtDays = fts(Tman)
	esRecord.MgtCost = fts((Tman * HoursInDay) * Rate)
	// Contingency Days
	//Tcon,_ := rtn(Total_Contingency, RoundingFactor)
	//esRecord.ConDays = fts(Tcon)
	//esRecord.ConCost = fts((Tcon * HoursInDay) * Rate)

	// Release Days
	Trel, _ := rtn(Total_Rel_Days, RoundingFactor)
	esRecord.RelDays = fts(Trel)
	esRecord.RelCost = fts((Trel * ReleaseHoursInDay) * Rate)

	Total_Days := TReqs + Timps + Tuat + Tman + Trel
	Total_Cost := Total_Days * HoursInDay * Rate

	SupportUplift := Total_Cost * (stf(profile.SupportUplift) / 100)
	esRecord.SupportUplift = fts(SupportUplift)

	esRecord.EffortTotal = fts(Total_Days)
	esRecord.Total = fts(Total_Cost)

	// Update the Estimation Session Record
	estUpdateTXT := dao.Translate("AuditMessage", "ESTIMATE UPDATE")

	esRecord.Notes = addActivitySystem(esRecord.Notes, estUpdateTXT)

	dao.EstimationSession_StoreSystem(esRecord)

	return esRecord
}

// rtn Rounds a number to the nearest multiple of the RoundingFactor
func rtn(number float64, RoundingFactor float64) (float64, error) {
	// Round to the nearest multiple of the RoundingFactor

	return math.Round(number/RoundingFactor) * RoundingFactor, nil
}

func EstimationSession_HandlerRemove(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	esID := core.GetURLparam(r, dm.EstimationSession_QueryString)

	_, esREC, _ := dao.EstimationSession_GetByID(esID)

	estimationSession_SoftDelete(esID)

	REDR := dm.EstimationSession_ByProject_PathList + "?" + dm.EstimationSession_ByProject_QueryString + "=" + esREC.ProjectID
	http.Redirect(w, r, REDR, http.StatusFound)
}

func estimationSession_SoftDelete(esID string) error {
	dao.EstimationSession_SoftDelete(esID)

	err := Feature_SoftDeleteByEstimationSessionID(esID)
	if err != nil {
		logs.Panic("Cannot SoftDelete Features {"+esID+"}", err)
		return err
	}
	return nil
}

func EstimationSession_SoftDeleteByProjectID(projectID string) error {
	// Get the list of Estimation Sessions
	_, esList, _ := dao.EstimationSession_Active_ByProject_GetList(projectID)

	for _, es := range esList {
		err := estimationSession_SoftDelete(es.EstimationSessionID)
		if err != nil {
			logs.Panic("Cannot SoftDelete Estimation Session {"+es.EstimationSessionID+"}", err)
			return err
		}
	}
	return nil
}

func EstimationSession_HandlerClone(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path + r.FormValue("EstimationSessionID"))

	esID := core.GetURLparam(r, dm.EstimationSession_QueryString)

	//var item dm.EstimationSession
	// START
	// Dynamically generated 28/11/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/11/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	// Clone Features
	err := EstimationSession_Clone(esID, r)
	if err != nil {
		logs.Panic("Cannot Clone Estimation Session {"+esID+"}", err)
		return
	}

	REDR := dm.EstimationSession_PathEdit + "/?" + dm.EstimationSession_QueryString + "=" + esID
	http.Redirect(w, r, REDR, http.StatusFound)
}

func EstimationSession_Clone(esID string, r *http.Request) error {
	_, esREC, _ := dao.EstimationSession_GetByID(esID)

	esREC.SYSId = ""
	cloneID := dao.EstimationSession_NewID(esREC)
	esREC.EstimationSessionID = cloneID
	esREC.Notes = r.FormValue(dm.EstimationSession_Notes_scrn)

	esREC.SYSCreated = r.FormValue(dm.EstimationSession_SYSCreated_scrn)
	esREC.SYSCreatedBy = r.FormValue(dm.EstimationSession_SYSCreatedBy_scrn)
	esREC.SYSCreatedHost = r.FormValue(dm.EstimationSession_SYSCreatedHost_scrn)
	esREC.SYSUpdated = r.FormValue(dm.EstimationSession_SYSUpdated_scrn)
	esREC.SYSUpdatedBy = r.FormValue(dm.EstimationSession_SYSUpdatedBy_scrn)
	esREC.SYSUpdatedHost = r.FormValue(dm.EstimationSession_SYSUpdatedHost_scrn)
	esREC.SYSDeleted = r.FormValue(dm.EstimationSession_SYSDeleted_scrn)
	esREC.SYSDeletedBy = r.FormValue(dm.EstimationSession_SYSDeletedBy_scrn)
	esREC.SYSDeletedHost = r.FormValue(dm.EstimationSession_SYSDeletedHost_scrn)

	esREC.Notes = addActivity(esREC.Notes, "CLONED -> "+esREC.EstimationStateID, r)

	dao.EstimationSession_Store(esREC, r)

	_, featureList, _ := dao.Feature_Active_ByEstimationSession_GetList(esID)
	for _, feature := range featureList {
		err := Feature_Clone(feature, cloneID, r)
		if err != nil {
			logs.Panic("Cannot Clone Feature {"+feature.FeatureID+"}", err)
			return err
		}
	}
	return nil
}

func stf(in string) float64 {
	out, _ := strconv.ParseFloat(in, 64)
	return out
}

func fts(in float64) string {
	out := strconv.FormatFloat(in, 'f', 2, 64)
	return out
}

func calculate(hrs float64, hrsInDay float64, rate float64) (float64, float64) {

	days := 0.00
	cost := 0.00

	if hrs != 0 {

		if hrsInDay != 0 {
			days = hrs / hrsInDay
		}

		if rate != 0 {
			cost = hrs * rate
		}
	}

	return days, cost
}

// EstimationSession_HandlerSave is the handler used process the saving of an EstimationSession
func EstimationSession_HandlerStore(w http.ResponseWriter, r *http.Request) {
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
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	item.SYSId = r.FormValue(dm.EstimationSession_SYSId_scrn)
	item.EstimationSessionID = r.FormValue(dm.EstimationSession_EstimationSessionID_scrn)
	item.ProjectID = r.FormValue(dm.EstimationSession_ProjectID_scrn)
	item.EstimationStateID = r.FormValue(dm.EstimationSession_EstimationStateID_scrn)
	item.Notes = r.FormValue(dm.EstimationSession_Notes_scrn)
	item.Releases = r.FormValue(dm.EstimationSession_Releases_scrn)
	item.Total = r.FormValue(dm.EstimationSession_Total_scrn)
	item.Contingency = r.FormValue(dm.EstimationSession_Contingency_scrn)
	item.ReqDays = r.FormValue(dm.EstimationSession_ReqDays_scrn)
	item.RegCost = r.FormValue(dm.EstimationSession_RegCost_scrn)
	item.ImpDays = r.FormValue(dm.EstimationSession_ImpDays_scrn)
	item.ImpCost = r.FormValue(dm.EstimationSession_ImpCost_scrn)
	item.UatDays = r.FormValue(dm.EstimationSession_UatDays_scrn)
	item.UatCost = r.FormValue(dm.EstimationSession_UatCost_scrn)
	item.MgtDays = r.FormValue(dm.EstimationSession_MgtDays_scrn)
	item.MgtCost = r.FormValue(dm.EstimationSession_MgtCost_scrn)
	item.RelDays = r.FormValue(dm.EstimationSession_RelDays_scrn)
	item.RelCost = r.FormValue(dm.EstimationSession_RelCost_scrn)
	item.SupportUplift = r.FormValue(dm.EstimationSession_SupportUplift_scrn)
	item.SYSCreated = r.FormValue(dm.EstimationSession_SYSCreated_scrn)
	item.SYSCreatedBy = r.FormValue(dm.EstimationSession_SYSCreatedBy_scrn)
	item.SYSCreatedHost = r.FormValue(dm.EstimationSession_SYSCreatedHost_scrn)
	item.SYSUpdated = r.FormValue(dm.EstimationSession_SYSUpdated_scrn)
	item.SYSUpdatedBy = r.FormValue(dm.EstimationSession_SYSUpdatedBy_scrn)
	item.SYSUpdatedHost = r.FormValue(dm.EstimationSession_SYSUpdatedHost_scrn)
	item.SYSDeleted = r.FormValue(dm.EstimationSession_SYSDeleted_scrn)
	item.SYSDeletedBy = r.FormValue(dm.EstimationSession_SYSDeletedBy_scrn)
	item.SYSDeletedHost = r.FormValue(dm.EstimationSession_SYSDeletedHost_scrn)
	item.Name = r.FormValue(dm.EstimationSession_Name_scrn)
	item.AdoID = r.FormValue(dm.EstimationSession_AdoID_scrn)
	item.FreshdeskID = r.FormValue(dm.EstimationSession_FreshdeskID_scrn)
	item.TrackerID = r.FormValue(dm.EstimationSession_TrackerID_scrn)
	item.EstRef = r.FormValue(dm.EstimationSession_EstRef_scrn)
	item.ExtRef = r.FormValue(dm.EstimationSession_ExtRef_scrn)
	item.SYSActivity = r.FormValue(dm.EstimationSession_SYSActivity_scrn)
	item.Origin = r.FormValue(dm.EstimationSession_Origin_scrn)
	item.OriginStateID = r.FormValue(dm.EstimationSession_OriginStateID_scrn)
	item.OriginState = r.FormValue(dm.EstimationSession_OriginState_scrn)
	item.OriginDocTypeID = r.FormValue(dm.EstimationSession_OriginDocTypeID_scrn)
	item.OriginDocType = r.FormValue(dm.EstimationSession_OriginDocType_scrn)
	item.OriginCode = r.FormValue(dm.EstimationSession_OriginCode_scrn)
	item.OriginName = r.FormValue(dm.EstimationSession_OriginName_scrn)
	item.OriginRate = r.FormValue(dm.EstimationSession_OriginRate_scrn)
	item.ProjectProfileID = r.FormValue(dm.EstimationSession_ProjectProfileID_scrn)
	item.ProjectProfile = r.FormValue(dm.EstimationSession_ProjectProfile_scrn)
	item.ProjectDefaultReleases = r.FormValue(dm.EstimationSession_ProjectDefaultReleases_scrn)
	item.ProjectDefaultReleaseHours = r.FormValue(dm.EstimationSession_ProjectDefaultReleaseHours_scrn)
	item.ProjectBlendedRate = r.FormValue(dm.EstimationSession_ProjectBlendedRate_scrn)
	item.ProjectStateID = r.FormValue(dm.EstimationSession_ProjectStateID_scrn)
	item.ProjectState = r.FormValue(dm.EstimationSession_ProjectState_scrn)
	item.ProjectName = r.FormValue(dm.EstimationSession_ProjectName_scrn)
	item.ProjectStartDate = r.FormValue(dm.EstimationSession_ProjectStartDate_scrn)
	item.ProjectEndDate = r.FormValue(dm.EstimationSession_ProjectEndDate_scrn)
	item.ProfileSupportUpliftPerc = r.FormValue(dm.EstimationSession_ProfileSupportUpliftPerc_scrn)
	item.CCY = r.FormValue(dm.EstimationSession_CCY_scrn)
	item.CCYCode = r.FormValue(dm.EstimationSession_CCYCode_scrn)
	item.EffortTotal = r.FormValue(dm.EstimationSession_EffortTotal_scrn)
	item.FreshDeskURI = r.FormValue(dm.EstimationSession_FreshDeskURI_scrn)
	item.ADOURI = r.FormValue(dm.EstimationSession_ADOURI_scrn)

	//
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	dao.EstimationSession_Store(item, r)
	REDR := dm.EstimationSession_ByProject_PathList + "?" + dm.EstimationSession_ByProject_QueryString + "=" + item.ProjectID
	http.Redirect(w, r, REDR, http.StatusFound)
}
