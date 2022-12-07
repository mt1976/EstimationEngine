package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/project.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Project (project)
// Endpoint 	        : Project (ProjectID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 27/11/2022 at 20:46:14
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

const (

	///
	/// Handler Defintions
	///
	Feature_ByEstimationSession_TemplateList = "/Feature/Feature_ByEstimationSessionList_impl"
	///
	/// Handler Monitor Paths
	///
	Feature_ByEstimationSession_PathList    = "/FeatureByEstimationSession/"
	Feature_ByEstimationSession_QueryString = "EstimationSessionID"

	Feature_SoftDelete_Path        = "/FeatureRemove/"
	Feature_SoftDelete_QueryString = Feature_QueryString
	Feature_PathStore              = "/FeatureStore/"
	///

)

// project_Page provides the information for the template for an individual Project
type Feature_ByEstimationSession_List struct {
	SessionInfo SessionInfo
	UserMenu    AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	Origin            Origin
	Project           Project
	EstimationSession EstimationSession
	ItemsOnPage       int
	ItemList          []Feature
	Context           appContext
	CCY               string
	CCYCode           string
}
