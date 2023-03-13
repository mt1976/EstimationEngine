package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/translation.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Translation (translation)
// Endpoint 	        : Translation (Message)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 13/03/2023 at 15:47:28
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Translation defines the datamodel for the Translation object
type Translation struct {
	// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	SYSId       string
	Id       string
	Class       string
	Message       string
	Translation       string
	SYSCreated       string
	SYSUpdated       string
	SYSCreatedBy       string
	SYSCreatedHost       string
	SYSUpdatedBy       string
	SYSUpdatedHost       string
	SYSDeleted       string
	SYSDeletedBy       string
	SYSDeletedHost       string
	SYSDbVersion       string
	Migrate       string
	//
	// Field Properties
	//
	SYSId_props FieldProperties
	Id_props FieldProperties
	Class_props FieldProperties
	Message_props FieldProperties
	Translation_props FieldProperties
	SYSCreated_props FieldProperties
	SYSUpdated_props FieldProperties
	SYSCreatedBy_props FieldProperties
	SYSCreatedHost_props FieldProperties
	SYSUpdatedBy_props FieldProperties
	SYSUpdatedHost_props FieldProperties
	SYSDeleted_props FieldProperties
	SYSDeletedBy_props FieldProperties
	SYSDeletedHost_props FieldProperties
	SYSDbVersion_props FieldProperties
	Migrate_props FieldProperties
	//
 	// Any lookups will be added below
	//
	Migrate_lookup []Lookup_Item
	}

const (
	// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// Page Infrastructure
	Translation_Name      = "Translation"
	Translation_Title       = "Message Translation"
	Translation_SQLTable    = "translationStore"
	Translation_SQLSearchID = "Id"
	Translation_QueryString = "Message"
	///
	/// Template Path Defintions
	///
	Translation_Template     = "Translation"
	Translation_TemplateList = "/Translation/TranslationList"
	Translation_TemplateView = "/Translation/TranslationView"
	Translation_TemplateEdit = "/Translation/TranslationEdit"
	Translation_TemplateNew  = "/Translation/TranslationNew"
	///
	/// URI Handler Paths
	///
	Translation_Path       = "/API/Translation/"
	Translation_PathList   = "/TranslationList/"
	Translation_PathView   = "/TranslationView/"
	Translation_PathEdit   = "/TranslationEdit/"
	Translation_PathNew    = "/TranslationNew/"
	Translation_PathSave   = "/TranslationSave/"
	Translation_PathDelete = "/TranslationDelete/"
	// Redirects - On Server Side Error
	Translation_PathEditException   = "/TranslationEditException/"
	Translation_PathNewException    = "/TranslationNewException/"
	//
	//Translation_Redirect provides a page to return to aftern an action
	Translation_Redirect = Translation_PathList
	
	//
	//
	// SQL Field Definitions
	//
	Translation_SYSId_sql   = "_id" // SYSId is a Int
	Translation_Id_sql   = "Id" // Id is a String
	Translation_Class_sql   = "Class" // Class is a String
	Translation_Message_sql   = "Message" // Message is a String
	Translation_Translation_sql   = "Translation" // Translation is a String
	Translation_SYSCreated_sql   = "_created" // SYSCreated is a String
	Translation_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	Translation_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	Translation_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	Translation_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
	Translation_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
	Translation_SYSDeleted_sql   = "_deleted" // SYSDeleted is a String
	Translation_SYSDeletedBy_sql   = "_deletedBy" // SYSDeletedBy is a String
	Translation_SYSDeletedHost_sql   = "_deletedHost" // SYSDeletedHost is a String
	Translation_SYSDbVersion_sql   = "_dbVersion" // SYSDbVersion is a String
	Translation_Migrate_sql   = "migrate" // Migrate is a String
///
	/// Application Field Definitions
	///
	Translation_SYSId_scrn   = "SYSId" // SYSId is a Int
	Translation_Id_scrn   = "Id" // Id is a String
	Translation_Class_scrn   = "Class" // Class is a String
	Translation_Message_scrn   = "Message" // Message is a String
	Translation_Translation_scrn   = "Translation" // Translation is a String
	Translation_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	Translation_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	Translation_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	Translation_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	Translation_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
	Translation_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
	Translation_SYSDeleted_scrn   = "SYSDeleted" // SYSDeleted is a String
	Translation_SYSDeletedBy_scrn   = "SYSDeletedBy" // SYSDeletedBy is a String
	Translation_SYSDeletedHost_scrn   = "SYSDeletedHost" // SYSDeletedHost is a String
	Translation_SYSDbVersion_scrn   = "SYSDbVersion" // SYSDbVersion is a String
	Translation_Migrate_scrn   = "Migrate" // Migrate is a String
///
)

//Translation_PageList provides the information for the template for a list of Translations
type Translation_PageList struct {
	// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	ItemList  		 []Translation
}

//Translation_Page provides the information for the template for an individual Translation
type Translation_Page struct {
	// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	Id         string
	Class         string
	Message         string
	Translation         string
	SYSCreated         string
	SYSUpdated         string
	SYSCreatedBy         string
	SYSCreatedHost         string
	SYSUpdatedBy         string
	SYSUpdatedHost         string
	SYSDeleted         string
	SYSDeletedBy         string
	SYSDeletedHost         string
	SYSDbVersion         string
	Migrate         string
	/// Field Properties
	SYSId_props     FieldProperties
	Id_props     FieldProperties
	Class_props     FieldProperties
	Message_props     FieldProperties
	Translation_props     FieldProperties
	SYSCreated_props     FieldProperties
	SYSUpdated_props     FieldProperties
	SYSCreatedBy_props     FieldProperties
	SYSCreatedHost_props     FieldProperties
	SYSUpdatedBy_props     FieldProperties
	SYSUpdatedHost_props     FieldProperties
	SYSDeleted_props     FieldProperties
	SYSDeletedBy_props     FieldProperties
	SYSDeletedHost_props     FieldProperties
	SYSDbVersion_props     FieldProperties
	Migrate_props     FieldProperties
	/// Lookups
	Migrate_lookup    []Lookup_Item
	
	/// END OF DEFINITIONS
	///
}