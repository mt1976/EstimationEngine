package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/resource.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Resource (resource)
// Endpoint 	        : Resource (ResourceID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 04/03/2023 at 20:14:14
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Resource defines the datamodel for the Resource object
type Resource struct {
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	SYSId       string
	ResourceID       string
	Code       string
	Name       string
	Email       string
	Class       string
	SYSCreated       string
	SYSCreatedBy       string
	SYSCreatedHost       string
	SYSUpdated       string
	SYSUpdatedBy       string
	SYSUpdatedHost       string
	SYSDeleted       string
	SYSDeletedBy       string
	SYSDeletedHost       string
	UserActive       string
	SYSActivity       string
	Notes       string
	SYSDbVersion       string
	Comments       string
	//
	// Field Properties
	//
	SYSId_props FieldProperties
	ResourceID_props FieldProperties
	Code_props FieldProperties
	Name_props FieldProperties
	Email_props FieldProperties
	Class_props FieldProperties
	SYSCreated_props FieldProperties
	SYSCreatedBy_props FieldProperties
	SYSCreatedHost_props FieldProperties
	SYSUpdated_props FieldProperties
	SYSUpdatedBy_props FieldProperties
	SYSUpdatedHost_props FieldProperties
	SYSDeleted_props FieldProperties
	SYSDeletedBy_props FieldProperties
	SYSDeletedHost_props FieldProperties
	UserActive_props FieldProperties
	SYSActivity_props FieldProperties
	Notes_props FieldProperties
	SYSDbVersion_props FieldProperties
	Comments_props FieldProperties
	//
 	// Any lookups will be added below
	//
	Class_lookup []Lookup_Item
	UserActive_lookup []Lookup_Item
	}

const (
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// Page Infrastructure
	Resource_Name      = "Resource"
	Resource_Title       = "Resource"
	Resource_SQLTable    = "resourceStore"
	Resource_SQLSearchID = "resourceID"
	Resource_QueryString = "ResourceID"
	///
	/// Template Path Defintions
	///
	Resource_Template     = "Resource"
	Resource_TemplateList = "/Resource/ResourceList"
	Resource_TemplateView = "/Resource/ResourceView"
	Resource_TemplateEdit = "/Resource/ResourceEdit"
	Resource_TemplateNew  = "/Resource/ResourceNew"
	///
	/// URI Handler Paths
	///
	Resource_Path       = "/API/Resource/"
	Resource_PathList   = "/ResourceList/"
	Resource_PathView   = "/ResourceView/"
	Resource_PathEdit   = "/ResourceEdit/"
	Resource_PathNew    = "/ResourceNew/"
	Resource_PathSave   = "/ResourceSave/"
	Resource_PathDelete = "/ResourceDelete/"
	// Redirects - On Server Side Error
	Resource_PathEditException   = "/ResourceEditException/"
	Resource_PathNewException    = "/ResourceNewException/"
	//
	//Resource_Redirect provides a page to return to aftern an action
	Resource_Redirect = Resource_PathList
	
	//
	//
	// SQL Field Definitions
	//
	Resource_SYSId_sql   = "_id" // SYSId is a Int
	Resource_ResourceID_sql   = "resourceID" // ResourceID is a String
	Resource_Code_sql   = "code" // Code is a String
	Resource_Name_sql   = "name" // Name is a String
	Resource_Email_sql   = "email" // Email is a String
	Resource_Class_sql   = "class" // Class is a String
	Resource_SYSCreated_sql   = "_created" // SYSCreated is a String
	Resource_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	Resource_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	Resource_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	Resource_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
	Resource_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
	Resource_SYSDeleted_sql   = "_deleted" // SYSDeleted is a String
	Resource_SYSDeletedBy_sql   = "_deletedBy" // SYSDeletedBy is a String
	Resource_SYSDeletedHost_sql   = "_deletedHost" // SYSDeletedHost is a String
	Resource_UserActive_sql   = "userActive" // UserActive is a String
	Resource_SYSActivity_sql   = "_activity" // SYSActivity is a String
	Resource_Notes_sql   = "notes" // Notes is a String
	Resource_SYSDbVersion_sql   = "_dbVersion" // SYSDbVersion is a String
	Resource_Comments_sql   = "comments" // Comments is a String
///
	/// Application Field Definitions
	///
	Resource_SYSId_scrn   = "SYSId" // SYSId is a Int
	Resource_ResourceID_scrn   = "ResourceID" // ResourceID is a String
	Resource_Code_scrn   = "Code" // Code is a String
	Resource_Name_scrn   = "Name" // Name is a String
	Resource_Email_scrn   = "Email" // Email is a String
	Resource_Class_scrn   = "Class" // Class is a String
	Resource_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	Resource_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	Resource_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	Resource_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	Resource_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
	Resource_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
	Resource_SYSDeleted_scrn   = "SYSDeleted" // SYSDeleted is a String
	Resource_SYSDeletedBy_scrn   = "SYSDeletedBy" // SYSDeletedBy is a String
	Resource_SYSDeletedHost_scrn   = "SYSDeletedHost" // SYSDeletedHost is a String
	Resource_UserActive_scrn   = "UserActive" // UserActive is a String
	Resource_SYSActivity_scrn   = "SYSActivity" // SYSActivity is a String
	Resource_Notes_scrn   = "Notes" // Notes is a String
	Resource_SYSDbVersion_scrn   = "SYSDbVersion" // SYSDbVersion is a String
	Resource_Comments_scrn   = "Comments" // Comments is a String
///
)

//Resource_PageList provides the information for the template for a list of Resources
type Resource_PageList struct {
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	ItemList  		 []Resource
}

//Resource_Page provides the information for the template for an individual Resource
type Resource_Page struct {
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	ResourceID         string
	Code         string
	Name         string
	Email         string
	Class         string
	SYSCreated         string
	SYSCreatedBy         string
	SYSCreatedHost         string
	SYSUpdated         string
	SYSUpdatedBy         string
	SYSUpdatedHost         string
	SYSDeleted         string
	SYSDeletedBy         string
	SYSDeletedHost         string
	UserActive         string
	SYSActivity         string
	Notes         string
	SYSDbVersion         string
	Comments         string
	/// Field Properties
	SYSId_props     FieldProperties
	ResourceID_props     FieldProperties
	Code_props     FieldProperties
	Name_props     FieldProperties
	Email_props     FieldProperties
	Class_props     FieldProperties
	SYSCreated_props     FieldProperties
	SYSCreatedBy_props     FieldProperties
	SYSCreatedHost_props     FieldProperties
	SYSUpdated_props     FieldProperties
	SYSUpdatedBy_props     FieldProperties
	SYSUpdatedHost_props     FieldProperties
	SYSDeleted_props     FieldProperties
	SYSDeletedBy_props     FieldProperties
	SYSDeletedHost_props     FieldProperties
	UserActive_props     FieldProperties
	SYSActivity_props     FieldProperties
	Notes_props     FieldProperties
	SYSDbVersion_props     FieldProperties
	Comments_props     FieldProperties
	/// Lookups
	Class_lookup    []Lookup_Item
	UserActive_lookup    []Lookup_Item
	
	/// END OF DEFINITIONS
	///
}