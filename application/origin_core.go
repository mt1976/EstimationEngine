package application
// ----------------------------------------------------------------
// Automatically generated  "/application/origin.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Origin (origin)
// Endpoint 	        : Origin (OriginID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 27/11/2022 at 20:46:14
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)





//Origin_Publish annouces the endpoints available for this object
func Origin_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Origin_Path, Origin_Handler)
	mux.HandleFunc(dm.Origin_PathList, Origin_HandlerList)
	mux.HandleFunc(dm.Origin_PathView, Origin_HandlerView)
	mux.HandleFunc(dm.Origin_PathEdit, Origin_HandlerEdit)
	mux.HandleFunc(dm.Origin_PathNew, Origin_HandlerNew)
	mux.HandleFunc(dm.Origin_PathSave, Origin_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Origin_Title)
    core.Catalog_Add(dm.Origin_Title, dm.Origin_Path, "", dm.Origin_QueryString, "Application")
}


//Origin_HandlerList is the handler for the list page
func Origin_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Origin
	noItems, returnList, _ := dao.Origin_GetList()

	pageDetail := dm.Origin_PageList{
		Title:            CardTitle(dm.Origin_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Origin_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Origin_TemplateList, w, r, pageDetail)

}


//Origin_HandlerView is the handler used to View a page
func Origin_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Origin_QueryString)
	_, rD, _ := dao.Origin_GetByID(searchID)

	pageDetail := dm.Origin_Page{
		Title:       CardTitle(dm.Origin_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Origin_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = origin_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Origin_TemplateView, w, r, pageDetail)

}


//Origin_HandlerEdit is the handler used generate the Edit page
func Origin_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Origin_QueryString)
	_, rD, _ := dao.Origin_GetByID(searchID)
	
	pageDetail := dm.Origin_Page{
		Title:       CardTitle(dm.Origin_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Origin_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = origin_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Origin_TemplateEdit, w, r, pageDetail)
}


//Origin_HandlerSave is the handler used process the saving of an Origin
func Origin_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("OriginID"))

	var item dm.Origin
	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.Origin_SYSId_scrn)
		item.OriginID = r.FormValue(dm.Origin_OriginID_scrn)
		item.StateID = r.FormValue(dm.Origin_StateID_scrn)
		item.DocTypeID = r.FormValue(dm.Origin_DocTypeID_scrn)
		item.Code = r.FormValue(dm.Origin_Code_scrn)
		item.FullName = r.FormValue(dm.Origin_FullName_scrn)
		item.Rate = r.FormValue(dm.Origin_Rate_scrn)
		item.Notes = r.FormValue(dm.Origin_Notes_scrn)
		item.StartDate = r.FormValue(dm.Origin_StartDate_scrn)
		item.EndDate = r.FormValue(dm.Origin_EndDate_scrn)
		item.SYSCreated = r.FormValue(dm.Origin_SYSCreated_scrn)
		item.SYSCreatedBy = r.FormValue(dm.Origin_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.Origin_SYSCreatedHost_scrn)
		item.SYSUpdated = r.FormValue(dm.Origin_SYSUpdated_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.Origin_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.Origin_SYSUpdatedHost_scrn)
		item.SYSDeleted = r.FormValue(dm.Origin_SYSDeleted_scrn)
		item.SYSDeletedBy = r.FormValue(dm.Origin_SYSDeletedBy_scrn)
		item.SYSDeletedHost = r.FormValue(dm.Origin_SYSDeletedHost_scrn)
	
	// 
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Origin_Store(item,r)	
	http.Redirect(w, r, dm.Origin_Redirect, http.StatusFound)
}


//Origin_HandlerNew is the handler used process the creation of an Origin
func Origin_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.Origin_New()

	pageDetail := dm.Origin_Page{
		Title:       CardTitle(dm.Origin_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Origin_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = origin_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Origin_TemplateNew, w, r, pageDetail)

}	



// Builds/Popuplates the Origin Page 
func origin_PopulatePage(rD dm.Origin, pageDetail dm.Origin_Page) dm.Origin_Page {
	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.OriginID = rD.OriginID
	pageDetail.StateID = rD.StateID
	pageDetail.DocTypeID = rD.DocTypeID
	pageDetail.Code = rD.Code
	pageDetail.FullName = rD.FullName
	pageDetail.Rate = rD.Rate
	pageDetail.Notes = rD.Notes
	pageDetail.StartDate = rD.StartDate
	pageDetail.EndDate = rD.EndDate
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
	// Automatically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	pageDetail.StateID_lookup = dao.OriginState_GetLookup()
	
	
	
	pageDetail.DocTypeID_lookup = dao.DocType_GetLookup()
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.OriginID_props = rD.OriginID_props
	pageDetail.StateID_props = rD.StateID_props
	pageDetail.DocTypeID_props = rD.DocTypeID_props
	pageDetail.Code_props = rD.Code_props
	pageDetail.FullName_props = rD.FullName_props
	pageDetail.Rate_props = rD.Rate_props
	pageDetail.Notes_props = rD.Notes_props
	pageDetail.StartDate_props = rD.StartDate_props
	pageDetail.EndDate_props = rD.EndDate_props
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
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	