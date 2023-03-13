package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/credentialsaction.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CredentialsAction (credentialsaction)
// Endpoint 	        : CredentialsAction (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 13/03/2023 at 14:22:25
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//CredentialsAction defines the datamodel for the CredentialsAction object
type CredentialsAction struct {
	// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	ID       string
	User       string
	Action       string
	Notes       string
	//
	// Field Properties
	//
	ID_props FieldProperties
	User_props FieldProperties
	Action_props FieldProperties
	Notes_props FieldProperties
	//
 	// Any lookups will be added below
	//
	User_lookup []Lookup_Item
	Action_lookup []Lookup_Item
	}

const (
	// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// Page Infrastructure
	CredentialsAction_Name      = "CredentialsAction"
	CredentialsAction_Title       = "Credentials Actions"
	CredentialsAction_SQLTable    = "inboxMessages"
	CredentialsAction_SQLSearchID = "MailId"
	CredentialsAction_QueryString = "ID"
	///
	/// Template Path Defintions
	///
	CredentialsAction_Template     = "CredentialsAction"
	CredentialsAction_TemplateList = "/CredentialsAction/CredentialsActionList"
	CredentialsAction_TemplateView = "/CredentialsAction/CredentialsActionView"
	CredentialsAction_TemplateEdit = "/CredentialsAction/CredentialsActionEdit"
	CredentialsAction_TemplateNew  = "/CredentialsAction/CredentialsActionNew"
	///
	/// URI Handler Paths
	///
	CredentialsAction_Path       = "/API/CredentialsAction/"
	CredentialsAction_PathList   = "/CredentialsActionList/"
	CredentialsAction_PathView   = "/CredentialsActionView/"
	CredentialsAction_PathEdit   = "/CredentialsActionEdit/"
	CredentialsAction_PathNew    = "/CredentialsActionNew/"
	CredentialsAction_PathSave   = "/CredentialsActionSave/"
	CredentialsAction_PathDelete = "/CredentialsActionDelete/"
	// Redirects - On Server Side Error
	CredentialsAction_PathEditException   = "/CredentialsActionEditException/"
	CredentialsAction_PathNewException    = "/CredentialsActionNewException/"
	//
	//CredentialsAction_Redirect provides a page to return to aftern an action
	CredentialsAction_Redirect = CredentialsAction_PathView
	//
	//
	// SQL Field Definitions
	//
	CredentialsAction_ID_sql   = "ID" // ID is a String
	CredentialsAction_User_sql   = "User" // User is a String
	CredentialsAction_Action_sql   = "Action" // Action is a String
	CredentialsAction_Notes_sql   = "Notes" // Notes is a String
///
	/// Application Field Definitions
	///
	CredentialsAction_ID_scrn   = "ID" // ID is a String
	CredentialsAction_User_scrn   = "User" // User is a String
	CredentialsAction_Action_scrn   = "Action" // Action is a String
	CredentialsAction_Notes_scrn   = "Notes" // Notes is a String
///
)

//CredentialsAction_PageList provides the information for the template for a list of CredentialsActions
type CredentialsAction_PageList struct {
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
	ItemList  		 []CredentialsAction
}

//CredentialsAction_Page provides the information for the template for an individual CredentialsAction
type CredentialsAction_Page struct {
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
	ID         string
	User         string
	Action         string
	Notes         string
	/// Field Properties
	ID_props     FieldProperties
	User_props     FieldProperties
	Action_props     FieldProperties
	Notes_props     FieldProperties
	/// Lookups
	User_lookup    []Lookup_Item
	Action_lookup    []Lookup_Item
	
	/// END OF DEFINITIONS
	///
}