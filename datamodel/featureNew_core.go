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
// Date & Time		    : 07/01/2023 at 23:01:29
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
Comments       string
Comments_props FieldProperties
Description       string
Description_props FieldProperties
DevOpsID       string
DevOpsID_props FieldProperties
FreshDeskID       string
FreshDeskID_props FieldProperties
RSCID       string
RSCID_props FieldProperties
OtherID       string
OtherID_props FieldProperties
OtherID2       string
OtherID2_props FieldProperties
Analyst       string
Analyst_props FieldProperties
ProductManager       string
ProductManager_props FieldProperties
ProjectManager       string
ProjectManager_props FieldProperties
DefaultProfile       string
DefaultProfile_props FieldProperties
ActualProfile       string
ActualProfile_props FieldProperties
 // Any lookups will be added below
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
	// Dynamically generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	Comments         string
	Comments_props     FieldProperties
	Description         string
	Description_props     FieldProperties
	DevOpsID         string
	DevOpsID_props     FieldProperties
	FreshDeskID         string
	FreshDeskID_props     FieldProperties
	RSCID         string
	RSCID_props     FieldProperties
	OtherID         string
	OtherID_props     FieldProperties
	OtherID2         string
	OtherID2_props     FieldProperties
	Analyst         string
	Analyst_lookup    []Lookup_Item
	Analyst_props     FieldProperties
	ProductManager         string
	ProductManager_lookup    []Lookup_Item
	ProductManager_props     FieldProperties
	ProjectManager         string
	ProjectManager_lookup    []Lookup_Item
	ProjectManager_props     FieldProperties
	DefaultProfile         string
	DefaultProfile_lookup    []Lookup_Item
	DefaultProfile_props     FieldProperties
	ActualProfile         string
	ActualProfile_lookup    []Lookup_Item
	ActualProfile_props     FieldProperties
	// 
	// Dynamically generated 07/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}