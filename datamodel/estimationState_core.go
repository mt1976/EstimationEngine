package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/estimationstate.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : EstimationState (estimationstate)
// Endpoint 	        : EstimationState (EstimationStateID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 08/12/2022 at 13:31:30
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//EstimationState defines the datamolde for the EstimationState object
type EstimationState struct {


SYSId       string
SYSId_props FieldProperties
EstimationStateID       string
EstimationStateID_props FieldProperties
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
 // Any lookups will be added below













}

const (
	EstimationState_Title       = "Estimation State"
	EstimationState_SQLTable    = "estimationStateStore"
	EstimationState_SQLSearchID = "estimationStateID"
	EstimationState_QueryString = "EstimationStateID"
	///
	/// Handler Defintions
	///
	EstimationState_Template     = "EstimationState"
	EstimationState_TemplateList = "/EstimationState/EstimationState_List"
	EstimationState_TemplateView = "/EstimationState/EstimationState_View"
	EstimationState_TemplateEdit = "/EstimationState/EstimationState_Edit"
	EstimationState_TemplateNew  = "/EstimationState/EstimationState_New"
	///
	/// Handler Monitor Paths
	///
	EstimationState_Path       = "/API/EstimationState/"
	EstimationState_PathList   = "/EstimationStateList/"
	EstimationState_PathView   = "/EstimationStateView/"
	EstimationState_PathEdit   = "/EstimationStateEdit/"
	EstimationState_PathNew    = "/EstimationStateNew/"
	EstimationState_PathSave   = "/EstimationStateSave/"
	EstimationState_PathDelete = "/EstimationStateDelete/"
	///
	//EstimationState_Redirect provides a page to return to aftern an action
	EstimationState_Redirect = EstimationState_PathList
	
	///
	///
	/// SQL Field Definitions
	///
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

	/// Definitions End
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

	/// Definitions End
	///
)

//estimationstate_PageList provides the information for the template for a list of EstimationStates
type EstimationState_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []EstimationState
	Context	 appContext
}

//estimationstate_Page provides the information for the template for an individual EstimationState
type EstimationState_Page struct {
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
	EstimationStateID         string
	EstimationStateID_props     FieldProperties
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
	// 
	// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}