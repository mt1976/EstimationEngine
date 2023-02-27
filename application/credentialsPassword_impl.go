package application

import (
	"net/http"

	core "github.com/mt1976/ebEstimates/core"
	dao "github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

func CredentialsPassword_Publish_Impl(mux http.ServeMux) {
	mux.HandleFunc(dm.CredentialsPassword_PathChange, CredentialsPassword_HandlerChange)
	mux.HandleFunc(dm.CredentialsPassword_PathStore, CredentialsPassword_HandlerStore)
	logs.Publish("Implementaion", "CredentialsPassword")
}

// CredentialsPassword_HandlerEdit is the handler used generate the Edit page
// Allows Editing for an existing CredentialsPassword record and then allows the user to save the changes
// CredentialsPassword_HandlerEdit - Auto generated 24/02/2023 by matttownsend (Matt Townsend) on silicon.local
func CredentialsPassword_HandlerChange(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 24/02/2023 by matttownsend (Matt Townsend) on silicon.local
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

	//searchID := core.GetURLparam(r, dm.CredentialsPassword_QueryString)
	userID := Session_GetUserUUID(r)
	//userName := Session_GetUserName(r)

	searchID := core.GetURLparam(r, dm.CredentialsPassword_QueryString)

	//fmt.Printf("userID: %v\n", userID)
	//fmt.Printf("userName: %v\n", userName)

	action := core.GetURLparam(r, core.ContextState)

	var rD dm.CredentialsPassword
	if action == core.ContextState_ERROR {
		//	logs.Information("CredentialsPassword_HandlerEdit", "Error State")
		//	logs.Information("CredentialsPassword_HandlerEdit", "Error State - searchID: "+searchID)
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.CredentialsPassword)
	} else {
		//	logs.Information("CredentialsPassword_HandlerEdit", "Normal State")
		rD.ID = core.GetUUID()
		rD.UserName = userID
		rD.PasswordOld = ""
		rD.PasswordNew = ""
		rD.PasswordConfirm = ""
	}

	pageDetail := dm.CredentialsPassword_Page{
		Title:     CardTitle(dm.CredentialsPassword_Title, core.Action_Edit),
		PageTitle: PageTitle(dm.CredentialsPassword_Title, core.Action_Edit),
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = credentialspassword_PopulatePage(rD, pageDetail)

	ExecuteTemplate(dm.CredentialsPassword_TemplateEdit, w, r, pageDetail)
}

// CredentialsPassword_HandlerSave is the handler used process the saving of an CredentialsPassword
// It is called from the Edit and New pages
// CredentialsPassword_HandlerSave  - Auto generated 24/02/2023 by matttownsend (Matt Townsend) on silicon.local
func CredentialsPassword_HandlerStore(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 24/02/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("ID")
	logs.Servicing(r.URL.Path + itemID)

	item := credentialspassword_DataFromRequest(r)

	item, errStore := dao.CredentialsPassword_Store(item, r)
	//logs.Information(dm.CredentialsPassword_Name, errStore.Error())
	//logs.Information(dm.CredentialsPassword_Name, "Redirecting to "+dm.CredentialsPassword_Redirect)
	if errStore == nil {
		core.SessionManager.Clear(r.Context())
		ExecuteTemplate(dm.CredentialsPassword_TemplateSuccess, w, r, item)
		//http.Redirect(w, r, dm.CredentialsPassword_Redirect, http.StatusFound)
	} else {
		logs.Warning(dm.CredentialsPassword_Name + errStore.Error())
		//http.Redirect(w, r, r.Referer(), http.StatusFound)
		ExecuteRedirect(r.Referer(), w, r, dm.CredentialsPassword_QueryString, itemID, item)
	}
	//
	// Auto generated 24/02/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}
