package application
// ----------------------------------------------------------------
// Automatically generated  "/application/profile.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Profile (profile)
// Endpoint 	        : Profile (ProfileID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 01/12/2022 at 09:40:02
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)





//Profile_Publish annouces the endpoints available for this object
func Profile_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Profile_Path, Profile_Handler)
	mux.HandleFunc(dm.Profile_PathList, Profile_HandlerList)
	mux.HandleFunc(dm.Profile_PathView, Profile_HandlerView)
	mux.HandleFunc(dm.Profile_PathEdit, Profile_HandlerEdit)
	mux.HandleFunc(dm.Profile_PathNew, Profile_HandlerNew)
	mux.HandleFunc(dm.Profile_PathSave, Profile_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Profile_Title)
    core.Catalog_Add(dm.Profile_Title, dm.Profile_Path, "", dm.Profile_QueryString, "Application")
}


//Profile_HandlerList is the handler for the list page
func Profile_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Profile
	noItems, returnList, _ := dao.Profile_GetList()

	pageDetail := dm.Profile_PageList{
		Title:            CardTitle(dm.Profile_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Profile_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Profile_TemplateList, w, r, pageDetail)

}


//Profile_HandlerView is the handler used to View a page
func Profile_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Profile_QueryString)
	_, rD, _ := dao.Profile_GetByID(searchID)

	pageDetail := dm.Profile_Page{
		Title:       CardTitle(dm.Profile_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Profile_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = profile_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Profile_TemplateView, w, r, pageDetail)

}


//Profile_HandlerEdit is the handler used generate the Edit page
func Profile_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Profile_QueryString)
	_, rD, _ := dao.Profile_GetByID(searchID)
	
	pageDetail := dm.Profile_Page{
		Title:       CardTitle(dm.Profile_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Profile_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = profile_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Profile_TemplateEdit, w, r, pageDetail)
}


//Profile_HandlerSave is the handler used process the saving of an Profile
func Profile_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("ProfileID"))

	var item dm.Profile
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.Profile_SYSId_scrn)
		item.ProfileID = r.FormValue(dm.Profile_ProfileID_scrn)
		item.Code = r.FormValue(dm.Profile_Code_scrn)
		item.Name = r.FormValue(dm.Profile_Name_scrn)
		item.StartDate = r.FormValue(dm.Profile_StartDate_scrn)
		item.EndDate = r.FormValue(dm.Profile_EndDate_scrn)
		item.DefaultReleases = r.FormValue(dm.Profile_DefaultReleases_scrn)
		item.DefaultReleaseHours = r.FormValue(dm.Profile_DefaultReleaseHours_scrn)
		item.BlendedRate = r.FormValue(dm.Profile_BlendedRate_scrn)
		item.Rounding = r.FormValue(dm.Profile_Rounding_scrn)
		item.HoursPerDay = r.FormValue(dm.Profile_HoursPerDay_scrn)
		item.REQPerc = r.FormValue(dm.Profile_REQPerc_scrn)
		item.ANAPerc = r.FormValue(dm.Profile_ANAPerc_scrn)
		item.DOCPerc = r.FormValue(dm.Profile_DOCPerc_scrn)
		item.PMPerc = r.FormValue(dm.Profile_PMPerc_scrn)
		item.UATPerc = r.FormValue(dm.Profile_UATPerc_scrn)
		item.GTMPerc = r.FormValue(dm.Profile_GTMPerc_scrn)
		item.SupportUplift = r.FormValue(dm.Profile_SupportUplift_scrn)
		item.ContigencyPerc = r.FormValue(dm.Profile_ContigencyPerc_scrn)
		item.SYSCreated = r.FormValue(dm.Profile_SYSCreated_scrn)
		item.SYSCreatedBy = r.FormValue(dm.Profile_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.Profile_SYSCreatedHost_scrn)
		item.SYSUpdated = r.FormValue(dm.Profile_SYSUpdated_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.Profile_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.Profile_SYSUpdatedHost_scrn)
		item.SYSDeleted = r.FormValue(dm.Profile_SYSDeleted_scrn)
		item.SYSDeletedBy = r.FormValue(dm.Profile_SYSDeletedBy_scrn)
		item.SYSDeletedHost = r.FormValue(dm.Profile_SYSDeletedHost_scrn)
		item.SYSActivity = r.FormValue(dm.Profile_SYSActivity_scrn)
		item.Notes = r.FormValue(dm.Profile_Notes_scrn)
	
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Profile_Store(item,r)	
	http.Redirect(w, r, dm.Profile_Redirect, http.StatusFound)
}


