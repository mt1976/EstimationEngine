package core

import (
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/hako/durafmt"
	logs "github.com/mt1976/ebEstimates/logs"
)

func Uptime() time.Duration {
	return time.Since(startTime)
}

func Log_uptime() {
	logs.System("App Uptime : " + ApplicationUpTime())
}

func ApplicationUpTime() string {
	return durafmt.Parse(Uptime()).String()
}

func Initialise() {
	//logs.Message("Initialisation", "Vroom")
	//logs.Information("Initialisation", "Vroom "+Bike)

	SessionToken = ""
	UUID = "authorAdjust"
	SecurityViolation = ""

	SystemHostname, _ = os.Hostname()

	PWD, _ = os.Getwd()

	PreInitialise()

	//ApplicationProperties = getPropertiesFromFile(APPCONFIG)
	//ApplicationPropertiesDB = getPropertiesFromFile(DATASTORECONFIG)
	//InstanceProperties = getPropertiesFromFile(INSTANCECONFIG)
	//ApplicationStubLists = getPropertiesFromFile(APPSTUBLISTS)

	//var ApplicationProperties Congiguration
	ApplicationProperties.Load(APPCONFIG)
	//var ApplicationPropertiesDB Congiguration
	ApplicationPropertiesDB.Load(DATASTORECONFIG)
	//var InstanceProperties Congiguration
	InstanceProperties.Load(INSTANCECONFIG)
	//ßvar ApplicationStubLists Congiguration
	ApplicationStubLists.Load(APPSTUBLISTS)

	//var API API
	//var API Cache
	API.Generate()
	//API.ListRoutes()
	//API = API.AddRoute("GET", "/api/author", "Get Author", "author", "author")
	//API = API.AddRoute("GET", "/api/author/{id}", "Get Author", "author", "author")
	//API.ListRoutes()
	// fmt.Printf("ApplicationProperties: %v\n", ApplicationProperties)
	// fmt.Printf("ApplicationPropertiesDB: %v\n", ApplicationPropertiesDB)
	// fmt.Printf("InstanceProperties: %v\n", InstanceProperties)
	// fmt.Printf("ApplicationStubLists: %v\n", ApplicationStubLists)
	var appCache Cache
	ApplicationCache = appCache.Start()
	//logs.Information("Initialisation", ApplicationCache.Stat())
	//fmt.Printf("ApplicationCache: %v\n", ApplicationCache)

	IsChildInstance = false
	if len(GetDatabaseProperty("instance")) != 0 {
		//	logs.Information("Initialisation", fmt.Sprintf("Child Instance Detected %v", len(ApplicationPropertiesDB["instance"])))
		IsChildInstance = true
	}
	//logs.Information("Initialisation", fmt.Sprintf("IsChildInstance: %v", IsChildInstance))
	//

	logs.Information("Connecting to application databases...", "")

	ApplicationDB, _ = Database_Connect(ApplicationPropertiesDB) //log.Printf("Initialisation", ApplicationDB.Stats())

	logs.Success("Connection established")

	//

	SessionManager = scs.New()
	//fmt.Printf("ApplicationSessionLife(): %v\n", ApplicationSessionLife())
	life, err := time.ParseDuration(ApplicationSessionLife())
	if err != nil {
		logs.Fatal("No Session Life Found", err)
	}
	SessionManager.Lifetime = life
	SessionManager.IdleTimeout = 20 * time.Minute
	SessionManager.Cookie.Name = GetCookieIdentity()
	SessionManager.Cookie.HttpOnly = true
	SessionManager.Cookie.Persist = true
	//SessionManager.Cookie.SameSite = http.SameSiteStrictMode
	SessionManager.Cookie.Secure = false
	//spew.Dump(SessionManager)
	//fmt.Printf("SessionManager: %v\n", SessionManager)

	Emailer = Email_init()

	//logs.Information("Initialisation", "Vroooom Vrooooom! "+Bike+Bike)
}
