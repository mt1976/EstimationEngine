package application
// ----------------------------------------------------------------
// Automatically generated  "/application/session.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Session (session)
// Endpoint 	        : Session (SessionID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 16/12/2022 at 16:47:12
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)

//Session_Publish annouces the endpoints available for this object
//Session_Publish - Auto generated 16/12/2022 by matttownsend (Matt Townsend) on silicon.local
func Session_Publish(mux http.ServeMux) {
	// START
	// Auto generated 16/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// 
	//No API
	mux.HandleFunc(dm.Session_PathList, Session_HandlerList)
	mux.HandleFunc(dm.Session_PathView, Session_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	mux.HandleFunc(dm.Session_PathSave, Session_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Session_Title)
    //No API
	// 
	// Auto generated 16/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//Session_HandlerList is the handler for the list page
//Allows Listing of Session records
//Session_HandlerList - Auto generated 16/12/2022 by matttownsend (Matt Townsend) on silicon.local
func Session_HandlerList(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 16/12/2022 by matttownsend (Matt Townsend) on silicon.local
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

	var returnList []dm.Session
	noItems, returnList, _ := dao.Session_GetList()

	pageDetail := dm.Session_PageList{
		Title:            CardTitle(dm.Session_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Session_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Session_TemplateList, w, r, pageDetail)
	// 
	// Auto generated 16/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

}


//Session_HandlerView is the handler used to View a page
//Allows Viewing for an existing Session record
//Session_HandlerView - Auto generated 16/12/2022 by matttownsend (Matt Townsend) on silicon.local 
func Session_HandlerView(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 16/12/2022 by matttownsend (Matt Townsend) on silicon.local
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

	searchID := core.GetURLparam(r, dm.Session_QueryString)
	_, rD, _ := dao.Session_GetByID(searchID)

	pageDetail := dm.Session_Page{
		Title:       CardTitle(dm.Session_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Session_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = session_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Session_TemplateView, w, r, pageDetail)
	// 
	// Auto generated 16/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}



//Session_HandlerSave is the handler used process the saving of an Session
//It is called from the Edit and New pages
//Session_HandlerSave  - Auto generated 16/12/2022 by matttownsend (Matt Townsend) on silicon.local 
func Session_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 16/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// 
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Id"))

	item := session_DataFromRequest(r)
	
	dao.Session_Store(item,r)	
	http.Redirect(w, r, dm.Session_Redirect, http.StatusFound)
	// 
	// Auto generated 16/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}




//session_PopulatePage Builds/Populates the Session Page 
func session_PopulatePage(rD dm.Session, pageDetail dm.Session_Page) dm.Session_Page {
	// START
	// Auto generated 16/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.Id = rD.Id
	pageDetail.Apptoken = rD.Apptoken
	pageDetail.Createdate = rD.Createdate
	pageDetail.Createtime = rD.Createtime
	pageDetail.Uniqueid = rD.Uniqueid
	pageDetail.Sessiontoken = rD.Sessiontoken
	pageDetail.Username = rD.Username
	pageDetail.Password = rD.Password
	pageDetail.Userip = rD.Userip
	pageDetail.Userhost = rD.Userhost
	pageDetail.Appip = rD.Appip
	pageDetail.Apphost = rD.Apphost
	pageDetail.Issued = rD.Issued
	pageDetail.Expiry = rD.Expiry
	pageDetail.Expiryraw = rD.Expiryraw
	pageDetail.Expires = rD.Expires
	pageDetail.Brand = rD.Brand
	pageDetail.SessionRole = rD.SessionRole
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
	
	
	//
	// Automatically generated 16/12/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.Id_props = rD.Id_props
	pageDetail.Apptoken_props = rD.Apptoken_props
	pageDetail.Createdate_props = rD.Createdate_props
	pageDetail.Createtime_props = rD.Createtime_props
	pageDetail.Uniqueid_props = rD.Uniqueid_props
	pageDetail.Sessiontoken_props = rD.Sessiontoken_props
	pageDetail.Username_props = rD.Username_props
	pageDetail.Password_props = rD.Password_props
	pageDetail.Userip_props = rD.Userip_props
	pageDetail.Userhost_props = rD.Userhost_props
	pageDetail.Appip_props = rD.Appip_props
	pageDetail.Apphost_props = rD.Apphost_props
	pageDetail.Issued_props = rD.Issued_props
	pageDetail.Expiry_props = rD.Expiry_props
	pageDetail.Expiryraw_props = rD.Expiryraw_props
	pageDetail.Expires_props = rD.Expires_props
	pageDetail.Brand_props = rD.Brand_props
	pageDetail.SessionRole_props = rD.SessionRole_props
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
	
	// 
	// Auto generated 16/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	


//session_DataFromRequest is used process the content of an HTTP Request and return an instance of an Session
func session_DataFromRequest(r *http.Request) dm.Session {
	// START
	// Auto generated 16/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	var item dm.Session
	// FIELD SET START
	// Auto generated 16/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.Session_SYSId_scrn)
		item.Id = r.FormValue(dm.Session_Id_scrn)
		item.Apptoken = r.FormValue(dm.Session_Apptoken_scrn)
		item.Createdate = r.FormValue(dm.Session_Createdate_scrn)
		item.Createtime = r.FormValue(dm.Session_Createtime_scrn)
		item.Uniqueid = r.FormValue(dm.Session_Uniqueid_scrn)
		item.Sessiontoken = r.FormValue(dm.Session_Sessiontoken_scrn)
		item.Username = r.FormValue(dm.Session_Username_scrn)
		item.Password = r.FormValue(dm.Session_Password_scrn)
		item.Userip = r.FormValue(dm.Session_Userip_scrn)
		item.Userhost = r.FormValue(dm.Session_Userhost_scrn)
		item.Appip = r.FormValue(dm.Session_Appip_scrn)
		item.Apphost = r.FormValue(dm.Session_Apphost_scrn)
		item.Issued = r.FormValue(dm.Session_Issued_scrn)
		item.Expiry = r.FormValue(dm.Session_Expiry_scrn)
		item.Expiryraw = r.FormValue(dm.Session_Expiryraw_scrn)
		item.Expires = r.FormValue(dm.Session_Expires_scrn)
		item.Brand = r.FormValue(dm.Session_Brand_scrn)
		item.SessionRole = r.FormValue(dm.Session_SessionRole_scrn)
		item.SYSCreated = r.FormValue(dm.Session_SYSCreated_scrn)
		item.SYSCreatedBy = r.FormValue(dm.Session_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.Session_SYSCreatedHost_scrn)
		item.SYSUpdated = r.FormValue(dm.Session_SYSUpdated_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.Session_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.Session_SYSUpdatedHost_scrn)
		item.SYSDeleted = r.FormValue(dm.Session_SYSDeleted_scrn)
		item.SYSDeletedBy = r.FormValue(dm.Session_SYSDeletedBy_scrn)
		item.SYSDeletedHost = r.FormValue(dm.Session_SYSDeletedHost_scrn)
		item.SYSDbVersion = r.FormValue(dm.Session_SYSDbVersion_scrn)
	
	// 
	// Auto generated 16/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return item
}
