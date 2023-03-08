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

	rD := dao.Estimationsession_Calculate(searchID, "View Estimation")

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

	if thisProject.DefaultRate != "" {
		pageDetail.OriginRate = thisProject.DefaultRate
	}

	if thisProject.ProjectRate != "" {
		pageDetail.OriginRate = thisProject.ProjectRate
	}

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

	pageDetail.RegCost = core.AmtFormatStr(pageDetail.RegCost)
	pageDetail.ImpCost = core.AmtFormatStr(pageDetail.ImpCost)
	pageDetail.UatCost = core.AmtFormatStr(pageDetail.UatCost)
	pageDetail.MgtCost = core.AmtFormatStr(pageDetail.MgtCost)
	pageDetail.RelCost = core.AmtFormatStr(pageDetail.RelCost)
	pageDetail.SupportUplift = core.AmtFormatStr(pageDetail.SupportUplift)
	pageDetail.Total = core.AmtFormatStr(pageDetail.Total)

	if pageDetail.ProfileSupportUpliftPerc == "" {
		pageDetail.ProfileSupportUpliftPerc = "0.00"
	}

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

	item := estimationsession_DataFromRequest(r)
	if item.Releases == "" {
		// Get the default from the project, if not there then the profile
		_, thisProject, _ := dao.Project_GetByID(item.ProjectID)
		item.Releases = thisProject.Releases
		if item.Releases == "" {
			_, thisProfile, _ := dao.Profile_GetByCode(thisProject.ProfileID)
			item.Releases = thisProfile.DefaultReleases
		}
		if item.Releases == "" {
			item.Releases = "1"
		}
	}

	msg_TXT := "CREATED -> %s"
	msg_TXT = dao.Translate("AuditMessage", msg_TXT)
	msg_TXT = fmt.Sprintf(msg_TXT, item.EstimationStateID)
	item.Activity = addActivity(item.Activity, msg_TXT, r)

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

	// Get the default from the project, if not there then the profile
	pageDetail.Releases = proj.Releases
	if pageDetail.Releases == "" {
		pageDetail.Releases = "1"
	}

	pageDetail.Name = proj.Name

	ExecuteTemplate(dm.EstimationSession_TemplateSetup, w, r, pageDetail)
}

