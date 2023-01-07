package application
// ----------------------------------------------------------------
// Automatically generated  "/application/schedule.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Schedule (schedule)
// Endpoint 	        : Schedule (Schedule)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 07/01/2023 at 23:01:31
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)

//Schedule_Publish annouces the endpoints available for this object
//Schedule_Publish - Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
func Schedule_Publish(mux http.ServeMux) {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	mux.HandleFunc(dm.Schedule_Path, Schedule_Handler)
	mux.HandleFunc(dm.Schedule_PathList, Schedule_HandlerList)
	mux.HandleFunc(dm.Schedule_PathView, Schedule_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	mux.HandleFunc(dm.Schedule_PathSave, Schedule_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Schedule_Title)
    core.Catalog_Add(dm.Schedule_Title, dm.Schedule_Path, "", dm.Schedule_QueryString, "Application")
	// 
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//Schedule_HandlerList is the handler for the list page
//Allows Listing of Schedule records
//Schedule_HandlerList - Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
func Schedule_HandlerList(w http.ResponseWriter, r *http.Request) {
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

	var returnList []dm.Schedule
	noItems, returnList, _ := dao.Schedule_GetList()

	pageDetail := dm.Schedule_PageList{
		Title:            CardTitle(dm.Schedule_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Schedule_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Schedule_TemplateList, w, r, pageDetail)
	// 
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END

}


//Schedule_HandlerView is the handler used to View a page
//Allows Viewing for an existing Schedule record
//Schedule_HandlerView - Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
func Schedule_HandlerView(w http.ResponseWriter, r *http.Request) {
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

	searchID := core.GetURLparam(r, dm.Schedule_QueryString)
	_, rD, _ := dao.Schedule_GetByID(searchID)

	pageDetail := dm.Schedule_Page{
		Title:       CardTitle(dm.Schedule_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Schedule_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = schedule_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Schedule_TemplateView, w, r, pageDetail)
	// 
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}



//Schedule_HandlerSave is the handler used process the saving of an Schedule
//It is called from the Edit and New pages
//Schedule_HandlerSave  - Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
func Schedule_HandlerSave(w http.ResponseWriter, r *http.Request) {
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
	logs.Servicing(r.URL.Path+r.FormValue("Id"))

	item := schedule_DataFromRequest(r)
	
	dao.Schedule_Store(item,r)	
	http.Redirect(w, r, dm.Schedule_Redirect, http.StatusFound)
	// 
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}




//schedule_PopulatePage Builds/Populates the Schedule Page 
func schedule_PopulatePage(rD dm.Schedule, pageDetail dm.Schedule_Page) dm.Schedule_Page {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.Id = rD.Id
	pageDetail.Name = rD.Name
	pageDetail.Description = rD.Description
	pageDetail.Schedule = rD.Schedule
	pageDetail.Started = rD.Started
	pageDetail.Lastrun = rD.Lastrun
	pageDetail.Message = rD.Message
	pageDetail.Type = rD.Type
	pageDetail.Human = rD.Human
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
	pageDetail.Id_props = rD.Id_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.Description_props = rD.Description_props
	pageDetail.Schedule_props = rD.Schedule_props
	pageDetail.Started_props = rD.Started_props
	pageDetail.Lastrun_props = rD.Lastrun_props
	pageDetail.Message_props = rD.Message_props
	pageDetail.Type_props = rD.Type_props
	pageDetail.Human_props = rD.Human_props
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


//schedule_DataFromRequest is used process the content of an HTTP Request and return an instance of an Schedule
func schedule_DataFromRequest(r *http.Request) dm.Schedule {
	// START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	var item dm.Schedule
	// FIELD SET START
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.Schedule_SYSId_scrn)
		item.Id = r.FormValue(dm.Schedule_Id_scrn)
		item.Name = r.FormValue(dm.Schedule_Name_scrn)
		item.Description = r.FormValue(dm.Schedule_Description_scrn)
		item.Schedule = r.FormValue(dm.Schedule_Schedule_scrn)
		item.Started = r.FormValue(dm.Schedule_Started_scrn)
		item.Lastrun = r.FormValue(dm.Schedule_Lastrun_scrn)
		item.Message = r.FormValue(dm.Schedule_Message_scrn)
		item.Type = r.FormValue(dm.Schedule_Type_scrn)
		item.Human = r.FormValue(dm.Schedule_Human_scrn)
		item.SYSCreated = r.FormValue(dm.Schedule_SYSCreated_scrn)
		item.SYSCreatedBy = r.FormValue(dm.Schedule_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.Schedule_SYSCreatedHost_scrn)
		item.SYSUpdated = r.FormValue(dm.Schedule_SYSUpdated_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.Schedule_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.Schedule_SYSUpdatedHost_scrn)
		item.SYSDeleted = r.FormValue(dm.Schedule_SYSDeleted_scrn)
		item.SYSDeletedBy = r.FormValue(dm.Schedule_SYSDeletedBy_scrn)
		item.SYSDeletedHost = r.FormValue(dm.Schedule_SYSDeletedHost_scrn)
		item.SYSDbVersion = r.FormValue(dm.Schedule_SYSDbVersion_scrn)
	
	// 
	// Auto generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return item
}
