package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/inbox.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Inbox (inbox)
// Endpoint 	        : Inbox (Message)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 04/03/2023 at 20:14:12
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Inbox defines the datamodel for the Inbox object
type Inbox struct {
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	SYSId       string
	SYSCreated       string
	SYSUpdated       string
	SYSCreatedBy       string
	SYSCreatedHost       string
	SYSUpdatedBy       string
	SYSUpdatedHost       string
	MailId       string
	MailTo       string
	MailFrom       string
	MailSource       string
	MailSent       string
	MailUnread       string
	MailSubject       string
	MailContent       string
	MailAcknowledged       string
	SYSDeleted       string
	SYSDeletedBy       string
	SYSDeletedHost       string
	SYSDbVersion       string
	//
	// Field Properties
	//
	SYSId_props FieldProperties
	SYSCreated_props FieldProperties
	SYSUpdated_props FieldProperties
	SYSCreatedBy_props FieldProperties
	SYSCreatedHost_props FieldProperties
	SYSUpdatedBy_props FieldProperties
	SYSUpdatedHost_props FieldProperties
	MailId_props FieldProperties
	MailTo_props FieldProperties
	MailFrom_props FieldProperties
	MailSource_props FieldProperties
	MailSent_props FieldProperties
	MailUnread_props FieldProperties
	MailSubject_props FieldProperties
	MailContent_props FieldProperties
	MailAcknowledged_props FieldProperties
	SYSDeleted_props FieldProperties
	SYSDeletedBy_props FieldProperties
	SYSDeletedHost_props FieldProperties
	SYSDbVersion_props FieldProperties
	//
 	// Any lookups will be added below
	//
	MailUnread_lookup []Lookup_Item
	MailAcknowledged_lookup []Lookup_Item
	}

const (
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// Page Infrastructure
	Inbox_Name      = "Inbox"
	Inbox_Title       = "Inbox"
	Inbox_SQLTable    = "inboxMessages"
	Inbox_SQLSearchID = "MailId"
	Inbox_QueryString = "Message"
	///
	/// Template Path Defintions
	///
	Inbox_Template     = "Inbox"
	Inbox_TemplateList = "/Inbox/InboxList"
	Inbox_TemplateView = "/Inbox/InboxView"
	Inbox_TemplateEdit = "/Inbox/InboxEdit"
	Inbox_TemplateNew  = "/Inbox/InboxNew"
	///
	/// URI Handler Paths
	///
	Inbox_Path       = "/API/Inbox/"
	Inbox_PathList   = "/InboxList/"
	Inbox_PathView   = "/InboxView/"
	Inbox_PathEdit   = "/InboxEdit/"
	Inbox_PathNew    = "/InboxNew/"
	Inbox_PathSave   = "/InboxSave/"
	Inbox_PathDelete = "/InboxDelete/"
	// Redirects - On Server Side Error
	Inbox_PathEditException   = "/InboxEditException/"
	Inbox_PathNewException    = "/InboxNewException/"
	//
	//Inbox_Redirect provides a page to return to aftern an action
	Inbox_Redirect = Inbox_PathList
	
	//
	//
	// SQL Field Definitions
	//
	Inbox_SYSId_sql   = "_id" // SYSId is a Int
	Inbox_SYSCreated_sql   = "_created" // SYSCreated is a String
	Inbox_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	Inbox_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	Inbox_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	Inbox_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
	Inbox_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
	Inbox_MailId_sql   = "MailId" // MailId is a String
	Inbox_MailTo_sql   = "MailTo" // MailTo is a String
	Inbox_MailFrom_sql   = "MailFrom" // MailFrom is a String
	Inbox_MailSource_sql   = "MailSource" // MailSource is a String
	Inbox_MailSent_sql   = "MailSent" // MailSent is a String
	Inbox_MailUnread_sql   = "MailUnread" // MailUnread is a String
	Inbox_MailSubject_sql   = "MailSubject" // MailSubject is a String
	Inbox_MailContent_sql   = "MailContent" // MailContent is a String
	Inbox_MailAcknowledged_sql   = "MailAcknowledged" // MailAcknowledged is a String
	Inbox_SYSDeleted_sql   = "_deleted" // SYSDeleted is a String
	Inbox_SYSDeletedBy_sql   = "_deletedBy" // SYSDeletedBy is a String
	Inbox_SYSDeletedHost_sql   = "_deletedHost" // SYSDeletedHost is a String
	Inbox_SYSDbVersion_sql   = "_dbVersion" // SYSDbVersion is a String
///
	/// Application Field Definitions
	///
	Inbox_SYSId_scrn   = "SYSId" // SYSId is a Int
	Inbox_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	Inbox_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	Inbox_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	Inbox_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	Inbox_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
	Inbox_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
	Inbox_MailId_scrn   = "MailId" // MailId is a String
	Inbox_MailTo_scrn   = "MailTo" // MailTo is a String
	Inbox_MailFrom_scrn   = "MailFrom" // MailFrom is a String
	Inbox_MailSource_scrn   = "MailSource" // MailSource is a String
	Inbox_MailSent_scrn   = "MailSent" // MailSent is a String
	Inbox_MailUnread_scrn   = "MailUnread" // MailUnread is a String
	Inbox_MailSubject_scrn   = "MailSubject" // MailSubject is a String
	Inbox_MailContent_scrn   = "MailContent" // MailContent is a String
	Inbox_MailAcknowledged_scrn   = "MailAcknowledged" // MailAcknowledged is a String
	Inbox_SYSDeleted_scrn   = "SYSDeleted" // SYSDeleted is a String
	Inbox_SYSDeletedBy_scrn   = "SYSDeletedBy" // SYSDeletedBy is a String
	Inbox_SYSDeletedHost_scrn   = "SYSDeletedHost" // SYSDeletedHost is a String
	Inbox_SYSDbVersion_scrn   = "SYSDbVersion" // SYSDbVersion is a String
///
)

