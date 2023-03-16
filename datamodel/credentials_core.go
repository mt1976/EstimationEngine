package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/credentials.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Credentials (credentials)
// Endpoint 	        : Credentials (Id)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/03/2023 at 19:24:47
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Credentials defines the datamodel for the Credentials object
type Credentials struct {
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	SYSId       string
	Id       string
	Username       string
	Password       string
	Firstname       string
	Lastname       string
	Knownas       string
	Email       string
	Issued       string
	Expiry       string
	RoleType       string
	Brand       string
	SYSCreated       string
	SYSUpdated       string
	SYSCreatedBy       string
	SYSCreatedHost       string
	SYSUpdatedBy       string
	SYSUpdatedHost       string
	State       string
	Notes       string
	SYSDeleted       string
	SYSDeletedBy       string
	SYSDeletedHost       string
	SYSActivity       string
	SYSDbVersion       string
	EmailNotifications       string
	PasswordExpiry       string
	//
	// Field Properties
	//
	SYSId_props FieldProperties
	Id_props FieldProperties
	Username_props FieldProperties
	Password_props FieldProperties
	Firstname_props FieldProperties
	Lastname_props FieldProperties
	Knownas_props FieldProperties
	Email_props FieldProperties
	Issued_props FieldProperties
	Expiry_props FieldProperties
	RoleType_props FieldProperties
	Brand_props FieldProperties
	SYSCreated_props FieldProperties
	SYSUpdated_props FieldProperties
	SYSCreatedBy_props FieldProperties
	SYSCreatedHost_props FieldProperties
	SYSUpdatedBy_props FieldProperties
	SYSUpdatedHost_props FieldProperties
	State_props FieldProperties
	Notes_props FieldProperties
	SYSDeleted_props FieldProperties
	SYSDeletedBy_props FieldProperties
	SYSDeletedHost_props FieldProperties
	SYSActivity_props FieldProperties
	SYSDbVersion_props FieldProperties
	EmailNotifications_props FieldProperties
	PasswordExpiry_props FieldProperties
	//
 	// Any lookups will be added below
	//
	RoleType_lookup []Lookup_Item
	State_lookup []Lookup_Item
	EmailNotifications_lookup []Lookup_Item
	}

