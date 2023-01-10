package jobs

import (
	core "github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/dao"
	"github.com/mt1976/ebEstimates/logs"
)

func SendMailToResource(resourceID string, MSG_SUBJECT string, MSG_BODY string) {
	_, who, _ := dao.Resource_GetByCode(resourceID)
	if resourceID == "" {
		logs.Warning("SendMail: No resourceID - not sending email")
		return
	}
	if MSG_BODY == "" || MSG_SUBJECT == "" {
		logs.Warning("SendMail: No message body or subject - not sending email")
		return
	}
	if who.Email == "" {
		logs.Warning("SendMail: No email address for " + core.DQuote(who.Name) + " " + core.DQuote(resourceID) + " - not sending email")
		return
	}
	if who.UserActive == core.FALSE {
		logs.Warning("SendMail: User " + core.DQuote(who.Name) + " " + core.DQuote(resourceID) + " is not active - not sending email")
		return
	}
	core.SendEmail(who.Email, who.Name, MSG_SUBJECT, MSG_BODY)
}
