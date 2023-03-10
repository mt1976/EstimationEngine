package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/data.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Data (data)
// Endpoint 	        : Data (DataID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 10/03/2023 at 19:54:31
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Data defines the datamodel for the Data object
type Data struct {
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	SYSId       string
	DataID       string
	Class       string
	Field       string
	Value       string
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
	Category       string
	Migrate       string
	//
	// Field Properties
	//
	SYSId_props FieldProperties
	DataID_props FieldProperties
	Class_props FieldProperties
	Field_props FieldProperties
	Value_props FieldProperties
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
	Category_props FieldProperties
	Migrate_props FieldProperties
	//
 	// Any lookups will be added below
	//
	Migrate_lookup []Lookup_Item
	}

const (
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// Page Infrastructure
	Data_Name      = "Data"
	Data_Title       = "System Properties"
	Data_SQLTable    = "dataStore"
	Data_SQLSearchID = "dataID"
	Data_QueryString = "DataID"
	///
	/// Template Path Defintions
	///
	Data_Template     = "Data"
	Data_TemplateList = "/Data/DataList"
	Data_TemplateView = "/Data/DataView"
	Data_TemplateEdit = "/Data/DataEdit"
	Data_TemplateNew  = "/Data/DataNew"
	///
	/// URI Handler Paths
	///
	Data_Path       = "/API/Data/"
	Data_PathList   = "/DataList/"
	Data_PathView   = "/DataView/"
	Data_PathEdit   = "/DataEdit/"
	Data_PathNew    = "/DataNew/"
	Data_PathSave   = "/DataSave/"
	Data_PathDelete = "/DataDelete/"
	// Redirects - On Server Side Error
	Data_PathEditException   = "/DataEditException/"
	Data_PathNewException    = "/DataNewException/"
	//
	//Data_Redirect provides a page to return to aftern an action
	Data_Redirect = Data_PathList
	
	//
	//
	// SQL Field Definitions
	//
	Data_SYSId_sql   = "_id" // SYSId is a Int
	Data_DataID_sql   = "dataID" // DataID is a String
	Data_Class_sql   = "class" // Class is a String
	Data_Field_sql   = "field" // Field is a String
	Data_Value_sql   = "value" // Value is a String
	Data_SYSCreated_sql   = "_created" // SYSCreated is a String
	Data_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	Data_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	Data_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	Data_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
	Data_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
	Data_SYSDeleted_sql   = "_deleted" // SYSDeleted is a String
	Data_SYSDeletedBy_sql   = "_deletedBy" // SYSDeletedBy is a String
	Data_SYSDeletedHost_sql   = "_deletedHost" // SYSDeletedHost is a String
	Data_SYSDbVersion_sql   = "_dbVersion" // SYSDbVersion is a String
	Data_Category_sql   = "category" // Category is a String
	Data_Migrate_sql   = "migrate" // Migrate is a String
///
	/// Application Field Definitions
	///
	Data_SYSId_scrn   = "SYSId" // SYSId is a Int
	Data_DataID_scrn   = "DataID" // DataID is a String
	Data_Class_scrn   = "Class" // Class is a String
	Data_Field_scrn   = "Field" // Field is a String
	Data_Value_scrn   = "Value" // Value is a String
	Data_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	Data_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	Data_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	Data_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	Data_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
	Data_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
	Data_SYSDeleted_scrn   = "SYSDeleted" // SYSDeleted is a String
	Data_SYSDeletedBy_scrn   = "SYSDeletedBy" // SYSDeletedBy is a String
	Data_SYSDeletedHost_scrn   = "SYSDeletedHost" // SYSDeletedHost is a String
	Data_SYSDbVersion_scrn   = "SYSDbVersion" // SYSDbVersion is a String
	Data_Category_scrn   = "Category" // Category is a String
	Data_Migrate_scrn   = "Migrate" // Migrate is a String
///
)

//Data_PageList provides the information for the template for a list of Datas
type Data_PageList struct {
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
	ItemList  		 []Data
}

//Data_Page provides the information for the template for an individual Data
type Data_Page struct {
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
	DataID         string
	Class         string
	Field         string
	Value         string
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
	Category         string
	Migrate         string
	/// Field Properties
	SYSId_props     FieldProperties
	DataID_props     FieldProperties
	Class_props     FieldProperties
	Field_props     FieldProperties
	Value_props     FieldProperties
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
	Category_props     FieldProperties
	Migrate_props     FieldProperties
	/// Lookups
	Migrate_lookup    []Lookup_Item
	
	/// END OF DEFINITIONS
	///
}