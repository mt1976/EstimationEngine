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
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 03/01/2023 at 19:18:08
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in credentials_impl.go

// If there are fields below, create the methods in adaptor\credentials_impl.go
// START - GET API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
func Credentials_State_OnStore_impl(fieldval string, rec dm.Credentials, usr string) (string, error) {
	logs.Callout("Credentials", dm.Credentials_State_scrn, PUT, rec.Id)
	return fieldval, nil
}

// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - GET API/Callout
//
// START - PUT API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
func Credentials_State_OnFetch_impl(rec dm.Credentials) string {
	logs.Callout("Credentials", dm.Credentials_State_scrn, GET, rec.Id)
	return rec.State
}

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - PUT API/Callout

// START - Validation API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
//
// Credentials_State_impl provides validation/actions for State
func Credentials_State_impl(iAction string, iId string, iValue string, iRec dm.Credentials, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Credentials", dm.Credentials_State_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout

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
			msg := fmt.Sprintf("You have been authorised to use %s on %s until %s", core.ApplicationName(), host, iRec.Expiry)
			_ = Inbox_SendMail(iRec.Username, "Welcome Message", msg, host, host)
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
		default:
			logs.Warning("Credentials_ObjectValidation_impl" + " - Invalid State")
		}
	case GET:

	default:
		logs.Warning("Tmpl_TDate_impl" + " - Invalid Action")
	}
	return iRec, "", nil
}
