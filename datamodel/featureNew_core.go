package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/featurenew.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : FeatureNew (featurenew)
// Endpoint 	        : FeatureNew (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 08/03/2023 at 18:42:24
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//FeatureNew defines the datamodel for the FeatureNew object
type FeatureNew struct {
	// Dynamically generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	ID       string
	EstimationSessionID       string
	Name       string
	DeveloperEstimate       string
	DeveloperResource       string
	ConfidenceCODE       string
	AnalystEstimate       string
	AnalystResource       string
	EstimateEffort       string
	DevOpsID       string
	FreshDeskID       string
	RSCID       string
	OtherID       string
	OtherID2       string
	ProfileDefault       string
	ProfileSelected       string
	//
	// Field Properties
	//
	ID_props FieldProperties
	EstimationSessionID_props FieldProperties
	Name_props FieldProperties
	DeveloperEstimate_props FieldProperties
	DeveloperResource_props FieldProperties
	ConfidenceCODE_props FieldProperties
	AnalystEstimate_props FieldProperties
	AnalystResource_props FieldProperties
	EstimateEffort_props FieldProperties
	DevOpsID_props FieldProperties
	FreshDeskID_props FieldProperties
	RSCID_props FieldProperties
	OtherID_props FieldProperties
	OtherID2_props FieldProperties
	ProfileDefault_props FieldProperties
	ProfileSelected_props FieldProperties
	//
 	// Any lookups will be added below
	//
	EstimationSessionID_lookup []Lookup_Item
	DeveloperResource_lookup []Lookup_Item
	ConfidenceCODE_lookup []Lookup_Item
	AnalystResource_lookup []Lookup_Item
	ProfileDefault_lookup []Lookup_Item
	ProfileSelected_lookup []Lookup_Item
	}

const (
	// Dynamically generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// Page Infrastructure
	FeatureNew_Name      = "FeatureNew"
	FeatureNew_Title       = "Create a New Feature"
	FeatureNew_SQLTable    = "n/a"
	FeatureNew_SQLSearchID = "n/a"
	FeatureNew_QueryString = "ID"
	///
	/// Template Path Defintions
	///
	FeatureNew_Template     = "FeatureNew"
	FeatureNew_TemplateList = "/FeatureNew/FeatureNewList"
	FeatureNew_TemplateView = "/FeatureNew/FeatureNewView"
	FeatureNew_TemplateEdit = "/FeatureNew/FeatureNewEdit"
	FeatureNew_TemplateNew  = "/FeatureNew/FeatureNewNew"
	///
	/// URI Handler Paths
	///
	FeatureNew_Path       = "/API/FeatureNew/"
	FeatureNew_PathList   = "/FeatureNewList/"
	FeatureNew_PathView   = "/FeatureNewView/"
	FeatureNew_PathEdit   = "/FeatureNewEdit/"
	FeatureNew_PathNew    = "/FeatureNewNew/"
	FeatureNew_PathSave   = "/FeatureNewSave/"
	FeatureNew_PathDelete = "/FeatureNewDelete/"
	// Redirects - On Server Side Error
	FeatureNew_PathEditException   = "/FeatureNewEditException/"
	FeatureNew_PathNewException    = "/FeatureNewNewException/"
	//
	//FeatureNew_Redirect provides a page to return to aftern an action
	FeatureNew_Redirect = FeatureNew_PathView
	//
	//
	// SQL Field Definitions
	//
	FeatureNew_ID_sql   = "ID" // ID is a String
	FeatureNew_EstimationSessionID_sql   = "EstimationSessionID" // EstimationSessionID is a String
	FeatureNew_Name_sql   = "Name" // Name is a String
	FeatureNew_DeveloperEstimate_sql   = "DeveloperEstimate" // DeveloperEstimate is a String
	FeatureNew_DeveloperResource_sql   = "DeveloperResource" // DeveloperResource is a String
	FeatureNew_ConfidenceCODE_sql   = "ConfidenceCODE" // ConfidenceCODE is a String
	FeatureNew_AnalystEstimate_sql   = "AnalystEstimate" // AnalystEstimate is a String
	FeatureNew_AnalystResource_sql   = "AnalystResource" // AnalystResource is a String
	FeatureNew_EstimateEffort_sql   = "EstimateEffort" // EstimateEffort is a String
	FeatureNew_DevOpsID_sql   = "DevOpsID" // DevOpsID is a String
	FeatureNew_FreshDeskID_sql   = "FreshDeskID" // FreshDeskID is a String
	FeatureNew_RSCID_sql   = "RSCID" // RSCID is a String
	FeatureNew_OtherID_sql   = "OtherID" // OtherID is a String
	FeatureNew_OtherID2_sql   = "OtherID2" // OtherID2 is a String
	FeatureNew_ProfileDefault_sql   = "ProfileDefault" // ProfileDefault is a String
	FeatureNew_ProfileSelected_sql   = "ProfileSelected" // ProfileSelected is a String
///
	/// Application Field Definitions
	///
	FeatureNew_ID_scrn   = "ID" // ID is a String
	FeatureNew_EstimationSessionID_scrn   = "EstimationSessionID" // EstimationSessionID is a String
	FeatureNew_Name_scrn   = "Name" // Name is a String
	FeatureNew_DeveloperEstimate_scrn   = "DeveloperEstimate" // DeveloperEstimate is a String
	FeatureNew_DeveloperResource_scrn   = "DeveloperResource" // DeveloperResource is a String
	FeatureNew_ConfidenceCODE_scrn   = "ConfidenceCODE" // ConfidenceCODE is a String
	FeatureNew_AnalystEstimate_scrn   = "AnalystEstimate" // AnalystEstimate is a String
	FeatureNew_AnalystResource_scrn   = "AnalystResource" // AnalystResource is a String
	FeatureNew_EstimateEffort_scrn   = "EstimateEffort" // EstimateEffort is a String
	FeatureNew_DevOpsID_scrn   = "DevOpsID" // DevOpsID is a String
	FeatureNew_FreshDeskID_scrn   = "FreshDeskID" // FreshDeskID is a String
	FeatureNew_RSCID_scrn   = "RSCID" // RSCID is a String
	FeatureNew_OtherID_scrn   = "OtherID" // OtherID is a String
	FeatureNew_OtherID2_scrn   = "OtherID2" // OtherID2 is a String
	FeatureNew_ProfileDefault_scrn   = "ProfileDefault" // ProfileDefault is a String
	FeatureNew_ProfileSelected_scrn   = "ProfileSelected" // ProfileSelected is a String
///
)

