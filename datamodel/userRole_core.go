package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/userrole.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : UserRole (userrole)
// Endpoint 	        : UserRole (Id)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 13/03/2023 at 14:22:32
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//UserRole defines the datamodel for the UserRole object
type UserRole struct {
	// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	SYSId       string
	Id       string
	Name       string
	SYSCreatedBy       string
	SYSCreatedHost       string
	SYSUpdatedBy       string
	SYSUpdatedHost       string
	SYSUpdated       string
	SYSCreated       string
	SYSDeleted       string
	SYSDeletedBy       string
	SYSDeletedHost       string
	SYSDbVersion       string
	//
	// Field Properties
	//
	SYSId_props FieldProperties
	Id_props FieldProperties
	Name_props FieldProperties
	SYSCreatedBy_props FieldProperties
	SYSCreatedHost_props FieldProperties
	SYSUpdatedBy_props FieldProperties
	SYSUpdatedHost_props FieldProperties
	SYSUpdated_props FieldProperties
	SYSCreated_props FieldProperties
	SYSDeleted_props FieldProperties
	SYSDeletedBy_props FieldProperties
	SYSDeletedHost_props FieldProperties
	SYSDbVersion_props FieldProperties
	//
 	// Any lookups will be added below
	//
	}

const (
	// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// Page Infrastructure
	UserRole_Name      = "UserRole"
	UserRole_Title       = "User Roles"
	UserRole_SQLTable    = "roleStore"
	UserRole_SQLSearchID = "Id"
	UserRole_QueryString = "Id"
	///
	/// Template Path Defintions
	///
	UserRole_Template     = "UserRole"
	UserRole_TemplateList = "/UserRole/UserRoleList"
	UserRole_TemplateView = "/UserRole/UserRoleView"
	UserRole_TemplateEdit = "/UserRole/UserRoleEdit"
	UserRole_TemplateNew  = "/UserRole/UserRoleNew"
	///
	/// URI Handler Paths
	///
	UserRole_Path       = "/API/UserRole/"
	UserRole_PathList   = "/UserRoleList/"
	UserRole_PathView   = "/UserRoleView/"
	UserRole_PathEdit   = "/UserRoleEdit/"
	UserRole_PathNew    = "/UserRoleNew/"
	UserRole_PathSave   = "/UserRoleSave/"
	UserRole_PathDelete = "/UserRoleDelete/"
	// Redirects - On Server Side Error
	UserRole_PathEditException   = "/UserRoleEditException/"
	UserRole_PathNewException    = "/UserRoleNewException/"
	//
	//UserRole_Redirect provides a page to return to aftern an action
	UserRole_Redirect = UserRole_PathList
	
	//
	//
	// SQL Field Definitions
	//
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
///
)

//UserRole_PageList provides the information for the template for a list of UserRoles
type UserRole_PageList struct {
	// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	ItemList  		 []UserRole
}

//UserRole_Page provides the information for the template for an individual UserRole
type UserRole_Page struct {
	// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	Id         string
	Name         string
	SYSCreatedBy         string
	SYSCreatedHost         string
	SYSUpdatedBy         string
	SYSUpdatedHost         string
	SYSUpdated         string
	SYSCreated         string
	SYSDeleted         string
	SYSDeletedBy         string
	SYSDeletedHost         string
	SYSDbVersion         string
	/// Field Properties
	SYSId_props     FieldProperties
	Id_props     FieldProperties
	Name_props     FieldProperties
	SYSCreatedBy_props     FieldProperties
	SYSCreatedHost_props     FieldProperties
	SYSUpdatedBy_props     FieldProperties
	SYSUpdatedHost_props     FieldProperties
	SYSUpdated_props     FieldProperties
	SYSCreated_props     FieldProperties
	SYSDeleted_props     FieldProperties
	SYSDeletedBy_props     FieldProperties
	SYSDeletedHost_props     FieldProperties
	SYSDbVersion_props     FieldProperties
	/// Lookups
	
	/// END OF DEFINITIONS
	///
}