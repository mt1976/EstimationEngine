package jobs

import (
	"fmt"
	"time"

	core "github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
	"golang.org/x/exp/slices"
)

func Project_Job_impl(j dm.JobDefinition) dm.JobDefinition {
	j.Name = "Project"
	j.Period = "10 0 * * *"
	j.Description = "Notify of Project Expiry"
	j.Type = core.General
	j.ID = "Project"
	return j
}

func Project_Run_impl() (string, error) {

	message := ""
	/// CONTENT STARTS
	period, err0 := dao.Data_GetInt("Project", "DueWarning")
	if err0 != nil {
		return message, err0
	}
	overdueState, err1 := dao.Data_GetString("Project", "OverdueState")
	if err1 != nil {
		return message, err1
	}
	closedStates, err2 := dao.Data_GetArray("Project", "ClosedStates")
	if err2 != nil {
		return message, err2
	}

	// Calculate Warning Date
	warn := period
	notifyDate := time.Now().AddDate(0, 0, warn)
	logs.Information("Project End Notice", notifyDate.Format(core.DATEFORMATUSER))
	// Get the projects list
	_, projects, err := dao.Project_GetList()
	if err != nil {
		logs.Error("Project Housekeeping", err)
		return message, err
	}
	// Scan the projects
	noProjects := 0
	noWarnings := 0
	noOverdue := 0
	for _, p := range projects {
		noProjects++
		if p.EndDate == "" {
			continue
		}
		// Check if the project is due to expire
		// Get Acount Manager from Origin
		warn := warnExpiry(p, notifyDate)
		if warn {
			noWarnings++
		}
		// Check if the project is overdue
		if (p.ProjectStateID != overdueState) && !(slices.Contains(closedStates, p.ProjectStateID)) {
			overdue := overdueProcessing(p, overdueState)
			if overdue {
				noOverdue++
			}
		}
		if p.ProjectStateID == overdueState {
			sendProjectOverdueMail(p)
		}
	}
	/// CONTENT ENDS
	MSG_TEXT := "Project Housekeeping: %d projects, %d warnings, %d overdue"
	MSG_TEXT = dao.Translate("ProjectJob", MSG_TEXT)
	message = fmt.Sprintf(MSG_TEXT, noProjects, noWarnings, noOverdue)
	logs.Information("Project Housekeeping", message)
	return message, nil

}

func overdueProcessing(p dm.Project, overdueState string) bool {
	// Get the project end date
	endDate, _ := time.Parse(core.DATEFORMAT, p.EndDate)
	if endDate.Before(time.Now()) {
		p.ProjectStateID = overdueState
		dao.Project_StoreSystem(p)
		return true
	}
	return false
}

func warnExpiry(p dm.Project, notifyDate time.Time) bool {
	endDate, _ := time.Parse(core.DATEFORMAT, p.EndDate)
	if endDate.Before(time.Now()) {
		// This will be handled by the Expiry Processing
		return false
	}
	if endDate.After(time.Now()) && endDate.Before(notifyDate) {
		// If its after today and before the notify date then send the warning
		_, o, _ := dao.Origin_GetByCode(p.OriginID)

		if p.ProjectManager != "" {
			sendProjectExpiryMail(p, p.ProjectManager)
		} else {
			sendProjectExpiryMail(p, o.AccountManager)
		}

		sendProjectExpiryMail(p, p.ProjectAnalyst)
		sendProjectExpiryMail(p, p.ProjectEngineer)
		sendProjectExpiryMail(p, o.AccountManager)
		return true
	}
	return false
}

func sendProjectOverdueMail(p dm.Project) {
	_, o, _ := dao.Origin_GetByCode(p.OriginID)

	if p.ProjectManager != "" {
		sendProjectExpiredMail(p, p.ProjectManager)
	} else {
		sendProjectExpiredMail(p, o.AccountManager)
	}

	sendProjectExpiredMail(p, p.ProjectAnalyst)
	sendProjectExpiredMail(p, p.ProjectEngineer)
	sendProjectExpiredMail(p, o.AccountManager)
}

func sendProjectExpiryMail(p dm.Project, resID string) {
	MSG_SUBJECT := "INFORMATION! %s Project %s is due to end on %s"
	MSG_SUBJECT = dao.Translate("Email", MSG_SUBJECT)
	MSG_SUBJECT = fmt.Sprintf(MSG_SUBJECT, p.OriginID, p.Name, p.EndDate)

	MSG_BODY := "%s Project %s is due to end on %s. Please review the project and update the end date if necessary."
	MSG_BODY = dao.Translate("Email", MSG_BODY)
	MSG_BODY = fmt.Sprintf(MSG_BODY, p.OriginID, p.Name, p.EndDate)

	SendEmail(resID, MSG_SUBJECT, MSG_BODY)
}

func sendProjectExpiredMail(p dm.Project, resID string) {
	MSG_SUBJECT := "WARNING! %s Project %s is due to have ended on %s"
	MSG_SUBJECT = dao.Translate("Email", MSG_SUBJECT)
	MSG_SUBJECT = fmt.Sprintf(MSG_SUBJECT, p.OriginID, p.Name, p.EndDate)

	MSG_BODY := "%s Project %s is due to have ended on %s. Please review the project and update the end date if necessary."
	MSG_BODY = dao.Translate("Email", MSG_BODY)
	MSG_BODY = fmt.Sprintf(MSG_BODY, p.OriginID, p.Name, p.EndDate)

	SendEmail(resID, MSG_SUBJECT, MSG_BODY)
}
