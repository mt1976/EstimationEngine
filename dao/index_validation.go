package dao

import (
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/index.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : Index (index)
// Endpoint 	        : Index (IndexID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/02/2023 at 10:44:44
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------
// The following functions should be created in index_impl.go
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Index_ObjectValidation_impl provides Record/Object level validation for Index
func Index_ObjectValidation_impl(iAction string, iId string, iRec dm.Index) (dm.Index, string, error) {
	logs.Callout("Index", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("Index" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}

//
// ----------------------------------------------------------------
// These are the the default implementations, which do nothing

// ----------------------------------------------------------------
// BEGIN Index_IndexID
// BEGIN Index_IndexID
// BEGIN Index_IndexID
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Index_IndexID_validate_impl provides validation/actions for IndexID
func Index_IndexID_validate_impl(iAction string, iId string, iValue string, iRec dm.Index, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Index", dm.Index_IndexID_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   Index_IndexID
// END   Index_IndexID
// END   Index_IndexID
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// BEGIN Index_Link
// BEGIN Index_Link
// BEGIN Index_Link
// ----------------------------------------------------------------
// Index_Link_OnStore_impl provides the implementation for the callout

// ----------------------------------------------------------------
// Index_Link_validate_impl provides validation/actions for Link
func Index_Link_validate_impl(iAction string, iId string, iValue string, iRec dm.Index, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Index", dm.Index_Link_scrn, VAL+"-"+iAction, iId)
	return iValue, fP
}

// ----------------------------------------------------------------
// END   Index_Link
// END   Index_Link
// END   Index_Link
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------
