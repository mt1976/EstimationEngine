package application
// ----------------------------------------------------------------
// Automatically generated  "/application/originstate.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : OriginState (originstate)
// Endpoint 	        : OriginState (OriginStateID)
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

//OriginState_Publish annouces the endpoints available for this object
func OriginState_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.OriginState_Path, OriginState_Handler)
	mux.HandleFunc(dm.OriginState_PathList, OriginState_HandlerList)
	mux.HandleFunc(dm.OriginState_PathView, OriginState_HandlerView)
	mux.HandleFunc(dm.OriginState_PathEdit, OriginState_HandlerEdit)
	mux.HandleFunc(dm.OriginState_PathNew, OriginState_HandlerNew)
	mux.HandleFunc(dm.OriginState_PathSave, OriginState_HandlerSave)
	mux.HandleFunc(dm.OriginState_PathDelete, OriginState_HandlerDelete)
    core.API = core.API.AddRoute(dm.OriginState_Title, dm.OriginState_Path, "", dm.OriginState_QueryString, "Application")
	logs.Publish("Application", dm.OriginState_Title)
}

//OriginState_HandlerList is the handler for the OriginState list page
func OriginState_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.OriginState

	objectName := dao.Translate("ObjectName", "OriginState")
	reqField := "Base"
	usage := "Defines a filter for the list of OriginState records." + core.TEXTAREA_CR
	usage = usage + "Fields can be any of those in the underlying DB table." + core.TEXTAREA_CR
	usage = usage + "Examples Below:" + core.TEXTAREA_CR
	usage = usage + "* datalength(_deleted) = 0 or " + core.TEXTAREA_CR 
	usage = usage + "* class IN ('x','y','z')"
	
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule,usage)
	if filter == "" {
		logs.Warning("No filter found : " + reqField + " for Object: " + objectName)
	} 

	noItems, returnList, _ := dao.OriginState_GetListFiltered(filter)

	pageDetail := dm.OriginState_PageList{
		Title:            CardTitle(dm.OriginState_Title, core.Action_List),
		PageTitle:        PageTitle(dm.OriginState_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	nextTemplate :=  NextTemplate("OriginState", "List", dm.OriginState_TemplateList)
	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//OriginState_HandlerView is the handler used to View a OriginState database record
func OriginState_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.OriginState_QueryString)
	_, rD, _ := dao.OriginState_GetByID(searchID)

	pageDetail := dm.OriginState_Page{
		Title:       CardTitle(dm.OriginState_Title, core.Action_View),
		PageTitle:   PageTitle(dm.OriginState_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = originstate_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("OriginState", "View", dm.OriginState_TemplateView)
	nextTemplate = originstate_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//OriginState_HandlerEdit is the handler used to Edit of an existing OriginState database record
func OriginState_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.OriginState_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.OriginState
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.OriginState)
	} else {
		_, rD, _ = dao.OriginState_GetByID(searchID)
	}
	
	pageDetail := dm.OriginState_Page{
		Title:       CardTitle(dm.OriginState_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.OriginState_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = originstate_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("OriginState", "Edit", dm.OriginState_TemplateEdit)
	nextTemplate = originstate_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//OriginState_HandlerSave is the handler used process the saving of an OriginState database record, either new or existing, referenced by Edit & New Handlers.
func OriginState_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("OriginStateID")
	logs.Servicing(r.URL.Path+itemID)

	item := originstate_DataFromRequest(r)
	
	item, errStore := dao.OriginState_Store(item,r)
	if errStore == nil {
		nextTemplate :=  NextTemplate("OriginState", "Save", dm.OriginState_Redirect)
		nextTemplate = originstate_URIQueryData(nextTemplate,item,itemID)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.OriginState_Name, errStore.Error())
		ExecuteRedirect(r.Referer(), w, r,dm.OriginState_QueryString,itemID,item)
	}
}

//OriginState_HandlerNew is the handler used process the creation of an new OriginState database record, then redirect to Edit
func OriginState_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.OriginState_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.OriginState
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.OriginState)
	} else {
		_, _, rD, _ = dao.OriginState_New()
	}

	pageDetail := dm.OriginState_Page{
		Title:       CardTitle(dm.OriginState_Title, core.Action_New),
		PageTitle:   PageTitle(dm.OriginState_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = originstate_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("OriginState", "New", dm.OriginState_TemplateNew)
	nextTemplate = originstate_URIQueryData(nextTemplate,dm.OriginState{},searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}	

//OriginState_HandlerDelete is the handler used process the deletion of an OriginState database record. May be Hard or SoftDelete.
func OriginState_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.OriginState_QueryString)
	// DAO Call to Delete OriginState Record, may be SoftDelete or HardDelete depending on the DAO implementation
	dao.OriginState_Delete(searchID)	

	nextTemplate :=  NextTemplate("OriginState", "Delete", dm.OriginState_Redirect)
	nextTemplate = originstate_URIQueryData(nextTemplate,dm.OriginState{},searchID)
	http.Redirect(w, r, nextTemplate, http.StatusFound)
}
//originstate_PopulatePage Builds/Populates the OriginState Page from an instance of OriginState from the Data Model
func originstate_PopulatePage(rD dm.OriginState, pageDetail dm.OriginState_Page) dm.OriginState_Page {
	// Real DB Fields
	pageDetail.SYSId = rD.SYSId
	pageDetail.OriginStateID = rD.OriginStateID
	pageDetail.Code = rD.Code
	pageDetail.Name = rD.Name
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.SYSDeleted = rD.SYSDeleted
	pageDetail.SYSDeletedBy = rD.SYSDeletedBy
	pageDetail.SYSDeletedHost = rD.SYSDeletedHost
	pageDetail.SYSDbVersion = rD.SYSDbVersion
	pageDetail.IsLocked = rD.IsLocked
	pageDetail.Notify = rD.Notify
	pageDetail.Migrate = rD.Migrate
	// Add Pseudo/Extra Fields, fields that are not in the DB but are used in the UI
	// Enrichment content, content used provide lookups,lists etc
	pageDetail.Migrate_lookup = dao.StubLists_Get("tf")
	// Add the Properties for the Fields
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.OriginStateID_props = rD.OriginStateID_props
	pageDetail.Code_props = rD.Code_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.SYSDeleted_props = rD.SYSDeleted_props
	pageDetail.SYSDeletedBy_props = rD.SYSDeletedBy_props
	pageDetail.SYSDeletedHost_props = rD.SYSDeletedHost_props
	pageDetail.SYSDbVersion_props = rD.SYSDbVersion_props
	pageDetail.IsLocked_props = rD.IsLocked_props
	pageDetail.Notify_props = rD.Notify_props
	pageDetail.Migrate_props = rD.Migrate_props
	return pageDetail
}
//originstate_DataFromRequest is used process the content of an HTTP Request and return an instance of an OriginState
func originstate_DataFromRequest(r *http.Request) dm.OriginState {
	var item dm.OriginState
	item.SYSId = r.FormValue(dm.OriginState_SYSId_scrn)
	item.OriginStateID = r.FormValue(dm.OriginState_OriginStateID_scrn)
	item.Code = r.FormValue(dm.OriginState_Code_scrn)
	item.Name = r.FormValue(dm.OriginState_Name_scrn)
	item.SYSCreated = r.FormValue(dm.OriginState_SYSCreated_scrn)
	item.SYSCreatedBy = r.FormValue(dm.OriginState_SYSCreatedBy_scrn)
	item.SYSCreatedHost = r.FormValue(dm.OriginState_SYSCreatedHost_scrn)
	item.SYSUpdated = r.FormValue(dm.OriginState_SYSUpdated_scrn)
	item.SYSUpdatedBy = r.FormValue(dm.OriginState_SYSUpdatedBy_scrn)
	item.SYSUpdatedHost = r.FormValue(dm.OriginState_SYSUpdatedHost_scrn)
	item.SYSDeleted = r.FormValue(dm.OriginState_SYSDeleted_scrn)
	item.SYSDeletedBy = r.FormValue(dm.OriginState_SYSDeletedBy_scrn)
	item.SYSDeletedHost = r.FormValue(dm.OriginState_SYSDeletedHost_scrn)
	item.SYSDbVersion = r.FormValue(dm.OriginState_SYSDbVersion_scrn)
	item.IsLocked = r.FormValue(dm.OriginState_IsLocked_scrn)
	item.Notify = r.FormValue(dm.OriginState_Notify_scrn)
	item.Migrate = r.FormValue(dm.OriginState_Migrate_scrn)
	return item
}
//originstate_URIQueryData is used to replace the wildcards in the URI Query Path with the values from the OriginState Data Model
func originstate_URIQueryData(queryPath string,item dm.OriginState,itemID string) string {
	if queryPath == "" {
		return ""
	}
	queryPath = core.ReplaceWildcard(queryPath, strings.ToUpper("ID"), itemID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.OriginState_OriginStateID_scrn), item.OriginStateID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.OriginState_Code_scrn), item.Code)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.OriginState_Name_scrn), item.Name)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.OriginState_IsLocked_scrn), item.IsLocked)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.OriginState_Notify_scrn), item.Notify)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.OriginState_Migrate_scrn), item.Migrate)
	return queryPath
}