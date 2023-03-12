package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/origin.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Origin (origin)
// Endpoint 	        : Origin (OriginID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 27/11/2022 at 20:46:14
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"fmt"

	"github.com/bojanz/currency"
	core "github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/das"
	logs "github.com/mt1976/ebEstimates/logs"

	dm "github.com/mt1976/ebEstimates/datamodel"
)

// Origin_GetByCode() returns a single Origin record
func Origin_GetByCode(id string) (int, dm.Origin, error) {
	//logs.Warning("Origin_GetByCode " + id + " " + core.GetSQLSchema(core.ApplicationPropertiesDB))
	tsql := das.SELECTALL + das.FROM + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Origin_SQLTable)
	tsql = tsql + " WHERE " + dm.Origin_Code_sql + "='" + id + "'"
	//logs.Warning("Origin GetByCode " + tsql)
	no, _, originItem, err := origin_Fetch(tsql)
	//logs.Warning("Query Done")
	if no == 0 {
		return no, dm.Origin{}, nil
	}
	if no > 1 {
		return no, dm.Origin{}, nil
	}
	//logs.Warning("Origin_GetByCode " + id + " " + originItem.FullName)
	// START
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return no, originItem, err
}

// Origin_GetList() returns a list of all Origin records
func Origin_GetActiveList() (int, []dm.Origin, error) {

	tsql := das.SELECTALL + das.FROM + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Origin_SQLTable)
	tsql = tsql + " " + das.WHERE + das.ISNULL(dm.Project_SYSDeleted_sql)
	tsql = tsql + " " + das.SORTBY + dm.Origin_Code_sql
	count, originList, _, _ := origin_Fetch(tsql)

	for i := 0; i < count; i++ {
		// START
		// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
		//
		symbol := "£"
		ok := false
		if originList[i].Currency != "" {
			locale := currency.NewLocale("en")
			symbol, ok = currency.GetSymbol(originList[i].Currency, locale)
			if !ok {
				symbol = originList[i].Currency
			}
		}
		originList[i].Currency = symbol
		//
		// Dynamically generated 27/11/2022 by matttownsend (Matt Townsend) on silicon.local
		// END
	}

	return count, originList, nil
}

// ----------------------------------------------------------------
// Origin_RateOnLoad_validate_impl provides validation/actions for RateOnLoad
func Origin_RateOnLoad_validate_impl(iAction string, iId string, iValue string, iRec dm.Origin, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Origin", dm.Origin_RateOnLoad_scrn, VAL+"-"+iAction, iId)

	return iValue, fP
}

func sendStateChangeEmail(who string, rec dm.Origin, newState string) dm.Origin {

	logs.Email("Origin", "Sending email to "+who+" about state change to "+newState)

	_, lk, _ := OriginState_GetByCode(newState)

	MSG_SUBJECT := Translate("Notification", "%s state has been updated")
	MSG_SUBJECT = fmt.Sprintf(MSG_SUBJECT, rec.FullName)
	MSG_BODY := Translate("Notification", "Origin %s (%s) state has been changed to %s (%s)")
	MSG_BODY = fmt.Sprintf(MSG_BODY, rec.FullName, rec.Code, lk.Name, rec.StateID)

	//fmt.Printf("lk: %v\n", lk)
	if lk.Notify == core.TRUE {

		Resource_SendMail(who, MSG_SUBJECT, MSG_BODY)
	} else {
		logs.Email("Origin", "No email sent to "+who+" about state change to "+newState)
	}
	rec.SYSActivity = core.AddActivity_ForUser(rec.SYSActivity, MSG_BODY, origin_UpdateBy(rec))

	return rec
}

func sendRateChangeEmail(who string, rec dm.Origin, newRate string) dm.Origin {

	logs.Email("Origin", "Sending email to "+who+" about rate change to "+rec.FullName)
	state := rec.StateID

	_, lk, _ := OriginState_GetByCode(state)

	MSG_SUBJECT := Translate("Notification", "%s rate has been updated")
	MSG_SUBJECT = fmt.Sprintf(MSG_SUBJECT, rec.FullName)

	MSG_BODY := Translate("Notification", "The rate for %s (%s) has been changed to from %s%s to %s%s")
	MSG_BODY = fmt.Sprintf(MSG_BODY, rec.FullName, rec.Code, rec.Currency, core.Financial_FormatAmount(rec.RateOnLoad, rec.Currency), rec.Currency, core.Financial_FormatAmount(newRate, rec.Currency))

	if lk.Notify == core.TRUE {

		Resource_SendMail(who, MSG_SUBJECT, MSG_BODY)
	} else {
		logs.Email("Origin", "No email sent to "+who+" about rate change to "+rec.FullName)
	}

	rec.SYSActivity = core.AddActivity_ForUser(rec.SYSActivity, MSG_BODY, origin_UpdateBy(rec))

	return rec
}

func origin_UpdateBy(rec dm.Origin) string {
	who := rec.SYSUpdatedBy
	if len(who) == 0 {
		who = rec.SYSCreatedBy
	}
	if len(who) == 0 {
		who = Audit_Host()
	}
	return who
}
