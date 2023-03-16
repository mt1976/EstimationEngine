package application
// ----------------------------------------------------------------
// Automatically generated  "/application/translation.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Translation (translation)
// Endpoint 	        : Translation (Message)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/03/2023 at 19:24:50
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"
	"strings"

	core    "github.com/mt1976/ebEstimates/core"
	dao     "github.com/mt1976/ebEstimates/dao"
	dm      "github.com/mt1976/ebEstimates/datamodel"
	logs    "github.com/mt1976/ebEstimates/logs"
)

//Translation_Publish annouces the endpoints available for this object
func Translation_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Translation_Path, Translation_Handler)
	mux.HandleFunc(dm.Translation_PathList, Translation_HandlerList)
	mux.HandleFunc(dm.Translation_PathView, Translation_HandlerView)
	mux.HandleFunc(dm.Translation_PathEdit, Translation_HandlerEdit)
	//Cannot Create via GUI
	mux.HandleFunc(dm.Translation_PathSave, Translation_HandlerSave)
	//Cannot Delete via GUI
    core.API = core.API.AddRoute(dm.Translation_Title, dm.Translation_Path, "", dm.Translation_QueryString, "Application")
	logs.Publish("Application", dm.Translation_Title)
}

//Translation_HandlerList is the handler for the Translation list page
func Translation_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Translation

	objectName := dao.Translate("ObjectName", "Translation")
	reqField := "Base"
	usage := "Defines a filter for the list of Translation records." + core.TEXTAREA_CR
	usage = usage + "Fields can be any of those in the underlying DB table." + core.TEXTAREA_CR
	usage = usage + "Examples Below:" + core.TEXTAREA_CR
	usage = usage + "* datalength(_deleted) = 0 or " + core.TEXTAREA_CR 
	usage = usage + "* class IN ('x','y','z')"
	
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule,usage)
	if filter == "" {
		logs.Warning("No filter found : " + reqField + " for Object: " + objectName)
	} 

	noItems, returnList, _ := dao.Translation_GetListFiltered(filter)

	pageDetail := dm.Translation_PageList{
		Title:            CardTitle(dm.Translation_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Translation_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	nextTemplate :=  NextTemplate("Translation", "List", dm.Translation_TemplateList)
	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Translation_HandlerView is the handler used to View a Translation database record
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

	nextTemplate :=  NextTemplate("Translation", "View", dm.Translation_TemplateView)
	nextTemplate = translation_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Translation_HandlerEdit is the handler used to Edit of an existing Translation database record
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
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.Translation
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.Translation)
	} else {
		_, rD, _ = dao.Translation_GetByID(searchID)
	}
	
	pageDetail := dm.Translation_Page{
		Title:       CardTitle(dm.Translation_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Translation_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = translation_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Translation", "Edit", dm.Translation_TemplateEdit)
	nextTemplate = translation_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Translation_HandlerSave is the handler used process the saving of an Translation database record, either new or existing, referenced by Edit & New Handlers.
func Translation_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("Id")
	logs.Servicing(r.URL.Path+itemID)

	item := translation_DataFromRequest(r)
	
	item, errStore := dao.Translation_Store(item,r)
	if errStore == nil {
		nextTemplate :=  NextTemplate("Translation", "Save", dm.Translation_Redirect)
		nextTemplate = translation_URIQueryData(nextTemplate,item,itemID)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.Translation_Name, errStore.Error())
		ExecuteRedirect(r.Referer(), w, r,dm.Translation_QueryString,itemID,item)
	}
}
//translation_PopulatePage Builds/Populates the Translation Page from an instance of Translation from the Data Model
func translation_PopulatePage(rD dm.Translation, pageDetail dm.Translation_Page) dm.Translation_Page {
	// Real DB Fields
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
	pageDetail.SYSDbVersion = rD.SYSDbVersion
	pageDetail.Migrate = rD.Migrate
	// Add Pseudo/Extra Fields, fields that are not in the DB but are used in the UI
	// Enrichment content, content used provide lookups,lists etc
	pageDetail.Migrate_lookup = dao.StubLists_Get("tf")
	// Add the Properties for the Fields
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
	pageDetail.SYSDbVersion_props = rD.SYSDbVersion_props
	pageDetail.Migrate_props = rD.Migrate_props
	return pageDetail
}
//translation_DataFromRequest is used process the content of an HTTP Request and return an instance of an Translation
func translation_DataFromRequest(r *http.Request) dm.Translation {
	var item dm.Translation
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
	item.SYSDbVersion = r.FormValue(dm.Translation_SYSDbVersion_scrn)
	item.Migrate = r.FormValue(dm.Translation_Migrate_scrn)
	return item
}
//translation_URIQueryData is used to replace the wildcards in the URI Query Path with the values from the Translation Data Model
func translation_URIQueryData(queryPath string,item dm.Translation,itemID string) string {
	if queryPath == "" {
		return ""
	}
	queryPath = core.ReplaceWildcard(queryPath, strings.ToUpper("ID"), itemID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Translation_Id_scrn), item.Id)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Translation_Class_scrn), item.Class)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Translation_Message_scrn), item.Message)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Translation_Translation_scrn), item.Translation)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Translation_Migrate_scrn), item.Migrate)
	return queryPath
}