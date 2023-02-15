package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/feature.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Feature (feature)
// Endpoint 	        : Feature (FeatureID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/02/2023 at 10:44:43
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Feature defines the datamodel for the Feature object
type Feature struct {
	// Dynamically generated 15/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	SYSId       string
	FeatureID       string
	EstimationSessionID       string
	ConfidenceID       string
	Name       string
	DevEstimate       string
	DevUplift       string
	Reqs       string
	AnalystTest       string
	Docs       string
	Mgt       string
	UatSupport       string
	Marketing       string
	Contingency       string
	TrackerID       string
	AdoID       string
	FreshdeskID       string
	ExtRef       string
	ExtRef2       string
	SYSCreated       string
	SYSCreatedBy       string
	SYSCreatedHost       string
	SYSUpdated       string
	SYSUpdatedBy       string
	SYSUpdatedHost       string
	SYSDeleted       string
	SYSDeletedBy       string
	SYSDeletedHost       string
	Developer       string
	Approver       string
	Notes       string
	OffProfile       string
	OffProfileJustification       string
	SYSActivity       string
	DfReqs       string
	DfAnalystTest       string
	DfDocs       string
	Dfmgt       string
	DfuatSupport       string
	Dfmarketing       string
	Dfcontingency       string
	DfdevUplift       string
	Total       string
	SYSDbVersion       string
	Comments       string
	Description       string
	Analyst       string
	ProductManager       string
	ProjectManager       string
	Training       string
	DfTraining       string
	DefaultProfile       string
	ActualProfile       string
	//
	// Field Properties
	//
	SYSId_props FieldProperties
	FeatureID_props FieldProperties
	EstimationSessionID_props FieldProperties
	ConfidenceID_props FieldProperties
	Name_props FieldProperties
	DevEstimate_props FieldProperties
	DevUplift_props FieldProperties
	Reqs_props FieldProperties
	AnalystTest_props FieldProperties
	Docs_props FieldProperties
	Mgt_props FieldProperties
	UatSupport_props FieldProperties
	Marketing_props FieldProperties
	Contingency_props FieldProperties
	TrackerID_props FieldProperties
	AdoID_props FieldProperties
	FreshdeskID_props FieldProperties
	ExtRef_props FieldProperties
	ExtRef2_props FieldProperties
	SYSCreated_props FieldProperties
	SYSCreatedBy_props FieldProperties
	SYSCreatedHost_props FieldProperties
	SYSUpdated_props FieldProperties
	SYSUpdatedBy_props FieldProperties
	SYSUpdatedHost_props FieldProperties
	SYSDeleted_props FieldProperties
	SYSDeletedBy_props FieldProperties
	SYSDeletedHost_props FieldProperties
	Developer_props FieldProperties
	Approver_props FieldProperties
	Notes_props FieldProperties
	OffProfile_props FieldProperties
	OffProfileJustification_props FieldProperties
	SYSActivity_props FieldProperties
	DfReqs_props FieldProperties
	DfAnalystTest_props FieldProperties
	DfDocs_props FieldProperties
	Dfmgt_props FieldProperties
	DfuatSupport_props FieldProperties
	Dfmarketing_props FieldProperties
	Dfcontingency_props FieldProperties
	DfdevUplift_props FieldProperties
	Total_props FieldProperties
	SYSDbVersion_props FieldProperties
	Comments_props FieldProperties
	Description_props FieldProperties
	Analyst_props FieldProperties
	ProductManager_props FieldProperties
	ProjectManager_props FieldProperties
	Training_props FieldProperties
	DfTraining_props FieldProperties
	DefaultProfile_props FieldProperties
	ActualProfile_props FieldProperties
	//
 	// Any lookups will be added below
	//
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
	// Dynamically generated 15/02/2023 by matttownsend (Matt Townsend) on silicon.local 
	// Page Infrastructure
	Feature_Name      = "Feature"
	Feature_Title       = "Feature"
	Feature_SQLTable    = "featureStore"
	Feature_SQLSearchID = "featureID"
	Feature_QueryString = "FeatureID"
	///
	/// Template Path Defintions
	///
	Feature_Template     = "Feature"
	Feature_TemplateList = "/Feature/FeatureList"
	Feature_TemplateView = "/Feature/FeatureView"
	Feature_TemplateEdit = "/Feature/FeatureEdit"
	Feature_TemplateNew  = "/Feature/FeatureNew"
	///
	/// URI Handler Paths
	///
	Feature_Path       = "/API/Feature/"
	Feature_PathList   = "/FeatureList/"
	Feature_PathView   = "/FeatureView/"
	Feature_PathEdit   = "/FeatureEdit/"
	Feature_PathNew    = "/FeatureNew/"
	Feature_PathSave   = "/FeatureSave/"
	Feature_PathDelete = "/FeatureDelete/"
	// Redirects - On Server Side Error
	Feature_PathEditException   = "/FeatureEditException/"
	Feature_PathNewException    = "/FeatureNewException/"
	//
	//Feature_Redirect provides a page to return to aftern an action
	Feature_Redirect = Feature_PathList
	
	//
	//
	// SQL Field Definitions
	//
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
///
)

