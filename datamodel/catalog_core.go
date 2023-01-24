package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/catalog.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Catalog (catalog)
// Endpoint 	        : Catalog (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 24/01/2023 at 13:18:07
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Catalog defines the datamolde for the Catalog object
type Catalog struct {


ID       string
ID_props FieldProperties
Endpoint       string
Endpoint_props FieldProperties
Descr       string
Descr_props FieldProperties
Query       string
Query_props FieldProperties
Source       string
Source_props FieldProperties
 // Any lookups will be added below





}

const (
	Catalog_Title       = "API Catalog"
	Catalog_SQLTable    = ""
	Catalog_SQLSearchID = "ID"
	Catalog_QueryString = "ID"
	///
	/// Handler Defintions
	///
	Catalog_Template     = "Catalog"
	Catalog_TemplateList = "/Catalog/Catalog_List"
	Catalog_TemplateView = "/Catalog/Catalog_View"
	Catalog_TemplateEdit = "/Catalog/Catalog_Edit"
	Catalog_TemplateNew  = "/Catalog/Catalog_New"
	///
	/// Handler Monitor Paths
	///
	Catalog_Path       = "/API/Catalog/"
	Catalog_PathList   = "/CatalogList/"
	Catalog_PathView   = "/CatalogView/"
	Catalog_PathEdit   = "/CatalogEdit/"
	Catalog_PathNew    = "/CatalogNew/"
	Catalog_PathSave   = "/CatalogSave/"
	Catalog_PathDelete = "/CatalogDelete/"
	///
	//Catalog_Redirect provides a page to return to aftern an action
	Catalog_Redirect = Catalog_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	Catalog_ID_sql   = "ID" // ID is a String
	Catalog_Endpoint_sql   = "Endpoint" // Endpoint is a String
	Catalog_Descr_sql   = "Descr" // Descr is a String
	Catalog_Query_sql   = "Query" // Query is a String
	Catalog_Source_sql   = "Source" // Source is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	Catalog_ID_scrn   = "ID" // ID is a String
	Catalog_Endpoint_scrn   = "Endpoint" // Endpoint is a String
	Catalog_Descr_scrn   = "Descr" // Descr is a String
	Catalog_Query_scrn   = "Query" // Query is a String
	Catalog_Source_scrn   = "Source" // Source is a String

	/// Definitions End
	///
)

//catalog_PageList provides the information for the template for a list of Catalogs
type Catalog_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []Catalog
	Context	 appContext
	BlockEdit		 bool
	BlockSave 		 bool
	BlockCreate		 bool
	BlockDelete		 bool
	BlockValidate	 bool
	BlockView		 bool
}

//catalog_Page provides the information for the template for an individual Catalog
type Catalog_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//	
	ID         string
	ID_props     FieldProperties
	Endpoint         string
	Endpoint_props     FieldProperties
	Descr         string
	Descr_props     FieldProperties
	Query         string
	Query_props     FieldProperties
	Source         string
	Source_props     FieldProperties
	// 
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 		 appContext
	BlockEdit		 bool
	BlockSave 		 bool
	BlockCreate		 bool
	BlockDelete		 bool
	BlockValidate	 bool
	BlockView		 bool
}