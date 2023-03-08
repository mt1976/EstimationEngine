package jobs

// ----------------------------------------------------------------
// Automatically generated  "/jobs/EstimationSession.go"
// ----------------------------------------------------------------
// Package              : jobs
// Object 			    : EstimationSession (EstimationSession)
// Endpoint 	        : EstimationSession (MessageID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/12/2021 at 09:31:50
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"fmt"
	"time"

	core "github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	"github.com/mt1976/ebEstimates/logs"
)

// EstimationSessionJob defines the job properties, name, period etc..
func EstimationSession_Job_impl(j dm.JobDefinition) dm.JobDefinition {
	// ----------------------------------------------------------------
	// Job Definition
	// ----------------------------------------------------------------
	j.ID = "EstimationSession"
	j.Name = "EstimationSession"
	j.Period = "15 00 * * *"
	j.Description = "EstimationSession processing"
	j.Type = core.General
	return j
}

// EstimationSession_Run initiates and runs the job as per the period.
func EstimationSession_Run_impl() (string, error) {

	objectName := dao.Translate("ObjectName", "EstimationSession")

	//GEt a list of unprocessed messages from EstimationSession
	expiredState, _ := dao.Data_GetString(objectName, "Quote_Expired_State", dm.Data_Category_State)
	issuedState, _ := dao.Data_GetString(objectName, "Quote_Issued_State", dm.Data_Category_State)
	expiryPeriod, _ := dao.Data_GetInt(objectName, "Quote_Expiry_Notification_Period", dm.Data_Category_Setting)
	logs.Information("Expired State: ", expiredState)
	logs.Information("Issued State: ", issuedState)
	logs.Information("Expiry Period: ", fmt.Sprintf("%d", expiryPeriod))

	//Get a list of issued sessions
	noissued, isList, _ := dao.EstimationSession_ListActive_ByState(issuedState)
	logs.Information("No of Estimates: ", fmt.Sprintf("%d", noissued))

	//Get Today's date
	today := time.Now().Format(core.DATEFORMAT)
	noUpdated := 0
	noIssuedDateAdded := 0
	noExpiryDateAdded := 0
	noExpired := 0
	//Iterate through the list
	for _, v := range isList {
		upd := false
		//Get the session
		if v.EstimationStateID == issuedState {
			if v.IssueDate == "" {
				logs.Warning("Estimate has no issue date : " + v.Name)
				v.EstimationStateID = expiredState
				v.IssueDate = today
				v.Activity = core.AddActivity(v.Activity, "Estimate Issued with no Issue Date - Adding "+today+" as Issue Date")
				noIssuedDateAdded++
				upd = true
			}
			if v.ExpiryDate == "" {
				logs.Warning("Estimate Issued but has no expiry date : " + v.Name)
				// Expiry Date = Issue Date + 30 days
				thisIssueDate, _ := time.Parse(core.DATEFORMAT, v.IssueDate)
				thisExpiryDate := thisIssueDate.AddDate(0, 0, expiryPeriod)
				v.ExpiryDate = thisExpiryDate.Format(core.DATEFORMAT)
				v.Activity = core.AddActivity(v.Activity, "Estimate Issued with no Expiry Date - Adding "+v.ExpiryDate+" as Expiry Date ")
				MSG_TXT := "%s + %d days = %s"
				MSG_TXT = fmt.Sprintf(MSG_TXT, v.IssueDate, expiryPeriod, v.ExpiryDate)
				logs.Information("Expiry Date: ", MSG_TXT)
				v.Activity = core.AddActivity(v.Activity, "Estimate Issued on "+v.IssueDate)
				v.Activity = core.AddActivity(v.Activity, MSG_TXT)
				if thisExpiryDate.Before(time.Now()) {
					v.EstimationStateID = expiredState
					//Add a note to the session
					v.Activity = core.AddActivity(v.Activity, "Estimate expired on "+v.ExpiryDate)
				}
				upd = true
				noExpiryDateAdded++
			}
			thisExpiryDate, _ := time.Parse(core.DATEFORMAT, v.ExpiryDate)
			logs.Information("Issue Date: ", v.IssueDate)
			logs.Information("Expiry Date: ", v.ExpiryDate)
			if thisExpiryDate.Before(time.Now()) {
				logs.Warning("Estimate has expired : " + v.Name)
				v.EstimationStateID = expiredState
				v.Activity = core.AddActivity(v.Activity, "Estimate expired on "+v.ExpiryDate)
				upd = true
				noExpired++
			}
			if upd {
				//spew.Dump(v)
				dao.EstimationSession_StoreSystem(v)
				noUpdated++
			}
		}
	}

	MSG_TEXT := "Estimations Issued: %d, Updated: %d, Issued Date Added: %d, Expiry Date Added: %d, Expired: %d"
	dao.Translate("JobMessage", MSG_TEXT)
	MSG_TEXT = fmt.Sprintf(MSG_TEXT, noissued, noUpdated, noIssuedDateAdded, noExpiryDateAdded, noExpired)
	logs.Information("Result", MSG_TEXT)

	return MSG_TEXT, nil
}
