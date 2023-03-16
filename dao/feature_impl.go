package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/Feature.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Feature (Feature)
// Endpoint 	        : Feature (FeatureID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 27/11/2022 at 20:46:13
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"fmt"
	"math/big"

	core "github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/logs"

	dm "github.com/mt1976/ebEstimates/datamodel"
)

// Feature_GetList() returns a list of all Feature records
func Feature_ByEstimationSession_GetList() (int, []dm.Feature, error) {

	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Feature_SQLTable)
	count, FeatureList, _, _ := feature_Fetch(tsql)

	return count, FeatureList, nil
}

// Feature_GetList() returns a list of all Feature records
func Feature_Active_ByEstimationSession_GetList(id string) (int, []dm.Feature, error) {

	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.Feature_SQLTable)
	tsql = tsql + " WHERE " + dm.Feature_EstimationSessionID_sql + "='" + id + "'"
	tsql = tsql + " AND datalength(" + dm.Feature_SYSDeleted_sql + ") = 0"
	count, FeatureList, _, _ := feature_Fetch(tsql)

	return count, FeatureList, nil
}

func Feature_CalcDefaults(item dm.Feature, updateBaseline bool, who string) dm.Feature {
	// START
	// Dynamically generated 29/11/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	logs.Information("Recaulcating Feature Defaults", item.Name)
	// logs.Warning("Feature_CalcDefaults")
	// fmt.Printf("item.ProfileSelected: %v\n", item.ProfileSelected)
	// fmt.Printf("item.ProfileSelectedOld: %v\n", item.ProfileSelectedOld)
	// logs.Warning("Feature_CalcDefaults")
	// logs.Warning("Feature_CalcDefaults")
	//fmt.Printf("item: %v\n", item)
	// Get Project Profile for this Estimation Session
	// Get the Estimation Session
	_, es, _ := EstimationSession_GetByID(item.EstimationSessionID)
	// Then Get the Project
	_, p, _ := Project_GetByID(es.ProjectID)
	if item.ProfileDefault == "" {
		item.ProfileDefault = p.ProfileID
	}
	if item.ProfileSelected == "" {
		item.ProfileSelected = item.ProfileDefault
	}
	if item.ProjectManagerResource == "" {
		item.ProjectManagerResource = es.ProjectManager
	}
	if item.ProductManagerResource == "" {
		item.ProductManagerResource = es.ProductManager
	}
	// Then Get the Profile
	_, pp, _ := Profile_GetByCode(item.ProfileSelected)
	// fmt.Printf("updateBaseline: %v\n", updateBaseline)
	// fmt.Printf("item.ProfileSelected: %v\n", item.ProfileSelected)
	// fmt.Printf("profile: %v\n", pp)
	//spew.Dump(pp)

	if who == "" {
		who = item.SYSCreatedBy
		if item.SYSUpdatedBy != "" {
			who = item.SYSUpdatedBy
		}
	}

	if who == "" {
		who = Audit_Host()
	}

	if item.DevelopmentEstimate == "" {
		item.Activity = core.AddActivity_ForUser(item.Activity, "Development Estimate Is Blank", who)
		item.DevelopmentEstimate = "0.0"
		logs.Warning("Development Estimate Is Blank")
	}

	baseEstimate := core.StringToFloat(item.DevelopmentEstimate)
	roundingFactor, _ := RoundingHours()

	reqPerc := core.StringToFloat(pp.REQPerc)
	anaPerc := core.StringToFloat(pp.ANAPerc)
	docPerc := core.StringToFloat(pp.DOCPerc)
	pmPerc := core.StringToFloat(pp.PMPerc)
	uatPerc := core.StringToFloat(pp.UATPerc)
	gtmPerc := core.StringToFloat(pp.GTMPerc)
	trainPerc := core.StringToFloat(pp.TrainingPerc)
	//supportUpliftPerc, _ := strconv.ParseFloat(pp.SupportUplift, 64)

	// Get the Confidence Factor
	_, cf, _ := Confidence_GetByCode(item.ConfidenceCODE)
	devConfPerc := core.StringToFloat(cf.Perc)

	var coreEstimate float64
	if devConfPerc == 0 {
		coreEstimate = baseEstimate
	} else {
		coreEstimate = baseEstimate + (baseEstimate * (devConfPerc / 100))
	}

	roundEngEstimate, _ := core.RoundUpTo(coreEstimate, roundingFactor)
	item.DevelopmentFactoredDefault = core.FloatToString(roundEngEstimate)

	reqEstimate, _ := core.RoundUpTo(coreEstimate*(reqPerc/100), roundingFactor)
	item.RequirementsDefault = core.FloatToString(reqEstimate)
	item.RequirementsPerc = core.FloatToString(reqPerc)

	anaEstimate, _ := core.RoundUpTo(coreEstimate*(anaPerc/100), roundingFactor)
	item.AnalystTestDefault = core.FloatToString(anaEstimate)
	item.AnalystTestPerc = core.FloatToString(anaPerc)

	docEstimate, _ := core.RoundUpTo(coreEstimate*(docPerc/100), roundingFactor)
	item.DocumentationDefault = core.FloatToString(docEstimate)
	item.DocumentationPerc = core.FloatToString(docPerc)

	uatEstimate, _ := core.RoundUpTo(coreEstimate*(uatPerc/100), roundingFactor)
	item.UatSupportDefault = core.FloatToString(uatEstimate)
	item.UatSupportPerc = core.FloatToString(uatPerc)

	gtmEstimate, _ := core.RoundUpTo(coreEstimate*(gtmPerc/100), roundingFactor)
	item.MarketingDefault = core.FloatToString(gtmEstimate)
	item.MarketingPerc = core.FloatToString(gtmPerc)

	trainEstimate, _ := core.RoundUpTo(coreEstimate*(trainPerc/100), roundingFactor)
	item.TrainingDefault = core.FloatToString(trainEstimate)
	item.TrainingPerc = core.FloatToString(trainPerc)

	item.EstimateEffortDefault = item.EstimateEffort

	pmBase := reqEstimate + roundEngEstimate + anaEstimate + docEstimate + uatEstimate + trainEstimate

	pmEstimate, _ := core.RoundUpTo(pmBase*(pmPerc/100), roundingFactor)
	item.ManagementDefault = core.FloatToString(pmEstimate)
	item.ManagementPerc = core.FloatToString(pmPerc)

	item.Total = feature_CalculateTotal(item)

	if updateBaseline {
		item.Requirements = item.RequirementsDefault
		item.DevelopmentFactored = item.DevelopmentFactoredDefault
		item.AnalystTest = item.AnalystTestDefault
		if item.AnalystEstimate != "" {
			item.AnalystTest = item.AnalystEstimate
		}
		item.Documentation = item.DocumentationDefault
		item.Management = item.ManagementDefault
		item.UatSupport = item.UatSupportDefault
		item.Marketing = item.MarketingDefault
		item.Training = item.TrainingDefault
		item.EstimateEffort = item.EstimateEffortDefault
		//item.Total = item.TotalDefault
		item.TotalDefault = feature_CalculateTotalDefault(item)
	}

	msgTXT := Translate("AuditMessage", "Feature Defaults Recalculated")
	item.Activity = core.AddActivity_ForUser(item.Activity, msgTXT, who)

	msgTXT = Translate("AuditMessage", "Estimate = Base Estimate + Confidence Factor")
	item.Activity = core.AddActivity_ForUser(item.Activity, msgTXT, who)

	msgTXT = Translate("AuditMessage", "Estimate eq %s plus %sperc eq %s")
	msgTXT = fmt.Sprintf(msgTXT, core.FloatToString(baseEstimate), core.FloatToString(devConfPerc), core.FloatToString(coreEstimate))
	item.Activity = core.AddActivity_ForUser(item.Activity, msgTXT, who)

	msgTXT = Translate("AuditMessage", "| %s | %s | %s | %s | %s | %s | %s | %s | %s |")
	msgBASE := msgTXT
	brkTXT := fmt.Sprintf(msgBASE, "---", "---", "---", "---", "---", "---", "---", "---", "---")

	//msgTXT = Translate("AuditMessage", baseTXT)
	hdrTXT := fmt.Sprintf(msgBASE, "TYPE", "DEV", "REQ", "ANA", "DOC", "UAT", "TRG", "MKT", "MGT")
	item.Activity = core.AddActivity_ForUser(item.Activity, hdrTXT, who)
	item.Activity = core.AddActivity_ForUser(item.Activity, brkTXT, who)

	msgTXT = fmt.Sprintf(msgBASE, "Profile",
		core.FloatToString(devConfPerc+100),
		core.FloatToString(reqPerc),
		core.FloatToString(anaPerc),
		core.FloatToString(docPerc),
		core.FloatToString(uatPerc),
		core.FloatToString(trainPerc),
		core.FloatToString(gtmPerc),
		core.FloatToString(pmPerc))
	item.Activity = core.AddActivity_ForUser(item.Activity, msgTXT, who)

	msgTXT = fmt.Sprintf(msgBASE, "Result",
		core.FloatToString(coreEstimate),
		core.FloatToString(reqEstimate),
		core.FloatToString(anaEstimate),
		core.FloatToString(docEstimate),
		core.FloatToString(uatEstimate),
		core.FloatToString(trainEstimate),
		core.FloatToString(gtmEstimate),
		core.FloatToString(pmEstimate))
	item.Activity = core.AddActivity_ForUser(item.Activity, msgTXT, who)

	msgTXT = Translate("AuditMessage", "Total: %s")
	msgTXT = fmt.Sprintf(msgTXT, item.Total)
	item.Activity = core.AddActivity_ForUser(item.Activity, msgTXT, who)

	//item.TotalDefault = feature_CalculateTotalDefault(item)

	return item
}

