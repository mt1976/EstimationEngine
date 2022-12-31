package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/featurenew.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : FeatureNew (featurenew)
// Endpoint 	        : FeatureNew (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 30/12/2022 at 12:18:21
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//FeatureNew defines the datamolde for the FeatureNew object
type FeatureNew struct {


ID       string
ID_props FieldProperties
EstimationSession       string
EstimationSession_props FieldProperties
Name       string
Name_props FieldProperties
DevEstimate       string
DevEstimate_props FieldProperties
Confidence       string
Confidence_props FieldProperties
Developer       string
Developer_props FieldProperties
 // Any lookups will be added below
EstimationSession_lookup []Lookup_Item


Confidence_lookup []Lookup_Item
Developer_lookup []Lookup_Item

}

const (
	FeatureNew_Title       = "Create a New Feature"
	FeatureNew_SQLTable    = "n/a"
	FeatureNew_SQLSearchID = "n/a"
	FeatureNew_QueryString = "ID"
	///
	/// Handler Defintions
	///
	FeatureNew_Template     = "FeatureNew"
	FeatureNew_TemplateList = "/FeatureNew/FeatureNew_List"
	FeatureNew_TemplateView = "/FeatureNew/FeatureNew_View"
	FeatureNew_TemplateEdit = "/FeatureNew/FeatureNew_Edit"
	FeatureNew_TemplateNew  = "/FeatureNew/FeatureNew_New"
	///
	/// Handler Monitor Paths
	///
	FeatureNew_Path       = "/API/FeatureNew/"
	FeatureNew_PathList   = "/FeatureNewList/"
	FeatureNew_PathView   = "/FeatureNewView/"
	FeatureNew_PathEdit   = "/FeatureNewEdit/"
	FeatureNew_PathNew    = "/FeatureNewNew/"
	FeatureNew_PathSave   = "/FeatureNewSave/"
	FeatureNew_PathDelete = "/FeatureNewDelete/"
	///
	//FeatureNew_Redirect provides a page to return to aftern an action
	FeatureNew_Redirect = FeatureNew_PathView
	///
	///
	/// SQL Field Definitions
	///
	FeatureNew_ID_sql   = "ID" // ID is a String
	FeatureNew_EstimationSession_sql   = "EstimationSession" // EstimationSession is a String
	FeatureNew_Name_sql   = "Name" // Name is a String
	FeatureNew_DevEstimate_sql   = "DevEstimate" // DevEstimate is a String
	FeatureNew_Confidence_sql   = "Confidence" // Confidence is a String
	FeatureNew_Developer_sql   = "Developer" // Developer is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	FeatureNew_ID_scrn   = "ID" // ID is a String
	FeatureNew_EstimationSession_scrn   = "EstimationSession" // EstimationSession is a String
	FeatureNew_Name_scrn   = "Name" // Name is a String
	FeatureNew_DevEstimate_scrn   = "DevEstimate" // DevEstimate is a String
	FeatureNew_Confidence_scrn   = "Confidence" // Confidence is a String
	FeatureNew_Developer_scrn   = "Developer" // Developer is a String

	/// Definitions End
	///
)

//featurenew_PageList provides the information for the template for a list of FeatureNews
type FeatureNew_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []FeatureNew
	Context	 appContext
}

//featurenew_Page provides the information for the template for an individual FeatureNew
type FeatureNew_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	ID         string
	ID_props     FieldProperties
	EstimationSession         string
	EstimationSession_lookup    []Lookup_Item
	EstimationSession_props     FieldProperties
	Name         string
	Name_props     FieldProperties
	DevEstimate         string
	DevEstimate_props     FieldProperties
	Confidence         string
	Confidence_lookup    []Lookup_Item
	Confidence_props     FieldProperties
	Developer         string
	Developer_lookup    []Lookup_Item
	Developer_props     FieldProperties
	// 
	// Dynamically generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}