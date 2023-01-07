package jobs

import (
	core "github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/dao"
)

func SendEmail(resourceID string, MSG_SUBJECT string, MSG_BODY string) {
	_, who, _ := dao.Resource_GetByCode(resourceID)
	core.SendEmail(who.Email, who.Name, MSG_SUBJECT, MSG_BODY)
}
