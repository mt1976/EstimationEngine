package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/estimationstate.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : EstimationState (estimationstate)
// Endpoint 	        : EstimationState (EstimationStateID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 03/01/2023 at 19:18:08
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"github.com/mt1976/ebEstimates/das"

	dm "github.com/mt1976/ebEstimates/datamodel"
)

// EstimationState_GetByCode() returns a single EstimationState record
func EstimationState_GetByCode(id string) (int, dm.EstimationState, error) {

	tsql := EstimationState_SQLbase
	tsql = tsql + " " + das.WHERE + dm.EstimationState_Code_sql + das.EQ + das.ID(id)
	_, _, estimationstateItem, _ := estimationstate_Fetch(tsql)

	// START
	// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 03/01/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
	return 1, estimationstateItem, nil
}
