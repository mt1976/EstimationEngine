package jobs

import (
	core "github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
)

func Index_Job_impl(j dm.JobDefinition) dm.JobDefinition {
	j.Name = "Indexer"
	j.Period = ""
	j.Description = "Rebuild Index"
	j.Type = core.Adhoc
	j.ID = "Indexer"
	return j
}

func Index_Run_impl() (string, error) {

	message := ""

	/// CONTENT STARTS

	// Rebuild the index
	err := dao.Indexer_Rebuild()
	if err != nil {
		return message, err
	}
	/// CONTENT ENDS

	message = "Index Rebuild Completed"
	dao.Translate("OriginWarning", message)
	//message = fmt.Sprintf(message, noOrigins, noWarnings, noDue, noOverdue)
	return message, nil
}
