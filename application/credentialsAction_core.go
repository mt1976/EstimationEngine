package application
// ----------------------------------------------------------------
// Automatically generated  "/application/credentialsaction.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CredentialsAction (credentialsaction)
// Endpoint 	        : CredentialsAction (ID)
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





//CredentialsAction_Publish annouces the endpoints available for this object
func CredentialsAction_Publish(mux http.ServeMux) {
	//No API
	//Cannot List via GUI
	mux.HandleFunc(dm.CredentialsAction_PathView, CredentialsAction_HandlerView)
	mux.HandleFunc(dm.CredentialsAction_PathEdit, CredentialsAction_HandlerEdit)
	mux.HandleFunc(dm.CredentialsAction_PathNew, CredentialsAction_HandlerNew)
	mux.HandleFunc(dm.CredentialsAction_PathSave, CredentialsAction_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.CredentialsAction_Title)
    //No API
}



//CredentialsAction_HandlerView is the handler used to View a page
func CredentialsAction_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CredentialsAction_QueryString)
	_, rD, _ := dao.CredentialsAction_GetByID(searchID)

	pageDetail := dm.CredentialsAction_Page{
		Title:       CardTitle(dm.CredentialsAction_Title, core.Action_View),
		PageTitle:   PageTitle(dm.CredentialsAction_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = credentialsaction_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.CredentialsAction_TemplateView, w, r, pageDetail)

}


//CredentialsAction_HandlerEdit is the handler used generate the Edit page
func CredentialsAction_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CredentialsAction_QueryString)
	_, rD, _ := dao.CredentialsAction_GetByID(searchID)
	
	pageDetail := dm.CredentialsAction_Page{
		Title:       CardTitle(dm.CredentialsAction_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.CredentialsAction_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = credentialsaction_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.CredentialsAction_TemplateEdit, w, r, pageDetail)
}


//CredentialsAction_HandlerSave is the handler used process the saving of an CredentialsAction
func CredentialsAction_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("ID"))

	var item dm.CredentialsAction
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.ID = r.FormValue(dm.CredentialsAction_ID_scrn)
		item.User = r.FormValue(dm.CredentialsAction_User_scrn)
		item.Action = r.FormValue(dm.CredentialsAction_Action_scrn)
		item.Notes = r.FormValue(dm.CredentialsAction_Notes_scrn)
	
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.CredentialsAction_Store(item,r)	
	http.Redirect(w, r, dm.CredentialsAction_Redirect, http.StatusFound)
}


//CredentialsAction_HandlerNew is the handler used process the creation of an CredentialsAction
func CredentialsAction_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.CredentialsAction_New()

	pageDetail := dm.CredentialsAction_Page{
		Title:       CardTitle(dm.CredentialsAction_Title, core.Action_New),
		PageTitle:   PageTitle(dm.CredentialsAction_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = credentialsaction_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.CredentialsAction_TemplateNew, w, r, pageDetail)

}	



// Builds/Popuplates the CredentialsAction Page 
func credentialsaction_PopulatePage(rD dm.CredentialsAction, pageDetail dm.CredentialsAction_Page) dm.CredentialsAction_Page {
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.ID = rD.ID
	pageDetail.User = rD.User
	pageDetail.Action = rD.Action
	pageDetail.Notes = rD.Notes
	
	
	//
	// Automatically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	pageDetail.User_lookup = dao.Credentials_GetLookup()
	
	
	
	
	pageDetail.Action_lookup = dao.StubLists_Get("credentialStates")
	
	
	
	
	pageDetail.ID_props = rD.ID_props
	pageDetail.User_props = rD.User_props
	pageDetail.Action_props = rD.Action_props
	pageDetail.Notes_props = rD.Notes_props
	
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	