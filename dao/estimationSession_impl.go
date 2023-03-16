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
	"strconv"

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
	Resource_SendMail(thisPM, MSG_TXT, MSG_BODY)
	// Send a notification to the Account Manager
	//_, am, _ := Resource_GetByCode(origin.AccountManager)
	//core.SendEmail(am.Email, am.Name, MSG_TXT, MSG_BODY)
	Resource_SendMail(origin.AccountManager, MSG_TXT, MSG_BODY)
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
	Resource_SendMail(pmEmail, MSG_SUBJECT, MSG_BODY)
	Resource_SendMail(pdEmail, MSG_SUBJECT, MSG_BODY)
	Resource_SendMail(acEmail, MSG_SUBJECT, MSG_BODY)

	return "", nil
}

func Estimationsession_Calculate(searchID string, trigger string) dm.EstimationSession {

	logs.Processing("estimationsession_ReCalculate-->" + searchID)

	_, esRecord, estErr := EstimationSession_GetByID(searchID)
	if estErr != nil {
		logs.Error("estimationsession_ReCalculate", estErr)
		return dm.EstimationSession{}
	}
	//Get Estimation State
	_, estState, esStateErr := EstimationState_GetByCode(esRecord.EstimationStateID)
	if esStateErr != nil {
		logs.Error("estimationsession_ReCalculate", esStateErr)
		return dm.EstimationSession{}
	}
	if estState.IsLocked == core.TRUE {
		logs.Information("Estimation State is Locked - No Update Required", esRecord.EstimationStateID)
		return esRecord
	}

	// Get the current projiect id
	_, proj, projErr := Project_GetByID(esRecord.ProjectID)
	if projErr != nil {
		logs.Error("estimationsession_ReCalculate", projErr)
		return dm.EstimationSession{}
	}
	// Get Project State
	_, projState, projStateErr := ProjectState_GetByCode(proj.ProjectStateID)
	if projStateErr != nil {
		logs.Error("estimationsession_ReCalculate", projStateErr)
		return dm.EstimationSession{}
	}
	if projState.IsLocked == core.TRUE {
		logs.Information("Project State is Locked - No Update Required", esRecord.ProjectStateID)
		return esRecord
	}

	//Get the current rate from the origin
	_, origin, originErr := Origin_GetByCode(proj.OriginID)
	if originErr != nil {
		logs.Error("estimationsession_ReCalculate", originErr)
		return dm.EstimationSession{}
	}
	// Get Origin State
	_, originState, originStateErr := OriginState_GetByCode(esRecord.OriginStateID)
	if originStateErr != nil {
		logs.Error("estimationsession_ReCalculate", originStateErr)
		return dm.EstimationSession{}
	}
	if originState.IsLocked == core.TRUE {
		logs.Information("Origin State is Locked - No Update Required", esRecord.OriginStateID)
		return esRecord
	}

	//Get the current profile
	profileID := proj.ProfileID
	originID := proj.OriginID
	logs.Information("ProjectID", esRecord.ProjectID)
	logs.Information("ProfileID", profileID)
	logs.Information("OriginID", originID)

	_, profile, profileErr := Profile_GetByCode(profileID)
	if profileErr != nil {
		logs.Error("estimationsession_ReCalculate", profileErr)
		return dm.EstimationSession{}
	}

	// Get the list of all Features for this EstimationSession
	foundFeatures, featureList, _ := Feature_Active_ByEstimationSession_GetList(searchID)
	logs.Information("Features Found", strconv.Itoa(foundFeatures))

	RoundingFactor := core.StringToFloat(profile.Rounding)

	HoursInDay := core.StringToFloat(profile.HoursPerDay)
	Rate := core.StringToFloat(origin.Rate)
	if proj.DefaultRate != "" {
		Rate = core.StringToFloat(proj.DefaultRate)
	}
	if proj.ProjectRate != "" {
		Rate = core.StringToFloat(proj.ProjectRate)
	}

	logs.Information("HoursInDay", core.FloatToString(HoursInDay))
	logs.Information("Rate", core.FloatToString(Rate))

	Total_Reqs := 0.00
	Total_AnaTest := 0.00
	Total_Docs := 0.00
	Total_Mgt := 0.00
	Total_UAT := 0.00
	Total_MKT := 0.00
	Total_Training := 0.00
	//Total_Contingency := 0.00
	Total_DevEstimate := 0.00
	Total_DevUplift := 0.00

	// Range through the list of Features
	for thisFeature, feature := range featureList {
		// Sum up the estimates in this featre
		// convert feature.Reqs to INT
		logs.Break()
		logs.Information("Feature", strconv.Itoa(thisFeature))
		//	logs.Information("Feature.Reqs", feature.Reqs)

		Total_Reqs += core.StringToFloat(feature.Requirements)
		Total_AnaTest += core.StringToFloat(feature.AnalystTest)
		Total_Docs += core.StringToFloat(feature.Documentation)
		Total_Mgt += core.StringToFloat(feature.Management)
		Total_UAT += core.StringToFloat(feature.UatSupport)
		Total_MKT += core.StringToFloat(feature.Marketing)
		//Total_Contingency += core.StringToFloat(feature.Contingency)
		Total_DevUplift += core.StringToFloat(feature.DevelopmentFactored)
		Total_DevEstimate += core.StringToFloat(feature.DevelopmentEstimate)
		Total_Training += core.StringToFloat(feature.Training)
	}

	//logs.Information("Total_Reqs: ", fmt.Sprintf("%.2f", Total_Reqs))

	ReleaseHoursInDay := core.StringToFloat(profile.DefaultReleaseHours)
	RelHours := core.StringToFloat(profile.DefaultReleases) * ReleaseHoursInDay

	if esRecord.Releases != "" {
		RelHours = core.StringToFloat(esRecord.Releases) * ReleaseHoursInDay
	}
	// Convert Hours to Days

	Total_Reqs_Days, _ := calculate(Total_Reqs, HoursInDay, Rate)
	Total_AnaTest_Days, _ := calculate(Total_AnaTest, HoursInDay, Rate)
	Total_Docs_Days, _ := calculate(Total_Docs, HoursInDay, Rate)
	Total_Mgt_Days, _ := calculate(Total_Mgt, HoursInDay, Rate)
	Total_UAT_Days, _ := calculate(Total_UAT, HoursInDay, Rate)
	Total_MKT_Days, _ := calculate(Total_MKT, HoursInDay, Rate)
	//Total_Contingency_Days, _ := calculate(Total_Contingency, HoursInDay, Rate)
	//Total_DevEstimate_Days, Total_DevEstimate_Cost := calculate(Total_DevEstimate, HoursInDay, Rate)
	Total_DevUplift_Days, _ := calculate(Total_DevUplift, HoursInDay, Rate)
	Total_Rel_Days, _ := calculate(RelHours, ReleaseHoursInDay, Rate)
	Total_Training_Days, _ := calculate(Total_Training, HoursInDay, Rate)

	TReqs, _ := rtn(Total_Reqs_Days, RoundingFactor)

	// Requirement Days
	esRecord.ReqDays = core.FloatToString(TReqs)
	esRecord.RegCost = core.FloatToString((TReqs * HoursInDay) * Rate)
	Timps, _ := rtn(Total_DevUplift_Days+Total_AnaTest_Days+Total_Docs_Days+Total_MKT_Days+Total_Training_Days, RoundingFactor)
	// Implementaion Days
	esRecord.ImpDays = core.FloatToString(Timps)
	esRecord.ImpCost = core.FloatToString((Timps * HoursInDay) * Rate)
	// UAT Days
	Tuat, _ := rtn(Total_UAT_Days, RoundingFactor)
	esRecord.UatDays = core.FloatToString(Tuat)
	esRecord.UatCost = core.FloatToString((Tuat * HoursInDay) * Rate)
	// Management Days
	Tman, _ := rtn(Total_Mgt_Days, RoundingFactor)
	esRecord.MgtDays = core.FloatToString(Tman)
	esRecord.MgtCost = core.FloatToString((Tman * HoursInDay) * Rate)
	// Contingency Days
	//Tcon, _ := rtn(Total_Contingency_Days, RoundingFactor)
	//esRecord.Contingency = core.FloatToString(Tcon)
	//esRecord.ConCost = core.FloatToString((Tcon * HoursInDay) * Rate)

	// Release Days
	Trel, _ := rtn(Total_Rel_Days, RoundingFactor)
	esRecord.RelDays = core.FloatToString(Trel)
	esRecord.RelCost = core.FloatToString((Trel * ReleaseHoursInDay) * Rate)

	Total_Days := TReqs + Timps + Tuat + Tman + Trel
	Total_Cost := Total_Days * HoursInDay * Rate

	SupportUplift := Total_Cost * (core.StringToFloat(profile.SupportUplift) / 100)
	esRecord.SupportUplift = core.FloatToString(SupportUplift)

	esRecord.EffortTotal = core.FloatToString(Total_Days)
	esRecord.Total = core.FloatToString(Total_Cost)

	// Update the Estimation Session Record
	estUpdateTXT := Translate("AuditMessage", "Refresh - %s")
	estUpdateTXT = fmt.Sprintf(estUpdateTXT, trigger)

	esRecord.Activity = core.AddActivity_ForProcess(esRecord.Activity, estUpdateTXT, "FeatureUpdate")

	EstimationSession_StoreSystem(esRecord)

	return esRecord
}