//Feature_PageList provides the information for the template for a list of Features
type Feature_PageList struct {
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
	ItemList  		 []Feature
}

//Feature_Page provides the information for the template for an individual Feature
type Feature_Page struct {
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
	SYSId         string
	FeatureID         string
	EstimationSessionID         string
	ConfidenceID         string
	Name         string
	DevEstimate         string
	DevUplift         string
	Reqs         string
	AnalystTest         string
	Docs         string
	Mgt         string
	UatSupport         string
	Marketing         string
	Contingency         string
	TrackerID         string
	AdoID         string
	FreshdeskID         string
	ExtRef         string
	ExtRef2         string
	SYSCreated         string
	SYSCreatedBy         string
	SYSCreatedHost         string
	SYSUpdated         string
	SYSUpdatedBy         string
	SYSUpdatedHost         string
	SYSDeleted         string
	SYSDeletedBy         string
	SYSDeletedHost         string
	Developer         string
	Approver         string
	Notes         string
	OffProfile         string
	OffProfileJustification         string
	SYSActivity         string
	DfReqs         string
	DfAnalystTest         string
	DfDocs         string
	Dfmgt         string
	DfuatSupport         string
	Dfmarketing         string
	Dfcontingency         string
	DfdevUplift         string
	Total         string
	SYSDbVersion         string
	Comments         string
	Description         string
	Analyst         string
	ProductManager         string
	ProjectManager         string
	Training         string
	DfTraining         string
	DefaultProfile         string
	ActualProfile         string
	/// Field Properties
	SYSId_props     FieldProperties
	FeatureID_props     FieldProperties
	EstimationSessionID_props     FieldProperties
	ConfidenceID_props     FieldProperties
	Name_props     FieldProperties
	DevEstimate_props     FieldProperties
	DevUplift_props     FieldProperties
	Reqs_props     FieldProperties
	AnalystTest_props     FieldProperties
	Docs_props     FieldProperties
	Mgt_props     FieldProperties
	UatSupport_props     FieldProperties
	Marketing_props     FieldProperties
	Contingency_props     FieldProperties
	TrackerID_props     FieldProperties
	AdoID_props     FieldProperties
	FreshdeskID_props     FieldProperties
	ExtRef_props     FieldProperties
	ExtRef2_props     FieldProperties
	SYSCreated_props     FieldProperties
	SYSCreatedBy_props     FieldProperties
	SYSCreatedHost_props     FieldProperties
	SYSUpdated_props     FieldProperties
	SYSUpdatedBy_props     FieldProperties
	SYSUpdatedHost_props     FieldProperties
	SYSDeleted_props     FieldProperties
	SYSDeletedBy_props     FieldProperties
	SYSDeletedHost_props     FieldProperties
	Developer_props     FieldProperties
	Approver_props     FieldProperties
	Notes_props     FieldProperties
	OffProfile_props     FieldProperties
	OffProfileJustification_props     FieldProperties
	SYSActivity_props     FieldProperties
	DfReqs_props     FieldProperties
	DfAnalystTest_props     FieldProperties
	DfDocs_props     FieldProperties
	Dfmgt_props     FieldProperties
	DfuatSupport_props     FieldProperties
	Dfmarketing_props     FieldProperties
	Dfcontingency_props     FieldProperties
	DfdevUplift_props     FieldProperties
	Total_props     FieldProperties
	SYSDbVersion_props     FieldProperties
	Comments_props     FieldProperties
	Description_props     FieldProperties
	Analyst_props     FieldProperties
	ProductManager_props     FieldProperties
	ProjectManager_props     FieldProperties
	Training_props     FieldProperties
	DfTraining_props     FieldProperties
	DefaultProfile_props     FieldProperties
	ActualProfile_props     FieldProperties
	/// Lookups
	EstimationSessionID_lookup    []Lookup_Item
	ConfidenceID_lookup    []Lookup_Item
	Developer_lookup    []Lookup_Item
	Approver_lookup    []Lookup_Item
	Analyst_lookup    []Lookup_Item
	ProductManager_lookup    []Lookup_Item
	ProjectManager_lookup    []Lookup_Item
	DefaultProfile_lookup    []Lookup_Item
	ActualProfile_lookup    []Lookup_Item
	
	/// END OF DEFINITIONS
	///
}