package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/confidence.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Confidence (confidence)
// Endpoint 	        : Confidence (ConfidenceID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 09/02/2023 at 15:57:32
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// Confidence defines the datamolde for the Confidence object
type Confidence struct {
	// Dynamically generated 09/02/2023 by matttownsend (Matt Townsend) on silicon.local
	// Field Definitions
	SYSId                string
	SYSId_props          FieldProperties
	ConfidenceID         string
	ConfidenceID_props   FieldProperties
	Code                 string
	Code_props           FieldProperties
	Name                 string
	Name_props           FieldProperties
	Perc                 string
	Perc_props           FieldProperties
	SYSCreated           string
	SYSCreated_props     FieldProperties
	SYSCreatedBy         string
	SYSCreatedBy_props   FieldProperties
	SYSCreatedHost       string
	SYSCreatedHost_props FieldProperties
	SYSUpdated           string
	SYSUpdated_props     FieldProperties
	SYSUpdatedBy         string
	SYSUpdatedBy_props   FieldProperties
	SYSUpdatedHost       string
	SYSUpdatedHost_props FieldProperties
	SYSDeleted           string
	SYSDeleted_props     FieldProperties
	SYSDeletedBy         string
	SYSDeletedBy_props   FieldProperties
	SYSDeletedHost       string
	SYSDeletedHost_props FieldProperties
	Notes                string
	Notes_props          FieldProperties
	SYSDbVersion         string
	SYSDbVersion_props   FieldProperties
	Comments             string
	Comments_props       FieldProperties // Any lookups will be added below
}

const (
	// Dynamically generated 09/02/2023 by matttownsend (Matt Townsend) on silicon.local
	// Page Infrastructure
	Confidence_Title       = "Confidence"
	Confidence_SQLTable    = "confidenceStore"
	Confidence_SQLSearchID = "confidenceID"
	Confidence_QueryString = "ConfidenceID"
	///
	/// Handler Defintions
	///
	Confidence_Template     = "Confidence"
	Confidence_TemplateList = "/Confidence/Confidence_List"
	Confidence_TemplateView = "/Confidence/Confidence_View"
	Confidence_TemplateEdit = "/Confidence/Confidence_Edit"
	Confidence_TemplateNew  = "/Confidence/Confidence_New"
	///
	/// Handler Monitor Paths
	///
	Confidence_Path       = "/API/Confidence/"
	Confidence_PathList   = "/ConfidenceList/"
	Confidence_PathView   = "/ConfidenceView/"
	Confidence_PathEdit   = "/ConfidenceEdit/"
	Confidence_PathNew    = "/ConfidenceNew/"
	Confidence_PathSave   = "/ConfidenceSave/"
	Confidence_PathDelete = "/ConfidenceDelete/"
	//
	//Confidence_Redirect provides a page to return to aftern an action
	Confidence_Redirect = Confidence_PathList

	//
	//
	// SQL Field Definitions
	//
	Confidence_SYSId_sql          = "_id"          // SYSId is a Int
	Confidence_ConfidenceID_sql   = "confidenceID" // ConfidenceID is a String
	Confidence_Code_sql           = "code"         // Code is a String
	Confidence_Name_sql           = "name"         // Name is a String
	Confidence_Perc_sql           = "perc"         // Perc is a String
	Confidence_SYSCreated_sql     = "_created"     // SYSCreated is a String
	Confidence_SYSCreatedBy_sql   = "_createdBy"   // SYSCreatedBy is a String
	Confidence_SYSCreatedHost_sql = "_createdHost" // SYSCreatedHost is a String
	Confidence_SYSUpdated_sql     = "_updated"     // SYSUpdated is a String
	Confidence_SYSUpdatedBy_sql   = "_updatedBy"   // SYSUpdatedBy is a String
	Confidence_SYSUpdatedHost_sql = "_updatedHost" // SYSUpdatedHost is a String
	Confidence_SYSDeleted_sql     = "_deleted"     // SYSDeleted is a String
	Confidence_SYSDeletedBy_sql   = "_deletedBy"   // SYSDeletedBy is a String
	Confidence_SYSDeletedHost_sql = "_deletedHost" // SYSDeletedHost is a String
	Confidence_Notes_sql          = "notes"        // Notes is a String
	Confidence_SYSDbVersion_sql   = "_dbVersion"   // SYSDbVersion is a String
	Confidence_Comments_sql       = "comments"     // Comments is a String
	///
	/// Application Field Definitions
	///
	Confidence_SYSId_scrn          = "SYSId"          // SYSId is a Int
	Confidence_ConfidenceID_scrn   = "ConfidenceID"   // ConfidenceID is a String
	Confidence_Code_scrn           = "Code"           // Code is a String
	Confidence_Name_scrn           = "Name"           // Name is a String
	Confidence_Perc_scrn           = "Perc"           // Perc is a String
	Confidence_SYSCreated_scrn     = "SYSCreated"     // SYSCreated is a String
	Confidence_SYSCreatedBy_scrn   = "SYSCreatedBy"   // SYSCreatedBy is a String
	Confidence_SYSCreatedHost_scrn = "SYSCreatedHost" // SYSCreatedHost is a String
	Confidence_SYSUpdated_scrn     = "SYSUpdated"     // SYSUpdated is a String
	Confidence_SYSUpdatedBy_scrn   = "SYSUpdatedBy"   // SYSUpdatedBy is a String
	Confidence_SYSUpdatedHost_scrn = "SYSUpdatedHost" // SYSUpdatedHost is a String
	Confidence_SYSDeleted_scrn     = "SYSDeleted"     // SYSDeleted is a String
	Confidence_SYSDeletedBy_scrn   = "SYSDeletedBy"   // SYSDeletedBy is a String
	Confidence_SYSDeletedHost_scrn = "SYSDeletedHost" // SYSDeletedHost is a String
	Confidence_Notes_scrn          = "Notes"          // Notes is a String
	Confidence_SYSDbVersion_scrn   = "SYSDbVersion"   // SYSDbVersion is a String
	Confidence_Comments_scrn       = "Comments"       // Comments is a String
// /
)

