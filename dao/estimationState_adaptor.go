package dao

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/estimationstate.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : EstimationState (estimationstate)
// Endpoint 	        : EstimationState (EstimationStateID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 03/01/2023 at 19:18:08
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in estimationstate_impl.go

// If there are fields below, create the methods in adaptor\estimationstate_impl.go
// START - GET API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - GET API/Callout
//
// START - PUT API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - PUT API/Callout

// START - Validation API/Callout
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout

func EstimationState_ObjectValidation_impl(iAction string, iId string, iRec dm.EstimationState) (dm.EstimationState, string, error) {
	logs.Callout("EstimationState", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("EstimationState" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}

// ----------------------------------------------------------------
// BEGIN EstimationSession_TrackerID
// BEGIN EstimationSession_TrackerID
// BEGIN EstimationSession_TrackerID
// ----------------------------------------------------------------
// EstimationSession_TrackerID_OnStore_impl provides the implementation for the callout
func EstimationSession_TrackerID_OnStore_impl(fieldval string, rec dm.EstimationSession, usr string) (string, error) {
	logs.Callout("EstimationSession", dm.EstimationSession_TrackerID_scrn, PUT, rec.EstimationSessionID)
	state := rec.EstimationStateID

	if state == "WRIT" && fieldval == "" {
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
