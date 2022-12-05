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
	EstimationSession_ByProject_TemplateList = "/EstimationSession/EstimationSession_ByProjectList_impl"
	EstimationSession_Formatted_TemplateView = "/EstimationSession/EstimationSession_FormattedView_impl"
	EstimationSession_TemplateSetup          = "/EstimationSession/EstimationSession_Setup_impl"
	EstimationSession_TemplateCreate         = "/EstimationSession/EstimationSession_View_impl"
	///
	/// Handler Monitor Paths
	///
	EstimationSession_ByProject_PathList    = "/EstimationSessionByProject/"
	EstimationSession_ByProject_QueryString = "ProjectID"
	EstimationSession_Formatted_PathView    = "/EstimationSessionFormatted/"
	EstimationSession_Formatted_QueryString = EstimationSession_QueryString

	EstimationSession_PathSetup  = "/EstimationSessionSetup/"
	EstimationSession_PathCreate = "/EstimationSessionCreate/"
	EstimationSession_PathRemove = "/EstimationSessionRemove/"
	EstimationSession_PathClone  = "/EstimationSessionClone/"
	///

)

// project_Page provides the information for the template for an individual Project
type EstimationSession_ByProject_List struct {
	SessionInfo SessionInfo
	UserMenu    AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	Origin      Origin
	Project     Project
	ItemsOnPage int
	ItemList    []EstimationSession
	Context     appContext
	CCY         string
	CCYCode     string
}

type EstimationSession_Formatted_View struct {
}
