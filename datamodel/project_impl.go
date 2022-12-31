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
	Project_Origin_TemplateList = "/Project/Project_ByOriginList_impl"
	///
	/// Handler Monitor Paths
	///
	Project_Origin_PathList    = "/ProjectsByOrigin/"
	Project_Origin_QueryString = "OriginID"
	Project_PathRemove         = "/ProjectRemove/"
	Project_PathUpdate         = "/ProjectUpdate/"
	///

)

// project_Page provides the information for the template for an individual Project
type Project_Origin_List struct {
	SessionInfo SessionInfo
	UserMenu    AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	Origin      Origin
	ItemsOnPage int
	ItemList    []Project
	Context     appContext
}
