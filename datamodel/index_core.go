package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/index.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Index (index)
// Endpoint 	        : Index (IndexID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 10/03/2023 at 19:54:34
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Index defines the datamodel for the Index object
type Index struct {
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	SYSId       string
	IndexID       string
	KeyClass       string
	KeyName       string
	KeyID       string
	Link       string
	SYSCreated       string
	SYSCreatedBy       string
	SYSCreatedHost       string
	SYSUpdated       string
	SYSUpdatedBy       string
	SYSUpdatedHost       string
	SYSDeleted       string
	SYSDeletedBy       string
	SYSDeletedHost       string
	SYSDbVersion       string
	KeyValue       string
	//
	// Field Properties
	//
	SYSId_props FieldProperties
	IndexID_props FieldProperties
	KeyClass_props FieldProperties
	KeyName_props FieldProperties
	KeyID_props FieldProperties
	Link_props FieldProperties
	SYSCreated_props FieldProperties
	SYSCreatedBy_props FieldProperties
	SYSCreatedHost_props FieldProperties
	SYSUpdated_props FieldProperties
	SYSUpdatedBy_props FieldProperties
	SYSUpdatedHost_props FieldProperties
	SYSDeleted_props FieldProperties
	SYSDeletedBy_props FieldProperties
	SYSDeletedHost_props FieldProperties
	SYSDbVersion_props FieldProperties
	KeyValue_props FieldProperties
	//
 	// Any lookups will be added below
	//
	}

const (
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// Page Infrastructure
	Index_Name      = "Index"
	Index_Title       = "Index"
	Index_SQLTable    = "indexStore"
	Index_SQLSearchID = "indexID"
	Index_QueryString = "IndexID"
	///
	/// Template Path Defintions
	///
	Index_Template     = "Index"
	Index_TemplateList = "/Index/IndexList"
	Index_TemplateView = "/Index/IndexView"
	Index_TemplateEdit = "/Index/IndexEdit"
	Index_TemplateNew  = "/Index/IndexNew"
	///
	/// URI Handler Paths
	///
	Index_Path       = "/API/Index/"
	Index_PathList   = "/IndexList/"
	Index_PathView   = "/IndexView/"
	Index_PathEdit   = "/IndexEdit/"
	Index_PathNew    = "/IndexNew/"
	Index_PathSave   = "/IndexSave/"
	Index_PathDelete = "/IndexDelete/"
	// Redirects - On Server Side Error
	Index_PathEditException   = "/IndexEditException/"
	Index_PathNewException    = "/IndexNewException/"
	//
	//Index_Redirect provides a page to return to aftern an action
	Index_Redirect = Index_PathList
	
	//
	//
	// SQL Field Definitions
	//
	Index_SYSId_sql   = "_id" // SYSId is a Int
	Index_IndexID_sql   = "indexID" // IndexID is a String
	Index_KeyClass_sql   = "keyClass" // KeyClass is a String
	Index_KeyName_sql   = "keyName" // KeyName is a String
	Index_KeyID_sql   = "keyID" // KeyID is a String
	Index_Link_sql   = "link" // Link is a String
	Index_SYSCreated_sql   = "_created" // SYSCreated is a String
	Index_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	Index_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	Index_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	Index_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
	Index_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
	Index_SYSDeleted_sql   = "_deleted" // SYSDeleted is a String
	Index_SYSDeletedBy_sql   = "_deletedBy" // SYSDeletedBy is a String
	Index_SYSDeletedHost_sql   = "_deletedHost" // SYSDeletedHost is a String
	Index_SYSDbVersion_sql   = "_dbVersion" // SYSDbVersion is a String
	Index_KeyValue_sql   = "KeyValue" // KeyValue is a String
///
	/// Application Field Definitions
	///
	Index_SYSId_scrn   = "SYSId" // SYSId is a Int
	Index_IndexID_scrn   = "IndexID" // IndexID is a String
	Index_KeyClass_scrn   = "KeyClass" // KeyClass is a String
	Index_KeyName_scrn   = "KeyName" // KeyName is a String
	Index_KeyID_scrn   = "KeyID" // KeyID is a String
	Index_Link_scrn   = "Link" // Link is a String
	Index_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	Index_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	Index_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	Index_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	Index_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
	Index_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
	Index_SYSDeleted_scrn   = "SYSDeleted" // SYSDeleted is a String
	Index_SYSDeletedBy_scrn   = "SYSDeletedBy" // SYSDeletedBy is a String
	Index_SYSDeletedHost_scrn   = "SYSDeletedHost" // SYSDeletedHost is a String
	Index_SYSDbVersion_scrn   = "SYSDbVersion" // SYSDbVersion is a String
	Index_KeyValue_scrn   = "KeyValue" // KeyValue is a String
///
)

//Index_PageList provides the information for the template for a list of Indexs
type Index_PageList struct {
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
	ItemList  		 []Index
}

//Index_Page provides the information for the template for an individual Index
type Index_Page struct {
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
	SYSId         string
	IndexID         string
	KeyClass         string
	KeyName         string
	KeyID         string
	Link         string
	SYSCreated         string
	SYSCreatedBy         string
	SYSCreatedHost         string
	SYSUpdated         string
	SYSUpdatedBy         string
	SYSUpdatedHost         string
	SYSDeleted         string
	SYSDeletedBy         string
	SYSDeletedHost         string
	SYSDbVersion         string
	KeyValue         string
	/// Field Properties
	SYSId_props     FieldProperties
	IndexID_props     FieldProperties
	KeyClass_props     FieldProperties
	KeyName_props     FieldProperties
	KeyID_props     FieldProperties
	Link_props     FieldProperties
	SYSCreated_props     FieldProperties
	SYSCreatedBy_props     FieldProperties
	SYSCreatedHost_props     FieldProperties
	SYSUpdated_props     FieldProperties
	SYSUpdatedBy_props     FieldProperties
	SYSUpdatedHost_props     FieldProperties
	SYSDeleted_props     FieldProperties
	SYSDeletedBy_props     FieldProperties
	SYSDeletedHost_props     FieldProperties
	SYSDbVersion_props     FieldProperties
	KeyValue_props     FieldProperties
	/// Lookups
	
	/// END OF DEFINITIONS
	///
}