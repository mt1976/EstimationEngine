package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/project.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Project (project)
// Endpoint 	        : Project (ProjectID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 24/01/2023 at 13:18:10
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Project defines the datamolde for the Project object
type Project struct {


SYSId       string
SYSId_props FieldProperties
ProjectID       string
ProjectID_props FieldProperties
OriginID       string
OriginID_props FieldProperties
ProjectStateID       string
ProjectStateID_props FieldProperties
ProfileID       string
ProfileID_props FieldProperties
Name       string
Name_props FieldProperties
Description       string
Description_props FieldProperties
StartDate       string
StartDate_props FieldProperties
EndDate       string
EndDate_props FieldProperties
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
SYSActivity       string
SYSActivity_props FieldProperties
SYSDbVersion       string
SYSDbVersion_props FieldProperties
Comments       string
Comments_props FieldProperties
ProjectRate       string
ProjectRate_props FieldProperties
DefaultRate       string
DefaultRate_props FieldProperties
ProjectAnalyst       string
ProjectAnalyst_props FieldProperties
ProjectEngineer       string
ProjectEngineer_props FieldProperties
ProjectManager       string
ProjectManager_props FieldProperties
Releases       string
Releases_props FieldProperties
Notes       string
Notes_props FieldProperties
NoEstimationSessions       string
NoEstimationSessions_props FieldProperties
 // Any lookups will be added below

OriginID_lookup []Lookup_Item
ProjectStateID_lookup []Lookup_Item
ProfileID_lookup []Lookup_Item


















ProjectAnalyst_lookup []Lookup_Item
ProjectEngineer_lookup []Lookup_Item
ProjectManager_lookup []Lookup_Item




}

const (
	Project_Title       = "Project"
	Project_SQLTable    = "projectStore"
	Project_SQLSearchID = "projectID"
	Project_QueryString = "ProjectID"
	///
	/// Handler Defintions
	///
	Project_Template     = "Project"
	Project_TemplateList = "/Project/Project_List"
	Project_TemplateView = "/Project/Project_View"
	Project_TemplateEdit = "/Project/Project_Edit"
	Project_TemplateNew  = "/Project/Project_New"
	///
	/// Handler Monitor Paths
	///
	Project_Path       = "/API/Project/"
	Project_PathList   = "/ProjectList/"
	Project_PathView   = "/ProjectView/"
	Project_PathEdit   = "/ProjectEdit/"
	Project_PathNew    = "/ProjectNew/"
	Project_PathSave   = "/ProjectSave/"
	Project_PathDelete = "/ProjectDelete/"
	///
	//Project_Redirect provides a page to return to aftern an action
	Project_Redirect = Project_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	Project_SYSId_sql   = "_id" // SYSId is a Int
	Project_ProjectID_sql   = "projectID" // ProjectID is a String
	Project_OriginID_sql   = "originID" // OriginID is a String
	Project_ProjectStateID_sql   = "projectStateID" // ProjectStateID is a String
	Project_ProfileID_sql   = "profileID" // ProfileID is a String
	Project_Name_sql   = "name" // Name is a String
	Project_Description_sql   = "description" // Description is a String
	Project_StartDate_sql   = "startDate" // StartDate is a String
	Project_EndDate_sql   = "endDate" // EndDate is a String
	Project_SYSCreated_sql   = "_created" // SYSCreated is a String
	Project_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	Project_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	Project_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	Project_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
	Project_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
	Project_SYSDeleted_sql   = "_deleted" // SYSDeleted is a String
	Project_SYSDeletedBy_sql   = "_deletedBy" // SYSDeletedBy is a String
	Project_SYSDeletedHost_sql   = "_deletedHost" // SYSDeletedHost is a String
	Project_SYSActivity_sql   = "_activity" // SYSActivity is a String
	Project_SYSDbVersion_sql   = "_dbVersion" // SYSDbVersion is a String
	Project_Comments_sql   = "comments" // Comments is a String
	Project_ProjectRate_sql   = "projectRate" // ProjectRate is a String
	Project_DefaultRate_sql   = "defaultRate" // DefaultRate is a String
	Project_ProjectAnalyst_sql   = "projectAnalyst" // ProjectAnalyst is a String
	Project_ProjectEngineer_sql   = "projectEngineer" // ProjectEngineer is a String
	Project_ProjectManager_sql   = "projectManager" // ProjectManager is a String
	Project_Releases_sql   = "releases" // Releases is a String
	Project_Notes_sql   = "notes" // Notes is a String
	Project_NoEstimationSessions_sql   = "NoEstimationSessions" // NoEstimationSessions is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	Project_SYSId_scrn   = "SYSId" // SYSId is a Int
	Project_ProjectID_scrn   = "ProjectID" // ProjectID is a String
	Project_OriginID_scrn   = "OriginID" // OriginID is a String
	Project_ProjectStateID_scrn   = "ProjectStateID" // ProjectStateID is a String
	Project_ProfileID_scrn   = "ProfileID" // ProfileID is a String
	Project_Name_scrn   = "Name" // Name is a String
	Project_Description_scrn   = "Description" // Description is a String
	Project_StartDate_scrn   = "StartDate" // StartDate is a String
	Project_EndDate_scrn   = "EndDate" // EndDate is a String
	Project_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	Project_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	Project_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	Project_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	Project_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
	Project_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
	Project_SYSDeleted_scrn   = "SYSDeleted" // SYSDeleted is a String
	Project_SYSDeletedBy_scrn   = "SYSDeletedBy" // SYSDeletedBy is a String
	Project_SYSDeletedHost_scrn   = "SYSDeletedHost" // SYSDeletedHost is a String
	Project_SYSActivity_scrn   = "SYSActivity" // SYSActivity is a String
	Project_SYSDbVersion_scrn   = "SYSDbVersion" // SYSDbVersion is a String
	Project_Comments_scrn   = "Comments" // Comments is a String
	Project_ProjectRate_scrn   = "ProjectRate" // ProjectRate is a String
	Project_DefaultRate_scrn   = "DefaultRate" // DefaultRate is a String
	Project_ProjectAnalyst_scrn   = "ProjectAnalyst" // ProjectAnalyst is a String
	Project_ProjectEngineer_scrn   = "ProjectEngineer" // ProjectEngineer is a String
	Project_ProjectManager_scrn   = "ProjectManager" // ProjectManager is a String
	Project_Releases_scrn   = "Releases" // Releases is a String
	Project_Notes_scrn   = "Notes" // Notes is a String
	Project_NoEstimationSessions_scrn   = "NoEstimationSessions" // NoEstimationSessions is a String

	/// Definitions End
	///
)

