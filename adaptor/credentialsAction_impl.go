package adaptor

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
)

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/credentialsaction.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : CredentialsAction (credentialsaction)
// Endpoint 	        : CredentialsAction (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 23/06/2022 at 10:15:54
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in credentialsaction_impl.go

func CredentialsAction_GetList_impl() (int, []dm.CredentialsAction, error) { return 0, nil, nil }
func CredentialsAction_GetByID_impl(id string) (int, dm.CredentialsAction, error) {
	//fmt.Printf("Get By id: %v\n", id)
	return 0, dm.CredentialsAction{}, nil
}

func CredentialsAction_NewID_impl(rec dm.CredentialsAction) string { return rec.ID }
func CredentialsAction_Delete_impl(id string) error                { return nil }

func CredentialsAction_Update_impl(id string, rec dm.CredentialsAction, usr string) error { return nil }

// If there are fields below, create the methods in adaptor\credentialsaction_impl.go
