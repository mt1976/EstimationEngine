package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/projectaction.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : ProjectAction (projectaction)
// Endpoint 	        : ProjectAction (ProjectID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 08/03/2023 at 18:42:25
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//ProjectAction defines the datamodel for the ProjectAction object
type ProjectAction struct {
	// Dynamically generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	SYSId       string
	ProjectID       string
	OriginID       string
	ProjectStateID       string
	ProfileID       string
	Name       string
	Description       string
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
	SYSDbVersion       string
	Comments       string
	ProjectRate       string
	DefaultRate       string
	ProjectAnalyst       string
	ProjectEngineer       string
	ProjectManager       string
	Releases       string
	Notes       string
	NoEstimationSessions       string
	//
	// Field Properties
	//
	SYSId_props FieldProperties
	ProjectID_props FieldProperties
	OriginID_props FieldProperties
	ProjectStateID_props FieldProperties
	ProfileID_props FieldProperties
	Name_props FieldProperties
	Description_props FieldProperties
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
	SYSDbVersion_props FieldProperties
	Comments_props FieldProperties
	ProjectRate_props FieldProperties
	DefaultRate_props FieldProperties
	ProjectAnalyst_props FieldProperties
	ProjectEngineer_props FieldProperties
	ProjectManager_props FieldProperties
	Releases_props FieldProperties
	Notes_props FieldProperties
	NoEstimationSessions_props FieldProperties
	//
 	// Any lookups will be added below
	//
	OriginID_lookup []Lookup_Item
	ProjectStateID_lookup []Lookup_Item
	ProfileID_lookup []Lookup_Item
	ProjectAnalyst_lookup []Lookup_Item
	ProjectEngineer_lookup []Lookup_Item
	ProjectManager_lookup []Lookup_Item
	}

const (
	// Dynamically generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// Page Infrastructure
	ProjectAction_Name      = "ProjectAction"
	ProjectAction_Title       = "ProjectAction"
	ProjectAction_SQLTable    = "projectStore"
	ProjectAction_SQLSearchID = "projectID"
	ProjectAction_QueryString = "ProjectID"
	///
	/// Template Path Defintions
	///
	ProjectAction_Template     = "ProjectAction"
	ProjectAction_TemplateList = "/ProjectAction/ProjectActionList"
	ProjectAction_TemplateView = "/ProjectAction/ProjectActionView"
	ProjectAction_TemplateEdit = "/ProjectAction/ProjectActionEdit"
	ProjectAction_TemplateNew  = "/ProjectAction/ProjectActionNew"
	///
	/// URI Handler Paths
	///
	ProjectAction_Path       = "/API/ProjectAction/"
	ProjectAction_PathList   = "/ProjectActionList/"
	ProjectAction_PathView   = "/ProjectActionView/"
	ProjectAction_PathEdit   = "/ProjectActionEdit/"
	ProjectAction_PathNew    = "/ProjectActionNew/"
	ProjectAction_PathSave   = "/ProjectActionSave/"
	ProjectAction_PathDelete = "/ProjectActionDelete/"
	// Redirects - On Server Side Error
	ProjectAction_PathEditException   = "/ProjectActionEditException/"
	ProjectAction_PathNewException    = "/ProjectActionNewException/"
	//
	//ProjectAction_Redirect provides a page to return to aftern an action
	ProjectAction_Redirect = ProjectAction_PathView
	//
	//
	// SQL Field Definitions
	//
	ProjectAction_SYSId_sql   = "_id" // SYSId is a Int
	ProjectAction_ProjectID_sql   = "projectID" // ProjectID is a String
	ProjectAction_OriginID_sql   = "originID" // OriginID is a String
	ProjectAction_ProjectStateID_sql   = "projectStateID" // ProjectStateID is a String
	ProjectAction_ProfileID_sql   = "profileID" // ProfileID is a String
	ProjectAction_Name_sql   = "name" // Name is a String
	ProjectAction_Description_sql   = "description" // Description is a String
	ProjectAction_StartDate_sql   = "startDate" // StartDate is a String
	ProjectAction_EndDate_sql   = "endDate" // EndDate is a String
	ProjectAction_SYSCreated_sql   = "_created" // SYSCreated is a String
	ProjectAction_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	ProjectAction_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	ProjectAction_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	ProjectAction_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
	ProjectAction_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
	ProjectAction_SYSDeleted_sql   = "_deleted" // SYSDeleted is a String
	ProjectAction_SYSDeletedBy_sql   = "_deletedBy" // SYSDeletedBy is a String
	ProjectAction_SYSDeletedHost_sql   = "_deletedHost" // SYSDeletedHost is a String
	ProjectAction_SYSActivity_sql   = "_activity" // SYSActivity is a String
	ProjectAction_SYSDbVersion_sql   = "_dbVersion" // SYSDbVersion is a String
	ProjectAction_Comments_sql   = "comments" // Comments is a String
	ProjectAction_ProjectRate_sql   = "projectRate" // ProjectRate is a String
	ProjectAction_DefaultRate_sql   = "defaultRate" // DefaultRate is a String
	ProjectAction_ProjectAnalyst_sql   = "projectAnalyst" // ProjectAnalyst is a String
	ProjectAction_ProjectEngineer_sql   = "projectEngineer" // ProjectEngineer is a String
	ProjectAction_ProjectManager_sql   = "projectManager" // ProjectManager is a String
	ProjectAction_Releases_sql   = "releases" // Releases is a String
	ProjectAction_Notes_sql   = "notes" // Notes is a String
	ProjectAction_NoEstimationSessions_sql   = "NoEstimationSessions" // NoEstimationSessions is a String
///
	/// Application Field Definitions
	///
	ProjectAction_SYSId_scrn   = "SYSId" // SYSId is a Int
	ProjectAction_ProjectID_scrn   = "ProjectID" // ProjectID is a String
	ProjectAction_OriginID_scrn   = "OriginID" // OriginID is a String
	ProjectAction_ProjectStateID_scrn   = "ProjectStateID" // ProjectStateID is a String
	ProjectAction_ProfileID_scrn   = "ProfileID" // ProfileID is a String
	ProjectAction_Name_scrn   = "Name" // Name is a String
	ProjectAction_Description_scrn   = "Description" // Description is a String
	ProjectAction_StartDate_scrn   = "StartDate" // StartDate is a String
	ProjectAction_EndDate_scrn   = "EndDate" // EndDate is a String
	ProjectAction_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	ProjectAction_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	ProjectAction_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	ProjectAction_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	ProjectAction_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
	ProjectAction_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
	ProjectAction_SYSDeleted_scrn   = "SYSDeleted" // SYSDeleted is a String
	ProjectAction_SYSDeletedBy_scrn   = "SYSDeletedBy" // SYSDeletedBy is a String
	ProjectAction_SYSDeletedHost_scrn   = "SYSDeletedHost" // SYSDeletedHost is a String
	ProjectAction_SYSActivity_scrn   = "SYSActivity" // SYSActivity is a String
	ProjectAction_SYSDbVersion_scrn   = "SYSDbVersion" // SYSDbVersion is a String
	ProjectAction_Comments_scrn   = "Comments" // Comments is a String
	ProjectAction_ProjectRate_scrn   = "ProjectRate" // ProjectRate is a String
	ProjectAction_DefaultRate_scrn   = "DefaultRate" // DefaultRate is a String
	ProjectAction_ProjectAnalyst_scrn   = "ProjectAnalyst" // ProjectAnalyst is a String
	ProjectAction_ProjectEngineer_scrn   = "ProjectEngineer" // ProjectEngineer is a String
	ProjectAction_ProjectManager_scrn   = "ProjectManager" // ProjectManager is a String
	ProjectAction_Releases_scrn   = "Releases" // Releases is a String
	ProjectAction_Notes_scrn   = "Notes" // Notes is a String
	ProjectAction_NoEstimationSessions_scrn   = "NoEstimationSessions" // NoEstimationSessions is a String
///
)

