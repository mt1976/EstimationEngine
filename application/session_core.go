package application
// ----------------------------------------------------------------
// Automatically generated  "/application/session.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Session (session)
// Endpoint 	        : Session (SessionID)
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

//Session_Publish annouces the endpoints available for this object
func Session_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.Session_PathList, Session_HandlerList)
	mux.HandleFunc(dm.Session_PathView, Session_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	mux.HandleFunc(dm.Session_PathSave, Session_HandlerSave)
	//Cannot Delete via GUI
    //No API
	logs.Publish("Application", dm.Session_Title)
}

//Session_HandlerList is the handler for the Session list page
func Session_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Session

	objectName := dao.Translate("ObjectName", "Session")
	reqField := "Base"
	usage := "Defines a filter for the list of Session records." + core.TEXTAREA_CR
	usage = usage + "Fields can be any of those in the underlying DB table." + core.TEXTAREA_CR
	usage = usage + "Examples Below:" + core.TEXTAREA_CR
	usage = usage + "* datalength(_deleted) = 0 or " + core.TEXTAREA_CR 
	usage = usage + "* class IN ('x','y','z')"
	
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule,usage)
	if filter == "" {
		logs.Warning("No filter found : " + reqField + " for Object: " + objectName)
	} 

	noItems, returnList, _ := dao.Session_GetListFiltered(filter)

	pageDetail := dm.Session_PageList{
		Title:            CardTitle(dm.Session_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Session_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	nextTemplate :=  NextTemplate("Session", "List", dm.Session_TemplateList)
	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Session_HandlerView is the handler used to View a Session database record
func Session_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Session_QueryString)
	_, rD, _ := dao.Session_GetByID(searchID)

	pageDetail := dm.Session_Page{
		Title:       CardTitle(dm.Session_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Session_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = session_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Session", "View", dm.Session_TemplateView)
	nextTemplate = session_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Session_HandlerSave is the handler used process the saving of an Session database record, either new or existing, referenced by Edit & New Handlers.
func Session_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	itemID := r.FormValue("Id")
	logs.Servicing(r.URL.Path+itemID)

	item := session_DataFromRequest(r)
	
	item, errStore := dao.Session_Store(item,r)
	if errStore == nil {
		nextTemplate :=  NextTemplate("Session", "Save", dm.Session_Redirect)
		nextTemplate = session_URIQueryData(nextTemplate,item,itemID)
		http.Redirect(w, r, nextTemplate, http.StatusFound)
	} else {
		logs.Information(dm.Session_Name, errStore.Error())
		ExecuteRedirect(r.Referer(), w, r,dm.Session_QueryString,itemID,item)
	}
}
//session_PopulatePage Builds/Populates the Session Page from an instance of Session from the Data Model
func session_PopulatePage(rD dm.Session, pageDetail dm.Session_Page) dm.Session_Page {
	// Real DB Fields
	pageDetail.SYSId = rD.SYSId
	pageDetail.Id = rD.Id
	pageDetail.Apptoken = rD.Apptoken
	pageDetail.Createdate = rD.Createdate
	pageDetail.Createtime = rD.Createtime
	pageDetail.Uniqueid = rD.Uniqueid
	pageDetail.Sessiontoken = rD.Sessiontoken
	pageDetail.Username = rD.Username
	pageDetail.Password = rD.Password
	pageDetail.Userip = rD.Userip
	pageDetail.Userhost = rD.Userhost
	pageDetail.Appip = rD.Appip
	pageDetail.Apphost = rD.Apphost
	pageDetail.Issued = rD.Issued
	pageDetail.Expiry = rD.Expiry
	pageDetail.Expiryraw = rD.Expiryraw
	pageDetail.Expires = rD.Expires
	pageDetail.Brand = rD.Brand
	pageDetail.SessionRole = rD.SessionRole
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
	// Add Pseudo/Extra Fields, fields that are not in the DB but are used in the UI
	// Enrichment content, content used provide lookups,lists etc
	// Add the Properties for the Fields
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.Id_props = rD.Id_props
	pageDetail.Apptoken_props = rD.Apptoken_props
	pageDetail.Createdate_props = rD.Createdate_props
	pageDetail.Createtime_props = rD.Createtime_props
	pageDetail.Uniqueid_props = rD.Uniqueid_props
	pageDetail.Sessiontoken_props = rD.Sessiontoken_props
	pageDetail.Username_props = rD.Username_props
	pageDetail.Password_props = rD.Password_props
	pageDetail.Userip_props = rD.Userip_props
	pageDetail.Userhost_props = rD.Userhost_props
	pageDetail.Appip_props = rD.Appip_props
	pageDetail.Apphost_props = rD.Apphost_props
	pageDetail.Issued_props = rD.Issued_props
	pageDetail.Expiry_props = rD.Expiry_props
	pageDetail.Expiryraw_props = rD.Expiryraw_props
	pageDetail.Expires_props = rD.Expires_props
	pageDetail.Brand_props = rD.Brand_props
	pageDetail.SessionRole_props = rD.SessionRole_props
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
	return pageDetail
}
//session_DataFromRequest is used process the content of an HTTP Request and return an instance of an Session
func session_DataFromRequest(r *http.Request) dm.Session {
	var item dm.Session
	item.SYSId = r.FormValue(dm.Session_SYSId_scrn)
	item.Id = r.FormValue(dm.Session_Id_scrn)
	item.Apptoken = r.FormValue(dm.Session_Apptoken_scrn)
	item.Createdate = r.FormValue(dm.Session_Createdate_scrn)
	item.Createtime = r.FormValue(dm.Session_Createtime_scrn)
	item.Uniqueid = r.FormValue(dm.Session_Uniqueid_scrn)
	item.Sessiontoken = r.FormValue(dm.Session_Sessiontoken_scrn)
	item.Username = r.FormValue(dm.Session_Username_scrn)
	item.Password = r.FormValue(dm.Session_Password_scrn)
	item.Userip = r.FormValue(dm.Session_Userip_scrn)
	item.Userhost = r.FormValue(dm.Session_Userhost_scrn)
	item.Appip = r.FormValue(dm.Session_Appip_scrn)
	item.Apphost = r.FormValue(dm.Session_Apphost_scrn)
	item.Issued = r.FormValue(dm.Session_Issued_scrn)
	item.Expiry = r.FormValue(dm.Session_Expiry_scrn)
	item.Expiryraw = r.FormValue(dm.Session_Expiryraw_scrn)
	item.Expires = r.FormValue(dm.Session_Expires_scrn)
	item.Brand = r.FormValue(dm.Session_Brand_scrn)
	item.SessionRole = r.FormValue(dm.Session_SessionRole_scrn)
	item.SYSCreated = r.FormValue(dm.Session_SYSCreated_scrn)
	item.SYSCreatedBy = r.FormValue(dm.Session_SYSCreatedBy_scrn)
	item.SYSCreatedHost = r.FormValue(dm.Session_SYSCreatedHost_scrn)
	item.SYSUpdated = r.FormValue(dm.Session_SYSUpdated_scrn)
	item.SYSUpdatedBy = r.FormValue(dm.Session_SYSUpdatedBy_scrn)
	item.SYSUpdatedHost = r.FormValue(dm.Session_SYSUpdatedHost_scrn)
	item.SYSDeleted = r.FormValue(dm.Session_SYSDeleted_scrn)
	item.SYSDeletedBy = r.FormValue(dm.Session_SYSDeletedBy_scrn)
	item.SYSDeletedHost = r.FormValue(dm.Session_SYSDeletedHost_scrn)
	item.SYSDbVersion = r.FormValue(dm.Session_SYSDbVersion_scrn)
	return item
}
//session_URIQueryData is used to replace the wildcards in the URI Query Path with the values from the Session Data Model
func session_URIQueryData(queryPath string,item dm.Session,itemID string) string {
	if queryPath == "" {
		return ""
	}
	queryPath = core.ReplaceWildcard(queryPath, strings.ToUpper("ID"), itemID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Session_Id_scrn), item.Id)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Session_Apptoken_scrn), item.Apptoken)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Session_Createdate_scrn), item.Createdate)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Session_Createtime_scrn), item.Createtime)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Session_Uniqueid_scrn), item.Uniqueid)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Session_Sessiontoken_scrn), item.Sessiontoken)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Session_Username_scrn), item.Username)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Session_Password_scrn), item.Password)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Session_Userip_scrn), item.Userip)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Session_Userhost_scrn), item.Userhost)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Session_Appip_scrn), item.Appip)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Session_Apphost_scrn), item.Apphost)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Session_Issued_scrn), item.Issued)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Session_Expiry_scrn), item.Expiry)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Session_Expiryraw_scrn), item.Expiryraw)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Session_Expires_scrn), item.Expires)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Session_Brand_scrn), item.Brand)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Session_SessionRole_scrn), item.SessionRole)
	return queryPath
}