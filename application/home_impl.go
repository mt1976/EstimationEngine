package application

import (
	"net/http"
	"time"

	core "github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

type homePage struct {
	UserMenu        dm.AppMenuItem
	UserNavi        string
	UserRole        string
	Title           string
	PageTitle       string
	AppReleaseID    string
	AppReleaseLevel string
	AppReleaseNo    string

	SQLServer          string
	SQLDB              string
	SQLSchema          string
	SQLParentSchema    string
	UserName           string
	UserKnowAs         string
	AppServerDate      string
	AppServerName      string
	DealImportInPath   string
	StaticDataInPath   string
	FundsCheckInPath   string
	RatesDataInPath    string
	DealImportOutPath  string
	StaticDataOutPath  string
	FundsCheckOutPath  string
	RatesDataOutPath   string
	InstanceState      string
	AppSQLServer       string
	AppSQLDatabase     string
	AppSQLSchema       string
	AppSQLParentSchema string
	DateSyncIssue      string
	SessionInfo        dm.SessionInfo
	NoOrigins          int
	OriginList         []dm.Origin
}

func Home_Publish_Impl(mux http.ServeMux) {
	mux.HandleFunc("/home", Home_HandlerView)
	logs.Publish("Core", "Home"+" Impl")

}

// Home_HandlerView
func Home_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	//fmt.Printf("Home_HandlerView\n")
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	//fmt.Printf("Home_HandlerView - Session_Validate\n")
	// Code Continues Below

	//	log.Println("IN HOMEPAGE")

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)
	//tmpHostname, _ := os.Hostname()

	homePage := homePage{
		UserMenu:        UserMenu_Get(r),
		UserRole:        Session_GetUserRole(r),
		UserNavi:        "NOT USED",
		Title:           core.ApplicationName(),
		PageTitle:       PageTitle("Home", ""),
		AppReleaseID:    core.ReleaseIdentity(),
		AppReleaseLevel: core.ReleaseLevel(),
		AppReleaseNo:    core.ReleaseNumber(),

		UserName:      Session_GetUserName(r),
		UserKnowAs:    Session_GetUserKnownAs(r),
		AppServerDate: time.Now().Format(core.DATEFORMAT),
		AppServerName: core.ApplicationHostname(),

		AppSQLServer:       core.ApplicationSQLServer(),
		AppSQLSchema:       core.ApplicationSQLSchema(),
		AppSQLParentSchema: core.ApplicationSQLSchemaParent(),
	}

	//spew.Dump(homePage.UserMenu)

	homePage.InstanceState = "Primary System"
	homePage.AppSQLDatabase = core.ApplicationSQLDatabase()
	homePage.SessionInfo, _ = Session_GetSessionInfo(r)
	if core.IsChildInstance {
		homePage.InstanceState = "Child System"
		//	homePage.AppSQLDatabase = ApplicationSQLDatabase() + "-" + core.GetDatabaseProperty("instance")
	}

	homePage.NoOrigins, homePage.OriginList, _ = dao.Origin_GetActiveList()

	ExecuteTemplate("Impl_Home", w, r, homePage)

}
