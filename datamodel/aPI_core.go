package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/api.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : API (api)
// Endpoint 	        : API (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:07
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

// Centre is cheese
type API struct {
	ID       string
	Endpoint string
	Descr    string
	Query    string
	Source   string
}

const (
	API_Title       = "API"
	API_SQLTable    = "API"
	API_SQLSearchID = "ID"
	API_QueryString = "ID"
	///
	/// Handler Defintions
	///
	API_Template     = "API"
	API_TemplateList = "API_List"
	API_TemplateView = "API_View"
	API_TemplateEdit = "API_Edit"
	API_TemplateNew  = "API_New"
	///
	/// Handler Monitor Paths
	///
	API_Path       = "/API/API/"
	API_PathList   = "/APIList/"
	API_PathView   = "/APIView/"
	API_PathEdit   = "/APIEdit/"
	API_PathNew    = "/APINew/"
	API_PathSave   = "/APISave/"
	API_PathDelete = "/APIDelete/"
	///
	/// SQL Field Definitions
	///
	API_ID       = "ID"       // ID is a String
	API_Endpoint = "Endpoint" // Endpoint is a String
	API_Descr    = "Descr"    // Descr is a String
	API_Query    = "Query"    // Query is a String
	API_Source   = "Source"   // Source is a String

	/// Definitions End
)
