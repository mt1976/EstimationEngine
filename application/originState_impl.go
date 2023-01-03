package application

// ----------------------------------------------------------------
// Automatically generated  "/application/originstate.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : OriginState (originstate)
// Endpoint 	        : OriginState (OriginStateID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 03/01/2023 at 10:31:00
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"net/http"

	core "github.com/mt1976/ebEstimates/core"
	dao "github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// OriginState_Publish annouces the endpoints available for this object
// OriginState_Publish - Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
func OriginState_Publish_Impl(mux http.ServeMux) {
	// START
	// Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
	//

	mux.HandleFunc(dm.OriginState_PathStore, OriginState_HandlerStore)
	logs.Publish("Implementation", dm.OriginState_Title)
	//
	// Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}

// OriginState_HandlerSave is the handler used process the saving of an OriginState
// It is called from the Edit and New pages
// OriginState_HandlerSave  - Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
func OriginState_HandlerStore(w http.ResponseWriter, r *http.Request) {
	// START
	// Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	// Mandatory Security Validation
	//
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path + r.FormValue("OriginStateID"))

	item := originstate_DataFromRequest(r)

	item.Notify = getCheckBoxValue(dm.OriginState_Notify_scrn, r)
	item.IsLocked = getCheckBoxValue(dm.OriginState_IsLocked_scrn, r)

	//spew.Dump(item)

	dao.OriginState_Store(item, r)
	http.Redirect(w, r, dm.OriginState_Redirect, http.StatusFound)
	//
	// Auto generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
}
