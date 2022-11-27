package dao

import (
	"strings"

	core "github.com/mt1976/ebEstimates/core"
)

// Returns a valid translation store ID
func Translation_BuildID(class string, message string) string {
	var translationStoreID string
	message = strings.ReplaceAll(message, " ", "-")
	message = core.RemoveSpecialChars(message)
	translationStoreID = class + core.IDSep + message
	return translationStoreID
}
