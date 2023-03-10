package dao

import (
	"strconv"
	"time"

	core "github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
	"golang.org/x/exp/slices"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/estimationsession.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : EstimationSession (estimationsession)
// Endpoint 	        : EstimationSession (EstimationSessionID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 03/01/2023 at 19:57:49
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

func EstimationSession_ObjectValidation_impl(iAction string, iId string, iRec dm.EstimationSession) (dm.EstimationSession, string, error) {
	logs.Callout("EstimationSession", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("EstimationSession" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}

// If there are fields below, create the methods in adaptor\estimationsession_impl.go
// START - GET API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
func EstimationSession_EstimationStateID_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_EstimationStateID_scrn, PUT, rec.EstimationSessionID)
	// Check the status of the Estimation
	status := rec.EstimationStateID
	// Get EstimationState from DB
	_, est, _ := EstimationState_GetByCode(status)
	if est.Notify == core.TRUE {
		_, err := StatusChangeMessage(est.Name, status, rec, usr)
		if err != nil {
			return fieldval, err
		}
	}
	return fieldval, nil
}

func EstimationSession_Approver_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_Approver_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_Origin_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_Origin_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_OriginStateID_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginStateID_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_OriginState_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginState_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_OriginDocTypeID_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginDocTypeID_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_OriginDocType_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginDocType_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_OriginCode_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginCode_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_OriginName_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginName_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_OriginRate_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginRate_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectProfileID_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectProfileID_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectProfile_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectProfile_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectDefaultReleases_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectDefaultReleases_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectDefaultReleaseHours_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectDefaultReleaseHours_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectBlendedRate_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectBlendedRate_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectStateID_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectStateID_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectState_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectState_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectName_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectName_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectStartDate_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectStartDate_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectEndDate_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectEndDate_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProfileSupportUpliftPerc_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProfileSupportUpliftPerc_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_CCY_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_CCY_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_CCYCode_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_CCYCode_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_EffortTotal_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_EffortTotal_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_FreshDeskURI_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_FreshDeskURI_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ADOURI_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ADOURI_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_NoActiveFeatures_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_NoActiveFeatures_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}

// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - GET API/Callout
//
// START - PUT API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
func EstimationSession_EstimationStateID_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_EstimationStateID_scrn, GET, rec.EstimationSessionID)
	return rec.EstimationStateID
}
func EstimationSession_Approver_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_Approver_scrn, GET, rec.EstimationSessionID)
	return rec.Approver
}
func EstimationSession_Origin_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_Origin_scrn, GET, rec.EstimationSessionID)
	if rec.Origin != "" {
		return rec.Origin
	}
	if rec.ProjectID != "" {
		project := CacheRead(dm.Project_Name, rec.ProjectID).(dm.Project)
		origin := CacheRead(dm.Origin_Name, project.OriginID).(dm.Origin)
		return origin.OriginID
	}
	return rec.Origin
}
func EstimationSession_OriginStateID_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginStateID_scrn, GET, rec.EstimationSessionID)
	return rec.OriginStateID
}
func EstimationSession_OriginState_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginState_scrn, GET, rec.EstimationSessionID)
	return rec.OriginState
}
func EstimationSession_OriginDocTypeID_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginDocTypeID_scrn, GET, rec.EstimationSessionID)
	return rec.OriginDocTypeID
}
func EstimationSession_OriginDocType_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginDocType_scrn, GET, rec.EstimationSessionID)
	return rec.OriginDocType
}
func EstimationSession_OriginCode_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginCode_scrn, GET, rec.EstimationSessionID)
	if rec.OriginCode != "" {
		return rec.OriginCode
	}
	if rec.ProjectID != "" {
		project := CacheRead(dm.Project_Name, rec.ProjectID).(dm.Project)
		originID := project.OriginID
		origRec := CacheRead(dm.Origin_Name, originID).(dm.Origin)
		return origRec.Code
	}
	return rec.OriginCode
}
func EstimationSession_OriginName_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginName_scrn, GET, rec.EstimationSessionID)
	//fmt.Printf("rec: %v\n", rec)
	if rec.OriginName != "" {
		return rec.OriginName
	}
	if rec.ProjectID != "" {
		project := CacheRead(dm.Project_Name, rec.ProjectID).(dm.Project)
		originID := project.OriginID
		origRec := CacheRead(dm.Origin_Name, originID).(dm.Origin)
		return origRec.FullName
	}
	return ""
}
func EstimationSession_OriginRate_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginRate_scrn, GET, rec.EstimationSessionID)
	return rec.OriginRate
}
func EstimationSession_ProjectProfileID_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectProfileID_scrn, GET, rec.EstimationSessionID)
	return rec.ProjectProfileID
}
func EstimationSession_ProjectProfile_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectProfile_scrn, GET, rec.EstimationSessionID)
	return rec.ProjectProfile
}
func EstimationSession_ProjectDefaultReleases_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectDefaultReleases_scrn, GET, rec.EstimationSessionID)
	return rec.ProjectDefaultReleases
}
func EstimationSession_ProjectDefaultReleaseHours_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectDefaultReleaseHours_scrn, GET, rec.EstimationSessionID)
	return rec.ProjectDefaultReleaseHours
}
func EstimationSession_ProjectBlendedRate_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectBlendedRate_scrn, GET, rec.EstimationSessionID)
	return rec.ProjectBlendedRate
}
func EstimationSession_ProjectStateID_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectStateID_scrn, GET, rec.EstimationSessionID)
	return rec.ProjectStateID
}
func EstimationSession_ProjectState_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectState_scrn, GET, rec.EstimationSessionID)
	return rec.ProjectState
}
func EstimationSession_ProjectName_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectName_scrn, GET, rec.EstimationSessionID)
	if rec.ProjectName != "" {
		return rec.ProjectName
	}
	if rec.ProjectID != "" {
		projID := rec.ProjectID
		projRec := CacheRead(dm.Project_Name, projID).(dm.Project)
		return projRec.Name
	}
	return ""
}
func EstimationSession_ProjectStartDate_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectStartDate_scrn, GET, rec.EstimationSessionID)
	return rec.ProjectStartDate
}
func EstimationSession_ProjectEndDate_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectEndDate_scrn, GET, rec.EstimationSessionID)
	return rec.ProjectEndDate
}
func EstimationSession_ProfileSupportUpliftPerc_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ProfileSupportUpliftPerc_scrn, GET, rec.EstimationSessionID)
	return rec.ProfileSupportUpliftPerc
}
func EstimationSession_CCY_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_CCY_scrn, GET, rec.EstimationSessionID)
	return rec.CCY
}
func EstimationSession_CCYCode_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_CCYCode_scrn, GET, rec.EstimationSessionID)
	return rec.CCYCode
}
func EstimationSession_EffortTotal_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_EffortTotal_scrn, GET, rec.EstimationSessionID)
	return rec.EffortTotal
}
func EstimationSession_FreshDeskURI_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "FreshDeskURI", GET, rec.EstimationSessionID)

	rtn := ""
	if rec.FreshdeskID != "" {
		stub, _ := Data_Get("System", "FreshDesk", dm.Data_Category_URI)
		rtn = core.ReplaceWildcard(stub, "ID", rec.FreshdeskID)
	}

	return rtn
}
func EstimationSession_ADOURI_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ADOURI_scrn, GET, rec.EstimationSessionID)

	rtn := ""
	if rec.AdoID != "" {
		stub, _ := Data_Get("System", "ADO", dm.Data_Category_URI)
		rtn = core.ReplaceWildcard(stub, "ID", rec.AdoID)
	}

	return rtn
}
func EstimationSession_NoActiveFeatures_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_NoActiveFeatures_scrn, GET, rec.EstimationSessionID)
	return rec.NoActiveFeatures
}

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - PUT API/Callout

// START - Validation API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
//

