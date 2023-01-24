package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/profile.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Profile (profile)
// Endpoint 	        : Profile (ProfileID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 24/01/2023 at 13:18:10
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Profile defines the datamolde for the Profile object
type Profile struct {


SYSId       string
SYSId_props FieldProperties
ProfileID       string
ProfileID_props FieldProperties
Code       string
Code_props FieldProperties
Name       string
Name_props FieldProperties
StartDate       string
StartDate_props FieldProperties
EndDate       string
EndDate_props FieldProperties
DefaultReleases       string
DefaultReleases_props FieldProperties
DefaultReleaseHours       string
DefaultReleaseHours_props FieldProperties
BlendedRate       string
BlendedRate_props FieldProperties
Rounding       string
Rounding_props FieldProperties
HoursPerDay       string
HoursPerDay_props FieldProperties
REQPerc       string
REQPerc_props FieldProperties
ANAPerc       string
ANAPerc_props FieldProperties
DOCPerc       string
DOCPerc_props FieldProperties
PMPerc       string
PMPerc_props FieldProperties
UATPerc       string
UATPerc_props FieldProperties
GTMPerc       string
GTMPerc_props FieldProperties
SupportUplift       string
SupportUplift_props FieldProperties
ContigencyPerc       string
ContigencyPerc_props FieldProperties
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
SYSActivity       string
SYSActivity_props FieldProperties
Notes       string
Notes_props FieldProperties
SYSDbVersion       string
SYSDbVersion_props FieldProperties
Comments       string
Comments_props FieldProperties
TrainingPerc       string
TrainingPerc_props FieldProperties
 // Any lookups will be added below

































}

const (
	Profile_Title       = "Profile"
	Profile_SQLTable    = "profileStore"
	Profile_SQLSearchID = "profileID"
	Profile_QueryString = "ProfileID"
	///
	/// Handler Defintions
	///
	Profile_Template     = "Profile"
	Profile_TemplateList = "/Profile/Profile_List"
	Profile_TemplateView = "/Profile/Profile_View"
	Profile_TemplateEdit = "/Profile/Profile_Edit"
	Profile_TemplateNew  = "/Profile/Profile_New"
	///
	/// Handler Monitor Paths
	///
	Profile_Path       = "/API/Profile/"
	Profile_PathList   = "/ProfileList/"
	Profile_PathView   = "/ProfileView/"
	Profile_PathEdit   = "/ProfileEdit/"
	Profile_PathNew    = "/ProfileNew/"
	Profile_PathSave   = "/ProfileSave/"
	Profile_PathDelete = "/ProfileDelete/"
	///
	//Profile_Redirect provides a page to return to aftern an action
	Profile_Redirect = Profile_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	Profile_SYSId_sql   = "_id" // SYSId is a Int
	Profile_ProfileID_sql   = "profileID" // ProfileID is a String
	Profile_Code_sql   = "code" // Code is a String
	Profile_Name_sql   = "name" // Name is a String
	Profile_StartDate_sql   = "startDate" // StartDate is a String
	Profile_EndDate_sql   = "endDate" // EndDate is a String
	Profile_DefaultReleases_sql   = "DefaultReleases" // DefaultReleases is a String
	Profile_DefaultReleaseHours_sql   = "DefaultReleaseHours" // DefaultReleaseHours is a String
	Profile_BlendedRate_sql   = "BlendedRate" // BlendedRate is a String
	Profile_Rounding_sql   = "Rounding" // Rounding is a String
	Profile_HoursPerDay_sql   = "HoursPerDay" // HoursPerDay is a String
	Profile_REQPerc_sql   = "REQPerc" // REQPerc is a String
	Profile_ANAPerc_sql   = "ANAPerc" // ANAPerc is a String
	Profile_DOCPerc_sql   = "DOCPerc" // DOCPerc is a String
	Profile_PMPerc_sql   = "PMPerc" // PMPerc is a String
	Profile_UATPerc_sql   = "UATPerc" // UATPerc is a String
	Profile_GTMPerc_sql   = "GTMPerc" // GTMPerc is a String
	Profile_SupportUplift_sql   = "SupportUplift" // SupportUplift is a String
	Profile_ContigencyPerc_sql   = "ContigencyPerc" // ContigencyPerc is a String
	Profile_SYSCreated_sql   = "_created" // SYSCreated is a String
	Profile_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	Profile_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	Profile_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	Profile_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
	Profile_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
	Profile_SYSDeleted_sql   = "_deleted" // SYSDeleted is a String
	Profile_SYSDeletedBy_sql   = "_deletedBy" // SYSDeletedBy is a String
	Profile_SYSDeletedHost_sql   = "_deletedHost" // SYSDeletedHost is a String
	Profile_SYSActivity_sql   = "_activity" // SYSActivity is a String
	Profile_Notes_sql   = "notes" // Notes is a String
	Profile_SYSDbVersion_sql   = "_dbVersion" // SYSDbVersion is a String
	Profile_Comments_sql   = "comments" // Comments is a String
	Profile_TrainingPerc_sql   = "TrainingPerc" // TrainingPerc is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	Profile_SYSId_scrn   = "SYSId" // SYSId is a Int
	Profile_ProfileID_scrn   = "ProfileID" // ProfileID is a String
	Profile_Code_scrn   = "Code" // Code is a String
	Profile_Name_scrn   = "Name" // Name is a String
	Profile_StartDate_scrn   = "StartDate" // StartDate is a String
	Profile_EndDate_scrn   = "EndDate" // EndDate is a String
	Profile_DefaultReleases_scrn   = "DefaultReleases" // DefaultReleases is a String
	Profile_DefaultReleaseHours_scrn   = "DefaultReleaseHours" // DefaultReleaseHours is a String
	Profile_BlendedRate_scrn   = "BlendedRate" // BlendedRate is a String
	Profile_Rounding_scrn   = "Rounding" // Rounding is a String
	Profile_HoursPerDay_scrn   = "HoursPerDay" // HoursPerDay is a String
	Profile_REQPerc_scrn   = "REQPerc" // REQPerc is a String
	Profile_ANAPerc_scrn   = "ANAPerc" // ANAPerc is a String
	Profile_DOCPerc_scrn   = "DOCPerc" // DOCPerc is a String
	Profile_PMPerc_scrn   = "PMPerc" // PMPerc is a String
	Profile_UATPerc_scrn   = "UATPerc" // UATPerc is a String
	Profile_GTMPerc_scrn   = "GTMPerc" // GTMPerc is a String
	Profile_SupportUplift_scrn   = "SupportUplift" // SupportUplift is a String
	Profile_ContigencyPerc_scrn   = "ContigencyPerc" // ContigencyPerc is a String
	Profile_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	Profile_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	Profile_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	Profile_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	Profile_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
	Profile_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
	Profile_SYSDeleted_scrn   = "SYSDeleted" // SYSDeleted is a String
	Profile_SYSDeletedBy_scrn   = "SYSDeletedBy" // SYSDeletedBy is a String
	Profile_SYSDeletedHost_scrn   = "SYSDeletedHost" // SYSDeletedHost is a String
	Profile_SYSActivity_scrn   = "SYSActivity" // SYSActivity is a String
	Profile_Notes_scrn   = "Notes" // Notes is a String
	Profile_SYSDbVersion_scrn   = "SYSDbVersion" // SYSDbVersion is a String
	Profile_Comments_scrn   = "Comments" // Comments is a String
	Profile_TrainingPerc_scrn   = "TrainingPerc" // TrainingPerc is a String

	/// Definitions End
	///
)