// confidence_PageList provides the information for the template for a list of Confidences
type Confidence_PageList struct {
	// Dynamically generated 09/02/2023 by matttownsend (Matt Townsend) on silicon.local
	// Page Infrastructure
	SessionInfo SessionInfo
	UserMenu    AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	// Context & Permisions
	Context       appContext
	BlockEdit     bool
	BlockSave     bool
	BlockCreate   bool
	BlockDelete   bool
	BlockValidate bool
	BlockView     bool
	// Page Data
	ItemsOnPage int
	ItemList    []Confidence
}

// confidence_Page provides the information for the template for an individual Confidence
type Confidence_Page struct {
	// Dynamically generated 09/02/2023 by matttownsend (Matt Townsend) on silicon.local
	// Page Infrastructure
	SessionInfo SessionInfo
	UserMenu    AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	// Context & Permisions
	Context       appContext
	BlockEdit     bool
	BlockSave     bool
	BlockCreate   bool
	BlockDelete   bool
	BlockValidate bool
	BlockView     bool
	// Field Definitions
	SYSId                string
	SYSId_props          FieldProperties
	ConfidenceID         string
	ConfidenceID_props   FieldProperties
	Code                 string
	Code_props           FieldProperties
	Name                 string
	Name_props           FieldProperties
	Perc                 string
	Perc_props           FieldProperties
	SYSCreated           string
	SYSCreated_props     FieldProperties
	SYSCreatedBy         string
	SYSCreatedBy_props   FieldProperties
	SYSCreatedHost       string
	SYSCreatedHost_props FieldProperties
	SYSUpdated           string
	SYSUpdated_props     FieldProperties
	SYSUpdatedBy         string
	SYSUpdatedBy_props   FieldProperties
	SYSUpdatedHost       string
	SYSUpdatedHost_props FieldProperties
	SYSDeleted           string
	SYSDeleted_props     FieldProperties
	SYSDeletedBy         string
	SYSDeletedBy_props   FieldProperties
	SYSDeletedHost       string
	SYSDeletedHost_props FieldProperties
	Notes                string
	Notes_props          FieldProperties
	SYSDbVersion         string
	SYSDbVersion_props   FieldProperties
	Comments             string
	Comments_props       FieldProperties
}
