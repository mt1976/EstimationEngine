package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/origin.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Origin (origin)
// Endpoint 	        : Origin (OriginID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 16/02/2023 at 10:11:36
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Origin defines the datamodel for the Origin object
type Origin struct {
	// Dynamically generated 16/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	SYSId       string
	OriginID       string
	StateID       string
	DocTypeID       string
	Code       string
	FullName       string
	Rate       string
	Notes       string
	StartDate       string
	EndDate       string
	SYSCreated       string
	SYSCreatedBy       string
	SYSCreatedHost       string
	SYSUpdated       string
	SYSUpdatedBy       string
	SYSUpdatedHost       string
	SYSDeleted       string
	SYSDeletedBy       string
	SYSDeletedHost       string
	SYSActivity       string
	Currency       string
	SYSDbVersion       string
	Comments       string
	ProjectManager       string
	AccountManager       string
	NoActiveProjects       string
	//
	// Field Properties
	//
	SYSId_props FieldProperties
	OriginID_props FieldProperties
	StateID_props FieldProperties
	DocTypeID_props FieldProperties
	Code_props FieldProperties
	FullName_props FieldProperties
	Rate_props FieldProperties
	Notes_props FieldProperties
	StartDate_props FieldProperties
	EndDate_props FieldProperties
	SYSCreated_props FieldProperties
	SYSCreatedBy_props FieldProperties
	SYSCreatedHost_props FieldProperties
	SYSUpdated_props FieldProperties
	SYSUpdatedBy_props FieldProperties
	SYSUpdatedHost_props FieldProperties
	SYSDeleted_props FieldProperties
	SYSDeletedBy_props FieldProperties
	SYSDeletedHost_props FieldProperties
	SYSActivity_props FieldProperties
	Currency_props FieldProperties
	SYSDbVersion_props FieldProperties
	Comments_props FieldProperties
	ProjectManager_props FieldProperties
	AccountManager_props FieldProperties
	NoActiveProjects_props FieldProperties
	//
 	// Any lookups will be added below
	//
	StateID_lookup []Lookup_Item
	DocTypeID_lookup []Lookup_Item
	Currency_lookup []Lookup_Item
	ProjectManager_lookup []Lookup_Item
	AccountManager_lookup []Lookup_Item
	}

const (
	// Dynamically generated 16/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	// Page Infrastructure
	Origin_Name      = "Origin"
	Origin_Title       = "Origin"
	Origin_SQLTable    = "originStore"
	Origin_SQLSearchID = "originID"
	Origin_QueryString = "OriginID"
	///
	/// Template Path Defintions
	///
	Origin_Template     = "Origin"
	Origin_TemplateList = "/Origin/OriginList"
	Origin_TemplateView = "/Origin/OriginView"
	Origin_TemplateEdit = "/Origin/OriginEdit"
	Origin_TemplateNew  = "/Origin/OriginNew"
	///
	/// URI Handler Paths
	///
	Origin_Path       = "/API/Origin/"
	Origin_PathList   = "/OriginList/"
	Origin_PathView   = "/OriginView/"
	Origin_PathEdit   = "/OriginEdit/"
	Origin_PathNew    = "/OriginNew/"
	Origin_PathSave   = "/OriginSave/"
	Origin_PathDelete = "/OriginDelete/"
	// Redirects - On Server Side Error
	Origin_PathEditException   = "/OriginEditException/"
	Origin_PathNewException    = "/OriginNewException/"
	//
	//Origin_Redirect provides a page to return to aftern an action
	Origin_Redirect = Origin_PathList
	
	//
	//
	// SQL Field Definitions
	//
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
	Origin_SYSDbVersion_sql   = "_dbVersion" // SYSDbVersion is a String
	Origin_Comments_sql   = "comments" // Comments is a String
	Origin_ProjectManager_sql   = "projectManager" // ProjectManager is a String
	Origin_AccountManager_sql   = "accountManager" // AccountManager is a String
	Origin_NoActiveProjects_sql   = "NoActiveProjects" // NoActiveProjects is a String
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
	Origin_SYSDbVersion_scrn   = "SYSDbVersion" // SYSDbVersion is a String
	Origin_Comments_scrn   = "Comments" // Comments is a String
	Origin_ProjectManager_scrn   = "ProjectManager" // ProjectManager is a String
	Origin_AccountManager_scrn   = "AccountManager" // AccountManager is a String
	Origin_NoActiveProjects_scrn   = "NoActiveProjects" // NoActiveProjects is a String
///
)

//Origin_PageList provides the information for the template for a list of Origins
type Origin_PageList struct {
	// Dynamically generated 16/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	// Page Infrastructure
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	// Context & Permisions
	Context	 appContext
	BlockEdit		 bool
	BlockSave 		 bool
	BlockCreate		 bool
	BlockDelete		 bool
	BlockValidate	 bool
	BlockView		 bool
	// Page Data
	ItemsOnPage 	 int
	ItemList  		 []Origin
}

//Origin_Page provides the information for the template for an individual Origin
type Origin_Page struct {
	// Dynamically generated 16/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	/// Page Infrastructure
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	/// Context & Permisions
	Context	 		 appContext
	BlockEdit		 bool
	BlockSave 		 bool
	BlockCreate		 bool
	BlockDelete		 bool
	BlockValidate	 bool
	BlockView		 bool
	/// Fields Definitions
	SYSId         string
	OriginID         string
	StateID         string
	DocTypeID         string
	Code         string
	FullName         string
	Rate         string
	Notes         string
	StartDate         string
	EndDate         string
	SYSCreated         string
	SYSCreatedBy         string
	SYSCreatedHost         string
	SYSUpdated         string
	SYSUpdatedBy         string
	SYSUpdatedHost         string
	SYSDeleted         string
	SYSDeletedBy         string
	SYSDeletedHost         string
	SYSActivity         string
	Currency         string
	SYSDbVersion         string
	Comments         string
	ProjectManager         string
	AccountManager         string
	NoActiveProjects         string
	/// Field Properties
	SYSId_props     FieldProperties
	OriginID_props     FieldProperties
	StateID_props     FieldProperties
	DocTypeID_props     FieldProperties
	Code_props     FieldProperties
	FullName_props     FieldProperties
	Rate_props     FieldProperties
	Notes_props     FieldProperties
	StartDate_props     FieldProperties
	EndDate_props     FieldProperties
	SYSCreated_props     FieldProperties
	SYSCreatedBy_props     FieldProperties
	SYSCreatedHost_props     FieldProperties
	SYSUpdated_props     FieldProperties
	SYSUpdatedBy_props     FieldProperties
	SYSUpdatedHost_props     FieldProperties
	SYSDeleted_props     FieldProperties
	SYSDeletedBy_props     FieldProperties
	SYSDeletedHost_props     FieldProperties
	SYSActivity_props     FieldProperties
	Currency_props     FieldProperties
	SYSDbVersion_props     FieldProperties
	Comments_props     FieldProperties
	ProjectManager_props     FieldProperties
	AccountManager_props     FieldProperties
	NoActiveProjects_props     FieldProperties
	/// Lookups
	StateID_lookup    []Lookup_Item
	DocTypeID_lookup    []Lookup_Item
	Currency_lookup    []Lookup_Item
	ProjectManager_lookup    []Lookup_Item
	AccountManager_lookup    []Lookup_Item
	
	/// END OF DEFINITIONS
	///
}