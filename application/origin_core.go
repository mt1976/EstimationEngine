package application
// ----------------------------------------------------------------
// Automatically generated  "/application/origin.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Origin (origin)
// Endpoint 	        : Origin (OriginID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/03/2023 at 19:24:49
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

//Origin_Publish annouces the endpoints available for this object
func Origin_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Origin_Path, Origin_Handler)
	mux.HandleFunc(dm.Origin_PathList, Origin_HandlerList)
	mux.HandleFunc(dm.Origin_PathView, Origin_HandlerView)
	mux.HandleFunc(dm.Origin_PathEdit, Origin_HandlerEdit)
	mux.HandleFunc(dm.Origin_PathNew, Origin_HandlerNew)
	mux.HandleFunc(dm.Origin_PathSave, Origin_HandlerSave)
	mux.HandleFunc(dm.Origin_PathDelete, Origin_HandlerDelete)
    core.API = core.API.AddRoute(dm.Origin_Title, dm.Origin_Path, "", dm.Origin_QueryString, "Application")
	logs.Publish("Application", dm.Origin_Title)
}

//Origin_HandlerList is the handler for the Origin list page
func Origin_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Origin

	objectName := dao.Translate("ObjectName", "Origin")
	reqField := "Base"
	usage := "Defines a filter for the list of Origin records." + core.TEXTAREA_CR
	usage = usage + "Fields can be any of those in the underlying DB table." + core.TEXTAREA_CR
	usage = usage + "Examples Below:" + core.TEXTAREA_CR
	usage = usage + "* datalength(_deleted) = 0 or " + core.TEXTAREA_CR 
	usage = usage + "* class IN ('x','y','z')"
	
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule,usage)
	if filter == "" {
		logs.Warning("No filter found : " + reqField + " for Object: " + objectName)
	} 

	noItems, returnList, _ := dao.Origin_GetListFiltered(filter)

	pageDetail := dm.Origin_PageList{
		Title:            CardTitle(dm.Origin_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Origin_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	nextTemplate :=  NextTemplate("Origin", "List", dm.Origin_TemplateList)
	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Origin_HandlerView is the handler used to View a Origin database record
func Origin_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Origin_QueryString)
	_, rD, _ := dao.Origin_GetByID(searchID)

	pageDetail := dm.Origin_Page{
		Title:       CardTitle(dm.Origin_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Origin_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = origin_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Origin", "View", dm.Origin_TemplateView)
	nextTemplate = origin_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Origin_HandlerEdit is the handler used to Edit of an existing Origin database record
func Origin_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Origin_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.Origin
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.Origin)
	} else {
		_, rD, _ = dao.Origin_GetByID(searchID)
	}
	
	pageDetail := dm.Origin_Page{
		Title:       CardTitle(dm.Origin_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Origin_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = origin_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Origin", "Edit", dm.Origin_TemplateEdit)
	nextTemplate = origin_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Origin_HandlerSave is the handler used process the saving of an Origin database record, either new or existing, referenced by Edit & New Handlers.
func Origin_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("OriginID")
	logs.Servicing(r.URL.Path+itemID)

	item := origin_DataFromRequest(r)
	
	item, errStore := dao.Origin_Store(item,r)
	if errStore == nil {
		nextTemplate :=  NextTemplate("Origin", "Save", dm.Origin_Redirect)
		nextTemplate = origin_URIQueryData(nextTemplate,item,itemID)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.Origin_Name, errStore.Error())
		ExecuteRedirect(r.Referer(), w, r,dm.Origin_QueryString,itemID,item)
	}
}

//Origin_HandlerNew is the handler used process the creation of an new Origin database record, then redirect to Edit
func Origin_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Origin_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.Origin
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.Origin)
	} else {
		_, _, rD, _ = dao.Origin_New()
	}

	pageDetail := dm.Origin_Page{
		Title:       CardTitle(dm.Origin_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Origin_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = origin_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Origin", "New", dm.Origin_TemplateNew)
	nextTemplate = origin_URIQueryData(nextTemplate,dm.Origin{},searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}	

//Origin_HandlerDelete is the handler used process the deletion of an Origin database record. May be Hard or SoftDelete.
func Origin_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Origin_QueryString)
	// DAO Call to Delete Origin Record, may be SoftDelete or HardDelete depending on the DAO implementation
	dao.Origin_Delete(searchID)	

	nextTemplate :=  NextTemplate("Origin", "Delete", dm.Origin_Redirect)
	nextTemplate = origin_URIQueryData(nextTemplate,dm.Origin{},searchID)
	http.Redirect(w, r, nextTemplate, http.StatusFound)
}
//origin_PopulatePage Builds/Populates the Origin Page from an instance of Origin from the Data Model
func origin_PopulatePage(rD dm.Origin, pageDetail dm.Origin_Page) dm.Origin_Page {
	// Real DB Fields
	pageDetail.SYSId = rD.SYSId
	pageDetail.OriginID = rD.OriginID
	pageDetail.StateID = rD.StateID
	pageDetail.DocTypeID = rD.DocTypeID
	pageDetail.Code = rD.Code
	pageDetail.FullName = rD.FullName
	pageDetail.Rate = rD.Rate
	pageDetail.Notes = rD.Notes
	pageDetail.StartDate = rD.StartDate
	pageDetail.EndDate = rD.EndDate
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.SYSDeleted = rD.SYSDeleted
	pageDetail.SYSDeletedBy = rD.SYSDeletedBy
	pageDetail.SYSDeletedHost = rD.SYSDeletedHost
	pageDetail.SYSActivity = rD.SYSActivity
	pageDetail.Currency = rD.Currency
	pageDetail.SYSDbVersion = rD.SYSDbVersion
	pageDetail.Comments = rD.Comments
	pageDetail.ProjectManager = rD.ProjectManager
	pageDetail.AccountManager = rD.AccountManager
	// Add Pseudo/Extra Fields, fields that are not in the DB but are used in the UI
	pageDetail.NoActiveProjects = rD.NoActiveProjects
	pageDetail.RateOnLoad = rD.RateOnLoad
	pageDetail.StatusOnLoad = rD.StatusOnLoad
	// Enrichment content, content used provide lookups,lists etc
	pageDetail.StateID_lookup = dao.OriginState_GetLookup()
	pageDetail.DocTypeID_lookup = dao.DocType_GetLookup()
	pageDetail.Currency_lookup = dao.StubLists_Get("ccy")
	pageDetail.ProjectManager_lookup = dao.Resource_GetFilteredLookup("Origin","ProjectManager")
	pageDetail.AccountManager_lookup = dao.Resource_GetFilteredLookup("Origin","AccountManager")
	// Add the Properties for the Fields
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.OriginID_props = rD.OriginID_props
	pageDetail.StateID_props = rD.StateID_props
	pageDetail.DocTypeID_props = rD.DocTypeID_props
	pageDetail.Code_props = rD.Code_props
	pageDetail.FullName_props = rD.FullName_props
	pageDetail.Rate_props = rD.Rate_props
	pageDetail.Notes_props = rD.Notes_props
	pageDetail.StartDate_props = rD.StartDate_props
	pageDetail.EndDate_props = rD.EndDate_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.SYSDeleted_props = rD.SYSDeleted_props
	pageDetail.SYSDeletedBy_props = rD.SYSDeletedBy_props
	pageDetail.SYSDeletedHost_props = rD.SYSDeletedHost_props
	pageDetail.SYSActivity_props = rD.SYSActivity_props
	pageDetail.Currency_props = rD.Currency_props
	pageDetail.SYSDbVersion_props = rD.SYSDbVersion_props
	pageDetail.Comments_props = rD.Comments_props
	pageDetail.ProjectManager_props = rD.ProjectManager_props
	pageDetail.AccountManager_props = rD.AccountManager_props
	pageDetail.NoActiveProjects_props = rD.NoActiveProjects_props
	pageDetail.RateOnLoad_props = rD.RateOnLoad_props
	pageDetail.StatusOnLoad_props = rD.StatusOnLoad_props
	return pageDetail
}
//origin_DataFromRequest is used process the content of an HTTP Request and return an instance of an Origin
func origin_DataFromRequest(r *http.Request) dm.Origin {
	var item dm.Origin
	item.SYSId = r.FormValue(dm.Origin_SYSId_scrn)
	item.OriginID = r.FormValue(dm.Origin_OriginID_scrn)
	item.StateID = r.FormValue(dm.Origin_StateID_scrn)
	item.DocTypeID = r.FormValue(dm.Origin_DocTypeID_scrn)
	item.Code = r.FormValue(dm.Origin_Code_scrn)
	item.FullName = r.FormValue(dm.Origin_FullName_scrn)
	item.Rate = r.FormValue(dm.Origin_Rate_scrn)
	item.Notes = r.FormValue(dm.Origin_Notes_scrn)
	item.StartDate = r.FormValue(dm.Origin_StartDate_scrn)
	item.EndDate = r.FormValue(dm.Origin_EndDate_scrn)
	item.SYSCreated = r.FormValue(dm.Origin_SYSCreated_scrn)
	item.SYSCreatedBy = r.FormValue(dm.Origin_SYSCreatedBy_scrn)
	item.SYSCreatedHost = r.FormValue(dm.Origin_SYSCreatedHost_scrn)
	item.SYSUpdated = r.FormValue(dm.Origin_SYSUpdated_scrn)
	item.SYSUpdatedBy = r.FormValue(dm.Origin_SYSUpdatedBy_scrn)
	item.SYSUpdatedHost = r.FormValue(dm.Origin_SYSUpdatedHost_scrn)
	item.SYSDeleted = r.FormValue(dm.Origin_SYSDeleted_scrn)
	item.SYSDeletedBy = r.FormValue(dm.Origin_SYSDeletedBy_scrn)
	item.SYSDeletedHost = r.FormValue(dm.Origin_SYSDeletedHost_scrn)
	item.SYSActivity = r.FormValue(dm.Origin_SYSActivity_scrn)
	item.Currency = r.FormValue(dm.Origin_Currency_scrn)
	item.SYSDbVersion = r.FormValue(dm.Origin_SYSDbVersion_scrn)
	item.Comments = r.FormValue(dm.Origin_Comments_scrn)
	item.ProjectManager = r.FormValue(dm.Origin_ProjectManager_scrn)
	item.AccountManager = r.FormValue(dm.Origin_AccountManager_scrn)
	item.NoActiveProjects = r.FormValue(dm.Origin_NoActiveProjects_scrn)
	item.RateOnLoad = r.FormValue(dm.Origin_RateOnLoad_scrn)
	item.StatusOnLoad = r.FormValue(dm.Origin_StatusOnLoad_scrn)
	return item
}
//origin_URIQueryData is used to replace the wildcards in the URI Query Path with the values from the Origin Data Model
func origin_URIQueryData(queryPath string,item dm.Origin,itemID string) string {
	if queryPath == "" {
		return ""
	}
	queryPath = core.ReplaceWildcard(queryPath, strings.ToUpper("ID"), itemID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Origin_OriginID_scrn), item.OriginID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Origin_StateID_scrn), item.StateID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Origin_DocTypeID_scrn), item.DocTypeID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Origin_Code_scrn), item.Code)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Origin_FullName_scrn), item.FullName)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Origin_Rate_scrn), item.Rate)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Origin_Notes_scrn), item.Notes)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Origin_StartDate_scrn), item.StartDate)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Origin_EndDate_scrn), item.EndDate)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Origin_Currency_scrn), item.Currency)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Origin_Comments_scrn), item.Comments)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Origin_ProjectManager_scrn), item.ProjectManager)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Origin_AccountManager_scrn), item.AccountManager)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Origin_NoActiveProjects_scrn), item.NoActiveProjects)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Origin_RateOnLoad_scrn), item.RateOnLoad)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Origin_StatusOnLoad_scrn), item.StatusOnLoad)
	return queryPath
}