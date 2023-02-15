package dao

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/credentialsaction.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : CredentialsAction (credentialsaction)
// Endpoint 	        : CredentialsAction (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 03/01/2023 at 19:18:08
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in credentialsaction_impl.go

func CredentialsAction_GetList_impl(filter string) (int, []dm.CredentialsAction, error) {
	return 0, nil, nil
}
func CredentialsAction_GetByID_impl(id string) (int, dm.CredentialsAction, error) {
	return 0, dm.CredentialsAction{}, nil
}

func CredentialsAction_NewID_impl(rec dm.CredentialsAction) string { return rec.ID }
func CredentialsAction_Delete_impl(id string) error                { return nil }

func CredentialsAction_Update_impl(id string, rec dm.CredentialsAction, usr string) error { return nil }

// If there are fields below, create the methods in adaptor\credentialsaction_impl.go
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