// EstimationSession_EstimationSessionID_validate_impl provides validation/actions for EstimationSessionID
func EstimationSession_EstimationSessionID_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_EstimationSessionID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_EstimationStateID_impl provides validation/actions for EstimationStateID
func EstimationSession_EstimationStateID_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_EstimationStateID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// EstimationSession_AdoID_validate_impl provides validation/actions for AdoID
func EstimationSession_AdoID_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_AdoID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// EstimationSession_FreshdeskID_validate_impl provides validation/actions for FreshdeskID
func EstimationSession_FreshdeskID_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_FreshdeskID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// EstimationSession_TrackerID_validate_impl provides validation/actions for TrackerID
func EstimationSession_TrackerID_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_TrackerID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_Approver_impl provides validation/actions for Approver
func EstimationSession_Approver_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_Approver_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// EstimationSession_IssueDate_validate_impl provides validation/actions for IssueDate
func EstimationSession_IssueDate_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_IssueDate_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// EstimationSession_ExpiryDate_validate_impl provides validation/actions for ExpiryDate
func EstimationSession_ExpiryDate_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ExpiryDate_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_Origin_impl provides validation/actions for Origin
func EstimationSession_Origin_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_Origin_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_OriginStateID_impl provides validation/actions for OriginStateID
func EstimationSession_OriginStateID_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginStateID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_OriginState_impl provides validation/actions for OriginState
func EstimationSession_OriginState_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginState_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_OriginDocTypeID_impl provides validation/actions for OriginDocTypeID
func EstimationSession_OriginDocTypeID_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginDocTypeID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_OriginDocType_impl provides validation/actions for OriginDocType
func EstimationSession_OriginDocType_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginDocType_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_OriginCode_impl provides validation/actions for OriginCode
func EstimationSession_OriginCode_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginCode_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_OriginName_impl provides validation/actions for OriginName
func EstimationSession_OriginName_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginName_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_OriginRate_impl provides validation/actions for OriginRate
func EstimationSession_OriginRate_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_OriginRate_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectProfileID_impl provides validation/actions for ProjectProfileID
func EstimationSession_ProjectProfileID_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectProfileID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectProfile_impl provides validation/actions for ProjectProfile
func EstimationSession_ProjectProfile_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectProfile_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectDefaultReleases_impl provides validation/actions for ProjectDefaultReleases
func EstimationSession_ProjectDefaultReleases_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectDefaultReleases_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectDefaultReleaseHours_impl provides validation/actions for ProjectDefaultReleaseHours
func EstimationSession_ProjectDefaultReleaseHours_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectDefaultReleaseHours_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectBlendedRate_impl provides validation/actions for ProjectBlendedRate
func EstimationSession_ProjectBlendedRate_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectBlendedRate_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectStateID_impl provides validation/actions for ProjectStateID
func EstimationSession_ProjectStateID_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectStateID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectState_impl provides validation/actions for ProjectState
func EstimationSession_ProjectState_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectState_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectName_impl provides validation/actions for ProjectName
func EstimationSession_ProjectName_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectName_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectStartDate_impl provides validation/actions for ProjectStartDate
func EstimationSession_ProjectStartDate_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectStartDate_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectEndDate_impl provides validation/actions for ProjectEndDate
func EstimationSession_ProjectEndDate_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProjectEndDate_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProfileSupportUpliftPerc_impl provides validation/actions for ProfileSupportUpliftPerc
func EstimationSession_ProfileSupportUpliftPerc_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ProfileSupportUpliftPerc_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_CCY_impl provides validation/actions for CCY
func EstimationSession_CCY_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_CCY_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_CCYCode_impl provides validation/actions for CCYCode
func EstimationSession_CCYCode_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_CCYCode_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_EffortTotal_impl provides validation/actions for EffortTotal
func EstimationSession_EffortTotal_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_EffortTotal_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_FreshDeskURI_impl provides validation/actions for FreshDeskURI
func EstimationSession_FreshDeskURI_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_FreshDeskURI_scrn, VAL+"-"+iAction, iId)
	rtn := ""
	//logs.Warning("Generate FreshDesk URI")
	if iRec.FreshdeskID != "" {
		stub, _ := Data_Get("System", "FreshDesk", dm.Data_Category_URI)
		rtn = core.ReplaceWildcard(stub, "ID", iRec.FreshdeskID)
	}
	return rtn, fP
}

