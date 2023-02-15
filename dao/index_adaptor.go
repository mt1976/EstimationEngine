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
// Date & Time		    : 23/01/2023 at 11:31:28
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------
// The following functions should be created in index_impl.go
// ----------------------------------------------------------------

// ----------------------------------------------------------------

// ----------------------------------------------------------------

//
// ----------------------------------------------------------------
// If there are fields below, create the methods in adaptor\index_impl.go
// These are the the default implementations, which do nothing

// ----------------------------------------------------------------
// BEGIN Index_IndexID
// BEGIN Index_IndexID
// BEGIN Index_IndexID
// ----------------------------------------------------------------
// Index_IndexID_OnStore_impl provides the implementation for the callout
func Index_IndexID_OnStore_impl(fieldval string, rec dm.Index, usr string) (string, error) {
	logs.Callout("Index", dm.Index_IndexID_scrn, PUT, rec.IndexID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Index_IndexID_OnFetch_impl provides the implementation for the callout
func Index_IndexID_OnFetch_impl(rec dm.Index) string {
	logs.Callout("Index", dm.Index_IndexID_scrn, GET, rec.IndexID)
	return rec.IndexID
}

// ----------------------------------------------------------------
// Index_IndexID_impl provides validation/actions for IndexID
func Index_IndexID_impl(iAction string, iId string, iValue string, iRec dm.Index, fP dm.FieldProperties) (string, dm.FieldProperties) {
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
func Index_Link_OnStore_impl(fieldval string, rec dm.Index, usr string) (string, error) {
	logs.Callout("Index", dm.Index_Link_scrn, PUT, rec.IndexID)
	return fieldval, nil
}

// ----------------------------------------------------------------
// Index_Link_OnFetch_impl provides the implementation for the callout
func Index_Link_OnFetch_impl(rec dm.Index) string {
	logs.Callout("Index", dm.Index_Link_scrn, GET, rec.IndexID)
	return rec.Link
}

// ----------------------------------------------------------------
// Index_Link_impl provides validation/actions for Link
func Index_Link_impl(iAction string, iId string, iValue string, iRec dm.Index, fP dm.FieldProperties) (string, dm.FieldProperties) {
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
