package dao

import (
	core "github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/doctype.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : DocType (doctype)
// Endpoint 	        : DocType (DocTypeID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/02/2023 at 09:39:55
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------
// The following functions should be created in doctype_impl.go
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// DocType_ObjectValidation_impl provides Record/Object level validation for DocType
func DocType_ObjectValidation_impl(iAction string, iId string, iRec dm.DocType) (dm.DocType, string, error) {
	logs.Callout("DocType", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("DocType" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}

//
// ----------------------------------------------------------------
// If there are fields below, create the methods in adaptor\doctype_impl.go
// These are the the default implementations, which do nothing

// ----------------------------------------------------------------
// BEGIN DocType_Code
// BEGIN DocType_Code
// BEGIN DocType_Code
// ----------------------------------------------------------------
// DocType_Code_OnStore_impl provides the implementation for the callout
func DocType_Code_OnStore_impl(fieldval string, rec dm.DocType, usr string) (string, error) {
	logs.Callout("DocType", dm.DocType_Code_scrn, PUT, rec.DocTypeID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// DocType_Code_OnFetch_impl provides the implementation for the callout
func DocType_Code_OnFetch_impl(rec dm.DocType) string {
	logs.Callout("DocType", dm.DocType_Code_scrn, GET, rec.DocTypeID)
	return rec.Code
}

// ----------------------------------------------------------------
// DocType_Code_validate_impl provides validation/actions for Code
func DocType_Code_validate_impl(iAction string, iId string, iValue string, iRec dm.DocType, fP dm.FieldProperties) (string, dm.FieldProperties) {
	//logs.Callout("DocType", dm.DocType_Code_scrn, VAL+"-"+iAction, iId)
	if iAction == PUT || iAction == NEW {
		if len(iValue) == 0 {
			fP = SetFieldError(fP, "is required")
		}
		if len(iValue) > 4 {
			fP = SetFieldError(fP, "must be less than 4 characters")
		}
		if !core.IsLetter(iValue) {
			fP = SetFieldError(fP, "must be letters only A-Z")
		}
		if !core.IsUpper(iValue) {
			fP = SetFieldError(fP, "must be upper case")

		}
	}
	//logs.Information("DocType", dm.DocType_Code_scrn+" - "+iAction+" - "+iId+" - "+iValue)
	//logs.Information("Returning", fP.MsgType+" - "+fP.MsgMessage)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   DocType_Code
// END   DocType_Code
// END   DocType_Code
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------
