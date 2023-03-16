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
// Date & Time		    : 15/03/2023 at 19:24:48
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"
	"strings"

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
	mux.HandleFunc(dm.Feature_PathDelete, Feature_HandlerDelete)
    core.API = core.API.AddRoute(dm.Feature_Title, dm.Feature_Path, "", dm.Feature_QueryString, "Application")
	logs.Publish("Application", dm.Feature_Title)
}

//Feature_HandlerList is the handler for the Feature list page
func Feature_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
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
	usage := "Defines a filter for the list of Feature records." + core.TEXTAREA_CR
	usage = usage + "Fields can be any of those in the underlying DB table." + core.TEXTAREA_CR
	usage = usage + "Examples Below:" + core.TEXTAREA_CR
	usage = usage + "* datalength(_deleted) = 0 or " + core.TEXTAREA_CR 
	usage = usage + "* class IN ('x','y','z')"
	
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule,usage)
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
}

//Feature_HandlerView is the handler used to View a Feature database record
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

	nextTemplate :=  NextTemplate("Feature", "View", dm.Feature_TemplateView)
	nextTemplate = feature_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Feature_HandlerEdit is the handler used to Edit of an existing Feature database record
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
	nextTemplate = feature_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Feature_HandlerSave is the handler used process the saving of an Feature database record, either new or existing, referenced by Edit & New Handlers.
func Feature_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
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
		nextTemplate = feature_URIQueryData(nextTemplate,item,itemID)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.Feature_Name, errStore.Error())
		ExecuteRedirect(r.Referer(), w, r,dm.Feature_QueryString,itemID,item)
	}
}

//Feature_HandlerNew is the handler used process the creation of an new Feature database record, then redirect to Edit
func Feature_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
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
	nextTemplate = feature_URIQueryData(nextTemplate,dm.Feature{},searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}	

