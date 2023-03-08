package application
// ----------------------------------------------------------------
// Automatically generated  "/application/doctype.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DocType (doctype)
// Endpoint 	        : DocType (DocTypeID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 08/03/2023 at 18:42:23
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
//DocType_Publish - Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local
func DocType_Publish(mux http.ServeMux) {
	// START
	// Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	mux.HandleFunc(dm.DocType_Path, DocType_Handler)
	mux.HandleFunc(dm.DocType_PathList, DocType_HandlerList)
	mux.HandleFunc(dm.DocType_PathView, DocType_HandlerView)
	mux.HandleFunc(dm.DocType_PathEdit, DocType_HandlerEdit)
	mux.HandleFunc(dm.DocType_PathNew, DocType_HandlerNew)
	mux.HandleFunc(dm.DocType_PathSave, DocType_HandlerSave)
	mux.HandleFunc(dm.DocType_PathDelete, DocType_HandlerDelete)
	logs.Publish("Application", dm.DocType_Title)
    core.API.AddRoute(dm.DocType_Title, dm.DocType_Path, "", dm.DocType_QueryString, "Application")
	// 
	// Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//DocType_HandlerList is the handler for the list page
//Allows Listing of DocType records
//DocType_HandlerList - Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local
func DocType_HandlerList(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local
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

	objectName := dao.Translate("ObjectName", "DocType")
	reqField := "Base"
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule)
	if filter == "" {
		logs.Warning("No filter found : " + reqField + " for Object: " + objectName)
	} 

	noItems, returnList, _ := dao.DocType_GetListFiltered(filter)

	pageDetail := dm.DocType_PageList{
		Title:            CardTitle(dm.DocType_Title, core.Action_List),
		PageTitle:        PageTitle(dm.DocType_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	nextTemplate :=  NextTemplate("DocType", "List", dm.DocType_TemplateList)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END

}


//DocType_HandlerView is the handler used to View a page
//Allows Viewing for an existing DocType record
//DocType_HandlerView - Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func DocType_HandlerView(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local
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

	nextTemplate :=  NextTemplate("DocType", "View", dm.DocType_TemplateView)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//DocType_HandlerEdit is the handler used generate the Edit page
//Allows Editing for an existing DocType record and then allows the user to save the changes
//DocType_HandlerEdit - Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func DocType_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local
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
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.DocType
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.DocType)
	} else {
		_, rD, _ = dao.DocType_GetByID(searchID)
	}

	
	pageDetail := dm.DocType_Page{
		Title:       CardTitle(dm.DocType_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.DocType_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = doctype_PopulatePage(rD , pageDetail) 


	nextTemplate :=  NextTemplate("DocType", "Edit", dm.DocType_TemplateEdit)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//DocType_HandlerSave is the handler used process the saving of an DocType
//It is called from the Edit and New pages
//DocType_HandlerSave  - Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func DocType_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// 
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("DocTypeID")
	logs.Servicing(r.URL.Path+itemID)

	item := doctype_DataFromRequest(r)
	
	item, errStore := dao.DocType_Store(item,r)
	if errStore == nil {
		nextTemplate :=  NextTemplate("DocType", "Save", dm.DocType_Redirect)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.DocType_Name, errStore.Error())
		//http.Redirect(w, r, r.Referer(), http.StatusFound)
		ExecuteRedirect(r.Referer(), w, r,dm.DocType_QueryString,itemID,item)
	}
	// 
	// Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//DocType_HandlerNew is the handler used process the creation of an DocType
//It will create a new DocType and then redirect to the Edit page
//DocType_HandlerNew  - Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func DocType_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.DocType
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.DocType)
	} else {
		_, _, rD, _ = dao.DocType_New()
	}



	pageDetail := dm.DocType_Page{
		Title:       CardTitle(dm.DocType_Title, core.Action_New),
		PageTitle:   PageTitle(dm.DocType_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = doctype_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("DocType", "New", dm.DocType_TemplateNew)
	ExecuteTemplate(nextTemplate, w, r, pageDetail)
	// 
	// Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}	


//DocType_HandlerDelete is the handler used process the deletion of an DocType
// It will delete the DocType and then redirect to the List page
//DocType_HandlerDelete - Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func DocType_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	//
	// Code Continues Below
	//
	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.DocType_QueryString)

	dao.DocType_Delete(searchID)	

	nextTemplate :=  NextTemplate("DocType", "Delete", dm.DocType_Redirect)
	http.Redirect(w, r, nextTemplate, http.StatusFound)
	// 
	// Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}


//doctype_PopulatePage Builds/Populates the DocType Page 
//doctype_PopulatePage Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func doctype_PopulatePage(rD dm.DocType, pageDetail dm.DocType_Page) dm.DocType_Page {
	// Real DB Fields
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
	// Add Pseudo/Extra Fields
	// Enrichment Fields 
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
	return pageDetail
}
//doctype_DataFromRequest is used process the content of an HTTP Request and return an instance of an DocType
//doctype_DataFromRequest Auto generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
func doctype_DataFromRequest(r *http.Request) dm.DocType {

	var item dm.DocType
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
	return item
}
