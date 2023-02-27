package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	_ "github.com/denisenkom/go-mssqldb"

	application "github.com/mt1976/ebEstimates/application"
	core "github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/dao"
	"github.com/mt1976/ebEstimates/jobs"
	logs "github.com/mt1976/ebEstimates/logs"
	upgrader "github.com/mt1976/ebEstimates/upgrade"
)

var err error

func init() {
	err = nil
	logs.Clear()
	logs.Break()
	logs.Header("Estimation Engine")
	logs.Break()
	logs.Information("Initialising...", "")
	core.Initialise()
	logs.Success("Initialised")
}

func main() {
	if err == nil {
		upgradeDatabase()
	}
	if err == nil {
		startScheduler()
	}

	mux := publishGUIEndpoints()

	if err == nil {
		//	rebuildIndex()
	}
	if err == nil {
		publishAPIEndpoints()
	}

	if err == nil {
		startMonitors()
	}

	if err == nil {
		displayAppInfo()
	}

	logs.Header("READY STEADY GO!!!")
	//logs.Information("Initialisation", "Vrooom, Vrooooom, Vroooooooo..."+logs.Character_Bike+logs.Character_Bike+logs.Character_Bike+logs.Character_Bike)

	if err == nil {
		alertSystemStarted()
	}
	//jobs.EstimationSession_Run()
	//jobs.Ssloader_Run()

	httpProtocol := core.ApplicationHTTPProtocol()
	logs.URI(httpProtocol + "://localhost:" + core.ApplicationHTTPPort())

	httpPort := ":" + core.ApplicationHTTPPort()

	if core.ApplicationEnvironment() != "PROD" {

		http.ListenAndServe(httpPort, core.SessionManager.LoadAndSave(mux))

	} else {

		certPath := core.ApplicationCertificatePath()
		certName := core.ApplicationCertificateName()

		pwd, _ := os.Getwd()
		certLocation := pwd + certPath + certName
		certKey := certLocation + ".key"
		certCrt := certLocation + ".crt"

		logs.Information("Certificate Path", certPath)
		logs.Information("Certificate Name", certName)
		logs.Information("Certificate Location", certLocation)
		logs.Information("Certificate Key", certKey)
		logs.Information("Certificate Crt", certCrt)

		log.Fatal(http.ListenAndServeTLS(httpPort, certCrt, certKey, core.SessionManager.LoadAndSave(mux)))
	}
	logs.Break()

}

func alertSystemStarted() {
	MSG_BODY := "System Online <br><br> %s Started %s at %s on %s"
	MSG_BODY = dao.Translate("Email", MSG_BODY)
	MSG_BODY = fmt.Sprintf(MSG_BODY, core.ApplicationName(), time.Now().Format("15:04:05"), time.Now().Format("02/01/2006"), core.ApplicationHostname())
	core.SendEmail(core.GetApplicationProperty("admin"), "Admin", "System Online - "+core.ApplicationName(), MSG_BODY)
}

func startMonitors() {
	logs.Break()
	logs.Header("Start Watchers")
	logs.Break()

	logs.Success("Watchers Started")
}

func publishAPIEndpoints() {
	logs.Break()
	logs.Header("Publish API")
	logs.Break()
	core.Catalog_List()

	logs.Success("API Published")
}

func rebuildIndex() {
	logs.Break()
	logs.Header("Rebuild Index")
	logs.Break()
	e2 := dao.Indexer_Rebuild()
	if e2 != nil {
		logs.Error("Indexer", e2)
	}
	logs.Success("Index Rebuilt")
}

