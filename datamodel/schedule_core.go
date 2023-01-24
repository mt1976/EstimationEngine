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
// Date & Time		    : 24/01/2023 at 13:18:11
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Schedule defines the datamolde for the Schedule object
type Schedule struct {


SYSId       string
SYSId_props FieldProperties
Id       string
Id_props FieldProperties
Name       string
Name_props FieldProperties
Description       string
Description_props FieldProperties
Schedule       string
Schedule_props FieldProperties
Started       string
Started_props FieldProperties
Lastrun       string
Lastrun_props FieldProperties
Message       string
Message_props FieldProperties
Type       string
Type_props FieldProperties
Human       string
Human_props FieldProperties
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
 // Any lookups will be added below




















}

const (
	Schedule_Title       = "Scheduler"
	Schedule_SQLTable    = "scheduleStore"
	Schedule_SQLSearchID = "id"
	Schedule_QueryString = "Schedule"
	///
	/// Handler Defintions
	///
	Schedule_Template     = "Schedule"
	Schedule_TemplateList = "/Schedule/Schedule_List"
	Schedule_TemplateView = "/Schedule/Schedule_View"
	Schedule_TemplateEdit = "/Schedule/Schedule_Edit"
	Schedule_TemplateNew  = "/Schedule/Schedule_New"
	///
	/// Handler Monitor Paths
	///
	Schedule_Path       = "/API/Schedule/"
	Schedule_PathList   = "/ScheduleList/"
	Schedule_PathView   = "/ScheduleView/"
	Schedule_PathEdit   = "/ScheduleEdit/"
	Schedule_PathNew    = "/ScheduleNew/"
	Schedule_PathSave   = "/ScheduleSave/"
	Schedule_PathDelete = "/ScheduleDelete/"
	///
	//Schedule_Redirect provides a page to return to aftern an action
	Schedule_Redirect = Schedule_PathList
	
	///
	///
	/// SQL Field Definitions
	///
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

	/// Definitions End
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

	/// Definitions End
	///
)

//schedule_PageList provides the information for the template for a list of Schedules
type Schedule_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []Schedule
	Context	 appContext
	BlockEdit		 bool
	BlockSave 		 bool
	BlockCreate		 bool
	BlockDelete		 bool
	BlockValidate	 bool
	BlockView		 bool
}

//schedule_Page provides the information for the template for an individual Schedule
type Schedule_Page struct {
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
	Id         string
	Id_props     FieldProperties
	Name         string
	Name_props     FieldProperties
	Description         string
	Description_props     FieldProperties
	Schedule         string
	Schedule_props     FieldProperties
	Started         string
	Started_props     FieldProperties
	Lastrun         string
	Lastrun_props     FieldProperties
	Message         string
	Message_props     FieldProperties
	Type         string
	Type_props     FieldProperties
	Human         string
	Human_props     FieldProperties
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