// EstimationSession_ADOURI_impl provides validation/actions for ADOURI
func EstimationSession_ADOURI_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ADOURI_scrn, VAL+"-"+iAction, iId)
	//logs.Warning("Generate ADO URI")
	rtn := ""
	if iRec.AdoID != "" {
		stub, _ := Data_Get("System", "ADO", dm.Data_Category_URI)
		rtn = core.ReplaceWildcard(stub, "ID", iRec.AdoID)
	}
	return rtn, fP
}

// EstimationSession_NoActiveFeatures_impl provides validation/actions for NoActiveFeatures
func EstimationSession_NoActiveFeatures_validate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_NoActiveFeatures_scrn, VAL+"-"+iAction, iId)
	return "999", fP
}

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout

func EstimationSession_IssueDate_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_IssueDate_scrn, PUT, rec.EstimationSessionID)
	issuedState, _ := Data_GetString("Quote", "Issued", dm.Data_Category_StateRule)
	if fieldval == "" && rec.ExpiryDate == "" && rec.EstimationStateID == issuedState {
		// Get Todays Date yyyy-mm-dd
		today := time.Now().Format(core.DATEFORMAT)
		return today, nil
	}

	return fieldval, nil
}

func EstimationSession_IssueDate_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_IssueDate_scrn, GET, rec.EstimationSessionID)
	return rec.IssueDate
}

// EstimationSession_IssueDate_impl provides validation/actions for IssueDate
func EstimationSession_IssueDate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_IssueDate_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

func EstimationSession_ExpiryDate_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_ExpiryDate_scrn, PUT, rec.EstimationSessionID)

	issuedState, _ := Data_GetString("Quote", "Issued", dm.Data_Category_StateRule)
	expiredState, _ := Data_GetString("Quote", "Expired", dm.Data_Category_StateRule)

	objectName := Translate("ObjectName", "EstimationSession")

	switch rec.EstimationStateID {
	case issuedState:
		// Get Expiroty Date 30 days from today
		if fieldval == "" {
			expPeriod, _ := Data_GetInt(objectName, "Quote_Expiry_Notification_Period", dm.Data_Category_Setting)
			logs.Information("Expiry Period: ", strconv.Itoa(expPeriod))
			expiry := time.Now().AddDate(0, 0, expPeriod).Format(core.DATEFORMAT)
			return expiry, nil
		}
	case expiredState:
		logs.Information("Quote is expired", " cannot change Expiry Date")
		return emailRelatedResources(rec)
	default:
	}

	return fieldval, nil
}

func EstimationSession_ExpiryDate_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_ExpiryDate_scrn, GET, rec.EstimationSessionID)
	return rec.ExpiryDate
}

// EstimationSession_ExpiryDate_impl provides validation/actions for ExpiryDate
func EstimationSession_ExpiryDate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_ExpiryDate_scrn, VAL+"-"+iAction, iId)

	return iValue, fP
}

// ----------------------------------------------------------------
// BEGIN EstimationSession_EstimationSessionID
// BEGIN EstimationSession_EstimationSessionID
// BEGIN EstimationSession_EstimationSessionID
// ----------------------------------------------------------------
// EstimationSession_EstimationSessionID_OnStore_impl provides the implementation for the callout
func EstimationSession_EstimationSessionID_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_EstimationSessionID_scrn, PUT, rec.EstimationSessionID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// EstimationSession_EstimationSessionID_OnFetch_impl provides the implementation for the callout
func EstimationSession_EstimationSessionID_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_EstimationSessionID_scrn, GET, rec.EstimationSessionID)
	return rec.EstimationSessionID
}

// ----------------------------------------------------------------
// EstimationSession_EstimationSessionID_impl provides validation/actions for EstimationSessionID
func EstimationSession_EstimationSessionID_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_EstimationSessionID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   EstimationSession_EstimationSessionID
// END   EstimationSession_EstimationSessionID
// END   EstimationSession_EstimationSessionID
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN EstimationSession_AdoID
// BEGIN EstimationSession_AdoID
// BEGIN EstimationSession_AdoID
// ----------------------------------------------------------------
// EstimationSession_AdoID_OnStore_impl provides the implementation for the callout
func EstimationSession_AdoID_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_AdoID_scrn, PUT, rec.EstimationSessionID)
	Indexer_Put("EstimationSession", dm.EstimationSession_AdoID_scrn, rec.EstimationSessionID, fieldval)

	return fieldval, nil
}

