package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/estimationsession.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : EstimationSession (estimationsession)
// Endpoint 	        : EstimationSession (EstimationSessionID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 03/01/2023 at 20:58:19
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//EstimationSession defines the datamolde for the EstimationSession object
type EstimationSession struct {


SYSId       string
SYSId_props FieldProperties
EstimationSessionID       string
EstimationSessionID_props FieldProperties
ProjectID       string
ProjectID_props FieldProperties
EstimationStateID       string
EstimationStateID_props FieldProperties
Notes       string
Notes_props FieldProperties
Releases       string
Releases_props FieldProperties
Total       string
Total_props FieldProperties
Contingency       string
Contingency_props FieldProperties
ReqDays       string
ReqDays_props FieldProperties
RegCost       string
RegCost_props FieldProperties
ImpDays       string
ImpDays_props FieldProperties
ImpCost       string
ImpCost_props FieldProperties
UatDays       string
UatDays_props FieldProperties
UatCost       string
UatCost_props FieldProperties
MgtDays       string
MgtDays_props FieldProperties
MgtCost       string
MgtCost_props FieldProperties
RelDays       string
RelDays_props FieldProperties
RelCost       string
RelCost_props FieldProperties
SupportUplift       string
SupportUplift_props FieldProperties
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
Name       string
Name_props FieldProperties
AdoID       string
AdoID_props FieldProperties
FreshdeskID       string
FreshdeskID_props FieldProperties
TrackerID       string
TrackerID_props FieldProperties
EstRef       string
EstRef_props FieldProperties
ExtRef       string
ExtRef_props FieldProperties
SYSActivity       string
SYSActivity_props FieldProperties
SYSDbVersion       string
SYSDbVersion_props FieldProperties
Comments       string
Comments_props FieldProperties
ProjectManager       string
ProjectManager_props FieldProperties
ProductManager       string
ProductManager_props FieldProperties
Approver       string
Approver_props FieldProperties
Origin       string
Origin_props FieldProperties
OriginStateID       string
OriginStateID_props FieldProperties
OriginState       string
OriginState_props FieldProperties
OriginDocTypeID       string
OriginDocTypeID_props FieldProperties
OriginDocType       string
OriginDocType_props FieldProperties
OriginCode       string
OriginCode_props FieldProperties
OriginName       string
OriginName_props FieldProperties
OriginRate       string
OriginRate_props FieldProperties
ProjectProfileID       string
ProjectProfileID_props FieldProperties
ProjectProfile       string
ProjectProfile_props FieldProperties
ProjectDefaultReleases       string
ProjectDefaultReleases_props FieldProperties
ProjectDefaultReleaseHours       string
ProjectDefaultReleaseHours_props FieldProperties
ProjectBlendedRate       string
ProjectBlendedRate_props FieldProperties
ProjectStateID       string
ProjectStateID_props FieldProperties
ProjectState       string
ProjectState_props FieldProperties
ProjectName       string
ProjectName_props FieldProperties
ProjectStartDate       string
ProjectStartDate_props FieldProperties
ProjectEndDate       string
ProjectEndDate_props FieldProperties
ProfileSupportUpliftPerc       string
ProfileSupportUpliftPerc_props FieldProperties
CCY       string
CCY_props FieldProperties
CCYCode       string
CCYCode_props FieldProperties
EffortTotal       string
EffortTotal_props FieldProperties
FreshDeskURI       string
FreshDeskURI_props FieldProperties
ADOURI       string
ADOURI_props FieldProperties
NoActiveFeatures       string
NoActiveFeatures_props FieldProperties
 // Any lookups will be added below

ProjectID_lookup []Lookup_Item
EstimationStateID_lookup []Lookup_Item

































ProjectManager_lookup []Lookup_Item
ProductManager_lookup []Lookup_Item
Approver_lookup []Lookup_Item


























}