//FeatureNew_PageList provides the information for the template for a list of FeatureNews
type FeatureNew_PageList struct {
	// Dynamically generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	ItemList  		 []FeatureNew
}

//FeatureNew_Page provides the information for the template for an individual FeatureNew
type FeatureNew_Page struct {
	// Dynamically generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	ID         string
	EstimationSessionID         string
	Name         string
	DeveloperEstimate         string
	DeveloperResource         string
	ConfidenceCODE         string
	AnalystEstimate         string
	AnalystResource         string
	EstimateEffort         string
	DevOpsID         string
	FreshDeskID         string
	RSCID         string
	OtherID         string
	OtherID2         string
	ProfileDefault         string
	ProfileSelected         string
	/// Field Properties
	ID_props     FieldProperties
	EstimationSessionID_props     FieldProperties
	Name_props     FieldProperties
	DeveloperEstimate_props     FieldProperties
	DeveloperResource_props     FieldProperties
	ConfidenceCODE_props     FieldProperties
	AnalystEstimate_props     FieldProperties
	AnalystResource_props     FieldProperties
	EstimateEffort_props     FieldProperties
	DevOpsID_props     FieldProperties
	FreshDeskID_props     FieldProperties
	RSCID_props     FieldProperties
	OtherID_props     FieldProperties
	OtherID2_props     FieldProperties
	ProfileDefault_props     FieldProperties
	ProfileSelected_props     FieldProperties
	/// Lookups
	EstimationSessionID_lookup    []Lookup_Item
	DeveloperResource_lookup    []Lookup_Item
	ConfidenceCODE_lookup    []Lookup_Item
	AnalystResource_lookup    []Lookup_Item
	ProfileDefault_lookup    []Lookup_Item
	ProfileSelected_lookup    []Lookup_Item
	
	/// END OF DEFINITIONS
	///
}