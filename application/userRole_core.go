package application
// ----------------------------------------------------------------
// Automatically generated  "/application/userrole.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : UserRole (userrole)
// Endpoint 	        : UserRole (Id)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 01/12/2022 at 09:40:04
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

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
	logs.Publish("Application", dm.UserRole_Title)
    core.Catalog_Add(dm.UserRole_Title, dm.UserRole_Path, "", dm.UserRole_QueryString, "Application")
}


//UserRole_HandlerList is the handler for the list page
func UserRole_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.UserRole
	noItems, returnList, _ := dao.UserRole_GetList()

	pageDetail := dm.UserRole_PageList{
		Title:            CardTitle(dm.UserRole_Title, core.Action_List),
		PageTitle:        PageTitle(dm.UserRole_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.UserRole_TemplateList, w, r, pageDetail)

}


//UserRole_HandlerView is the handler used to View a page
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

	ExecuteTemplate(dm.UserRole_TemplateView, w, r, pageDetail)

}


//UserRole_HandlerEdit is the handler used generate the Edit page
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
	_, rD, _ := dao.UserRole_GetByID(searchID)
	
	pageDetail := dm.UserRole_Page{
		Title:       CardTitle(dm.UserRole_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.UserRole_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = userrole_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.UserRole_TemplateEdit, w, r, pageDetail)
}


//UserRole_HandlerSave is the handler used process the saving of an UserRole
func UserRole_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Id"))

	var item dm.UserRole
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
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
	
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.UserRole_Store(item,r)	
	http.Redirect(w, r, dm.UserRole_Redirect, http.StatusFound)
}


//UserRole_HandlerNew is the handler used process the creation of an UserRole
func UserRole_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.UserRole_New()

	pageDetail := dm.UserRole_Page{
		Title:       CardTitle(dm.UserRole_Title, core.Action_New),
		PageTitle:   PageTitle(dm.UserRole_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = userrole_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.UserRole_TemplateNew, w, r, pageDetail)

}	


//UserRole_HandlerDelete is the handler used process the deletion of an UserRole
func UserRole_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.UserRole_QueryString)

	dao.UserRole_Delete(searchID)	

	http.Redirect(w, r, dm.UserRole_Redirect, http.StatusFound)
}


// Builds/Popuplates the UserRole Page 
func userrole_PopulatePage(rD dm.UserRole, pageDetail dm.UserRole_Page) dm.UserRole_Page {
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
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
	
	
	//
	// Automatically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
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
	
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	