//Feature_HandlerDelete is the handler used process the deletion of an Feature database record. May be Hard or SoftDelete.
func Feature_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Feature_QueryString)
	// DAO Call to Delete Feature Record, may be SoftDelete or HardDelete depending on the DAO implementation
	dao.Feature_Delete(searchID)	

	nextTemplate :=  NextTemplate("Feature", "Delete", dm.Feature_Redirect)
	nextTemplate = feature_URIQueryData(nextTemplate,dm.Feature{},searchID)
	http.Redirect(w, r, nextTemplate, http.StatusFound)
}
//feature_PopulatePage Builds/Populates the Feature Page from an instance of Feature from the Data Model
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
	pageDetail.RequirementsPerc = rD.RequirementsPerc
	pageDetail.AnalystTestPerc = rD.AnalystTestPerc
	pageDetail.DocumentationPerc = rD.DocumentationPerc
	pageDetail.ManagementPerc = rD.ManagementPerc
	pageDetail.UatSupportPerc = rD.UatSupportPerc
	pageDetail.MarketingPerc = rD.MarketingPerc
	pageDetail.ContingencyPerc = rD.ContingencyPerc
	pageDetail.TrainingPerc = rD.TrainingPerc
	pageDetail.RoundedTo = rD.RoundedTo
	// Add Pseudo/Extra Fields, fields that are not in the DB but are used in the UI
	pageDetail.OriginName = rD.OriginName
	pageDetail.ProjectName = rD.ProjectName
	pageDetail.OriginID = rD.OriginID
	pageDetail.ProjectID = rD.ProjectID
	pageDetail.ProfileSelectedOld = rD.ProfileSelectedOld
	pageDetail.CCY = rD.CCY
	pageDetail.OriginCode = rD.OriginCode
	pageDetail.EstimationSessionName = rD.EstimationSessionName
	// Enrichment content, content used provide lookups,lists etc
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
	// Add the Properties for the Fields
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
	pageDetail.RequirementsPerc_props = rD.RequirementsPerc_props
	pageDetail.AnalystTestPerc_props = rD.AnalystTestPerc_props
	pageDetail.DocumentationPerc_props = rD.DocumentationPerc_props
	pageDetail.ManagementPerc_props = rD.ManagementPerc_props
	pageDetail.UatSupportPerc_props = rD.UatSupportPerc_props
	pageDetail.MarketingPerc_props = rD.MarketingPerc_props
	pageDetail.ContingencyPerc_props = rD.ContingencyPerc_props
	pageDetail.TrainingPerc_props = rD.TrainingPerc_props
	pageDetail.RoundedTo_props = rD.RoundedTo_props
	pageDetail.OriginName_props = rD.OriginName_props
	pageDetail.ProjectName_props = rD.ProjectName_props
	pageDetail.OriginID_props = rD.OriginID_props
	pageDetail.ProjectID_props = rD.ProjectID_props
	pageDetail.ProfileSelectedOld_props = rD.ProfileSelectedOld_props
	pageDetail.CCY_props = rD.CCY_props
	pageDetail.OriginCode_props = rD.OriginCode_props
	pageDetail.EstimationSessionName_props = rD.EstimationSessionName_props
	return pageDetail
}
//feature_DataFromRequest is used process the content of an HTTP Request and return an instance of an Feature
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
	item.RequirementsPerc = r.FormValue(dm.Feature_RequirementsPerc_scrn)
	item.AnalystTestPerc = r.FormValue(dm.Feature_AnalystTestPerc_scrn)
	item.DocumentationPerc = r.FormValue(dm.Feature_DocumentationPerc_scrn)
	item.ManagementPerc = r.FormValue(dm.Feature_ManagementPerc_scrn)
	item.UatSupportPerc = r.FormValue(dm.Feature_UatSupportPerc_scrn)
	item.MarketingPerc = r.FormValue(dm.Feature_MarketingPerc_scrn)
	item.ContingencyPerc = r.FormValue(dm.Feature_ContingencyPerc_scrn)
	item.TrainingPerc = r.FormValue(dm.Feature_TrainingPerc_scrn)
	item.RoundedTo = r.FormValue(dm.Feature_RoundedTo_scrn)
	item.OriginName = r.FormValue(dm.Feature_OriginName_scrn)
	item.ProjectName = r.FormValue(dm.Feature_ProjectName_scrn)
	item.OriginID = r.FormValue(dm.Feature_OriginID_scrn)
	item.ProjectID = r.FormValue(dm.Feature_ProjectID_scrn)
	item.ProfileSelectedOld = r.FormValue(dm.Feature_ProfileSelectedOld_scrn)
	item.CCY = r.FormValue(dm.Feature_CCY_scrn)
	item.OriginCode = r.FormValue(dm.Feature_OriginCode_scrn)
	item.EstimationSessionName = r.FormValue(dm.Feature_EstimationSessionName_scrn)
	return item
}
//feature_URIQueryData is used to replace the wildcards in the URI Query Path with the values from the Feature Data Model
func feature_URIQueryData(queryPath string,item dm.Feature,itemID string) string {
	if queryPath == "" {
		return ""
	}
	queryPath = core.ReplaceWildcard(queryPath, strings.ToUpper("ID"), itemID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_FeatureID_scrn), item.FeatureID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_EstimationSessionID_scrn), item.EstimationSessionID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_ConfidenceCODE_scrn), item.ConfidenceCODE)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_Name_scrn), item.Name)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_DevelopmentEstimate_scrn), item.DevelopmentEstimate)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_DevelopmentFactored_scrn), item.DevelopmentFactored)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_Requirements_scrn), item.Requirements)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_AnalystTest_scrn), item.AnalystTest)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_Documentation_scrn), item.Documentation)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_Management_scrn), item.Management)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_UatSupport_scrn), item.UatSupport)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_Marketing_scrn), item.Marketing)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_Contingency_scrn), item.Contingency)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_TrackerID_scrn), item.TrackerID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_AdoID_scrn), item.AdoID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_FreshdeskID_scrn), item.FreshdeskID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_ExtRef_scrn), item.ExtRef)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_ExtRef2_scrn), item.ExtRef2)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_DeveloperResource_scrn), item.DeveloperResource)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_ApproverResource_scrn), item.ApproverResource)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_Activity_scrn), item.Activity)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_OffProfile_scrn), item.OffProfile)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_OffProfileJustification_scrn), item.OffProfileJustification)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_RequirementsDefault_scrn), item.RequirementsDefault)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_AnalystTestDefault_scrn), item.AnalystTestDefault)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_DocumentationDefault_scrn), item.DocumentationDefault)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_ManagementDefault_scrn), item.ManagementDefault)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_UatSupportDefault_scrn), item.UatSupportDefault)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_MarketingDefault_scrn), item.MarketingDefault)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_ContingencyDefault_scrn), item.ContingencyDefault)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_DevelopmentFactoredDefault_scrn), item.DevelopmentFactoredDefault)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_Total_scrn), item.Total)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_Comments_scrn), item.Comments)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_Description_scrn), item.Description)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_AnalystResource_scrn), item.AnalystResource)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_ProductManagerResource_scrn), item.ProductManagerResource)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_ProjectManagerResource_scrn), item.ProjectManagerResource)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_Training_scrn), item.Training)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_TrainingDefault_scrn), item.TrainingDefault)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_ProfileDefault_scrn), item.ProfileDefault)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_ProfileSelected_scrn), item.ProfileSelected)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_EstimateEffort_scrn), item.EstimateEffort)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_EstimateEffortDefault_scrn), item.EstimateEffortDefault)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_AnalystEstimate_scrn), item.AnalystEstimate)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_TotalDefault_scrn), item.TotalDefault)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_RequirementsPerc_scrn), item.RequirementsPerc)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_AnalystTestPerc_scrn), item.AnalystTestPerc)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_DocumentationPerc_scrn), item.DocumentationPerc)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_ManagementPerc_scrn), item.ManagementPerc)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_UatSupportPerc_scrn), item.UatSupportPerc)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_MarketingPerc_scrn), item.MarketingPerc)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_ContingencyPerc_scrn), item.ContingencyPerc)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_TrainingPerc_scrn), item.TrainingPerc)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_RoundedTo_scrn), item.RoundedTo)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_OriginName_scrn), item.OriginName)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_ProjectName_scrn), item.ProjectName)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_OriginID_scrn), item.OriginID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_ProjectID_scrn), item.ProjectID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_ProfileSelectedOld_scrn), item.ProfileSelectedOld)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_CCY_scrn), item.CCY)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_OriginCode_scrn), item.OriginCode)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Feature_EstimationSessionName_scrn), item.EstimationSessionName)
	return queryPath
}