package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/originstate.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : OriginState (originstate)
// Endpoint 	        : OriginState (OriginStateID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 04/03/2023 at 20:14:13
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//OriginState defines the datamodel for the OriginState object
type OriginState struct {
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	SYSId       string
	OriginStateID       string
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
	//
	// Field Properties
	//
	SYSId_props FieldProperties
	OriginStateID_props FieldProperties
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
	//
 	// Any lookups will be added below
	//
	}

const (
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// Page Infrastructure
	OriginState_Name      = "OriginState"
	OriginState_Title       = "Origin State"
	OriginState_SQLTable    = "originStateStore"
	OriginState_SQLSearchID = "originStateID"
	OriginState_QueryString = "OriginStateID"
	///
	/// Template Path Defintions
	///
	OriginState_Template     = "OriginState"
	OriginState_TemplateList = "/OriginState/OriginStateList"
	OriginState_TemplateView = "/OriginState/OriginStateView"
	OriginState_TemplateEdit = "/OriginState/OriginStateEdit"
	OriginState_TemplateNew  = "/OriginState/OriginStateNew"
	///
	/// URI Handler Paths
	///
	OriginState_Path       = "/API/OriginState/"
	OriginState_PathList   = "/OriginStateList/"
	OriginState_PathView   = "/OriginStateView/"
	OriginState_PathEdit   = "/OriginStateEdit/"
	OriginState_PathNew    = "/OriginStateNew/"
	OriginState_PathSave   = "/OriginStateSave/"
	OriginState_PathDelete = "/OriginStateDelete/"
	// Redirects - On Server Side Error
	OriginState_PathEditException   = "/OriginStateEditException/"
	OriginState_PathNewException    = "/OriginStateNewException/"
	//
	//OriginState_Redirect provides a page to return to aftern an action
	OriginState_Redirect = OriginState_PathList
	
	//
	//
	// SQL Field Definitions
	//
	OriginState_SYSId_sql   = "_id" // SYSId is a Int
	OriginState_OriginStateID_sql   = "originStateID" // OriginStateID is a String
	OriginState_Code_sql   = "code" // Code is a String
	OriginState_Name_sql   = "name" // Name is a String
	OriginState_SYSCreated_sql   = "_created" // SYSCreated is a String
	OriginState_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	OriginState_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	OriginState_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	OriginState_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
	OriginState_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
	OriginState_SYSDeleted_sql   = "_deleted" // SYSDeleted is a String
	OriginState_SYSDeletedBy_sql   = "_deletedBy" // SYSDeletedBy is a String
	OriginState_SYSDeletedHost_sql   = "_deletedHost" // SYSDeletedHost is a String
	OriginState_SYSDbVersion_sql   = "_dbVersion" // SYSDbVersion is a String
	OriginState_IsLocked_sql   = "isLocked" // IsLocked is a String
	OriginState_Notify_sql   = "notify" // Notify is a String
///
	/// Application Field Definitions
	///
	OriginState_SYSId_scrn   = "SYSId" // SYSId is a Int
	OriginState_OriginStateID_scrn   = "OriginStateID" // OriginStateID is a String
	OriginState_Code_scrn   = "Code" // Code is a String
	OriginState_Name_scrn   = "Name" // Name is a String
	OriginState_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	OriginState_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	OriginState_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	OriginState_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	OriginState_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
	OriginState_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
	OriginState_SYSDeleted_scrn   = "SYSDeleted" // SYSDeleted is a String
	OriginState_SYSDeletedBy_scrn   = "SYSDeletedBy" // SYSDeletedBy is a String
	OriginState_SYSDeletedHost_scrn   = "SYSDeletedHost" // SYSDeletedHost is a String
	OriginState_SYSDbVersion_scrn   = "SYSDbVersion" // SYSDbVersion is a String
	OriginState_IsLocked_scrn   = "IsLocked" // IsLocked is a String
	OriginState_Notify_scrn   = "Notify" // Notify is a String
///
)

//OriginState_PageList provides the information for the template for a list of OriginStates
type OriginState_PageList struct {
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
	ItemList  		 []OriginState
}

//OriginState_Page provides the information for the template for an individual OriginState
type OriginState_Page struct {
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
	OriginStateID         string
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
	/// Field Properties
	SYSId_props     FieldProperties
	OriginStateID_props     FieldProperties
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
	/// Lookups
	
	/// END OF DEFINITIONS
	///
}