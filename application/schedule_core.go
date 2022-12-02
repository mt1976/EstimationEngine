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
// Date & Time		    : 01/12/2022 at 09:40:03
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
func Schedule_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Schedule_Path, Schedule_Handler)
	mux.HandleFunc(dm.Schedule_PathList, Schedule_HandlerList)
	mux.HandleFunc(dm.Schedule_PathView, Schedule_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	mux.HandleFunc(dm.Schedule_PathSave, Schedule_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Schedule_Title)
    core.Catalog_Add(dm.Schedule_Title, dm.Schedule_Path, "", dm.Schedule_QueryString, "Application")
}


//Schedule_HandlerList is the handler for the list page
func Schedule_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

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

}


//Schedule_HandlerView is the handler used to View a page
func Schedule_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
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

}



//Schedule_HandlerSave is the handler used process the saving of an Schedule
func Schedule_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Id"))

	var item dm.Schedule
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.Schedule_SYSId_scrn)
		item.Id = r.FormValue(dm.Schedule_Id_scrn)
		item.Name = r.FormValue(dm.Schedule_Name_scrn)
		item.Description = r.FormValue(dm.Schedule_Description_scrn)
		item.Schedule = r.FormValue(dm.Schedule_Schedule_scrn)
		item.Started = r.FormValue(dm.Schedule_Started_scrn)
		item.Lastrun = r.FormValue(dm.Schedule_Lastrun_scrn)
		item.Message = r.FormValue(dm.Schedule_Message_scrn)
		item.SYSCreated = r.FormValue(dm.Schedule_SYSCreated_scrn)
		item.SYSWho = r.FormValue(dm.Schedule_SYSWho_scrn)
		item.SYSHost = r.FormValue(dm.Schedule_SYSHost_scrn)
		item.SYSUpdated = r.FormValue(dm.Schedule_SYSUpdated_scrn)
		item.Type = r.FormValue(dm.Schedule_Type_scrn)
		item.SYSCreatedBy = r.FormValue(dm.Schedule_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.Schedule_SYSCreatedHost_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.Schedule_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.Schedule_SYSUpdatedHost_scrn)
		item.Human = r.FormValue(dm.Schedule_Human_scrn)
	
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Schedule_Store(item,r)	
	http.Redirect(w, r, dm.Schedule_Redirect, http.StatusFound)
}




// Builds/Popuplates the Schedule Page 
func schedule_PopulatePage(rD dm.Schedule, pageDetail dm.Schedule_Page) dm.Schedule_Page {
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.Id = rD.Id
	pageDetail.Name = rD.Name
	pageDetail.Description = rD.Description
	pageDetail.Schedule = rD.Schedule
	pageDetail.Started = rD.Started
	pageDetail.Lastrun = rD.Lastrun
	pageDetail.Message = rD.Message
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSWho = rD.SYSWho
	pageDetail.SYSHost = rD.SYSHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.Type = rD.Type
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.Human = rD.Human
	
	
	//
	// Automatically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.Id_props = rD.Id_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.Description_props = rD.Description_props
	pageDetail.Schedule_props = rD.Schedule_props
	pageDetail.Started_props = rD.Started_props
	pageDetail.Lastrun_props = rD.Lastrun_props
	pageDetail.Message_props = rD.Message_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSWho_props = rD.SYSWho_props
	pageDetail.SYSHost_props = rD.SYSHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.Type_props = rD.Type_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.Human_props = rD.Human_props
	
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	