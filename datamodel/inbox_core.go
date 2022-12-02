package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/inbox.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Inbox (inbox)
// Endpoint 	        : Inbox (Message)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 01/12/2022 at 09:40:01
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Inbox defines the datamolde for the Inbox object
type Inbox struct {


SYSId       string
SYSId_props FieldProperties
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
MailId       string
MailId_props FieldProperties
MailTo       string
MailTo_props FieldProperties
MailFrom       string
MailFrom_props FieldProperties
MailSource       string
MailSource_props FieldProperties
MailSent       string
MailSent_props FieldProperties
MailUnread       string
MailUnread_props FieldProperties
MailSubject       string
MailSubject_props FieldProperties
MailContent       string
MailContent_props FieldProperties
MailAcknowledged       string
MailAcknowledged_props FieldProperties
SYSDeleted       string
SYSDeleted_props FieldProperties
SYSDeletedBy       string
SYSDeletedBy_props FieldProperties
SYSDeletedHost       string
SYSDeletedHost_props FieldProperties
 // Any lookups will be added below











MailUnread_lookup []Lookup_Item


MailAcknowledged_lookup []Lookup_Item




}

const (
	Inbox_Title       = "Inbox"
	Inbox_SQLTable    = "inboxMessages"
	Inbox_SQLSearchID = "MailId"
	Inbox_QueryString = "Message"
	///
	/// Handler Defintions
	///
	Inbox_Template     = "Inbox"
	Inbox_TemplateList = "/Inbox/Inbox_List"
	Inbox_TemplateView = "/Inbox/Inbox_View"
	Inbox_TemplateEdit = "/Inbox/Inbox_Edit"
	Inbox_TemplateNew  = "/Inbox/Inbox_New"
	///
	/// Handler Monitor Paths
	///
	Inbox_Path       = "/API/Inbox/"
	Inbox_PathList   = "/InboxList/"
	Inbox_PathView   = "/InboxView/"
	Inbox_PathEdit   = "/InboxEdit/"
	Inbox_PathNew    = "/InboxNew/"
	Inbox_PathSave   = "/InboxSave/"
	Inbox_PathDelete = "/InboxDelete/"
	///
	//Inbox_Redirect provides a page to return to aftern an action
	Inbox_Redirect = Inbox_PathList
	
	///
	///
	/// SQL Field Definitions
	///
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

	/// Definitions End
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

	/// Definitions End
	///
)

//inbox_PageList provides the information for the template for a list of Inboxs
type Inbox_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []Inbox
	Context	 appContext
}

//inbox_Page provides the information for the template for an individual Inbox
type Inbox_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	SYSId_props     FieldProperties
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
	MailId         string
	MailId_props     FieldProperties
	MailTo         string
	MailTo_props     FieldProperties
	MailFrom         string
	MailFrom_props     FieldProperties
	MailSource         string
	MailSource_props     FieldProperties
	MailSent         string
	MailSent_props     FieldProperties
	MailUnread         string
	MailUnread_lookup    []Lookup_Item
	MailUnread_props     FieldProperties
	MailSubject         string
	MailSubject_props     FieldProperties
	MailContent         string
	MailContent_props     FieldProperties
	MailAcknowledged         string
	MailAcknowledged_lookup    []Lookup_Item
	MailAcknowledged_props     FieldProperties
	SYSDeleted         string
	SYSDeleted_props     FieldProperties
	SYSDeletedBy         string
	SYSDeletedBy_props     FieldProperties
	SYSDeletedHost         string
	SYSDeletedHost_props     FieldProperties
	// 
	// Dynamically generated 01/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}