package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/session.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Session (session)
// Endpoint 	        : Session (SessionID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 10/03/2023 at 23:49:56
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Session defines the datamodel for the Session object
type Session struct {
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	SYSId       string
	Id       string
	Apptoken       string
	Createdate       string
	Createtime       string
	Uniqueid       string
	Sessiontoken       string
	Username       string
	Password       string
	Userip       string
	Userhost       string
	Appip       string
	Apphost       string
	Issued       string
	Expiry       string
	Expiryraw       string
	Expires       string
	Brand       string
	SessionRole       string
	SYSCreated       string
	SYSCreatedBy       string
	SYSCreatedHost       string
	SYSUpdated       string
	SYSUpdatedBy       string
	SYSUpdatedHost       string
	SYSDeleted       string
	SYSDeletedBy       string
	SYSDeletedHost       string
	SYSDbVersion       string
	//
	// Field Properties
	//
	SYSId_props FieldProperties
	Id_props FieldProperties
	Apptoken_props FieldProperties
	Createdate_props FieldProperties
	Createtime_props FieldProperties
	Uniqueid_props FieldProperties
	Sessiontoken_props FieldProperties
	Username_props FieldProperties
	Password_props FieldProperties
	Userip_props FieldProperties
	Userhost_props FieldProperties
	Appip_props FieldProperties
	Apphost_props FieldProperties
	Issued_props FieldProperties
	Expiry_props FieldProperties
	Expiryraw_props FieldProperties
	Expires_props FieldProperties
	Brand_props FieldProperties
	SessionRole_props FieldProperties
	SYSCreated_props FieldProperties
	SYSCreatedBy_props FieldProperties
	SYSCreatedHost_props FieldProperties
	SYSUpdated_props FieldProperties
	SYSUpdatedBy_props FieldProperties
	SYSUpdatedHost_props FieldProperties
	SYSDeleted_props FieldProperties
	SYSDeletedBy_props FieldProperties
	SYSDeletedHost_props FieldProperties
	SYSDbVersion_props FieldProperties
	//
 	// Any lookups will be added below
	//
	}

const (
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// Page Infrastructure
	Session_Name      = "Session"
	Session_Title       = "Session"
	Session_SQLTable    = "sessionStore"
	Session_SQLSearchID = "Id"
	Session_QueryString = "SessionID"
	///
	/// Template Path Defintions
	///
	Session_Template     = "Session"
	Session_TemplateList = "/Session/SessionList"
	Session_TemplateView = "/Session/SessionView"
	Session_TemplateEdit = "/Session/SessionEdit"
	Session_TemplateNew  = "/Session/SessionNew"
	///
	/// URI Handler Paths
	///
	Session_Path       = "/API/Session/"
	Session_PathList   = "/SessionList/"
	Session_PathView   = "/SessionView/"
	Session_PathEdit   = "/SessionEdit/"
	Session_PathNew    = "/SessionNew/"
	Session_PathSave   = "/SessionSave/"
	Session_PathDelete = "/SessionDelete/"
	// Redirects - On Server Side Error
	Session_PathEditException   = "/SessionEditException/"
	Session_PathNewException    = "/SessionNewException/"
	//
	//Session_Redirect provides a page to return to aftern an action
	Session_Redirect = Session_PathList
	
	//
	//
	// SQL Field Definitions
	//
	Session_SYSId_sql   = "_id" // SYSId is a Int
	Session_Id_sql   = "Id" // Id is a String
	Session_Apptoken_sql   = "Apptoken" // Apptoken is a String
	Session_Createdate_sql   = "Createdate" // Createdate is a String
	Session_Createtime_sql   = "Createtime" // Createtime is a String
	Session_Uniqueid_sql   = "Uniqueid" // Uniqueid is a String
	Session_Sessiontoken_sql   = "Sessiontoken" // Sessiontoken is a String
	Session_Username_sql   = "Username" // Username is a String
	Session_Password_sql   = "Password" // Password is a String
	Session_Userip_sql   = "Userip" // Userip is a String
	Session_Userhost_sql   = "Userhost" // Userhost is a String
	Session_Appip_sql   = "Appip" // Appip is a String
	Session_Apphost_sql   = "Apphost" // Apphost is a String
	Session_Issued_sql   = "Issued" // Issued is a String
	Session_Expiry_sql   = "Expiry" // Expiry is a String
	Session_Expiryraw_sql   = "Expiryraw" // Expiryraw is a String
	Session_Expires_sql   = "Expires" // Expires is a Time
	Session_Brand_sql   = "Brand" // Brand is a String
	Session_SessionRole_sql   = "SessionRole" // SessionRole is a String
	Session_SYSCreated_sql   = "_created" // SYSCreated is a String
	Session_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	Session_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	Session_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	Session_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
	Session_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
	Session_SYSDeleted_sql   = "_deleted" // SYSDeleted is a String
	Session_SYSDeletedBy_sql   = "_deletedBy" // SYSDeletedBy is a String
	Session_SYSDeletedHost_sql   = "_deletedHost" // SYSDeletedHost is a String
	Session_SYSDbVersion_sql   = "_dbVersion" // SYSDbVersion is a String
///
	/// Application Field Definitions
	///
	Session_SYSId_scrn   = "SYSId" // SYSId is a Int
	Session_Id_scrn   = "Id" // Id is a String
	Session_Apptoken_scrn   = "Apptoken" // Apptoken is a String
	Session_Createdate_scrn   = "Createdate" // Createdate is a String
	Session_Createtime_scrn   = "Createtime" // Createtime is a String
	Session_Uniqueid_scrn   = "Uniqueid" // Uniqueid is a String
	Session_Sessiontoken_scrn   = "Sessiontoken" // Sessiontoken is a String
	Session_Username_scrn   = "Username" // Username is a String
	Session_Password_scrn   = "Password" // Password is a String
	Session_Userip_scrn   = "Userip" // Userip is a String
	Session_Userhost_scrn   = "Userhost" // Userhost is a String
	Session_Appip_scrn   = "Appip" // Appip is a String
	Session_Apphost_scrn   = "Apphost" // Apphost is a String
	Session_Issued_scrn   = "Issued" // Issued is a String
	Session_Expiry_scrn   = "Expiry" // Expiry is a String
	Session_Expiryraw_scrn   = "Expiryraw" // Expiryraw is a String
	Session_Expires_scrn   = "Expires" // Expires is a Time
	Session_Brand_scrn   = "Brand" // Brand is a String
	Session_SessionRole_scrn   = "SessionRole" // SessionRole is a String
	Session_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	Session_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	Session_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	Session_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	Session_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
	Session_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
	Session_SYSDeleted_scrn   = "SYSDeleted" // SYSDeleted is a String
	Session_SYSDeletedBy_scrn   = "SYSDeletedBy" // SYSDeletedBy is a String
	Session_SYSDeletedHost_scrn   = "SYSDeletedHost" // SYSDeletedHost is a String
	Session_SYSDbVersion_scrn   = "SYSDbVersion" // SYSDbVersion is a String
///
)

