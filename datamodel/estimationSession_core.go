package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/estimationsession.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : EstimationSession (estimationsession)
// Endpoint 	        : EstimationSession (EstimationSessionID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 27/11/2022 at 20:46:13
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//EstimationSession defines the datamolde for the EstimationSession object
type EstimationSession struct {


SYSId       string
SYSId_props FieldProperties
EstimationSessionID       string
EstimationSessionID_props FieldProperties
ProjectID       string
ProjectID_props FieldProperties
EstimationStateID       string
EstimationStateID_props FieldProperties
Notes       string
Notes_props FieldProperties
Releases       string
Releases_props FieldProperties
Total       string
Total_props FieldProperties
Contingency       string
Contingency_props FieldProperties
ReqDays       string
ReqDays_props FieldProperties
RegCost       string
RegCost_props FieldProperties
ImpDays       string
ImpDays_props FieldProperties
ImpCost       string
ImpCost_props FieldProperties
UatDays       string
UatDays_props FieldProperties
UatCost       string
UatCost_props FieldProperties
MgtDays       string
MgtDays_props FieldProperties
MgtCost       string
MgtCost_props FieldProperties
RelDays       string
RelDays_props FieldProperties
RelCost       string
RelCost_props FieldProperties
SupportUplift       string
SupportUplift_props FieldProperties
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
Name       string
Name_props FieldProperties
 // Any lookups will be added below

ProjectID_lookup []Lookup_Item
EstimationStateID_lookup []Lookup_Item


























}

const (
	EstimationSession_Title       = "Estimation Session"
	EstimationSession_SQLTable    = "estimationSessionStore"
	EstimationSession_SQLSearchID = "estimationSessionID"
	EstimationSession_QueryString = "EstimationSessionID"
	///
	/// Handler Defintions
	///
	EstimationSession_Template     = "EstimationSession"
	EstimationSession_TemplateList = "/EstimationSession/EstimationSession_List"
	EstimationSession_TemplateView = "/EstimationSession/EstimationSession_View"
	EstimationSession_TemplateEdit = "/EstimationSession/EstimationSession_Edit"
	EstimationSession_TemplateNew  = "/EstimationSession/EstimationSession_New"
	///
	/// Handler Monitor Paths
	///
	EstimationSession_Path       = "/API/EstimationSession/"
	EstimationSession_PathList   = "/EstimationSessionList/"
	EstimationSession_PathView   = "/EstimationSessionView/"
	EstimationSession_PathEdit   = "/EstimationSessionEdit/"
	EstimationSession_PathNew    = "/EstimationSessionNew/"
	EstimationSession_PathSave   = "/EstimationSessionSave/"
	EstimationSession_PathDelete = "/EstimationSessionDelete/"
	///
	//EstimationSession_Redirect provides a page to return to aftern an action
	EstimationSession_Redirect = EstimationSession_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	EstimationSession_SYSId_sql   = "_id" // SYSId is a Int
	EstimationSession_EstimationSessionID_sql   = "estimationSessionID" // EstimationSessionID is a String
	EstimationSession_ProjectID_sql   = "projectID" // ProjectID is a String
	EstimationSession_EstimationStateID_sql   = "estimationStateID" // EstimationStateID is a String
	EstimationSession_Notes_sql   = "notes" // Notes is a String
	EstimationSession_Releases_sql   = "releases" // Releases is a String
	EstimationSession_Total_sql   = "total" // Total is a String
	EstimationSession_Contingency_sql   = "contingency" // Contingency is a String
	EstimationSession_ReqDays_sql   = "reqDays" // ReqDays is a String
	EstimationSession_RegCost_sql   = "regCost" // RegCost is a String
	EstimationSession_ImpDays_sql   = "impDays" // ImpDays is a String
	EstimationSession_ImpCost_sql   = "impCost" // ImpCost is a String
	EstimationSession_UatDays_sql   = "uatDays" // UatDays is a String
	EstimationSession_UatCost_sql   = "uatCost" // UatCost is a String
	EstimationSession_MgtDays_sql   = "mgtDays" // MgtDays is a String
	EstimationSession_MgtCost_sql   = "mgtCost" // MgtCost is a String
	EstimationSession_RelDays_sql   = "relDays" // RelDays is a String
	EstimationSession_RelCost_sql   = "relCost" // RelCost is a String
	EstimationSession_SupportUplift_sql   = "supportUplift" // SupportUplift is a String
	EstimationSession_SYSCreated_sql   = "_created" // SYSCreated is a String
	EstimationSession_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	EstimationSession_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	EstimationSession_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	EstimationSession_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
	EstimationSession_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
	EstimationSession_SYSDeleted_sql   = "_deleted" // SYSDeleted is a String
	EstimationSession_SYSDeletedBy_sql   = "_deletedBy" // SYSDeletedBy is a String
	EstimationSession_SYSDeletedHost_sql   = "_deletedHost" // SYSDeletedHost is a String
	EstimationSession_Name_sql   = "name" // Name is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	EstimationSession_SYSId_scrn   = "SYSId" // SYSId is a Int
	EstimationSession_EstimationSessionID_scrn   = "EstimationSessionID" // EstimationSessionID is a String
	EstimationSession_ProjectID_scrn   = "ProjectID" // ProjectID is a String
	EstimationSession_EstimationStateID_scrn   = "EstimationStateID" // EstimationStateID is a String
	EstimationSession_Notes_scrn   = "Notes" // Notes is a String
	EstimationSession_Releases_scrn   = "Releases" // Releases is a String
	EstimationSession_Total_scrn   = "Total" // Total is a String
	EstimationSession_Contingency_scrn   = "Contingency" // Contingency is a String
	EstimationSession_ReqDays_scrn   = "ReqDays" // ReqDays is a String
	EstimationSession_RegCost_scrn   = "RegCost" // RegCost is a String
	EstimationSession_ImpDays_scrn   = "ImpDays" // ImpDays is a String
	EstimationSession_ImpCost_scrn   = "ImpCost" // ImpCost is a String
	EstimationSession_UatDays_scrn   = "UatDays" // UatDays is a String
	EstimationSession_UatCost_scrn   = "UatCost" // UatCost is a String
	EstimationSession_MgtDays_scrn   = "MgtDays" // MgtDays is a String
	EstimationSession_MgtCost_scrn   = "MgtCost" // MgtCost is a String
	EstimationSession_RelDays_scrn   = "RelDays" // RelDays is a String
	EstimationSession_RelCost_scrn   = "RelCost" // RelCost is a String
	EstimationSession_SupportUplift_scrn   = "SupportUplift" // SupportUplift is a String
	EstimationSession_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	EstimationSession_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	EstimationSession_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	EstimationSession_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	EstimationSession_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
	EstimationSession_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
	EstimationSession_SYSDeleted_scrn   = "SYSDeleted" // SYSDeleted is a String
	EstimationSession_SYSDeletedBy_scrn   = "SYSDeletedBy" // SYSDeletedBy is a String
	EstimationSession_SYSDeletedHost_scrn   = "SYSDeletedHost" // SYSDeletedHost is a String
	EstimationSession_Name_scrn   = "Name" // Name is a String

	/// Definitions End
	///
)

