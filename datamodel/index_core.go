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
// Date & Time		    : 25/01/2023 at 14:40:46
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Index defines the datamolde for the Index object
type Index struct {


SYSId       string
SYSId_props FieldProperties
IndexID       string
IndexID_props FieldProperties
KeyClass       string
KeyClass_props FieldProperties
KeyName       string
KeyName_props FieldProperties
KeyID       string
KeyID_props FieldProperties
Link       string
Link_props FieldProperties
SYSCreated       string
SYSCreated_props FieldProperties
SYSCreatedBy       string
SYSCreatedBy_props FieldProperties
SYSCreatedHost       string
SYSCreatedHost_props FieldProperties
SYSUpdated       string
SYSUpdated_props FieldProperties
SYSUpdatedBy       string
SYSUpdatedBy_props FieldProperties
SYSUpdatedHost       string
SYSUpdatedHost_props FieldProperties
SYSDeleted       string
SYSDeleted_props FieldProperties
SYSDeletedBy       string
SYSDeletedBy_props FieldProperties
SYSDeletedHost       string
SYSDeletedHost_props FieldProperties
SYSDbVersion       string
SYSDbVersion_props FieldProperties
KeyValue       string
KeyValue_props FieldProperties
 // Any lookups will be added below

















}

const (
	Index_Title       = "Index"
	Index_SQLTable    = "indexStore"
	Index_SQLSearchID = "indexID"
	Index_QueryString = "IndexID"
	///
	/// Handler Defintions
	///
	Index_Template     = "Index"
	Index_TemplateList = "/Index/Index_List"
	Index_TemplateView = "/Index/Index_View"
	Index_TemplateEdit = "/Index/Index_Edit"
	Index_TemplateNew  = "/Index/Index_New"
	///
	/// Handler Monitor Paths
	///
	Index_Path       = "/API/Index/"
	Index_PathList   = "/IndexList/"
	Index_PathView   = "/IndexView/"
	Index_PathEdit   = "/IndexEdit/"
	Index_PathNew    = "/IndexNew/"
	Index_PathSave   = "/IndexSave/"
	Index_PathDelete = "/IndexDelete/"
	///
	//Index_Redirect provides a page to return to aftern an action
	Index_Redirect = Index_PathList
	
	///
	///
	/// SQL Field Definitions
	///
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

	/// Definitions End
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

	/// Definitions End
	///
)

//index_PageList provides the information for the template for a list of Indexs
type Index_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []Index
	Context	 appContext
	BlockEdit		 bool
	BlockSave 		 bool
	BlockCreate		 bool
	BlockDelete		 bool
	BlockValidate	 bool
	BlockView		 bool
}

//index_Page provides the information for the template for an individual Index
type Index_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 25/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	SYSId_props     FieldProperties
	IndexID         string
	IndexID_props     FieldProperties
	KeyClass         string
	KeyClass_props     FieldProperties
	KeyName         string
	KeyName_props     FieldProperties
	KeyID         string
	KeyID_props     FieldProperties
	Link         string
	Link_props     FieldProperties
	SYSCreated         string
	SYSCreated_props     FieldProperties
	SYSCreatedBy         string
	SYSCreatedBy_props     FieldProperties
	SYSCreatedHost         string
	SYSCreatedHost_props     FieldProperties
	SYSUpdated         string
	SYSUpdated_props     FieldProperties
	SYSUpdatedBy         string
	SYSUpdatedBy_props     FieldProperties
	SYSUpdatedHost         string
	SYSUpdatedHost_props     FieldProperties
	SYSDeleted         string
	SYSDeleted_props     FieldProperties
	SYSDeletedBy         string
	SYSDeletedBy_props     FieldProperties
	SYSDeletedHost         string
	SYSDeletedHost_props     FieldProperties
	SYSDbVersion         string
	SYSDbVersion_props     FieldProperties
	KeyValue         string
	KeyValue_props     FieldProperties
	// 
	// Dynamically generated 25/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 		 appContext
	BlockEdit		 bool
	BlockSave 		 bool
	BlockCreate		 bool
	BlockDelete		 bool
	BlockValidate	 bool
	BlockView		 bool
}