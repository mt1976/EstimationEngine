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
// Date & Time		    : 30/12/2022 at 12:18:20
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
//DocType_Publish - Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
func DocType_Publish(mux http.ServeMux) {
	// START
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// 
	mux.HandleFunc(dm.DocType_Path, DocType_Handler)
	mux.HandleFunc(dm.DocType_PathList, DocType_HandlerList)
	mux.HandleFunc(dm.DocType_PathView, DocType_HandlerView)
	mux.HandleFunc(dm.DocType_PathEdit, DocType_HandlerEdit)
	mux.HandleFunc(dm.DocType_PathNew, DocType_HandlerNew)
	mux.HandleFunc(dm.DocType_PathSave, DocType_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.DocType_Title)
    core.Catalog_Add(dm.DocType_Title, dm.DocType_Path, "", dm.DocType_QueryString, "Application")
	// 
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//DocType_HandlerList is the handler for the list page
//Allows Listing of DocType records
//DocType_HandlerList - Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
func DocType_HandlerList(w http.ResponseWriter, r *http.Request) {
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
	// 
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

}


//DocType_HandlerView is the handler used to View a page
//Allows Viewing for an existing DocType record
//DocType_HandlerView - Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
func DocType_HandlerView(w http.ResponseWriter, r *http.Request) {
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
	// 
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//DocType_HandlerEdit is the handler used generate the Edit page
//Allows Editing for an existing DocType record and then allows the user to save the changes
//DocType_HandlerEdit - Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
func DocType_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
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
	// 
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//DocType_HandlerSave is the handler used process the saving of an DocType
//It is called from the Edit and New pages
//DocType_HandlerSave  - Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
func DocType_HandlerSave(w http.ResponseWriter, r *http.Request) {
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
	logs.Servicing(r.URL.Path+r.FormValue("DocTypeID"))

	item := doctype_DataFromRequest(r)
	
	dao.DocType_Store(item,r)	
	http.Redirect(w, r, dm.DocType_Redirect, http.StatusFound)
	// 
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//DocType_HandlerNew is the handler used process the creation of an DocType
//It will create a new DocType and then redirect to the Edit page
//DocType_HandlerNew  - Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
func DocType_HandlerNew(w http.ResponseWriter, r *http.Request) {
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
	// 
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}	



//doctype_PopulatePage Builds/Populates the DocType Page 
func doctype_PopulatePage(rD dm.DocType, pageDetail dm.DocType_Page) dm.DocType_Page {
	// START
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
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
	pageDetail.SYSDbVersion = rD.SYSDbVersion
	pageDetail.Comments = rD.Comments
	
	
	//
	// Automatically generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
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
	pageDetail.SYSDbVersion_props = rD.SYSDbVersion_props
	pageDetail.Comments_props = rD.Comments_props
	
	// 
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	


//doctype_DataFromRequest is used process the content of an HTTP Request and return an instance of an DocType
func doctype_DataFromRequest(r *http.Request) dm.DocType {
	// START
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	var item dm.DocType
	// FIELD SET START
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
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
		item.SYSDbVersion = r.FormValue(dm.DocType_SYSDbVersion_scrn)
		item.Comments = r.FormValue(dm.DocType_Comments_scrn)
	
	// 
	// Auto generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return item
}
