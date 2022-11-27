package dao

import (
	"net/http"

	core "github.com/mt1976/ebEstimates/core"
)

// getappMenuData
func Session_GetUserName(r *http.Request) string {
	return core.SessionManager.GetString(r.Context(), core.SessionUserName)
}
