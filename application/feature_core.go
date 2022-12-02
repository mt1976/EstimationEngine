package application
// ----------------------------------------------------------------
// Automatically generated  "/application/feature.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Feature (feature)
// Endpoint 	        : Feature (FeatureID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 01/12/2022 at 09:40:01
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)





//Feature_Publish annouces the endpoints available for this object
func Feature_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Feature_Path, Feature_Handler)
	mux.HandleFunc(dm.Feature_PathList, Feature_HandlerList)
	mux.HandleFunc(dm.Feature_PathView, Feature_HandlerView)
	mux.HandleFunc(dm.Feature_PathEdit, Feature_HandlerEdit)
	mux.HandleFunc(dm.Feature_PathNew, Feature_HandlerNew)
	mux.HandleFunc(dm.Feature_PathSave, Feature_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Feature_Title)
    core.Catalog_Add(dm.Feature_Title, dm.Feature_Path, "", dm.Feature_QueryString, "Application")
}


//Feature_HandlerList is the handler for the list page
func Feature_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Feature
	noItems, returnList, _ := dao.Feature_GetList()

	pageDetail := dm.Feature_PageList{
		Title:            CardTitle(dm.Feature_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Feature_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Feature_TemplateList, w, r, pageDetail)

}


//Feature_HandlerView is the handler used to View a page
func Feature_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Feature_QueryString)
	_, rD, _ := dao.Feature_GetByID(searchID)

	pageDetail := dm.Feature_Page{
		Title:       CardTitle(dm.Feature_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Feature_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = feature_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Feature_TemplateView, w, r, pageDetail)

}


//Feature_HandlerEdit is the handler used generate the Edit page
func Feature_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Feature_QueryString)
	_, rD, _ := dao.Feature_GetByID(searchID)
	
	pageDetail := dm.Feature_Page{
		Title:       CardTitle(dm.Feature_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Feature_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = feature_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Feature_TemplateEdit, w, r, pageDetail)
}


//Feature_HandlerSave is the handler used process the saving of an Feature
func Feature_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("FeatureID"))

	var item dm.Feature
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.Feature_SYSId_scrn)
		item.FeatureID = r.FormValue(dm.Feature_FeatureID_scrn)
		item.EstimationSessionID = r.FormValue(dm.Feature_EstimationSessionID_scrn)
		item.ConfidenceID = r.FormValue(dm.Feature_ConfidenceID_scrn)
		item.Name = r.FormValue(dm.Feature_Name_scrn)
		item.DevEstimate = r.FormValue(dm.Feature_DevEstimate_scrn)
		item.DevUplift = r.FormValue(dm.Feature_DevUplift_scrn)
		item.Reqs = r.FormValue(dm.Feature_Reqs_scrn)
		item.AnalystTest = r.FormValue(dm.Feature_AnalystTest_scrn)
		item.Docs = r.FormValue(dm.Feature_Docs_scrn)
		item.Mgt = r.FormValue(dm.Feature_Mgt_scrn)
		item.UatSupport = r.FormValue(dm.Feature_UatSupport_scrn)
		item.Marketing = r.FormValue(dm.Feature_Marketing_scrn)
		item.Contingency = r.FormValue(dm.Feature_Contingency_scrn)
		item.TrackerID = r.FormValue(dm.Feature_TrackerID_scrn)
		item.AdoID = r.FormValue(dm.Feature_AdoID_scrn)
		item.FreshdeskID = r.FormValue(dm.Feature_FreshdeskID_scrn)
		item.ExtRef = r.FormValue(dm.Feature_ExtRef_scrn)
		item.ExtRef2 = r.FormValue(dm.Feature_ExtRef2_scrn)
		item.SYSCreated = r.FormValue(dm.Feature_SYSCreated_scrn)
		item.SYSCreatedBy = r.FormValue(dm.Feature_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.Feature_SYSCreatedHost_scrn)
		item.SYSUpdated = r.FormValue(dm.Feature_SYSUpdated_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.Feature_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.Feature_SYSUpdatedHost_scrn)
		item.SYSDeleted = r.FormValue(dm.Feature_SYSDeleted_scrn)
		item.SYSDeletedBy = r.FormValue(dm.Feature_SYSDeletedBy_scrn)
		item.SYSDeletedHost = r.FormValue(dm.Feature_SYSDeletedHost_scrn)
		item.Developer = r.FormValue(dm.Feature_Developer_scrn)
		item.Approver = r.FormValue(dm.Feature_Approver_scrn)
		item.Notes = r.FormValue(dm.Feature_Notes_scrn)
		item.OffProfile = r.FormValue(dm.Feature_OffProfile_scrn)
		item.OffProfileJustification = r.FormValue(dm.Feature_OffProfileJustification_scrn)
		item.SYSActivity = r.FormValue(dm.Feature_SYSActivity_scrn)
	
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Feature_Store(item,r)	
	http.Redirect(w, r, dm.Feature_Redirect, http.StatusFound)
}


