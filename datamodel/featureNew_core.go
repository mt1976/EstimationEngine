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
// Date & Time		    : 15/02/2023 at 10:44:44
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//FeatureNew defines the datamodel for the FeatureNew object
type FeatureNew struct {
	// Dynamically generated 15/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	ID       string
	EstimationSession       string
	Name       string
	DevEstimate       string
	Confidence       string
	Developer       string
	Comments       string
	Description       string
	DevOpsID       string
	FreshDeskID       string
	RSCID       string
	OtherID       string
	OtherID2       string
	Analyst       string
	ProductManager       string
	ProjectManager       string
	DefaultProfile       string
	ActualProfile       string
	//
	// Field Properties
	//
	ID_props FieldProperties
	EstimationSession_props FieldProperties
	Name_props FieldProperties
	DevEstimate_props FieldProperties
	Confidence_props FieldProperties
	Developer_props FieldProperties
	Comments_props FieldProperties
	Description_props FieldProperties
	DevOpsID_props FieldProperties
	FreshDeskID_props FieldProperties
	RSCID_props FieldProperties
	OtherID_props FieldProperties
	OtherID2_props FieldProperties
	Analyst_props FieldProperties
	ProductManager_props FieldProperties
	ProjectManager_props FieldProperties
	DefaultProfile_props FieldProperties
	ActualProfile_props FieldProperties
	//
 	// Any lookups will be added below
	//
	EstimationSession_lookup []Lookup_Item
	Confidence_lookup []Lookup_Item
	Developer_lookup []Lookup_Item
	Analyst_lookup []Lookup_Item
	ProductManager_lookup []Lookup_Item
	ProjectManager_lookup []Lookup_Item
	DefaultProfile_lookup []Lookup_Item
	ActualProfile_lookup []Lookup_Item
	}

const (
	// Dynamically generated 15/02/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	FeatureNew_EstimationSession_sql   = "EstimationSession" // EstimationSession is a String
	FeatureNew_Name_sql   = "Name" // Name is a String
	FeatureNew_DevEstimate_sql   = "DevEstimate" // DevEstimate is a String
	FeatureNew_Confidence_sql   = "Confidence" // Confidence is a String
	FeatureNew_Developer_sql   = "Developer" // Developer is a String
	FeatureNew_Comments_sql   = "Comments" // Comments is a String
	FeatureNew_Description_sql   = "Description" // Description is a String
	FeatureNew_DevOpsID_sql   = "DevOpsID" // DevOpsID is a String
	FeatureNew_FreshDeskID_sql   = "FreshDeskID" // FreshDeskID is a String
	FeatureNew_RSCID_sql   = "RSCID" // RSCID is a String
	FeatureNew_OtherID_sql   = "OtherID" // OtherID is a String
	FeatureNew_OtherID2_sql   = "OtherID2" // OtherID2 is a String
	FeatureNew_Analyst_sql   = "Analyst" // Analyst is a String
	FeatureNew_ProductManager_sql   = "ProductManager" // ProductManager is a String
	FeatureNew_ProjectManager_sql   = "ProjectManager" // ProjectManager is a String
	FeatureNew_DefaultProfile_sql   = "DefaultProfile" // DefaultProfile is a String
	FeatureNew_ActualProfile_sql   = "ActualProfile" // ActualProfile is a String
///
	/// Application Field Definitions
	///
	FeatureNew_ID_scrn   = "ID" // ID is a String
	FeatureNew_EstimationSession_scrn   = "EstimationSession" // EstimationSession is a String
	FeatureNew_Name_scrn   = "Name" // Name is a String
	FeatureNew_DevEstimate_scrn   = "DevEstimate" // DevEstimate is a String
	FeatureNew_Confidence_scrn   = "Confidence" // Confidence is a String
	FeatureNew_Developer_scrn   = "Developer" // Developer is a String
	FeatureNew_Comments_scrn   = "Comments" // Comments is a String
	FeatureNew_Description_scrn   = "Description" // Description is a String
	FeatureNew_DevOpsID_scrn   = "DevOpsID" // DevOpsID is a String
	FeatureNew_FreshDeskID_scrn   = "FreshDeskID" // FreshDeskID is a String
	FeatureNew_RSCID_scrn   = "RSCID" // RSCID is a String
	FeatureNew_OtherID_scrn   = "OtherID" // OtherID is a String
	FeatureNew_OtherID2_scrn   = "OtherID2" // OtherID2 is a String
	FeatureNew_Analyst_scrn   = "Analyst" // Analyst is a String
	FeatureNew_ProductManager_scrn   = "ProductManager" // ProductManager is a String
	FeatureNew_ProjectManager_scrn   = "ProjectManager" // ProjectManager is a String
	FeatureNew_DefaultProfile_scrn   = "DefaultProfile" // DefaultProfile is a String
	FeatureNew_ActualProfile_scrn   = "ActualProfile" // ActualProfile is a String
///
)

//FeatureNew_PageList provides the information for the template for a list of FeatureNews
type FeatureNew_PageList struct {
	// Dynamically generated 15/02/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	// Dynamically generated 15/02/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	EstimationSession         string
	Name         string
	DevEstimate         string
	Confidence         string
	Developer         string
	Comments         string
	Description         string
	DevOpsID         string
	FreshDeskID         string
	RSCID         string
	OtherID         string
	OtherID2         string
	Analyst         string
	ProductManager         string
	ProjectManager         string
	DefaultProfile         string
	ActualProfile         string
	/// Field Properties
	ID_props     FieldProperties
	EstimationSession_props     FieldProperties
	Name_props     FieldProperties
	DevEstimate_props     FieldProperties
	Confidence_props     FieldProperties
	Developer_props     FieldProperties
	Comments_props     FieldProperties
	Description_props     FieldProperties
	DevOpsID_props     FieldProperties
	FreshDeskID_props     FieldProperties
	RSCID_props     FieldProperties
	OtherID_props     FieldProperties
	OtherID2_props     FieldProperties
	Analyst_props     FieldProperties
	ProductManager_props     FieldProperties
	ProjectManager_props     FieldProperties
	DefaultProfile_props     FieldProperties
	ActualProfile_props     FieldProperties
	/// Lookups
	EstimationSession_lookup    []Lookup_Item
	Confidence_lookup    []Lookup_Item
	Developer_lookup    []Lookup_Item
	Analyst_lookup    []Lookup_Item
	ProductManager_lookup    []Lookup_Item
	ProjectManager_lookup    []Lookup_Item
	DefaultProfile_lookup    []Lookup_Item
	ActualProfile_lookup    []Lookup_Item
	
	/// END OF DEFINITIONS
	///
}