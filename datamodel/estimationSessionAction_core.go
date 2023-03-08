package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/estimationsessionaction.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : EstimationSessionAction (estimationsessionaction)
// Endpoint 	        : EstimationSessionAction (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 08/03/2023 at 18:42:23
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//EstimationSessionAction defines the datamodel for the EstimationSessionAction object
type EstimationSessionAction struct {
	// Dynamically generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	ID       string
	EstimationSession       string
	Action       string
	Notes       string
	//
	// Field Properties
	//
	ID_props FieldProperties
	EstimationSession_props FieldProperties
	Action_props FieldProperties
	Notes_props FieldProperties
	//
 	// Any lookups will be added below
	//
	EstimationSession_lookup []Lookup_Item
	Action_lookup []Lookup_Item
	}

const (
	// Dynamically generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// Page Infrastructure
	EstimationSessionAction_Name      = "EstimationSessionAction"
	EstimationSessionAction_Title       = "Estimation Session Actions"
	EstimationSessionAction_SQLTable    = "n/a"
	EstimationSessionAction_SQLSearchID = "n/a"
	EstimationSessionAction_QueryString = "ID"
	///
	/// Template Path Defintions
	///
	EstimationSessionAction_Template     = "EstimationSessionAction"
	EstimationSessionAction_TemplateList = "/EstimationSessionAction/EstimationSessionActionList"
	EstimationSessionAction_TemplateView = "/EstimationSessionAction/EstimationSessionActionView"
	EstimationSessionAction_TemplateEdit = "/EstimationSessionAction/EstimationSessionActionEdit"
	EstimationSessionAction_TemplateNew  = "/EstimationSessionAction/EstimationSessionActionNew"
	///
	/// URI Handler Paths
	///
	EstimationSessionAction_Path       = "/API/EstimationSessionAction/"
	EstimationSessionAction_PathList   = "/EstimationSessionActionList/"
	EstimationSessionAction_PathView   = "/EstimationSessionActionView/"
	EstimationSessionAction_PathEdit   = "/EstimationSessionActionEdit/"
	EstimationSessionAction_PathNew    = "/EstimationSessionActionNew/"
	EstimationSessionAction_PathSave   = "/EstimationSessionActionSave/"
	EstimationSessionAction_PathDelete = "/EstimationSessionActionDelete/"
	// Redirects - On Server Side Error
	EstimationSessionAction_PathEditException   = "/EstimationSessionActionEditException/"
	EstimationSessionAction_PathNewException    = "/EstimationSessionActionNewException/"
	//
	//EstimationSessionAction_Redirect provides a page to return to aftern an action
	EstimationSessionAction_Redirect = EstimationSessionAction_PathView
	//
	//
	// SQL Field Definitions
	//
	EstimationSessionAction_ID_sql   = "ID" // ID is a String
	EstimationSessionAction_EstimationSession_sql   = "EstimationSession" // EstimationSession is a String
	EstimationSessionAction_Action_sql   = "Action" // Action is a String
	EstimationSessionAction_Notes_sql   = "Notes" // Notes is a String
///
	/// Application Field Definitions
	///
	EstimationSessionAction_ID_scrn   = "ID" // ID is a String
	EstimationSessionAction_EstimationSession_scrn   = "EstimationSession" // EstimationSession is a String
	EstimationSessionAction_Action_scrn   = "Action" // Action is a String
	EstimationSessionAction_Notes_scrn   = "Notes" // Notes is a String
///
)

//EstimationSessionAction_PageList provides the information for the template for a list of EstimationSessionActions
type EstimationSessionAction_PageList struct {
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
	ItemList  		 []EstimationSessionAction
}

//EstimationSessionAction_Page provides the information for the template for an individual EstimationSessionAction
type EstimationSessionAction_Page struct {
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
	ID         string
	EstimationSession         string
	Action         string
	Notes         string
	/// Field Properties
	ID_props     FieldProperties
	EstimationSession_props     FieldProperties
	Action_props     FieldProperties
	Notes_props     FieldProperties
	/// Lookups
	EstimationSession_lookup    []Lookup_Item
	Action_lookup    []Lookup_Item
	
	/// END OF DEFINITIONS
	///
}