//project_PageList provides the information for the template for a list of Projects
type Project_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []Project
	Context	 appContext
	BlockEdit		 bool
	BlockSave 		 bool
	BlockCreate		 bool
	BlockDelete		 bool
	BlockValidate	 bool
	BlockView		 bool
}

//project_Page provides the information for the template for an individual Project
type Project_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	SYSId_props     FieldProperties
	ProjectID         string
	ProjectID_props     FieldProperties
	OriginID         string
	OriginID_lookup    []Lookup_Item
	OriginID_props     FieldProperties
	ProjectStateID         string
	ProjectStateID_lookup    []Lookup_Item
	ProjectStateID_props     FieldProperties
	ProfileID         string
	ProfileID_lookup    []Lookup_Item
	ProfileID_props     FieldProperties
	Name         string
	Name_props     FieldProperties
	Description         string
	Description_props     FieldProperties
	StartDate         string
	StartDate_props     FieldProperties
	EndDate         string
	EndDate_props     FieldProperties
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
	SYSActivity         string
	SYSActivity_props     FieldProperties
	SYSDbVersion         string
	SYSDbVersion_props     FieldProperties
	Comments         string
	Comments_props     FieldProperties
	ProjectRate         string
	ProjectRate_props     FieldProperties
	DefaultRate         string
	DefaultRate_props     FieldProperties
	ProjectAnalyst         string
	ProjectAnalyst_lookup    []Lookup_Item
	ProjectAnalyst_props     FieldProperties
	ProjectEngineer         string
	ProjectEngineer_lookup    []Lookup_Item
	ProjectEngineer_props     FieldProperties
	ProjectManager         string
	ProjectManager_lookup    []Lookup_Item
	ProjectManager_props     FieldProperties
	Releases         string
	Releases_props     FieldProperties
	Notes         string
	Notes_props     FieldProperties
	NoEstimationSessions         string
	NoEstimationSessions_props     FieldProperties
	// 
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 		 appContext
	BlockEdit		 bool
	BlockSave 		 bool
	BlockCreate		 bool
	BlockDelete		 bool
	BlockValidate	 bool
	BlockView		 bool
}