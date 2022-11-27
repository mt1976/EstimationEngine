package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/message.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Message (message)
// Endpoint 	        : Message (Message)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:54
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// Message defines the datamolde for the Message object
type Message struct {
	SYSId                string
	SYSId_props          FieldProperties
	Id                   string
	Id_props             FieldProperties
	Message              string
	Message_props        FieldProperties
	SYSCreated           string
	SYSCreated_props     FieldProperties
	SYSWho               string
	SYSWho_props         FieldProperties
	SYSHost              string
	SYSHost_props        FieldProperties
	SYSUpdated           string
	SYSUpdated_props     FieldProperties
	SYSCreatedBy         string
	SYSCreatedBy_props   FieldProperties
	SYSCreatedHost       string
	SYSCreatedHost_props FieldProperties
	SYSUpdatedBy         string
	SYSUpdatedBy_props   FieldProperties
	SYSUpdatedHost       string
	SYSUpdatedHost_props FieldProperties
	// Any lookups will be added below

}

const (
	Message_Title       = "System Message"
	Message_SQLTable    = "messageStore"
	Message_SQLSearchID = "id"
	Message_QueryString = "Message"
	///
	/// Handler Defintions
	///
	Message_Template     = "Message"
	Message_TemplateList = "/Message/Message_List"
	Message_TemplateView = "/Message/Message_View"
	Message_TemplateEdit = "/Message/Message_Edit"
	Message_TemplateNew  = "/Message/Message_New"
	///
	/// Handler Monitor Paths
	///
	Message_Path       = "/API/Message/"
	Message_PathList   = "/MessageList/"
	Message_PathView   = "/MessageView/"
	Message_PathEdit   = "/MessageEdit/"
	Message_PathNew    = "/MessageNew/"
	Message_PathSave   = "/MessageSave/"
	Message_PathDelete = "/MessageDelete/"
	///
	//Message_Redirect provides a page to return to aftern an action
	Message_Redirect = Message_PathList

	///
	///
	/// SQL Field Definitions
	///
	Message_SYSId_sql          = "_id"          // SYSId is a Int
	Message_Id_sql             = "Id"           // Id is a String
	Message_Message_sql        = "Message"      // Message is a String
	Message_SYSCreated_sql     = "_created"     // SYSCreated is a String
	Message_SYSWho_sql         = "_who"         // SYSWho is a String
	Message_SYSHost_sql        = "_host"        // SYSHost is a String
	Message_SYSUpdated_sql     = "_updated"     // SYSUpdated is a String
	Message_SYSCreatedBy_sql   = "_createdBy"   // SYSCreatedBy is a String
	Message_SYSCreatedHost_sql = "_createdHost" // SYSCreatedHost is a String
	Message_SYSUpdatedBy_sql   = "_updatedBy"   // SYSUpdatedBy is a String
	Message_SYSUpdatedHost_sql = "_updatedHost" // SYSUpdatedHost is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	Message_SYSId_scrn          = "SYSId"          // SYSId is a Int
	Message_Id_scrn             = "Id"             // Id is a String
	Message_Message_scrn        = "Message"        // Message is a String
	Message_SYSCreated_scrn     = "SYSCreated"     // SYSCreated is a String
	Message_SYSWho_scrn         = "SYSWho"         // SYSWho is a String
	Message_SYSHost_scrn        = "SYSHost"        // SYSHost is a String
	Message_SYSUpdated_scrn     = "SYSUpdated"     // SYSUpdated is a String
	Message_SYSCreatedBy_scrn   = "SYSCreatedBy"   // SYSCreatedBy is a String
	Message_SYSCreatedHost_scrn = "SYSCreatedHost" // SYSCreatedHost is a String
	Message_SYSUpdatedBy_scrn   = "SYSUpdatedBy"   // SYSUpdatedBy is a String
	Message_SYSUpdatedHost_scrn = "SYSUpdatedHost" // SYSUpdatedHost is a String

	/// Definitions End
	///
)

// message_PageList provides the information for the template for a list of Messages
type Message_PageList struct {
	SessionInfo SessionInfo
	UserMenu    AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	ItemsOnPage int
	ItemList    []Message
	Context     appContext
}

// message_Page provides the information for the template for an individual Message
type Message_Page struct {
	SessionInfo SessionInfo
	UserMenu    AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	SYSId                string
	SYSId_props          FieldProperties
	Id                   string
	Id_props             FieldProperties
	Message              string
	Message_props        FieldProperties
	SYSCreated           string
	SYSCreated_props     FieldProperties
	SYSWho               string
	SYSWho_props         FieldProperties
	SYSHost              string
	SYSHost_props        FieldProperties
	SYSUpdated           string
	SYSUpdated_props     FieldProperties
	SYSCreatedBy         string
	SYSCreatedBy_props   FieldProperties
	SYSCreatedHost       string
	SYSCreatedHost_props FieldProperties
	SYSUpdatedBy         string
	SYSUpdatedBy_props   FieldProperties
	SYSUpdatedHost       string
	SYSUpdatedHost_props FieldProperties
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	Context appContext
}