//ProjectAction_PageList provides the information for the template for a list of ProjectActions
type ProjectAction_PageList struct {
	// Dynamically generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	ItemList  		 []ProjectAction
}

//ProjectAction_Page provides the information for the template for an individual ProjectAction
type ProjectAction_Page struct {
	// Dynamically generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	ProjectID         string
	OriginID         string
	ProjectStateID         string
	ProfileID         string
	Name         string
	Description         string
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
	SYSDbVersion         string
	Comments         string
	ProjectRate         string
	DefaultRate         string
	ProjectAnalyst         string
	ProjectEngineer         string
	ProjectManager         string
	Releases         string
	Notes         string
	NoEstimationSessions         string
	/// Field Properties
	SYSId_props     FieldProperties
	ProjectID_props     FieldProperties
	OriginID_props     FieldProperties
	ProjectStateID_props     FieldProperties
	ProfileID_props     FieldProperties
	Name_props     FieldProperties
	Description_props     FieldProperties
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
	SYSDbVersion_props     FieldProperties
	Comments_props     FieldProperties
	ProjectRate_props     FieldProperties
	DefaultRate_props     FieldProperties
	ProjectAnalyst_props     FieldProperties
	ProjectEngineer_props     FieldProperties
	ProjectManager_props     FieldProperties
	Releases_props     FieldProperties
	Notes_props     FieldProperties
	NoEstimationSessions_props     FieldProperties
	/// Lookups
	OriginID_lookup    []Lookup_Item
	ProjectStateID_lookup    []Lookup_Item
	ProfileID_lookup    []Lookup_Item
	ProjectAnalyst_lookup    []Lookup_Item
	ProjectEngineer_lookup    []Lookup_Item
	ProjectManager_lookup    []Lookup_Item
	
	/// END OF DEFINITIONS
	///
}