func publishGUIEndpoints() *http.ServeMux {
	logs.Break()
	logs.Header("Publish Endpoints")
	logs.Break()

	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	mux.Handle("/data/images/avatars/", http.StripPrefix("/data/images/avatars/", http.FileServer(http.Dir("data/images/avatars"))))

	core.LoginLogout_Publish_Impl(*mux)
	application.Resources_Publish_Impl(*mux)
	application.Home_Publish_Impl(*mux)
	application.Session_Publish_Impl(*mux)
	application.Translation_Publish(*mux)

	application.Configuration_Publish_Impl(*mux)

	application.Credentials_Publish(*mux)
	application.CredentialsAction_Publish(*mux)
	application.CredentialsPassword_Publish(*mux)
	application.CredentialsPassword_Publish_Impl(*mux)

	application.Message_Publish(*mux)

	application.Schedule_Publish(*mux)

	jobs.Schedule_PublishImpl(*mux)

	application.Session_Publish(*mux)

	application.SQLInjection_Publish(*mux)

	application.Catalog_Publish_impl(*mux)
	application.Inbox_Publish(*mux)
	application.Inbox_Publish_Impl(*mux)
	application.UserRole_Publish(*mux)

	application.Resource_Publish(*mux)
	application.Confidence_Publish(*mux)
	application.DocType_Publish(*mux)
	application.EstimationSession_Publish(*mux)
	application.EstimationSession_Publish_Impl(*mux)
	application.EstimationSessionAction_Publish(*mux)
	application.EstimationSessionAction_Publish_Impl(*mux)
	application.EstimationState_Publish(*mux)
	application.ExternalMessage_Publish(*mux)
	application.Feature_Publish(*mux)
	application.Feature_Publish_Impl(*mux)
	application.FeatureNew_Publish(*mux)
	application.FeatureNew_Publish_Impl(*mux)
	application.Origin_Publish(*mux)
	application.Origin_Publish_Impl(*mux)

	application.OriginState_Publish(*mux)
	application.OriginState_Publish_Impl(*mux)
	application.Profile_Publish(*mux)
	application.Project_Publish(*mux)
	application.Project_Publish_Impl(*mux)
	application.ProjectState_Publish(*mux)
	application.ProjectAction_Publish_Impl(*mux)

	application.Data_Publish(*mux)

	application.Index_Publish(*mux)
	logs.Success("Endpoints Published")

	return mux
}

func startScheduler() {
	logs.Break()
	logs.Header("Scheduling Jobs")
	logs.Break()
	jobs.Start()
	logs.Success("Jobs Scheduled")
}

func upgradeDatabase() {
	logs.Break()
	logs.Header("Database Upgrade")
	logs.Break()
	message, e := upgrader.Upgrade()
	if e != nil {
		logs.Error("Upgrade", e)
	}
	logs.Success(message)
	logs.Break()
}

func displayAppInfo() {

	logs.Break()
	logs.Header("Application Information")
	logs.Break()

	logs.Header("Application")
	if core.ApplicationCompanyName() != "" {
		logs.Information("Company", core.ApplicationCompanyName())
	}
	logs.Information("Name", core.ApplicationName())
	logs.Information("Host Name", core.ApplicationHostname())
	logs.Information("Server Release", core.ReleaseIdentityVerbose())
	logs.Information("Server Date", time.Now().Format(core.DATEFORMATUSER))
	logs.Information("Server Time", time.Now().Format(core.TIMEHMS))
	if !core.IsChildInstance {
		logs.Success("Server Mode: " + "Primary")
	} else {
		logs.Warning("Server Mode: " + "Secondary")
	}
	logs.Information("Licence", core.ApplicationGetLicenseName())
	logs.Information("Lic URL", core.ApplicationGetLicenseLink())
	logs.Header("Runtime")
	logs.Information("GO Version", runtime.Version())
	logs.Information("Operating System", runtime.GOOS+" ("+runtime.GOARCH+")")
	logs.Header("Application Database (MSSQL)")
	logs.Information("Server", core.ApplicationSQLServer())
	logs.Information("Database", core.ApplicationSQLDatabase())
	logs.Information("Schema", core.ApplicationSQLSchema())
	logs.Information("Parent Schema", core.ApplicationSQLSchemaParent())
	logs.Information("DB Model Version", core.DB_Version())

	logs.Header("Sessions")
	logs.Information("Session Life", core.ApplicationSessionLife())

	//core.GetApplicationProperty("admin")
	//core.GetDatabaseProperty("admin")

	logs.Break()
	logs.Header("Contacts")
	logs.Break()
	logs.Information("Admin", core.GetApplicationProperty("admin"))
	logs.Break()

}
