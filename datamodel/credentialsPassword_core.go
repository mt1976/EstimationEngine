package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/credentialspassword.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CredentialsPassword (credentialspassword)
// Endpoint 	        : CredentialsPassword (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 08/03/2023 at 18:42:23
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//CredentialsPassword defines the datamodel for the CredentialsPassword object
type CredentialsPassword struct {
	// Dynamically generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// Field Definitions
	//
	ID       string
	UserName       string
	PasswordOld       string
	PasswordNew       string
	PasswordConfirm       string
	//
	// Field Properties
	//
	ID_props FieldProperties
	UserName_props FieldProperties
	PasswordOld_props FieldProperties
	PasswordNew_props FieldProperties
	PasswordConfirm_props FieldProperties
	//
 	// Any lookups will be added below
	//
	}

const (
	// Dynamically generated 08/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// Page Infrastructure
	CredentialsPassword_Name      = "CredentialsPassword"
	CredentialsPassword_Title       = "Credentials Password"
	CredentialsPassword_SQLTable    = "inboxMessages"
	CredentialsPassword_SQLSearchID = "MailId"
	CredentialsPassword_QueryString = "ID"
	///
	/// Template Path Defintions
	///
	CredentialsPassword_Template     = "CredentialsPassword"
	CredentialsPassword_TemplateList = "/CredentialsPassword/CredentialsPasswordList"
	CredentialsPassword_TemplateView = "/CredentialsPassword/CredentialsPasswordView"
	CredentialsPassword_TemplateEdit = "/CredentialsPassword/CredentialsPasswordEdit"
	CredentialsPassword_TemplateNew  = "/CredentialsPassword/CredentialsPasswordNew"
	///
	/// URI Handler Paths
	///
	CredentialsPassword_Path       = "/API/CredentialsPassword/"
	CredentialsPassword_PathList   = "/CredentialsPasswordList/"
	CredentialsPassword_PathView   = "/CredentialsPasswordView/"
	CredentialsPassword_PathEdit   = "/CredentialsPasswordEdit/"
	CredentialsPassword_PathNew    = "/CredentialsPasswordNew/"
	CredentialsPassword_PathSave   = "/CredentialsPasswordSave/"
	CredentialsPassword_PathDelete = "/CredentialsPasswordDelete/"
	// Redirects - On Server Side Error
	CredentialsPassword_PathEditException   = "/CredentialsPasswordEditException/"
	CredentialsPassword_PathNewException    = "/CredentialsPasswordNewException/"
	//
	//CredentialsPassword_Redirect provides a page to return to aftern an action
	CredentialsPassword_Redirect = CredentialsPassword_PathView
	//
	//
	// SQL Field Definitions
	//
	CredentialsPassword_ID_sql   = "ID" // ID is a String
	CredentialsPassword_UserName_sql   = "UserName" // UserName is a String
	CredentialsPassword_PasswordOld_sql   = "PasswordOld" // PasswordOld is a String
	CredentialsPassword_PasswordNew_sql   = "PasswordNew" // PasswordNew is a String
	CredentialsPassword_PasswordConfirm_sql   = "PasswordConfirm" // PasswordConfirm is a String
///
	/// Application Field Definitions
	///
	CredentialsPassword_ID_scrn   = "ID" // ID is a String
	CredentialsPassword_UserName_scrn   = "UserName" // UserName is a String
	CredentialsPassword_PasswordOld_scrn   = "PasswordOld" // PasswordOld is a String
	CredentialsPassword_PasswordNew_scrn   = "PasswordNew" // PasswordNew is a String
	CredentialsPassword_PasswordConfirm_scrn   = "PasswordConfirm" // PasswordConfirm is a String
///
)

//CredentialsPassword_PageList provides the information for the template for a list of CredentialsPasswords
type CredentialsPassword_PageList struct {
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
	ItemList  		 []CredentialsPassword
}

//CredentialsPassword_Page provides the information for the template for an individual CredentialsPassword
type CredentialsPassword_Page struct {
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
	UserName         string
	PasswordOld         string
	PasswordNew         string
	PasswordConfirm         string
	/// Field Properties
	ID_props     FieldProperties
	UserName_props     FieldProperties
	PasswordOld_props     FieldProperties
	PasswordNew_props     FieldProperties
	PasswordConfirm_props     FieldProperties
	/// Lookups
	
	/// END OF DEFINITIONS
	///
}