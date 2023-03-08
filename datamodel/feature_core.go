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
// Date & Time		    : 08/03/2023 at 18:42:23
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Feature defines the datamodel for the Feature object
type Feature struct {
	// Dynamically generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	SYSId       string
	FeatureID       string
	EstimationSessionID       string
	ConfidenceCODE       string
	Name       string
	DevelopmentEstimate       string
	DevelopmentFactored       string
	Requirements       string
	AnalystTest       string
	Documentation       string
	Management       string
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
	DeveloperResource       string
	ApproverResource       string
	Activity       string
	OffProfile       string
	OffProfileJustification       string
	SYSActivity       string
	RequirementsDefault       string
	AnalystTestDefault       string
	DocumentationDefault       string
	ManagementDefault       string
	UatSupportDefault       string
	MarketingDefault       string
	ContingencyDefault       string
	DevelopmentFactoredDefault       string
	Total       string
	SYSDbVersion       string
	Comments       string
	Description       string
	AnalystResource       string
	ProductManagerResource       string
	ProjectManagerResource       string
	Training       string
	TrainingDefault       string
	ProfileDefault       string
	ProfileSelected       string
	EstimateEffort       string
	EstimateEffortDefault       string
	AnalystEstimate       string
	TotalDefault       string
	OriginName       string
	ProjectName       string
	OriginID       string
	ProjectID       string
	ProfileSelectedOld       string
	CCY       string
	RoundTo       string
	OriginCode       string
	EstimationSessionName       string
	//
	// Field Properties
	//
	SYSId_props FieldProperties
	FeatureID_props FieldProperties
	EstimationSessionID_props FieldProperties
	ConfidenceCODE_props FieldProperties
	Name_props FieldProperties
	DevelopmentEstimate_props FieldProperties
	DevelopmentFactored_props FieldProperties
	Requirements_props FieldProperties
	AnalystTest_props FieldProperties
	Documentation_props FieldProperties
	Management_props FieldProperties
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
	DeveloperResource_props FieldProperties
	ApproverResource_props FieldProperties
	Activity_props FieldProperties
	OffProfile_props FieldProperties
	OffProfileJustification_props FieldProperties
	SYSActivity_props FieldProperties
	RequirementsDefault_props FieldProperties
	AnalystTestDefault_props FieldProperties
	DocumentationDefault_props FieldProperties
	ManagementDefault_props FieldProperties
	UatSupportDefault_props FieldProperties
	MarketingDefault_props FieldProperties
	ContingencyDefault_props FieldProperties
	DevelopmentFactoredDefault_props FieldProperties
	Total_props FieldProperties
	SYSDbVersion_props FieldProperties
	Comments_props FieldProperties
	Description_props FieldProperties
	AnalystResource_props FieldProperties
	ProductManagerResource_props FieldProperties
	ProjectManagerResource_props FieldProperties
	Training_props FieldProperties
	TrainingDefault_props FieldProperties
	ProfileDefault_props FieldProperties
	ProfileSelected_props FieldProperties
	EstimateEffort_props FieldProperties
	EstimateEffortDefault_props FieldProperties
	AnalystEstimate_props FieldProperties
	TotalDefault_props FieldProperties
	OriginName_props FieldProperties
	ProjectName_props FieldProperties
	OriginID_props FieldProperties
	ProjectID_props FieldProperties
	ProfileSelectedOld_props FieldProperties
	CCY_props FieldProperties
	RoundTo_props FieldProperties
	OriginCode_props FieldProperties
	EstimationSessionName_props FieldProperties
	//
 	// Any lookups will be added below
	//
	EstimationSessionID_lookup []Lookup_Item
	ConfidenceCODE_lookup []Lookup_Item
	DeveloperResource_lookup []Lookup_Item
	ApproverResource_lookup []Lookup_Item
	OffProfile_lookup []Lookup_Item
	AnalystResource_lookup []Lookup_Item
	ProductManagerResource_lookup []Lookup_Item
	ProjectManagerResource_lookup []Lookup_Item
	ProfileDefault_lookup []Lookup_Item
	ProfileSelected_lookup []Lookup_Item
	}