// feature_CalculateTotal is used to calculate the total of the feature's estimate
func feature_CalculateTotal(item dm.Feature) string {
	total := 0.0
	total += core.StringToFloat(item.DevelopmentFactored)
	total += core.StringToFloat(item.Requirements)
	total += core.StringToFloat(item.AnalystTest)
	total += core.StringToFloat(item.Documentation)
	total += core.StringToFloat(item.Management)
	total += core.StringToFloat(item.UatSupport)
	total += core.StringToFloat(item.Marketing)
	total += core.StringToFloat(item.Contingency)
	total += core.StringToFloat(item.Training)
	total += core.StringToFloat(item.EstimateEffort)
	return core.FloatToString(total)
}

// feature_CalculateTotal is used to calculate the total of the feature's estimate
func feature_CalculateTotalDefault(item dm.Feature) string {
	//fmt.Printf("item: %v\n", item)
	total := 0.0
	total += core.StringToFloat(item.DevelopmentFactoredDefault)
	total += core.StringToFloat(item.RequirementsDefault)
	total += core.StringToFloat(item.AnalystTestDefault)
	total += core.StringToFloat(item.DocumentationDefault)
	total += core.StringToFloat(item.ManagementDefault)
	total += core.StringToFloat(item.UatSupportDefault)
	total += core.StringToFloat(item.MarketingDefault)
	total += core.StringToFloat(item.ContingencyDefault)
	total += core.StringToFloat(item.TrainingDefault)
	total += core.StringToFloat(item.EstimateEffortDefault)
	//fmt.Printf("total: %v\n", total)
	return core.FloatToString(total)
}

