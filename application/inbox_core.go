package application
// ----------------------------------------------------------------
// Automatically generated  "/application/inbox.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Inbox (inbox)
// Endpoint 	        : Inbox (Message)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 30/12/2022 at 12:18:21
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
//Inbox_Publish - Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
func Inbox_Publish(mux http.ServeMux) {
	// START
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// 
	mux.HandleFunc(dm.Inbox_Path, Inbox_Handler)
	mux.HandleFunc(dm.Inbox_PathList, Inbox_HandlerList)
	mux.HandleFunc(dm.Inbox_PathView, Inbox_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	mux.HandleFunc(dm.Inbox_PathSave, Inbox_HandlerSave)
	mux.HandleFunc(dm.Inbox_PathDelete, Inbox_HandlerDelete)
	logs.Publish("Application", dm.Inbox_Title)
    core.Catalog_Add(dm.Inbox_Title, dm.Inbox_Path, "", dm.Inbox_QueryString, "Application")
	// 
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//Inbox_HandlerList is the handler for the list page
//Allows Listing of Inbox records
//Inbox_HandlerList - Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
func Inbox_HandlerList(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
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
	noItems, returnList, _ := dao.Inbox_GetList()

	pageDetail := dm.Inbox_PageList{
		Title:            CardTitle(dm.Inbox_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Inbox_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Inbox_TemplateList, w, r, pageDetail)
	// 
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

}


//Inbox_HandlerView is the handler used to View a page
//Allows Viewing for an existing Inbox record
//Inbox_HandlerView - Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
func Inbox_HandlerView(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
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

	ExecuteTemplate(dm.Inbox_TemplateView, w, r, pageDetail)
	// 
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}



//Inbox_HandlerSave is the handler used process the saving of an Inbox
//It is called from the Edit and New pages
//Inbox_HandlerSave  - Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
func Inbox_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// 
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("MailId"))

	item := inbox_DataFromRequest(r)
	
	dao.Inbox_Store(item,r)	
	http.Redirect(w, r, dm.Inbox_Redirect, http.StatusFound)
	// 
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}



//Inbox_HandlerDelete is the handler used process the deletion of an Inbox
// It will delete the Inbox and then redirect to the List page
//Inbox_HandlerDelete - Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
func Inbox_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
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

	http.Redirect(w, r, dm.Inbox_Redirect, http.StatusFound)
	// 
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//inbox_PopulatePage Builds/Populates the Inbox Page 
func inbox_PopulatePage(rD dm.Inbox, pageDetail dm.Inbox_Page) dm.Inbox_Page {
	// START
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
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
	
	
	//
	// Automatically generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
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
	
	// 
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	


//inbox_DataFromRequest is used process the content of an HTTP Request and return an instance of an Inbox
func inbox_DataFromRequest(r *http.Request) dm.Inbox {
	// START
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	var item dm.Inbox
	// FIELD SET START
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
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
	
	// 
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return item
}