const (
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// Page Infrastructure
	Credentials_Name      = "Credentials"
	Credentials_Title       = "Credentials"
	Credentials_SQLTable    = "credentialsStore"
	Credentials_SQLSearchID = "id"
	Credentials_QueryString = "Id"
	///
	/// Template Path Defintions
	///
	Credentials_Template     = "Credentials"
	Credentials_TemplateList = "/Credentials/CredentialsList"
	Credentials_TemplateView = "/Credentials/CredentialsView"
	Credentials_TemplateEdit = "/Credentials/CredentialsEdit"
	Credentials_TemplateNew  = "/Credentials/CredentialsNew"
	///
	/// URI Handler Paths
	///
	Credentials_Path       = "/API/Credentials/"
	Credentials_PathList   = "/CredentialsList/"
	Credentials_PathView   = "/CredentialsView/"
	Credentials_PathEdit   = "/CredentialsEdit/"
	Credentials_PathNew    = "/CredentialsNew/"
	Credentials_PathSave   = "/CredentialsSave/"
	Credentials_PathDelete = "/CredentialsDelete/"
	// Redirects - On Server Side Error
	Credentials_PathEditException   = "/CredentialsEditException/"
	Credentials_PathNewException    = "/CredentialsNewException/"
	//
	//Credentials_Redirect provides a page to return to aftern an action
	Credentials_Redirect = Credentials_PathList
	
	//
	//
	// SQL Field Definitions
	//
	Credentials_SYSId_sql   = "_id" // SYSId is a Int
	Credentials_Id_sql   = "Id" // Id is a String
	Credentials_Username_sql   = "Username" // Username is a String
	Credentials_Password_sql   = "Password" // Password is a String
	Credentials_Firstname_sql   = "Firstname" // Firstname is a String
	Credentials_Lastname_sql   = "Lastname" // Lastname is a String
	Credentials_Knownas_sql   = "Knownas" // Knownas is a String
	Credentials_Email_sql   = "Email" // Email is a String
	Credentials_Issued_sql   = "Issued" // Issued is a String
	Credentials_Expiry_sql   = "Expiry" // Expiry is a String
	Credentials_RoleType_sql   = "RoleType" // RoleType is a String
	Credentials_Brand_sql   = "Brand" // Brand is a String
	Credentials_SYSCreated_sql   = "_created" // SYSCreated is a String
	Credentials_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	Credentials_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	Credentials_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	Credentials_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
	Credentials_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
	Credentials_State_sql   = "State" // State is a String
	Credentials_Notes_sql   = "Notes" // Notes is a String
	Credentials_SYSDeleted_sql   = "_deleted" // SYSDeleted is a String
	Credentials_SYSDeletedBy_sql   = "_deletedBy" // SYSDeletedBy is a String
	Credentials_SYSDeletedHost_sql   = "_deletedHost" // SYSDeletedHost is a String
	Credentials_SYSActivity_sql   = "_activity" // SYSActivity is a String
	Credentials_SYSDbVersion_sql   = "_dbVersion" // SYSDbVersion is a String
	Credentials_EmailNotifications_sql   = "emailNotifications" // EmailNotifications is a String
	Credentials_PasswordExpiry_sql   = "passwordExpiry" // PasswordExpiry is a String
///
	/// Application Field Definitions
	///
	Credentials_SYSId_scrn   = "SYSId" // SYSId is a Int
	Credentials_Id_scrn   = "Id" // Id is a String
	Credentials_Username_scrn   = "Username" // Username is a String
	Credentials_Password_scrn   = "Password" // Password is a String
	Credentials_Firstname_scrn   = "Firstname" // Firstname is a String
	Credentials_Lastname_scrn   = "Lastname" // Lastname is a String
	Credentials_Knownas_scrn   = "Knownas" // Knownas is a String
	Credentials_Email_scrn   = "Email" // Email is a String
	Credentials_Issued_scrn   = "Issued" // Issued is a String
	Credentials_Expiry_scrn   = "Expiry" // Expiry is a String
	Credentials_RoleType_scrn   = "RoleType" // RoleType is a String
	Credentials_Brand_scrn   = "Brand" // Brand is a String
	Credentials_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	Credentials_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	Credentials_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	Credentials_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	Credentials_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
	Credentials_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
	Credentials_State_scrn   = "State" // State is a String
	Credentials_Notes_scrn   = "Notes" // Notes is a String
	Credentials_SYSDeleted_scrn   = "SYSDeleted" // SYSDeleted is a String
	Credentials_SYSDeletedBy_scrn   = "SYSDeletedBy" // SYSDeletedBy is a String
	Credentials_SYSDeletedHost_scrn   = "SYSDeletedHost" // SYSDeletedHost is a String
	Credentials_SYSActivity_scrn   = "SYSActivity" // SYSActivity is a String
	Credentials_SYSDbVersion_scrn   = "SYSDbVersion" // SYSDbVersion is a String
	Credentials_EmailNotifications_scrn   = "EmailNotifications" // EmailNotifications is a String
	Credentials_PasswordExpiry_scrn   = "PasswordExpiry" // PasswordExpiry is a String
///
)

//Credentials_PageList provides the information for the template for a list of Credentialss
type Credentials_PageList struct {
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	ItemList  		 []Credentials
}

//Credentials_Page provides the information for the template for an individual Credentials
type Credentials_Page struct {
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	Username         string
	Password         string
	Firstname         string
	Lastname         string
	Knownas         string
	Email         string
	Issued         string
	Expiry         string
	RoleType         string
	Brand         string
	SYSCreated         string
	SYSUpdated         string
	SYSCreatedBy         string
	SYSCreatedHost         string
	SYSUpdatedBy         string
	SYSUpdatedHost         string
	State         string
	Notes         string
	SYSDeleted         string
	SYSDeletedBy         string
	SYSDeletedHost         string
	SYSActivity         string
	SYSDbVersion         string
	EmailNotifications         string
	PasswordExpiry         string
	/// Field Properties
	SYSId_props     FieldProperties
	Id_props     FieldProperties
	Username_props     FieldProperties
	Password_props     FieldProperties
	Firstname_props     FieldProperties
	Lastname_props     FieldProperties
	Knownas_props     FieldProperties
	Email_props     FieldProperties
	Issued_props     FieldProperties
	Expiry_props     FieldProperties
	RoleType_props     FieldProperties
	Brand_props     FieldProperties
	SYSCreated_props     FieldProperties
	SYSUpdated_props     FieldProperties
	SYSCreatedBy_props     FieldProperties
	SYSCreatedHost_props     FieldProperties
	SYSUpdatedBy_props     FieldProperties
	SYSUpdatedHost_props     FieldProperties
	State_props     FieldProperties
	Notes_props     FieldProperties
	SYSDeleted_props     FieldProperties
	SYSDeletedBy_props     FieldProperties
	SYSDeletedHost_props     FieldProperties
	SYSActivity_props     FieldProperties
	SYSDbVersion_props     FieldProperties
	EmailNotifications_props     FieldProperties
	PasswordExpiry_props     FieldProperties
	/// Lookups
	RoleType_lookup    []Lookup_Item
	State_lookup    []Lookup_Item
	EmailNotifications_lookup    []Lookup_Item
	
	/// END OF DEFINITIONS
	///
}