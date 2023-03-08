package jobs

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
	application "github.com/mt1976/ebEstimates/application"
	core "github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

type RSC struct {
	// Customer,RSC,ADO,EXT REF,Status,Description,RSC Expiry,"Value / Cost","Logged/Raised","Last Update",
	Customer string `csv:"Customer"`
	RSC      string `csv:"RSC"`
	ADO      string `csv:"ADO"`
	EXTREF   string `csv:"EXT REF"`
	Status   string `csv:"Status"`
	Desc     string `csv:"Description"`
	Expiry   string `csv:"RSC Expiry"`
	Value    string `csv:"Value / Cost"`
	Logged   string `csv:"Logged/Raised"`
	Updated  string `csv:"Last Update"`
	Project  string
}

type RSCs []*RSC

func ssloader_Job_impl(j dm.JobDefinition) dm.JobDefinition {
	j.Name = "SpreadsheetLoader"
	j.Period = ""
	j.Description = "Load Data from Old Tracker Spreadsheet"
	j.Type = core.Adhoc
	j.ID = j.Name
	return j
}

func ssloader_Run_impl() (string, error) {

	message := ""
	jobName := ssloader_Job().Name

	//Get Path from Data to Spreadsheet
	ssPath, err := dao.Data_Get(jobName, JOB_IN_PATH, dm.Data_Category_Path)
	if err != nil {
		return "No Location Specified for " + jobName, err
	}
	if ssPath == "" {
		return "No Spreadsheet at path" + core.DQuote(ssPath), nil
	}

	//Get Mode from Data
	ssMode, err := dao.Data_Get(jobName, JOB_MODE, dm.Data_Category_State)
	if err != nil {
		return "No Mode Specified for " + jobName, err
	}
	if ssMode != "Trial" && ssMode != "Live" {
		return "Invalid Mode " + core.DQuote(ssMode) + " for " + jobName, err
	}

	trialMode := TrialMode(ssMode)

	if ssMode == "Done" {
		return jobName + " is Done", nil
	}
	if ssMode != "Trial" && ssMode != "Live" {
		return "Invalid Mode " + core.DQuote(ssMode) + " for " + jobName, err
	}
	//Load Spreadsheet
	ss, err := os.ReadFile(ssPath)
	if err != nil {
		return "Unable to Find Spreadsheet for " + jobName, err
	}
	content := string(ss)
	if len(content) == 0 {
		return "No Content in Spreadsheet for " + jobName, err
	}
	//logs.Information(jobName, strconv.Itoa(len(content)))
	in, err := os.Open(ssPath)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	today := time.Now().Format(core.DATEFORMAT)

	workItems := []*RSC{}

	if err := gocsv.UnmarshalFile(in, &workItems); err != nil {
		panic(err)
	}
	noWorkItems := 0
	totalWorkItems := len(workItems)
	for _, workItem := range workItems {
		if workItem.Desc == "" {
			continue
		}
		noWorkItems++
		//fmt.Println("Hello, ", workItem.Desc)
		// Process Work Item
		var newRSC RSC
		newRSC.Customer = dao.Translate("RSCCustomer", workItem.Customer)
		newRSC.RSC = workItem.RSC
		newRSC.ADO = workItem.ADO
		newRSC.EXTREF = workItem.EXTREF
		newRSC.Status = dao.Translate("RSCStatus", workItem.Status)
		newRSC.Desc = core.ReplaceSpecialChars(strings.TrimSpace(workItem.Desc))
		newRSC.Expiry = convertDate(workItem.Expiry)
		newRSC.Value = cleanValue(workItem.Value)
		newRSC.Logged = convertDate(workItem.Logged)
		newRSC.Updated = convertDate(workItem.Updated)
		newRSC.Project = dao.Translate("RSCProject", workItem.Customer)

		dumpRSC(newRSC)

		// Lookup Origin using Customer
		oCount, origin, err := dao.Origin_GetByCode(newRSC.Customer)
		if err != nil {
			return "Unable to Find Origin for " + core.DQuote(newRSC.Customer), err
		}
		if oCount == 0 {
			return "Unable to Find Origin for " + core.DQuote(newRSC.Customer), err
		}

		proj, projErr := getProject(newRSC.Customer, newRSC.Project, origin, trialMode)
		if projErr != nil {
			return "Unable to Find/Generate Project for " + core.DQuote(newRSC.Customer), projErr
		}
		proj.Notes = core.AddActivity_ForProcess(proj.Notes, core.DQuote(newRSC.Desc)+" loaded from Spreadsheet "+today, jobName)
		proj.ProjectAnalyst = "--"
		proj.ProjectEngineer = "--"
		proj.ProjectManager = origin.ProjectManager

		estim := newEstimationSession(proj, newRSC, today, noWorkItems)
		estim.ProjectManager = origin.ProjectManager
		feature := newFeature(estim, newRSC, proj, today)
		feature.ProjectManagerResource = origin.ProjectManager

		//Save Work Item
		//if !trialMode {
		//Save Work Item
		//logs.Storing("Origin", dumpOrigin(origin))
		//dao.Origin_StoreSystem(origin)
		logs.Storing("Project", dumpProject(proj))
		logs.Storing("EstimationSession", dumpEstimationSession(estim))
		logs.Storing("Feature", dumpFeature(feature))

		if !trialMode {
			dao.Project_StoreSystem(proj)
			dao.EstimationSession_StoreSystem(estim)
			dao.Feature_StoreSystem(feature)

		}
		//logs.Break()
		MSG := "Item %d of %d Processed"
		MSG = fmt.Sprintf(MSG, noWorkItems, totalWorkItems)
		application.Schedule_Update(ssloader_Job(), MSG)
	}

	//logs.Information("Spreadsheet", content)
	//Load Data into Database
	ret := RecordJobStatusInfo(message, noWorkItems, JOB_ACTION_IMPORT, jobName, ssPath, "", ssMode)
	if ret != "" {
		message = "Unable to Update Data for " + jobName + " " + ret
	}

	return message, nil
}

