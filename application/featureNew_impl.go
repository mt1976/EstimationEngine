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

	item = Feature_Recalculate(item)
	//
	// Dynamically generated 29/11/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	//spew.Dump(item)

	tm := time.Now().Format("02/01/2006 ")
	item.Notes = tm + Session_GetUserName(r) + ": Created"

	dao.Feature_Store(item, r)

	REDR := dm.Feature_ByEstimationSession_PathList + "?" + dm.Feature_ByEstimationSession_QueryString + "=" + item.EstimationSessionID
	http.Redirect(w, r, REDR, http.StatusFound)
}

func Feature_Recalculate(item dm.Feature) dm.Feature {
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

	reqEstimate := coreEstimate * (reqPerc / 100)
	item.Reqs = strconv.FormatFloat(reqEstimate, 'f', 2, 64)

	anaEstimate := coreEstimate * (anaPerc / 100)
	item.AnalystTest = strconv.FormatFloat(anaEstimate, 'f', 2, 64)

	docEstimate := coreEstimate * (docPerc / 100)
	item.Docs = strconv.FormatFloat(docEstimate, 'f', 2, 64)

	pmEstimate := coreEstimate * (pmPerc / 100)
	item.Mgt = strconv.FormatFloat(pmEstimate, 'f', 2, 64)

	uatEstimate := coreEstimate * (uatPerc / 100)
	item.UatSupport = strconv.FormatFloat(uatEstimate, 'f', 2, 64)

	gtmEstimate := coreEstimate * (gtmPerc / 100)
	item.Marketing = strconv.FormatFloat(gtmEstimate, 'f', 2, 64)

	//totalEstimate := coreEstimate + reqEstimate + anaEstimate + docEstimate + pmEstimate + uatEstimate + gtmEstimate

	//supportUplift := totalEstimate * (supportUpliftPerc / 100)
	//item.SupportUplift = strconv.FormatFloat(supportUplift, 'f', 2, 64)
	item.Notes = "Recalculated"
	item.Notes = item.Notes + "\nBase Estimate: " + strconv.FormatFloat(baseEstimate, 'f', 2, 64)
	item.Notes = item.Notes + "\nDev Confidence: " + strconv.FormatFloat(devConfPerc, 'f', 2, 64)
	item.Notes = item.Notes + "\nCore Estimate: " + strconv.FormatFloat(coreEstimate, 'f', 2, 64)
	item.Notes = item.Notes + "\nDev: " + item.DevEstimate + " | " + "Conf: " + item.ConfidenceID + " | " + "DevUplift: " + item.DevUplift + " | " + "Req: " + item.Reqs + " | " + "Ana: " + item.AnalystTest + " | " + "Doc: " + item.Docs + " | " + "PM: " + item.Mgt + " | " + "UAT: " + item.UatSupport + " | " + "GTM: " + item.Marketing
	item.Notes = item.Notes + "\nDevConf: " + cf.Perc + " | " + "ReqPerc: " + pp.REQPerc + " | " + "AnaPerc: " + pp.ANAPerc + " | " + "DocPerc: " + pp.DOCPerc + " | " + "PMPerc: " + pp.PMPerc + " | " + "UATPerc: " + pp.UATPerc + " | " + "GTM: " + pp.GTMPerc + " | " + "SupportUplift: " + pp.SupportUplift
	item.Notes = item.Notes + "\nUpdated: " + time.Now().Format("2006-01-02 at 15:04:05")
	// Call Feature_RefreshEstimates(item)
	//
	// Dynamically generated 29/11/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return item
}
