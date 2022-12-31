package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/feature.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Feature (feature)
// Endpoint 	        : Feature (FeatureID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 30/12/2022 at 12:58:05
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Feature defines the datamolde for the Feature object
type Feature struct {


SYSId       string
SYSId_props FieldProperties
FeatureID       string
FeatureID_props FieldProperties
EstimationSessionID       string
EstimationSessionID_props FieldProperties
ConfidenceID       string
ConfidenceID_props FieldProperties
Name       string
Name_props FieldProperties
DevEstimate       string
DevEstimate_props FieldProperties
DevUplift       string
DevUplift_props FieldProperties
Reqs       string
Reqs_props FieldProperties
AnalystTest       string
AnalystTest_props FieldProperties
Docs       string
Docs_props FieldProperties
Mgt       string
Mgt_props FieldProperties
UatSupport       string
UatSupport_props FieldProperties
Marketing       string
Marketing_props FieldProperties
Contingency       string
Contingency_props FieldProperties
TrackerID       string
TrackerID_props FieldProperties
AdoID       string
AdoID_props FieldProperties
FreshdeskID       string
FreshdeskID_props FieldProperties
ExtRef       string
ExtRef_props FieldProperties
ExtRef2       string
ExtRef2_props FieldProperties
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
Developer       string
Developer_props FieldProperties
Approver       string
Approver_props FieldProperties
Notes       string
Notes_props FieldProperties
OffProfile       string
OffProfile_props FieldProperties
OffProfileJustification       string
OffProfileJustification_props FieldProperties
SYSActivity       string
SYSActivity_props FieldProperties
DfReqs       string
DfReqs_props FieldProperties
DfAnalystTest       string
DfAnalystTest_props FieldProperties
DfDocs       string
DfDocs_props FieldProperties
Dfmgt       string
Dfmgt_props FieldProperties
DfuatSupport       string
DfuatSupport_props FieldProperties
Dfmarketing       string
Dfmarketing_props FieldProperties
Dfcontingency       string
Dfcontingency_props FieldProperties
DfdevUplift       string
DfdevUplift_props FieldProperties
Total       string
Total_props FieldProperties
SYSDbVersion       string
SYSDbVersion_props FieldProperties
Comments       string
Comments_props FieldProperties
Description       string
Description_props FieldProperties
Analyst       string
Analyst_props FieldProperties
ProductManager       string
ProductManager_props FieldProperties
ProjectManager       string
ProjectManager_props FieldProperties
Training       string
Training_props FieldProperties
DfTraining       string
DfTraining_props FieldProperties
DefaultProfile       string
DefaultProfile_props FieldProperties
ActualProfile       string
ActualProfile_props FieldProperties
 // Any lookups will be added below

EstimationSessionID_lookup []Lookup_Item
ConfidenceID_lookup []Lookup_Item
























Developer_lookup []Lookup_Item
Approver_lookup []Lookup_Item
















Analyst_lookup []Lookup_Item
ProductManager_lookup []Lookup_Item
ProjectManager_lookup []Lookup_Item


DefaultProfile_lookup []Lookup_Item
ActualProfile_lookup []Lookup_Item

}

