package application

import (
	"net/http"

	core "github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/logs"
)

// SQLInjection_Publish Publishes the endpoint
func SQLInjection_Publish(mux http.ServeMux) {
	logs.Publish("Application", "SQLInjection")
	mux.HandleFunc("/injectSQLViews/", SQLInjection_HandlerRun)

}

// SQLInjectionHander
func SQLInjection_HandlerRun(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	//core.Database_CreateObjects(core.ApplicationDB, core.ApplicationPropertiesDB, "/config/database/views", false)
	http.Redirect(w, r, "/", http.StatusFound)

}