const (
	// Dynamically generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	Feature_ConfidenceCODE_sql   = "confidenceCODE" // ConfidenceCODE is a String
	Feature_Name_sql   = "name" // Name is a String
	Feature_DevelopmentEstimate_sql   = "developmentEstimate" // DevelopmentEstimate is a String
	Feature_DevelopmentFactored_sql   = "developmentFactored" // DevelopmentFactored is a String
	Feature_Requirements_sql   = "requirements" // Requirements is a String
	Feature_AnalystTest_sql   = "analystTest" // AnalystTest is a String
	Feature_Documentation_sql   = "documentation" // Documentation is a String
	Feature_Management_sql   = "management" // Management is a String
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
	Feature_DeveloperResource_sql   = "developerResource" // DeveloperResource is a String
	Feature_ApproverResource_sql   = "approverResource" // ApproverResource is a String
	Feature_Activity_sql   = "activity" // Activity is a String
	Feature_OffProfile_sql   = "offProfile" // OffProfile is a String
	Feature_OffProfileJustification_sql   = "offProfileJustification" // OffProfileJustification is a String
	Feature_SYSActivity_sql   = "_activity" // SYSActivity is a String
	Feature_RequirementsDefault_sql   = "requirementsDefault" // RequirementsDefault is a String
	Feature_AnalystTestDefault_sql   = "analystTestDefault" // AnalystTestDefault is a String
	Feature_DocumentationDefault_sql   = "documentationDefault" // DocumentationDefault is a String
	Feature_ManagementDefault_sql   = "managementDefault" // ManagementDefault is a String
	Feature_UatSupportDefault_sql   = "uatSupportDefault" // UatSupportDefault is a String
	Feature_MarketingDefault_sql   = "marketingDefault" // MarketingDefault is a String
	Feature_ContingencyDefault_sql   = "contingencyDefault" // ContingencyDefault is a String
	Feature_DevelopmentFactoredDefault_sql   = "developmentFactoredDefault" // DevelopmentFactoredDefault is a String
	Feature_Total_sql   = "total" // Total is a String
	Feature_SYSDbVersion_sql   = "_dbVersion" // SYSDbVersion is a String
	Feature_Comments_sql   = "comments" // Comments is a String
	Feature_Description_sql   = "description" // Description is a String
	Feature_AnalystResource_sql   = "analystResource" // AnalystResource is a String
	Feature_ProductManagerResource_sql   = "productManagerResource" // ProductManagerResource is a String
	Feature_ProjectManagerResource_sql   = "projectManagerResource" // ProjectManagerResource is a String
	Feature_Training_sql   = "training" // Training is a String
	Feature_TrainingDefault_sql   = "trainingDefault" // TrainingDefault is a String
	Feature_ProfileDefault_sql   = "profileDefault" // ProfileDefault is a String
	Feature_ProfileSelected_sql   = "profileSelected" // ProfileSelected is a String
	Feature_EstimateEffort_sql   = "estimateEffort" // EstimateEffort is a String
	Feature_EstimateEffortDefault_sql   = "estimateEffortDefault" // EstimateEffortDefault is a String
	Feature_AnalystEstimate_sql   = "analystEstimate" // AnalystEstimate is a String
	Feature_TotalDefault_sql   = "totalDefault" // TotalDefault is a String
	Feature_OriginName_sql   = "OriginName" // OriginName is a String
	Feature_ProjectName_sql   = "ProjectName" // ProjectName is a String
	Feature_OriginID_sql   = "OriginID" // OriginID is a String
	Feature_ProjectID_sql   = "ProjectID" // ProjectID is a String
	Feature_ProfileSelectedOld_sql   = "ProfileSelectedOld" // ProfileSelectedOld is a String
	Feature_CCY_sql   = "CCY" // CCY is a String
	Feature_RoundTo_sql   = "RoundTo" // RoundTo is a String
	Feature_OriginCode_sql   = "OriginCode" // OriginCode is a String
	Feature_EstimationSessionName_sql   = "EstimationSessionName" // EstimationSessionName is a String
///
	/// Application Field Definitions
	///
	Feature_SYSId_scrn   = "SYSId" // SYSId is a Int
	Feature_FeatureID_scrn   = "FeatureID" // FeatureID is a String
	Feature_EstimationSessionID_scrn   = "EstimationSessionID" // EstimationSessionID is a String
	Feature_ConfidenceCODE_scrn   = "ConfidenceCODE" // ConfidenceCODE is a String
	Feature_Name_scrn   = "Name" // Name is a String
	Feature_DevelopmentEstimate_scrn   = "DevelopmentEstimate" // DevelopmentEstimate is a String
	Feature_DevelopmentFactored_scrn   = "DevelopmentFactored" // DevelopmentFactored is a String
	Feature_Requirements_scrn   = "Requirements" // Requirements is a String
	Feature_AnalystTest_scrn   = "AnalystTest" // AnalystTest is a String
	Feature_Documentation_scrn   = "Documentation" // Documentation is a String
	Feature_Management_scrn   = "Management" // Management is a String
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
	Feature_DeveloperResource_scrn   = "DeveloperResource" // DeveloperResource is a String
	Feature_ApproverResource_scrn   = "ApproverResource" // ApproverResource is a String
	Feature_Activity_scrn   = "Activity" // Activity is a String
	Feature_OffProfile_scrn   = "OffProfile" // OffProfile is a String
	Feature_OffProfileJustification_scrn   = "OffProfileJustification" // OffProfileJustification is a String
	Feature_SYSActivity_scrn   = "SYSActivity" // SYSActivity is a String
	Feature_RequirementsDefault_scrn   = "RequirementsDefault" // RequirementsDefault is a String
	Feature_AnalystTestDefault_scrn   = "AnalystTestDefault" // AnalystTestDefault is a String
	Feature_DocumentationDefault_scrn   = "DocumentationDefault" // DocumentationDefault is a String
	Feature_ManagementDefault_scrn   = "ManagementDefault" // ManagementDefault is a String
	Feature_UatSupportDefault_scrn   = "UatSupportDefault" // UatSupportDefault is a String
	Feature_MarketingDefault_scrn   = "MarketingDefault" // MarketingDefault is a String
	Feature_ContingencyDefault_scrn   = "ContingencyDefault" // ContingencyDefault is a String
	Feature_DevelopmentFactoredDefault_scrn   = "DevelopmentFactoredDefault" // DevelopmentFactoredDefault is a String
	Feature_Total_scrn   = "Total" // Total is a String
	Feature_SYSDbVersion_scrn   = "SYSDbVersion" // SYSDbVersion is a String
	Feature_Comments_scrn   = "Comments" // Comments is a String
	Feature_Description_scrn   = "Description" // Description is a String
	Feature_AnalystResource_scrn   = "AnalystResource" // AnalystResource is a String
	Feature_ProductManagerResource_scrn   = "ProductManagerResource" // ProductManagerResource is a String
	Feature_ProjectManagerResource_scrn   = "ProjectManagerResource" // ProjectManagerResource is a String
	Feature_Training_scrn   = "Training" // Training is a String
	Feature_TrainingDefault_scrn   = "TrainingDefault" // TrainingDefault is a String
	Feature_ProfileDefault_scrn   = "ProfileDefault" // ProfileDefault is a String
	Feature_ProfileSelected_scrn   = "ProfileSelected" // ProfileSelected is a String
	Feature_EstimateEffort_scrn   = "EstimateEffort" // EstimateEffort is a String
	Feature_EstimateEffortDefault_scrn   = "EstimateEffortDefault" // EstimateEffortDefault is a String
	Feature_AnalystEstimate_scrn   = "AnalystEstimate" // AnalystEstimate is a String
	Feature_TotalDefault_scrn   = "TotalDefault" // TotalDefault is a String
	Feature_OriginName_scrn   = "OriginName" // OriginName is a String
	Feature_ProjectName_scrn   = "ProjectName" // ProjectName is a String
	Feature_OriginID_scrn   = "OriginID" // OriginID is a String
	Feature_ProjectID_scrn   = "ProjectID" // ProjectID is a String
	Feature_ProfileSelectedOld_scrn   = "ProfileSelectedOld" // ProfileSelectedOld is a String
	Feature_CCY_scrn   = "CCY" // CCY is a String
	Feature_RoundTo_scrn   = "RoundTo" // RoundTo is a String
	Feature_OriginCode_scrn   = "OriginCode" // OriginCode is a String
	Feature_EstimationSessionName_scrn   = "EstimationSessionName" // EstimationSessionName is a String
///
)