const (
	Feature_Title       = "Feature"
	Feature_SQLTable    = "featureStore"
	Feature_SQLSearchID = "featureID"
	Feature_QueryString = "FeatureID"
	///
	/// Handler Defintions
	///
	Feature_Template     = "Feature"
	Feature_TemplateList = "/Feature/Feature_List"
	Feature_TemplateView = "/Feature/Feature_View"
	Feature_TemplateEdit = "/Feature/Feature_Edit"
	Feature_TemplateNew  = "/Feature/Feature_New"
	///
	/// Handler Monitor Paths
	///
	Feature_Path       = "/API/Feature/"
	Feature_PathList   = "/FeatureList/"
	Feature_PathView   = "/FeatureView/"
	Feature_PathEdit   = "/FeatureEdit/"
	Feature_PathNew    = "/FeatureNew/"
	Feature_PathSave   = "/FeatureSave/"
	Feature_PathDelete = "/FeatureDelete/"
	///
	//Feature_Redirect provides a page to return to aftern an action
	Feature_Redirect = Feature_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	Feature_SYSId_sql   = "_id" // SYSId is a Int
	Feature_FeatureID_sql   = "featureID" // FeatureID is a String
	Feature_EstimationSessionID_sql   = "estimationSessionID" // EstimationSessionID is a String
	Feature_ConfidenceID_sql   = "confidenceID" // ConfidenceID is a String
	Feature_Name_sql   = "name" // Name is a String
	Feature_DevEstimate_sql   = "devEstimate" // DevEstimate is a String
	Feature_DevUplift_sql   = "devUplift" // DevUplift is a String
	Feature_Reqs_sql   = "reqs" // Reqs is a String
	Feature_AnalystTest_sql   = "analystTest" // AnalystTest is a String
	Feature_Docs_sql   = "docs" // Docs is a String
	Feature_Mgt_sql   = "mgt" // Mgt is a String
	Feature_UatSupport_sql   = "uatSupport" // UatSupport is a String
	Feature_Marketing_sql   = "marketing" // Marketing is a String
	Feature_Contingency_sql   = "contingency" // Contingency is a String
	Feature_TrackerID_sql   = "trackerID" // TrackerID is a String
	Feature_AdoID_sql   = "adoID" // AdoID is a String
	Feature_FreshdeskID_sql   = "freshdeskID" // FreshdeskID is a String
	Feature_ExtRef_sql   = "extRef" // ExtRef is a String
	Feature_ExtRef2_sql   = "extRef2" // ExtRef2 is a String
	Feature_SYSCreated_sql   = "_created" // SYSCreated is a String
	Feature_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	Feature_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	Feature_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	Feature_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
	Feature_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
	Feature_SYSDeleted_sql   = "_deleted" // SYSDeleted is a String
	Feature_SYSDeletedBy_sql   = "_deletedBy" // SYSDeletedBy is a String
	Feature_SYSDeletedHost_sql   = "_deletedHost" // SYSDeletedHost is a String
	Feature_Developer_sql   = "developer" // Developer is a String
	Feature_Approver_sql   = "approver" // Approver is a String
	Feature_Notes_sql   = "notes" // Notes is a String
	Feature_OffProfile_sql   = "offProfile" // OffProfile is a String
	Feature_OffProfileJustification_sql   = "offProfileJustification" // OffProfileJustification is a String
	Feature_SYSActivity_sql   = "_activity" // SYSActivity is a String
	Feature_DfReqs_sql   = "dfReqs" // DfReqs is a String
	Feature_DfAnalystTest_sql   = "dfAnalystTest" // DfAnalystTest is a String
	Feature_DfDocs_sql   = "dfDocs" // DfDocs is a String
	Feature_Dfmgt_sql   = "dfmgt" // Dfmgt is a String
	Feature_DfuatSupport_sql   = "dfuatSupport" // DfuatSupport is a String
	Feature_Dfmarketing_sql   = "dfmarketing" // Dfmarketing is a String
	Feature_Dfcontingency_sql   = "dfcontingency" // Dfcontingency is a String
	Feature_DfdevUplift_sql   = "dfdevUplift" // DfdevUplift is a String
	Feature_Total_sql   = "total" // Total is a String
	Feature_SYSDbVersion_sql   = "_dbVersion" // SYSDbVersion is a String
	Feature_Comments_sql   = "comments" // Comments is a String
	Feature_Description_sql   = "description" // Description is a String
	Feature_Analyst_sql   = "analyst" // Analyst is a String
	Feature_ProductManager_sql   = "productManager" // ProductManager is a String
	Feature_ProjectManager_sql   = "projectManager" // ProjectManager is a String
	Feature_Training_sql   = "training" // Training is a String
	Feature_DfTraining_sql   = "dfTraining" // DfTraining is a String
	Feature_DefaultProfile_sql   = "defaultProfile" // DefaultProfile is a String
	Feature_ActualProfile_sql   = "actualProfile" // ActualProfile is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	Feature_SYSId_scrn   = "SYSId" // SYSId is a Int
	Feature_FeatureID_scrn   = "FeatureID" // FeatureID is a String
	Feature_EstimationSessionID_scrn   = "EstimationSessionID" // EstimationSessionID is a String
	Feature_ConfidenceID_scrn   = "ConfidenceID" // ConfidenceID is a String
	Feature_Name_scrn   = "Name" // Name is a String
	Feature_DevEstimate_scrn   = "DevEstimate" // DevEstimate is a String
	Feature_DevUplift_scrn   = "DevUplift" // DevUplift is a String
	Feature_Reqs_scrn   = "Reqs" // Reqs is a String
	Feature_AnalystTest_scrn   = "AnalystTest" // AnalystTest is a String
	Feature_Docs_scrn   = "Docs" // Docs is a String
	Feature_Mgt_scrn   = "Mgt" // Mgt is a String
	Feature_UatSupport_scrn   = "UatSupport" // UatSupport is a String
	Feature_Marketing_scrn   = "Marketing" // Marketing is a String
	Feature_Contingency_scrn   = "Contingency" // Contingency is a String
	Feature_TrackerID_scrn   = "TrackerID" // TrackerID is a String
	Feature_AdoID_scrn   = "AdoID" // AdoID is a String
	Feature_FreshdeskID_scrn   = "FreshdeskID" // FreshdeskID is a String
	Feature_ExtRef_scrn   = "ExtRef" // ExtRef is a String
	Feature_ExtRef2_scrn   = "ExtRef2" // ExtRef2 is a String
	Feature_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	Feature_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	Feature_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	Feature_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	Feature_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
	Feature_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
	Feature_SYSDeleted_scrn   = "SYSDeleted" // SYSDeleted is a String
	Feature_SYSDeletedBy_scrn   = "SYSDeletedBy" // SYSDeletedBy is a String
	Feature_SYSDeletedHost_scrn   = "SYSDeletedHost" // SYSDeletedHost is a String
	Feature_Developer_scrn   = "Developer" // Developer is a String
	Feature_Approver_scrn   = "Approver" // Approver is a String
	Feature_Notes_scrn   = "Notes" // Notes is a String
	Feature_OffProfile_scrn   = "OffProfile" // OffProfile is a String
	Feature_OffProfileJustification_scrn   = "OffProfileJustification" // OffProfileJustification is a String
	Feature_SYSActivity_scrn   = "SYSActivity" // SYSActivity is a String
	Feature_DfReqs_scrn   = "DfReqs" // DfReqs is a String
	Feature_DfAnalystTest_scrn   = "DfAnalystTest" // DfAnalystTest is a String
	Feature_DfDocs_scrn   = "DfDocs" // DfDocs is a String
	Feature_Dfmgt_scrn   = "Dfmgt" // Dfmgt is a String
	Feature_DfuatSupport_scrn   = "DfuatSupport" // DfuatSupport is a String
	Feature_Dfmarketing_scrn   = "Dfmarketing" // Dfmarketing is a String
	Feature_Dfcontingency_scrn   = "Dfcontingency" // Dfcontingency is a String
	Feature_DfdevUplift_scrn   = "DfdevUplift" // DfdevUplift is a String
	Feature_Total_scrn   = "Total" // Total is a String
	Feature_SYSDbVersion_scrn   = "SYSDbVersion" // SYSDbVersion is a String
	Feature_Comments_scrn   = "Comments" // Comments is a String
	Feature_Description_scrn   = "Description" // Description is a String
	Feature_Analyst_scrn   = "Analyst" // Analyst is a String
	Feature_ProductManager_scrn   = "ProductManager" // ProductManager is a String
	Feature_ProjectManager_scrn   = "ProjectManager" // ProjectManager is a String
	Feature_Training_scrn   = "Training" // Training is a String
	Feature_DfTraining_scrn   = "DfTraining" // DfTraining is a String
	Feature_DefaultProfile_scrn   = "DefaultProfile" // DefaultProfile is a String
	Feature_ActualProfile_scrn   = "ActualProfile" // ActualProfile is a String

	/// Definitions End
	///
)