const (
	EstimationSession_Title       = "Estimation Session"
	EstimationSession_SQLTable    = "estimationSessionStore"
	EstimationSession_SQLSearchID = "estimationSessionID"
	EstimationSession_QueryString = "EstimationSessionID"
	///
	/// Handler Defintions
	///
	EstimationSession_Template     = "EstimationSession"
	EstimationSession_TemplateList = "/EstimationSession/EstimationSession_List"
	EstimationSession_TemplateView = "/EstimationSession/EstimationSession_View"
	EstimationSession_TemplateEdit = "/EstimationSession/EstimationSession_Edit"
	EstimationSession_TemplateNew  = "/EstimationSession/EstimationSession_New"
	///
	/// Handler Monitor Paths
	///
	EstimationSession_Path       = "/API/EstimationSession/"
	EstimationSession_PathList   = "/EstimationSessionList/"
	EstimationSession_PathView   = "/EstimationSessionView/"
	EstimationSession_PathEdit   = "/EstimationSessionEdit/"
	EstimationSession_PathNew    = "/EstimationSessionNew/"
	EstimationSession_PathSave   = "/EstimationSessionSave/"
	EstimationSession_PathDelete = "/EstimationSessionDelete/"
	///
	//EstimationSession_Redirect provides a page to return to aftern an action
	EstimationSession_Redirect = EstimationSession_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	EstimationSession_SYSId_sql   = "_id" // SYSId is a Int
	EstimationSession_EstimationSessionID_sql   = "estimationSessionID" // EstimationSessionID is a String
	EstimationSession_ProjectID_sql   = "projectID" // ProjectID is a String
	EstimationSession_EstimationStateID_sql   = "estimationStateID" // EstimationStateID is a String
	EstimationSession_Notes_sql   = "notes" // Notes is a String
	EstimationSession_Releases_sql   = "releases" // Releases is a String
	EstimationSession_Total_sql   = "total" // Total is a String
	EstimationSession_Contingency_sql   = "contingency" // Contingency is a String
	EstimationSession_ReqDays_sql   = "reqDays" // ReqDays is a String
	EstimationSession_RegCost_sql   = "regCost" // RegCost is a String
	EstimationSession_ImpDays_sql   = "impDays" // ImpDays is a String
	EstimationSession_ImpCost_sql   = "impCost" // ImpCost is a String
	EstimationSession_UatDays_sql   = "uatDays" // UatDays is a String
	EstimationSession_UatCost_sql   = "uatCost" // UatCost is a String
	EstimationSession_MgtDays_sql   = "mgtDays" // MgtDays is a String
	EstimationSession_MgtCost_sql   = "mgtCost" // MgtCost is a String
	EstimationSession_RelDays_sql   = "relDays" // RelDays is a String
	EstimationSession_RelCost_sql   = "relCost" // RelCost is a String
	EstimationSession_SupportUplift_sql   = "supportUplift" // SupportUplift is a String
	EstimationSession_SYSCreated_sql   = "_created" // SYSCreated is a String
	EstimationSession_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	EstimationSession_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	EstimationSession_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	EstimationSession_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
	EstimationSession_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
	EstimationSession_SYSDeleted_sql   = "_deleted" // SYSDeleted is a String
	EstimationSession_SYSDeletedBy_sql   = "_deletedBy" // SYSDeletedBy is a String
	EstimationSession_SYSDeletedHost_sql   = "_deletedHost" // SYSDeletedHost is a String
	EstimationSession_Name_sql   = "name" // Name is a String
	EstimationSession_AdoID_sql   = "adoID" // AdoID is a String
	EstimationSession_FreshdeskID_sql   = "freshdeskID" // FreshdeskID is a String
	EstimationSession_TrackerID_sql   = "trackerID" // TrackerID is a String
	EstimationSession_EstRef_sql   = "estRef" // EstRef is a String
	EstimationSession_ExtRef_sql   = "extRef" // ExtRef is a String
	EstimationSession_SYSActivity_sql   = "_activity" // SYSActivity is a String
	EstimationSession_SYSDbVersion_sql   = "_dbVersion" // SYSDbVersion is a String
	EstimationSession_Comments_sql   = "comments" // Comments is a String
	EstimationSession_ProjectManager_sql   = "projectManager" // ProjectManager is a String
	EstimationSession_ProductManager_sql   = "productManager" // ProductManager is a String
	EstimationSession_Approver_sql   = "approver" // Approver is a String
	EstimationSession_Origin_sql   = "Origin" // Origin is a String
	EstimationSession_OriginStateID_sql   = "OriginStateID" // OriginStateID is a String
	EstimationSession_OriginState_sql   = "OriginState" // OriginState is a String
	EstimationSession_OriginDocTypeID_sql   = "OriginDocTypeID" // OriginDocTypeID is a String
	EstimationSession_OriginDocType_sql   = "OriginDocType" // OriginDocType is a String
	EstimationSession_OriginCode_sql   = "OriginCode" // OriginCode is a String
	EstimationSession_OriginName_sql   = "OriginName" // OriginName is a String
	EstimationSession_OriginRate_sql   = "OriginRate" // OriginRate is a String
	EstimationSession_ProjectProfileID_sql   = "ProjectProfileID" // ProjectProfileID is a String
	EstimationSession_ProjectProfile_sql   = "ProjectProfile" // ProjectProfile is a String
	EstimationSession_ProjectDefaultReleases_sql   = "ProjectDefaultReleases" // ProjectDefaultReleases is a String
	EstimationSession_ProjectDefaultReleaseHours_sql   = "ProjectDefaultReleaseHours" // ProjectDefaultReleaseHours is a String
	EstimationSession_ProjectBlendedRate_sql   = "ProjectBlendedRate" // ProjectBlendedRate is a String
	EstimationSession_ProjectStateID_sql   = "ProjectStateID" // ProjectStateID is a String
	EstimationSession_ProjectState_sql   = "ProjectState" // ProjectState is a String
	EstimationSession_ProjectName_sql   = "ProjectName" // ProjectName is a String
	EstimationSession_ProjectStartDate_sql   = "ProjectStartDate" // ProjectStartDate is a String
	EstimationSession_ProjectEndDate_sql   = "ProjectEndDate" // ProjectEndDate is a String
	EstimationSession_ProfileSupportUpliftPerc_sql   = "ProfileSupportUpliftPerc" // ProfileSupportUpliftPerc is a String
	EstimationSession_CCY_sql   = "CCY" // CCY is a String
	EstimationSession_CCYCode_sql   = "CCYCode" // CCYCode is a String
	EstimationSession_EffortTotal_sql   = "EffortTotal" // EffortTotal is a String
	EstimationSession_FreshDeskURI_sql   = "FreshDeskURI" // FreshDeskURI is a String
	EstimationSession_ADOURI_sql   = "ADOURI" // ADOURI is a String
	EstimationSession_NoActiveFeatures_sql   = "NoActiveFeatures" // NoActiveFeatures is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	EstimationSession_SYSId_scrn   = "SYSId" // SYSId is a Int
	EstimationSession_EstimationSessionID_scrn   = "EstimationSessionID" // EstimationSessionID is a String
	EstimationSession_ProjectID_scrn   = "ProjectID" // ProjectID is a String
	EstimationSession_EstimationStateID_scrn   = "EstimationStateID" // EstimationStateID is a String
	EstimationSession_Notes_scrn   = "Notes" // Notes is a String
	EstimationSession_Releases_scrn   = "Releases" // Releases is a String
	EstimationSession_Total_scrn   = "Total" // Total is a String
	EstimationSession_Contingency_scrn   = "Contingency" // Contingency is a String
	EstimationSession_ReqDays_scrn   = "ReqDays" // ReqDays is a String
	EstimationSession_RegCost_scrn   = "RegCost" // RegCost is a String
	EstimationSession_ImpDays_scrn   = "ImpDays" // ImpDays is a String
	EstimationSession_ImpCost_scrn   = "ImpCost" // ImpCost is a String
	EstimationSession_UatDays_scrn   = "UatDays" // UatDays is a String
	EstimationSession_UatCost_scrn   = "UatCost" // UatCost is a String
	EstimationSession_MgtDays_scrn   = "MgtDays" // MgtDays is a String
	EstimationSession_MgtCost_scrn   = "MgtCost" // MgtCost is a String
	EstimationSession_RelDays_scrn   = "RelDays" // RelDays is a String
	EstimationSession_RelCost_scrn   = "RelCost" // RelCost is a String
	EstimationSession_SupportUplift_scrn   = "SupportUplift" // SupportUplift is a String
	EstimationSession_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	EstimationSession_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	EstimationSession_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	EstimationSession_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	EstimationSession_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
	EstimationSession_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
	EstimationSession_SYSDeleted_scrn   = "SYSDeleted" // SYSDeleted is a String
	EstimationSession_SYSDeletedBy_scrn   = "SYSDeletedBy" // SYSDeletedBy is a String
	EstimationSession_SYSDeletedHost_scrn   = "SYSDeletedHost" // SYSDeletedHost is a String
	EstimationSession_Name_scrn   = "Name" // Name is a String
	EstimationSession_AdoID_scrn   = "AdoID" // AdoID is a String
	EstimationSession_FreshdeskID_scrn   = "FreshdeskID" // FreshdeskID is a String
	EstimationSession_TrackerID_scrn   = "TrackerID" // TrackerID is a String
	EstimationSession_EstRef_scrn   = "EstRef" // EstRef is a String
	EstimationSession_ExtRef_scrn   = "ExtRef" // ExtRef is a String
	EstimationSession_SYSActivity_scrn   = "SYSActivity" // SYSActivity is a String
	EstimationSession_SYSDbVersion_scrn   = "SYSDbVersion" // SYSDbVersion is a String
	EstimationSession_Comments_scrn   = "Comments" // Comments is a String
	EstimationSession_ProjectManager_scrn   = "ProjectManager" // ProjectManager is a String
	EstimationSession_ProductManager_scrn   = "ProductManager" // ProductManager is a String
	EstimationSession_Approver_scrn   = "Approver" // Approver is a String
	EstimationSession_Origin_scrn   = "Origin" // Origin is a String
	EstimationSession_OriginStateID_scrn   = "OriginStateID" // OriginStateID is a String
	EstimationSession_OriginState_scrn   = "OriginState" // OriginState is a String
	EstimationSession_OriginDocTypeID_scrn   = "OriginDocTypeID" // OriginDocTypeID is a String
	EstimationSession_OriginDocType_scrn   = "OriginDocType" // OriginDocType is a String
	EstimationSession_OriginCode_scrn   = "OriginCode" // OriginCode is a String
	EstimationSession_OriginName_scrn   = "OriginName" // OriginName is a String
	EstimationSession_OriginRate_scrn   = "OriginRate" // OriginRate is a String
	EstimationSession_ProjectProfileID_scrn   = "ProjectProfileID" // ProjectProfileID is a String
	EstimationSession_ProjectProfile_scrn   = "ProjectProfile" // ProjectProfile is a String
	EstimationSession_ProjectDefaultReleases_scrn   = "ProjectDefaultReleases" // ProjectDefaultReleases is a String
	EstimationSession_ProjectDefaultReleaseHours_scrn   = "ProjectDefaultReleaseHours" // ProjectDefaultReleaseHours is a String
	EstimationSession_ProjectBlendedRate_scrn   = "ProjectBlendedRate" // ProjectBlendedRate is a String
	EstimationSession_ProjectStateID_scrn   = "ProjectStateID" // ProjectStateID is a String
	EstimationSession_ProjectState_scrn   = "ProjectState" // ProjectState is a String
	EstimationSession_ProjectName_scrn   = "ProjectName" // ProjectName is a String
	EstimationSession_ProjectStartDate_scrn   = "ProjectStartDate" // ProjectStartDate is a String
	EstimationSession_ProjectEndDate_scrn   = "ProjectEndDate" // ProjectEndDate is a String
	EstimationSession_ProfileSupportUpliftPerc_scrn   = "ProfileSupportUpliftPerc" // ProfileSupportUpliftPerc is a String
	EstimationSession_CCY_scrn   = "CCY" // CCY is a String
	EstimationSession_CCYCode_scrn   = "CCYCode" // CCYCode is a String
	EstimationSession_EffortTotal_scrn   = "EffortTotal" // EffortTotal is a String
	EstimationSession_FreshDeskURI_scrn   = "FreshDeskURI" // FreshDeskURI is a String
	EstimationSession_ADOURI_scrn   = "ADOURI" // ADOURI is a String
	EstimationSession_NoActiveFeatures_scrn   = "NoActiveFeatures" // NoActiveFeatures is a String

	/// Definitions End
	///
)

