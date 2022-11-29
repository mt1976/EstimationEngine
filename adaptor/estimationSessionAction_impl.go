package adaptor

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/estimationsessionaction.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : EstimationSessionAction (estimationsessionaction)
// Endpoint 	        : EstimationSessionAction (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 29/11/2022 at 12:48:16
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in estimationsessionaction_impl.go

func EstimationSessionAction_GetList_impl() (int, []dm.EstimationSessionAction, error) {
	return 0, nil, nil
}
func EstimationSessionAction_GetByID_impl(id string) (int, dm.EstimationSessionAction, error) {
	return 0, dm.EstimationSessionAction{}, nil
}

func EstimationSessionAction_NewID_impl(rec dm.EstimationSessionAction) string { return rec.ID }
func EstimationSessionAction_Delete_impl(id string) error                      { return nil }

func EstimationSessionAction_Update_impl(id string, rec dm.EstimationSessionAction, usr string) error {
	return nil
}

// If there are fields below, create the methods in adaptor\estimationsessionaction_impl.go
// START - GET API/Callout
// Dynamically generated 29/11/2022 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 29/11/2022 by matttownsend (Matt Townsend) on silicon.local
// END - GET API/Callout
//
// START - PUT API/Callout
// Dynamically generated 29/11/2022 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 29/11/2022 by matttownsend (Matt Townsend) on silicon.local
// END - PUT API/Callout

// START - Validation API/Callout
// Dynamically generated 29/11/2022 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 29/11/2022 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout

func EstimationSessionAction_ObjectValidation_impl(iAction string, iId string, iRec dm.EstimationSessionAction) (dm.EstimationSessionAction, error, string) {
	logs.Callout("EstimationSessionAction", "ObjectValidation", VAL+"-"+iAction, iId)
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
