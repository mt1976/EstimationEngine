package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/estimationsession.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : EstimationSession (estimationsession)
// Endpoint 	        : EstimationSession (EstimationSessionID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 27/11/2022 at 20:46:13
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"errors"
	"fmt"

	core "github.com/mt1976/ebEstimates/core"
	logs "github.com/mt1976/ebEstimates/logs"

	dm "github.com/mt1976/ebEstimates/datamodel"
)

// EstimationSession_GetList() returns a list of all EstimationSession records
func EstimationSession_ByProject_GetList() (int, []dm.EstimationSession, error) {

	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.EstimationSession_SQLTable)
	count, estimationsessionList, _, _ := estimationsession_Fetch(tsql)

	return count, estimationsessionList, nil
}

// EstimationSession_GetList() returns a list of all EstimationSession records
func EstimationSession_Active_ByProject_GetList(id string) (int, []dm.EstimationSession, error) {

	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.EstimationSession_SQLTable)
	tsql = tsql + " WHERE " + dm.EstimationSession_ProjectID_sql + "='" + id + "'"
	tsql = tsql + " AND datalength(" + dm.EstimationSession_SYSDeleted_sql + ") = 0"
	count, estimationsessionList, _, _ := estimationsession_Fetch(tsql)

	return count, estimationsessionList, nil
}

// EstimationSession_GetList() returns a list of all EstimationSession records
func EstimationSession_Get_ByADO(id string) (int, []dm.EstimationSession, error) {

	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.EstimationSession_SQLTable)
	tsql = tsql + " WHERE " + dm.EstimationSession_AdoID_sql + "='" + id + "'"
	tsql = tsql + " AND datalength(" + dm.EstimationSession_SYSDeleted_sql + ") = 0"
	count, estimationsessionList, _, _ := estimationsession_Fetch(tsql)
	if count > 1 {
		return count, estimationsessionList, errors.New("duplicate records found")
	}
	return count, estimationsessionList, nil
}

// EstimationSession_GetList() returns a list of all EstimationSession records
func EstimationSession_Get_ByFreshDesk(id string) (int, []dm.EstimationSession, error) {

	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.EstimationSession_SQLTable)
	tsql = tsql + " WHERE " + dm.EstimationSession_FreshdeskID_sql + "='" + id + "'"
	tsql = tsql + " AND datalength(" + dm.EstimationSession_SYSDeleted_sql + ") = 0"
	count, estimationsessionList, _, _ := estimationsession_Fetch(tsql)
	if count > 1 {
		return count, estimationsessionList, errors.New("duplicate records found")
	}
	return count, estimationsessionList, nil
}

// EstimationSession_List_ByStatus() returns a list of all EstimationSession records
func EstimationSession_ListActive_ByState(id string) (int, []dm.EstimationSession, error) {
	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.EstimationSession_SQLTable)
	tsql = tsql + " WHERE " + dm.EstimationSession_EstimationStateID_sql + "='" + id + "'"
	tsql = tsql + " AND datalength(" + dm.EstimationSession_SYSDeleted_sql + ") = 0"
	count, estimationsessionList, _, _ := estimationsession_Fetch(tsql)
	return count, estimationsessionList, nil
}

func StatusChangeMessage(action string, val string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "StatusChangeMessage", PUT, rec.EstimationSessionID)
	// Get Project from DB
	_, project, _ := Project_GetByID(rec.ProjectID)

	action = Translate("EstimationSessionStatus", action)
	// Send a notification to the Estimation Owner
	MSG_TXT := Translate("Notification", "%s: Estimation ''%s'' has been updated. %s (%s)")
	MSG_TXT = fmt.Sprintf(MSG_TXT, project.OriginID, rec.Name, action, val)
	MSG_BODY := Translate("NotificationBody", "%s: Estimation ''%s'' has been updated to ''%s'' (%s). Please review and take appropriate action. %s")
	// Get the Estimation from the database
	MSG_BODY = fmt.Sprintf(MSG_BODY, project.OriginID, rec.Name, action, val, rec.Comments)
	URI := "/EstimationSessionFormatted/?EstimationSessionID=" + rec.EstimationSessionID
	core.Notification_URL(MSG_TXT, MSG_BODY, URI)

	// Get the Origin
	_, origin, _ := Origin_GetByCode(rec.Origin)

	// Get the Project Manager
	thisPM := project.ProjectManager
	if thisPM == "" {
		thisPM = origin.ProjectManager
	}

	// Send a notification to the Project Manager
	//_, pm, _ := Resource_GetByCode(thisPM)
	//core.SendEmail(pm.Email, pm.Name, MSG_TXT, MSG_BODY)
	SendMailToResource(thisPM, MSG_TXT, MSG_BODY)
	// Send a notification to the Account Manager
	//_, am, _ := Resource_GetByCode(origin.AccountManager)
	//core.SendEmail(am.Email, am.Name, MSG_TXT, MSG_BODY)
	SendMailToResource(origin.AccountManager, MSG_TXT, MSG_BODY)
	return "", nil
}

func emailRelatedResources(rec dm.EstimationSession) (string, error) {

	_, project, _ := Project_GetByID(rec.ProjectID)
	_, origin, _ := Origin_GetByCode(project.OriginID)

	// Get the email addresses
	var pmEmail string

	if rec.ProjectManager == "" {
		// Get the Project Manager from the Project
		pmEmail = project.ProjectManager
	} else {
		pmEmail = rec.ProjectManager
	}
	_, pmResource, _ := Resource_GetByID(pmEmail)

	pdEmail := rec.ProductManager
	acEmail := origin.AccountManager

	MSG_SUBJECT := "%s Quote Expired - %s"
	MSG_SUBJECT = Translate("QuoteExpiredSubject", MSG_SUBJECT)
	MSG_SUBJECT = fmt.Sprintf(MSG_SUBJECT, project.OriginID, rec.Name, pmResource.Name, pmResource.Email)

	MSG_BODY := "Quote for %s %s has expired. Please contact the Project Manager %s (%s) for further information."
	MSG_SUBJECT = Translate("QuoteExpiredBody", MSG_BODY)
	MSG_BODY = fmt.Sprintf(MSG_BODY, project.OriginID, rec.Name, pmResource.Name, pmResource.Email)

	// Send the email
	// core.SendEmail(pmEmail.Email, pmEmail.Name, MSG_SUBJECT, MSG_BODY)
	// core.SendEmail(pdEmail.Email, pdEmail.Name, MSG_SUBJECT, MSG_BODY)
	// core.SendEmail(acEmail.Email, acEmail.Name, MSG_SUBJECT, MSG_BODY)
	SendMailToResource(pmEmail, MSG_SUBJECT, MSG_BODY)
	SendMailToResource(pdEmail, MSG_SUBJECT, MSG_BODY)
	SendMailToResource(acEmail, MSG_SUBJECT, MSG_BODY)

	return "", nil
}
