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
// Date & Time		    : 07/02/2023 at 18:52:38
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//ProjectState defines the datamolde for the ProjectState object
type ProjectState struct {


SYSId       string
SYSId_props FieldProperties
ProjectStateID       string
ProjectStateID_props FieldProperties
Code       string
Code_props FieldProperties
Name       string
Name_props FieldProperties
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
SYSDbVersion       string
SYSDbVersion_props FieldProperties
IsLocked       string
IsLocked_props FieldProperties
Notify       string
Notify_props FieldProperties
 // Any lookups will be added below













IsLocked_lookup []Lookup_Item
Notify_lookup []Lookup_Item

}

const (
	ProjectState_Title       = "Project State"
	ProjectState_SQLTable    = "projectStateStore"
	ProjectState_SQLSearchID = "projectStateID"
	ProjectState_QueryString = "ProjectStateID"
	///
	/// Handler Defintions
	///
	ProjectState_Template     = "ProjectState"
	ProjectState_TemplateList = "/ProjectState/ProjectState_List"
	ProjectState_TemplateView = "/ProjectState/ProjectState_View"
	ProjectState_TemplateEdit = "/ProjectState/ProjectState_Edit"
	ProjectState_TemplateNew  = "/ProjectState/ProjectState_New"
	///
	/// Handler Monitor Paths
	///
	ProjectState_Path       = "/API/ProjectState/"
	ProjectState_PathList   = "/ProjectStateList/"
	ProjectState_PathView   = "/ProjectStateView/"
	ProjectState_PathEdit   = "/ProjectStateEdit/"
	ProjectState_PathNew    = "/ProjectStateNew/"
	ProjectState_PathSave   = "/ProjectStateSave/"
	ProjectState_PathDelete = "/ProjectStateDelete/"
	///
	//ProjectState_Redirect provides a page to return to aftern an action
	ProjectState_Redirect = ProjectState_PathList
	
	///
	///
	/// SQL Field Definitions
	///
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

	/// Definitions End
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

	/// Definitions End
	///
)

//projectstate_PageList provides the information for the template for a list of ProjectStates
type ProjectState_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []ProjectState
	Context	 appContext
	BlockEdit		 bool
	BlockSave 		 bool
	BlockCreate		 bool
	BlockDelete		 bool
	BlockValidate	 bool
	BlockView		 bool
}

//projectstate_Page provides the information for the template for an individual ProjectState
type ProjectState_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	SYSId_props     FieldProperties
	ProjectStateID         string
	ProjectStateID_props     FieldProperties
	Code         string
	Code_props     FieldProperties
	Name         string
	Name_props     FieldProperties
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
	SYSDbVersion         string
	SYSDbVersion_props     FieldProperties
	IsLocked         string
	IsLocked_lookup    []Lookup_Item
	IsLocked_props     FieldProperties
	Notify         string
	Notify_lookup    []Lookup_Item
	Notify_props     FieldProperties
	// 
	// Dynamically generated 07/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 		 appContext
	BlockEdit		 bool
	BlockSave 		 bool
	BlockCreate		 bool
	BlockDelete		 bool
	BlockValidate	 bool
	BlockView		 bool
}