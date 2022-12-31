package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/userrole.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : UserRole (userrole)
// Endpoint 	        : UserRole (Id)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 30/12/2022 at 12:18:23
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//UserRole defines the datamolde for the UserRole object
type UserRole struct {


SYSId       string
SYSId_props FieldProperties
Id       string
Id_props FieldProperties
Name       string
Name_props FieldProperties
SYSCreatedBy       string
SYSCreatedBy_props FieldProperties
SYSCreatedHost       string
SYSCreatedHost_props FieldProperties
SYSUpdatedBy       string
SYSUpdatedBy_props FieldProperties
SYSUpdatedHost       string
SYSUpdatedHost_props FieldProperties
SYSUpdated       string
SYSUpdated_props FieldProperties
SYSCreated       string
SYSCreated_props FieldProperties
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
	UserRole_Title       = "User Roles"
	UserRole_SQLTable    = "roleStore"
	UserRole_SQLSearchID = "Id"
	UserRole_QueryString = "Id"
	///
	/// Handler Defintions
	///
	UserRole_Template     = "UserRole"
	UserRole_TemplateList = "/UserRole/UserRole_List"
	UserRole_TemplateView = "/UserRole/UserRole_View"
	UserRole_TemplateEdit = "/UserRole/UserRole_Edit"
	UserRole_TemplateNew  = "/UserRole/UserRole_New"
	///
	/// Handler Monitor Paths
	///
	UserRole_Path       = "/API/UserRole/"
	UserRole_PathList   = "/UserRoleList/"
	UserRole_PathView   = "/UserRoleView/"
	UserRole_PathEdit   = "/UserRoleEdit/"
	UserRole_PathNew    = "/UserRoleNew/"
	UserRole_PathSave   = "/UserRoleSave/"
	UserRole_PathDelete = "/UserRoleDelete/"
	///
	//UserRole_Redirect provides a page to return to aftern an action
	UserRole_Redirect = UserRole_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	UserRole_SYSId_sql   = "_id" // SYSId is a Int
	UserRole_Id_sql   = "Id" // Id is a String
	UserRole_Name_sql   = "Name" // Name is a String
	UserRole_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	UserRole_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	UserRole_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
	UserRole_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
	UserRole_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	UserRole_SYSCreated_sql   = "_created" // SYSCreated is a String
	UserRole_SYSDeleted_sql   = "_deleted" // SYSDeleted is a String
	UserRole_SYSDeletedBy_sql   = "_deletedBy" // SYSDeletedBy is a String
	UserRole_SYSDeletedHost_sql   = "_deletedHost" // SYSDeletedHost is a String
	UserRole_SYSDbVersion_sql   = "_dbVersion" // SYSDbVersion is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	UserRole_SYSId_scrn   = "SYSId" // SYSId is a Int
	UserRole_Id_scrn   = "Id" // Id is a String
	UserRole_Name_scrn   = "Name" // Name is a String
	UserRole_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	UserRole_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	UserRole_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
	UserRole_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
	UserRole_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	UserRole_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	UserRole_SYSDeleted_scrn   = "SYSDeleted" // SYSDeleted is a String
	UserRole_SYSDeletedBy_scrn   = "SYSDeletedBy" // SYSDeletedBy is a String
	UserRole_SYSDeletedHost_scrn   = "SYSDeletedHost" // SYSDeletedHost is a String
	UserRole_SYSDbVersion_scrn   = "SYSDbVersion" // SYSDbVersion is a String

	/// Definitions End
	///
)

//userrole_PageList provides the information for the template for a list of UserRoles
type UserRole_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []UserRole
	Context	 appContext
}

//userrole_Page provides the information for the template for an individual UserRole
type UserRole_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	SYSId_props     FieldProperties
	Id         string
	Id_props     FieldProperties
	Name         string
	Name_props     FieldProperties
	SYSCreatedBy         string
	SYSCreatedBy_props     FieldProperties
	SYSCreatedHost         string
	SYSCreatedHost_props     FieldProperties
	SYSUpdatedBy         string
	SYSUpdatedBy_props     FieldProperties
	SYSUpdatedHost         string
	SYSUpdatedHost_props     FieldProperties
	SYSUpdated         string
	SYSUpdated_props     FieldProperties
	SYSCreated         string
	SYSCreated_props     FieldProperties
	SYSDeleted         string
	SYSDeleted_props     FieldProperties
	SYSDeletedBy         string
	SYSDeletedBy_props     FieldProperties
	SYSDeletedHost         string
	SYSDeletedHost_props     FieldProperties
	SYSDbVersion         string
	SYSDbVersion_props     FieldProperties
	// 
	// Dynamically generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}