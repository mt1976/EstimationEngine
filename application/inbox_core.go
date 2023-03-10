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
// Date & Time		    : 10/03/2023 at 19:54:34
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)

//Inbox_Publish annouces the endpoints available for this object
//Inbox_Publish - Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
func Inbox_Publish(mux http.ServeMux) {
	// START
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	mux.HandleFunc(dm.Inbox_Path, Inbox_Handler)
	mux.HandleFunc(dm.Inbox_PathList, Inbox_HandlerList)
	mux.HandleFunc(dm.Inbox_PathView, Inbox_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	mux.HandleFunc(dm.Inbox_PathSave, Inbox_HandlerSave)
	mux.HandleFunc(dm.Inbox_PathDelete, Inbox_HandlerDelete)
	logs.Publish("Application", dm.Inbox_Title)
    core.API = core.API.AddRoute(dm.Inbox_Title, dm.Inbox_Path, "", dm.Inbox_QueryString, "Application")
	// 
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//Inbox_HandlerList is the handler for the list page
//Allows Listing of Inbox records
//Inbox_HandlerList - Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
func Inbox_HandlerList(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
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

	var returnList []dm.Inbox

	objectName := dao.Translate("ObjectName", "Inbox")
	reqField := "Base"
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule)
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
	// 
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END

}


//Inbox_HandlerView is the handler used to View a page
//Allows Viewing for an existing Inbox record
//Inbox_HandlerView - Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func Inbox_HandlerView(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
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

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}



//Inbox_HandlerSave is the handler used process the saving of an Inbox
//It is called from the Edit and New pages
//Inbox_HandlerSave  - Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func Inbox_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	// Mandatory Security Validation
	//
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
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.Inbox_Name, errStore.Error())
		//http.Redirect(w, r, r.Referer(), http.StatusFound)
		ExecuteRedirect(r.Referer(), w, r,dm.Inbox_QueryString,itemID,item)
	}
	// 
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}



//Inbox_HandlerDelete is the handler used process the deletion of an Inbox
// It will delete the Inbox and then redirect to the List page
//Inbox_HandlerDelete - Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func Inbox_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	searchID := core.GetURLparam(r, dm.Inbox_QueryString)

	dao.Inbox_Delete(searchID)	

	nextTemplate :=  NextTemplate("Inbox", "Delete", dm.Inbox_Redirect)
	http.Redirect(w, r, nextTemplate, http.StatusFound)
	// 
	// Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//inbox_PopulatePage Builds/Populates the Inbox Page 
//inbox_PopulatePage Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	// Add Pseudo/Extra Fields
	// Enrichment Fields 
	
	pageDetail.MailUnread_lookup = dao.StubLists_Get("tf")
	
	pageDetail.MailAcknowledged_lookup = dao.StubLists_Get("tf")
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
//inbox_DataFromRequest Auto generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
