package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/resource.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Resource (resource)
// Endpoint 	        : Resource (ResourceID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 30/12/2022 at 12:18:22
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Resource defines the datamolde for the Resource object
type Resource struct {


SYSId       string
SYSId_props FieldProperties
ResourceID       string
ResourceID_props FieldProperties
Code       string
Code_props FieldProperties
Name       string
Name_props FieldProperties
Email       string
Email_props FieldProperties
Class       string
Class_props FieldProperties
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
UserActive       string
UserActive_props FieldProperties
SYSActivity       string
SYSActivity_props FieldProperties
Notes       string
Notes_props FieldProperties
SYSDbVersion       string
SYSDbVersion_props FieldProperties
Comments       string
Comments_props FieldProperties
 // Any lookups will be added below




Class_lookup []Lookup_Item









UserActive_lookup []Lookup_Item





}

const (
	Resource_Title       = "Resource"
	Resource_SQLTable    = "resourceStore"
	Resource_SQLSearchID = "resourceID"
	Resource_QueryString = "ResourceID"
	///
	/// Handler Defintions
	///
	Resource_Template     = "Resource"
	Resource_TemplateList = "/Resource/Resource_List"
	Resource_TemplateView = "/Resource/Resource_View"
	Resource_TemplateEdit = "/Resource/Resource_Edit"
	Resource_TemplateNew  = "/Resource/Resource_New"
	///
	/// Handler Monitor Paths
	///
	Resource_Path       = "/API/Resource/"
	Resource_PathList   = "/ResourceList/"
	Resource_PathView   = "/ResourceView/"
	Resource_PathEdit   = "/ResourceEdit/"
	Resource_PathNew    = "/ResourceNew/"
	Resource_PathSave   = "/ResourceSave/"
	Resource_PathDelete = "/ResourceDelete/"
	///
	//Resource_Redirect provides a page to return to aftern an action
	Resource_Redirect = Resource_PathList
	
	///
	///
	/// SQL Field Definitions
	///
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

	/// Definitions End
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

	/// Definitions End
	///
)

//resource_PageList provides the information for the template for a list of Resources
type Resource_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []Resource
	Context	 appContext
}

//resource_Page provides the information for the template for an individual Resource
type Resource_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	SYSId_props     FieldProperties
	ResourceID         string
	ResourceID_props     FieldProperties
	Code         string
	Code_props     FieldProperties
	Name         string
	Name_props     FieldProperties
	Email         string
	Email_props     FieldProperties
	Class         string
	Class_lookup    []Lookup_Item
	Class_props     FieldProperties
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
	UserActive         string
	UserActive_lookup    []Lookup_Item
	UserActive_props     FieldProperties
	SYSActivity         string
	SYSActivity_props     FieldProperties
	Notes         string
	Notes_props     FieldProperties
	SYSDbVersion         string
	SYSDbVersion_props     FieldProperties
	Comments         string
	Comments_props     FieldProperties
	// 
	// Dynamically generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}