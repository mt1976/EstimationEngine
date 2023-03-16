package application
// ----------------------------------------------------------------
// Automatically generated  "/application/estimationsession.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : EstimationSession (estimationsession)
// Endpoint 	        : EstimationSession (EstimationSessionID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/03/2023 at 19:24:47
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

//EstimationSession_Publish annouces the endpoints available for this object
func EstimationSession_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.EstimationSession_Path, EstimationSession_Handler)
	mux.HandleFunc(dm.EstimationSession_PathList, EstimationSession_HandlerList)
	mux.HandleFunc(dm.EstimationSession_PathView, EstimationSession_HandlerView)
	mux.HandleFunc(dm.EstimationSession_PathEdit, EstimationSession_HandlerEdit)
	mux.HandleFunc(dm.EstimationSession_PathNew, EstimationSession_HandlerNew)
	mux.HandleFunc(dm.EstimationSession_PathSave, EstimationSession_HandlerSave)
	mux.HandleFunc(dm.EstimationSession_PathDelete, EstimationSession_HandlerDelete)
    core.API = core.API.AddRoute(dm.EstimationSession_Title, dm.EstimationSession_Path, "", dm.EstimationSession_QueryString, "Application")
	logs.Publish("Application", dm.EstimationSession_Title)
}

//EstimationSession_HandlerList is the handler for the EstimationSession list page
func EstimationSession_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.EstimationSession

	objectName := dao.Translate("ObjectName", "EstimationSession")
	reqField := "Base"
	usage := "Defines a filter for the list of EstimationSession records." + core.TEXTAREA_CR
	usage = usage + "Fields can be any of those in the underlying DB table." + core.TEXTAREA_CR
	usage = usage + "Examples Below:" + core.TEXTAREA_CR
	usage = usage + "* datalength(_deleted) = 0 or " + core.TEXTAREA_CR 
	usage = usage + "* class IN ('x','y','z')"
	
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule,usage)
	if filter == "" {
		logs.Warning("No filter found : " + reqField + " for Object: " + objectName)
	} 

	noItems, returnList, _ := dao.EstimationSession_GetListFiltered(filter)

	pageDetail := dm.EstimationSession_PageList{
		Title:            CardTitle(dm.EstimationSession_Title, core.Action_List),
		PageTitle:        PageTitle(dm.EstimationSession_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	nextTemplate :=  NextTemplate("EstimationSession", "List", dm.EstimationSession_TemplateList)
	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//EstimationSession_HandlerView is the handler used to View a EstimationSession database record
func EstimationSession_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.EstimationSession_QueryString)
	_, rD, _ := dao.EstimationSession_GetByID(searchID)

	pageDetail := dm.EstimationSession_Page{
		Title:       CardTitle(dm.EstimationSession_Title, core.Action_View),
		PageTitle:   PageTitle(dm.EstimationSession_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = estimationsession_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("EstimationSession", "View", dm.EstimationSession_TemplateView)
	nextTemplate = estimationsession_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//EstimationSession_HandlerEdit is the handler used to Edit of an existing EstimationSession database record
func EstimationSession_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.EstimationSession_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.EstimationSession
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.EstimationSession)
	} else {
		_, rD, _ = dao.EstimationSession_GetByID(searchID)
	}
	
	pageDetail := dm.EstimationSession_Page{
		Title:       CardTitle(dm.EstimationSession_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.EstimationSession_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = estimationsession_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("EstimationSession", "Edit", dm.EstimationSession_TemplateEdit)
	nextTemplate = estimationsession_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//EstimationSession_HandlerSave is the handler used process the saving of an EstimationSession database record, either new or existing, referenced by Edit & New Handlers.
func EstimationSession_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("EstimationSessionID")
	logs.Servicing(r.URL.Path+itemID)

	item := estimationsession_DataFromRequest(r)
	
	item, errStore := dao.EstimationSession_Store(item,r)
	if errStore == nil {
		nextTemplate :=  NextTemplate("EstimationSession", "Save", dm.EstimationSession_Redirect)
		nextTemplate = estimationsession_URIQueryData(nextTemplate,item,itemID)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.EstimationSession_Name, errStore.Error())
		ExecuteRedirect(r.Referer(), w, r,dm.EstimationSession_QueryString,itemID,item)
	}
}

//EstimationSession_HandlerNew is the handler used process the creation of an new EstimationSession database record, then redirect to Edit
func EstimationSession_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.EstimationSession_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.EstimationSession
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.EstimationSession)
	} else {
		_, _, rD, _ = dao.EstimationSession_New()
	}

	pageDetail := dm.EstimationSession_Page{
		Title:       CardTitle(dm.EstimationSession_Title, core.Action_New),
		PageTitle:   PageTitle(dm.EstimationSession_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = estimationsession_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("EstimationSession", "New", dm.EstimationSession_TemplateNew)
	nextTemplate = estimationsession_URIQueryData(nextTemplate,dm.EstimationSession{},searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}	

//EstimationSession_HandlerDelete is the handler used process the deletion of an EstimationSession database record. May be Hard or SoftDelete.
func EstimationSession_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.EstimationSession_QueryString)
	// DAO Call to Delete EstimationSession Record, may be SoftDelete or HardDelete depending on the DAO implementation
	dao.EstimationSession_Delete(searchID)	

	nextTemplate :=  NextTemplate("EstimationSession", "Delete", dm.EstimationSession_Redirect)
	nextTemplate = estimationsession_URIQueryData(nextTemplate,dm.EstimationSession{},searchID)
	http.Redirect(w, r, nextTemplate, http.StatusFound)
}
//estimationsession_PopulatePage Builds/Populates the EstimationSession Page from an instance of EstimationSession from the Data Model
func estimationsession_PopulatePage(rD dm.EstimationSession, pageDetail dm.EstimationSession_Page) dm.EstimationSession_Page {
	// Real DB Fields
	pageDetail.SYSId = rD.SYSId
	pageDetail.EstimationSessionID = rD.EstimationSessionID
	pageDetail.ProjectID = rD.ProjectID
	pageDetail.EstimationStateID = rD.EstimationStateID
	pageDetail.Activity = rD.Activity
	pageDetail.Releases = rD.Releases
	pageDetail.Total = rD.Total
	pageDetail.Contingency = rD.Contingency
	pageDetail.ReqDays = rD.ReqDays
	pageDetail.RegCost = rD.RegCost
	pageDetail.ImpDays = rD.ImpDays
	pageDetail.ImpCost = rD.ImpCost
	pageDetail.UatDays = rD.UatDays
	pageDetail.UatCost = rD.UatCost
	pageDetail.MgtDays = rD.MgtDays
	pageDetail.MgtCost = rD.MgtCost
	pageDetail.RelDays = rD.RelDays
	pageDetail.RelCost = rD.RelCost
	pageDetail.SupportUplift = rD.SupportUplift
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.SYSDeleted = rD.SYSDeleted
	pageDetail.SYSDeletedBy = rD.SYSDeletedBy
	pageDetail.SYSDeletedHost = rD.SYSDeletedHost
	pageDetail.Name = rD.Name
	pageDetail.AdoID = rD.AdoID
	pageDetail.FreshdeskID = rD.FreshdeskID
	pageDetail.TrackerID = rD.TrackerID
	pageDetail.EstRef = rD.EstRef
	pageDetail.ExtRef = rD.ExtRef
	pageDetail.SYSActivity = rD.SYSActivity
	pageDetail.SYSDbVersion = rD.SYSDbVersion
	pageDetail.Comments = rD.Comments
	pageDetail.ProjectManager = rD.ProjectManager
	pageDetail.ProductManager = rD.ProductManager
	pageDetail.Approver = rD.Approver
	pageDetail.IssueDate = rD.IssueDate
	pageDetail.ExpiryDate = rD.ExpiryDate
	// Add Pseudo/Extra Fields, fields that are not in the DB but are used in the UI
	pageDetail.Origin = rD.Origin
	pageDetail.OriginStateID = rD.OriginStateID
	pageDetail.OriginState = rD.OriginState
	pageDetail.OriginDocTypeID = rD.OriginDocTypeID
	pageDetail.OriginDocType = rD.OriginDocType
	pageDetail.OriginCode = rD.OriginCode
	pageDetail.OriginName = rD.OriginName
	pageDetail.OriginRate = rD.OriginRate
	pageDetail.ProjectProfileID = rD.ProjectProfileID
	pageDetail.ProjectProfile = rD.ProjectProfile
	pageDetail.ProjectDefaultReleases = rD.ProjectDefaultReleases
	pageDetail.ProjectDefaultReleaseHours = rD.ProjectDefaultReleaseHours
	pageDetail.ProjectBlendedRate = rD.ProjectBlendedRate
	pageDetail.ProjectStateID = rD.ProjectStateID
	pageDetail.ProjectState = rD.ProjectState
	pageDetail.ProjectName = rD.ProjectName
	pageDetail.ProjectStartDate = rD.ProjectStartDate
	pageDetail.ProjectEndDate = rD.ProjectEndDate
	pageDetail.ProfileSupportUpliftPerc = rD.ProfileSupportUpliftPerc
	pageDetail.CCY = rD.CCY
	pageDetail.CCYCode = rD.CCYCode
	pageDetail.EffortTotal = rD.EffortTotal
	pageDetail.FreshDeskURI = rD.FreshDeskURI
	pageDetail.ADOURI = rD.ADOURI
	pageDetail.NoActiveFeatures = rD.NoActiveFeatures
	pageDetail.OriginKey = rD.OriginKey
	// Enrichment content, content used provide lookups,lists etc
	pageDetail.ProjectID_lookup = dao.Project_GetLookup()
	pageDetail.EstimationStateID_lookup = dao.EstimationState_GetLookup()
	pageDetail.ProjectManager_lookup = dao.Resource_GetFilteredLookup("EstimationSession","ProjectManager")
	pageDetail.ProductManager_lookup = dao.Resource_GetFilteredLookup("EstimationSession","ProductManager")
	pageDetail.Approver_lookup = dao.Resource_GetFilteredLookup("EstimationSession","Approver")
	// Add the Properties for the Fields
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.EstimationSessionID_props = rD.EstimationSessionID_props
	pageDetail.ProjectID_props = rD.ProjectID_props
	pageDetail.EstimationStateID_props = rD.EstimationStateID_props
	pageDetail.Activity_props = rD.Activity_props
	pageDetail.Releases_props = rD.Releases_props
	pageDetail.Total_props = rD.Total_props
	pageDetail.Contingency_props = rD.Contingency_props
	pageDetail.ReqDays_props = rD.ReqDays_props
	pageDetail.RegCost_props = rD.RegCost_props
	pageDetail.ImpDays_props = rD.ImpDays_props
	pageDetail.ImpCost_props = rD.ImpCost_props
	pageDetail.UatDays_props = rD.UatDays_props
	pageDetail.UatCost_props = rD.UatCost_props
	pageDetail.MgtDays_props = rD.MgtDays_props
	pageDetail.MgtCost_props = rD.MgtCost_props
	pageDetail.RelDays_props = rD.RelDays_props
	pageDetail.RelCost_props = rD.RelCost_props
	pageDetail.SupportUplift_props = rD.SupportUplift_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.SYSDeleted_props = rD.SYSDeleted_props
	pageDetail.SYSDeletedBy_props = rD.SYSDeletedBy_props
	pageDetail.SYSDeletedHost_props = rD.SYSDeletedHost_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.AdoID_props = rD.AdoID_props
	pageDetail.FreshdeskID_props = rD.FreshdeskID_props
	pageDetail.TrackerID_props = rD.TrackerID_props
	pageDetail.EstRef_props = rD.EstRef_props
	pageDetail.ExtRef_props = rD.ExtRef_props
	pageDetail.SYSActivity_props = rD.SYSActivity_props
	pageDetail.SYSDbVersion_props = rD.SYSDbVersion_props
	pageDetail.Comments_props = rD.Comments_props
	pageDetail.ProjectManager_props = rD.ProjectManager_props
	pageDetail.ProductManager_props = rD.ProductManager_props
	pageDetail.Approver_props = rD.Approver_props
	pageDetail.IssueDate_props = rD.IssueDate_props
	pageDetail.ExpiryDate_props = rD.ExpiryDate_props
	pageDetail.Origin_props = rD.Origin_props
	pageDetail.OriginStateID_props = rD.OriginStateID_props
	pageDetail.OriginState_props = rD.OriginState_props
	pageDetail.OriginDocTypeID_props = rD.OriginDocTypeID_props
	pageDetail.OriginDocType_props = rD.OriginDocType_props
	pageDetail.OriginCode_props = rD.OriginCode_props
	pageDetail.OriginName_props = rD.OriginName_props
	pageDetail.OriginRate_props = rD.OriginRate_props
	pageDetail.ProjectProfileID_props = rD.ProjectProfileID_props
	pageDetail.ProjectProfile_props = rD.ProjectProfile_props
	pageDetail.ProjectDefaultReleases_props = rD.ProjectDefaultReleases_props
	pageDetail.ProjectDefaultReleaseHours_props = rD.ProjectDefaultReleaseHours_props
	pageDetail.ProjectBlendedRate_props = rD.ProjectBlendedRate_props
	pageDetail.ProjectStateID_props = rD.ProjectStateID_props
	pageDetail.ProjectState_props = rD.ProjectState_props
	pageDetail.ProjectName_props = rD.ProjectName_props
	pageDetail.ProjectStartDate_props = rD.ProjectStartDate_props
	pageDetail.ProjectEndDate_props = rD.ProjectEndDate_props
	pageDetail.ProfileSupportUpliftPerc_props = rD.ProfileSupportUpliftPerc_props
	pageDetail.CCY_props = rD.CCY_props
	pageDetail.CCYCode_props = rD.CCYCode_props
	pageDetail.EffortTotal_props = rD.EffortTotal_props
	pageDetail.FreshDeskURI_props = rD.FreshDeskURI_props
	pageDetail.ADOURI_props = rD.ADOURI_props
	pageDetail.NoActiveFeatures_props = rD.NoActiveFeatures_props
	pageDetail.OriginKey_props = rD.OriginKey_props
	return pageDetail
}
//estimationsession_DataFromRequest is used process the content of an HTTP Request and return an instance of an EstimationSession
func estimationsession_DataFromRequest(r *http.Request) dm.EstimationSession {
	var item dm.EstimationSession
	item.SYSId = r.FormValue(dm.EstimationSession_SYSId_scrn)
	item.EstimationSessionID = r.FormValue(dm.EstimationSession_EstimationSessionID_scrn)
	item.ProjectID = r.FormValue(dm.EstimationSession_ProjectID_scrn)
	item.EstimationStateID = r.FormValue(dm.EstimationSession_EstimationStateID_scrn)
	item.Activity = r.FormValue(dm.EstimationSession_Activity_scrn)
	item.Releases = r.FormValue(dm.EstimationSession_Releases_scrn)
	item.Total = r.FormValue(dm.EstimationSession_Total_scrn)
	item.Contingency = r.FormValue(dm.EstimationSession_Contingency_scrn)
	item.ReqDays = r.FormValue(dm.EstimationSession_ReqDays_scrn)
	item.RegCost = r.FormValue(dm.EstimationSession_RegCost_scrn)
	item.ImpDays = r.FormValue(dm.EstimationSession_ImpDays_scrn)
	item.ImpCost = r.FormValue(dm.EstimationSession_ImpCost_scrn)
	item.UatDays = r.FormValue(dm.EstimationSession_UatDays_scrn)
	item.UatCost = r.FormValue(dm.EstimationSession_UatCost_scrn)
	item.MgtDays = r.FormValue(dm.EstimationSession_MgtDays_scrn)
	item.MgtCost = r.FormValue(dm.EstimationSession_MgtCost_scrn)
	item.RelDays = r.FormValue(dm.EstimationSession_RelDays_scrn)
	item.RelCost = r.FormValue(dm.EstimationSession_RelCost_scrn)
	item.SupportUplift = r.FormValue(dm.EstimationSession_SupportUplift_scrn)
	item.SYSCreated = r.FormValue(dm.EstimationSession_SYSCreated_scrn)
	item.SYSCreatedBy = r.FormValue(dm.EstimationSession_SYSCreatedBy_scrn)
	item.SYSCreatedHost = r.FormValue(dm.EstimationSession_SYSCreatedHost_scrn)
	item.SYSUpdated = r.FormValue(dm.EstimationSession_SYSUpdated_scrn)
	item.SYSUpdatedBy = r.FormValue(dm.EstimationSession_SYSUpdatedBy_scrn)
	item.SYSUpdatedHost = r.FormValue(dm.EstimationSession_SYSUpdatedHost_scrn)
	item.SYSDeleted = r.FormValue(dm.EstimationSession_SYSDeleted_scrn)
	item.SYSDeletedBy = r.FormValue(dm.EstimationSession_SYSDeletedBy_scrn)
	item.SYSDeletedHost = r.FormValue(dm.EstimationSession_SYSDeletedHost_scrn)
	item.Name = r.FormValue(dm.EstimationSession_Name_scrn)
	item.AdoID = r.FormValue(dm.EstimationSession_AdoID_scrn)
	item.FreshdeskID = r.FormValue(dm.EstimationSession_FreshdeskID_scrn)
	item.TrackerID = r.FormValue(dm.EstimationSession_TrackerID_scrn)
	item.EstRef = r.FormValue(dm.EstimationSession_EstRef_scrn)
	item.ExtRef = r.FormValue(dm.EstimationSession_ExtRef_scrn)
	item.SYSActivity = r.FormValue(dm.EstimationSession_SYSActivity_scrn)
	item.SYSDbVersion = r.FormValue(dm.EstimationSession_SYSDbVersion_scrn)
	item.Comments = r.FormValue(dm.EstimationSession_Comments_scrn)
	item.ProjectManager = r.FormValue(dm.EstimationSession_ProjectManager_scrn)
	item.ProductManager = r.FormValue(dm.EstimationSession_ProductManager_scrn)
	item.Approver = r.FormValue(dm.EstimationSession_Approver_scrn)
	item.IssueDate = r.FormValue(dm.EstimationSession_IssueDate_scrn)
	item.ExpiryDate = r.FormValue(dm.EstimationSession_ExpiryDate_scrn)
	item.Origin = r.FormValue(dm.EstimationSession_Origin_scrn)
	item.OriginStateID = r.FormValue(dm.EstimationSession_OriginStateID_scrn)
	item.OriginState = r.FormValue(dm.EstimationSession_OriginState_scrn)
	item.OriginDocTypeID = r.FormValue(dm.EstimationSession_OriginDocTypeID_scrn)
	item.OriginDocType = r.FormValue(dm.EstimationSession_OriginDocType_scrn)
	item.OriginCode = r.FormValue(dm.EstimationSession_OriginCode_scrn)
	item.OriginName = r.FormValue(dm.EstimationSession_OriginName_scrn)
	item.OriginRate = r.FormValue(dm.EstimationSession_OriginRate_scrn)
	item.ProjectProfileID = r.FormValue(dm.EstimationSession_ProjectProfileID_scrn)
	item.ProjectProfile = r.FormValue(dm.EstimationSession_ProjectProfile_scrn)
	item.ProjectDefaultReleases = r.FormValue(dm.EstimationSession_ProjectDefaultReleases_scrn)
	item.ProjectDefaultReleaseHours = r.FormValue(dm.EstimationSession_ProjectDefaultReleaseHours_scrn)
	item.ProjectBlendedRate = r.FormValue(dm.EstimationSession_ProjectBlendedRate_scrn)
	item.ProjectStateID = r.FormValue(dm.EstimationSession_ProjectStateID_scrn)
	item.ProjectState = r.FormValue(dm.EstimationSession_ProjectState_scrn)
	item.ProjectName = r.FormValue(dm.EstimationSession_ProjectName_scrn)
	item.ProjectStartDate = r.FormValue(dm.EstimationSession_ProjectStartDate_scrn)
	item.ProjectEndDate = r.FormValue(dm.EstimationSession_ProjectEndDate_scrn)
	item.ProfileSupportUpliftPerc = r.FormValue(dm.EstimationSession_ProfileSupportUpliftPerc_scrn)
	item.CCY = r.FormValue(dm.EstimationSession_CCY_scrn)
	item.CCYCode = r.FormValue(dm.EstimationSession_CCYCode_scrn)
	item.EffortTotal = r.FormValue(dm.EstimationSession_EffortTotal_scrn)
	item.FreshDeskURI = r.FormValue(dm.EstimationSession_FreshDeskURI_scrn)
	item.ADOURI = r.FormValue(dm.EstimationSession_ADOURI_scrn)
	item.NoActiveFeatures = r.FormValue(dm.EstimationSession_NoActiveFeatures_scrn)
	item.OriginKey = r.FormValue(dm.EstimationSession_OriginKey_scrn)
	return item
}
//estimationsession_URIQueryData is used to replace the wildcards in the URI Query Path with the values from the EstimationSession Data Model
func estimationsession_URIQueryData(queryPath string,item dm.EstimationSession,itemID string) string {
	if queryPath == "" {
		return ""
	}
	queryPath = core.ReplaceWildcard(queryPath, strings.ToUpper("ID"), itemID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_EstimationSessionID_scrn), item.EstimationSessionID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_ProjectID_scrn), item.ProjectID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_EstimationStateID_scrn), item.EstimationStateID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_Activity_scrn), item.Activity)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_Releases_scrn), item.Releases)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_Total_scrn), item.Total)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_Contingency_scrn), item.Contingency)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_ReqDays_scrn), item.ReqDays)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_RegCost_scrn), item.RegCost)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_ImpDays_scrn), item.ImpDays)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_ImpCost_scrn), item.ImpCost)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_UatDays_scrn), item.UatDays)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_UatCost_scrn), item.UatCost)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_MgtDays_scrn), item.MgtDays)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_MgtCost_scrn), item.MgtCost)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_RelDays_scrn), item.RelDays)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_RelCost_scrn), item.RelCost)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_SupportUplift_scrn), item.SupportUplift)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_Name_scrn), item.Name)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_AdoID_scrn), item.AdoID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_FreshdeskID_scrn), item.FreshdeskID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_TrackerID_scrn), item.TrackerID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_EstRef_scrn), item.EstRef)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_ExtRef_scrn), item.ExtRef)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_Comments_scrn), item.Comments)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_ProjectManager_scrn), item.ProjectManager)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_ProductManager_scrn), item.ProductManager)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_Approver_scrn), item.Approver)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_IssueDate_scrn), item.IssueDate)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_ExpiryDate_scrn), item.ExpiryDate)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_Origin_scrn), item.Origin)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_OriginStateID_scrn), item.OriginStateID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_OriginState_scrn), item.OriginState)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_OriginDocTypeID_scrn), item.OriginDocTypeID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_OriginDocType_scrn), item.OriginDocType)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_OriginCode_scrn), item.OriginCode)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_OriginName_scrn), item.OriginName)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_OriginRate_scrn), item.OriginRate)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_ProjectProfileID_scrn), item.ProjectProfileID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_ProjectProfile_scrn), item.ProjectProfile)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_ProjectDefaultReleases_scrn), item.ProjectDefaultReleases)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_ProjectDefaultReleaseHours_scrn), item.ProjectDefaultReleaseHours)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_ProjectBlendedRate_scrn), item.ProjectBlendedRate)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_ProjectStateID_scrn), item.ProjectStateID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_ProjectState_scrn), item.ProjectState)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_ProjectName_scrn), item.ProjectName)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_ProjectStartDate_scrn), item.ProjectStartDate)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_ProjectEndDate_scrn), item.ProjectEndDate)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_ProfileSupportUpliftPerc_scrn), item.ProfileSupportUpliftPerc)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_CCY_scrn), item.CCY)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_CCYCode_scrn), item.CCYCode)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_EffortTotal_scrn), item.EffortTotal)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_FreshDeskURI_scrn), item.FreshDeskURI)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_ADOURI_scrn), item.ADOURI)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_NoActiveFeatures_scrn), item.NoActiveFeatures)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.EstimationSession_OriginKey_scrn), item.OriginKey)
	return queryPath
}