package jobs

import (
	"fmt"

	core "github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	"github.com/mt1976/ebEstimates/logs"
)

func ClearFiles_Job_impl(j dm.JobDefinition) dm.JobDefinition {
	j.Name = "System Clear Files"
	j.Period = ""
	j.Description = "Clearsdown Business Data From The System - Testing *ONLY*"
	j.Type = core.Adhoc
	j.ID = j.Name
	return j
}

func ClearFiles_Run_impl() (string, error) {

	message := ""
	mode := core.ApplicationEnvironment()
	if mode == "production" {
		message = "This Job is not available in production mode"
		return message, nil
	}

	// Clear the data

	_, projects, _ := dao.Project_GetList()
	for _, project := range projects {
		logs.Warning("Deleting Project: " + project.ProjectID + " " + project.Name)
		dao.Project_HardDelete(project.ProjectID)
	}

	_, estimations, _ := dao.EstimationSession_GetList()
	for _, estimation := range estimations {
		logs.Warning("Deleting Estimation: " + estimation.EstimationSessionID + " " + estimation.Name)
		dao.EstimationSession_HardDelete(estimation.EstimationSessionID)
	}

	_, features, _ := dao.Feature_GetList()
	for _, feature := range features {
		logs.Warning("Deleting Feature: " + feature.FeatureID + " " + feature.Name)
		dao.Feature_HardDelete(feature.FeatureID)
	}

	MSG := "Deleted - Projects : %v  Estimations : %v Features : %v"
	MSG = dao.Translate("Message", MSG)
	message = fmt.Sprintf(MSG, len(projects), len(estimations), len(features))
	logs.Information("ClearFiles", message)
	return message, nil
}