//feature_PageList provides the information for the template for a list of Features
type Feature_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []Feature
	Context	 appContext
}

//feature_Page provides the information for the template for an individual Feature
type Feature_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	SYSId_props     FieldProperties
	FeatureID         string
	FeatureID_props     FieldProperties
	EstimationSessionID         string
	EstimationSessionID_lookup    []Lookup_Item
	EstimationSessionID_props     FieldProperties
	ConfidenceID         string
	ConfidenceID_lookup    []Lookup_Item
	ConfidenceID_props     FieldProperties
	Name         string
	Name_props     FieldProperties
	DevEstimate         string
	DevEstimate_props     FieldProperties
	DevUplift         string
	DevUplift_props     FieldProperties
	Reqs         string
	Reqs_props     FieldProperties
	AnalystTest         string
	AnalystTest_props     FieldProperties
	Docs         string
	Docs_props     FieldProperties
	Mgt         string
	Mgt_props     FieldProperties
	UatSupport         string
	UatSupport_props     FieldProperties
	Marketing         string
	Marketing_props     FieldProperties
	Contingency         string
	Contingency_props     FieldProperties
	TrackerID         string
	TrackerID_props     FieldProperties
	AdoID         string
	AdoID_props     FieldProperties
	FreshdeskID         string
	FreshdeskID_props     FieldProperties
	ExtRef         string
	ExtRef_props     FieldProperties
	ExtRef2         string
	ExtRef2_props     FieldProperties
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
	Developer         string
	Developer_lookup    []Lookup_Item
	Developer_props     FieldProperties
	Approver         string
	Approver_lookup    []Lookup_Item
	Approver_props     FieldProperties
	Notes         string
	Notes_props     FieldProperties
	OffProfile         string
	OffProfile_props     FieldProperties
	OffProfileJustification         string
	OffProfileJustification_props     FieldProperties
	SYSActivity         string
	SYSActivity_props     FieldProperties
	DfReqs         string
	DfReqs_props     FieldProperties
	DfAnalystTest         string
	DfAnalystTest_props     FieldProperties
	DfDocs         string
	DfDocs_props     FieldProperties
	Dfmgt         string
	Dfmgt_props     FieldProperties
	DfuatSupport         string
	DfuatSupport_props     FieldProperties
	Dfmarketing         string
	Dfmarketing_props     FieldProperties
	Dfcontingency         string
	Dfcontingency_props     FieldProperties
	DfdevUplift         string
	DfdevUplift_props     FieldProperties
	Total         string
	Total_props     FieldProperties
	SYSDbVersion         string
	SYSDbVersion_props     FieldProperties
	Comments         string
	Comments_props     FieldProperties
	Description         string
	Description_props     FieldProperties
	Analyst         string
	Analyst_lookup    []Lookup_Item
	Analyst_props     FieldProperties
	ProductManager         string
	ProductManager_lookup    []Lookup_Item
	ProductManager_props     FieldProperties
	ProjectManager         string
	ProjectManager_lookup    []Lookup_Item
	ProjectManager_props     FieldProperties
	Training         string
	Training_props     FieldProperties
	DfTraining         string
	DfTraining_props     FieldProperties
	DefaultProfile         string
	DefaultProfile_lookup    []Lookup_Item
	DefaultProfile_props     FieldProperties
	ActualProfile         string
	ActualProfile_lookup    []Lookup_Item
	ActualProfile_props     FieldProperties
	// 
	// Dynamically generated 30/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}