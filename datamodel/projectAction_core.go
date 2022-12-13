package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/projectaction.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : ProjectAction (projectaction)
// Endpoint 	        : ProjectAction (ProjectID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 11/12/2022 at 19:24:02
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//ProjectAction defines the datamolde for the ProjectAction object
type ProjectAction struct {


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
 // Any lookups will be added below

OriginID_lookup []Lookup_Item
ProjectStateID_lookup []Lookup_Item
ProfileID_lookup []Lookup_Item

















}

const (
	ProjectAction_Title       = "ProjectAction"
	ProjectAction_SQLTable    = "projectStore"
	ProjectAction_SQLSearchID = "projectID"
	ProjectAction_QueryString = "ProjectID"
	///
	/// Handler Defintions
	///
	ProjectAction_Template     = "ProjectAction"
	ProjectAction_TemplateList = "/ProjectAction/ProjectAction_List"
	ProjectAction_TemplateView = "/ProjectAction/ProjectAction_View"
	ProjectAction_TemplateEdit = "/ProjectAction/ProjectAction_Edit"
	ProjectAction_TemplateNew  = "/ProjectAction/ProjectAction_New"
	///
	/// Handler Monitor Paths
	///
	ProjectAction_Path       = "/API/ProjectAction/"
	ProjectAction_PathList   = "/ProjectActionList/"
	ProjectAction_PathView   = "/ProjectActionView/"
	ProjectAction_PathEdit   = "/ProjectActionEdit/"
	ProjectAction_PathNew    = "/ProjectActionNew/"
	ProjectAction_PathSave   = "/ProjectActionSave/"
	ProjectAction_PathDelete = "/ProjectActionDelete/"
	///
	//ProjectAction_Redirect provides a page to return to aftern an action
	ProjectAction_Redirect = ProjectAction_PathView
	///
	///
	/// SQL Field Definitions
	///
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

	/// Definitions End
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

	/// Definitions End
	///
)

//projectaction_PageList provides the information for the template for a list of ProjectActions
type ProjectAction_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []ProjectAction
	Context	 appContext
}

//projectaction_Page provides the information for the template for an individual ProjectAction
type ProjectAction_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
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
	// 
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}