//Session_PageList provides the information for the template for a list of Sessions
type Session_PageList struct {
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	ItemList  		 []Session
}

//Session_Page provides the information for the template for an individual Session
type Session_Page struct {
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	Apptoken         string
	Createdate         string
	Createtime         string
	Uniqueid         string
	Sessiontoken         string
	Username         string
	Password         string
	Userip         string
	Userhost         string
	Appip         string
	Apphost         string
	Issued         string
	Expiry         string
	Expiryraw         string
	Expires         string
	Brand         string
	SessionRole         string
	SYSCreated         string
	SYSCreatedBy         string
	SYSCreatedHost         string
	SYSUpdated         string
	SYSUpdatedBy         string
	SYSUpdatedHost         string
	SYSDeleted         string
	SYSDeletedBy         string
	SYSDeletedHost         string
	SYSDbVersion         string
	/// Field Properties
	SYSId_props     FieldProperties
	Id_props     FieldProperties
	Apptoken_props     FieldProperties
	Createdate_props     FieldProperties
	Createtime_props     FieldProperties
	Uniqueid_props     FieldProperties
	Sessiontoken_props     FieldProperties
	Username_props     FieldProperties
	Password_props     FieldProperties
	Userip_props     FieldProperties
	Userhost_props     FieldProperties
	Appip_props     FieldProperties
	Apphost_props     FieldProperties
	Issued_props     FieldProperties
	Expiry_props     FieldProperties
	Expiryraw_props     FieldProperties
	Expires_props     FieldProperties
	Brand_props     FieldProperties
	SessionRole_props     FieldProperties
	SYSCreated_props     FieldProperties
	SYSCreatedBy_props     FieldProperties
	SYSCreatedHost_props     FieldProperties
	SYSUpdated_props     FieldProperties
	SYSUpdatedBy_props     FieldProperties
	SYSUpdatedHost_props     FieldProperties
	SYSDeleted_props     FieldProperties
	SYSDeletedBy_props     FieldProperties
	SYSDeletedHost_props     FieldProperties
	SYSDbVersion_props     FieldProperties
	/// Lookups
	
	/// END OF DEFINITIONS
	///
}