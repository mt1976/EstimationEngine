package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/origin.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Origin (origin)
// Endpoint 	        : Origin (OriginID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 08/12/2022 at 13:31:30
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Origin defines the datamolde for the Origin object
type Origin struct {


SYSId       string
SYSId_props FieldProperties
OriginID       string
OriginID_props FieldProperties
StateID       string
StateID_props FieldProperties
DocTypeID       string
DocTypeID_props FieldProperties
Code       string
Code_props FieldProperties
FullName       string
FullName_props FieldProperties
Rate       string
Rate_props FieldProperties
Notes       string
Notes_props FieldProperties
StartDate       string
StartDate_props FieldProperties
EndDate       string
EndDate_props FieldProperties
SYSCreated       string
SYSCreated_props FieldProperties
SYSCreatedBy       string
SYSCreatedBy_props FieldProperties
SYSCreatedHost       string
SYSCreatedHost_props FieldProperties
SYSUpdated       string
SYSUpdated_props FieldProperties
SYSUpdatedBy       string
SYSUpdatedBy_props FieldProperties
SYSUpdatedHost       string
SYSUpdatedHost_props FieldProperties
SYSDeleted       string
SYSDeleted_props FieldProperties
SYSDeletedBy       string
SYSDeletedBy_props FieldProperties
SYSDeletedHost       string
SYSDeletedHost_props FieldProperties
SYSActivity       string
SYSActivity_props FieldProperties
Currency       string
Currency_props FieldProperties
NoActiveProjects       string
NoActiveProjects_props FieldProperties
 // Any lookups will be added below

StateID_lookup []Lookup_Item
DocTypeID_lookup []Lookup_Item
















Currency_lookup []Lookup_Item


}

const (
	Origin_Title       = "Origin"
	Origin_SQLTable    = "originStore"
	Origin_SQLSearchID = "originID"
	Origin_QueryString = "OriginID"
	///
	/// Handler Defintions
	///
	Origin_Template     = "Origin"
	Origin_TemplateList = "/Origin/Origin_List"
	Origin_TemplateView = "/Origin/Origin_View"
	Origin_TemplateEdit = "/Origin/Origin_Edit"
	Origin_TemplateNew  = "/Origin/Origin_New"
	///
	/// Handler Monitor Paths
	///
	Origin_Path       = "/API/Origin/"
	Origin_PathList   = "/OriginList/"
	Origin_PathView   = "/OriginView/"
	Origin_PathEdit   = "/OriginEdit/"
	Origin_PathNew    = "/OriginNew/"
	Origin_PathSave   = "/OriginSave/"
	Origin_PathDelete = "/OriginDelete/"
	///
	//Origin_Redirect provides a page to return to aftern an action
	Origin_Redirect = Origin_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	Origin_SYSId_sql   = "_id" // SYSId is a Int
	Origin_OriginID_sql   = "originID" // OriginID is a String
	Origin_StateID_sql   = "stateID" // StateID is a String
	Origin_DocTypeID_sql   = "docTypeID" // DocTypeID is a String
	Origin_Code_sql   = "code" // Code is a String
	Origin_FullName_sql   = "fullName" // FullName is a String
	Origin_Rate_sql   = "rate" // Rate is a String
	Origin_Notes_sql   = "notes" // Notes is a String
	Origin_StartDate_sql   = "startDate" // StartDate is a String
	Origin_EndDate_sql   = "endDate" // EndDate is a String
	Origin_SYSCreated_sql   = "_created" // SYSCreated is a String
	Origin_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	Origin_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	Origin_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	Origin_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
	Origin_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
	Origin_SYSDeleted_sql   = "_deleted" // SYSDeleted is a String
	Origin_SYSDeletedBy_sql   = "_deletedBy" // SYSDeletedBy is a String
	Origin_SYSDeletedHost_sql   = "_deletedHost" // SYSDeletedHost is a String
	Origin_SYSActivity_sql   = "_activity" // SYSActivity is a String
	Origin_Currency_sql   = "currency" // Currency is a String
	Origin_NoActiveProjects_sql   = "NoActiveProjects" // NoActiveProjects is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	Origin_SYSId_scrn   = "SYSId" // SYSId is a Int
	Origin_OriginID_scrn   = "OriginID" // OriginID is a String
	Origin_StateID_scrn   = "StateID" // StateID is a String
	Origin_DocTypeID_scrn   = "DocTypeID" // DocTypeID is a String
	Origin_Code_scrn   = "Code" // Code is a String
	Origin_FullName_scrn   = "FullName" // FullName is a String
	Origin_Rate_scrn   = "Rate" // Rate is a String
	Origin_Notes_scrn   = "Notes" // Notes is a String
	Origin_StartDate_scrn   = "StartDate" // StartDate is a String
	Origin_EndDate_scrn   = "EndDate" // EndDate is a String
	Origin_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	Origin_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	Origin_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	Origin_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	Origin_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
	Origin_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
	Origin_SYSDeleted_scrn   = "SYSDeleted" // SYSDeleted is a String
	Origin_SYSDeletedBy_scrn   = "SYSDeletedBy" // SYSDeletedBy is a String
	Origin_SYSDeletedHost_scrn   = "SYSDeletedHost" // SYSDeletedHost is a String
	Origin_SYSActivity_scrn   = "SYSActivity" // SYSActivity is a String
	Origin_Currency_scrn   = "Currency" // Currency is a String
	Origin_NoActiveProjects_scrn   = "NoActiveProjects" // NoActiveProjects is a String

	/// Definitions End
	///
)

//origin_PageList provides the information for the template for a list of Origins
type Origin_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []Origin
	Context	 appContext
}

//origin_Page provides the information for the template for an individual Origin
type Origin_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	SYSId_props     FieldProperties
	OriginID         string
	OriginID_props     FieldProperties
	StateID         string
	StateID_lookup    []Lookup_Item
	StateID_props     FieldProperties
	DocTypeID         string
	DocTypeID_lookup    []Lookup_Item
	DocTypeID_props     FieldProperties
	Code         string
	Code_props     FieldProperties
	FullName         string
	FullName_props     FieldProperties
	Rate         string
	Rate_props     FieldProperties
	Notes         string
	Notes_props     FieldProperties
	StartDate         string
	StartDate_props     FieldProperties
	EndDate         string
	EndDate_props     FieldProperties
	SYSCreated         string
	SYSCreated_props     FieldProperties
	SYSCreatedBy         string
	SYSCreatedBy_props     FieldProperties
	SYSCreatedHost         string
	SYSCreatedHost_props     FieldProperties
	SYSUpdated         string
	SYSUpdated_props     FieldProperties
	SYSUpdatedBy         string
	SYSUpdatedBy_props     FieldProperties
	SYSUpdatedHost         string
	SYSUpdatedHost_props     FieldProperties
	SYSDeleted         string
	SYSDeleted_props     FieldProperties
	SYSDeletedBy         string
	SYSDeletedBy_props     FieldProperties
	SYSDeletedHost         string
	SYSDeletedHost_props     FieldProperties
	SYSActivity         string
	SYSActivity_props     FieldProperties
	Currency         string
	Currency_lookup    []Lookup_Item
	Currency_props     FieldProperties
	NoActiveProjects         string
	NoActiveProjects_props     FieldProperties
	// 
	// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}