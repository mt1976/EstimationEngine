package dao

import (
	"strings"

	core "github.com/mt1976/ebEstimates/core"
	das "github.com/mt1976/ebEstimates/das"
	dm "github.com/mt1976/ebEstimates/datamodel"
	"github.com/mt1976/ebEstimates/logs"
)

// Returns a valid translation store ID
func Translation_BuildID(class string, message string) string {
	var translationStoreID string
	message = strings.ReplaceAll(message, " ", "-")
	message = core.ReplaceSpecialChars(message)
	translationStoreID = class + core.ID_SEP + message

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
	if t.Id == "" {
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

func Translation_GetMigrateable() (int, []dm.Translation, error) {

	tsql := Translation_SQLbase
	tsql = tsql + " " + das.WHERE + dm.Translation_Migrate_sql + das.EQ + das.ID(das.TRUE)
	_, items, _, _ := translation_Fetch(tsql)
	return len(items), items, nil
}
