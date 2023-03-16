package application
// ----------------------------------------------------------------
// Automatically generated  "/application/catalog.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Catalog (catalog)
// Endpoint 	        : Catalog (ID)
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

//Catalog_Publish annouces the endpoints available for this object
func Catalog_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Catalog_Path, Catalog_Handler)
	mux.HandleFunc(dm.Catalog_PathList, Catalog_HandlerList)
	mux.HandleFunc(dm.Catalog_PathView, Catalog_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
    core.API = core.API.AddRoute(dm.Catalog_Title, dm.Catalog_Path, "", dm.Catalog_QueryString, "Application")
	logs.Publish("Application", dm.Catalog_Title)
}

//Catalog_HandlerList is the handler for the Catalog list page
func Catalog_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Catalog

	objectName := dao.Translate("ObjectName", "Catalog")
	reqField := "Base"
	usage := "Defines a filter for the list of Catalog records." + core.TEXTAREA_CR
	usage = usage + "Fields can be any of those in the underlying DB table." + core.TEXTAREA_CR
	usage = usage + "Examples Below:" + core.TEXTAREA_CR
	usage = usage + "* datalength(_deleted) = 0 or " + core.TEXTAREA_CR 
	usage = usage + "* class IN ('x','y','z')"
	
	filter,_ := dao.Data_GetString(objectName, reqField, dm.Data_Category_FilterRule,usage)
	if filter == "" {
		logs.Warning("No filter found : " + reqField + " for Object: " + objectName)
	} 

	noItems, returnList, _ := dao.Catalog_GetListFiltered(filter)

	pageDetail := dm.Catalog_PageList{
		Title:            CardTitle(dm.Catalog_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Catalog_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	nextTemplate :=  NextTemplate("Catalog", "List", dm.Catalog_TemplateList)
	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}

//Catalog_HandlerView is the handler used to View a Catalog database record
func Catalog_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Catalog_QueryString)
	_, rD, _ := dao.Catalog_GetByID(searchID)

	pageDetail := dm.Catalog_Page{
		Title:       CardTitle(dm.Catalog_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Catalog_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	pageDetail = catalog_PopulatePage(rD , pageDetail) 

	nextTemplate :=  NextTemplate("Catalog", "View", dm.Catalog_TemplateView)
	nextTemplate = catalog_URIQueryData(nextTemplate,rD,searchID)

	ExecuteTemplate(nextTemplate, w, r, pageDetail)
}
//catalog_PopulatePage Builds/Populates the Catalog Page from an instance of Catalog from the Data Model
func catalog_PopulatePage(rD dm.Catalog, pageDetail dm.Catalog_Page) dm.Catalog_Page {
	// Real DB Fields
	pageDetail.ID = rD.ID
	pageDetail.Endpoint = rD.Endpoint
	pageDetail.Descr = rD.Descr
	pageDetail.Query = rD.Query
	pageDetail.Source = rD.Source
	// Add Pseudo/Extra Fields, fields that are not in the DB but are used in the UI
	// Enrichment content, content used provide lookups,lists etc
	// Add the Properties for the Fields
	pageDetail.ID_props = rD.ID_props
	pageDetail.Endpoint_props = rD.Endpoint_props
	pageDetail.Descr_props = rD.Descr_props
	pageDetail.Query_props = rD.Query_props
	pageDetail.Source_props = rD.Source_props
	return pageDetail
}//catalog_URIQueryData is used to replace the wildcards in the URI Query Path with the values from the Catalog Data Model
func catalog_URIQueryData(queryPath string,item dm.Catalog,itemID string) string {
	if queryPath == "" {
		return ""
	}
	queryPath = core.ReplaceWildcard(queryPath, strings.ToUpper("ID"), itemID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Catalog_ID_scrn), item.ID)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Catalog_Endpoint_scrn), item.Endpoint)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Catalog_Descr_scrn), item.Descr)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Catalog_Query_scrn), item.Query)
	queryPath = core.ReplaceWildcard(queryPath, "!"+strings.ToUpper(dm.Catalog_Source_scrn), item.Source)
	return queryPath
}