//Feature_PageList provides the information for the template for a list of Features
type Feature_PageList struct {
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
	ItemList  		 []Feature
}

//Feature_Page provides the information for the template for an individual Feature
type Feature_Page struct {
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
	SYSId         string
	FeatureID         string
	EstimationSessionID         string
	ConfidenceCODE         string
	Name         string
	DevelopmentEstimate         string
	DevelopmentFactored         string
	Requirements         string
	AnalystTest         string
	Documentation         string
	Management         string
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
	DeveloperResource         string
	ApproverResource         string
	Activity         string
	OffProfile         string
	OffProfileJustification         string
	SYSActivity         string
	RequirementsDefault         string
	AnalystTestDefault         string
	DocumentationDefault         string
	ManagementDefault         string
	UatSupportDefault         string
	MarketingDefault         string
	ContingencyDefault         string
	DevelopmentFactoredDefault         string
	Total         string
	SYSDbVersion         string
	Comments         string
	Description         string
	AnalystResource         string
	ProductManagerResource         string
	ProjectManagerResource         string
	Training         string
	TrainingDefault         string
	ProfileDefault         string
	ProfileSelected         string
	EstimateEffort         string
	EstimateEffortDefault         string
	AnalystEstimate         string
	TotalDefault         string
	OriginName         string
	ProjectName         string
	OriginID         string
	ProjectID         string
	ProfileSelectedOld         string
	CCY         string
	RoundTo         string
	OriginCode         string
	EstimationSessionName         string
	/// Field Properties
	SYSId_props     FieldProperties
	FeatureID_props     FieldProperties
	EstimationSessionID_props     FieldProperties
	ConfidenceCODE_props     FieldProperties
	Name_props     FieldProperties
	DevelopmentEstimate_props     FieldProperties
	DevelopmentFactored_props     FieldProperties
	Requirements_props     FieldProperties
	AnalystTest_props     FieldProperties
	Documentation_props     FieldProperties
	Management_props     FieldProperties
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
	DeveloperResource_props     FieldProperties
	ApproverResource_props     FieldProperties
	Activity_props     FieldProperties
	OffProfile_props     FieldProperties
	OffProfileJustification_props     FieldProperties
	SYSActivity_props     FieldProperties
	RequirementsDefault_props     FieldProperties
	AnalystTestDefault_props     FieldProperties
	DocumentationDefault_props     FieldProperties
	ManagementDefault_props     FieldProperties
	UatSupportDefault_props     FieldProperties
	MarketingDefault_props     FieldProperties
	ContingencyDefault_props     FieldProperties
	DevelopmentFactoredDefault_props     FieldProperties
	Total_props     FieldProperties
	SYSDbVersion_props     FieldProperties
	Comments_props     FieldProperties
	Description_props     FieldProperties
	AnalystResource_props     FieldProperties
	ProductManagerResource_props     FieldProperties
	ProjectManagerResource_props     FieldProperties
	Training_props     FieldProperties
	TrainingDefault_props     FieldProperties
	ProfileDefault_props     FieldProperties
	ProfileSelected_props     FieldProperties
	EstimateEffort_props     FieldProperties
	EstimateEffortDefault_props     FieldProperties
	AnalystEstimate_props     FieldProperties
	TotalDefault_props     FieldProperties
	OriginName_props     FieldProperties
	ProjectName_props     FieldProperties
	OriginID_props     FieldProperties
	ProjectID_props     FieldProperties
	ProfileSelectedOld_props     FieldProperties
	CCY_props     FieldProperties
	RoundTo_props     FieldProperties
	OriginCode_props     FieldProperties
	EstimationSessionName_props     FieldProperties
	/// Lookups
	EstimationSessionID_lookup    []Lookup_Item
	ConfidenceCODE_lookup    []Lookup_Item
	DeveloperResource_lookup    []Lookup_Item
	ApproverResource_lookup    []Lookup_Item
	OffProfile_lookup    []Lookup_Item
	AnalystResource_lookup    []Lookup_Item
	ProductManagerResource_lookup    []Lookup_Item
	ProjectManagerResource_lookup    []Lookup_Item
	ProfileDefault_lookup    []Lookup_Item
	ProfileSelected_lookup    []Lookup_Item
	
	/// END OF DEFINITIONS
	///
}