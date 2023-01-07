package jobs

import (
	"fmt"
	"time"

	core "github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

func Origin_Job_impl(j dm.JobDefinition) dm.JobDefinition {
	j.Name = "Origin"
	j.Period = "5 0 * * *"
	j.Description = "Revalidate Origins"
	j.Type = core.HouseKeeping
	j.ID = "Origin"
	return j
}

func Origin_Run_impl() (string, error) {

	message := ""

	/// CONTENT STARTS
	contractExpiryPeriod, err0 := dao.Data_GetInt("Contract", "Expiry")
	if err0 != nil {
		return message, err0
	}
	contractExpiryDate := time.Now().AddDate(0, 0, contractExpiryPeriod)
	// Get the origins list
	_, origins, err := dao.Origin_GetActiveList()
	if err != nil {
		logs.Error("Origin Housekeeping", err)
		return message, err
	}
	// Scan the origins
	noOrigins := 0
	noWarnings := 0
	noDue := 0
	noOverdue := 0
	for _, o := range origins {
		noOrigins++

		// Get Acount Manager from Origin
		_, am, err := dao.Resource_GetByCode(o.AccountManager)
		if err != nil {
			logs.Error("Origin Housekeeping", err)
			return message, err
		}
		// Get Project Manager from Origin
		_, pm, err := dao.Resource_GetByCode(o.ProjectManager)
		if err != nil {
			logs.Error("Origin Housekeeping", err)
			return message, err
		}

		if o.EndDate == "" {
			logs.Information("Origin Housekeeping", "Origin "+o.Code+" has no contract expiry date")
			MSG_TEXT := "%s - %s has no contractual expiry date."
			MSG_BODY := MSG_TEXT + " Please check the contract and update."
			MSG_TEXT = dao.Translate("OriginWarning", MSG_TEXT)
			MSG_BODY = dao.Translate("OriginWarning", MSG_BODY)
			MSG_TEXT = fmt.Sprintf(MSG_TEXT, o.Code, o.FullName)
			MSG_BODY = fmt.Sprintf(MSG_BODY, o.Code, o.FullName)
			core.SendEmail(am.Email, am.Name, MSG_TEXT, MSG_BODY)
			core.SendEmail(pm.Email, pm.Name, MSG_TEXT, MSG_BODY)
			noWarnings++
			continue
		}
		// Check if the contract is due to expire
		if o.EndDate != "" {
			contractExpiry, err := time.Parse(core.DATEFORMAT, o.EndDate)
			if err != nil {
				logs.Error("Origin Housekeeping", err)
				return message, err
			}
			if contractExpiry.Before(contractExpiryDate) {

				logs.Warning("Origin Housekeeping - Origin " + o.Code + " contract expires on " + o.EndDate)
				noDue++

				MSG_TEXT := "%s - %s contract expires on %s."
				MSG_BODY := MSG_TEXT + " Please check notify the client re: extension/renewal."
				MSG_TEXT = dao.Translate("OriginWarning", MSG_TEXT)
				MSG_BODY = dao.Translate("OriginWarning", MSG_BODY)
				MSG_TEXT = fmt.Sprintf(MSG_TEXT, o.Code, o.FullName, o.EndDate)
				MSG_BODY = fmt.Sprintf(MSG_BODY, o.Code, o.FullName, o.EndDate)
				core.SendEmail(am.Email, am.Name, MSG_TEXT, MSG_BODY)
				core.SendEmail(pm.Email, pm.Name, MSG_TEXT, MSG_BODY)

			}
			if contractExpiry.Before(time.Now()) {

				logs.Warning("Origin Housekeeping - Origin " + o.Code + " contract expired on " + o.EndDate)
				noOverdue++

				MSG_TEXT := "%s - %s contract expired on %s."
				MSG_BODY := MSG_TEXT + " Please check the notify the client as their license file may be invalidated. Please also check the contract and update"
				MSG_TEXT = dao.Translate("OriginWarning", MSG_TEXT)
				MSG_BODY = dao.Translate("OriginWarning", MSG_BODY)
				MSG_TEXT = fmt.Sprintf(MSG_TEXT, o.Code, o.FullName, o.EndDate)
				MSG_BODY = fmt.Sprintf(MSG_BODY, o.Code, o.FullName, o.EndDate)
				core.SendEmail(am.Email, am.Name, MSG_TEXT, MSG_BODY)
				core.SendEmail(pm.Email, pm.Name, MSG_TEXT, MSG_BODY)
			}
		}

		/// CONTENT ENDS

	}
	message = "Origins: %d Warnings: %d Due: %d Overdue: %d"
	dao.Translate("OriginWarning", message)
	message = fmt.Sprintf(message, noOrigins, noWarnings, noDue, noOverdue)
	return message, nil
}
