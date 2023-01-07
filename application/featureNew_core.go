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
// Date & Time		    : 07/01/2023 at 23:01:29
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
//FeatureNew_Publish - Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
func FeatureNew_Publish(mux http.ServeMux) {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
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
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}



//FeatureNew_HandlerView is the handler used to View a page
//Allows Viewing for an existing FeatureNew record
//FeatureNew_HandlerView - Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
func FeatureNew_HandlerView(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
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
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//FeatureNew_HandlerEdit is the handler used generate the Edit page
//Allows Editing for an existing FeatureNew record and then allows the user to save the changes
//FeatureNew_HandlerEdit - Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
func FeatureNew_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
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
	_, rD, _ := dao.FeatureNew_GetByID(searchID)
	
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
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//FeatureNew_HandlerSave is the handler used process the saving of an FeatureNew
//It is called from the Edit and New pages
//FeatureNew_HandlerSave  - Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
func FeatureNew_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("ID"))

	item := featurenew_DataFromRequest(r)
	
	dao.FeatureNew_Store(item,r)	
	http.Redirect(w, r, dm.FeatureNew_Redirect, http.StatusFound)
	// 
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//FeatureNew_HandlerNew is the handler used process the creation of an FeatureNew
//It will create a new FeatureNew and then redirect to the Edit page
//FeatureNew_HandlerNew  - Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
func FeatureNew_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	_, _, rD, _ := dao.FeatureNew_New()

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
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}	



//featurenew_PopulatePage Builds/Populates the FeatureNew Page 
func featurenew_PopulatePage(rD dm.FeatureNew, pageDetail dm.FeatureNew_Page) dm.FeatureNew_Page {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.ID = rD.ID
	pageDetail.EstimationSession = rD.EstimationSession
	pageDetail.Name = rD.Name
	pageDetail.DevEstimate = rD.DevEstimate
	pageDetail.Confidence = rD.Confidence
	pageDetail.Developer = rD.Developer
	pageDetail.Comments = rD.Comments
	pageDetail.Description = rD.Description
	pageDetail.DevOpsID = rD.DevOpsID
	pageDetail.FreshDeskID = rD.FreshDeskID
	pageDetail.RSCID = rD.RSCID
	pageDetail.OtherID = rD.OtherID
	pageDetail.OtherID2 = rD.OtherID2
	pageDetail.Analyst = rD.Analyst
	pageDetail.ProductManager = rD.ProductManager
	pageDetail.ProjectManager = rD.ProjectManager
	pageDetail.DefaultProfile = rD.DefaultProfile
	pageDetail.ActualProfile = rD.ActualProfile
	
	
	//
	// Automatically generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	pageDetail.EstimationSession_lookup = dao.EstimationSession_GetLookup()
	
	
	
	
	
	
	
	pageDetail.Confidence_lookup = dao.Confidence_GetLookup()
	
	
	
	pageDetail.Developer_lookup = dao.Resource_GetLookup()
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.Analyst_lookup = dao.Resource_GetLookup()
	
	
	
	pageDetail.ProductManager_lookup = dao.Resource_GetLookup()
	
	
	
	pageDetail.ProjectManager_lookup = dao.Resource_GetLookup()
	
	
	
	pageDetail.DefaultProfile_lookup = dao.Profile_GetLookup()
	
	
	
	pageDetail.ActualProfile_lookup = dao.Profile_GetLookup()
	
	
	
	pageDetail.ID_props = rD.ID_props
	pageDetail.EstimationSession_props = rD.EstimationSession_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.DevEstimate_props = rD.DevEstimate_props
	pageDetail.Confidence_props = rD.Confidence_props
	pageDetail.Developer_props = rD.Developer_props
	pageDetail.Comments_props = rD.Comments_props
	pageDetail.Description_props = rD.Description_props
	pageDetail.DevOpsID_props = rD.DevOpsID_props
	pageDetail.FreshDeskID_props = rD.FreshDeskID_props
	pageDetail.RSCID_props = rD.RSCID_props
	pageDetail.OtherID_props = rD.OtherID_props
	pageDetail.OtherID2_props = rD.OtherID2_props
	pageDetail.Analyst_props = rD.Analyst_props
	pageDetail.ProductManager_props = rD.ProductManager_props
	pageDetail.ProjectManager_props = rD.ProjectManager_props
	pageDetail.DefaultProfile_props = rD.DefaultProfile_props
	pageDetail.ActualProfile_props = rD.ActualProfile_props
	
	// 
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	


//featurenew_DataFromRequest is used process the content of an HTTP Request and return an instance of an FeatureNew
func featurenew_DataFromRequest(r *http.Request) dm.FeatureNew {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	var item dm.FeatureNew
	// FIELD SET START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.ID = r.FormValue(dm.FeatureNew_ID_scrn)
		item.EstimationSession = r.FormValue(dm.FeatureNew_EstimationSession_scrn)
		item.Name = r.FormValue(dm.FeatureNew_Name_scrn)
		item.DevEstimate = r.FormValue(dm.FeatureNew_DevEstimate_scrn)
		item.Confidence = r.FormValue(dm.FeatureNew_Confidence_scrn)
		item.Developer = r.FormValue(dm.FeatureNew_Developer_scrn)
		item.Comments = r.FormValue(dm.FeatureNew_Comments_scrn)
		item.Description = r.FormValue(dm.FeatureNew_Description_scrn)
		item.DevOpsID = r.FormValue(dm.FeatureNew_DevOpsID_scrn)
		item.FreshDeskID = r.FormValue(dm.FeatureNew_FreshDeskID_scrn)
		item.RSCID = r.FormValue(dm.FeatureNew_RSCID_scrn)
		item.OtherID = r.FormValue(dm.FeatureNew_OtherID_scrn)
		item.OtherID2 = r.FormValue(dm.FeatureNew_OtherID2_scrn)
		item.Analyst = r.FormValue(dm.FeatureNew_Analyst_scrn)
		item.ProductManager = r.FormValue(dm.FeatureNew_ProductManager_scrn)
		item.ProjectManager = r.FormValue(dm.FeatureNew_ProjectManager_scrn)
		item.DefaultProfile = r.FormValue(dm.FeatureNew_DefaultProfile_scrn)
		item.ActualProfile = r.FormValue(dm.FeatureNew_ActualProfile_scrn)
	
	// 
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return item
}