func newFeature(estim dm.EstimationSession, newRSC RSC, proj dm.Project, today string) dm.Feature {
	var feature dm.Feature
	feature.FeatureID = dao.Feature_NewID(feature)
	feature.EstimationSessionID = estim.EstimationSessionID
	feature.ConfidenceCODE = "HIGH"
	feature.Name = newRSC.Desc
	feature.DevelopmentEstimate = newRSC.Value
	feature.AdoID = newRSC.ADO
	feature.TrackerID = newRSC.RSC
	feature.ExtRef = newRSC.EXTREF
	feature.ProfileDefault = proj.ProfileID
	feature.ProfileSelected = proj.ProfileID
	feature.Activity = core.AddActivity_ForProcess(feature.Activity, core.DQuote(newRSC.Desc)+" loaded from Spreadsheet "+today, ssloader_Job().Name)
	feature.DeveloperResource = "--"
	feature.ApproverResource = "--"
	feature.ProductManagerResource = "--"
	feature.ProjectManagerResource = "--"
	feature.AnalystEstimate = "100"
	return feature
}

func newEstimationSession(proj dm.Project, newRSC RSC, today string, noWorkItems int) dm.EstimationSession {
	var estim dm.EstimationSession
	estim.EstimationSessionID = dao.EstimationSession_NewID(estim)
	estim.ProjectID = proj.ProjectID
	estim.EstimationStateID = newRSC.Status
	estim.ProjectProfileID = "ADHC"
	estim.Activity = core.AddActivity_ForProcess(estim.Activity, core.DQuote(newRSC.Desc)+" loaded from Spreadsheet "+today, ssloader_Job().Name)
	estim.Name = newRSC.Desc
	estim.Total = newRSC.Value
	estim.Name = newRSC.Desc
	estim.AdoID = newRSC.ADO
	estim.TrackerID = newRSC.EXTREF
	estim.EstRef = fmt.Sprintf("%d", noWorkItems)
	estim.ExpiryDate = newRSC.Expiry
	estim.ProductManager = "--"
	estim.ProjectManager = "--"
	estim.Approver = "--"
	estim.ProjectProfile = proj.ProfileID
	return estim
}

