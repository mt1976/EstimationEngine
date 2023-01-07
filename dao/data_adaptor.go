package dao

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/data.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : Data (data)
// Endpoint 	        : Data (DataID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 04/01/2023 at 13:08:46
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in data_impl.go

func Data_GetList_impl() (int, []dm.Data, error)        { return 0, nil, nil }
func Data_GetByID_impl(id string) (int, dm.Data, error) { return 0, dm.Data{}, nil }

func Data_NewID_impl(rec dm.Data) string { return rec.DataID }
func Data_Delete_impl(id string) error   { return nil }

func Data_Update_impl(id string, rec dm.Data, usr string) error { return nil }

// If there are fields below, create the methods in adaptor\data_impl.go
// START - GET API/Callout
// Dynamically generated 04/01/2023 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 04/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - GET API/Callout
//
// START - PUT API/Callout
// Dynamically generated 04/01/2023 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 04/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - PUT API/Callout

// START - Validation API/Callout
// Dynamically generated 04/01/2023 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 04/01/2023 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout

func Data_ObjectValidation_impl(iAction string, iId string, iRec dm.Data) (dm.Data, string, error) {
	logs.Callout("Data", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("Data" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}

func Data_DataID_OnStore_impl(fieldval string, rec dm.Data, usr string) (string, error) {
	logs.Callout("Data", dm.Data_DataID_scrn, PUT, rec.DataID)
	fieldval = data_BuildID(rec.Class, rec.Field)
	return fieldval, nil
}

func Data_DataID_OnFetch_impl(rec dm.Data) string {
	logs.Callout("Data", dm.Data_DataID_scrn, GET, rec.DataID)
	return rec.DataID
}

// Data_DataID_impl provides validation/actions for DataID
func Data_DataID_impl(iAction string, iId string, iValue string, iRec dm.Data, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Data", dm.Data_DataID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}
