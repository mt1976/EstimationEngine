package application
// ----------------------------------------------------------------
// Automatically generated  "/application/translation.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Translation (translation)
// Endpoint 	        : Translation (Message)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 27/11/2022 at 20:46:15
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)





//Translation_Publish annouces the endpoints available for this object
func Translation_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.Translation_PathList, Translation_HandlerList)
	mux.HandleFunc(dm.Translation_PathView, Translation_HandlerView)
	mux.HandleFunc(dm.Translation_PathEdit, Translation_HandlerEdit)
	//Cannot Create via GUI
	mux.HandleFunc(dm.Translation_PathSave, Translation_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Translation_Title)
    //No API
}


//Translation_HandlerList is the handler for the list page
func Translation_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Translation
	noItems, returnList, _ := dao.Translation_GetList()

	pageDetail := dm.Translation_PageList{
		Title:            CardTitle(dm.Translation_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Translation_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Translation_TemplateList, w, r, pageDetail)

}


//Translation_HandlerView is the handler used to View a page
func Translation_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Translation_QueryString)
	_, rD, _ := dao.Translation_GetByID(searchID)

	pageDetail := dm.Translation_Page{
		Title:       CardTitle(dm.Translation_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Translation_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = translation_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Translation_TemplateView, w, r, pageDetail)

}


//Translation_HandlerEdit is the handler used generate the Edit page
func Translation_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Translation_QueryString)
	_, rD, _ := dao.Translation_GetByID(searchID)
	
	pageDetail := dm.Translation_Page{
		Title:       CardTitle(dm.Translation_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Translation_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = translation_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Translation_TemplateEdit, w, r, pageDetail)
}


//Translation_HandlerSave is the handler used process the saving of an Translation
func Translation_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Id"))

	var item dm.Translation
	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.Translation_SYSId_scrn)
		item.Id = r.FormValue(dm.Translation_Id_scrn)
		item.Class = r.FormValue(dm.Translation_Class_scrn)
		item.Message = r.FormValue(dm.Translation_Message_scrn)
		item.Translation = r.FormValue(dm.Translation_Translation_scrn)
		item.SYSCreated = r.FormValue(dm.Translation_SYSCreated_scrn)
		item.SYSUpdated = r.FormValue(dm.Translation_SYSUpdated_scrn)
		item.SYSCreatedBy = r.FormValue(dm.Translation_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.Translation_SYSCreatedHost_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.Translation_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.Translation_SYSUpdatedHost_scrn)
		item.SYSDeleted = r.FormValue(dm.Translation_SYSDeleted_scrn)
		item.SYSDeletedBy = r.FormValue(dm.Translation_SYSDeletedBy_scrn)
		item.SYSDeletedHost = r.FormValue(dm.Translation_SYSDeletedHost_scrn)
	
	// 
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Translation_Store(item,r)	
	http.Redirect(w, r, dm.Translation_Redirect, http.StatusFound)
}




// Builds/Popuplates the Translation Page 
func translation_PopulatePage(rD dm.Translation, pageDetail dm.Translation_Page) dm.Translation_Page {
	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.Id = rD.Id
	pageDetail.Class = rD.Class
	pageDetail.Message = rD.Message
	pageDetail.Translation = rD.Translation
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.SYSDeleted = rD.SYSDeleted
	pageDetail.SYSDeletedBy = rD.SYSDeletedBy
	pageDetail.SYSDeletedHost = rD.SYSDeletedHost
	
	
	//
	// Automatically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.Id_props = rD.Id_props
	pageDetail.Class_props = rD.Class_props
	pageDetail.Message_props = rD.Message_props
	pageDetail.Translation_props = rD.Translation_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
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