package dao

import (
	core "github.com/mt1976/ebEstimates/core"
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
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/02/2023 at 09:57:37
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------
// The following functions should be created in data_impl.go
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Data_ObjectValidation_impl provides Record/Object level validation for Data
func Data_ObjectValidation_impl(iAction string, iId string, iRec dm.Data) (dm.Data, string, error) {
	logs.Callout("Data", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:
		iRec.Value = core.SQL_Escape(iRec.Value)
	case GET:
		iRec.Value = core.SQL_UnEscape(iRec.Value)
	default:
		logs.Warning("Data" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}

//
// ----------------------------------------------------------------
// These are the the default implementations, which do nothing

// ----------------------------------------------------------------
// BEGIN Data_DataID
// BEGIN Data_DataID
// BEGIN Data_DataID
// ----------------------------------------------------------------
// Data_DataID_OnStore_impl provides the implementation for the callout

// ----------------------------------------------------------------
// Data_DataID_validate_impl provides validation/actions for DataID
func Data_DataID_validate_impl(iAction string, iId string, iValue string, iRec dm.Data, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Data", dm.Data_DataID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   Data_DataID
// END   Data_DataID
// END   Data_DataID
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------