//estimationsession_PageList provides the information for the template for a list of EstimationSessions
type EstimationSession_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []EstimationSession
	Context	 appContext
}

//estimationsession_Page provides the information for the template for an individual EstimationSession
type EstimationSession_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	SYSId_props     FieldProperties
	EstimationSessionID         string
	EstimationSessionID_props     FieldProperties
	ProjectID         string
	ProjectID_lookup    []Lookup_Item
	ProjectID_props     FieldProperties
	EstimationStateID         string
	EstimationStateID_lookup    []Lookup_Item
	EstimationStateID_props     FieldProperties
	Notes         string
	Notes_props     FieldProperties
	Releases         string
	Releases_props     FieldProperties
	Total         string
	Total_props     FieldProperties
	Contingency         string
	Contingency_props     FieldProperties
	ReqDays         string
	ReqDays_props     FieldProperties
	RegCost         string
	RegCost_props     FieldProperties
	ImpDays         string
	ImpDays_props     FieldProperties
	ImpCost         string
	ImpCost_props     FieldProperties
	UatDays         string
	UatDays_props     FieldProperties
	UatCost         string
	UatCost_props     FieldProperties
	MgtDays         string
	MgtDays_props     FieldProperties
	MgtCost         string
	MgtCost_props     FieldProperties
	RelDays         string
	RelDays_props     FieldProperties
	RelCost         string
	RelCost_props     FieldProperties
	SupportUplift         string
	SupportUplift_props     FieldProperties
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
	Name         string
	Name_props     FieldProperties
	AdoID         string
	AdoID_props     FieldProperties
	FreshdeskID         string
	FreshdeskID_props     FieldProperties
	TrackerID         string
	TrackerID_props     FieldProperties
	EstRef         string
	EstRef_props     FieldProperties
	ExtRef         string
	ExtRef_props     FieldProperties
	SYSActivity         string
	SYSActivity_props     FieldProperties
	SYSDbVersion         string
	SYSDbVersion_props     FieldProperties
	Comments         string
	Comments_props     FieldProperties
	ProjectManager         string
	ProjectManager_lookup    []Lookup_Item
	ProjectManager_props     FieldProperties
	ProductManager         string
	ProductManager_lookup    []Lookup_Item
	ProductManager_props     FieldProperties
	Approver         string
	Approver_lookup    []Lookup_Item
	Approver_props     FieldProperties
	Origin         string
	Origin_props     FieldProperties
	OriginStateID         string
	OriginStateID_props     FieldProperties
	OriginState         string
	OriginState_props     FieldProperties
	OriginDocTypeID         string
	OriginDocTypeID_props     FieldProperties
	OriginDocType         string
	OriginDocType_props     FieldProperties
	OriginCode         string
	OriginCode_props     FieldProperties
	OriginName         string
	OriginName_props     FieldProperties
	OriginRate         string
	OriginRate_props     FieldProperties
	ProjectProfileID         string
	ProjectProfileID_props     FieldProperties
	ProjectProfile         string
	ProjectProfile_props     FieldProperties
	ProjectDefaultReleases         string
	ProjectDefaultReleases_props     FieldProperties
	ProjectDefaultReleaseHours         string
	ProjectDefaultReleaseHours_props     FieldProperties
	ProjectBlendedRate         string
	ProjectBlendedRate_props     FieldProperties
	ProjectStateID         string
	ProjectStateID_props     FieldProperties
	ProjectState         string
	ProjectState_props     FieldProperties
	ProjectName         string
	ProjectName_props     FieldProperties
	ProjectStartDate         string
	ProjectStartDate_props     FieldProperties
	ProjectEndDate         string
	ProjectEndDate_props     FieldProperties
	ProfileSupportUpliftPerc         string
	ProfileSupportUpliftPerc_props     FieldProperties
	CCY         string
	CCY_props     FieldProperties
	CCYCode         string
	CCYCode_props     FieldProperties
	EffortTotal         string
	EffortTotal_props     FieldProperties
	FreshDeskURI         string
	FreshDeskURI_props     FieldProperties
	ADOURI         string
	ADOURI_props     FieldProperties
	NoActiveFeatures         string
	NoActiveFeatures_props     FieldProperties
	// 
	// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}