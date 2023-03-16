package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/estimationstate.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : EstimationState (estimationstate)
// Endpoint 	        : EstimationState (EstimationStateID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/03/2023 at 19:24:48
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//EstimationState defines the datamodel for the EstimationState object
type EstimationState struct {
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	SYSId       string
	EstimationStateID       string
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
	EstimationStateID_props FieldProperties
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
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// Page Infrastructure
	EstimationState_Name      = "EstimationState"
	EstimationState_Title       = "Estimation State"
	EstimationState_SQLTable    = "estimationStateStore"
	EstimationState_SQLSearchID = "estimationStateID"
	EstimationState_QueryString = "EstimationStateID"
	///
	/// Template Path Defintions
	///
	EstimationState_Template     = "EstimationState"
	EstimationState_TemplateList = "/EstimationState/EstimationStateList"
	EstimationState_TemplateView = "/EstimationState/EstimationStateView"
	EstimationState_TemplateEdit = "/EstimationState/EstimationStateEdit"
	EstimationState_TemplateNew  = "/EstimationState/EstimationStateNew"
	///
	/// URI Handler Paths
	///
	EstimationState_Path       = "/API/EstimationState/"
	EstimationState_PathList   = "/EstimationStateList/"
	EstimationState_PathView   = "/EstimationStateView/"
	EstimationState_PathEdit   = "/EstimationStateEdit/"
	EstimationState_PathNew    = "/EstimationStateNew/"
	EstimationState_PathSave   = "/EstimationStateSave/"
	EstimationState_PathDelete = "/EstimationStateDelete/"
	// Redirects - On Server Side Error
	EstimationState_PathEditException   = "/EstimationStateEditException/"
	EstimationState_PathNewException    = "/EstimationStateNewException/"
	//
	//EstimationState_Redirect provides a page to return to aftern an action
	EstimationState_Redirect = EstimationState_PathList
	
	//
	//
	// SQL Field Definitions
	//
	EstimationState_SYSId_sql   = "_id" // SYSId is a Int
	EstimationState_EstimationStateID_sql   = "estimationStateID" // EstimationStateID is a String
	EstimationState_Code_sql   = "code" // Code is a String
	EstimationState_Name_sql   = "name" // Name is a String
	EstimationState_SYSCreated_sql   = "_created" // SYSCreated is a String
	EstimationState_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	EstimationState_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	EstimationState_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	EstimationState_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
	EstimationState_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
	EstimationState_SYSDeleted_sql   = "_deleted" // SYSDeleted is a String
	EstimationState_SYSDeletedBy_sql   = "_deletedBy" // SYSDeletedBy is a String
	EstimationState_SYSDeletedHost_sql   = "_deletedHost" // SYSDeletedHost is a String
	EstimationState_SYSDbVersion_sql   = "_dbVersion" // SYSDbVersion is a String
	EstimationState_IsLocked_sql   = "isLocked" // IsLocked is a String
	EstimationState_Notify_sql   = "notify" // Notify is a String
	EstimationState_Migrate_sql   = "migrate" // Migrate is a String
///
	/// Application Field Definitions
	///
	EstimationState_SYSId_scrn   = "SYSId" // SYSId is a Int
	EstimationState_EstimationStateID_scrn   = "EstimationStateID" // EstimationStateID is a String
	EstimationState_Code_scrn   = "Code" // Code is a String
	EstimationState_Name_scrn   = "Name" // Name is a String
	EstimationState_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	EstimationState_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	EstimationState_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	EstimationState_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	EstimationState_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
	EstimationState_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
	EstimationState_SYSDeleted_scrn   = "SYSDeleted" // SYSDeleted is a String
	EstimationState_SYSDeletedBy_scrn   = "SYSDeletedBy" // SYSDeletedBy is a String
	EstimationState_SYSDeletedHost_scrn   = "SYSDeletedHost" // SYSDeletedHost is a String
	EstimationState_SYSDbVersion_scrn   = "SYSDbVersion" // SYSDbVersion is a String
	EstimationState_IsLocked_scrn   = "IsLocked" // IsLocked is a String
	EstimationState_Notify_scrn   = "Notify" // Notify is a String
	EstimationState_Migrate_scrn   = "Migrate" // Migrate is a String
///
)

//EstimationState_PageList provides the information for the template for a list of EstimationStates
type EstimationState_PageList struct {
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	ItemList  		 []EstimationState
}

//EstimationState_Page provides the information for the template for an individual EstimationState
type EstimationState_Page struct {
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	EstimationStateID         string
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
	EstimationStateID_props     FieldProperties
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