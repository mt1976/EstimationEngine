package application
// ----------------------------------------------------------------
// Automatically generated  "/application/feature.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Feature (feature)
// Endpoint 	        : Feature (FeatureID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 07/03/2023 at 16:42:06
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
//Feature_Publish - Auto generated 07/03/2023 by matttownsend (Matt Townsend) on silicon.local
func Feature_Publish(mux http.ServeMux) {
	// START
	// Auto generated 07/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	mux.HandleFunc(dm.Feature_Path, Feature_Handler)
	mux.HandleFunc(dm.Feature_PathList, Feature_HandlerList)
	mux.HandleFunc(dm.Feature_PathView, Feature_HandlerView)
	mux.HandleFunc(dm.Feature_PathEdit, Feature_HandlerEdit)
	mux.HandleFunc(dm.Feature_PathNew, Feature_HandlerNew)
	mux.HandleFunc(dm.Feature_PathSave, Feature_HandlerSave)
	mux.HandleFunc(dm.Feature_PathDelete, Feature_HandlerDelete)
	logs.Publish("Application", dm.Feature_Title)
    core.Catalog_Add(dm.Feature_Title, dm.Feature_Path, "", dm.Feature_QueryString, "Application")
	// 
	// Auto generated 07/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//Feature_HandlerList is the handler for the list page
//Allows Listing of Feature records
//Feature_HandlerList - Auto generated 07/03/2023 by matttownsend (Matt Townsend) on silicon.local
func Feature_HandlerList(w http.ResponseWriter, r *http.Request) {
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

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Feature

	objectName := dao.Translate("ObjectName", "Feature")
	reqField := "Base"
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule)
	if filter == "" {
		logs.Warning("No filter found : " + reqField + " for Object: " + objectName)
	} 

	noItems, returnList, _ := dao.Feature_GetListFiltered(filter)

	pageDetail := dm.Feature_PageList{
		Title:            CardTitle(dm.Feature_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Feature_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	nextTemplate :=  NextTemplate("Feature", "List", dm.Feature_TemplateList)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 07/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END

}


//Feature_HandlerView is the handler used to View a page
//Allows Viewing for an existing Feature record
//Feature_HandlerView - Auto generated 07/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func Feature_HandlerView(w http.ResponseWriter, r *http.Request) {
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

	nextTemplate :=  NextTemplate("Feature", "View", dm.Feature_TemplateView)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 07/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//Feature_HandlerEdit is the handler used generate the Edit page
//Allows Editing for an existing Feature record and then allows the user to save the changes
//Feature_HandlerEdit - Auto generated 07/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func Feature_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 07/03/2023 by matttownsend (Matt Townsend) on silicon.local
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

	searchID := core.GetURLparam(r, dm.Feature_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.Feature
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.Feature)
	} else {
		_, rD, _ = dao.Feature_GetByID(searchID)
	}

	
	pageDetail := dm.Feature_Page{
		Title:       CardTitle(dm.Feature_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Feature_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = feature_PopulatePage(rD , pageDetail) 


	nextTemplate :=  NextTemplate("Feature", "Edit", dm.Feature_TemplateEdit)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 07/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//Feature_HandlerSave is the handler used process the saving of an Feature
//It is called from the Edit and New pages
//Feature_HandlerSave  - Auto generated 07/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func Feature_HandlerSave(w http.ResponseWriter, r *http.Request) {
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
	logs.Servicing(r.URL.Path+itemID)

	item := feature_DataFromRequest(r)
	
	item, errStore := dao.Feature_Store(item,r)
	if errStore == nil {
		nextTemplate :=  NextTemplate("Feature", "Save", dm.Feature_Redirect)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.Feature_Name, errStore.Error())
		//http.Redirect(w, r, r.Referer(), http.StatusFound)
		ExecuteRedirect(r.Referer(), w, r,dm.Feature_QueryString,itemID,item)
	}
	// 
	// Auto generated 07/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//Feature_HandlerNew is the handler used process the creation of an Feature
//It will create a new Feature and then redirect to the Edit page
//Feature_HandlerNew  - Auto generated 07/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func Feature_HandlerNew(w http.ResponseWriter, r *http.Request) {
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
		Title:       CardTitle(dm.Feature_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Feature_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = feature_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Feature", "New", dm.Feature_TemplateNew)
	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 07/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}	


//Feature_HandlerDelete is the handler used process the deletion of an Feature
// It will delete the Feature and then redirect to the List page
//Feature_HandlerDelete - Auto generated 07/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func Feature_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 07/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	//
	// Code Continues Below
	//
	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Feature_QueryString)

	dao.Feature_Delete(searchID)	

	nextTemplate :=  NextTemplate("Feature", "Delete", dm.Feature_Redirect)
	http.Redirect(w, r, nextTemplate, http.StatusFound)
	// 
	// Auto generated 07/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//feature_PopulatePage Builds/Populates the Feature Page 
//feature_PopulatePage Auto generated 07/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func feature_PopulatePage(rD dm.Feature, pageDetail dm.Feature_Page) dm.Feature_Page {
	// Real DB Fields
	pageDetail.SYSId = rD.SYSId
	pageDetail.FeatureID = rD.FeatureID
	pageDetail.EstimationSessionID = rD.EstimationSessionID
	pageDetail.ConfidenceCODE = rD.ConfidenceCODE
	pageDetail.Name = rD.Name
	pageDetail.DevelopmentEstimate = rD.DevelopmentEstimate
	pageDetail.DevelopmentFactored = rD.DevelopmentFactored
	pageDetail.Requirements = rD.Requirements
	pageDetail.AnalystTest = rD.AnalystTest
	pageDetail.Documentation = rD.Documentation
	pageDetail.Management = rD.Management
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
	pageDetail.DeveloperResource = rD.DeveloperResource
	pageDetail.ApproverResource = rD.ApproverResource
	pageDetail.Activity = rD.Activity
	pageDetail.OffProfile = rD.OffProfile
	pageDetail.OffProfileJustification = rD.OffProfileJustification
	pageDetail.SYSActivity = rD.SYSActivity
	pageDetail.RequirementsDefault = rD.RequirementsDefault
	pageDetail.AnalystTestDefault = rD.AnalystTestDefault
	pageDetail.DocumentationDefault = rD.DocumentationDefault
	pageDetail.ManagementDefault = rD.ManagementDefault
	pageDetail.UatSupportDefault = rD.UatSupportDefault
	pageDetail.MarketingDefault = rD.MarketingDefault
	pageDetail.ContingencyDefault = rD.ContingencyDefault
	pageDetail.DevelopmentFactoredDefault = rD.DevelopmentFactoredDefault
	pageDetail.Total = rD.Total
	pageDetail.SYSDbVersion = rD.SYSDbVersion
	pageDetail.Comments = rD.Comments
	pageDetail.Description = rD.Description
	pageDetail.AnalystResource = rD.AnalystResource
	pageDetail.ProductManagerResource = rD.ProductManagerResource
	pageDetail.ProjectManagerResource = rD.ProjectManagerResource
	pageDetail.Training = rD.Training
	pageDetail.TrainingDefault = rD.TrainingDefault
	pageDetail.ProfileDefault = rD.ProfileDefault
	pageDetail.ProfileSelected = rD.ProfileSelected
	pageDetail.EstimateEffort = rD.EstimateEffort
	pageDetail.EstimateEffortDefault = rD.EstimateEffortDefault
	pageDetail.AnalystEstimate = rD.AnalystEstimate
	pageDetail.TotalDefault = rD.TotalDefault
	// Add Pseudo/Extra Fields
	pageDetail.OriginName = rD.OriginName
	pageDetail.ProjectName = rD.ProjectName
	pageDetail.OriginID = rD.OriginID
	pageDetail.ProjectID = rD.ProjectID
	pageDetail.ProfileSelectedOld = rD.ProfileSelectedOld
	pageDetail.CCY = rD.CCY
	pageDetail.RoundTo = rD.RoundTo
	pageDetail.OriginCode = rD.OriginCode
	pageDetail.EstimationSessionName = rD.EstimationSessionName
	// Enrichment Fields 
	 
	pageDetail.EstimationSessionID_lookup = dao.EstimationSession_GetLookup()
	 
	pageDetail.ConfidenceCODE_lookup = dao.Confidence_GetFilteredLookup("Feature","ConfidenceCODE")
	 
	pageDetail.DeveloperResource_lookup = dao.Resource_GetFilteredLookup("Feature","DeveloperResource")
	 
	pageDetail.ApproverResource_lookup = dao.Resource_GetFilteredLookup("Feature","ApproverResource")
	
	pageDetail.OffProfile_lookup = dao.StubLists_Get("tf")
	 
	pageDetail.AnalystResource_lookup = dao.Resource_GetFilteredLookup("Feature","AnalystResource")
	 
	pageDetail.ProductManagerResource_lookup = dao.Resource_GetFilteredLookup("Feature","ProductManagerResource")
	 
	pageDetail.ProjectManagerResource_lookup = dao.Resource_GetFilteredLookup("Feature","ProjectManagerResource")
	 
	pageDetail.ProfileDefault_lookup = dao.Profile_GetLookup()
	 
	pageDetail.ProfileSelected_lookup = dao.Profile_GetFilteredLookup("Feature","ProfileSelected")
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.FeatureID_props = rD.FeatureID_props
	pageDetail.EstimationSessionID_props = rD.EstimationSessionID_props
	pageDetail.ConfidenceCODE_props = rD.ConfidenceCODE_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.DevelopmentEstimate_props = rD.DevelopmentEstimate_props
	pageDetail.DevelopmentFactored_props = rD.DevelopmentFactored_props
	pageDetail.Requirements_props = rD.Requirements_props
	pageDetail.AnalystTest_props = rD.AnalystTest_props
	pageDetail.Documentation_props = rD.Documentation_props
	pageDetail.Management_props = rD.Management_props
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
	pageDetail.DeveloperResource_props = rD.DeveloperResource_props
	pageDetail.ApproverResource_props = rD.ApproverResource_props
	pageDetail.Activity_props = rD.Activity_props
	pageDetail.OffProfile_props = rD.OffProfile_props
	pageDetail.OffProfileJustification_props = rD.OffProfileJustification_props
	pageDetail.SYSActivity_props = rD.SYSActivity_props
	pageDetail.RequirementsDefault_props = rD.RequirementsDefault_props
	pageDetail.AnalystTestDefault_props = rD.AnalystTestDefault_props
	pageDetail.DocumentationDefault_props = rD.DocumentationDefault_props
	pageDetail.ManagementDefault_props = rD.ManagementDefault_props
	pageDetail.UatSupportDefault_props = rD.UatSupportDefault_props
	pageDetail.MarketingDefault_props = rD.MarketingDefault_props
	pageDetail.ContingencyDefault_props = rD.ContingencyDefault_props
	pageDetail.DevelopmentFactoredDefault_props = rD.DevelopmentFactoredDefault_props
	pageDetail.Total_props = rD.Total_props
	pageDetail.SYSDbVersion_props = rD.SYSDbVersion_props
	pageDetail.Comments_props = rD.Comments_props
	pageDetail.Description_props = rD.Description_props
	pageDetail.AnalystResource_props = rD.AnalystResource_props
	pageDetail.ProductManagerResource_props = rD.ProductManagerResource_props
	pageDetail.ProjectManagerResource_props = rD.ProjectManagerResource_props
	pageDetail.Training_props = rD.Training_props
	pageDetail.TrainingDefault_props = rD.TrainingDefault_props
	pageDetail.ProfileDefault_props = rD.ProfileDefault_props
	pageDetail.ProfileSelected_props = rD.ProfileSelected_props
	pageDetail.EstimateEffort_props = rD.EstimateEffort_props
	pageDetail.EstimateEffortDefault_props = rD.EstimateEffortDefault_props
	pageDetail.AnalystEstimate_props = rD.AnalystEstimate_props
	pageDetail.TotalDefault_props = rD.TotalDefault_props
	pageDetail.OriginName_props = rD.OriginName_props
	pageDetail.ProjectName_props = rD.ProjectName_props
	pageDetail.OriginID_props = rD.OriginID_props
	pageDetail.ProjectID_props = rD.ProjectID_props
	pageDetail.ProfileSelectedOld_props = rD.ProfileSelectedOld_props
	pageDetail.CCY_props = rD.CCY_props
	pageDetail.RoundTo_props = rD.RoundTo_props
	pageDetail.OriginCode_props = rD.OriginCode_props
	pageDetail.EstimationSessionName_props = rD.EstimationSessionName_props
	return pageDetail
}
//feature_DataFromRequest is used process the content of an HTTP Request and return an instance of an Feature
//feature_DataFromRequest Auto generated 07/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func feature_DataFromRequest(r *http.Request) dm.Feature {

	var item dm.Feature
		item.SYSId = r.FormValue(dm.Feature_SYSId_scrn)
		item.FeatureID = r.FormValue(dm.Feature_FeatureID_scrn)
		item.EstimationSessionID = r.FormValue(dm.Feature_EstimationSessionID_scrn)
		item.ConfidenceCODE = r.FormValue(dm.Feature_ConfidenceCODE_scrn)
		item.Name = r.FormValue(dm.Feature_Name_scrn)
		item.DevelopmentEstimate = r.FormValue(dm.Feature_DevelopmentEstimate_scrn)
		item.DevelopmentFactored = r.FormValue(dm.Feature_DevelopmentFactored_scrn)
		item.Requirements = r.FormValue(dm.Feature_Requirements_scrn)
		item.AnalystTest = r.FormValue(dm.Feature_AnalystTest_scrn)
		item.Documentation = r.FormValue(dm.Feature_Documentation_scrn)
		item.Management = r.FormValue(dm.Feature_Management_scrn)
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
		item.DeveloperResource = r.FormValue(dm.Feature_DeveloperResource_scrn)
		item.ApproverResource = r.FormValue(dm.Feature_ApproverResource_scrn)
		item.Activity = r.FormValue(dm.Feature_Activity_scrn)
		item.OffProfile = r.FormValue(dm.Feature_OffProfile_scrn)
		item.OffProfileJustification = r.FormValue(dm.Feature_OffProfileJustification_scrn)
		item.SYSActivity = r.FormValue(dm.Feature_SYSActivity_scrn)
		item.RequirementsDefault = r.FormValue(dm.Feature_RequirementsDefault_scrn)
		item.AnalystTestDefault = r.FormValue(dm.Feature_AnalystTestDefault_scrn)
		item.DocumentationDefault = r.FormValue(dm.Feature_DocumentationDefault_scrn)
		item.ManagementDefault = r.FormValue(dm.Feature_ManagementDefault_scrn)
		item.UatSupportDefault = r.FormValue(dm.Feature_UatSupportDefault_scrn)
		item.MarketingDefault = r.FormValue(dm.Feature_MarketingDefault_scrn)
		item.ContingencyDefault = r.FormValue(dm.Feature_ContingencyDefault_scrn)
		item.DevelopmentFactoredDefault = r.FormValue(dm.Feature_DevelopmentFactoredDefault_scrn)
		item.Total = r.FormValue(dm.Feature_Total_scrn)
		item.SYSDbVersion = r.FormValue(dm.Feature_SYSDbVersion_scrn)
		item.Comments = r.FormValue(dm.Feature_Comments_scrn)
		item.Description = r.FormValue(dm.Feature_Description_scrn)
		item.AnalystResource = r.FormValue(dm.Feature_AnalystResource_scrn)
		item.ProductManagerResource = r.FormValue(dm.Feature_ProductManagerResource_scrn)
		item.ProjectManagerResource = r.FormValue(dm.Feature_ProjectManagerResource_scrn)
		item.Training = r.FormValue(dm.Feature_Training_scrn)
		item.TrainingDefault = r.FormValue(dm.Feature_TrainingDefault_scrn)
		item.ProfileDefault = r.FormValue(dm.Feature_ProfileDefault_scrn)
		item.ProfileSelected = r.FormValue(dm.Feature_ProfileSelected_scrn)
		item.EstimateEffort = r.FormValue(dm.Feature_EstimateEffort_scrn)
		item.EstimateEffortDefault = r.FormValue(dm.Feature_EstimateEffortDefault_scrn)
		item.AnalystEstimate = r.FormValue(dm.Feature_AnalystEstimate_scrn)
		item.TotalDefault = r.FormValue(dm.Feature_TotalDefault_scrn)
		item.OriginName = r.FormValue(dm.Feature_OriginName_scrn)
		item.ProjectName = r.FormValue(dm.Feature_ProjectName_scrn)
		item.OriginID = r.FormValue(dm.Feature_OriginID_scrn)
		item.ProjectID = r.FormValue(dm.Feature_ProjectID_scrn)
		item.ProfileSelectedOld = r.FormValue(dm.Feature_ProfileSelectedOld_scrn)
		item.CCY = r.FormValue(dm.Feature_CCY_scrn)
		item.RoundTo = r.FormValue(dm.Feature_RoundTo_scrn)
		item.OriginCode = r.FormValue(dm.Feature_OriginCode_scrn)
		item.EstimationSessionName = r.FormValue(dm.Feature_EstimationSessionName_scrn)
	return item
}
