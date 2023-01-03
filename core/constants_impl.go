package core

const (
	DATEFORMATPICK          = "20060102T150405"
	DATEFORMAT              = "2006-01-02"
	TIMEFORMAT              = "15:04:05"
	DATEFORMATYMD           = "20060102"
	DATEFORMATUSER          = "02/01/2006"
	DATETIME                = "2006-01-02T15:04:05.999999"
	DFDD                    = "02"
	DFMM                    = "01"
	DFYY                    = "06"
	DFYYYY                  = "2006"
	DFmm                    = "04"
	DFss                    = "05"
	DFhh                    = "15"
	DATETIMEFORMATSQLSERVER = "2006-01-02 15:04:05"
	DATETIMEFORMATUSER      = DATEFORMATUSER + " 15:04:05"
	TIMEHMS                 = "15:04:05"
	DATEMSG                 = "2006-01-02 at 15:04:05"
)
const (
	APPCONFIG       = "application.cfg"
	DATASTORECONFIG = "applicationDB.cfg"
	INSTANCECONFIG  = "instance.cfg"
	APPSTUBLISTS    = "lists.cfg"
)
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)
const (
	SessionRole              = "1891835972"
	SessionNavi              = "6782444386"
	SessionKnowAs            = "0627218437"
	SessionUserName          = "9214790474"
	SessionTokenID           = "1516099114"
	SessionUUID              = "0663644127"
	SessionSecurityViolation = "4097340829"
	SessionAppToken          = "1117429826"
)
const (
	IDSep             = "|"
	ListSeperator     = ","
	ListItemSeperator = "|"
)
const (
	Tick         = "☑️"
	WarningLabel = "⚠️"
	Bike         = "🚴‍♂️"
	Scheduled    = "Scheduled"
	Adhoc        = "AdHoc"
	Monitor      = "Monitor"
	Dispatcher   = "Dispatcher"
	Aquirer      = "Aquirer"
	HouseKeeping = "HouseKeeping"
	General      = "General"
)
const (
	Character_MapTo        = "⇄"
	Character_Job          = "⚙️"
	Character_Heart        = "🫀"
	Character_Poke         = "👉"
	Character_Time         = "🕒"
	Character_Break        = "≫"
	Character_Monitor      = "👁"
	Character_Aquirer      = "🛒"
	Character_Dispatcher   = "📡"
	Character_Gears        = "⚙️"
	Character_HouseKeeping = "🧼"
)
const (
	CatalogType_API = "API"
	CatalogType_GUI = "USER"
	DataSource_App  = "APP"
)
const (
	Action_Edit       = "Edit"
	Action_View       = "View"
	Action_Delete     = "Delete"
	Action_New        = "New"
	Action_Ban        = "Ban"
	Action_Activate   = "Activate"
	Action_Deactivate = "Deactivate"
	Action_List       = "List"
	Action_Save       = "Save"
	Action_Request    = "Request"
	Action_Requests   = Action_Request + "s"
	Action_Process    = "Process"
	Action_Approve    = "Approve"
	Action_Configure  = "Configure"
	Action_Import     = "Import"
)

const (
	// FieldProperties_MsgType is the field name for the message type
	FieldMessage_ERROR    = "is-invalid"
	FieldMessage_NORMAL   = ""
	FieldMessage_POSITIVE = "is-valid"
)

const (
	ContextState_ERROR  = "action-error"
	ContextState_NORMAL = "action-success"
	ContextState_OK     = "action-success"
)

const (
	TRUE  = "true"
	FALSE = "false"
)
