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
// Date & Time		    : 24/01/2023 at 14:03:00
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Data defines the datamolde for the Data object
type Data struct {


SYSId       string
SYSId_props FieldProperties
DataID       string
DataID_props FieldProperties
Class       string
Class_props FieldProperties
Field       string
Field_props FieldProperties
Value       string
Value_props FieldProperties
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
Category       string
Category_props FieldProperties
 // Any lookups will be added below
















}

const (
	Data_Title       = "System Properties"
	Data_SQLTable    = "dataStore"
	Data_SQLSearchID = "dataID"
	Data_QueryString = "DataID"
	///
	/// Handler Defintions
	///
	Data_Template     = "Data"
	Data_TemplateList = "/Data/Data_List"
	Data_TemplateView = "/Data/Data_View"
	Data_TemplateEdit = "/Data/Data_Edit"
	Data_TemplateNew  = "/Data/Data_New"
	///
	/// Handler Monitor Paths
	///
	Data_Path       = "/API/Data/"
	Data_PathList   = "/DataList/"
	Data_PathView   = "/DataView/"
	Data_PathEdit   = "/DataEdit/"
	Data_PathNew    = "/DataNew/"
	Data_PathSave   = "/DataSave/"
	Data_PathDelete = "/DataDelete/"
	///
	//Data_Redirect provides a page to return to aftern an action
	Data_Redirect = Data_PathList
	
	///
	///
	/// SQL Field Definitions
	///
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

	/// Definitions End
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

	/// Definitions End
	///
)

//data_PageList provides the information for the template for a list of Datas
type Data_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []Data
	Context	 appContext
	BlockEdit		 bool
	BlockSave 		 bool
	BlockCreate		 bool
	BlockDelete		 bool
	BlockValidate	 bool
	BlockView		 bool
}

//data_Page provides the information for the template for an individual Data
type Data_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	SYSId_props     FieldProperties
	DataID         string
	DataID_props     FieldProperties
	Class         string
	Class_props     FieldProperties
	Field         string
	Field_props     FieldProperties
	Value         string
	Value_props     FieldProperties
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
	Category         string
	Category_props     FieldProperties
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