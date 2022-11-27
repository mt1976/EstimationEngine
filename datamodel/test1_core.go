package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/test1.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Test1 (test1)
// Endpoint 	        : Test1 (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:58
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// Test1 defines the datamolde for the Test1 object
type Test1 struct {
	ID             string
	ID_props       FieldProperties
	Endpoint       string
	Endpoint_props FieldProperties
	Descr          string
	Descr_props    FieldProperties
	Query          string
	Query_props    FieldProperties
	Source         string
	Source_props   FieldProperties
	Firm           string
	Firm_props     FieldProperties
	YN             string
	YN_props       FieldProperties
	User           string
	User_props     FieldProperties
	Cheese         string
	Cheese_props   FieldProperties
	Onion          string
	Onion_props    FieldProperties
	// Any lookups will be added below

	Firm_lookup []Lookup_Item
	YN_lookup   []Lookup_Item
	User_lookup []Lookup_Item
}

const (
	Test1_Title       = "Test1Template"
	Test1_SQLTable    = ""
	Test1_SQLSearchID = "ID"
	Test1_QueryString = "ID"
	///
	/// Handler Defintions
	///
	Test1_Template     = "Test1"
	Test1_TemplateList = "/Test1/Test1_List"
	Test1_TemplateView = "/Test1/Test1_View"
	Test1_TemplateEdit = "/Test1/Test1_Edit"
	Test1_TemplateNew  = "/Test1/Test1_New"
	///
	/// Handler Monitor Paths
	///
	Test1_Path       = "/API/Test1/"
	Test1_PathList   = "/Test1List/"
	Test1_PathView   = "/Test1View/"
	Test1_PathEdit   = "/Test1Edit/"
	Test1_PathNew    = "/Test1New/"
	Test1_PathSave   = "/Test1Save/"
	Test1_PathDelete = "/Test1Delete/"
	///
	//Test1_Redirect provides a page to return to aftern an action
	Test1_Redirect = Test1_PathList

	///
	///
	/// SQL Field Definitions
	///
	Test1_ID_sql       = "ID"       // ID is a String
	Test1_Endpoint_sql = "Endpoint" // Endpoint is a String
	Test1_Descr_sql    = "Descr"    // Descr is a String
	Test1_Query_sql    = "Query"    // Query is a String
	Test1_Source_sql   = "Source"   // Source is a String
	Test1_Firm_sql     = "Firm"     // Firm is a String
	Test1_YN_sql       = "YN"       // YN is a String
	Test1_User_sql     = "User"     // User is a String
	Test1_Cheese_sql   = "Cheese"   // Cheese is a String
	Test1_Onion_sql    = "Onion"    // Onion is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	Test1_ID_scrn       = "ID"       // ID is a String
	Test1_Endpoint_scrn = "Endpoint" // Endpoint is a String
	Test1_Descr_scrn    = "Descr"    // Descr is a String
	Test1_Query_scrn    = "Query"    // Query is a String
	Test1_Source_scrn   = "Source"   // Source is a String
	Test1_Firm_scrn     = "Firm"     // Firm is a String
	Test1_YN_scrn       = "YN"       // YN is a String
	Test1_User_scrn     = "User"     // User is a String
	Test1_Cheese_scrn   = "Cheese"   // Cheese is a String
	Test1_Onion_scrn    = "Onion"    // Onion is a String

	/// Definitions End
	///
)

// test1_PageList provides the information for the template for a list of Test1s
type Test1_PageList struct {
	SessionInfo SessionInfo
	UserMenu    AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	ItemsOnPage int
	ItemList    []Test1
	Context     appContext
}

// test1_Page provides the information for the template for an individual Test1
type Test1_Page struct {
	SessionInfo SessionInfo
	UserMenu    AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	ID             string
	ID_props       FieldProperties
	Endpoint       string
	Endpoint_props FieldProperties
	Descr          string
	Descr_props    FieldProperties
	Query          string
	Query_props    FieldProperties
	Source         string
	Source_props   FieldProperties
	Firm           string
	Firm_lookup    []Lookup_Item
	Firm_props     FieldProperties
	YN             string
	YN_lookup      []Lookup_Item
	YN_props       FieldProperties
	User           string
	User_lookup    []Lookup_Item
	User_props     FieldProperties
	Cheese         string
	Cheese_props   FieldProperties
	Onion          string
	Onion_props    FieldProperties
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	Context appContext
}
