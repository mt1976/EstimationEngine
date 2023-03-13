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
// Date & Time		    : 13/03/2023 at 14:22:29
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Profile defines the datamodel for the Profile object
type Profile struct {
	// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	SYSId       string
	ProfileID       string
	Code       string
	Name       string
	StartDate       string
	EndDate       string
	DefaultReleases       string
	DefaultReleaseHours       string
	BlendedRate       string
	Rounding       string
	HoursPerDay       string
	REQPerc       string
	ANAPerc       string
	DOCPerc       string
	PMPerc       string
	UATPerc       string
	GTMPerc       string
	SupportUplift       string
	ContigencyPerc       string
	SYSCreated       string
	SYSCreatedBy       string
	SYSCreatedHost       string
	SYSUpdated       string
	SYSUpdatedBy       string
	SYSUpdatedHost       string
	SYSDeleted       string
	SYSDeletedBy       string
	SYSDeletedHost       string
	SYSActivity       string
	Notes       string
	SYSDbVersion       string
	Comments       string
	TrainingPerc       string
	//
	// Field Properties
	//
	SYSId_props FieldProperties
	ProfileID_props FieldProperties
	Code_props FieldProperties
	Name_props FieldProperties
	StartDate_props FieldProperties
	EndDate_props FieldProperties
	DefaultReleases_props FieldProperties
	DefaultReleaseHours_props FieldProperties
	BlendedRate_props FieldProperties
	Rounding_props FieldProperties
	HoursPerDay_props FieldProperties
	REQPerc_props FieldProperties
	ANAPerc_props FieldProperties
	DOCPerc_props FieldProperties
	PMPerc_props FieldProperties
	UATPerc_props FieldProperties
	GTMPerc_props FieldProperties
	SupportUplift_props FieldProperties
	ContigencyPerc_props FieldProperties
	SYSCreated_props FieldProperties
	SYSCreatedBy_props FieldProperties
	SYSCreatedHost_props FieldProperties
	SYSUpdated_props FieldProperties
	SYSUpdatedBy_props FieldProperties
	SYSUpdatedHost_props FieldProperties
	SYSDeleted_props FieldProperties
	SYSDeletedBy_props FieldProperties
	SYSDeletedHost_props FieldProperties
	SYSActivity_props FieldProperties
	Notes_props FieldProperties
	SYSDbVersion_props FieldProperties
	Comments_props FieldProperties
	TrainingPerc_props FieldProperties
	//
 	// Any lookups will be added below
	//
	}

const (
	// Dynamically generated 13/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// Page Infrastructure
	Profile_Name      = "Profile"
	Profile_Title       = "Profile"
	Profile_SQLTable    = "profileStore"
	Profile_SQLSearchID = "profileID"
	Profile_QueryString = "ProfileID"
	///
	/// Template Path Defintions
	///
	Profile_Template     = "Profile"
	Profile_TemplateList = "/Profile/ProfileList"
	Profile_TemplateView = "/Profile/ProfileView"
	Profile_TemplateEdit = "/Profile/ProfileEdit"
	Profile_TemplateNew  = "/Profile/ProfileNew"
	///
	/// URI Handler Paths
	///
	Profile_Path       = "/API/Profile/"
	Profile_PathList   = "/ProfileList/"
	Profile_PathView   = "/ProfileView/"
	Profile_PathEdit   = "/ProfileEdit/"
	Profile_PathNew    = "/ProfileNew/"
	Profile_PathSave   = "/ProfileSave/"
	Profile_PathDelete = "/ProfileDelete/"
	// Redirects - On Server Side Error
	Profile_PathEditException   = "/ProfileEditException/"
	Profile_PathNewException    = "/ProfileNewException/"
	//
	//Profile_Redirect provides a page to return to aftern an action
	Profile_Redirect = Profile_PathList
	
	//
	//
	// SQL Field Definitions
	//
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
///
)

//Profile_PageList provides the information for the template for a list of Profiles
type Profile_PageList struct {
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
	ItemList  		 []Profile
}

//Profile_Page provides the information for the template for an individual Profile
type Profile_Page struct {
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
	ProfileID         string
	Code         string
	Name         string
	StartDate         string
	EndDate         string
	DefaultReleases         string
	DefaultReleaseHours         string
	BlendedRate         string
	Rounding         string
	HoursPerDay         string
	REQPerc         string
	ANAPerc         string
	DOCPerc         string
	PMPerc         string
	UATPerc         string
	GTMPerc         string
	SupportUplift         string
	ContigencyPerc         string
	SYSCreated         string
	SYSCreatedBy         string
	SYSCreatedHost         string
	SYSUpdated         string
	SYSUpdatedBy         string
	SYSUpdatedHost         string
	SYSDeleted         string
	SYSDeletedBy         string
	SYSDeletedHost         string
	SYSActivity         string
	Notes         string
	SYSDbVersion         string
	Comments         string
	TrainingPerc         string
	/// Field Properties
	SYSId_props     FieldProperties
	ProfileID_props     FieldProperties
	Code_props     FieldProperties
	Name_props     FieldProperties
	StartDate_props     FieldProperties
	EndDate_props     FieldProperties
	DefaultReleases_props     FieldProperties
	DefaultReleaseHours_props     FieldProperties
	BlendedRate_props     FieldProperties
	Rounding_props     FieldProperties
	HoursPerDay_props     FieldProperties
	REQPerc_props     FieldProperties
	ANAPerc_props     FieldProperties
	DOCPerc_props     FieldProperties
	PMPerc_props     FieldProperties
	UATPerc_props     FieldProperties
	GTMPerc_props     FieldProperties
	SupportUplift_props     FieldProperties
	ContigencyPerc_props     FieldProperties
	SYSCreated_props     FieldProperties
	SYSCreatedBy_props     FieldProperties
	SYSCreatedHost_props     FieldProperties
	SYSUpdated_props     FieldProperties
	SYSUpdatedBy_props     FieldProperties
	SYSUpdatedHost_props     FieldProperties
	SYSDeleted_props     FieldProperties
	SYSDeletedBy_props     FieldProperties
	SYSDeletedHost_props     FieldProperties
	SYSActivity_props     FieldProperties
	Notes_props     FieldProperties
	SYSDbVersion_props     FieldProperties
	Comments_props     FieldProperties
	TrainingPerc_props     FieldProperties
	/// Lookups
	
	/// END OF DEFINITIONS
	///
}