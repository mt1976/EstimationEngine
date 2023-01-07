package main

import (
	"context"
	"fmt"
	"io"
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
)

func main() {

	logs.Break()
	logs.Header("MISSION CONTROL")
	logs.Break()

	logs.Information("Initialising...", "")

	core.Initialise()

	logs.Success("Initialised")
	logs.Break()
	logs.Header("Scheduling Jobs")
	logs.Break()

	jobs.Start()

	logs.Success("Jobs Scheduled")
	logs.Break()
	logs.Header("Publish Endpoints")
	logs.Break()
	// Setup Endpoints
	mux := http.NewServeMux()
	// At least one "mux" handler is required - Dont remove this
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	// At least one "mux" handler is required - Dont remove this

	Main_Publish(*mux)
	core.LoginLogout_Publish_Impl(*mux)
	application.Resources_Publish_Impl(*mux)
	application.Home_Publish_Impl(*mux)
	application.Session_Publish_Impl(*mux)
	application.Translation_Publish(*mux)

	application.Configuration_Publish_Impl(*mux)

	application.Credentials_Publish(*mux)
	application.CredentialsAction_Publish(*mux)

	application.Message_Publish(*mux)

	//application.Translation_Publish(*mux)

	application.Schedule_Publish(*mux)
	// Special Case - Schedule is a special case as it is the only object that has a job that runs it
	jobs.Schedule_PublishImpl(*mux)

	application.Session_Publish(*mux)

	application.SQLInjection_Publish(*mux)

	application.Catalog_Publish_impl(*mux)
	application.Inbox_Publish(*mux)
	application.Inbox_Publish_Impl(*mux)
	application.UserRole_Publish(*mux)
	// User Defined EndPoints
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

	// End of Endpoints
	logs.Break()
	logs.Header("Publish API")
	logs.Break()
	core.Catalog_List()

	logs.Success("Endpoints Published")
	logs.Break()
	logs.Header("Start Watchers")
	logs.Break()
	//go monitors.StaticDataImporter_Watch()

	//	monitors.Start()
	logs.Success("Watchers Started")
	Application_Info()
	logs.Break()
	logs.Header("Contacts")
	logs.Information("Admin", core.GetApplicationProperty("admin"))
	logs.Header("READY STEADY GO!!!")
	logs.Information("Initialisation", "Vrooom, Vrooooom, Vroooooooo..."+logs.Character_Bike+logs.Character_Bike+logs.Character_Bike+logs.Character_Bike)
	logs.Break()

	MSG_BODY := "System Online <br><br> %s Started %s at %s on %s"
	MSG_BODY = dao.Translate("Email", MSG_BODY)
	MSG_BODY = fmt.Sprintf(MSG_BODY, core.ApplicationName(), time.Now().Format("15:04:05"), time.Now().Format("02/01/2006"), core.ApplicationHostname())
	core.SendEmail(core.GetApplicationProperty("admin"), "Admin", "System Online - "+core.ApplicationName(), MSG_BODY)

	//jobs.EstimationSession_Run()

	httpProtocol := core.ApplicationHTTPProtocol()
	logs.URI(httpProtocol + "://localhost:" + core.ApplicationHTTPPort())
	logs.Break()

	httpPort := ":" + core.ApplicationHTTPPort()

	if core.ApplicationEnvironment() == "development" {

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
	core.Log_uptime()
	core.Log_uptime()

}

func Main_Publish(mux http.ServeMux) {

	mux.HandleFunc("/shutdown/", application_HandlerShutdown)
	mux.HandleFunc("/clearQueues/", application_HandlerClearQueues)
	mux.HandleFunc("/put", application_HandlerPUT)
	mux.HandleFunc("/get", application_HandlerGET)
	logs.Publish("Main", "Application")
}

// // TODO: migrage the following three functions to appsupport
func application_HandlerShutdown(w http.ResponseWriter, r *http.Request) {
	//	wctProperties := application.GetProperties(APPCONFIG)

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	//requestID := uuid.New()http.ListenAndServe(":8080", nil)

	log.Println("Servicing :", inUTL)
	//requestMessage := application.BuildRequestMessage(requestID.String(), "SHUTDOWN", line, line, line, wctProperties)

	//fmt.Println("requestMessage", requestMessage)
	//fmt.Println("SEND MESSAGE")
	//application.SendRequest(requestMessage, requestID.String(), core.ApplicationProperties)
	m := http.NewServeMux()

	s := http.Server{Addr: core.ApplicationHTTPPort(), Handler: m}
	s.Shutdown(context.Background())
	//	r.URL.Path = "/viewResponse?uuid=" + requestID.String()
	//	viewResponseHandler(w, r)

}
func application_HandlerClearQueues(w http.ResponseWriter, r *http.Request) {

	//var propertiesFileName = "config/properties.cfg"
	//wctProperties := application.GetProperties(APPCONFIG)
	//	tmpl := "viewResponse"
	inUTL := r.URL.Path
	//requestID := uuid.New()
	log.Println("Servicing :", inUTL)
	//fmt.Println("delivPath", wctProperties["deliverpath"])
	err1 := core.RemoveContents(core.OBSOLETE_MessageQueueDeliveryPath())
	if err1 != nil {
		fmt.Println(err1)
	}
	//fmt.Println("recPath", wctProperties["receivepath"])
	err2 := core.RemoveContents(core.OBSOLETE_MessageQueueRecievePath())
	if err2 != nil {
		fmt.Println(err2)
	}
	//fmt.Println("procPath", wctProperties["processedpath"])
	err3 := core.RemoveContents(core.GetApplicationProperty("processedpath"))
	if err3 != nil {
		fmt.Println(err3)
	}
	application.Home_HandlerView(w, r)
}

func application_HandlerPUT(w http.ResponseWriter, r *http.Request) {
	// Store a new key and value in the session data.
	core.SessionManager.Put(r.Context(), "message", "Hello from a session!")
}

func application_HandlerGET(w http.ResponseWriter, r *http.Request) {
	// Use the GetString helper to retrieve the string value associated with a
	// key. The zero value is returned if the key does not exist.
	msg := core.SessionManager.GetString(r.Context(), "message")
	io.WriteString(w, msg)
}

func Application_Info() {

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
		logs.Information("Server Mode", "Primary System")
	} else {
		logs.Information("Server Mode", "Secondary System")
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

	core.GetApplicationProperty("admin")
	core.GetDatabaseProperty("admin")

}
