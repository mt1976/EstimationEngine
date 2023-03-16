package application
// ----------------------------------------------------------------
// Automatically generated  "/application/inbox.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Inbox (inbox)
// Endpoint 	        : Inbox (Message)
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

//Inbox_Publish annouces the endpoints available for this object
func Inbox_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Inbox_Path, Inbox_Handler)
	mux.HandleFunc(dm.Inbox_PathList, Inbox_HandlerList)
	mux.HandleFunc(dm.Inbox_PathView, Inbox_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	mux.HandleFunc(dm.Inbox_PathSave, Inbox_HandlerSave)
	mux.HandleFunc(dm.Inbox_PathDelete, Inbox_HandlerDelete)
    core.API = core.API.AddRoute(dm.Inbox_Title, dm.Inbox_Path, "", dm.Inbox_QueryString, "Application")
	logs.Publish("Application", dm.Inbox_Title)
}

//Inbox_HandlerList is the handler for the Inbox list page
func Inbox_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Inbox

	objectName := dao.Translate("ObjectName", "Inbox")
	reqField := "Base"
	usage := "Defines a filter for the list of Inbox records." + core.TEXTAREA_CR
	usage = usage + "Fields can be any of those in the underlying DB table." + core.TEXTAREA_CR
	usage = usage + "Examples Below:" + core.TEXTAREA_CR
	usage = usage + "* datalength(_deleted) = 0 or " + core.TEXTAREA_CR 
	usage = usage + "* class IN ('x','y','z')"
	
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule,usage)
	if filter == "" {
		logs.Warning("No filter found : " + reqField + " for Object: " + objectName)
	} 

	noItems, returnList, _ := dao.Inbox_GetListFiltered(filter)

	pageDetail := dm.Inbox_PageList{
		Title:            CardTitle(dm.Inbox_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Inbox_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	nextTemplate :=  NextTemplate("Inbox", "List", dm.Inbox_TemplateList)
	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Inbox_HandlerView is the handler used to View a Inbox database record
func Inbox_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Inbox_QueryString)
	_, rD, _ := dao.Inbox_GetByID(searchID)

	pageDetail := dm.Inbox_Page{
		Title:       CardTitle(dm.Inbox_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Inbox_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = inbox_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Inbox", "View", dm.Inbox_TemplateView)
	nextTemplate = inbox_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Inbox_HandlerSave is the handler used process the saving of an Inbox database record, either new or existing, referenced by Edit & New Handlers.
func Inbox_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("MailId")
	logs.Servicing(r.URL.Path+itemID)

	item := inbox_DataFromRequest(r)
	
	item, errStore := dao.Inbox_Store(item,r)
	if errStore == nil {
		nextTemplate :=  NextTemplate("Inbox", "Save", dm.Inbox_Redirect)
		nextTemplate = inbox_URIQueryData(nextTemplate,item,itemID)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.Inbox_Name, errStore.Error())
		ExecuteRedirect(r.Referer(), w, r,dm.Inbox_QueryString,itemID,item)
	}
}

//Inbox_HandlerDelete is the handler used process the deletion of an Inbox database record. May be Hard or SoftDelete.
func Inbox_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Inbox_QueryString)
	// DAO Call to Delete Inbox Record, may be SoftDelete or HardDelete depending on the DAO implementation
	dao.Inbox_Delete(searchID)	

	nextTemplate :=  NextTemplate("Inbox", "Delete", dm.Inbox_Redirect)
	nextTemplate = inbox_URIQueryData(nextTemplate,dm.Inbox{},searchID)
	http.Redirect(w, r, nextTemplate, http.StatusFound)
}
//inbox_PopulatePage Builds/Populates the Inbox Page from an instance of Inbox from the Data Model
func inbox_PopulatePage(rD dm.Inbox, pageDetail dm.Inbox_Page) dm.Inbox_Page {
	// Real DB Fields
	pageDetail.SYSId = rD.SYSId
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.MailId = rD.MailId
	pageDetail.MailTo = rD.MailTo
	pageDetail.MailFrom = rD.MailFrom
	pageDetail.MailSource = rD.MailSource
	pageDetail.MailSent = rD.MailSent
	pageDetail.MailUnread = rD.MailUnread
	pageDetail.MailSubject = rD.MailSubject
	pageDetail.MailContent = rD.MailContent
	pageDetail.MailAcknowledged = rD.MailAcknowledged
	pageDetail.SYSDeleted = rD.SYSDeleted
	pageDetail.SYSDeletedBy = rD.SYSDeletedBy
	pageDetail.SYSDeletedHost = rD.SYSDeletedHost
	pageDetail.SYSDbVersion = rD.SYSDbVersion
	// Add Pseudo/Extra Fields, fields that are not in the DB but are used in the UI
	// Enrichment content, content used provide lookups,lists etc
	pageDetail.MailUnread_lookup = dao.StubLists_Get("tf")
	pageDetail.MailAcknowledged_lookup = dao.StubLists_Get("tf")
	// Add the Properties for the Fields
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.MailId_props = rD.MailId_props
	pageDetail.MailTo_props = rD.MailTo_props
	pageDetail.MailFrom_props = rD.MailFrom_props
	pageDetail.MailSource_props = rD.MailSource_props
	pageDetail.MailSent_props = rD.MailSent_props
	pageDetail.MailUnread_props = rD.MailUnread_props
	pageDetail.MailSubject_props = rD.MailSubject_props
	pageDetail.MailContent_props = rD.MailContent_props
	pageDetail.MailAcknowledged_props = rD.MailAcknowledged_props
	pageDetail.SYSDeleted_props = rD.SYSDeleted_props
	pageDetail.SYSDeletedBy_props = rD.SYSDeletedBy_props
	pageDetail.SYSDeletedHost_props = rD.SYSDeletedHost_props
	pageDetail.SYSDbVersion_props = rD.SYSDbVersion_props
	return pageDetail
}
//inbox_DataFromRequest is used process the content of an HTTP Request and return an instance of an Inbox
func inbox_DataFromRequest(r *http.Request) dm.Inbox {
	var item dm.Inbox
	item.SYSId = r.FormValue(dm.Inbox_SYSId_scrn)
	item.SYSCreated = r.FormValue(dm.Inbox_SYSCreated_scrn)
	item.SYSUpdated = r.FormValue(dm.Inbox_SYSUpdated_scrn)
	item.SYSCreatedBy = r.FormValue(dm.Inbox_SYSCreatedBy_scrn)
	item.SYSCreatedHost = r.FormValue(dm.Inbox_SYSCreatedHost_scrn)
	item.SYSUpdatedBy = r.FormValue(dm.Inbox_SYSUpdatedBy_scrn)
	item.SYSUpdatedHost = r.FormValue(dm.Inbox_SYSUpdatedHost_scrn)
	item.MailId = r.FormValue(dm.Inbox_MailId_scrn)
	item.MailTo = r.FormValue(dm.Inbox_MailTo_scrn)
	item.MailFrom = r.FormValue(dm.Inbox_MailFrom_scrn)
	item.MailSource = r.FormValue(dm.Inbox_MailSource_scrn)
	item.MailSent = r.FormValue(dm.Inbox_MailSent_scrn)
	item.MailUnread = r.FormValue(dm.Inbox_MailUnread_scrn)
	item.MailSubject = r.FormValue(dm.Inbox_MailSubject_scrn)
	item.MailContent = r.FormValue(dm.Inbox_MailContent_scrn)
	item.MailAcknowledged = r.FormValue(dm.Inbox_MailAcknowledged_scrn)
	item.SYSDeleted = r.FormValue(dm.Inbox_SYSDeleted_scrn)
	item.SYSDeletedBy = r.FormValue(dm.Inbox_SYSDeletedBy_scrn)
	item.SYSDeletedHost = r.FormValue(dm.Inbox_SYSDeletedHost_scrn)
	item.SYSDbVersion = r.FormValue(dm.Inbox_SYSDbVersion_scrn)
	return item
}
//inbox_URIQueryData is used to replace the wildcards in the URI Query Path with the values from the Inbox Data Model
func inbox_URIQueryData(queryPath string,item dm.Inbox,itemID string) string {
	if queryPath == "" {
		return ""
	}
	queryPath = core.ReplaceWildcard(queryPath, strings.ToUpper("ID"), itemID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Inbox_MailId_scrn), item.MailId)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Inbox_MailTo_scrn), item.MailTo)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Inbox_MailFrom_scrn), item.MailFrom)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Inbox_MailSource_scrn), item.MailSource)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Inbox_MailSent_scrn), item.MailSent)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Inbox_MailUnread_scrn), item.MailUnread)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Inbox_MailSubject_scrn), item.MailSubject)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Inbox_MailContent_scrn), item.MailContent)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Inbox_MailAcknowledged_scrn), item.MailAcknowledged)
	return queryPath
}