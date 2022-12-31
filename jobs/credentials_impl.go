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
	// ------------------
	// Get the credentials prompt period
	// ------------------
	period, err0 := strconv.Atoi(core.GetApplicationProperty("credentialsprompt"))
	warn := period
	if err0 != nil {
		logs.Error("Credentials Housekeeping", err0)
		return message, err0
	}

	notifyDate := time.Now().AddDate(0, 0, warn)
	logs.Information("Credentials Expiry Notice", notifyDate.Format(core.DATEFORMATUSER))

	// ------------------
	// Get the credentials list
	// ------------------
	_, creds, err := dao.Credentials_GetList()
	if err != nil {
		logs.Error("Credentials Housekeeping", err)
		return message, err
	}
	// ------------------
	// Scan the credentials
	// ------------------
	for _, c := range creds {
		expiry, err2 := time.Parse(core.DATETIMEFORMATUSER, c.Expiry)
		if err2 != nil {
			logs.Error("Credentials Housekeeping", err2)
			return message, err2
		}
		// ------------------
		// Check for expired credentials
		// ------------------
		if expiry.Before(time.Now()) {
			shouldReturn, returnValue, returnErr := expireCredentials(c, message)
			if shouldReturn {
				return returnValue, returnErr
			}
		}
		// ------------------
		// Check for credentials due to expire
		// ------------------
		if expiry.Before(notifyDate) && c.State != "EXPD" {
			shouldReturn, returnValue, returnErr1 := warnExpiryOfCredentials(c, period, message)
			if shouldReturn {
				return returnValue, returnErr1
			}
		}
	}
	/// CONTENT ENDS

	return message, nil

}

// warnExpiryOfCredentials sends a warning to the user that their credentials are due to expire
func warnExpiryOfCredentials(c dm.Credentials, period int, message string) (bool, string, error) {
	logs.Information("Credentials Housekeeping", "Credentials due to expire : "+c.Username)
	// ------------------
	// Add a note to the credentials
	// ------------------
	MSG_TXT := "Credentials will expire in %v days on %s."
	dao.Translate("Credentials", MSG_TXT)
	MSG_TXT = fmt.Sprintf(MSG_TXT, period, c.Expiry)
	c.Notes = core.AddActivity_ForProcess(c.Notes, MSG_TXT, "Credentials Housekeeping")
	// ------------------
	// Send a message to the inbox
	// ------------------
	MSG_INBX := "Your credentials %s will expire in %v days at %s. <br><br>Please contact the administrator for %s to renew them"
	MSG_INBX = dao.Translate("Credentials", MSG_INBX)
	MSG_INBX = fmt.Sprintf(MSG_INBX, c.Username, core.ApplicationName(), core.ApplicationHostname())
	err4 := application.Inbox_SendMailSystem(c.Email, MSG_INBX, "Credentials Housekeeping")
	if err4 != nil {
		logs.Error("Credentials Housekeeping", err4)
		return true, message, err4
	}
	// ------------------
	// Update the credentials
	// ------------------
	err3 := dao.Credentials_StoreSystem(c)
	if err3 != nil {
		logs.Error("Credentials Housekeeping", err3)
		return true, message, err3
	}
	// ------------------
	// Email the user
	// ------------------
	MSG_BODY := "Your credentials %s will expire in %v days at %s. <br><br>Please contact the administrator for %s to renew them"
	MSG_BODY = dao.Translate("Credentials", MSG_BODY)
	MSG_BODY = fmt.Sprintf(MSG_BODY, c.Username, period, c.Expiry, core.ApplicationHostname())
	core.SendEmail(c.Email, c.Firstname, core.ApplicationName()+" Credentials Due to Expire", MSG_BODY)
	return false, "", nil
}

// expireCredentials expires the credentials
func expireCredentials(c dm.Credentials, message string) (bool, string, error) {
	logs.Information("Credentials Housekeeping", "Expired credentials removed/warned : "+c.Username)
	// ------------------
	// Add a note to the credentials
	// ------------------
	MSG_TXT := "Credentials %s have expired"
	dao.Translate("Credentials", MSG_TXT)
	MSG_TXT = fmt.Sprintf(MSG_TXT, c.Username)
	c.Notes = core.AddActivity_ForProcess(c.Notes, MSG_TXT, "Credentials Housekeeping")
	// ------------------
	// Update the credentials
	// ------------------
	c.State = "EXPD"
	c.EmailNotifications = "false"
	err3 := dao.Credentials_StoreSystem(c)
	if err3 != nil {
		logs.Error("Credentials Housekeeping", err3)
		return true, message, err3
	}
	logs.Expired("Credentials " + c.Username + " expired " + c.Email)
	// ------------------
	// Email the user
	// ------------------
	MSG_BODY := "Your credentials %s have expired. <br><br>Please contact the administrator for %s to renew them"
	MSG_BODY = dao.Translate("Credentials", MSG_BODY)
	MSG_BODY = fmt.Sprintf(MSG_BODY, c.Username, core.ApplicationHostname())
	core.SendEmail(c.Email, c.Firstname, core.ApplicationName()+" Credentials Expired", MSG_BODY)
	return false, "", nil
}
