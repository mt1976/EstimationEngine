package application
// ----------------------------------------------------------------
// Automatically generated  "/application/credentials.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Credentials (credentials)
// Endpoint 	        : Credentials (Id)
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

//Credentials_Publish annouces the endpoints available for this object
func Credentials_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Credentials_Path, Credentials_Handler)
	mux.HandleFunc(dm.Credentials_PathList, Credentials_HandlerList)
	mux.HandleFunc(dm.Credentials_PathView, Credentials_HandlerView)
	mux.HandleFunc(dm.Credentials_PathEdit, Credentials_HandlerEdit)
	mux.HandleFunc(dm.Credentials_PathNew, Credentials_HandlerNew)
	mux.HandleFunc(dm.Credentials_PathSave, Credentials_HandlerSave)
	mux.HandleFunc(dm.Credentials_PathDelete, Credentials_HandlerDelete)
    core.API = core.API.AddRoute(dm.Credentials_Title, dm.Credentials_Path, "", dm.Credentials_QueryString, "Application")
	logs.Publish("Application", dm.Credentials_Title)
}

//Credentials_HandlerList is the handler for the Credentials list page
func Credentials_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Credentials

	objectName := dao.Translate("ObjectName", "Credentials")
	reqField := "Base"
	usage := "Defines a filter for the list of Credentials records." + core.TEXTAREA_CR
	usage = usage + "Fields can be any of those in the underlying DB table." + core.TEXTAREA_CR
	usage = usage + "Examples Below:" + core.TEXTAREA_CR
	usage = usage + "* datalength(_deleted) = 0 or " + core.TEXTAREA_CR 
	usage = usage + "* class IN ('x','y','z')"
	
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule,usage)
	if filter == "" {
		logs.Warning("No filter found : " + reqField + " for Object: " + objectName)
	} 

	noItems, returnList, _ := dao.Credentials_GetListFiltered(filter)

	pageDetail := dm.Credentials_PageList{
		Title:            CardTitle(dm.Credentials_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Credentials_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	nextTemplate :=  NextTemplate("Credentials", "List", dm.Credentials_TemplateList)
	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Credentials_HandlerView is the handler used to View a Credentials database record
func Credentials_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Credentials_QueryString)
	_, rD, _ := dao.Credentials_GetByID(searchID)

	pageDetail := dm.Credentials_Page{
		Title:       CardTitle(dm.Credentials_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Credentials_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = credentials_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Credentials", "View", dm.Credentials_TemplateView)
	nextTemplate = credentials_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Credentials_HandlerEdit is the handler used to Edit of an existing Credentials database record
func Credentials_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Credentials_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.Credentials
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.Credentials)
	} else {
		_, rD, _ = dao.Credentials_GetByID(searchID)
	}
	
	pageDetail := dm.Credentials_Page{
		Title:       CardTitle(dm.Credentials_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Credentials_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = credentials_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Credentials", "Edit", dm.Credentials_TemplateEdit)
	nextTemplate = credentials_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Credentials_HandlerSave is the handler used process the saving of an Credentials database record, either new or existing, referenced by Edit & New Handlers.
func Credentials_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("Id")
	logs.Servicing(r.URL.Path+itemID)

	item := credentials_DataFromRequest(r)
	
	item, errStore := dao.Credentials_Store(item,r)
	if errStore == nil {
		nextTemplate :=  NextTemplate("Credentials", "Save", dm.Credentials_Redirect)
		nextTemplate = credentials_URIQueryData(nextTemplate,item,itemID)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.Credentials_Name, errStore.Error())
		ExecuteRedirect(r.Referer(), w, r,dm.Credentials_QueryString,itemID,item)
	}
}

//Credentials_HandlerNew is the handler used process the creation of an new Credentials database record, then redirect to Edit
func Credentials_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Credentials_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.Credentials
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.Credentials)
	} else {
		_, _, rD, _ = dao.Credentials_New()
	}

	pageDetail := dm.Credentials_Page{
		Title:       CardTitle(dm.Credentials_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Credentials_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = credentials_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Credentials", "New", dm.Credentials_TemplateNew)
	nextTemplate = credentials_URIQueryData(nextTemplate,dm.Credentials{},searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}	

//Credentials_HandlerDelete is the handler used process the deletion of an Credentials database record. May be Hard or SoftDelete.
func Credentials_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Credentials_QueryString)
	// DAO Call to Delete Credentials Record, may be SoftDelete or HardDelete depending on the DAO implementation
	dao.Credentials_Delete(searchID)	

	nextTemplate :=  NextTemplate("Credentials", "Delete", dm.Credentials_Redirect)
	nextTemplate = credentials_URIQueryData(nextTemplate,dm.Credentials{},searchID)
	http.Redirect(w, r, nextTemplate, http.StatusFound)
}
//credentials_PopulatePage Builds/Populates the Credentials Page from an instance of Credentials from the Data Model
func credentials_PopulatePage(rD dm.Credentials, pageDetail dm.Credentials_Page) dm.Credentials_Page {
	// Real DB Fields
	pageDetail.SYSId = rD.SYSId
	pageDetail.Id = rD.Id
	pageDetail.Username = rD.Username
	pageDetail.Password = rD.Password
	pageDetail.Firstname = rD.Firstname
	pageDetail.Lastname = rD.Lastname
	pageDetail.Knownas = rD.Knownas
	pageDetail.Email = rD.Email
	pageDetail.Issued = rD.Issued
	pageDetail.Expiry = rD.Expiry
	pageDetail.RoleType = rD.RoleType
	pageDetail.Brand = rD.Brand
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.State = rD.State
	pageDetail.Notes = rD.Notes
	pageDetail.SYSDeleted = rD.SYSDeleted
	pageDetail.SYSDeletedBy = rD.SYSDeletedBy
	pageDetail.SYSDeletedHost = rD.SYSDeletedHost
	pageDetail.SYSActivity = rD.SYSActivity
	pageDetail.SYSDbVersion = rD.SYSDbVersion
	pageDetail.EmailNotifications = rD.EmailNotifications
	pageDetail.PasswordExpiry = rD.PasswordExpiry
	// Add Pseudo/Extra Fields, fields that are not in the DB but are used in the UI
	// Enrichment content, content used provide lookups,lists etc
	pageDetail.RoleType_lookup = dao.UserRole_GetLookup()
	pageDetail.State_lookup = dao.StubLists_Get("credentialStates")
	pageDetail.EmailNotifications_lookup = dao.StubLists_Get("tf")
	// Add the Properties for the Fields
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.Id_props = rD.Id_props
	pageDetail.Username_props = rD.Username_props
	pageDetail.Password_props = rD.Password_props
	pageDetail.Firstname_props = rD.Firstname_props
	pageDetail.Lastname_props = rD.Lastname_props
	pageDetail.Knownas_props = rD.Knownas_props
	pageDetail.Email_props = rD.Email_props
	pageDetail.Issued_props = rD.Issued_props
	pageDetail.Expiry_props = rD.Expiry_props
	pageDetail.RoleType_props = rD.RoleType_props
	pageDetail.Brand_props = rD.Brand_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.State_props = rD.State_props
	pageDetail.Notes_props = rD.Notes_props
	pageDetail.SYSDeleted_props = rD.SYSDeleted_props
	pageDetail.SYSDeletedBy_props = rD.SYSDeletedBy_props
	pageDetail.SYSDeletedHost_props = rD.SYSDeletedHost_props
	pageDetail.SYSActivity_props = rD.SYSActivity_props
	pageDetail.SYSDbVersion_props = rD.SYSDbVersion_props
	pageDetail.EmailNotifications_props = rD.EmailNotifications_props
	pageDetail.PasswordExpiry_props = rD.PasswordExpiry_props
	return pageDetail
}
//credentials_DataFromRequest is used process the content of an HTTP Request and return an instance of an Credentials
func credentials_DataFromRequest(r *http.Request) dm.Credentials {
	var item dm.Credentials
	item.SYSId = r.FormValue(dm.Credentials_SYSId_scrn)
	item.Id = r.FormValue(dm.Credentials_Id_scrn)
	item.Username = r.FormValue(dm.Credentials_Username_scrn)
	item.Password = r.FormValue(dm.Credentials_Password_scrn)
	item.Firstname = r.FormValue(dm.Credentials_Firstname_scrn)
	item.Lastname = r.FormValue(dm.Credentials_Lastname_scrn)
	item.Knownas = r.FormValue(dm.Credentials_Knownas_scrn)
	item.Email = r.FormValue(dm.Credentials_Email_scrn)
	item.Issued = r.FormValue(dm.Credentials_Issued_scrn)
	item.Expiry = r.FormValue(dm.Credentials_Expiry_scrn)
	item.RoleType = r.FormValue(dm.Credentials_RoleType_scrn)
	item.Brand = r.FormValue(dm.Credentials_Brand_scrn)
	item.SYSCreated = r.FormValue(dm.Credentials_SYSCreated_scrn)
	item.SYSUpdated = r.FormValue(dm.Credentials_SYSUpdated_scrn)
	item.SYSCreatedBy = r.FormValue(dm.Credentials_SYSCreatedBy_scrn)
	item.SYSCreatedHost = r.FormValue(dm.Credentials_SYSCreatedHost_scrn)
	item.SYSUpdatedBy = r.FormValue(dm.Credentials_SYSUpdatedBy_scrn)
	item.SYSUpdatedHost = r.FormValue(dm.Credentials_SYSUpdatedHost_scrn)
	item.State = r.FormValue(dm.Credentials_State_scrn)
	item.Notes = r.FormValue(dm.Credentials_Notes_scrn)
	item.SYSDeleted = r.FormValue(dm.Credentials_SYSDeleted_scrn)
	item.SYSDeletedBy = r.FormValue(dm.Credentials_SYSDeletedBy_scrn)
	item.SYSDeletedHost = r.FormValue(dm.Credentials_SYSDeletedHost_scrn)
	item.SYSActivity = r.FormValue(dm.Credentials_SYSActivity_scrn)
	item.SYSDbVersion = r.FormValue(dm.Credentials_SYSDbVersion_scrn)
	item.EmailNotifications = r.FormValue(dm.Credentials_EmailNotifications_scrn)
	item.PasswordExpiry = r.FormValue(dm.Credentials_PasswordExpiry_scrn)
	return item
}
//credentials_URIQueryData is used to replace the wildcards in the URI Query Path with the values from the Credentials Data Model
func credentials_URIQueryData(queryPath string,item dm.Credentials,itemID string) string {
	if queryPath == "" {
		return ""
	}
	queryPath = core.ReplaceWildcard(queryPath, strings.ToUpper("ID"), itemID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Credentials_Id_scrn), item.Id)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Credentials_Username_scrn), item.Username)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Credentials_Password_scrn), item.Password)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Credentials_Firstname_scrn), item.Firstname)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Credentials_Lastname_scrn), item.Lastname)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Credentials_Knownas_scrn), item.Knownas)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Credentials_Email_scrn), item.Email)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Credentials_Issued_scrn), item.Issued)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Credentials_Expiry_scrn), item.Expiry)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Credentials_RoleType_scrn), item.RoleType)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Credentials_Brand_scrn), item.Brand)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Credentials_State_scrn), item.State)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Credentials_Notes_scrn), item.Notes)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Credentials_EmailNotifications_scrn), item.EmailNotifications)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Credentials_PasswordExpiry_scrn), item.PasswordExpiry)
	return queryPath
}