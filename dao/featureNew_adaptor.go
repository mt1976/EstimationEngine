package dao

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/featurenew.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : FeatureNew (featurenew)
// Endpoint 	        : FeatureNew (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 03/01/2023 at 19:18:09
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in featurenew_impl.go

func FeatureNew_GetList_impl(filter string) (int, []dm.FeatureNew, error) { return 0, nil, nil }
func FeatureNew_GetByID_impl(id string) (int, dm.FeatureNew, error)       { return 0, dm.FeatureNew{}, nil }

func FeatureNew_NewID_impl(rec dm.FeatureNew) string { return rec.ID }
func FeatureNew_Delete_impl(id string) error         { return nil }

func FeatureNew_Update_impl(id string, rec dm.FeatureNew, usr string) error { return nil }

// If there are fields below, create the methods in adaptor\featurenew_impl.go
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
