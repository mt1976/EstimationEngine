package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/doctype.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DocType (doctype)
// Endpoint 	        : DocType (DocTypeID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 10/03/2023 at 19:54:32
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//DocType defines the datamodel for the DocType object
type DocType struct {
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	SYSId       string
	DocTypeID       string
	Code       string
	Name       string
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
	Comments       string
	//
	// Field Properties
	//
	SYSId_props FieldProperties
	DocTypeID_props FieldProperties
	Code_props FieldProperties
	Name_props FieldProperties
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
	Comments_props FieldProperties
	//
 	// Any lookups will be added below
	//
	}

const (
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// Page Infrastructure
	DocType_Name      = "DocType"
	DocType_Title       = "Document Types"
	DocType_SQLTable    = "docTypeStore"
	DocType_SQLSearchID = "docTypeID"
	DocType_QueryString = "DocTypeID"
	///
	/// Template Path Defintions
	///
	DocType_Template     = "DocType"
	DocType_TemplateList = "/DocType/DocTypeList"
	DocType_TemplateView = "/DocType/DocTypeView"
	DocType_TemplateEdit = "/DocType/DocTypeEdit"
	DocType_TemplateNew  = "/DocType/DocTypeNew"
	///
	/// URI Handler Paths
	///
	DocType_Path       = "/API/DocType/"
	DocType_PathList   = "/DocTypeList/"
	DocType_PathView   = "/DocTypeView/"
	DocType_PathEdit   = "/DocTypeEdit/"
	DocType_PathNew    = "/DocTypeNew/"
	DocType_PathSave   = "/DocTypeSave/"
	DocType_PathDelete = "/DocTypeDelete/"
	// Redirects - On Server Side Error
	DocType_PathEditException   = "/DocTypeEditException/"
	DocType_PathNewException    = "/DocTypeNewException/"
	//
	//DocType_Redirect provides a page to return to aftern an action
	DocType_Redirect = DocType_PathList
	
	//
	//
	// SQL Field Definitions
	//
	DocType_SYSId_sql   = "_id" // SYSId is a Int
	DocType_DocTypeID_sql   = "docTypeID" // DocTypeID is a String
	DocType_Code_sql   = "code" // Code is a String
	DocType_Name_sql   = "name" // Name is a String
	DocType_SYSCreated_sql   = "_created" // SYSCreated is a String
	DocType_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	DocType_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	DocType_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	DocType_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
	DocType_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
	DocType_SYSDeleted_sql   = "_deleted" // SYSDeleted is a String
	DocType_SYSDeletedBy_sql   = "_deletedBy" // SYSDeletedBy is a String
	DocType_SYSDeletedHost_sql   = "_deletedHost" // SYSDeletedHost is a String
	DocType_SYSDbVersion_sql   = "_dbVersion" // SYSDbVersion is a String
	DocType_Comments_sql   = "comments" // Comments is a String
///
	/// Application Field Definitions
	///
	DocType_SYSId_scrn   = "SYSId" // SYSId is a Int
	DocType_DocTypeID_scrn   = "DocTypeID" // DocTypeID is a String
	DocType_Code_scrn   = "Code" // Code is a String
	DocType_Name_scrn   = "Name" // Name is a String
	DocType_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	DocType_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	DocType_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	DocType_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	DocType_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
	DocType_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
	DocType_SYSDeleted_scrn   = "SYSDeleted" // SYSDeleted is a String
	DocType_SYSDeletedBy_scrn   = "SYSDeletedBy" // SYSDeletedBy is a String
	DocType_SYSDeletedHost_scrn   = "SYSDeletedHost" // SYSDeletedHost is a String
	DocType_SYSDbVersion_scrn   = "SYSDbVersion" // SYSDbVersion is a String
	DocType_Comments_scrn   = "Comments" // Comments is a String
///
)

//DocType_PageList provides the information for the template for a list of DocTypes
type DocType_PageList struct {
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
	ItemList  		 []DocType
}

//DocType_Page provides the information for the template for an individual DocType
type DocType_Page struct {
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
	DocTypeID         string
	Code         string
	Name         string
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
	Comments         string
	/// Field Properties
	SYSId_props     FieldProperties
	DocTypeID_props     FieldProperties
	Code_props     FieldProperties
	Name_props     FieldProperties
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
	Comments_props     FieldProperties
	/// Lookups
	
	/// END OF DEFINITIONS
	///
}