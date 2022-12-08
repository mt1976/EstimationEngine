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
	"math/big"
	"net/http"
	"strings"

	core "github.com/mt1976/ebEstimates/core"
	dao "github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// Project_Publish annouces the endpoints available for this object
func Feature_Publish_Impl(mux http.ServeMux) {
	mux.HandleFunc(dm.Feature_ByEstimationSession_PathList, Feature_ByEstimationSession_HandlerList)
	mux.HandleFunc(dm.Feature_SoftDelete_Path, Feature_SoftDelete_Handler)
	mux.HandleFunc(dm.Feature_PathStore, Feature_HandlerStore)
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

	Estimationsession_Calculate(featureREC.EstimationSessionID)

	REDR := dm.Feature_ByEstimationSession_PathList + "?" + dm.Feature_ByEstimationSession_QueryString + "=" + featureREC.EstimationSessionID
	http.Redirect(w, r, REDR, http.StatusFound)
}

// Feature_SoftDeleteByEstimationSessionID soft deletes all features for a given estimation session
func Feature_SoftDeleteByEstimationSessionID(esID string) error {
	_, featureList, err := dao.Feature_Active_ByEstimationSession_GetList(esID)
	for _, feature := range featureList {
		dao.Feature_SoftDelete(feature.FeatureID)
	}
	return err
}

// Feature_Clone clones a feature
func Feature_Clone(feature dm.Feature, esID string, r *http.Request) error {
	// START

	newID := dao.Feature_NewID(feature)

	feature.SYSId = ""
	feature.EstimationSessionID = esID
	feature.FeatureID = newID
	feature.SYSCreated = ""
	feature.SYSCreatedBy = ""
	feature.SYSCreatedHost = ""
	feature.SYSUpdated = ""
	feature.SYSUpdatedBy = ""
	feature.SYSUpdatedHost = ""
	feature.SYSDeleted = ""
	feature.SYSDeletedBy = ""
	feature.SYSDeletedHost = ""

	msgTXT := "CLONED -> %s"
	msgTXT = dao.Translate("AuditMessage", msgTXT)
	feature.Notes = addActivity(feature.Notes, fmt.Sprintf(msgTXT, feature.FeatureID), r)

	err := dao.Feature_Store(feature, r)
	if err != nil {
		logs.Panic("Cannot Store Feature {"+feature.FeatureID+"}", err)
		return err
	}

	Estimationsession_Calculate(esID)

	// END

	return nil
}

// Feature_HandlerSave is the handler used process the saving of an Feature
func Feature_HandlerStore(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path + r.FormValue("FeatureID"))

	item := feature_DataFromRequest(r)

	mTXT_0 := dao.Translate("AuditMessage", "VALIDATED %s")
	mTXT_0 = fmt.Sprintf(mTXT_0, strings.Repeat("-", 40))
	item.Notes = addActivity_System(item.Notes, mTXT_0)

	offProfile := false

	offProfile, item.Notes = feature_ValidateEstimate(item.DevUplift, item.DfdevUplift, offProfile, item.Notes, "Dev + Uplift")
	offProfile, item.Notes = feature_ValidateEstimate(item.Reqs, item.DfReqs, offProfile, item.Notes, "Requirements")
	offProfile, item.Notes = feature_ValidateEstimate(item.AnalystTest, item.DfAnalystTest, offProfile, item.Notes, "Analyst Test")
	offProfile, item.Notes = feature_ValidateEstimate(item.Docs, item.DfDocs, offProfile, item.Notes, "Documentation")
	offProfile, item.Notes = feature_ValidateEstimate(item.Mgt, item.Dfmgt, offProfile, item.Notes, "Management")
	offProfile, item.Notes = feature_ValidateEstimate(item.UatSupport, item.DfuatSupport, offProfile, item.Notes, "UAT Support")
	offProfile, item.Notes = feature_ValidateEstimate(item.Marketing, item.Dfmarketing, offProfile, item.Notes, "Marketing")
	offProfile, item.Notes = feature_ValidateEstimate(item.Contingency, item.Dfcontingency, offProfile, item.Notes, "Contingency")

	item.Total = feature_CalculateTotal(item)

	if offProfile {
		item.OffProfile = "true"
	}

	if item.OffProfile != "" {
		mTXT_1 := dao.Translate("AuditMessage", "ESTIMATE OFF PROFILE")
		item.Notes = addActivity(item.Notes, mTXT_1, r)
	}
	if item.OffProfileJustification != "" {
		mTXT_2 := dao.Translate("AuditMessage", "JUSTIFICATION [%s]")
		mTXT_2 = fmt.Sprintf(mTXT_2, item.OffProfileJustification)
		item.Notes = addActivity(item.Notes, mTXT_2, r)
	}
	if item.Approver != "" {
		mTXT_3 := dao.Translate("AuditMessage", "APPROVED BY [%s]")
		mTXT_3 = fmt.Sprintf(mTXT_3, item.Approver)
		item.Notes = addActivity(item.Notes, mTXT_3, r)
	}
	mTXT_4 := dao.Translate("AuditMessage", "UPDATED BY [%s]")
	mTXT_4 = fmt.Sprintf(mTXT_4, Session_GetUserName(r))
	item.Notes = addActivity(item.Notes, mTXT_4, r)

	dao.Feature_Store(item, r)

	esREC := Estimationsession_Calculate(item.EstimationSessionID)

	REDR := dm.Feature_ByEstimationSession_PathList + "/?" + dm.EstimationSession_QueryString + "=" + esREC.EstimationSessionID
	http.Redirect(w, r, REDR, http.StatusFound)
}

// feature_CalculateTotal is used to calculate the total of the feature's estimate
func feature_CalculateTotal(item dm.Feature) string {
	total := 0.0
	total += stf(item.DevUplift)
	total += stf(item.Reqs)
	total += stf(item.AnalystTest)
	total += stf(item.Docs)
	total += stf(item.Mgt)
	total += stf(item.UatSupport)
	total += stf(item.Marketing)
	total += stf(item.Contingency)
	return fts(total)
}

// feature_ValidateEstimate is used to validate the estimate of the feature
func feature_ValidateEstimate(baseEstimate string, userEstimate string, mismatch bool, notes string, info string) (bool, string) {
	//if mismatch {
	//	return mismatch, notes
	//}

	result := big.NewFloat(stf(baseEstimate)).Cmp(big.NewFloat(stf(userEstimate)))

	messageTXT := "ESTIMATE MISMATCH [%s] [%s] != [%s]"
	messageTXT = dao.Translate("AuditMessage", messageTXT)
	messageTXT = fmt.Sprintf(messageTXT, info, baseEstimate, userEstimate)

	if result != 0 {
		mismatch = true
		notes = addActivity_System(notes, messageTXT)
	}
	return mismatch, notes
}
