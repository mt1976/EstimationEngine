package adaptor

import (
	"strings"

	"github.com/google/uuid"
	core "github.com/mt1976/ebEstimates/core"
)

func getNewFileID() string {
	nID := uuid.New().String()
	nID = core.RemoveSpecialChars(nID)
	nID = strings.ReplaceAll(nID, "-", "")
	return nID
}
