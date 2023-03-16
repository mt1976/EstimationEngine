package application

// ----------------------------------------------------------------
// Automatically generated  "/application/featurenew.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : FeatureNew (featurenew)
// Endpoint 	        : FeatureNew (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 29/11/2022 at 19:40:01
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"net/http"

	core "github.com/mt1976/ebEstimates/core"
	dao "github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// FeatureNew_Publish annouces the endpoints available for this object
func FeatureNew_Publish_Impl(mux http.ServeMux) {

	//Cannot List via GUI
	mux.HandleFunc(dm.FeatureNew_PathSetup, FeatureNew_HandlerSetup)
	mux.HandleFunc(dm.FeatureNew_PathCreate, FeatureNew_HandlerCreate)
	//Cannot Delete via GUI
	logs.Publish("Implementation", dm.FeatureNew_Title)
	core.API.AddRoute(dm.FeatureNew_Title, dm.FeatureNew_Path, "", dm.FeatureNew_QueryString, "Application")
}

// FeatureNew_HandlerSetup is the handler used generate the Setup page
func FeatureNew_HandlerSetup(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path + r.FormValue(dm.EstimationSession_QueryString))

	//searchID := core.GetURLparam(r, dm.FeatureNew_QueryString)
	//_, rD, _ := dao.FeatureNew_GetByID(searchID)

	esID := core.GetURLparam(r, dm.EstimationSession_QueryString)
	logs.Processing("Estimation Session ID: " + esID)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.FeatureNew
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), esID).(dm.FeatureNew)
	} else {
		_, _, rD, _ = dao.FeatureNew_New()
	}

	pageDetail := dm.FeatureNew_Page{
		Title:     CardTitle(dm.FeatureNew_Title, core.Action_Edit),
		PageTitle: PageTitle(dm.FeatureNew_Title, core.Action_Edit),
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
	}

	// Take the data from the request (FeatureNew) and populate the page
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = featurenew_PopulatePage(rD, pageDetail)
	pageDetail.EstimationSessionID = esID

	_, es, _ := dao.EstimationSession_GetByID(esID)
	_, proj, _ := dao.Project_GetByID(es.ProjectID)

	pageDetail.Name = proj.Name
	pageDetail.ProfileDefault = proj.ProfileID
	pageDetail.ProfileSelected = proj.ProfileID

	pageDetail.ProjectName = proj.Name
	pageDetail.ProjectID = proj.ProjectID
	pageDetail.OriginCode = proj.OriginID
	pageDetail.OriginID = dao.CacheRead(dm.Origin_Name, proj.OriginID).(dm.Origin).OriginID
	pageDetail.EstimationSessionName = es.Name

	nextTemplate := NextTemplate("FeatureNew", "Setup", dm.FeatureNew_TemplateSetup)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

// FeatureNew_HandlerSave is the handler used process the saving of an FeatureNew
func FeatureNew_HandlerCreate(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	logs.Processing(r.URL.Path)

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path + r.FormValue(dm.Feature_QueryString))

	var item dm.Feature
	//var fn dm.FeatureNew

	fn := featurenew_DataFromRequest(r)
	// START
	// Dynamically generated 29/11/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	newID := dao.Feature_NewID(item)

	logs.Processing("New ID: " + newID)

	item.FeatureID = newID
	item.EstimationSessionID = r.FormValue(dm.FeatureNew_EstimationSessionID_scrn)
	item.Name = r.FormValue(dm.FeatureNew_Name_scrn)
	item.DevelopmentEstimate = r.FormValue(dm.FeatureNew_DeveloperEstimate_scrn)
	item.ConfidenceCODE = r.FormValue(dm.FeatureNew_ConfidenceCODE_scrn)
	item.DeveloperResource = r.FormValue(dm.FeatureNew_DeveloperResource_scrn)

	fop := featurenew_DataFromRequest(r)

	est := dao.CacheRead(dm.EstimationSession_Name, fop.EstimationSessionID).(dm.EstimationSession)

	item.EstimationSessionID = fop.EstimationSessionID
	item.Name = fop.Name
	item.DevelopmentEstimate = fop.DeveloperEstimate
	item.AnalystEstimate = fop.AnalystEstimate
	item.ConfidenceCODE = fop.ConfidenceCODE
	item.DeveloperResource = fop.DeveloperResource
	item.AnalystResource = fop.AnalystResource
	item.ProjectManagerResource = est.ProjectManager
	item.ProductManagerResource = est.ProductManager
	item.AdoID = fop.DevOpsID
	item.FreshdeskID = fop.FreshDeskID
	item.TrackerID = fop.RSCID
	item.ExtRef = fop.OtherID
	item.ExtRef2 = fop.OtherID2
	item.ProfileDefault = fop.ProfileDefault
	item.ProfileSelected = fop.ProfileSelected
	item.EstimateEffort = fop.EstimateEffort
	if item.EstimateEffort == "" {
		item.EstimateEffort = "0"
	}
	item.OffProfile = core.FALSE

	item = dao.Feature_CalcDefaults(item, true, Session_GetUserName(r))

	msgTXT := "New Feature Created"
	msgTXT = dao.Translate("Message", msgTXT)
	item.Activity = addActivity(item.Activity, msgTXT, r)
	item.SYSActivity = addActivity(item.SYSActivity, msgTXT, r)

	// logs.Warning("FeatureNew_HandlerCreate: IN  " + item.FeatureID)
	// logs.Warning("FeatureNew_HandlerCreate: IN  " + item.FeatureID)
	// logs.Warning("FeatureNew_HandlerCreate: IN  " + item.FeatureID)
	// logs.Warning("FeatureNew_HandlerCreate: IN  " + item.FeatureID)
	// _, err := dao.Feature_Store(item, r)
	// logs.Warning("FeatureNew_HandlerCreate: OUT " + item.FeatureID)
	// logs.Warning("FeatureNew_HandlerCreate: OUT " + item.FeatureID)
	// logs.Warning("FeatureNew_HandlerCreate: OUT " + item.FeatureID)
	// logs.Warning("FeatureNew_HandlerCreate: OUT " + item.FeatureID)

	// if err != nil {
	// 	logs.Warning(err.Error())
	// }

	// //time.Sleep(30 * time.Second) // Sleeps for 30 seconds
	// //dao.Feature_Validate(item, err)

	// //go Estimationsession_Calculate(item.EstimationSessionID)

	// //REDR := dm.Feature_ByEstimationSession_PathList + "?" + dm.Feature_ByEstimationSession_QueryString + "=" + item.EstimationSessionID

	// nextTemplate := NextTemplate("FeatureNew", "Create", dm.Feature_PathEdit)

	// REDR := nextTemplate + "?" + dm.Feature_QueryString + "=" + item.FeatureID
	// http.Redirect(w, r, REDR, http.StatusFound)

	item, errStore := dao.Feature_Store(item, r)
	if errStore == nil {
		REDR := dm.Feature_PathEdit + "?" + dm.Feature_QueryString + "=" + item.FeatureID
		nextTemplate := NextTemplate("FeatureNew", "Save", REDR)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.Feature_Name, errStore.Error())
		//http.Redirect(w, r, r.Referer(), http.StatusFound)
		ExecuteRedirect(r.Referer(), w, r, dm.EstimationSession_QueryString, item.EstimationSessionID, fn)
	}

}
