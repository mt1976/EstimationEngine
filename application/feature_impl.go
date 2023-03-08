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
	mux.HandleFunc(dm.Feature_PathStore, Feature_HandlerStore)
	mux.HandleFunc(dm.Feature_PathCreate, Feature_HandlerCreate)
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

	pageDetail.CCY = Origin.Currency
	pageDetail.CCYCode = dao.Financial_CCYtoSymbol(pageDetail.CCY)

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
	trigger := "Deleted Feature " + featureREC.Name + " (" + featureREC.FeatureID + ")"
	go dao.Estimationsession_Calculate(featureREC.EstimationSessionID, trigger)

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

// Feature_HandlerStore is the handler used process the saving of an Feature
func Feature_HandlerStoreOLD(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Information("Feature_HandlerStore", "Feature_HandlerStore")
	logs.Information("Feature_HandlerStore", "Feature_HandlerStore")
	logs.Information("Feature_HandlerStore", "Feature_HandlerStore")
	logs.Information("Feature_HandlerStore", "Feature_HandlerStore")
	logs.Information("Feature_HandlerStore", "Feature_HandlerStore")
	logs.Information("Feature_HandlerStore", "Feature_HandlerStore")

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path + r.FormValue("FeatureID"))

	item := feature_DataFromRequest(r)

	dao.Feature_Store(item, r)
	trigger := "Update " + dm.Feature_Title + " " + item.FeatureID
	esREC := dao.Estimationsession_Calculate(item.EstimationSessionID, trigger)

	REDR := dm.Feature_ByEstimationSession_PathList + "/?" + dm.EstimationSession_QueryString + "=" + esREC.EstimationSessionID
	http.Redirect(w, r, REDR, http.StatusFound)

}

// Feature_HandlerSave is the handler used process the saving of an Feature
// It is called from the Edit and New pages
// Feature_HandlerSave  - Auto generated 07/03/2023 by matttownsend (Matt Townsend) on silicon.local
func Feature_HandlerStore(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 07/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("FeatureID")
	logs.Servicing(r.URL.Path + itemID)

	item := feature_DataFromRequest(r)

	item, errStore := dao.Feature_Store(item, r)
	if errStore == nil {
		nextTemplate := NextTemplate("Feature", "Save", dm.Feature_Redirect)
		nextTemplate = core.ReplaceWildcard(nextTemplate, "!ID", item.EstimationSessionID)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.Feature_Name, errStore.Error())
		//http.Redirect(w, r, r.Referer(), http.StatusFound)
		ExecuteRedirect(r.Referer(), w, r, dm.Feature_QueryString, itemID, item)
	}
	//
	// Auto generated 07/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}

// Feature_HandlerCreate is the handler used process the creation of an Feature
// It will create a new Feature and then redirect to the Edit page
// Feature_HandlerNew  - Auto generated 05/03/2023 by matttownsend (Matt Townsend) on silicon.local
func Feature_HandlerCreate(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 05/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Feature_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.Feature
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.Feature)
	} else {
		_, _, rD, _ = dao.Feature_New()
	}

	pageDetail := dm.Feature_Page{
		Title:     CardTitle(dm.Feature_Title, core.Action_New),
		PageTitle: PageTitle(dm.Feature_Title, core.Action_New),
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = feature_PopulatePage(rD, pageDetail)

	ExecuteTemplate(dm.Feature_TemplateCreate, w, r, pageDetail)
	//
	// Auto generated 05/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}

// Feature_Clone clones a feature
func Feature_Clone(feature dm.Feature, esID string, r *http.Request) error {
	// START

	uid := Session_GetUserName(r)
	err := dao.Feature_Clone(feature, esID, uid)
	if err != nil {
		logs.Error("Feature_Clone", err)
		return err
	}
	// END

	return nil
}
