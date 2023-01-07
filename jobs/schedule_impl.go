package jobs

// ----------------------------------------------------------------
// Automatically generated  "/application/schedule.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Schedule (schedule)
// Endpoint 	        : Schedule (Schedule)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 21/11/2021 at 12:24:49
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"net/http"

	application "github.com/mt1976/ebEstimates/application"
	core "github.com/mt1976/ebEstimates/core"
	dao "github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"

	logs "github.com/mt1976/ebEstimates/logs"
)

// Schedule_Publish annouces the endpoints available for this object
func Schedule_PublishImpl(mux http.ServeMux) http.ServeMux {
	mux.HandleFunc(dm.Schedule_PathDO, Schedule_HandlerDO_Impl)
	//Cannot Delete via GUI
	logs.Publish("Implementation", dm.Schedule_Title)
	return mux
}

func Schedule_HandlerDO_Impl(w http.ResponseWriter, r *http.Request) {

	// Mandatory Security Validation
	if !(application.Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	searchID := core.GetURLparam(r, dm.Schedule_QueryString)

	_, thisJob, _ := dao.Schedule_GetByID(searchID)

	logs.Information("Adhoc Job", "Job: "+thisJob.Name+" Type: "+thisJob.Type)
	switch thisJob.Name {
	case Credentials_Job().Name:
		if thisJob.Type == Credentials_Job().Type {
			go Credentials_Run()
		}

	case Database_Job().Name:
		if thisJob.Type == Database_Job().Type {
			go Database_Run()
		}
	case EstimationSession_Job().Name:
		if thisJob.Type == EstimationSession_Job().Type {
			go EstimationSession_Run()
		}
	case HeartBeat_Job().Name:
		if thisJob.Type == HeartBeat_Job().Type {
			go HeartBeat_Run()
		}
	case Origin_Job().Name:
		if thisJob.Type == Origin_Job().Type {
			go Origin_Run()
		}
	case Project_Job().Name:
		if thisJob.Type == Project_Job().Type {
			go Project_Run()
		}
	case SessionHouseKeeping_Job().Name:
		if thisJob.Type == SessionHouseKeeping_Job().Type {
			go SessionHouseKeeping_Run()
		}
	default:
		//Do nothing
	}

	REDR := dm.Schedule_Redirect
	http.Redirect(w, r, REDR, http.StatusFound)

}
