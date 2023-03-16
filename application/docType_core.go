package application
// ----------------------------------------------------------------
// Automatically generated  "/application/doctype.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DocType (doctype)
// Endpoint 	        : DocType (DocTypeID)
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

//DocType_Publish annouces the endpoints available for this object
func DocType_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.DocType_Path, DocType_Handler)
	mux.HandleFunc(dm.DocType_PathList, DocType_HandlerList)
	mux.HandleFunc(dm.DocType_PathView, DocType_HandlerView)
	mux.HandleFunc(dm.DocType_PathEdit, DocType_HandlerEdit)
	mux.HandleFunc(dm.DocType_PathNew, DocType_HandlerNew)
	mux.HandleFunc(dm.DocType_PathSave, DocType_HandlerSave)
	mux.HandleFunc(dm.DocType_PathDelete, DocType_HandlerDelete)
    core.API = core.API.AddRoute(dm.DocType_Title, dm.DocType_Path, "", dm.DocType_QueryString, "Application")
	logs.Publish("Application", dm.DocType_Title)
}

//DocType_HandlerList is the handler for the DocType list page
func DocType_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.DocType

	objectName := dao.Translate("ObjectName", "DocType")
	reqField := "Base"
	usage := "Defines a filter for the list of DocType records." + core.TEXTAREA_CR
	usage = usage + "Fields can be any of those in the underlying DB table." + core.TEXTAREA_CR
	usage = usage + "Examples Below:" + core.TEXTAREA_CR
	usage = usage + "* datalength(_deleted) = 0 or " + core.TEXTAREA_CR 
	usage = usage + "* class IN ('x','y','z')"
	
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule,usage)
	if filter == "" {
		logs.Warning("No filter found : " + reqField + " for Object: " + objectName)
	} 

	noItems, returnList, _ := dao.DocType_GetListFiltered(filter)

	pageDetail := dm.DocType_PageList{
		Title:            CardTitle(dm.DocType_Title, core.Action_List),
		PageTitle:        PageTitle(dm.DocType_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	nextTemplate :=  NextTemplate("DocType", "List", dm.DocType_TemplateList)
	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//DocType_HandlerView is the handler used to View a DocType database record
func DocType_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DocType_QueryString)
	_, rD, _ := dao.DocType_GetByID(searchID)

	pageDetail := dm.DocType_Page{
		Title:       CardTitle(dm.DocType_Title, core.Action_View),
		PageTitle:   PageTitle(dm.DocType_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = doctype_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("DocType", "View", dm.DocType_TemplateView)
	nextTemplate = doctype_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//DocType_HandlerEdit is the handler used to Edit of an existing DocType database record
func DocType_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DocType_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.DocType
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.DocType)
	} else {
		_, rD, _ = dao.DocType_GetByID(searchID)
	}
	
	pageDetail := dm.DocType_Page{
		Title:       CardTitle(dm.DocType_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.DocType_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = doctype_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("DocType", "Edit", dm.DocType_TemplateEdit)
	nextTemplate = doctype_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//DocType_HandlerSave is the handler used process the saving of an DocType database record, either new or existing, referenced by Edit & New Handlers.
func DocType_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("DocTypeID")
	logs.Servicing(r.URL.Path+itemID)

	item := doctype_DataFromRequest(r)
	
	item, errStore := dao.DocType_Store(item,r)
	if errStore == nil {
		nextTemplate :=  NextTemplate("DocType", "Save", dm.DocType_Redirect)
		nextTemplate = doctype_URIQueryData(nextTemplate,item,itemID)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.DocType_Name, errStore.Error())
		ExecuteRedirect(r.Referer(), w, r,dm.DocType_QueryString,itemID,item)
	}
}

//DocType_HandlerNew is the handler used process the creation of an new DocType database record, then redirect to Edit
func DocType_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DocType_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.DocType
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.DocType)
	} else {
		_, _, rD, _ = dao.DocType_New()
	}

	pageDetail := dm.DocType_Page{
		Title:       CardTitle(dm.DocType_Title, core.Action_New),
		PageTitle:   PageTitle(dm.DocType_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = doctype_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("DocType", "New", dm.DocType_TemplateNew)
	nextTemplate = doctype_URIQueryData(nextTemplate,dm.DocType{},searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}	

//DocType_HandlerDelete is the handler used process the deletion of an DocType database record. May be Hard or SoftDelete.
func DocType_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.DocType_QueryString)
	// DAO Call to Delete DocType Record, may be SoftDelete or HardDelete depending on the DAO implementation
	dao.DocType_Delete(searchID)	

	nextTemplate :=  NextTemplate("DocType", "Delete", dm.DocType_Redirect)
	nextTemplate = doctype_URIQueryData(nextTemplate,dm.DocType{},searchID)
	http.Redirect(w, r, nextTemplate, http.StatusFound)
}
//doctype_PopulatePage Builds/Populates the DocType Page from an instance of DocType from the Data Model
func doctype_PopulatePage(rD dm.DocType, pageDetail dm.DocType_Page) dm.DocType_Page {
	// Real DB Fields
	pageDetail.SYSId = rD.SYSId
	pageDetail.DocTypeID = rD.DocTypeID
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
	pageDetail.Comments = rD.Comments
	pageDetail.Migrate = rD.Migrate
	// Add Pseudo/Extra Fields, fields that are not in the DB but are used in the UI
	// Enrichment content, content used provide lookups,lists etc
	pageDetail.Migrate_lookup = dao.StubLists_Get("tf")
	// Add the Properties for the Fields
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.DocTypeID_props = rD.DocTypeID_props
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
	pageDetail.Comments_props = rD.Comments_props
	pageDetail.Migrate_props = rD.Migrate_props
	return pageDetail
}
//doctype_DataFromRequest is used process the content of an HTTP Request and return an instance of an DocType
func doctype_DataFromRequest(r *http.Request) dm.DocType {
	var item dm.DocType
	item.SYSId = r.FormValue(dm.DocType_SYSId_scrn)
	item.DocTypeID = r.FormValue(dm.DocType_DocTypeID_scrn)
	item.Code = r.FormValue(dm.DocType_Code_scrn)
	item.Name = r.FormValue(dm.DocType_Name_scrn)
	item.SYSCreated = r.FormValue(dm.DocType_SYSCreated_scrn)
	item.SYSCreatedBy = r.FormValue(dm.DocType_SYSCreatedBy_scrn)
	item.SYSCreatedHost = r.FormValue(dm.DocType_SYSCreatedHost_scrn)
	item.SYSUpdated = r.FormValue(dm.DocType_SYSUpdated_scrn)
	item.SYSUpdatedBy = r.FormValue(dm.DocType_SYSUpdatedBy_scrn)
	item.SYSUpdatedHost = r.FormValue(dm.DocType_SYSUpdatedHost_scrn)
	item.SYSDeleted = r.FormValue(dm.DocType_SYSDeleted_scrn)
	item.SYSDeletedBy = r.FormValue(dm.DocType_SYSDeletedBy_scrn)
	item.SYSDeletedHost = r.FormValue(dm.DocType_SYSDeletedHost_scrn)
	item.SYSDbVersion = r.FormValue(dm.DocType_SYSDbVersion_scrn)
	item.Comments = r.FormValue(dm.DocType_Comments_scrn)
	item.Migrate = r.FormValue(dm.DocType_Migrate_scrn)
	return item
}
//doctype_URIQueryData is used to replace the wildcards in the URI Query Path with the values from the DocType Data Model
func doctype_URIQueryData(queryPath string,item dm.DocType,itemID string) string {
	if queryPath == "" {
		return ""
	}
	queryPath = core.ReplaceWildcard(queryPath, strings.ToUpper("ID"), itemID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.DocType_DocTypeID_scrn), item.DocTypeID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.DocType_Code_scrn), item.Code)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.DocType_Name_scrn), item.Name)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.DocType_Comments_scrn), item.Comments)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.DocType_Migrate_scrn), item.Migrate)
	return queryPath
}