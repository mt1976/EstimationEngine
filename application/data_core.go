package application
// ----------------------------------------------------------------
// Automatically generated  "/application/data.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Data (data)
// Endpoint 	        : Data (DataID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 07/01/2023 at 23:01:28
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)

//Data_Publish annouces the endpoints available for this object
//Data_Publish - Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
func Data_Publish(mux http.ServeMux) {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	//No API
	mux.HandleFunc(dm.Data_PathList, Data_HandlerList)
	mux.HandleFunc(dm.Data_PathView, Data_HandlerView)
	mux.HandleFunc(dm.Data_PathEdit, Data_HandlerEdit)
	mux.HandleFunc(dm.Data_PathNew, Data_HandlerNew)
	mux.HandleFunc(dm.Data_PathSave, Data_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Data_Title)
    //No API
	// 
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//Data_HandlerList is the handler for the list page
//Allows Listing of Data records
//Data_HandlerList - Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
func Data_HandlerList(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
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

	var returnList []dm.Data
	noItems, returnList, _ := dao.Data_GetList()

	pageDetail := dm.Data_PageList{
		Title:            CardTitle(dm.Data_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Data_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Data_TemplateList, w, r, pageDetail)
	// 
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END

}


//Data_HandlerView is the handler used to View a page
//Allows Viewing for an existing Data record
//Data_HandlerView - Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
func Data_HandlerView(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
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

	searchID := core.GetURLparam(r, dm.Data_QueryString)
	_, rD, _ := dao.Data_GetByID(searchID)

	pageDetail := dm.Data_Page{
		Title:       CardTitle(dm.Data_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Data_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = data_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Data_TemplateView, w, r, pageDetail)
	// 
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//Data_HandlerEdit is the handler used generate the Edit page
//Allows Editing for an existing Data record and then allows the user to save the changes
//Data_HandlerEdit - Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
func Data_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
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

	searchID := core.GetURLparam(r, dm.Data_QueryString)
	_, rD, _ := dao.Data_GetByID(searchID)
	
	pageDetail := dm.Data_Page{
		Title:       CardTitle(dm.Data_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Data_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = data_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Data_TemplateEdit, w, r, pageDetail)
	// 
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//Data_HandlerSave is the handler used process the saving of an Data
//It is called from the Edit and New pages
//Data_HandlerSave  - Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
func Data_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("DataID"))

	item := data_DataFromRequest(r)
	
	dao.Data_Store(item,r)	
	http.Redirect(w, r, dm.Data_Redirect, http.StatusFound)
	// 
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//Data_HandlerNew is the handler used process the creation of an Data
//It will create a new Data and then redirect to the Edit page
//Data_HandlerNew  - Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
func Data_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	_, _, rD, _ := dao.Data_New()

	pageDetail := dm.Data_Page{
		Title:       CardTitle(dm.Data_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Data_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = data_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Data_TemplateNew, w, r, pageDetail)
	// 
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}	



//data_PopulatePage Builds/Populates the Data Page 
func data_PopulatePage(rD dm.Data, pageDetail dm.Data_Page) dm.Data_Page {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.DataID = rD.DataID
	pageDetail.Class = rD.Class
	pageDetail.Field = rD.Field
	pageDetail.Value = rD.Value
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
	// Automatically generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.DataID_props = rD.DataID_props
	pageDetail.Class_props = rD.Class_props
	pageDetail.Field_props = rD.Field_props
	pageDetail.Value_props = rD.Value_props
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
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	


//data_DataFromRequest is used process the content of an HTTP Request and return an instance of an Data
func data_DataFromRequest(r *http.Request) dm.Data {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	var item dm.Data
	// FIELD SET START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.Data_SYSId_scrn)
		item.DataID = r.FormValue(dm.Data_DataID_scrn)
		item.Class = r.FormValue(dm.Data_Class_scrn)
		item.Field = r.FormValue(dm.Data_Field_scrn)
		item.Value = r.FormValue(dm.Data_Value_scrn)
		item.SYSCreated = r.FormValue(dm.Data_SYSCreated_scrn)
		item.SYSCreatedBy = r.FormValue(dm.Data_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.Data_SYSCreatedHost_scrn)
		item.SYSUpdated = r.FormValue(dm.Data_SYSUpdated_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.Data_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.Data_SYSUpdatedHost_scrn)
		item.SYSDeleted = r.FormValue(dm.Data_SYSDeleted_scrn)
		item.SYSDeletedBy = r.FormValue(dm.Data_SYSDeletedBy_scrn)
		item.SYSDeletedHost = r.FormValue(dm.Data_SYSDeletedHost_scrn)
		item.SYSDbVersion = r.FormValue(dm.Data_SYSDbVersion_scrn)
	
	// 
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return item
}