//profile_PageList provides the information for the template for a list of Profiles
type Profile_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []Profile
	Context	 appContext
	BlockEdit		 bool
	BlockSave 		 bool
	BlockCreate		 bool
	BlockDelete		 bool
	BlockValidate	 bool
	BlockView		 bool
}

//profile_Page provides the information for the template for an individual Profile
type Profile_Page struct {
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
	ProfileID         string
	ProfileID_props     FieldProperties
	Code         string
	Code_props     FieldProperties
	Name         string
	Name_props     FieldProperties
	StartDate         string
	StartDate_props     FieldProperties
	EndDate         string
	EndDate_props     FieldProperties
	DefaultReleases         string
	DefaultReleases_props     FieldProperties
	DefaultReleaseHours         string
	DefaultReleaseHours_props     FieldProperties
	BlendedRate         string
	BlendedRate_props     FieldProperties
	Rounding         string
	Rounding_props     FieldProperties
	HoursPerDay         string
	HoursPerDay_props     FieldProperties
	REQPerc         string
	REQPerc_props     FieldProperties
	ANAPerc         string
	ANAPerc_props     FieldProperties
	DOCPerc         string
	DOCPerc_props     FieldProperties
	PMPerc         string
	PMPerc_props     FieldProperties
	UATPerc         string
	UATPerc_props     FieldProperties
	GTMPerc         string
	GTMPerc_props     FieldProperties
	SupportUplift         string
	SupportUplift_props     FieldProperties
	ContigencyPerc         string
	ContigencyPerc_props     FieldProperties
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
	SYSActivity         string
	SYSActivity_props     FieldProperties
	Notes         string
	Notes_props     FieldProperties
	SYSDbVersion         string
	SYSDbVersion_props     FieldProperties
	Comments         string
	Comments_props     FieldProperties
	TrainingPerc         string
	TrainingPerc_props     FieldProperties
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