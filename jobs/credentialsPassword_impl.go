package jobs

import (
	"fmt"
	"time"

	application "github.com/mt1976/ebEstimates/application"
	core "github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	"github.com/mt1976/ebEstimates/logs"
)

func CredentialsPassword_Job_impl(j dm.JobDefinition) dm.JobDefinition {
	j.Name = "PasswordCredentials"
	j.Period = "0 0 * * *"
	j.Description = "Password Housekeeping"
	j.Type = core.HouseKeeping
	j.ID = "PasswordCredentials"
	return j
}

func CredentialsPassword_Run_impl() (string, error) {

	objectName := dao.Translate("ObjectName", "Credentials")

	message := ""
	/// CONTENT STARTS
	msg := "The notification period for passwords due to expire. In days."
	warningPeriod, err := dao.Data_GetInt(objectName, "Password_Expiry_Notification_Period", dm.Data_Category_Setting, msg)
	if err != nil {
		return message, err
	}
	if warningPeriod == 0 {
		logs.Warning("Password Expiry Warning Period not set using default of 7 days")
		warningPeriod = 7
	}

	warningDate := time.Now().AddDate(0, 0, warningPeriod*-1)

	//Get All Credentials
	_, creds, err := dao.Credentials_GetList()
	if err != nil {
		return message, err
	}
	today := time.Now()
	for _, cred := range creds {
		expiry, err := time.Parse(core.DATEFORMAT, cred.PasswordExpiry)
		if err != nil {
			return message, err
		}
		if today.After(expiry) || today.Equal(expiry) {
			message = "Password expired for user: " + cred.Username + " on " + cred.PasswordExpiry
			cred.Password = ""
			cred.PasswordExpiry = ""
			cred.Notes = core.AddActivity_ForUser(cred.Notes, "Password expired", "System")
			_, err := dao.Credentials_StoreSystem(cred)
			if err != nil {
				return message, err
			}
			MSG_BODY := "Your password for %s has expired. <br><br>Please update you password a.s.a.p. <br><br>Thank you."
			MSG_BODY = dao.Translate("Passwords", MSG_BODY)
			MSG_BODY = fmt.Sprintf(MSG_BODY, core.ApplicationName())
			core.SendEmail(cred.Email, "", core.ApplicationName()+" Password expired", MSG_BODY)

		}
		if cred.Password == "" && cred.PasswordExpiry != "" {
			message = "Password expired for user: " + cred.Username + " on " + cred.PasswordExpiry
			cred.PasswordExpiry = ""
			cred.Notes = core.AddActivity_ForUser(cred.Notes, "Password expired", "System")
			_, err := dao.Credentials_StoreSystem(cred)
			if err != nil {
				return message, err
			}
		}
		if today.Before(expiry) || today.After(warningDate) {
			message = "Password expiry warning for user: " + cred.Username + " on " + cred.PasswordExpiry
			cred.Notes = core.AddActivity_ForUser(cred.Notes, "Password expiry warning", "System")
			_, err := dao.Credentials_StoreSystem(cred)
			if err != nil {
				return message, err
			}
			MSG_BODY := "Your password for %s will expire on %s. <br><br>Please update you password a.s.a.p. <br><br>Thank you."
			MSG_BODY = dao.Translate("Passwords", MSG_BODY)
			MSG_BODY = fmt.Sprintf(MSG_BODY, core.ApplicationName(), cred.PasswordExpiry)
			core.SendEmail(cred.Email, "", core.ApplicationName()+" Password due to expire", MSG_BODY)
			err4 := application.Inbox_SendMailSystem(cred.Email, MSG_BODY, "Password Housekeeping")
			if err4 != nil {
				logs.Error("Password Housekeeping", err4)
				return message, err4
			}
		}
	}
	/// CONTENT ENDS

	return message, nil

}
