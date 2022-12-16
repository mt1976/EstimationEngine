package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/originstate.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : OriginState (originstate)
// Endpoint 	        : OriginState (OriginStateID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 16/12/2022 at 16:47:11
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//OriginState defines the datamolde for the OriginState object
type OriginState struct {


SYSId       string
SYSId_props FieldProperties
OriginStateID       string
OriginStateID_props FieldProperties
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
 // Any lookups will be added below














}

const (
	OriginState_Title       = "Origin State"
	OriginState_SQLTable    = "originStateStore"
	OriginState_SQLSearchID = "originStateID"
	OriginState_QueryString = "OriginStateID"
	///
	/// Handler Defintions
	///
	OriginState_Template     = "OriginState"
	OriginState_TemplateList = "/OriginState/OriginState_List"
	OriginState_TemplateView = "/OriginState/OriginState_View"
	OriginState_TemplateEdit = "/OriginState/OriginState_Edit"
	OriginState_TemplateNew  = "/OriginState/OriginState_New"
	///
	/// Handler Monitor Paths
	///
	OriginState_Path       = "/API/OriginState/"
	OriginState_PathList   = "/OriginStateList/"
	OriginState_PathView   = "/OriginStateView/"
	OriginState_PathEdit   = "/OriginStateEdit/"
	OriginState_PathNew    = "/OriginStateNew/"
	OriginState_PathSave   = "/OriginStateSave/"
	OriginState_PathDelete = "/OriginStateDelete/"
	///
	//OriginState_Redirect provides a page to return to aftern an action
	OriginState_Redirect = OriginState_PathList
	
	///
	///
	/// SQL Field Definitions
	///
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

	/// Definitions End
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

	/// Definitions End
	///
)

//originstate_PageList provides the information for the template for a list of OriginStates
type OriginState_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []OriginState
	Context	 appContext
}

//originstate_Page provides the information for the template for an individual OriginState
type OriginState_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 16/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	SYSId_props     FieldProperties
	OriginStateID         string
	OriginStateID_props     FieldProperties
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
	// 
	// Dynamically generated 16/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}