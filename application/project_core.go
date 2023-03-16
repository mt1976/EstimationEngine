package application
// ----------------------------------------------------------------
// Automatically generated  "/application/project.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Project (project)
// Endpoint 	        : Project (ProjectID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/03/2023 at 19:24:49
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

//Project_Publish annouces the endpoints available for this object
func Project_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Project_Path, Project_Handler)
	mux.HandleFunc(dm.Project_PathList, Project_HandlerList)
	mux.HandleFunc(dm.Project_PathView, Project_HandlerView)
	mux.HandleFunc(dm.Project_PathEdit, Project_HandlerEdit)
	mux.HandleFunc(dm.Project_PathNew, Project_HandlerNew)
	mux.HandleFunc(dm.Project_PathSave, Project_HandlerSave)
	mux.HandleFunc(dm.Project_PathDelete, Project_HandlerDelete)
    core.API = core.API.AddRoute(dm.Project_Title, dm.Project_Path, "", dm.Project_QueryString, "Application")
	logs.Publish("Application", dm.Project_Title)
}

//Project_HandlerList is the handler for the Project list page
func Project_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Project

	objectName := dao.Translate("ObjectName", "Project")
	reqField := "Base"
	usage := "Defines a filter for the list of Project records." + core.TEXTAREA_CR
	usage = usage + "Fields can be any of those in the underlying DB table." + core.TEXTAREA_CR
	usage = usage + "Examples Below:" + core.TEXTAREA_CR
	usage = usage + "* datalength(_deleted) = 0 or " + core.TEXTAREA_CR 
	usage = usage + "* class IN ('x','y','z')"
	
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule,usage)
	if filter == "" {
		logs.Warning("No filter found : " + reqField + " for Object: " + objectName)
	} 

	noItems, returnList, _ := dao.Project_GetListFiltered(filter)

	pageDetail := dm.Project_PageList{
		Title:            CardTitle(dm.Project_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Project_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	nextTemplate :=  NextTemplate("Project", "List", dm.Project_TemplateList)
	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Project_HandlerView is the handler used to View a Project database record
func Project_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Project_QueryString)
	_, rD, _ := dao.Project_GetByID(searchID)

	pageDetail := dm.Project_Page{
		Title:       CardTitle(dm.Project_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Project_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = project_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Project", "View", dm.Project_TemplateView)
	nextTemplate = project_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Project_HandlerEdit is the handler used to Edit of an existing Project database record
func Project_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Project_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.Project
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.Project)
	} else {
		_, rD, _ = dao.Project_GetByID(searchID)
	}
	
	pageDetail := dm.Project_Page{
		Title:       CardTitle(dm.Project_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Project_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = project_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Project", "Edit", dm.Project_TemplateEdit)
	nextTemplate = project_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Project_HandlerSave is the handler used process the saving of an Project database record, either new or existing, referenced by Edit & New Handlers.
func Project_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("ProjectID")
	logs.Servicing(r.URL.Path+itemID)

	item := project_DataFromRequest(r)
	
	item, errStore := dao.Project_Store(item,r)
	if errStore == nil {
		nextTemplate :=  NextTemplate("Project", "Save", dm.Project_Redirect)
		nextTemplate = project_URIQueryData(nextTemplate,item,itemID)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.Project_Name, errStore.Error())
		ExecuteRedirect(r.Referer(), w, r,dm.Project_QueryString,itemID,item)
	}
}

//Project_HandlerNew is the handler used process the creation of an new Project database record, then redirect to Edit
func Project_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Project_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.Project
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.Project)
	} else {
		_, _, rD, _ = dao.Project_New()
	}

	pageDetail := dm.Project_Page{
		Title:       CardTitle(dm.Project_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Project_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = project_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Project", "New", dm.Project_TemplateNew)
	nextTemplate = project_URIQueryData(nextTemplate,dm.Project{},searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}	

//Project_HandlerDelete is the handler used process the deletion of an Project database record. May be Hard or SoftDelete.
func Project_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Project_QueryString)
	// DAO Call to Delete Project Record, may be SoftDelete or HardDelete depending on the DAO implementation
	dao.Project_Delete(searchID)	

	nextTemplate :=  NextTemplate("Project", "Delete", dm.Project_Redirect)
	nextTemplate = project_URIQueryData(nextTemplate,dm.Project{},searchID)
	http.Redirect(w, r, nextTemplate, http.StatusFound)
}
//project_PopulatePage Builds/Populates the Project Page from an instance of Project from the Data Model
func project_PopulatePage(rD dm.Project, pageDetail dm.Project_Page) dm.Project_Page {
	// Real DB Fields
	pageDetail.SYSId = rD.SYSId
	pageDetail.ProjectID = rD.ProjectID
	pageDetail.OriginID = rD.OriginID
	pageDetail.ProjectStateID = rD.ProjectStateID
	pageDetail.ProfileID = rD.ProfileID
	pageDetail.Name = rD.Name
	pageDetail.Description = rD.Description
	pageDetail.StartDate = rD.StartDate
	pageDetail.EndDate = rD.EndDate
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
	pageDetail.SYSDbVersion = rD.SYSDbVersion
	pageDetail.Comments = rD.Comments
	pageDetail.ProjectRate = rD.ProjectRate
	pageDetail.DefaultRate = rD.DefaultRate
	pageDetail.ProjectAnalyst = rD.ProjectAnalyst
	pageDetail.ProjectEngineer = rD.ProjectEngineer
	pageDetail.ProjectManager = rD.ProjectManager
	pageDetail.Releases = rD.Releases
	pageDetail.Notes = rD.Notes
	// Add Pseudo/Extra Fields, fields that are not in the DB but are used in the UI
	pageDetail.NoEstimationSessions = rD.NoEstimationSessions
	pageDetail.OriginName = rD.OriginName
	pageDetail.OriginKey = rD.OriginKey
	// Enrichment content, content used provide lookups,lists etc
	pageDetail.OriginID_lookup = dao.Origin_GetLookup()
	pageDetail.ProjectStateID_lookup = dao.ProjectState_GetLookup()
	pageDetail.ProfileID_lookup = dao.Profile_GetLookup()
	pageDetail.ProjectAnalyst_lookup = dao.Resource_GetFilteredLookup("Project","ProjectAnalyst")
	pageDetail.ProjectEngineer_lookup = dao.Resource_GetFilteredLookup("Project","ProjectEngineer")
	pageDetail.ProjectManager_lookup = dao.Resource_GetFilteredLookup("Project","ProjectManager")
	// Add the Properties for the Fields
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.ProjectID_props = rD.ProjectID_props
	pageDetail.OriginID_props = rD.OriginID_props
	pageDetail.ProjectStateID_props = rD.ProjectStateID_props
	pageDetail.ProfileID_props = rD.ProfileID_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.Description_props = rD.Description_props
	pageDetail.StartDate_props = rD.StartDate_props
	pageDetail.EndDate_props = rD.EndDate_props
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
	pageDetail.SYSDbVersion_props = rD.SYSDbVersion_props
	pageDetail.Comments_props = rD.Comments_props
	pageDetail.ProjectRate_props = rD.ProjectRate_props
	pageDetail.DefaultRate_props = rD.DefaultRate_props
	pageDetail.ProjectAnalyst_props = rD.ProjectAnalyst_props
	pageDetail.ProjectEngineer_props = rD.ProjectEngineer_props
	pageDetail.ProjectManager_props = rD.ProjectManager_props
	pageDetail.Releases_props = rD.Releases_props
	pageDetail.Notes_props = rD.Notes_props
	pageDetail.NoEstimationSessions_props = rD.NoEstimationSessions_props
	pageDetail.OriginName_props = rD.OriginName_props
	pageDetail.OriginKey_props = rD.OriginKey_props
	return pageDetail
}
//project_DataFromRequest is used process the content of an HTTP Request and return an instance of an Project
func project_DataFromRequest(r *http.Request) dm.Project {
	var item dm.Project
	item.SYSId = r.FormValue(dm.Project_SYSId_scrn)
	item.ProjectID = r.FormValue(dm.Project_ProjectID_scrn)
	item.OriginID = r.FormValue(dm.Project_OriginID_scrn)
	item.ProjectStateID = r.FormValue(dm.Project_ProjectStateID_scrn)
	item.ProfileID = r.FormValue(dm.Project_ProfileID_scrn)
	item.Name = r.FormValue(dm.Project_Name_scrn)
	item.Description = r.FormValue(dm.Project_Description_scrn)
	item.StartDate = r.FormValue(dm.Project_StartDate_scrn)
	item.EndDate = r.FormValue(dm.Project_EndDate_scrn)
	item.SYSCreated = r.FormValue(dm.Project_SYSCreated_scrn)
	item.SYSCreatedBy = r.FormValue(dm.Project_SYSCreatedBy_scrn)
	item.SYSCreatedHost = r.FormValue(dm.Project_SYSCreatedHost_scrn)
	item.SYSUpdated = r.FormValue(dm.Project_SYSUpdated_scrn)
	item.SYSUpdatedBy = r.FormValue(dm.Project_SYSUpdatedBy_scrn)
	item.SYSUpdatedHost = r.FormValue(dm.Project_SYSUpdatedHost_scrn)
	item.SYSDeleted = r.FormValue(dm.Project_SYSDeleted_scrn)
	item.SYSDeletedBy = r.FormValue(dm.Project_SYSDeletedBy_scrn)
	item.SYSDeletedHost = r.FormValue(dm.Project_SYSDeletedHost_scrn)
	item.SYSActivity = r.FormValue(dm.Project_SYSActivity_scrn)
	item.SYSDbVersion = r.FormValue(dm.Project_SYSDbVersion_scrn)
	item.Comments = r.FormValue(dm.Project_Comments_scrn)
	item.ProjectRate = r.FormValue(dm.Project_ProjectRate_scrn)
	item.DefaultRate = r.FormValue(dm.Project_DefaultRate_scrn)
	item.ProjectAnalyst = r.FormValue(dm.Project_ProjectAnalyst_scrn)
	item.ProjectEngineer = r.FormValue(dm.Project_ProjectEngineer_scrn)
	item.ProjectManager = r.FormValue(dm.Project_ProjectManager_scrn)
	item.Releases = r.FormValue(dm.Project_Releases_scrn)
	item.Notes = r.FormValue(dm.Project_Notes_scrn)
	item.NoEstimationSessions = r.FormValue(dm.Project_NoEstimationSessions_scrn)
	item.OriginName = r.FormValue(dm.Project_OriginName_scrn)
	item.OriginKey = r.FormValue(dm.Project_OriginKey_scrn)
	return item
}
//project_URIQueryData is used to replace the wildcards in the URI Query Path with the values from the Project Data Model
func project_URIQueryData(queryPath string,item dm.Project,itemID string) string {
	if queryPath == "" {
		return ""
	}
	queryPath = core.ReplaceWildcard(queryPath, strings.ToUpper("ID"), itemID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Project_ProjectID_scrn), item.ProjectID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Project_OriginID_scrn), item.OriginID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Project_ProjectStateID_scrn), item.ProjectStateID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Project_ProfileID_scrn), item.ProfileID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Project_Name_scrn), item.Name)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Project_Description_scrn), item.Description)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Project_StartDate_scrn), item.StartDate)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Project_EndDate_scrn), item.EndDate)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Project_Comments_scrn), item.Comments)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Project_ProjectRate_scrn), item.ProjectRate)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Project_DefaultRate_scrn), item.DefaultRate)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Project_ProjectAnalyst_scrn), item.ProjectAnalyst)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Project_ProjectEngineer_scrn), item.ProjectEngineer)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Project_ProjectManager_scrn), item.ProjectManager)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Project_Releases_scrn), item.Releases)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Project_Notes_scrn), item.Notes)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Project_NoEstimationSessions_scrn), item.NoEstimationSessions)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Project_OriginName_scrn), item.OriginName)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Project_OriginKey_scrn), item.OriginKey)
	return queryPath
}