func calculate(hrs float64, hrsInDay float64, rate float64) (float64, float64) {

	days := 0.00
	cost := 0.00

	if hrs != 0 {

		if hrsInDay != 0 {
			days = hrs / hrsInDay
		}

		if rate != 0 {
			cost = hrs * rate
		}
	}

	return days, cost
}

// rtn Rounds a number to the nearest multiple of the RoundingFactor
func rtn(number float64, RoundingFactor float64) (float64, error) {
	// Round to the nearest multiple of the RoundingFactor
	rtnVal, rtnErr := core.RoundUpTo(number, RoundingFactor)
	return rtnVal, rtnErr
}

func EstimationSession_Clone(esID string, userID string) (string, error) {
	if esID == "" {
		logs.Warning("EstimationSession_Clone: No Estimation Session ID Supplied")
		return "", errors.New("no estimation session ID supplied")
	}

	_, esREC, _ := EstimationSession_GetByID(esID)

	esREC.SYSId = ""
	cloneID := EstimationSession_NewID(esREC)
	esREC.EstimationSessionID = cloneID
	origName := esREC.Name
	esREC.Name = fmt.Sprintf(Translate("ActionMessage", "%s (Clone)"), origName)
	esREC.Activity = core.AddActivity_ForUser(esREC.Activity, Translate("AuditMessage", "ESTIMATION CLONED"), userID)
	//esREC.Activity = r.FormValue(dm.EstimationSession_Activity_scrn)

	esREC.SYSCreated = Audit_TimeStamp()
	esREC.SYSCreatedBy = userID
	esREC.SYSCreatedHost = Audit_Host()
	esREC.SYSUpdated = ""
	esREC.SYSUpdatedBy = ""
	esREC.SYSUpdatedHost = ""
	esREC.SYSDeleted = ""
	esREC.SYSDeletedBy = ""
	esREC.SYSDeletedHost = ""

	esREC.Activity = core.AddActivity_ForUser(esREC.Activity, fmt.Sprintf(Translate("Audit Message", "Cloned from %s (%s)"), origName, esREC.EstimationStateID), userID)

	EstimationSession_StoreSystem(esREC)

	_, featureList, _ := Feature_Active_ByEstimationSession_GetList(esID)
	for _, feature := range featureList {
		err := Feature_Clone(feature, cloneID, userID)
		if err != nil {
			logs.Panic("Cannot Clone Feature {"+feature.FeatureID+"}", err)
			return cloneID, err
		}
	}
	return cloneID, nil
}