//Inbox_PageList provides the information for the template for a list of Inboxs
type Inbox_PageList struct {
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	ItemList  		 []Inbox
}

//Inbox_Page provides the information for the template for an individual Inbox
type Inbox_Page struct {
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	SYSCreated         string
	SYSUpdated         string
	SYSCreatedBy         string
	SYSCreatedHost         string
	SYSUpdatedBy         string
	SYSUpdatedHost         string
	MailId         string
	MailTo         string
	MailFrom         string
	MailSource         string
	MailSent         string
	MailUnread         string
	MailSubject         string
	MailContent         string
	MailAcknowledged         string
	SYSDeleted         string
	SYSDeletedBy         string
	SYSDeletedHost         string
	SYSDbVersion         string
	/// Field Properties
	SYSId_props     FieldProperties
	SYSCreated_props     FieldProperties
	SYSUpdated_props     FieldProperties
	SYSCreatedBy_props     FieldProperties
	SYSCreatedHost_props     FieldProperties
	SYSUpdatedBy_props     FieldProperties
	SYSUpdatedHost_props     FieldProperties
	MailId_props     FieldProperties
	MailTo_props     FieldProperties
	MailFrom_props     FieldProperties
	MailSource_props     FieldProperties
	MailSent_props     FieldProperties
	MailUnread_props     FieldProperties
	MailSubject_props     FieldProperties
	MailContent_props     FieldProperties
	MailAcknowledged_props     FieldProperties
	SYSDeleted_props     FieldProperties
	SYSDeletedBy_props     FieldProperties
	SYSDeletedHost_props     FieldProperties
	SYSDbVersion_props     FieldProperties
	/// Lookups
	MailUnread_lookup    []Lookup_Item
	MailAcknowledged_lookup    []Lookup_Item
	
	/// END OF DEFINITIONS
	///
}