//Feature_HandlerNew is the handler used process the creation of an Feature
func Feature_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.Feature_New()

	pageDetail := dm.Feature_Page{
		Title:       CardTitle(dm.Feature_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Feature_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = feature_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Feature_TemplateNew, w, r, pageDetail)

}	



// Builds/Popuplates the Feature Page 
func feature_PopulatePage(rD dm.Feature, pageDetail dm.Feature_Page) dm.Feature_Page {
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.FeatureID = rD.FeatureID
	pageDetail.EstimationSessionID = rD.EstimationSessionID
	pageDetail.ConfidenceID = rD.ConfidenceID
	pageDetail.Name = rD.Name
	pageDetail.DevEstimate = rD.DevEstimate
	pageDetail.DevUplift = rD.DevUplift
	pageDetail.Reqs = rD.Reqs
	pageDetail.AnalystTest = rD.AnalystTest
	pageDetail.Docs = rD.Docs
	pageDetail.Mgt = rD.Mgt
	pageDetail.UatSupport = rD.UatSupport
	pageDetail.Marketing = rD.Marketing
	pageDetail.Contingency = rD.Contingency
	pageDetail.TrackerID = rD.TrackerID
	pageDetail.AdoID = rD.AdoID
	pageDetail.FreshdeskID = rD.FreshdeskID
	pageDetail.ExtRef = rD.ExtRef
	pageDetail.ExtRef2 = rD.ExtRef2
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.SYSDeleted = rD.SYSDeleted
	pageDetail.SYSDeletedBy = rD.SYSDeletedBy
	pageDetail.SYSDeletedHost = rD.SYSDeletedHost
	pageDetail.Developer = rD.Developer
	pageDetail.Approver = rD.Approver
	pageDetail.Notes = rD.Notes
	pageDetail.OffProfile = rD.OffProfile
	pageDetail.OffProfileJustification = rD.OffProfileJustification
	pageDetail.SYSActivity = rD.SYSActivity
	
	
	//
	// Automatically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	pageDetail.EstimationSessionID_lookup = dao.EstimationSession_GetLookup()
	
	
	
	pageDetail.ConfidenceID_lookup = dao.Confidence_GetLookup()
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.Developer_lookup = dao.Credentials_GetLookup()
	
	
	
	pageDetail.Approver_lookup = dao.Credentials_GetLookup()
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.FeatureID_props = rD.FeatureID_props
	pageDetail.EstimationSessionID_props = rD.EstimationSessionID_props
	pageDetail.ConfidenceID_props = rD.ConfidenceID_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.DevEstimate_props = rD.DevEstimate_props
	pageDetail.DevUplift_props = rD.DevUplift_props
	pageDetail.Reqs_props = rD.Reqs_props
	pageDetail.AnalystTest_props = rD.AnalystTest_props
	pageDetail.Docs_props = rD.Docs_props
	pageDetail.Mgt_props = rD.Mgt_props
	pageDetail.UatSupport_props = rD.UatSupport_props
	pageDetail.Marketing_props = rD.Marketing_props
	pageDetail.Contingency_props = rD.Contingency_props
	pageDetail.TrackerID_props = rD.TrackerID_props
	pageDetail.AdoID_props = rD.AdoID_props
	pageDetail.FreshdeskID_props = rD.FreshdeskID_props
	pageDetail.ExtRef_props = rD.ExtRef_props
	pageDetail.ExtRef2_props = rD.ExtRef2_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.SYSDeleted_props = rD.SYSDeleted_props
	pageDetail.SYSDeletedBy_props = rD.SYSDeletedBy_props
	pageDetail.SYSDeletedHost_props = rD.SYSDeletedHost_props
	pageDetail.Developer_props = rD.Developer_props
	pageDetail.Approver_props = rD.Approver_props
	pageDetail.Notes_props = rD.Notes_props
	pageDetail.OffProfile_props = rD.OffProfile_props
	pageDetail.OffProfileJustification_props = rD.OffProfileJustification_props
	pageDetail.SYSActivity_props = rD.SYSActivity_props
	
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	