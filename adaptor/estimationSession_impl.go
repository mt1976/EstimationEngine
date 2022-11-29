package adaptor

import (
	"github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/estimationsession.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : EstimationSession (estimationsession)
// Endpoint 	        : EstimationSession (EstimationSessionID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 28/11/2022 at 18:56:33
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in estimationsession_impl.go

// If there are fields below, create the methods in adaptor\estimationsession_impl.go
// START - GET API/Callout
// Dynamically generated 28/11/2022 by matttownsend (Matt Townsend) on silicon.local
func EstimationSession_Origin_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "Origin", PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_OriginStateID_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "OriginStateID", PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_OriginState_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "OriginState", PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_OriginDocTypeID_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "OriginDocTypeID", PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_OriginDocType_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "OriginDocType", PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_OriginCode_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "OriginCode", PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_OriginName_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "OriginName", PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_OriginRate_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "OriginRate", PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectProfileID_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "ProjectProfileID", PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectProfile_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "ProjectProfile", PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectDefaultReleases_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "ProjectDefaultReleases", PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectDefaultReleaseHours_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "ProjectDefaultReleaseHours", PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectBlendedRate_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "ProjectBlendedRate", PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectStateID_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "ProjectStateID", PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectState_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "ProjectState", PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectName_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "ProjectName", PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectStartDate_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "ProjectStartDate", PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProjectEndDate_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "ProjectEndDate", PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ProfileSupportUpliftPerc_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "ProfileSupportUpliftPerc", PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_CCY_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "CCY", PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_CCYCode_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "CCYCode", PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_EffortTotal_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "EffortTotal", PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_FreshDeskURI_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "FreshDeskURI", PUT, rec.EstimationSessionID)
	return fieldval, nil
}
func EstimationSession_ADOURI_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", "ADOURI", PUT, rec.EstimationSessionID)
	return fieldval, nil
}

// Dynamically generated 28/11/2022 by matttownsend (Matt Townsend) on silicon.local
// END - GET API/Callout
//
// START - PUT API/Callout
// Dynamically generated 28/11/2022 by matttownsend (Matt Townsend) on silicon.local
func EstimationSession_Origin_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "Origin", GET, rec.EstimationSessionID)
	return rec.Origin
}
func EstimationSession_OriginStateID_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "OriginStateID", GET, rec.EstimationSessionID)
	return rec.OriginStateID
}
func EstimationSession_OriginState_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "OriginState", GET, rec.EstimationSessionID)
	return rec.OriginState
}
func EstimationSession_OriginDocTypeID_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "OriginDocTypeID", GET, rec.EstimationSessionID)
	return rec.OriginDocTypeID
}
func EstimationSession_OriginDocType_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "OriginDocType", GET, rec.EstimationSessionID)
	return rec.OriginDocType
}
func EstimationSession_OriginCode_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "OriginCode", GET, rec.EstimationSessionID)
	return rec.OriginCode
}
func EstimationSession_OriginName_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "OriginName", GET, rec.EstimationSessionID)
	return rec.OriginName
}
func EstimationSession_OriginRate_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "OriginRate", GET, rec.EstimationSessionID)
	return rec.OriginRate
}
func EstimationSession_ProjectProfileID_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "ProjectProfileID", GET, rec.EstimationSessionID)
	return rec.ProjectProfileID
}
func EstimationSession_ProjectProfile_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "ProjectProfile", GET, rec.EstimationSessionID)
	return rec.ProjectProfile
}
func EstimationSession_ProjectDefaultReleases_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "ProjectDefaultReleases", GET, rec.EstimationSessionID)
	return rec.ProjectDefaultReleases
}
func EstimationSession_ProjectDefaultReleaseHours_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "ProjectDefaultReleaseHours", GET, rec.EstimationSessionID)
	return rec.ProjectDefaultReleaseHours
}
func EstimationSession_ProjectBlendedRate_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "ProjectBlendedRate", GET, rec.EstimationSessionID)
	return rec.ProjectBlendedRate
}
func EstimationSession_ProjectStateID_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "ProjectStateID", GET, rec.EstimationSessionID)
	return rec.ProjectStateID
}
func EstimationSession_ProjectState_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "ProjectState", GET, rec.EstimationSessionID)
	return rec.ProjectState
}
func EstimationSession_ProjectName_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "ProjectName", GET, rec.EstimationSessionID)
	return rec.ProjectName
}
func EstimationSession_ProjectStartDate_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "ProjectStartDate", GET, rec.EstimationSessionID)
	return rec.ProjectStartDate
}
func EstimationSession_ProjectEndDate_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "ProjectEndDate", GET, rec.EstimationSessionID)
	return rec.ProjectEndDate
}
func EstimationSession_ProfileSupportUpliftPerc_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "ProfileSupportUpliftPerc", GET, rec.EstimationSessionID)
	return rec.ProfileSupportUpliftPerc
}
func EstimationSession_CCY_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "CCY", GET, rec.EstimationSessionID)
	return rec.CCY
}
func EstimationSession_CCYCode_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "CCYCode", GET, rec.EstimationSessionID)
	return rec.CCYCode
}
func EstimationSession_EffortTotal_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "EffortTotal", GET, rec.EstimationSessionID)
	return rec.EffortTotal
}
func EstimationSession_FreshDeskURI_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "FreshDeskURI", GET, rec.EstimationSessionID)
	rtn := core.ApplicationProperties["freshdeskticketuri"]
	rtn = core.ReplaceWildcard(rtn, "ID", rec.FreshdeskID)
	//logs.Result("FreshDeskURI", rtn)
	return rtn
}
func EstimationSession_ADOURI_OnFetch_impl(rec dm.EstimationSession) string {
	logs.Callout("EstimationSession", "ADOURI", GET, rec.EstimationSessionID)
	rtn := core.ApplicationProperties["adoticketuri"]
	rtn = core.ReplaceWildcard(rtn, "ID", rec.AdoID)
	return rtn
}

