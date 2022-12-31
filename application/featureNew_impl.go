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
	"fmt"
	"net/http"
	"strconv"
	"time"

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
	core.Catalog_Add(dm.FeatureNew_Title, dm.FeatureNew_Path, "", dm.FeatureNew_QueryString, "Application")
}

// FeatureNew_HandlerEdit is the handler used generate the Edit page
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
	var rD dm.FeatureNew

	esID := core.GetURLparam(r, dm.EstimationSession_QueryString)
	logs.Processing("Estimation Session ID: " + esID)

	pageDetail := dm.FeatureNew_Page{
		Title:     CardTitle(dm.FeatureNew_Title, core.Action_Edit),
		PageTitle: PageTitle(dm.FeatureNew_Title, core.Action_Edit),
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = featurenew_PopulatePage(rD, pageDetail)
	pageDetail.EstimationSession = esID

	_, es, _ := dao.EstimationSession_GetByID(esID)
	_, proj, _ := dao.Project_GetByID(es.ProjectID)

	pageDetail.Name = proj.Name

	ExecuteTemplate(dm.FeatureNew_TemplateSetup, w, r, pageDetail)
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
	// START
	// Dynamically generated 29/11/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	newID := dao.Feature_NewID(item)

	logs.Processing("New ID: " + newID)

	item.FeatureID = newID
	item.EstimationSessionID = r.FormValue(dm.FeatureNew_EstimationSession_scrn)
	item.Name = r.FormValue(dm.FeatureNew_Name_scrn)
	item.DevEstimate = r.FormValue(dm.FeatureNew_DevEstimate_scrn)
	item.ConfidenceID = r.FormValue(dm.FeatureNew_Confidence_scrn)
	item.Developer = r.FormValue(dm.FeatureNew_Developer_scrn)

	item = Feature_CalcDefaults(item)
	//
	// Dynamically generated 29/11/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	//spew.Dump(item)

	//tm := time.Now().Format("02/01/2006 ")
	//item.Notes = tm + Session_GetUserName(r) + ": Created"

	msgTXT := "New Feature Created"
	item.Notes = addActivity(item.Notes, msgTXT, r)

	dao.Feature_Store(item, r)

	Estimationsession_Calculate(item.EstimationSessionID)

	REDR := dm.Feature_ByEstimationSession_PathList + "?" + dm.Feature_ByEstimationSession_QueryString + "=" + item.EstimationSessionID
	http.Redirect(w, r, REDR, http.StatusFound)
}

