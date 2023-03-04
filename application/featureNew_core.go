package application
// ----------------------------------------------------------------
// Automatically generated  "/application/featurenew.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : FeatureNew (featurenew)
// Endpoint 	        : FeatureNew (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 03/03/2023 at 13:56:51
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)

//FeatureNew_Publish annouces the endpoints available for this object
//FeatureNew_Publish - Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local
func FeatureNew_Publish(mux http.ServeMux) {
	// START
	// Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	mux.HandleFunc(dm.FeatureNew_Path, FeatureNew_Handler)
	//Cannot List via GUI
	mux.HandleFunc(dm.FeatureNew_PathView, FeatureNew_HandlerView)
	mux.HandleFunc(dm.FeatureNew_PathEdit, FeatureNew_HandlerEdit)
	mux.HandleFunc(dm.FeatureNew_PathNew, FeatureNew_HandlerNew)
	mux.HandleFunc(dm.FeatureNew_PathSave, FeatureNew_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.FeatureNew_Title)
    core.Catalog_Add(dm.FeatureNew_Title, dm.FeatureNew_Path, "", dm.FeatureNew_QueryString, "Application")
	// 
	// Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}



//FeatureNew_HandlerView is the handler used to View a page
//Allows Viewing for an existing FeatureNew record
//FeatureNew_HandlerView - Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func FeatureNew_HandlerView(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local
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

	searchID := core.GetURLparam(r, dm.FeatureNew_QueryString)
	_, rD, _ := dao.FeatureNew_GetByID(searchID)

	pageDetail := dm.FeatureNew_Page{
		Title:       CardTitle(dm.FeatureNew_Title, core.Action_View),
		PageTitle:   PageTitle(dm.FeatureNew_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = featurenew_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.FeatureNew_TemplateView, w, r, pageDetail)
	// 
	// Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//FeatureNew_HandlerEdit is the handler used generate the Edit page
//Allows Editing for an existing FeatureNew record and then allows the user to save the changes
//FeatureNew_HandlerEdit - Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func FeatureNew_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.FeatureNew_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.FeatureNew
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.FeatureNew)
	} else {
		_, rD, _ = dao.FeatureNew_GetByID(searchID)
	}

	
	pageDetail := dm.FeatureNew_Page{
		Title:       CardTitle(dm.FeatureNew_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.FeatureNew_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = featurenew_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.FeatureNew_TemplateEdit, w, r, pageDetail)
	// 
	// Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//FeatureNew_HandlerSave is the handler used process the saving of an FeatureNew
//It is called from the Edit and New pages
//FeatureNew_HandlerSave  - Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func FeatureNew_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("ID")
	logs.Servicing(r.URL.Path+itemID)

	item := featurenew_DataFromRequest(r)
	
	item, errStore := dao.FeatureNew_Store(item,r)
	if errStore == nil {	
		http.Redirect(w, r, dm.FeatureNew_Redirect, http.StatusFound)
	} else {
		logs.Information(dm.FeatureNew_Name, errStore.Error())
		//http.Redirect(w, r, r.Referer(), http.StatusFound)
		ExecuteRedirect(r.Referer(), w, r,dm.FeatureNew_QueryString,itemID,item)
	}
	// 
	// Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//FeatureNew_HandlerNew is the handler used process the creation of an FeatureNew
//It will create a new FeatureNew and then redirect to the Edit page
//FeatureNew_HandlerNew  - Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func FeatureNew_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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

	searchID := core.GetURLparam(r, dm.FeatureNew_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.FeatureNew
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.FeatureNew)
	} else {
		_, _, rD, _ = dao.FeatureNew_New()
	}



	pageDetail := dm.FeatureNew_Page{
		Title:       CardTitle(dm.FeatureNew_Title, core.Action_New),
		PageTitle:   PageTitle(dm.FeatureNew_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = featurenew_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.FeatureNew_TemplateNew, w, r, pageDetail)
	// 
	// Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}	



//featurenew_PopulatePage Builds/Populates the FeatureNew Page 
//featurenew_PopulatePage Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func featurenew_PopulatePage(rD dm.FeatureNew, pageDetail dm.FeatureNew_Page) dm.FeatureNew_Page {
	// Real DB Fields
	pageDetail.ID = rD.ID
	pageDetail.EstimationSessionID = rD.EstimationSessionID
	pageDetail.Name = rD.Name
	pageDetail.DeveloperEstimate = rD.DeveloperEstimate
	pageDetail.DeveloperResource = rD.DeveloperResource
	pageDetail.AnalystEstimate = rD.AnalystEstimate
	pageDetail.AnalystResource = rD.AnalystResource
	pageDetail.ConfidenceCODE = rD.ConfidenceCODE
	pageDetail.ProductManagerResource = rD.ProductManagerResource
	pageDetail.ProjectManagerResource = rD.ProjectManagerResource
	pageDetail.Description = rD.Description
	pageDetail.Comments = rD.Comments
	pageDetail.DevOpsID = rD.DevOpsID
	pageDetail.FreshDeskID = rD.FreshDeskID
	pageDetail.RSCID = rD.RSCID
	pageDetail.OtherID = rD.OtherID
	pageDetail.OtherID2 = rD.OtherID2
	pageDetail.DefaultProfile = rD.DefaultProfile
	pageDetail.ActualProfile = rD.ActualProfile
	// Add Pseudo/Extra Fields
	// Enrichment Fields 
	 
	pageDetail.EstimationSessionID_lookup = dao.EstimationSession_GetFilteredLookup("FeatureNew","EstimationSessionID")
	 
	pageDetail.DeveloperResource_lookup = dao.Resource_GetFilteredLookup("FeatureNew","DeveloperResource")
	 
	pageDetail.AnalystResource_lookup = dao.Resource_GetFilteredLookup("FeatureNew","AnalystResource")
	 
	pageDetail.ConfidenceCODE_lookup = dao.Confidence_GetFilteredLookup("FeatureNew","ConfidenceCODE")
	 
	pageDetail.ProductManagerResource_lookup = dao.Resource_GetFilteredLookup("FeatureNew","ProductManagerResource")
	 
	pageDetail.ProjectManagerResource_lookup = dao.Resource_GetFilteredLookup("FeatureNew","ProjectManagerResource")
	 
	pageDetail.DefaultProfile_lookup = dao.Profile_GetFilteredLookup("FeatureNew","DefaultProfile")
	 
	pageDetail.ActualProfile_lookup = dao.Profile_GetFilteredLookup("FeatureNew","ActualProfile")
	pageDetail.ID_props = rD.ID_props
	pageDetail.EstimationSessionID_props = rD.EstimationSessionID_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.DeveloperEstimate_props = rD.DeveloperEstimate_props
	pageDetail.DeveloperResource_props = rD.DeveloperResource_props
	pageDetail.AnalystEstimate_props = rD.AnalystEstimate_props
	pageDetail.AnalystResource_props = rD.AnalystResource_props
	pageDetail.ConfidenceCODE_props = rD.ConfidenceCODE_props
	pageDetail.ProductManagerResource_props = rD.ProductManagerResource_props
	pageDetail.ProjectManagerResource_props = rD.ProjectManagerResource_props
	pageDetail.Description_props = rD.Description_props
	pageDetail.Comments_props = rD.Comments_props
	pageDetail.DevOpsID_props = rD.DevOpsID_props
	pageDetail.FreshDeskID_props = rD.FreshDeskID_props
	pageDetail.RSCID_props = rD.RSCID_props
	pageDetail.OtherID_props = rD.OtherID_props
	pageDetail.OtherID2_props = rD.OtherID2_props
	pageDetail.DefaultProfile_props = rD.DefaultProfile_props
	pageDetail.ActualProfile_props = rD.ActualProfile_props
	return pageDetail
}
//featurenew_DataFromRequest is used process the content of an HTTP Request and return an instance of an FeatureNew
//featurenew_DataFromRequest Auto generated 03/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func featurenew_DataFromRequest(r *http.Request) dm.FeatureNew {

	var item dm.FeatureNew
		item.ID = r.FormValue(dm.FeatureNew_ID_scrn)
		item.EstimationSessionID = r.FormValue(dm.FeatureNew_EstimationSessionID_scrn)
		item.Name = r.FormValue(dm.FeatureNew_Name_scrn)
		item.DeveloperEstimate = r.FormValue(dm.FeatureNew_DeveloperEstimate_scrn)
		item.DeveloperResource = r.FormValue(dm.FeatureNew_DeveloperResource_scrn)
		item.AnalystEstimate = r.FormValue(dm.FeatureNew_AnalystEstimate_scrn)
		item.AnalystResource = r.FormValue(dm.FeatureNew_AnalystResource_scrn)
		item.ConfidenceCODE = r.FormValue(dm.FeatureNew_ConfidenceCODE_scrn)
		item.ProductManagerResource = r.FormValue(dm.FeatureNew_ProductManagerResource_scrn)
		item.ProjectManagerResource = r.FormValue(dm.FeatureNew_ProjectManagerResource_scrn)
		item.Description = r.FormValue(dm.FeatureNew_Description_scrn)
		item.Comments = r.FormValue(dm.FeatureNew_Comments_scrn)
		item.DevOpsID = r.FormValue(dm.FeatureNew_DevOpsID_scrn)
		item.FreshDeskID = r.FormValue(dm.FeatureNew_FreshDeskID_scrn)
		item.RSCID = r.FormValue(dm.FeatureNew_RSCID_scrn)
		item.OtherID = r.FormValue(dm.FeatureNew_OtherID_scrn)
		item.OtherID2 = r.FormValue(dm.FeatureNew_OtherID2_scrn)
		item.DefaultProfile = r.FormValue(dm.FeatureNew_DefaultProfile_scrn)
		item.ActualProfile = r.FormValue(dm.FeatureNew_ActualProfile_scrn)
	return item
}
