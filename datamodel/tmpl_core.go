package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/tmpl.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Tmpl (tmpl)
// Endpoint 	        : Tmpl (TemplateID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:58
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// Tmpl defines the datamolde for the Tmpl object
type Tmpl struct {
	SYSId                string
	SYSId_props          FieldProperties
	FIELD1               string
	FIELD1_props         FieldProperties
	FIELD2               string
	FIELD2_props         FieldProperties
	SYSCreated           string
	SYSCreated_props     FieldProperties
	SYSCreatedBy         string
	SYSCreatedBy_props   FieldProperties
	SYSCreatedHost       string
	SYSCreatedHost_props FieldProperties
	SYSUpdated           string
	SYSUpdated_props     FieldProperties
	SYSUpdatedHost       string
	SYSUpdatedHost_props FieldProperties
	SYSUpdatedBy         string
	SYSUpdatedBy_props   FieldProperties
	ID                   string
	ID_props             FieldProperties
	ExtraField           string
	ExtraField_props     FieldProperties
	ExtraField2          string
	ExtraField2_props    FieldProperties
	ExtraField3          string
	ExtraField3_props    FieldProperties
	TDate                string
	TDate_props          FieldProperties
	// Any lookups will be added below
	FIELD1_lookup []Lookup_Item
	FIELD2_lookup []Lookup_Item
}

const (
	Tmpl_Title       = "Template Contents"
	Tmpl_SQLTable    = "template"
	Tmpl_SQLSearchID = "ID"
	Tmpl_QueryString = "TemplateID"
	///
	/// Handler Defintions
	///
	Tmpl_Template     = "Tmpl"
	Tmpl_TemplateList = "/Tmpl/Tmpl_List"
	Tmpl_TemplateView = "/Tmpl/Tmpl_View"
	Tmpl_TemplateEdit = "/Tmpl/Tmpl_Edit"
	Tmpl_TemplateNew  = "/Tmpl/Tmpl_New"
	///
	/// Handler Monitor Paths
	///
	Tmpl_Path       = "/API/Tmpl/"
	Tmpl_PathList   = "/TmplList/"
	Tmpl_PathView   = "/TmplView/"
	Tmpl_PathEdit   = "/TmplEdit/"
	Tmpl_PathNew    = "/TmplNew/"
	Tmpl_PathSave   = "/TmplSave/"
	Tmpl_PathDelete = "/TmplDelete/"
	///
	//Tmpl_Redirect provides a page to return to aftern an action
	Tmpl_Redirect = Tmpl_PathList

	///
	///
	/// SQL Field Definitions
	///
	Tmpl_SYSId_sql          = "_id"          // SYSId is a Int
	Tmpl_FIELD1_sql         = "FIELD1"       // FIELD1 is a String
	Tmpl_FIELD2_sql         = "FIELD2"       // FIELD2 is a String
	Tmpl_SYSCreated_sql     = "_created"     // SYSCreated is a String
	Tmpl_SYSCreatedBy_sql   = "_createdBy"   // SYSCreatedBy is a String
	Tmpl_SYSCreatedHost_sql = "_createdHost" // SYSCreatedHost is a String
	Tmpl_SYSUpdated_sql     = "_updated"     // SYSUpdated is a String
	Tmpl_SYSUpdatedHost_sql = "_updatedHost" // SYSUpdatedHost is a String
	Tmpl_SYSUpdatedBy_sql   = "_updatedBy"   // SYSUpdatedBy is a String
	Tmpl_ID_sql             = "ID"           // ID is a String
	Tmpl_ExtraField_sql     = "ExtraField"   // ExtraField is a String
	Tmpl_ExtraField2_sql    = "ExtraField2"  // ExtraField2 is a String
	Tmpl_ExtraField3_sql    = "ExtraField3"  // ExtraField3 is a String
	Tmpl_TDate_sql          = "TDate"        // TDate is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	Tmpl_SYSId_scrn          = "SYSId"          // SYSId is a Int
	Tmpl_FIELD1_scrn         = "FIELD1"         // FIELD1 is a String
	Tmpl_FIELD2_scrn         = "FIELD2"         // FIELD2 is a String
	Tmpl_SYSCreated_scrn     = "SYSCreated"     // SYSCreated is a String
	Tmpl_SYSCreatedBy_scrn   = "SYSCreatedBy"   // SYSCreatedBy is a String
	Tmpl_SYSCreatedHost_scrn = "SYSCreatedHost" // SYSCreatedHost is a String
	Tmpl_SYSUpdated_scrn     = "SYSUpdated"     // SYSUpdated is a String
	Tmpl_SYSUpdatedHost_scrn = "SYSUpdatedHost" // SYSUpdatedHost is a String
	Tmpl_SYSUpdatedBy_scrn   = "SYSUpdatedBy"   // SYSUpdatedBy is a String
	Tmpl_ID_scrn             = "ID"             // ID is a String
	Tmpl_ExtraField_scrn     = "ExtraField"     // ExtraField is a String
	Tmpl_ExtraField2_scrn    = "ExtraField2"    // ExtraField2 is a String
	Tmpl_ExtraField3_scrn    = "ExtraField3"    // ExtraField3 is a String
	Tmpl_TDate_scrn          = "TDate"          // TDate is a String

	/// Definitions End
	///
)

// tmpl_PageList provides the information for the template for a list of Tmpls
type Tmpl_PageList struct {
	SessionInfo SessionInfo
	UserMenu    AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	ItemsOnPage int
	ItemList    []Tmpl
	Context     appContext
}

// tmpl_Page provides the information for the template for an individual Tmpl
type Tmpl_Page struct {
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
	FIELD1               string
	FIELD1_lookup        []Lookup_Item
	FIELD1_props         FieldProperties
	FIELD2               string
	FIELD2_lookup        []Lookup_Item
	FIELD2_props         FieldProperties
	SYSCreated           string
	SYSCreated_props     FieldProperties
	SYSCreatedBy         string
	SYSCreatedBy_props   FieldProperties
	SYSCreatedHost       string
	SYSCreatedHost_props FieldProperties
	SYSUpdated           string
	SYSUpdated_props     FieldProperties
	SYSUpdatedHost       string
	SYSUpdatedHost_props FieldProperties
	SYSUpdatedBy         string
	SYSUpdatedBy_props   FieldProperties
	ID                   string
	ID_props             FieldProperties
	ExtraField           string
	ExtraField_props     FieldProperties
	ExtraField2          string
	ExtraField2_props    FieldProperties
	ExtraField3          string
	ExtraField3_props    FieldProperties
	TDate                string
	TDate_props          FieldProperties
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	Context appContext
}