func Feature_CalcDefaults(item dm.Feature) dm.Feature {
	// START
	// Dynamically generated 29/11/2022 by matttownsend (Matt Townsend) on silicon.local
	//

	// Get Project Profile for this Estimation Session
	// Get the Estimation Session
	_, es, _ := dao.EstimationSession_GetByID(item.EstimationSessionID)
	// Then Get the Project
	_, p, _ := dao.Project_GetByID(es.ProjectID)
	// Then Get the Profile
	_, pp, _ := dao.Profile_GetByCode(p.ProfileID)
	item.DefaultProfile = p.ProfileID
	if item.ActualProfile == "" {
		item.ActualProfile = p.ProfileID
	}
	//spew.Dump(pp)

	if item.DevEstimate == "" {
		return item
	}

	baseEstimate, _ := strconv.ParseFloat(item.DevEstimate, 64)

	reqPerc, _ := strconv.ParseFloat(pp.REQPerc, 64)
	anaPerc, _ := strconv.ParseFloat(pp.ANAPerc, 64)
	docPerc, _ := strconv.ParseFloat(pp.DOCPerc, 64)
	pmPerc, _ := strconv.ParseFloat(pp.PMPerc, 64)
	uatPerc, _ := strconv.ParseFloat(pp.UATPerc, 64)
	gtmPerc, _ := strconv.ParseFloat(pp.GTMPerc, 64)
	trainPerc, _ := strconv.ParseFloat(pp.TrainingPerc, 64)
	//supportUpliftPerc, _ := strconv.ParseFloat(pp.SupportUplift, 64)

	// Get the Confidence Factor
	_, cf, _ := dao.Confidence_GetByCode(item.ConfidenceID)
	devConfPerc, _ := strconv.ParseFloat(cf.Perc, 64)

	var coreEstimate float64
	if devConfPerc == 0 {
		coreEstimate = baseEstimate
	} else {
		coreEstimate = baseEstimate + (baseEstimate * (devConfPerc / 100))
	}

	item.DevUplift = strconv.FormatFloat(coreEstimate, 'f', 2, 64)
	item.DfdevUplift = item.DevUplift

	reqEstimate := coreEstimate * (reqPerc / 100)
	item.Reqs = strconv.FormatFloat(reqEstimate, 'f', 2, 64)
	item.DfReqs = item.Reqs

	anaEstimate := coreEstimate * (anaPerc / 100)
	item.AnalystTest = strconv.FormatFloat(anaEstimate, 'f', 2, 64)
	item.DfAnalystTest = item.AnalystTest

	docEstimate := coreEstimate * (docPerc / 100)
	item.Docs = strconv.FormatFloat(docEstimate, 'f', 2, 64)
	item.DfDocs = item.Docs

	pmEstimate := coreEstimate * (pmPerc / 100)
	item.Mgt = strconv.FormatFloat(pmEstimate, 'f', 2, 64)
	item.Dfmgt = item.Mgt

	uatEstimate := coreEstimate * (uatPerc / 100)
	item.UatSupport = strconv.FormatFloat(uatEstimate, 'f', 2, 64)
	item.DfuatSupport = item.UatSupport

	gtmEstimate := coreEstimate * (gtmPerc / 100)
	item.Marketing = strconv.FormatFloat(gtmEstimate, 'f', 2, 64)
	item.Dfmarketing = item.Marketing

	trainEstimate := coreEstimate * (trainPerc / 100)
	item.Training = strconv.FormatFloat(trainEstimate, 'f', 2, 64)
	item.DfTraining = item.Training

	//totalEstimate := coreEstimate + reqEstimate + anaEstimate + docEstimate + pmEstimate + uatEstimate + gtmEstimate
	item.Total = feature_CalculateTotal(item)
	//supportUplift := totalEstimate * (supportUpliftPerc / 100)
	//item.SupportUplift = strconv.FormatFloat(supportUplift, 'f', 2, 64)
	msgTXT := "Recalculated"
	item.Notes = addActivity_System(item.Notes, msgTXT)

	//item.Notes = item.Notes + "\nBase Estimate: " + strconv.FormatFloat(baseEstimate, 'f', 2, 64)
	msgTXT = "Base Estimate: %s"
	msgTXT = dao.Translate("AuditMessage", msgTXT)
	msgTXT = fmt.Sprintf(msgTXT, fts(baseEstimate))
	item.Notes = addActivity_System(item.Notes, msgTXT)

	//item.Notes = item.Notes + "\nDev Confidence: " + strconv.FormatFloat(devConfPerc, 'f', 2, 64)
	msgTXT = "Dev Confidence: %s"
	msgTXT = dao.Translate("AuditMessage", msgTXT)
	msgTXT = fmt.Sprintf(msgTXT, fts(devConfPerc))
	item.Notes = addActivity_System(item.Notes, msgTXT)

	//item.Notes = item.Notes + "\nCore Estimate: " + strconv.FormatFloat(coreEstimate, 'f', 2, 64)
	msgTXT = "Core Estimate: %s"
	msgTXT = dao.Translate("AuditMessage", msgTXT)
	msgTXT = fmt.Sprintf(msgTXT, fts(coreEstimate))
	item.Notes = addActivity_System(item.Notes, msgTXT)

	//item.Notes = item.Notes + "\nDev: " + item.DevEstimate + " | " + "Conf: " + item.ConfidenceID + " | " + "DevUplift: " + item.DevUplift + " | " + "Req: " + item.Reqs + " | " + "Ana: " + item.AnalystTest + " | " + "Doc: " + item.Docs + " | " + "PM: " + item.Mgt + " | " + "UAT: " + item.UatSupport + " | " + "GTM: " + item.Marketing
	msgTXT = "Dev: %s | Conf: %s | DevUplift: %s | Req: %s | Ana: %s | Doc: %s | PM: %s | UAT: %s | GTM: %s"
	msgTXT = dao.Translate("AuditMessage", msgTXT)
	msgTXT = fmt.Sprintf(msgTXT, item.DevEstimate, item.ConfidenceID, item.DevUplift, item.Reqs, item.AnalystTest, item.Docs, item.Mgt, item.UatSupport, item.Marketing)
	item.Notes = addActivity_System(item.Notes, msgTXT)

	//item.Notes = item.Notes + "\nDevConf: " + cf.Perc + " | " + "ReqPerc: " + pp.REQPerc + " | " + "AnaPerc: " + pp.ANAPerc + " | " + "DocPerc: " + pp.DOCPerc + " | " + "PMPerc: " + pp.PMPerc + " | " + "UATPerc: " + pp.UATPerc + " | " + "GTM: " + pp.GTMPerc + " | " + "SupportUplift: " + pp.SupportUplift
	msgTXT = "DevConf: %s | ReqPerc: %s | AnaPerc: %s | DocPerc: %s | PMPerc: %s | UATPerc: %s | GTM: %s | SupportUplift: %s"
	msgTXT = dao.Translate("AuditMessage", msgTXT)
	msgTXT = fmt.Sprintf(msgTXT, cf.Perc, pp.REQPerc, pp.ANAPerc, pp.DOCPerc, pp.PMPerc, pp.UATPerc, pp.GTMPerc, pp.SupportUplift)
	item.Notes = addActivity_System(item.Notes, msgTXT)

	//Total
	//item.Notes = item.Notes + "\nTotal: " + item.Total
	msgTXT = "Total: %s"
	msgTXT = dao.Translate("AuditMessage", msgTXT)
	msgTXT = fmt.Sprintf(msgTXT, item.Total)
	item.Notes = addActivity_System(item.Notes, msgTXT)

	//item.Notes = item.Notes + "\nUpdated: " + time.Now().Format("2006-01-02 at 15:04:05")
	msgTXT = "Updated: %s"
	msgTXT = dao.Translate("AuditMessage", msgTXT)
	msgTXT = fmt.Sprintf(msgTXT, time.Now().Format("2006-01-02 at 15:04:05"))
	item.Notes = addActivity_System(item.Notes, msgTXT)
	// Call Feature_RefreshEstimates(item)
	//
	// Dynamically generated 29/11/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return item
}