// feature_ValidateEstimate is used to validate the estimate of the feature
func feature_ValidateEstimate(baseEstimate string, userEstimate string, mismatch string, notes string, info string) (string, string) {
	//if mismatch {
	//	return mismatch, notes
	//}

	//fmt.Printf("baseEstimate: %v\n", baseEstimate)
	//fmt.Printf("userEstimate: %v\n", userEstimate)

	if baseEstimate == "" {
		baseEstimate = "0"
	}

	result := big.NewFloat(core.StringToFloat(baseEstimate)).Cmp(big.NewFloat(core.StringToFloat(userEstimate)))

	messageTXT := Translate("AuditMessage", "ESTIMATE MISMATCH [%s] [%s] != [%s]")
	messageTXT = fmt.Sprintf(messageTXT, info, baseEstimate, userEstimate)

	if result != 0 {
		notes = core.AddActivity(notes, messageTXT)
		return core.TRUE, notes
	}
	return core.FALSE, notes
}

// Feature_Clone clones a feature
func Feature_Clone(feature dm.Feature, esID string, userID string) error {
	// START

	newID := Feature_NewID(feature)

	feature.SYSId = ""
	feature.EstimationSessionID = esID
	feature.FeatureID = newID
	feature.SYSCreated = Audit_TimeStamp()
	feature.SYSCreatedBy = userID
	feature.SYSCreatedHost = Audit_Host()
	feature.SYSUpdated = ""
	feature.SYSUpdatedBy = ""
	feature.SYSUpdatedHost = ""
	feature.SYSDeleted = ""
	feature.SYSDeletedBy = ""
	feature.SYSDeletedHost = ""

	feature.Activity = core.AddActivity_ForUser(feature.Activity, fmt.Sprintf(Translate("Audit Message", "Cloned from %s (%s)"), feature.Name, esID), userID)

	_, err := Feature_StoreSystem(feature)
	if err != nil {
		logs.Panic("Cannot Store Feature {"+feature.FeatureID+"}", err)
		return err
	}

	trigger := "Feature Cloned " + feature.Name + " (" + feature.FeatureID + ")"
	go Estimationsession_Calculate(esID, trigger)

	// END

	return nil
}
