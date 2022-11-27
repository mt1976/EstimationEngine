package application

// ----------------------------------------------------------------
// Automatically generated  "/application/message.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Message (message)
// Endpoint 	        : Message (Message)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:54
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"net/http"

	core "github.com/mt1976/ebEstimates/core"
	dao "github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// Message_Publish annouces the endpoints available for this object
func Message_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.Message_PathList, Message_HandlerList)
	mux.HandleFunc(dm.Message_PathView, Message_HandlerView)
	mux.HandleFunc(dm.Message_PathEdit, Message_HandlerEdit)
	//Cannot Create via GUI
	mux.HandleFunc(dm.Message_PathSave, Message_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Message_Title)
	//No API
}

// Message_HandlerList is the handler for the list page
func Message_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Message
	noItems, returnList, _ := dao.Message_GetList()

	pageDetail := dm.Message_PageList{
		Title:       CardTitle(dm.Message_Title, core.Action_List),
		PageTitle:   PageTitle(dm.Message_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:    returnList,
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Message_TemplateList, w, r, pageDetail)

}

// Message_HandlerView is the handler used to View a page
func Message_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Message_QueryString)
	_, rD, _ := dao.Message_GetByID(searchID)

	pageDetail := dm.Message_Page{
		Title:     CardTitle(dm.Message_Title, core.Action_View),
		PageTitle: PageTitle(dm.Message_Title, core.Action_View),
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = message_PopulatePage(rD, pageDetail)

	ExecuteTemplate(dm.Message_TemplateView, w, r, pageDetail)

}

// Message_HandlerEdit is the handler used generate the Edit page
func Message_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Message_QueryString)
	_, rD, _ := dao.Message_GetByID(searchID)

	pageDetail := dm.Message_Page{
		Title:     CardTitle(dm.Message_Title, core.Action_Edit),
		PageTitle: PageTitle(dm.Message_Title, core.Action_Edit),
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = message_PopulatePage(rD, pageDetail)

	ExecuteTemplate(dm.Message_TemplateEdit, w, r, pageDetail)
}

// Message_HandlerSave is the handler used process the saving of an Message
func Message_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path + r.FormValue("Id"))

	var item dm.Message
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	item.SYSId = r.FormValue(dm.Message_SYSId_scrn)
	item.Id = r.FormValue(dm.Message_Id_scrn)
	item.Message = r.FormValue(dm.Message_Message_scrn)
	item.SYSCreated = r.FormValue(dm.Message_SYSCreated_scrn)
	item.SYSWho = r.FormValue(dm.Message_SYSWho_scrn)
	item.SYSHost = r.FormValue(dm.Message_SYSHost_scrn)
	item.SYSUpdated = r.FormValue(dm.Message_SYSUpdated_scrn)
	item.SYSCreatedBy = r.FormValue(dm.Message_SYSCreatedBy_scrn)
	item.SYSCreatedHost = r.FormValue(dm.Message_SYSCreatedHost_scrn)
	item.SYSUpdatedBy = r.FormValue(dm.Message_SYSUpdatedBy_scrn)
	item.SYSUpdatedHost = r.FormValue(dm.Message_SYSUpdatedHost_scrn)

	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	dao.Message_Store(item, r)
	http.Redirect(w, r, dm.Message_Redirect, http.StatusFound)
}

// Builds/Popuplates the Message Page
func message_PopulatePage(rD dm.Message, pageDetail dm.Message_Page) dm.Message_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.Id = rD.Id
	pageDetail.Message = rD.Message
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSWho = rD.SYSWho
	pageDetail.SYSHost = rD.SYSHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost

	//
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//

	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.Id_props = rD.Id_props
	pageDetail.Message_props = rD.Message_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSWho_props = rD.SYSWho_props
	pageDetail.SYSHost_props = rD.SYSHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props

	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
	return pageDetail
}