//
// Dynamically generated 28/11/2022 by matttownsend (Matt Townsend) on silicon.local
// END - PUT API/Callout

// START - Validation API/Callout
// Dynamically generated 28/11/2022 by matttownsend (Matt Townsend) on silicon.local
//
// EstimationSession_Origin_impl provides validation/actions for Origin
func EstimationSession_Origin_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "Origin", VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_OriginStateID_impl provides validation/actions for OriginStateID
func EstimationSession_OriginStateID_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "OriginStateID", VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_OriginState_impl provides validation/actions for OriginState
func EstimationSession_OriginState_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "OriginState", VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_OriginDocTypeID_impl provides validation/actions for OriginDocTypeID
func EstimationSession_OriginDocTypeID_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "OriginDocTypeID", VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_OriginDocType_impl provides validation/actions for OriginDocType
func EstimationSession_OriginDocType_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "OriginDocType", VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_OriginCode_impl provides validation/actions for OriginCode
func EstimationSession_OriginCode_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "OriginCode", VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_OriginName_impl provides validation/actions for OriginName
func EstimationSession_OriginName_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "OriginName", VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_OriginRate_impl provides validation/actions for OriginRate
func EstimationSession_OriginRate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "OriginRate", VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectProfileID_impl provides validation/actions for ProjectProfileID
func EstimationSession_ProjectProfileID_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "ProjectProfileID", VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectProfile_impl provides validation/actions for ProjectProfile
func EstimationSession_ProjectProfile_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "ProjectProfile", VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectDefaultReleases_impl provides validation/actions for ProjectDefaultReleases
func EstimationSession_ProjectDefaultReleases_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "ProjectDefaultReleases", VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectDefaultReleaseHours_impl provides validation/actions for ProjectDefaultReleaseHours
func EstimationSession_ProjectDefaultReleaseHours_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "ProjectDefaultReleaseHours", VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectBlendedRate_impl provides validation/actions for ProjectBlendedRate
func EstimationSession_ProjectBlendedRate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "ProjectBlendedRate", VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectStateID_impl provides validation/actions for ProjectStateID
func EstimationSession_ProjectStateID_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "ProjectStateID", VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectState_impl provides validation/actions for ProjectState
func EstimationSession_ProjectState_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "ProjectState", VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectName_impl provides validation/actions for ProjectName
func EstimationSession_ProjectName_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "ProjectName", VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectStartDate_impl provides validation/actions for ProjectStartDate
func EstimationSession_ProjectStartDate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "ProjectStartDate", VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProjectEndDate_impl provides validation/actions for ProjectEndDate
func EstimationSession_ProjectEndDate_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "ProjectEndDate", VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ProfileSupportUpliftPerc_impl provides validation/actions for ProfileSupportUpliftPerc
func EstimationSession_ProfileSupportUpliftPerc_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "ProfileSupportUpliftPerc", VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_CCY_impl provides validation/actions for CCY
func EstimationSession_CCY_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "CCY", VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_CCYCode_impl provides validation/actions for CCYCode
func EstimationSession_CCYCode_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "CCYCode", VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_EffortTotal_impl provides validation/actions for EffortTotal
func EstimationSession_EffortTotal_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "EffortTotal", VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_FreshDeskURI_impl provides validation/actions for FreshDeskURI
func EstimationSession_FreshDeskURI_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "FreshDeskURI", VAL+"-"+iAction, iId)
	return iValue, fP
}

// EstimationSession_ADOURI_impl provides validation/actions for ADOURI
func EstimationSession_ADOURI_impl(iAction string, iId string, iValue string, iRec dm.EstimationSession, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("EstimationSession", "ADOURI", VAL+"-"+iAction, iId)
	return iValue, fP
}

//
// Dynamically generated 28/11/2022 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout

func EstimationSession_ObjectValidation_impl(iAction string, iId string, iRec dm.EstimationSession) (dm.EstimationSession, error, string) {
	logs.Callout("EstimationSession", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("Tmpl_TDate_impl" + " - Invalid Action")
	}
	return iRec, nil, ""
}
