package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/projectstate.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : ProjectState (projectstate)
// Endpoint 	        : ProjectState (ProjectStateID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 13/03/2023 at 14:22:30
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//ProjectState defines the datamodel for the ProjectState object
type ProjectState struct {
	// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	SYSId       string
	ProjectStateID       string
	Code       string
	Name       string
	SYSCreated       string
	SYSCreatedBy       string
	SYSCreatedHost       string
	SYSUpdated       string
	SYSUpdatedBy       string
	SYSUpdatedHost       string
	SYSDeleted       string
	SYSDeletedBy       string
	SYSDeletedHost       string
	SYSDbVersion       string
	IsLocked       string
	Notify       string
	Migrate       string
	//
	// Field Properties
	//
	SYSId_props FieldProperties
	ProjectStateID_props FieldProperties
	Code_props FieldProperties
	Name_props FieldProperties
	SYSCreated_props FieldProperties
	SYSCreatedBy_props FieldProperties
	SYSCreatedHost_props FieldProperties
	SYSUpdated_props FieldProperties
	SYSUpdatedBy_props FieldProperties
	SYSUpdatedHost_props FieldProperties
	SYSDeleted_props FieldProperties
	SYSDeletedBy_props FieldProperties
	SYSDeletedHost_props FieldProperties
	SYSDbVersion_props FieldProperties
	IsLocked_props FieldProperties
	Notify_props FieldProperties
	Migrate_props FieldProperties
	//
 	// Any lookups will be added below
	//
	IsLocked_lookup []Lookup_Item
	Notify_lookup []Lookup_Item
	Migrate_lookup []Lookup_Item
	}

const (
	// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// Page Infrastructure
	ProjectState_Name      = "ProjectState"
	ProjectState_Title       = "Project State"
	ProjectState_SQLTable    = "projectStateStore"
	ProjectState_SQLSearchID = "projectStateID"
	ProjectState_QueryString = "ProjectStateID"
	///
	/// Template Path Defintions
	///
	ProjectState_Template     = "ProjectState"
	ProjectState_TemplateList = "/ProjectState/ProjectStateList"
	ProjectState_TemplateView = "/ProjectState/ProjectStateView"
	ProjectState_TemplateEdit = "/ProjectState/ProjectStateEdit"
	ProjectState_TemplateNew  = "/ProjectState/ProjectStateNew"
	///
	/// URI Handler Paths
	///
	ProjectState_Path       = "/API/ProjectState/"
	ProjectState_PathList   = "/ProjectStateList/"
	ProjectState_PathView   = "/ProjectStateView/"
	ProjectState_PathEdit   = "/ProjectStateEdit/"
	ProjectState_PathNew    = "/ProjectStateNew/"
	ProjectState_PathSave   = "/ProjectStateSave/"
	ProjectState_PathDelete = "/ProjectStateDelete/"
	// Redirects - On Server Side Error
	ProjectState_PathEditException   = "/ProjectStateEditException/"
	ProjectState_PathNewException    = "/ProjectStateNewException/"
	//
	//ProjectState_Redirect provides a page to return to aftern an action
	ProjectState_Redirect = ProjectState_PathList
	
	//
	//
	// SQL Field Definitions
	//
	ProjectState_SYSId_sql   = "_id" // SYSId is a Int
	ProjectState_ProjectStateID_sql   = "projectStateID" // ProjectStateID is a String
	ProjectState_Code_sql   = "code" // Code is a String
	ProjectState_Name_sql   = "name" // Name is a String
	ProjectState_SYSCreated_sql   = "_created" // SYSCreated is a String
	ProjectState_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	ProjectState_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	ProjectState_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	ProjectState_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
	ProjectState_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
	ProjectState_SYSDeleted_sql   = "_deleted" // SYSDeleted is a String
	ProjectState_SYSDeletedBy_sql   = "_deletedBy" // SYSDeletedBy is a String
	ProjectState_SYSDeletedHost_sql   = "_deletedHost" // SYSDeletedHost is a String
	ProjectState_SYSDbVersion_sql   = "_dbVersion" // SYSDbVersion is a String
	ProjectState_IsLocked_sql   = "isLocked" // IsLocked is a String
	ProjectState_Notify_sql   = "notify" // Notify is a String
	ProjectState_Migrate_sql   = "migrate" // Migrate is a String
///
	/// Application Field Definitions
	///
	ProjectState_SYSId_scrn   = "SYSId" // SYSId is a Int
	ProjectState_ProjectStateID_scrn   = "ProjectStateID" // ProjectStateID is a String
	ProjectState_Code_scrn   = "Code" // Code is a String
	ProjectState_Name_scrn   = "Name" // Name is a String
	ProjectState_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	ProjectState_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	ProjectState_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	ProjectState_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	ProjectState_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
	ProjectState_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
	ProjectState_SYSDeleted_scrn   = "SYSDeleted" // SYSDeleted is a String
	ProjectState_SYSDeletedBy_scrn   = "SYSDeletedBy" // SYSDeletedBy is a String
	ProjectState_SYSDeletedHost_scrn   = "SYSDeletedHost" // SYSDeletedHost is a String
	ProjectState_SYSDbVersion_scrn   = "SYSDbVersion" // SYSDbVersion is a String
	ProjectState_IsLocked_scrn   = "IsLocked" // IsLocked is a String
	ProjectState_Notify_scrn   = "Notify" // Notify is a String
	ProjectState_Migrate_scrn   = "Migrate" // Migrate is a String
///
)

//ProjectState_PageList provides the information for the template for a list of ProjectStates
type ProjectState_PageList struct {
	// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	ItemList  		 []ProjectState
}

//ProjectState_Page provides the information for the template for an individual ProjectState
type ProjectState_Page struct {
	// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	ProjectStateID         string
	Code         string
	Name         string
	SYSCreated         string
	SYSCreatedBy         string
	SYSCreatedHost         string
	SYSUpdated         string
	SYSUpdatedBy         string
	SYSUpdatedHost         string
	SYSDeleted         string
	SYSDeletedBy         string
	SYSDeletedHost         string
	SYSDbVersion         string
	IsLocked         string
	Notify         string
	Migrate         string
	/// Field Properties
	SYSId_props     FieldProperties
	ProjectStateID_props     FieldProperties
	Code_props     FieldProperties
	Name_props     FieldProperties
	SYSCreated_props     FieldProperties
	SYSCreatedBy_props     FieldProperties
	SYSCreatedHost_props     FieldProperties
	SYSUpdated_props     FieldProperties
	SYSUpdatedBy_props     FieldProperties
	SYSUpdatedHost_props     FieldProperties
	SYSDeleted_props     FieldProperties
	SYSDeletedBy_props     FieldProperties
	SYSDeletedHost_props     FieldProperties
	SYSDbVersion_props     FieldProperties
	IsLocked_props     FieldProperties
	Notify_props     FieldProperties
	Migrate_props     FieldProperties
	/// Lookups
	IsLocked_lookup    []Lookup_Item
	Notify_lookup    []Lookup_Item
	Migrate_lookup    []Lookup_Item
	
	/// END OF DEFINITIONS
	///
}