func dumpEstimationSession(estim dm.EstimationSession) string {
	// EstimationSession[Name|Project|Origin|Status|Notes]
	MSG := "EstimationSession[%s|%s|%s|%s|%s|%s]"
	return fmt.Sprintf(MSG, estim.Name, estim.EstimationSessionID, estim.ProjectID, estim.EstimationStateID, estim.AdoID, estim.FreshdeskID)
}

func dumpFeature(feature dm.Feature) string {
	// Feature[Name|Project|Origin|Status]
	MSG := "Feature[%s|%s|%s|%s]"
	return fmt.Sprintf(MSG, feature.Name, feature.ConfidenceCODE, feature.EstimationSessionID, feature.DevelopmentEstimate)
}

func dumpOrigin(origin dm.Origin) string {
	// Origin[Code|Name|Customer|Status|Notes]
	MSG := "Origin[%s|%s|%s|%s]"
	return fmt.Sprintf(MSG, origin.OriginID, origin.Code, origin.FullName, origin.SYSUpdatedBy)
}

func dumpProject(proj dm.Project) string {
	// Project[Code|Name|Customer|Status|Notes]
	MSG := "Project[%s|%s|%s|%s|%s]"
	return fmt.Sprintf(MSG, proj.ProjectID, proj.Name, proj.OriginID, proj.ProfileID, proj.ProjectStateID)
}

func getProject(cust string, project string, origin dm.Origin, trailMode bool) (dm.Project, error) {
	// Lookup Project by Customer and Project
	// If not found create Project
	// Re-Lookup Project by Customer and Project
	// End
	logs.Information("Project", "Looking for "+cust+"|"+project)
	count, proj, err := dao.Project_GetList_ByCustomerAndName(cust, project)
	if err != nil {
		logs.Error("Project", err)
		return proj, err
	}
	if count == 0 {
		logs.Information("Project", "Unable to Find Project for "+cust+"|"+project)

		proj.ProjectID = dao.Project_NewID(proj)
		proj.OriginID = origin.Code
		proj.Name = project
		proj.ProjectStateID = "TAKE"
		proj.ProfileID = "ADHC"
		proj.Notes = core.AddActivity_ForProcess(proj.Notes, "Created from RSC Spreadsheet", ssloader_Job().Name)
		if !trailMode {
			dao.Project_StoreSystem(proj)
			return proj, nil
		}
		if trailMode {
			logs.Information("Project", "Trail Mode - Not Updating Project for "+cust+"|"+project)
			//logs.Created("Project :" + dumpProject(proj))
		}
		//logs.Generate("Created Project for " + cust + "|" + project + "|" + proj.ProjectID + dumpProject(proj))
		if !trailMode {
			_, proj, err := dao.Project_GetList_ByCustomerAndName(cust, project)
			if err != nil {
				logs.Error("Project", err)
				return proj, err
			}
		}
		return proj, err

	}

	return proj, nil
}

func dumpRSC(rsc RSC) {
	MSG := "RSC[%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s]"
	logs.Generate("RSC:" + fmt.Sprintf(MSG, rsc.Customer, rsc.RSC, rsc.ADO, rsc.EXTREF, rsc.Status, rsc.Desc, rsc.Expiry, rsc.Value, rsc.Logged, rsc.Updated, rsc.Project))
}

func convertDate(date string) string {
	//Convert Date
	if date == "" {
		return ""
	}
	dt, err := time.Parse(core.DFSS, date)
	if err != nil {
		return ""
	}
	return dt.Format(core.DATEFORMAT)
}

func cleanValue(value string) string {
	//Convert Value
	if value == "" {
		return ""
	}
	value = strings.ReplaceAll(value, "£", "")
	value = strings.ReplaceAll(value, ",", "")
	return value
}
