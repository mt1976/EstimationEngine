package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/doctype.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DocType (doctype)
// Endpoint 	        : DocType (DocTypeID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 16/12/2022 at 16:47:10
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//DocType defines the datamolde for the DocType object
type DocType struct {


SYSId       string
SYSId_props FieldProperties
DocTypeID       string
DocTypeID_props FieldProperties
Code       string
Code_props FieldProperties
Name       string
Name_props FieldProperties
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
Comments       string
Comments_props FieldProperties
 // Any lookups will be added below















}

const (
	DocType_Title       = "Document Types"
	DocType_SQLTable    = "docTypeStore"
	DocType_SQLSearchID = "docTypeID"
	DocType_QueryString = "DocTypeID"
	///
	/// Handler Defintions
	///
	DocType_Template     = "DocType"
	DocType_TemplateList = "/DocType/DocType_List"
	DocType_TemplateView = "/DocType/DocType_View"
	DocType_TemplateEdit = "/DocType/DocType_Edit"
	DocType_TemplateNew  = "/DocType/DocType_New"
	///
	/// Handler Monitor Paths
	///
	DocType_Path       = "/API/DocType/"
	DocType_PathList   = "/DocTypeList/"
	DocType_PathView   = "/DocTypeView/"
	DocType_PathEdit   = "/DocTypeEdit/"
	DocType_PathNew    = "/DocTypeNew/"
	DocType_PathSave   = "/DocTypeSave/"
	DocType_PathDelete = "/DocTypeDelete/"
	///
	//DocType_Redirect provides a page to return to aftern an action
	DocType_Redirect = DocType_PathList
	
	///
	///
	/// SQL Field Definitions
	///
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

	/// Definitions End
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

	/// Definitions End
	///
)

//doctype_PageList provides the information for the template for a list of DocTypes
type DocType_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []DocType
	Context	 appContext
}

//doctype_Page provides the information for the template for an individual DocType
type DocType_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 16/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	SYSId_props     FieldProperties
	DocTypeID         string
	DocTypeID_props     FieldProperties
	Code         string
	Code_props     FieldProperties
	Name         string
	Name_props     FieldProperties
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
	Comments         string
	Comments_props     FieldProperties
	// 
	// Dynamically generated 16/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}