//Profile_HandlerNew is the handler used process the creation of an Profile
func Profile_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.Profile_New()

	pageDetail := dm.Profile_Page{
		Title:       CardTitle(dm.Profile_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Profile_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = profile_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Profile_TemplateNew, w, r, pageDetail)

}	



// Builds/Popuplates the Profile Page 
func profile_PopulatePage(rD dm.Profile, pageDetail dm.Profile_Page) dm.Profile_Page {
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.ProfileID = rD.ProfileID
	pageDetail.Code = rD.Code
	pageDetail.Name = rD.Name
	pageDetail.StartDate = rD.StartDate
	pageDetail.EndDate = rD.EndDate
	pageDetail.DefaultReleases = rD.DefaultReleases
	pageDetail.DefaultReleaseHours = rD.DefaultReleaseHours
	pageDetail.BlendedRate = rD.BlendedRate
	pageDetail.Rounding = rD.Rounding
	pageDetail.HoursPerDay = rD.HoursPerDay
	pageDetail.REQPerc = rD.REQPerc
	pageDetail.ANAPerc = rD.ANAPerc
	pageDetail.DOCPerc = rD.DOCPerc
	pageDetail.PMPerc = rD.PMPerc
	pageDetail.UATPerc = rD.UATPerc
	pageDetail.GTMPerc = rD.GTMPerc
	pageDetail.SupportUplift = rD.SupportUplift
	pageDetail.ContigencyPerc = rD.ContigencyPerc
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.SYSDeleted = rD.SYSDeleted
	pageDetail.SYSDeletedBy = rD.SYSDeletedBy
	pageDetail.SYSDeletedHost = rD.SYSDeletedHost
	pageDetail.SYSActivity = rD.SYSActivity
	pageDetail.Notes = rD.Notes
	
	
	//
	// Automatically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.ProfileID_props = rD.ProfileID_props
	pageDetail.Code_props = rD.Code_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.StartDate_props = rD.StartDate_props
	pageDetail.EndDate_props = rD.EndDate_props
	pageDetail.DefaultReleases_props = rD.DefaultReleases_props
	pageDetail.DefaultReleaseHours_props = rD.DefaultReleaseHours_props
	pageDetail.BlendedRate_props = rD.BlendedRate_props
	pageDetail.Rounding_props = rD.Rounding_props
	pageDetail.HoursPerDay_props = rD.HoursPerDay_props
	pageDetail.REQPerc_props = rD.REQPerc_props
	pageDetail.ANAPerc_props = rD.ANAPerc_props
	pageDetail.DOCPerc_props = rD.DOCPerc_props
	pageDetail.PMPerc_props = rD.PMPerc_props
	pageDetail.UATPerc_props = rD.UATPerc_props
	pageDetail.GTMPerc_props = rD.GTMPerc_props
	pageDetail.SupportUplift_props = rD.SupportUplift_props
	pageDetail.ContigencyPerc_props = rD.ContigencyPerc_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.SYSDeleted_props = rD.SYSDeleted_props
	pageDetail.SYSDeletedBy_props = rD.SYSDeletedBy_props
	pageDetail.SYSDeletedHost_props = rD.SYSDeletedHost_props
	pageDetail.SYSActivity_props = rD.SYSActivity_props
	pageDetail.Notes_props = rD.Notes_props
	
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	