func Estimationsession_Calculate(searchID string) dm.EstimationSession {

	logs.Processing("estimationsession_ReCalculate-->" + searchID)

	_, esRecord, estErr := dao.EstimationSession_GetByID(searchID)
	if estErr != nil {
		logs.Error("estimationsession_ReCalculate", estErr)
		return dm.EstimationSession{}
	}
	//Get Estimation State
	_, estState, esStateErr := dao.EstimationState_GetByCode(esRecord.EstimationStateID)
	if esStateErr != nil {
		logs.Error("estimationsession_ReCalculate", esStateErr)
		return dm.EstimationSession{}
	}
	if estState.IsLocked == core.TRUE {
		logs.Information("Estimation State is Locked - No Update Required", esRecord.EstimationStateID)
		return esRecord
	}

	// Get the current projiect id
	_, proj, projErr := dao.Project_GetByID(esRecord.ProjectID)
	if projErr != nil {
		logs.Error("estimationsession_ReCalculate", projErr)
		return dm.EstimationSession{}
	}
	// Get Project State
	_, projState, projStateErr := dao.ProjectState_GetByCode(proj.ProjectStateID)
	if projStateErr != nil {
		logs.Error("estimationsession_ReCalculate", projStateErr)
		return dm.EstimationSession{}
	}
	if projState.IsLocked == core.TRUE {
		logs.Information("Project State is Locked - No Update Required", esRecord.ProjectStateID)
		return esRecord
	}

	//Get the current rate from the origin
	_, origin, originErr := dao.Origin_GetByCode(proj.OriginID)
	if originErr != nil {
		logs.Error("estimationsession_ReCalculate", originErr)
		return dm.EstimationSession{}
	}
	// Get Origin State
	_, originState, originStateErr := dao.OriginState_GetByCode(esRecord.OriginStateID)
	if originStateErr != nil {
		logs.Error("estimationsession_ReCalculate", originStateErr)
		return dm.EstimationSession{}
	}
	if originState.IsLocked == core.TRUE {
		logs.Information("Origin State is Locked - No Update Required", esRecord.OriginStateID)
		return esRecord
	}

	//Get the current profile
	profileID := proj.ProfileID
	originID := proj.OriginID
	logs.Information("ProjectID", esRecord.ProjectID)
	logs.Information("ProfileID", profileID)
	logs.Information("OriginID", originID)

	_, profile, profileErr := dao.Profile_GetByCode(profileID)
	if profileErr != nil {
		logs.Error("estimationsession_ReCalculate", profileErr)
		return dm.EstimationSession{}
	}

	// Get the list of all Features for this EstimationSession
	foundFeatures, featureList, _ := dao.Feature_Active_ByEstimationSession_GetList(searchID)
	logs.Information("Features Found", strconv.Itoa(foundFeatures))

	RoundingFactor := stf(profile.Rounding)

	HoursInDay := stf(profile.HoursPerDay)
	Rate := stf(origin.Rate)
	if proj.DefaultRate != "" {
		Rate = stf(proj.DefaultRate)
	}
	if proj.ProjectRate != "" {
		Rate = stf(proj.ProjectRate)
	}

	logs.Information("HoursInDay", fts(HoursInDay))
	logs.Information("Rate", fts(Rate))

	Total_Reqs := 0.00
	Total_AnaTest := 0.00
	Total_Docs := 0.00
	Total_Mgt := 0.00
	Total_UAT := 0.00
	Total_MKT := 0.00
	Total_Training := 0.00
	//Total_Contingency := 0.00
	Total_DevEstimate := 0.00
	Total_DevUplift := 0.00

	// Range through the list of Features
	for thisFeature, feature := range featureList {
		// Sum up the estimates in this featre
		// convert feature.Reqs to INT
		logs.Break()
		logs.Information("Feature", strconv.Itoa(thisFeature))
		//	logs.Information("Feature.Reqs", feature.Reqs)

		Total_Reqs += stf(feature.Requirements)
		Total_AnaTest += stf(feature.AnalystTest)
		Total_Docs += stf(feature.Documentation)
		Total_Mgt += stf(feature.Management)
		Total_UAT += stf(feature.UatSupport)
		Total_MKT += stf(feature.Marketing)
		//Total_Contingency += stf(feature.Contingency)
		Total_DevUplift += stf(feature.DevelopmentFactored)
		Total_DevEstimate += stf(feature.DevelopmentEstimate)
		Total_Training += stf(feature.Training)
	}

	//logs.Information("Total_Reqs: ", fmt.Sprintf("%.2f", Total_Reqs))

	ReleaseHoursInDay := stf(profile.DefaultReleaseHours)
	RelHours := stf(profile.DefaultReleases) * ReleaseHoursInDay

	if esRecord.Releases != "" {
		RelHours = stf(esRecord.Releases) * ReleaseHoursInDay
	}
	// Convert Hours to Days

	Total_Reqs_Days, _ := calculate(Total_Reqs, HoursInDay, Rate)
	Total_AnaTest_Days, _ := calculate(Total_AnaTest, HoursInDay, Rate)
	Total_Docs_Days, _ := calculate(Total_Docs, HoursInDay, Rate)
	Total_Mgt_Days, _ := calculate(Total_Mgt, HoursInDay, Rate)
	Total_UAT_Days, _ := calculate(Total_UAT, HoursInDay, Rate)
	Total_MKT_Days, _ := calculate(Total_MKT, HoursInDay, Rate)
	//Total_Contingency_Days, _ := calculate(Total_Contingency, HoursInDay, Rate)
	//Total_DevEstimate_Days, Total_DevEstimate_Cost := calculate(Total_DevEstimate, HoursInDay, Rate)
	Total_DevUplift_Days, _ := calculate(Total_DevUplift, HoursInDay, Rate)
	Total_Rel_Days, _ := calculate(RelHours, ReleaseHoursInDay, Rate)
	Total_Training_Days, _ := calculate(Total_Training, HoursInDay, Rate)

	TReqs, _ := rtn(Total_Reqs_Days, RoundingFactor)

	// Requirement Days
	esRecord.ReqDays = fts(TReqs)
	esRecord.RegCost = fts((TReqs * HoursInDay) * Rate)
	Timps, _ := rtn(Total_DevUplift_Days+Total_AnaTest_Days+Total_Docs_Days+Total_MKT_Days+Total_Training_Days, RoundingFactor)
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
	//Tcon, _ := rtn(Total_Contingency_Days, RoundingFactor)
	//esRecord.Contingency = fts(Tcon)
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
	estUpdateTXT := dao.Translate("AuditMessage", "ESTIMATE UPDATED")

	esRecord.Activity = addActivity_System(esRecord.Activity, estUpdateTXT)

	dao.EstimationSession_StoreSystem(esRecord)

	return esRecord
}

// rtn Rounds a number to the nearest multiple of the RoundingFactor
func rtn(number float64, RoundingFactor float64) (float64, error) {
	// Round to the nearest multiple of the RoundingFactor
	rtnVal, rtnErr := core.RoundToNearest(number, RoundingFactor)
	return rtnVal, rtnErr
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
	uid := Session_GetUserName(r)
	newID, err := dao.EstimationSession_Clone(esID, uid)
	if err != nil {
		logs.Panic("Cannot Clone Estimation Session {"+esID+"}", err)
		return
	}

	REDR := dm.EstimationSession_PathEdit + "/?" + dm.EstimationSession_QueryString + "=" + newID
	http.Redirect(w, r, REDR, http.StatusFound)
}

func stf(in string) float64 {
	return core.StringToFloat(in)
}

func fts(in float64) string {
	return core.FloatToString(in)
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

	item := estimationsession_DataFromRequest(r)

	dao.EstimationSession_Store(item, r)
	REDR := dm.EstimationSession_ByProject_PathList + "?" + dm.EstimationSession_ByProject_QueryString + "=" + item.ProjectID
	http.Redirect(w, r, REDR, http.StatusFound)
}
