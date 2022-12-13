package jobs

import (
	"fmt"
	"strconv"
	"time"

	application "github.com/mt1976/ebEstimates/application"
	core "github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

func Credentials_Job_impl(j dm.JobDefinition) dm.JobDefinition {
	j.Name = "Credentials"
	j.Period = "0 0 * * *"
	j.Description = "Credentials Housekeeping"
	j.Type = core.HouseKeeping
	j.ID = "Credentials"
	return j
}

func Credentials_Run_impl() (string, error) {

	message := ""
	/// CONTENT STARTS
	logs.Information("Credentials Housekeeping", "Expired credentials removed/warned")

	period, err0 := strconv.Atoi(core.GetApplicationProperty("credentialsprompt"))
	warn := period
	if err0 != nil {
		logs.Error("Credentials Housekeeping", err0)
		return message, err0
	}

	notifyDate := time.Now().AddDate(0, 0, warn)

	logs.Information("Credentials Expiry Notice", notifyDate.Format(core.DATEFORMATUSER))

	_, creds, err := dao.Credentials_GetList()
	if err != nil {
		logs.Error("Credentials Housekeeping", err)
		return message, err
	}
	for _, c := range creds {
		expiry, err2 := time.Parse(c.Expiry, core.DATETIMEFORMATSQLSERVER)
		if err2 != nil {
			logs.Error("Credentials Housekeeping", err2)
			return message, err2
		}

		if expiry.Before(time.Now()) {
			c.State = "EXPD"
			c.EmailNotifications = "false"
			MSG_TXT := "Your credentials %s for %s have expired. Please contact the administrator for %s to renew them."
			dao.Translate("Credentials", MSG_TXT)
			MSG_TXT = fmt.Sprintf(MSG_TXT, c.Username, core.ApplicationName(), core.ApplicationHostname())
			c.Notes = core.AddActivity_ForProcess(c.Notes, MSG_TXT, "Credentials Housekeeping")
			err3 := dao.Credentials_StoreSystem(c)
			if err3 != nil {
				logs.Error("Credentials Housekeeping", err3)
				return message, err3
			}
			logs.Expired("Credentials " + c.Username + " expired " + c.Email)
			MSG_BODY := "Your credentials %s have expired. <br><br>Please contact the administrator for %s to renew them"
			MSG_BODY = dao.Translate("Credentials", MSG_BODY)
			MSG_BODY = fmt.Sprintf(MSG_BODY, c.Username, core.ApplicationHostname())
			core.SendEmail(c.Email, c.Firstname, core.ApplicationName()+" Credentials Expired", MSG_BODY)
		}

		if expiry.Before(notifyDate) && c.State != "EXPD" {
			MSG_TXT := "Your credentials %s for %s will expire in %v days at %s. Please contact the administrator for %s to renew them."
			dao.Translate("Credentials", MSG_TXT)
			MSG_TXT = fmt.Sprintf(MSG_TXT, c.Username, core.ApplicationName(), period, c.Expiry, core.ApplicationHostname())
			c.Notes = core.AddActivity_ForProcess(c.Notes, MSG_TXT, "Credentials Housekeeping")
			err4 := application.Inbox_SendMailSystem(c.Email, MSG_TXT, "Credentials Housekeeping")
			if err4 != nil {
				logs.Error("Credentials Housekeeping", err4)
				return message, err4
			}
			err3 := dao.Credentials_StoreSystem(c)
			if err3 != nil {
				logs.Error("Credentials Housekeeping", err3)
				return message, err3
			}
			MSG_BODY := "Your credentials %s will expire in %v days at %s. <br><br>Please contact the administrator for %s to renew them"
			MSG_BODY = dao.Translate("Credentials", MSG_BODY)
			MSG_BODY = fmt.Sprintf(MSG_BODY, c.Username, period, c.Expiry, core.ApplicationHostname())
			core.SendEmail(c.Email, c.Firstname, core.ApplicationName()+" Credentials Due to Expire", MSG_BODY)
		}
	}
	/// CONTENT ENDS

	return message, nil

}
