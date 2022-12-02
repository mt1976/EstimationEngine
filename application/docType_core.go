package application
// ----------------------------------------------------------------
// Automatically generated  "/application/doctype.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DocType (doctype)
// Endpoint 	        : DocType (DocTypeID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 01/12/2022 at 09:40:00
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

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
	//Cannot Delete via GUI
	logs.Publish("Application", dm.DocType_Title)
    core.Catalog_Add(dm.DocType_Title, dm.DocType_Path, "", dm.DocType_QueryString, "Application")
}


//DocType_HandlerList is the handler for the list page
func DocType_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.DocType
	noItems, returnList, _ := dao.DocType_GetList()

	pageDetail := dm.DocType_PageList{
		Title:            CardTitle(dm.DocType_Title, core.Action_List),
		PageTitle:        PageTitle(dm.DocType_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.DocType_TemplateList, w, r, pageDetail)

}


//DocType_HandlerView is the handler used to View a page
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

	ExecuteTemplate(dm.DocType_TemplateView, w, r, pageDetail)

}


//DocType_HandlerEdit is the handler used generate the Edit page
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
	_, rD, _ := dao.DocType_GetByID(searchID)
	
	pageDetail := dm.DocType_Page{
		Title:       CardTitle(dm.DocType_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.DocType_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = doctype_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.DocType_TemplateEdit, w, r, pageDetail)
}


//DocType_HandlerSave is the handler used process the saving of an DocType
func DocType_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("DocTypeID"))

	var item dm.DocType
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
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
	
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.DocType_Store(item,r)	
	http.Redirect(w, r, dm.DocType_Redirect, http.StatusFound)
}


//DocType_HandlerNew is the handler used process the creation of an DocType
func DocType_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.DocType_New()

	pageDetail := dm.DocType_Page{
		Title:       CardTitle(dm.DocType_Title, core.Action_New),
		PageTitle:   PageTitle(dm.DocType_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = doctype_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.DocType_TemplateNew, w, r, pageDetail)

}	



// Builds/Popuplates the DocType Page 
func doctype_PopulatePage(rD dm.DocType, pageDetail dm.DocType_Page) dm.DocType_Page {
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
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
	
	
	//
	// Automatically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
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
	
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	