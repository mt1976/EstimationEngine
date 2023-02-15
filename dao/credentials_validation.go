package dao

import (
	"fmt"
	"os"
	"time"

	core "github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/credentials.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : Credentials (credentials)
// Endpoint 	        : Credentials (Id)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/02/2023 at 10:31:01
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------
// The following functions should be created in credentials_impl.go
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Credentials_ObjectValidation_impl provides Record/Object level validation for Credentials
func Credentials_ObjectValidation_impl(iAction string, iId string, iRec dm.Credentials) (dm.Credentials, string, error) {
	logs.Callout("Credentials", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:
		switch iRec.State {
		case "ACTV":
			iRec.Expiry = getExpiryDate()
			iRec.Issued = time.Now().Format(core.DATETIMEFORMATUSER)
			host, _ := os.Hostname()
			MSG_BODY := "You have been authorised to use %s until %s"
			Translate("CredentialsMessage", MSG_BODY)
			MSG_BODY = fmt.Sprintf(MSG_BODY, core.ApplicationName(), iRec.Expiry)
			MSG_SUBJECT := "Welcome to %s"
			Translate("CredentialsMessage", MSG_SUBJECT)
			MSG_SUBJECT = fmt.Sprintf(MSG_SUBJECT, core.ApplicationName())
			_ = Inbox_SendMail(iRec.Username, MSG_SUBJECT, MSG_BODY, host, host)
			core.SendEmail(iRec.Username, iRec.Firstname, MSG_SUBJECT, MSG_BODY)

		case "LOCK":
			iRec.Expiry = ""
			iRec.Issued = ""
		case "ULCK":
			iRec.Expiry = getExpiryDate()
			iRec.Issued = time.Now().Format(core.DATETIMEFORMATUSER)
		case "REJE":
			iRec.Expiry = ""
			iRec.Issued = ""
			// Send an email telling the use to go bugger off???
			MSG_SUBJECT := "Your request to access %s has not been approved"
			MSG_BODY := MSG_SUBJECT + ". Please Contact the System Admin if you feel that this is not justified"
			Translate("CredentialsMessage", MSG_SUBJECT)
			Translate("CredentialsMessage", MSG_BODY)
			MSG_SUBJECT = fmt.Sprintf(MSG_SUBJECT, core.ApplicationName())
			MSG_BODY = fmt.Sprintf(MSG_BODY, core.ApplicationName())
			core.SendEmail(iRec.Username, iRec.Firstname, MSG_SUBJECT, MSG_BODY)
		default:
			logs.Warning("Credentials_ObjectValidation_impl" + " - Invalid State")
		}
	case GET:

	default:
		logs.Warning("Tmpl_TDate_impl" + " - Invalid Action")
	}
	return iRec, "", nil
}

//
// ----------------------------------------------------------------
// These are the the default implementations, which do nothing

// ----------------------------------------------------------------
// BEGIN Credentials_State
// BEGIN Credentials_State
// BEGIN Credentials_State
// ----------------------------------------------------------------
// Credentials_State_OnStore_impl provides the implementation for the callout
func Credentials_State_OnStore_impl(fieldval string, rec dm.Credentials, usr string) (string, error) {
	logs.Callout("Credentials", dm.Credentials_State_scrn, PUT, rec.Id)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Credentials_State_OnFetch_impl provides the implementation for the callout
func Credentials_State_OnFetch_impl(rec dm.Credentials) string {
	logs.Callout("Credentials", dm.Credentials_State_scrn, GET, rec.Id)
	return rec.State
}

// ----------------------------------------------------------------
// Credentials_State_validate_impl provides validation/actions for State
func Credentials_State_validate_impl(iAction string, iId string, iValue string, iRec dm.Credentials, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Credentials", dm.Credentials_State_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   Credentials_State
// END   Credentials_State
// END   Credentials_State
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------
