package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/credentials.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Credentials (credentials)
// Endpoint 	        : Credentials (Id)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 03/01/2023 at 19:18:08
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Credentials defines the datamolde for the Credentials object
type Credentials struct {


SYSId       string
SYSId_props FieldProperties
Id       string
Id_props FieldProperties
Username       string
Username_props FieldProperties
Password       string
Password_props FieldProperties
Firstname       string
Firstname_props FieldProperties
Lastname       string
Lastname_props FieldProperties
Knownas       string
Knownas_props FieldProperties
Email       string
Email_props FieldProperties
Issued       string
Issued_props FieldProperties
Expiry       string
Expiry_props FieldProperties
RoleType       string
RoleType_props FieldProperties
Brand       string
Brand_props FieldProperties
SYSCreated       string
SYSCreated_props FieldProperties
SYSUpdated       string
SYSUpdated_props FieldProperties
SYSCreatedBy       string
SYSCreatedBy_props FieldProperties
SYSCreatedHost       string
SYSCreatedHost_props FieldProperties
SYSUpdatedBy       string
SYSUpdatedBy_props FieldProperties
SYSUpdatedHost       string
SYSUpdatedHost_props FieldProperties
State       string
State_props FieldProperties
Notes       string
Notes_props FieldProperties
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
EmailNotifications       string
EmailNotifications_props FieldProperties
 // Any lookups will be added below









RoleType_lookup []Lookup_Item







State_lookup []Lookup_Item








}

const (
	Credentials_Title       = "Credentials"
	Credentials_SQLTable    = "credentialsStore"
	Credentials_SQLSearchID = "id"
	Credentials_QueryString = "Id"
	///
	/// Handler Defintions
	///
	Credentials_Template     = "Credentials"
	Credentials_TemplateList = "/Credentials/Credentials_List"
	Credentials_TemplateView = "/Credentials/Credentials_View"
	Credentials_TemplateEdit = "/Credentials/Credentials_Edit"
	Credentials_TemplateNew  = "/Credentials/Credentials_New"
	///
	/// Handler Monitor Paths
	///
	Credentials_Path       = "/API/Credentials/"
	Credentials_PathList   = "/CredentialsList/"
	Credentials_PathView   = "/CredentialsView/"
	Credentials_PathEdit   = "/CredentialsEdit/"
	Credentials_PathNew    = "/CredentialsNew/"
	Credentials_PathSave   = "/CredentialsSave/"
	Credentials_PathDelete = "/CredentialsDelete/"
	///
	//Credentials_Redirect provides a page to return to aftern an action
	Credentials_Redirect = Credentials_PathList
	
	///
	///
	/// SQL Field Definitions
	///
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

	/// Definitions End
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

	/// Definitions End
	///
)

//credentials_PageList provides the information for the template for a list of Credentialss
type Credentials_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []Credentials
	Context	 appContext
}

//credentials_Page provides the information for the template for an individual Credentials
type Credentials_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	SYSId_props     FieldProperties
	Id         string
	Id_props     FieldProperties
	Username         string
	Username_props     FieldProperties
	Password         string
	Password_props     FieldProperties
	Firstname         string
	Firstname_props     FieldProperties
	Lastname         string
	Lastname_props     FieldProperties
	Knownas         string
	Knownas_props     FieldProperties
	Email         string
	Email_props     FieldProperties
	Issued         string
	Issued_props     FieldProperties
	Expiry         string
	Expiry_props     FieldProperties
	RoleType         string
	RoleType_lookup    []Lookup_Item
	RoleType_props     FieldProperties
	Brand         string
	Brand_props     FieldProperties
	SYSCreated         string
	SYSCreated_props     FieldProperties
	SYSUpdated         string
	SYSUpdated_props     FieldProperties
	SYSCreatedBy         string
	SYSCreatedBy_props     FieldProperties
	SYSCreatedHost         string
	SYSCreatedHost_props     FieldProperties
	SYSUpdatedBy         string
	SYSUpdatedBy_props     FieldProperties
	SYSUpdatedHost         string
	SYSUpdatedHost_props     FieldProperties
	State         string
	State_lookup    []Lookup_Item
	State_props     FieldProperties
	Notes         string
	Notes_props     FieldProperties
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
	EmailNotifications         string
	EmailNotifications_props     FieldProperties
	// 
	// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}