//estimationsession_PageList provides the information for the template for a list of EstimationSessions
type EstimationSession_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []EstimationSession
	Context	 appContext
}

//estimationsession_Page provides the information for the template for an individual EstimationSession
type EstimationSession_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	SYSId_props     FieldProperties
	EstimationSessionID         string
	EstimationSessionID_props     FieldProperties
	ProjectID         string
	ProjectID_lookup    []Lookup_Item
	ProjectID_props     FieldProperties
	EstimationStateID         string
	EstimationStateID_lookup    []Lookup_Item
	EstimationStateID_props     FieldProperties
	Notes         string
	Notes_props     FieldProperties
	Releases         string
	Releases_props     FieldProperties
	Total         string
	Total_props     FieldProperties
	Contingency         string
	Contingency_props     FieldProperties
	ReqDays         string
	ReqDays_props     FieldProperties
	RegCost         string
	RegCost_props     FieldProperties
	ImpDays         string
	ImpDays_props     FieldProperties
	ImpCost         string
	ImpCost_props     FieldProperties
	UatDays         string
	UatDays_props     FieldProperties
	UatCost         string
	UatCost_props     FieldProperties
	MgtDays         string
	MgtDays_props     FieldProperties
	MgtCost         string
	MgtCost_props     FieldProperties
	RelDays         string
	RelDays_props     FieldProperties
	RelCost         string
	RelCost_props     FieldProperties
	SupportUplift         string
	SupportUplift_props     FieldProperties
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
	Name         string
	Name_props     FieldProperties
	// 
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}