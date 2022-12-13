package jobs

import (
	core "github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
)

func Origin_Job_impl(j dm.JobDefinition) dm.JobDefinition {
	j.Name = "Origin"
	j.Period = "5 0 * * *"
	j.Description = "Revalidate Origins"
	j.Type = core.HouseKeeping
	j.ID = "Origin"
	return j
}

func Origin_Run_impl() (string, error) {

	message := ""
	/// CONTENT STARTS

	/// CONTENT ENDS
	return message, nil

}
