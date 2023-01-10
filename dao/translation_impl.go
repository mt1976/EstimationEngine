package dao

import (
	"strings"

	core "github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
	"github.com/mt1976/ebEstimates/logs"
)

// Returns a valid translation store ID
func Translation_BuildID(class string, message string) string {
	var translationStoreID string
	message = strings.ReplaceAll(message, " ", "-")
	message = core.ReplaceSpecialChars(message)
	translationStoreID = class + core.IDSep + message

	return core.EncodeString(translationStoreID)
}

// Translation_Lookup  will translate a string from its source version to user-defined version.
func Translate(class string, message string) string {
	translation_ID := Translation_BuildID(class, message)
	//logs.Accessing("translationStoreID" + translation_ID)
	_, t, err := Translation_GetByID(translation_ID)
	// //spew.Dump(translationItem)
	if err != nil {
		logs.Information("Translation", err.Error())
	}
	// // Check if tis null
	if (dm.Translation{}) == t {
		// Create One
		t.Id = translation_ID
		t.Class = class
		t.Message = message
		t.Translation = message
		Translation_StoreSystem(t)
	}
	// //fmt.Printf("message: %v\n", message)
	// return t.Translation
	return t.Translation
}
