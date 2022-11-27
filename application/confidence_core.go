package application
// ----------------------------------------------------------------
// Automatically generated  "/application/confidence.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Confidence (confidence)
// Endpoint 	        : Confidence (ConfidenceID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 27/11/2022 at 20:46:12
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)





//Confidence_Publish annouces the endpoints available for this object
func Confidence_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Confidence_Path, Confidence_Handler)
	mux.HandleFunc(dm.Confidence_PathList, Confidence_HandlerList)
	mux.HandleFunc(dm.Confidence_PathView, Confidence_HandlerView)
	mux.HandleFunc(dm.Confidence_PathEdit, Confidence_HandlerEdit)
	mux.HandleFunc(dm.Confidence_PathNew, Confidence_HandlerNew)
	mux.HandleFunc(dm.Confidence_PathSave, Confidence_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Confidence_Title)
    core.Catalog_Add(dm.Confidence_Title, dm.Confidence_Path, "", dm.Confidence_QueryString, "Application")
}


//Confidence_HandlerList is the handler for the list page
func Confidence_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Confidence
	noItems, returnList, _ := dao.Confidence_GetList()

	pageDetail := dm.Confidence_PageList{
		Title:            CardTitle(dm.Confidence_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Confidence_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Confidence_TemplateList, w, r, pageDetail)

}


//Confidence_HandlerView is the handler used to View a page
func Confidence_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Confidence_QueryString)
	_, rD, _ := dao.Confidence_GetByID(searchID)

	pageDetail := dm.Confidence_Page{
		Title:       CardTitle(dm.Confidence_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Confidence_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = confidence_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Confidence_TemplateView, w, r, pageDetail)

}


//Confidence_HandlerEdit is the handler used generate the Edit page
func Confidence_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Confidence_QueryString)
	_, rD, _ := dao.Confidence_GetByID(searchID)
	
	pageDetail := dm.Confidence_Page{
		Title:       CardTitle(dm.Confidence_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Confidence_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = confidence_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Confidence_TemplateEdit, w, r, pageDetail)
}


//Confidence_HandlerSave is the handler used process the saving of an Confidence
func Confidence_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("ConfidenceID"))

	var item dm.Confidence
	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.Confidence_SYSId_scrn)
		item.ConfidenceID = r.FormValue(dm.Confidence_ConfidenceID_scrn)
		item.Code = r.FormValue(dm.Confidence_Code_scrn)
		item.Name = r.FormValue(dm.Confidence_Name_scrn)
		item.Perc = r.FormValue(dm.Confidence_Perc_scrn)
		item.SYSCreated = r.FormValue(dm.Confidence_SYSCreated_scrn)
		item.SYSCreatedBy = r.FormValue(dm.Confidence_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.Confidence_SYSCreatedHost_scrn)
		item.SYSUpdated = r.FormValue(dm.Confidence_SYSUpdated_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.Confidence_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.Confidence_SYSUpdatedHost_scrn)
		item.SYSDeleted = r.FormValue(dm.Confidence_SYSDeleted_scrn)
		item.SYSDeletedBy = r.FormValue(dm.Confidence_SYSDeletedBy_scrn)
		item.SYSDeletedHost = r.FormValue(dm.Confidence_SYSDeletedHost_scrn)
	
	// 
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Confidence_Store(item,r)	
	http.Redirect(w, r, dm.Confidence_Redirect, http.StatusFound)
}


//Confidence_HandlerNew is the handler used process the creation of an Confidence
func Confidence_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.Confidence_New()

	pageDetail := dm.Confidence_Page{
		Title:       CardTitle(dm.Confidence_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Confidence_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = confidence_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Confidence_TemplateNew, w, r, pageDetail)

}	



// Builds/Popuplates the Confidence Page 
func confidence_PopulatePage(rD dm.Confidence, pageDetail dm.Confidence_Page) dm.Confidence_Page {
	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.ConfidenceID = rD.ConfidenceID
	pageDetail.Code = rD.Code
	pageDetail.Name = rD.Name
	pageDetail.Perc = rD.Perc
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
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.ConfidenceID_props = rD.ConfidenceID_props
	pageDetail.Code_props = rD.Code_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.Perc_props = rD.Perc_props
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