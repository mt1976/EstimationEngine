package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/schedule.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Schedule (schedule)
// Endpoint 	        : Schedule (Schedule)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 04/03/2023 at 20:14:14
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Schedule defines the datamodel for the Schedule object
type Schedule struct {
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	SYSId       string
	Id       string
	Name       string
	Description       string
	Schedule       string
	Started       string
	Lastrun       string
	Message       string
	Type       string
	Human       string
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
	//
	// Field Properties
	//
	SYSId_props FieldProperties
	Id_props FieldProperties
	Name_props FieldProperties
	Description_props FieldProperties
	Schedule_props FieldProperties
	Started_props FieldProperties
	Lastrun_props FieldProperties
	Message_props FieldProperties
	Type_props FieldProperties
	Human_props FieldProperties
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
	//
 	// Any lookups will be added below
	//
	}

const (
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// Page Infrastructure
	Schedule_Name      = "Schedule"
	Schedule_Title       = "Scheduler"
	Schedule_SQLTable    = "scheduleStore"
	Schedule_SQLSearchID = "id"
	Schedule_QueryString = "Schedule"
	///
	/// Template Path Defintions
	///
	Schedule_Template     = "Schedule"
	Schedule_TemplateList = "/Schedule/ScheduleList"
	Schedule_TemplateView = "/Schedule/ScheduleView"
	Schedule_TemplateEdit = "/Schedule/ScheduleEdit"
	Schedule_TemplateNew  = "/Schedule/ScheduleNew"
	///
	/// URI Handler Paths
	///
	Schedule_Path       = "/API/Schedule/"
	Schedule_PathList   = "/ScheduleList/"
	Schedule_PathView   = "/ScheduleView/"
	Schedule_PathEdit   = "/ScheduleEdit/"
	Schedule_PathNew    = "/ScheduleNew/"
	Schedule_PathSave   = "/ScheduleSave/"
	Schedule_PathDelete = "/ScheduleDelete/"
	// Redirects - On Server Side Error
	Schedule_PathEditException   = "/ScheduleEditException/"
	Schedule_PathNewException    = "/ScheduleNewException/"
	//
	//Schedule_Redirect provides a page to return to aftern an action
	Schedule_Redirect = Schedule_PathList
	
	//
	//
	// SQL Field Definitions
	//
	Schedule_SYSId_sql   = "_id" // SYSId is a Int
	Schedule_Id_sql   = "id" // Id is a String
	Schedule_Name_sql   = "name" // Name is a String
	Schedule_Description_sql   = "description" // Description is a String
	Schedule_Schedule_sql   = "schedule" // Schedule is a String
	Schedule_Started_sql   = "started" // Started is a String
	Schedule_Lastrun_sql   = "lastrun" // Lastrun is a String
	Schedule_Message_sql   = "message" // Message is a String
	Schedule_Type_sql   = "type" // Type is a String
	Schedule_Human_sql   = "human" // Human is a String
	Schedule_SYSCreated_sql   = "_created" // SYSCreated is a String
	Schedule_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	Schedule_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	Schedule_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	Schedule_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
	Schedule_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
	Schedule_SYSDeleted_sql   = "_deleted" // SYSDeleted is a String
	Schedule_SYSDeletedBy_sql   = "_deletedBy" // SYSDeletedBy is a String
	Schedule_SYSDeletedHost_sql   = "_deletedHost" // SYSDeletedHost is a String
	Schedule_SYSDbVersion_sql   = "_dbVersion" // SYSDbVersion is a String
///
	/// Application Field Definitions
	///
	Schedule_SYSId_scrn   = "SYSId" // SYSId is a Int
	Schedule_Id_scrn   = "Id" // Id is a String
	Schedule_Name_scrn   = "Name" // Name is a String
	Schedule_Description_scrn   = "Description" // Description is a String
	Schedule_Schedule_scrn   = "Schedule" // Schedule is a String
	Schedule_Started_scrn   = "Started" // Started is a String
	Schedule_Lastrun_scrn   = "Lastrun" // Lastrun is a String
	Schedule_Message_scrn   = "Message" // Message is a String
	Schedule_Type_scrn   = "Type" // Type is a String
	Schedule_Human_scrn   = "Human" // Human is a String
	Schedule_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	Schedule_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	Schedule_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	Schedule_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	Schedule_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
	Schedule_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
	Schedule_SYSDeleted_scrn   = "SYSDeleted" // SYSDeleted is a String
	Schedule_SYSDeletedBy_scrn   = "SYSDeletedBy" // SYSDeletedBy is a String
	Schedule_SYSDeletedHost_scrn   = "SYSDeletedHost" // SYSDeletedHost is a String
	Schedule_SYSDbVersion_scrn   = "SYSDbVersion" // SYSDbVersion is a String
///
)

//Schedule_PageList provides the information for the template for a list of Schedules
type Schedule_PageList struct {
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	ItemList  		 []Schedule
}

//Schedule_Page provides the information for the template for an individual Schedule
type Schedule_Page struct {
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	Name         string
	Description         string
	Schedule         string
	Started         string
	Lastrun         string
	Message         string
	Type         string
	Human         string
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
	/// Field Properties
	SYSId_props     FieldProperties
	Id_props     FieldProperties
	Name_props     FieldProperties
	Description_props     FieldProperties
	Schedule_props     FieldProperties
	Started_props     FieldProperties
	Lastrun_props     FieldProperties
	Message_props     FieldProperties
	Type_props     FieldProperties
	Human_props     FieldProperties
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
	/// Lookups
	
	/// END OF DEFINITIONS
	///
}