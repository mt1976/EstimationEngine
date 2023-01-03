package application
// ----------------------------------------------------------------
// Automatically generated  "/application/externalmessage.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : ExternalMessage (externalmessage)
// Endpoint 	        : ExternalMessage (Message)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 03/01/2023 at 19:18:09
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)

//ExternalMessage_Publish annouces the endpoints available for this object
//ExternalMessage_Publish - Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
func ExternalMessage_Publish(mux http.ServeMux) {
	// START
	// Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	mux.HandleFunc(dm.ExternalMessage_Path, ExternalMessage_Handler)
	mux.HandleFunc(dm.ExternalMessage_PathList, ExternalMessage_HandlerList)
	mux.HandleFunc(dm.ExternalMessage_PathView, ExternalMessage_HandlerView)
	mux.HandleFunc(dm.ExternalMessage_PathEdit, ExternalMessage_HandlerEdit)
	//Cannot Create via GUI
	mux.HandleFunc(dm.ExternalMessage_PathSave, ExternalMessage_HandlerSave)
	mux.HandleFunc(dm.ExternalMessage_PathDelete, ExternalMessage_HandlerDelete)
	logs.Publish("Application", dm.ExternalMessage_Title)
    core.Catalog_Add(dm.ExternalMessage_Title, dm.ExternalMessage_Path, "", dm.ExternalMessage_QueryString, "Application")
	// 
	// Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//ExternalMessage_HandlerList is the handler for the list page
//Allows Listing of ExternalMessage records
//ExternalMessage_HandlerList - Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
func ExternalMessage_HandlerList(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
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

	var returnList []dm.ExternalMessage
	noItems, returnList, _ := dao.ExternalMessage_GetList()

	pageDetail := dm.ExternalMessage_PageList{
		Title:            CardTitle(dm.ExternalMessage_Title, core.Action_List),
		PageTitle:        PageTitle(dm.ExternalMessage_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.ExternalMessage_TemplateList, w, r, pageDetail)
	// 
	// Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END

}


//ExternalMessage_HandlerView is the handler used to View a page
//Allows Viewing for an existing ExternalMessage record
//ExternalMessage_HandlerView - Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local 
func ExternalMessage_HandlerView(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
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

	searchID := core.GetURLparam(r, dm.ExternalMessage_QueryString)
	_, rD, _ := dao.ExternalMessage_GetByID(searchID)

	pageDetail := dm.ExternalMessage_Page{
		Title:       CardTitle(dm.ExternalMessage_Title, core.Action_View),
		PageTitle:   PageTitle(dm.ExternalMessage_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = externalmessage_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.ExternalMessage_TemplateView, w, r, pageDetail)
	// 
	// Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//ExternalMessage_HandlerEdit is the handler used generate the Edit page
//Allows Editing for an existing ExternalMessage record and then allows the user to save the changes
//ExternalMessage_HandlerEdit - Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local 
func ExternalMessage_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.ExternalMessage_QueryString)
	_, rD, _ := dao.ExternalMessage_GetByID(searchID)
	
	pageDetail := dm.ExternalMessage_Page{
		Title:       CardTitle(dm.ExternalMessage_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.ExternalMessage_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = externalmessage_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.ExternalMessage_TemplateEdit, w, r, pageDetail)
	// 
	// Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//ExternalMessage_HandlerSave is the handler used process the saving of an ExternalMessage
//It is called from the Edit and New pages
//ExternalMessage_HandlerSave  - Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local 
func ExternalMessage_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("MessageID"))

	item := externalmessage_DataFromRequest(r)
	
	dao.ExternalMessage_Store(item,r)	
	http.Redirect(w, r, dm.ExternalMessage_Redirect, http.StatusFound)
	// 
	// Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}



//ExternalMessage_HandlerDelete is the handler used process the deletion of an ExternalMessage
// It will delete the ExternalMessage and then redirect to the List page
//ExternalMessage_HandlerDelete - Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local 
func ExternalMessage_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	searchID := core.GetURLparam(r, dm.ExternalMessage_QueryString)

	dao.ExternalMessage_Delete(searchID)	

	http.Redirect(w, r, dm.ExternalMessage_Redirect, http.StatusFound)
	// 
	// Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//externalmessage_PopulatePage Builds/Populates the ExternalMessage Page 
func externalmessage_PopulatePage(rD dm.ExternalMessage, pageDetail dm.ExternalMessage_Page) dm.ExternalMessage_Page {
	// START
	// Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.MessageID = rD.MessageID
	pageDetail.MessageFormat = rD.MessageFormat
	pageDetail.MessageDeliveredTo = rD.MessageDeliveredTo
	pageDetail.MessageBody = rD.MessageBody
	pageDetail.MessageFilename = rD.MessageFilename
	pageDetail.MessageLife = rD.MessageLife
	pageDetail.MessageDate = rD.MessageDate
	pageDetail.MessageTime = rD.MessageTime
	pageDetail.MessageTimeoutAction = rD.MessageTimeoutAction
	pageDetail.MessageACKNAK = rD.MessageACKNAK
	pageDetail.ResponseID = rD.ResponseID
	pageDetail.ResponseFilename = rD.ResponseFilename
	pageDetail.ResponseBody = rD.ResponseBody
	pageDetail.ResponseDate = rD.ResponseDate
	pageDetail.ResponseTime = rD.ResponseTime
	pageDetail.ResponseAdditionalInfo = rD.ResponseAdditionalInfo
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.MessageTimeout = rD.MessageTimeout
	pageDetail.MessageEmitted = rD.MessageEmitted
	pageDetail.ResponseRecieved = rD.ResponseRecieved
	pageDetail.MessageClass = rD.MessageClass
	pageDetail.AppID = rD.AppID
	pageDetail.SYSDeleted = rD.SYSDeleted
	pageDetail.SYSDeletedBy = rD.SYSDeletedBy
	pageDetail.SYSDeletedHost = rD.SYSDeletedHost
	pageDetail.SYSDbVersion = rD.SYSDbVersion
	
	
	//
	// Automatically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.MessageACKNAK_lookup = dao.StubLists_Get("messageACKNAK")
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.ResponseRecieved_lookup = dao.StubLists_Get("tf")
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.MessageID_props = rD.MessageID_props
	pageDetail.MessageFormat_props = rD.MessageFormat_props
	pageDetail.MessageDeliveredTo_props = rD.MessageDeliveredTo_props
	pageDetail.MessageBody_props = rD.MessageBody_props
	pageDetail.MessageFilename_props = rD.MessageFilename_props
	pageDetail.MessageLife_props = rD.MessageLife_props
	pageDetail.MessageDate_props = rD.MessageDate_props
	pageDetail.MessageTime_props = rD.MessageTime_props
	pageDetail.MessageTimeoutAction_props = rD.MessageTimeoutAction_props
	pageDetail.MessageACKNAK_props = rD.MessageACKNAK_props
	pageDetail.ResponseID_props = rD.ResponseID_props
	pageDetail.ResponseFilename_props = rD.ResponseFilename_props
	pageDetail.ResponseBody_props = rD.ResponseBody_props
	pageDetail.ResponseDate_props = rD.ResponseDate_props
	pageDetail.ResponseTime_props = rD.ResponseTime_props
	pageDetail.ResponseAdditionalInfo_props = rD.ResponseAdditionalInfo_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.MessageTimeout_props = rD.MessageTimeout_props
	pageDetail.MessageEmitted_props = rD.MessageEmitted_props
	pageDetail.ResponseRecieved_props = rD.ResponseRecieved_props
	pageDetail.MessageClass_props = rD.MessageClass_props
	pageDetail.AppID_props = rD.AppID_props
	pageDetail.SYSDeleted_props = rD.SYSDeleted_props
	pageDetail.SYSDeletedBy_props = rD.SYSDeletedBy_props
	pageDetail.SYSDeletedHost_props = rD.SYSDeletedHost_props
	pageDetail.SYSDbVersion_props = rD.SYSDbVersion_props
	
	// 
	// Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	


//externalmessage_DataFromRequest is used process the content of an HTTP Request and return an instance of an ExternalMessage
func externalmessage_DataFromRequest(r *http.Request) dm.ExternalMessage {
	// START
	// Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	var item dm.ExternalMessage
	// FIELD SET START
	// Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.ExternalMessage_SYSId_scrn)
		item.MessageID = r.FormValue(dm.ExternalMessage_MessageID_scrn)
		item.MessageFormat = r.FormValue(dm.ExternalMessage_MessageFormat_scrn)
		item.MessageDeliveredTo = r.FormValue(dm.ExternalMessage_MessageDeliveredTo_scrn)
		item.MessageBody = r.FormValue(dm.ExternalMessage_MessageBody_scrn)
		item.MessageFilename = r.FormValue(dm.ExternalMessage_MessageFilename_scrn)
		item.MessageLife = r.FormValue(dm.ExternalMessage_MessageLife_scrn)
		item.MessageDate = r.FormValue(dm.ExternalMessage_MessageDate_scrn)
		item.MessageTime = r.FormValue(dm.ExternalMessage_MessageTime_scrn)
		item.MessageTimeoutAction = r.FormValue(dm.ExternalMessage_MessageTimeoutAction_scrn)
		item.MessageACKNAK = r.FormValue(dm.ExternalMessage_MessageACKNAK_scrn)
		item.ResponseID = r.FormValue(dm.ExternalMessage_ResponseID_scrn)
		item.ResponseFilename = r.FormValue(dm.ExternalMessage_ResponseFilename_scrn)
		item.ResponseBody = r.FormValue(dm.ExternalMessage_ResponseBody_scrn)
		item.ResponseDate = r.FormValue(dm.ExternalMessage_ResponseDate_scrn)
		item.ResponseTime = r.FormValue(dm.ExternalMessage_ResponseTime_scrn)
		item.ResponseAdditionalInfo = r.FormValue(dm.ExternalMessage_ResponseAdditionalInfo_scrn)
		item.SYSCreated = r.FormValue(dm.ExternalMessage_SYSCreated_scrn)
		item.SYSCreatedBy = r.FormValue(dm.ExternalMessage_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.ExternalMessage_SYSCreatedHost_scrn)
		item.SYSUpdated = r.FormValue(dm.ExternalMessage_SYSUpdated_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.ExternalMessage_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.ExternalMessage_SYSUpdatedHost_scrn)
		item.MessageTimeout = r.FormValue(dm.ExternalMessage_MessageTimeout_scrn)
		item.MessageEmitted = r.FormValue(dm.ExternalMessage_MessageEmitted_scrn)
		item.ResponseRecieved = r.FormValue(dm.ExternalMessage_ResponseRecieved_scrn)
		item.MessageClass = r.FormValue(dm.ExternalMessage_MessageClass_scrn)
		item.AppID = r.FormValue(dm.ExternalMessage_AppID_scrn)
		item.SYSDeleted = r.FormValue(dm.ExternalMessage_SYSDeleted_scrn)
		item.SYSDeletedBy = r.FormValue(dm.ExternalMessage_SYSDeletedBy_scrn)
		item.SYSDeletedHost = r.FormValue(dm.ExternalMessage_SYSDeletedHost_scrn)
		item.SYSDbVersion = r.FormValue(dm.ExternalMessage_SYSDbVersion_scrn)
	
	// 
	// Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return item
}
