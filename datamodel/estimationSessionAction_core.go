package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/estimationsessionaction.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : EstimationSessionAction (estimationsessionaction)
// Endpoint 	        : EstimationSessionAction (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 08/12/2022 at 13:31:30
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//EstimationSessionAction defines the datamolde for the EstimationSessionAction object
type EstimationSessionAction struct {


ID       string
ID_props FieldProperties
EstimationSession       string
EstimationSession_props FieldProperties
Action       string
Action_props FieldProperties
Notes       string
Notes_props FieldProperties
 // Any lookups will be added below
EstimationSession_lookup []Lookup_Item
Action_lookup []Lookup_Item


}

const (
	EstimationSessionAction_Title       = "Estimation Session Actions"
	EstimationSessionAction_SQLTable    = "n/a"
	EstimationSessionAction_SQLSearchID = "n/a"
	EstimationSessionAction_QueryString = "ID"
	///
	/// Handler Defintions
	///
	EstimationSessionAction_Template     = "EstimationSessionAction"
	EstimationSessionAction_TemplateList = "/EstimationSessionAction/EstimationSessionAction_List"
	EstimationSessionAction_TemplateView = "/EstimationSessionAction/EstimationSessionAction_View"
	EstimationSessionAction_TemplateEdit = "/EstimationSessionAction/EstimationSessionAction_Edit"
	EstimationSessionAction_TemplateNew  = "/EstimationSessionAction/EstimationSessionAction_New"
	///
	/// Handler Monitor Paths
	///
	EstimationSessionAction_Path       = "/API/EstimationSessionAction/"
	EstimationSessionAction_PathList   = "/EstimationSessionActionList/"
	EstimationSessionAction_PathView   = "/EstimationSessionActionView/"
	EstimationSessionAction_PathEdit   = "/EstimationSessionActionEdit/"
	EstimationSessionAction_PathNew    = "/EstimationSessionActionNew/"
	EstimationSessionAction_PathSave   = "/EstimationSessionActionSave/"
	EstimationSessionAction_PathDelete = "/EstimationSessionActionDelete/"
	///
	//EstimationSessionAction_Redirect provides a page to return to aftern an action
	EstimationSessionAction_Redirect = EstimationSessionAction_PathView
	///
	///
	/// SQL Field Definitions
	///
	EstimationSessionAction_ID_sql   = "ID" // ID is a String
	EstimationSessionAction_EstimationSession_sql   = "EstimationSession" // EstimationSession is a String
	EstimationSessionAction_Action_sql   = "Action" // Action is a String
	EstimationSessionAction_Notes_sql   = "Notes" // Notes is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	EstimationSessionAction_ID_scrn   = "ID" // ID is a String
	EstimationSessionAction_EstimationSession_scrn   = "EstimationSession" // EstimationSession is a String
	EstimationSessionAction_Action_scrn   = "Action" // Action is a String
	EstimationSessionAction_Notes_scrn   = "Notes" // Notes is a String

	/// Definitions End
	///
)

//estimationsessionaction_PageList provides the information for the template for a list of EstimationSessionActions
type EstimationSessionAction_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []EstimationSessionAction
	Context	 appContext
}

//estimationsessionaction_Page provides the information for the template for an individual EstimationSessionAction
type EstimationSessionAction_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	ID         string
	ID_props     FieldProperties
	EstimationSession         string
	EstimationSession_lookup    []Lookup_Item
	EstimationSession_props     FieldProperties
	Action         string
	Action_lookup    []Lookup_Item
	Action_props     FieldProperties
	Notes         string
	Notes_props     FieldProperties
	// 
	// Dynamically generated 08/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}