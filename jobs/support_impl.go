package jobs

import (
	"strconv"
	"strings"
	"time"

	core "github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	"github.com/mt1976/ebEstimates/logs"
)

const (
	JOB_ACTION_IMPORT     string = "IMPORT"
	JOB_ACTION_EXPORT     string = "EXPORT"
	JOB_ACTION                   = "Action"
	JOB_MODE                     = "Mode"
	JOB_MODE_TRIAL               = "Trial"
	JOB_LAST_RUN                 = "LastRun"
	JOB_LAST_RUN_IN_PATH         = "LastRun_Path_In"
	JOB_LAST_RUN_OUT_PATH        = "LastRun_Path_Out"
	JOB_LAST_RUN_MESSAGE         = "LastRun_Message"
	JOB_LAST_RUN_MODE            = "LastRun_Mode"
	JOB_LAST_RUN_COUNT           = "LastRun_Count"
	JOB_LAST_RUN_ACTION          = "LastRun_Action"
	JOB_DONE                     = "Done"
	JOB_IN_PATH                  = "In_Path"
	JOB_OUT_PATH                 = "Out_Path"
	JOB_WHAT                     = "What"
)

// SendMailToResource sends an email to a resource, if the resource has an email address and is active
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

// RecordJobStatusInfo records the status of a job
func RecordJobStatusInfo(message string, noWorkItems int, action string, jobName string, inPath string, outPath string, ssMode string) string {
	noItemsProcessed := strconv.Itoa(noWorkItems) + " Work Items Processed"
	dao.Data_Put(jobName, JOB_LAST_RUN, dm.Data_Category_Info, time.Now().Format(core.DATEMSG))
	dao.Data_Put(jobName, JOB_LAST_RUN_COUNT, dm.Data_Category_Info, noItemsProcessed)
	if action != "" {
		dao.Data_Put(jobName, JOB_LAST_RUN_ACTION, dm.Data_Category_State, action)
	}
	if inPath != "" {
		dao.Data_Put(jobName, JOB_LAST_RUN_IN_PATH, dm.Data_Category_Info, inPath)
	}
	if outPath != "" {
		dao.Data_Put(jobName, JOB_LAST_RUN_OUT_PATH, dm.Data_Category_Path, outPath)
	}
	dao.Data_Put(jobName, JOB_LAST_RUN_MESSAGE, dm.Data_Category_Info, message)
	dao.Data_Put(jobName, JOB_LAST_RUN_MODE, dm.Data_Category_Info, ssMode)
	dao.Data_Put(jobName, JOB_MODE, dm.Data_Category_State, JOB_DONE)
	return ""
}

// TrialMode returns true if the job is in trial mode
func TrialMode(in string) bool {
	if strings.ToUpper(in) == JOB_MODE_TRIAL {
		logs.Warning("Trial Mode : No Updates will be made!")
		return true
	}
	return false
}
