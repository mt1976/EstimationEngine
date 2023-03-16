package application
// ----------------------------------------------------------------
// Automatically generated  "/application/data.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Data (data)
// Endpoint 	        : Data (DataID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/03/2023 at 19:24:47
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

//Data_Publish annouces the endpoints available for this object
func Data_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.Data_PathList, Data_HandlerList)
	mux.HandleFunc(dm.Data_PathView, Data_HandlerView)
	mux.HandleFunc(dm.Data_PathEdit, Data_HandlerEdit)
	mux.HandleFunc(dm.Data_PathNew, Data_HandlerNew)
	mux.HandleFunc(dm.Data_PathSave, Data_HandlerSave)
	mux.HandleFunc(dm.Data_PathDelete, Data_HandlerDelete)
    //No API
	logs.Publish("Application", dm.Data_Title)
}

//Data_HandlerList is the handler for the Data list page
func Data_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Data

	objectName := dao.Translate("ObjectName", "Data")
	reqField := "Base"
	usage := "Defines a filter for the list of Data records." + core.TEXTAREA_CR
	usage = usage + "Fields can be any of those in the underlying DB table." + core.TEXTAREA_CR
	usage = usage + "Examples Below:" + core.TEXTAREA_CR
	usage = usage + "* datalength(_deleted) = 0 or " + core.TEXTAREA_CR 
	usage = usage + "* class IN ('x','y','z')"
	
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule,usage)
	if filter == "" {
		logs.Warning("No filter found : " + reqField + " for Object: " + objectName)
	} 

	noItems, returnList, _ := dao.Data_GetListFiltered(filter)

	pageDetail := dm.Data_PageList{
		Title:            CardTitle(dm.Data_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Data_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	nextTemplate :=  NextTemplate("Data", "List", dm.Data_TemplateList)
	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Data_HandlerView is the handler used to View a Data database record
func Data_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
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

	nextTemplate :=  NextTemplate("Data", "View", dm.Data_TemplateView)
	nextTemplate = data_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Data_HandlerEdit is the handler used to Edit of an existing Data database record
func Data_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Data_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.Data
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.Data)
	} else {
		_, rD, _ = dao.Data_GetByID(searchID)
	}
	
	pageDetail := dm.Data_Page{
		Title:       CardTitle(dm.Data_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Data_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = data_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Data", "Edit", dm.Data_TemplateEdit)
	nextTemplate = data_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Data_HandlerSave is the handler used process the saving of an Data database record, either new or existing, referenced by Edit & New Handlers.
func Data_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("DataID")
	logs.Servicing(r.URL.Path+itemID)

	item := data_DataFromRequest(r)
	
	item, errStore := dao.Data_Store(item,r)
	if errStore == nil {
		nextTemplate :=  NextTemplate("Data", "Save", dm.Data_Redirect)
		nextTemplate = data_URIQueryData(nextTemplate,item,itemID)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.Data_Name, errStore.Error())
		ExecuteRedirect(r.Referer(), w, r,dm.Data_QueryString,itemID,item)
	}
}

//Data_HandlerNew is the handler used process the creation of an new Data database record, then redirect to Edit
func Data_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Data_QueryString)
	action := core.GetURLparam(r, core.ContextState)

	var rD dm.Data
	if action == core.ContextState_ERROR {
		rD = core.SessionManager.Get(r.Context(), searchID).(dm.Data)
	} else {
		_, _, rD, _ = dao.Data_New()
	}

	pageDetail := dm.Data_Page{
		Title:       CardTitle(dm.Data_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Data_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = data_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Data", "New", dm.Data_TemplateNew)
	nextTemplate = data_URIQueryData(nextTemplate,dm.Data{},searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}	

//Data_HandlerDelete is the handler used process the deletion of an Data database record. May be Hard or SoftDelete.
func Data_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Data_QueryString)
	// DAO Call to Delete Data Record, may be SoftDelete or HardDelete depending on the DAO implementation
	dao.Data_Delete(searchID)	

	nextTemplate :=  NextTemplate("Data", "Delete", dm.Data_Redirect)
	nextTemplate = data_URIQueryData(nextTemplate,dm.Data{},searchID)
	http.Redirect(w, r, nextTemplate, http.StatusFound)
}
//data_PopulatePage Builds/Populates the Data Page from an instance of Data from the Data Model
func data_PopulatePage(rD dm.Data, pageDetail dm.Data_Page) dm.Data_Page {
	// Real DB Fields
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
	pageDetail.Category = rD.Category
	pageDetail.Migrate = rD.Migrate
	pageDetail.Usage = rD.Usage
	// Add Pseudo/Extra Fields, fields that are not in the DB but are used in the UI
	// Enrichment content, content used provide lookups,lists etc
	pageDetail.Migrate_lookup = dao.StubLists_Get("tf")
	// Add the Properties for the Fields
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
	pageDetail.Category_props = rD.Category_props
	pageDetail.Migrate_props = rD.Migrate_props
	pageDetail.Usage_props = rD.Usage_props
	return pageDetail
}
//data_DataFromRequest is used process the content of an HTTP Request and return an instance of an Data
func data_DataFromRequest(r *http.Request) dm.Data {
	var item dm.Data
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
	item.Category = r.FormValue(dm.Data_Category_scrn)
	item.Migrate = r.FormValue(dm.Data_Migrate_scrn)
	item.Usage = r.FormValue(dm.Data_Usage_scrn)
	return item
}
//data_URIQueryData is used to replace the wildcards in the URI Query Path with the values from the Data Data Model
func data_URIQueryData(queryPath string,item dm.Data,itemID string) string {
	if queryPath == "" {
		return ""
	}
	queryPath = core.ReplaceWildcard(queryPath, strings.ToUpper("ID"), itemID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Data_DataID_scrn), item.DataID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Data_Class_scrn), item.Class)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Data_Field_scrn), item.Field)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Data_Value_scrn), item.Value)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Data_Category_scrn), item.Category)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Data_Migrate_scrn), item.Migrate)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Data_Usage_scrn), item.Usage)
	return queryPath
}