// ----------------------------------------------------------------
// EstimationSession_AdoID_OnFetch_impl provides the implementation for the callout
func EstimationSession_AdoID_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_AdoID_scrn, GET, rec.EstimationSessionID)
	return rec.AdoID
}

// ----------------------------------------------------------------
// EstimationSession_AdoID_impl provides validation/actions for AdoID
func EstimationSession_AdoID_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_AdoID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   EstimationSession_AdoID
// END   EstimationSession_AdoID
// END   EstimationSession_AdoID
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN EstimationSession_FreshdeskID
// BEGIN EstimationSession_FreshdeskID
// BEGIN EstimationSession_FreshdeskID
// ----------------------------------------------------------------
// EstimationSession_FreshdeskID_OnStore_impl provides the implementation for the callout
func EstimationSession_FreshdeskID_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_FreshdeskID_scrn, PUT, rec.EstimationSessionID)
	Indexer_Put("EstimationSession", dm.EstimationSession_FreshdeskID_scrn, rec.EstimationSessionID, fieldval)

	return fieldval, nil
}

// ----------------------------------------------------------------
// EstimationSession_FreshdeskID_OnFetch_impl provides the implementation for the callout
func EstimationSession_FreshdeskID_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_FreshdeskID_scrn, GET, rec.EstimationSessionID)
	return rec.FreshdeskID
}

// ----------------------------------------------------------------
// EstimationSession_FreshdeskID_impl provides validation/actions for FreshdeskID
func EstimationSession_FreshdeskID_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_FreshdeskID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   EstimationSession_FreshdeskID
// END   EstimationSession_FreshdeskID
// END   EstimationSession_FreshdeskID
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN EstimationSession_TrackerID
// BEGIN EstimationSession_TrackerID
// BEGIN EstimationSession_TrackerID
// ----------------------------------------------------------------
// EstimationSession_TrackerID_OnStore_impl provides the implementation for the callout
func EstimationSession_TrackerID_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_TrackerID_scrn, PUT, rec.EstimationSessionID)

	Indexer_Put("EstimationSession", dm.EstimationSession_TrackerID_scrn, rec.EstimationSessionID, fieldval)

	//spew.Dump(rec)
	state := rec.EstimationStateID
	objectName := Translate("ObjectName", "EstimationSession")
	assignRSCstatus, err := Data_GetArray(objectName, "Assign_RSC_No_OnState", dm.Data_Category_State)
	if err != nil {
		//	logs.Information("Assign_RSC_States", err.Error())
		assignRSCstatus = []string{"WRIT"}
	}

	//fmt.Printf("assignRSCstatus: %v\n", assignRSCstatus)

	testVal := rec.TrackerID

	//logs.Information("fieldval", fieldval)
	//logs.Information("testVal", testVal)
	//logs.Information("state", state)
	//fmt.Printf("assignRSCstatus: %v\n", assignRSCstatus)
	if slices.Contains(assignRSCstatus, state) && (testVal == "" || len(testVal) == 0) {
		logs.Information("Assign_RSC_States", "Assigning RSC")
		_, proj, err := Project_GetByID(rec.ProjectID)
		if err != nil {
			return "", err
		}
		_, origin, errorg := Origin_GetByCode(proj.OriginID)
		if errorg != nil {
			return "", errorg
		}
		docType := origin.DocTypeID
		_, newIDString, err := Data_NextSEQ(docType)
		if err != nil {
			return "", err
		}
		fieldval = newIDString
	}

	return fieldval, nil
}

// ----------------------------------------------------------------
// EstimationSession_TrackerID_OnFetch_impl provides the implementation for the callout
func EstimationSession_TrackerID_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", dm.EstimationSession_TrackerID_scrn, GET, rec.EstimationSessionID)
	return rec.TrackerID
}

// ----------------------------------------------------------------
// EstimationSession_TrackerID_impl provides validation/actions for TrackerID
func EstimationSession_TrackerID_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", dm.EstimationSession_TrackerID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   EstimationSession_TrackerID
// END   EstimationSession_TrackerID
// END   EstimationSession_TrackerID
// ----------------------------------------------------------------
