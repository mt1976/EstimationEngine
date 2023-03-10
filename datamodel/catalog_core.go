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
// Date & Time		    : 10/03/2023 at 19:54:29
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Catalog defines the datamodel for the Catalog object
type Catalog struct {
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	ID       string
	Endpoint       string
	Descr       string
	Query       string
	Source       string
	//
	// Field Properties
	//
	ID_props FieldProperties
	Endpoint_props FieldProperties
	Descr_props FieldProperties
	Query_props FieldProperties
	Source_props FieldProperties
	//
 	// Any lookups will be added below
	//
	}

const (
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// Page Infrastructure
	Catalog_Name      = "Catalog"
	Catalog_Title       = "API Catalog"
	Catalog_SQLTable    = ""
	Catalog_SQLSearchID = "ID"
	Catalog_QueryString = "ID"
	///
	/// Template Path Defintions
	///
	Catalog_Template     = "Catalog"
	Catalog_TemplateList = "/Catalog/CatalogList"
	Catalog_TemplateView = "/Catalog/CatalogView"
	Catalog_TemplateEdit = "/Catalog/CatalogEdit"
	Catalog_TemplateNew  = "/Catalog/CatalogNew"
	///
	/// URI Handler Paths
	///
	Catalog_Path       = "/API/Catalog/"
	Catalog_PathList   = "/CatalogList/"
	Catalog_PathView   = "/CatalogView/"
	Catalog_PathEdit   = "/CatalogEdit/"
	Catalog_PathNew    = "/CatalogNew/"
	Catalog_PathSave   = "/CatalogSave/"
	Catalog_PathDelete = "/CatalogDelete/"
	// Redirects - On Server Side Error
	Catalog_PathEditException   = "/CatalogEditException/"
	Catalog_PathNewException    = "/CatalogNewException/"
	//
	//Catalog_Redirect provides a page to return to aftern an action
	Catalog_Redirect = Catalog_PathList
	
	//
	//
	// SQL Field Definitions
	//
	Catalog_ID_sql   = "ID" // ID is a String
	Catalog_Endpoint_sql   = "Endpoint" // Endpoint is a String
	Catalog_Descr_sql   = "Descr" // Descr is a String
	Catalog_Query_sql   = "Query" // Query is a String
	Catalog_Source_sql   = "Source" // Source is a String
///
	/// Application Field Definitions
	///
	Catalog_ID_scrn   = "ID" // ID is a String
	Catalog_Endpoint_scrn   = "Endpoint" // Endpoint is a String
	Catalog_Descr_scrn   = "Descr" // Descr is a String
	Catalog_Query_scrn   = "Query" // Query is a String
	Catalog_Source_scrn   = "Source" // Source is a String
///
)

//Catalog_PageList provides the information for the template for a list of Catalogs
type Catalog_PageList struct {
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
	ItemList  		 []Catalog
}

//Catalog_Page provides the information for the template for an individual Catalog
type Catalog_Page struct {
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
	ID         string
	Endpoint         string
	Descr         string
	Query         string
	Source         string
	/// Field Properties
	ID_props     FieldProperties
	Endpoint_props     FieldProperties
	Descr_props     FieldProperties
	Query_props     FieldProperties
	Source_props     FieldProperties
	/// Lookups
	
	/// END OF DEFINITIONS
	///
}