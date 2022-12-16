package application

import (
	"net/http"
	"runtime"
	"strings"

	core "github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// AppConfigurationPage is cheese
type AppConfigurationPage struct {
	SessionInfo            dm.SessionInfo
	UserMenu               dm.AppMenuItem
	UserRole               string
	UserNavi               string
	Title                  string
	PageTitle              string
	RequestPath            string
	ResponsePath           string
	ProcessedPath          string
	ResponseType           string
	AppServerReleaseID     string
	AppServerReleaseLevel  string
	AppServerReleaseNumber string

	AppDBServer        string
	AppDBPort          string
	AppDBUser          string
	AppDBPassword      string
	AppDBDatabase      string
	AppDBSchema        string
	AppCredentialsLife string
	AppSessionLife     string
	GoVersion          string
	OS                 string
	AppName            string
	CompanyName        string
	DBModel            string
	LicenseName        string
	LicenseLink        string
	AppEnvironment     string
	AppDocker          string
}

func Configuration_Publish_Impl(mux http.ServeMux) {
	mux.HandleFunc("/viewAppConfiguration/", Configuration_HandlerView)
	logs.Publish("Core", "Configuration"+" Impl")
}

func Configuration_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	tmpl := "Configuration/Impl_Configuration_View"
	// Impl - defines this is hand crafted html page
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	//title := core.GetAppName()

	// Get Data Here

	pageAppConfigView := AppConfigurationPage{
		UserMenu:               UserMenu_Get(r),
		UserRole:               Session_GetUserRole(r),
		UserNavi:               "NOT USED",
		Title:                  CardTitle("Application", core.Action_Configure),
		PageTitle:              PageTitle("Application", core.Action_Configure),
		RequestPath:            "OBSOLETE",
		ResponsePath:           "OBSOLETE",
		ProcessedPath:          "OBSOLETE",
		ResponseType:           "OBSOLETE",
		AppServerReleaseID:     core.ReleaseID(),
		AppServerReleaseLevel:  core.ReleaseLevel(),
		AppServerReleaseNumber: core.ReleaseNumber(),

		GoVersion: runtime.Version(),
		OS:        runtime.GOOS,
	}
	pageAppConfigView.SessionInfo, _ = Session_GetSessionInfo(r)
	pageAppConfigView.AppDBServer = core.ApplicationSQLServer()
	pageAppConfigView.AppDBPort = core.GetDatabaseProperty("port")
	pageAppConfigView.AppDBUser = core.GetDatabaseProperty("user")
	pageAppConfigView.AppDBPassword = strings.Repeat("*", len(core.GetDatabaseProperty("password")))
	pageAppConfigView.AppDBDatabase = core.ApplicationSQLDatabase()
	pageAppConfigView.AppDBSchema = core.ApplicationSQLSchema()
	pageAppConfigView.AppCredentialsLife = core.ApplicationCredentialsLife()
	pageAppConfigView.AppSessionLife = core.ApplicationSessionLife()

	//fmt.Printf("pageAppConfigView: %v\n", pageAppConfigView)
	//thisTemplate:= core.GetTemplateID(tmpl,core.GetUserRole(r))

	pageAppConfigView.SessionInfo, _ = Session_GetSessionInfo(r)
	pageAppConfigView.AppName = core.ApplicationName()
	pageAppConfigView.CompanyName = core.ApplicationCompanyName()
	pageAppConfigView.DBModel = core.DB_Version()
	pageAppConfigView.LicenseName = core.ApplicationLicenseName()
	pageAppConfigView.LicenseLink = core.ApplicationLicenseLink()
	pageAppConfigView.AppEnvironment = core.ApplicationEnvironment()
	if core.RunningInDockerContainer() {
		pageAppConfigView.AppDocker = "Yes"
	} else {
		pageAppConfigView.AppDocker = "No"
	}

	ExecuteTemplate(tmpl, w, r, pageAppConfigView)

}
