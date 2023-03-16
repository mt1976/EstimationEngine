package application
// ----------------------------------------------------------------
// Automatically generated  "/application/userrole.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : UserRole (userrole)
// Endpoint 	        : UserRole (Id)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/03/2023 at 19:24:50
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

//UserRole_Publish annouces the endpoints available for this object
func UserRole_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.UserRole_Path, UserRole_Handler)
	mux.HandleFunc(dm.UserRole_PathList, UserRole_HandlerList)
	mux.HandleFunc(dm.UserRole_PathView, UserRole_HandlerView)
	mux.HandleFunc(dm.UserRole_PathEdit, UserRole_HandlerEdit)
	mux.HandleFunc(dm.UserRole_PathNew, UserRole_HandlerNew)
	mux.HandleFunc(dm.UserRole_PathSave, UserRole_HandlerSave)
	mux.HandleFunc(dm.UserRole_PathDelete, UserRole_HandlerDelete)
    core.API = core.API.AddRoute(dm.UserRole_Title, dm.UserRole_Path, "", dm.UserRole_QueryString, "Application")
	logs.Publish("Application", dm.UserRole_Title)
}

//UserRole_HandlerList is the handler for the UserRole list page
func UserRole_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.UserRole

	objectName := dao.Translate("ObjectName", "UserRole")
	reqField := "Base"
	usage := "Defines a filter for the list of UserRole records." + core.TEXTAREA_CR
	usage = usage + "Fields can be any of those in the underlying DB table." + core.TEXTAREA_CR
	usage = usage + "Examples Below:" + core.TEXTAREA_CR
	usage = usage + "* datalength(_deleted) = 0 or " + core.TEXTAREA_CR 
	usage = usage + "* class IN ('x','y','z')"
	
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule,usage)
	if filter == "" {
		logs.Warning("No filter found : " + reqField + " for Object: " + objectName)
	} 

	noItems, returnList, _ := dao.UserRole_GetListFiltered(filter)

	pageDetail := dm.UserRole_PageList{
		Title:            CardTitle(dm.UserRole_Title, core.Action_List),
		PageTitle:        PageTitle(dm.UserRole_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	nextTemplate :=  NextTemplate("UserRole", "List", dm.UserRole_TemplateList)
	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//UserRole_HandlerView is the handler used to View a UserRole database record
func UserRole_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.UserRole_QueryString)
	_, rD, _ := dao.UserRole_GetByID(searchID)

	pageDetail := dm.UserRole_Page{
		Title:       CardTitle(dm.UserRole_Title, core.Action_View),
		PageTitle:   PageTitle(dm.UserRole_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = userrole_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("UserRole", "View", dm.UserRole_TemplateView)
	nextTemplate = userrole_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//UserRole_HandlerEdit is the handler used to Edit of an existing UserRole database record
func UserRole_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.UserRole_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.UserRole
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.UserRole)
	} else {
		_, rD, _ = dao.UserRole_GetByID(searchID)
	}
	
	pageDetail := dm.UserRole_Page{
		Title:       CardTitle(dm.UserRole_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.UserRole_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = userrole_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("UserRole", "Edit", dm.UserRole_TemplateEdit)
	nextTemplate = userrole_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//UserRole_HandlerSave is the handler used process the saving of an UserRole database record, either new or existing, referenced by Edit & New Handlers.
func UserRole_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("Id")
	logs.Servicing(r.URL.Path+itemID)

	item := userrole_DataFromRequest(r)
	
	item, errStore := dao.UserRole_Store(item,r)
	if errStore == nil {
		nextTemplate :=  NextTemplate("UserRole", "Save", dm.UserRole_Redirect)
		nextTemplate = userrole_URIQueryData(nextTemplate,item,itemID)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.UserRole_Name, errStore.Error())
		ExecuteRedirect(r.Referer(), w, r,dm.UserRole_QueryString,itemID,item)
	}
}

//UserRole_HandlerNew is the handler used process the creation of an new UserRole database record, then redirect to Edit
func UserRole_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.UserRole_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.UserRole
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.UserRole)
	} else {
		_, _, rD, _ = dao.UserRole_New()
	}

	pageDetail := dm.UserRole_Page{
		Title:       CardTitle(dm.UserRole_Title, core.Action_New),
		PageTitle:   PageTitle(dm.UserRole_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = userrole_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("UserRole", "New", dm.UserRole_TemplateNew)
	nextTemplate = userrole_URIQueryData(nextTemplate,dm.UserRole{},searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}	

//UserRole_HandlerDelete is the handler used process the deletion of an UserRole database record. May be Hard or SoftDelete.
func UserRole_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.UserRole_QueryString)
	// DAO Call to Delete UserRole Record, may be SoftDelete or HardDelete depending on the DAO implementation
	dao.UserRole_Delete(searchID)	

	nextTemplate :=  NextTemplate("UserRole", "Delete", dm.UserRole_Redirect)
	nextTemplate = userrole_URIQueryData(nextTemplate,dm.UserRole{},searchID)
	http.Redirect(w, r, nextTemplate, http.StatusFound)
}
//userrole_PopulatePage Builds/Populates the UserRole Page from an instance of UserRole from the Data Model
func userrole_PopulatePage(rD dm.UserRole, pageDetail dm.UserRole_Page) dm.UserRole_Page {
	// Real DB Fields
	pageDetail.SYSId = rD.SYSId
	pageDetail.Id = rD.Id
	pageDetail.Name = rD.Name
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSDeleted = rD.SYSDeleted
	pageDetail.SYSDeletedBy = rD.SYSDeletedBy
	pageDetail.SYSDeletedHost = rD.SYSDeletedHost
	pageDetail.SYSDbVersion = rD.SYSDbVersion
	// Add Pseudo/Extra Fields, fields that are not in the DB but are used in the UI
	// Enrichment content, content used provide lookups,lists etc
	// Add the Properties for the Fields
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.Id_props = rD.Id_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSDeleted_props = rD.SYSDeleted_props
	pageDetail.SYSDeletedBy_props = rD.SYSDeletedBy_props
	pageDetail.SYSDeletedHost_props = rD.SYSDeletedHost_props
	pageDetail.SYSDbVersion_props = rD.SYSDbVersion_props
	return pageDetail
}
//userrole_DataFromRequest is used process the content of an HTTP Request and return an instance of an UserRole
func userrole_DataFromRequest(r *http.Request) dm.UserRole {
	var item dm.UserRole
	item.SYSId = r.FormValue(dm.UserRole_SYSId_scrn)
	item.Id = r.FormValue(dm.UserRole_Id_scrn)
	item.Name = r.FormValue(dm.UserRole_Name_scrn)
	item.SYSCreatedBy = r.FormValue(dm.UserRole_SYSCreatedBy_scrn)
	item.SYSCreatedHost = r.FormValue(dm.UserRole_SYSCreatedHost_scrn)
	item.SYSUpdatedBy = r.FormValue(dm.UserRole_SYSUpdatedBy_scrn)
	item.SYSUpdatedHost = r.FormValue(dm.UserRole_SYSUpdatedHost_scrn)
	item.SYSUpdated = r.FormValue(dm.UserRole_SYSUpdated_scrn)
	item.SYSCreated = r.FormValue(dm.UserRole_SYSCreated_scrn)
	item.SYSDeleted = r.FormValue(dm.UserRole_SYSDeleted_scrn)
	item.SYSDeletedBy = r.FormValue(dm.UserRole_SYSDeletedBy_scrn)
	item.SYSDeletedHost = r.FormValue(dm.UserRole_SYSDeletedHost_scrn)
	item.SYSDbVersion = r.FormValue(dm.UserRole_SYSDbVersion_scrn)
	return item
}
//userrole_URIQueryData is used to replace the wildcards in the URI Query Path with the values from the UserRole Data Model
func userrole_URIQueryData(queryPath string,item dm.UserRole,itemID string) string {
	if queryPath == "" {
		return ""
	}
	queryPath = core.ReplaceWildcard(queryPath, strings.ToUpper("ID"), itemID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.UserRole_Id_scrn), item.Id)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.UserRole_Name_scrn), item.Name)
	return queryPath
}