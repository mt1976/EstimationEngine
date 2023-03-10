package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/feature.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Feature (feature)
// Endpoint 	        : Feature (FeatureID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 10/03/2023 at 19:54:33
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"fmt"
	"net/http"
	core "github.com/mt1976/ebEstimates/core"
	"github.com/google/uuid"
	das  "github.com/mt1976/ebEstimates/das"
	dm   "github.com/mt1976/ebEstimates/datamodel"
	logs   "github.com/mt1976/ebEstimates/logs"
	"github.com/pkg/errors"
)

var Feature_SQLbase string
var Feature_QualifiedName string
func init(){
	Feature_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.Feature_SQLTable)
	Feature_SQLbase =  das.SELECTALL + das.FROM + Feature_QualifiedName
}

// Feature_GetList() returns a list of all Feature records
func Feature_GetList() (int, []dm.Feature, error) {
	count, featureList, err := Feature_GetListFiltered("")
	return count, featureList, err
}

// Feature_GetListFiltered() returns a filtered list of all Feature records
func Feature_GetListFiltered(filter string) (int, []dm.Feature, error) {
	
	tsql := Feature_SQLbase
	if filter != "" {
		tsql = tsql + " " + das.WHERE + filter
	}
	count, featureList, _, _ := feature_Fetch(tsql)
	
	return count, featureList, nil
}


// Feature_GetLookup() returns a lookup list of all Feature items in lookup format
func Feature_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, featureList, _ := Feature_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: featureList[i].FeatureID, Name: featureList[i].Name})
	}
	return returnList
}

// Feature_GetFilteredLookup() returns a lookup list of all Feature items in lookup format
func Feature_GetFilteredLookup(requestObject string,requestField string) []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	objectName := Translate("ObjectName", requestObject)
	reqField := requestField+"_Feature_Filter"
	filter,_ := Data_GetString(objectName, reqField, dm.Data_Category_FilterRule)
	if filter == "" {
		logs.Warning("Feature_GetFilteredLookup() - No filter found : " + reqField + " for Object: " + objectName)
	} 
	count, featureList, _ := Feature_GetListFiltered(filter)
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: featureList[i].FeatureID, Name: featureList[i].Name})
	}
	return returnList
}



// Feature_GetByID() returns a single Feature record
func Feature_GetByID(id string) (int, dm.Feature, error) {


	tsql := Feature_SQLbase
	tsql = tsql + " " + das.WHERE + dm.Feature_SQLSearchID + das.EQ + das.ID(id)
	_, _, featureItem, _ := feature_Fetch(tsql)


	featureItem = Feature_PostGet(featureItem,id)

	return 1, featureItem, nil
}

func Feature_PostGet(featureItem dm.Feature,id string) dm.Feature {
	// START
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	featureItem.FeatureID,featureItem.FeatureID_props = Feature_FeatureID_validate_impl (GET,id,featureItem.FeatureID,featureItem,featureItem.FeatureID_props)
	featureItem.EstimationSessionID,featureItem.EstimationSessionID_props = Feature_EstimationSessionID_validate_impl (GET,id,featureItem.EstimationSessionID,featureItem,featureItem.EstimationSessionID_props)
	featureItem.Name,featureItem.Name_props = Feature_Name_validate_impl (GET,id,featureItem.Name,featureItem,featureItem.Name_props)
	featureItem.DevelopmentEstimate,featureItem.DevelopmentEstimate_props = Feature_DevelopmentEstimate_validate_impl (GET,id,featureItem.DevelopmentEstimate,featureItem,featureItem.DevelopmentEstimate_props)
	featureItem.DevelopmentFactored,featureItem.DevelopmentFactored_props = Feature_DevelopmentFactored_validate_impl (GET,id,featureItem.DevelopmentFactored,featureItem,featureItem.DevelopmentFactored_props)
	featureItem.Requirements,featureItem.Requirements_props = Feature_Requirements_validate_impl (GET,id,featureItem.Requirements,featureItem,featureItem.Requirements_props)
	featureItem.AnalystTest,featureItem.AnalystTest_props = Feature_AnalystTest_validate_impl (GET,id,featureItem.AnalystTest,featureItem,featureItem.AnalystTest_props)
	featureItem.Documentation,featureItem.Documentation_props = Feature_Documentation_validate_impl (GET,id,featureItem.Documentation,featureItem,featureItem.Documentation_props)
	featureItem.Management,featureItem.Management_props = Feature_Management_validate_impl (GET,id,featureItem.Management,featureItem,featureItem.Management_props)
	featureItem.Marketing,featureItem.Marketing_props = Feature_Marketing_validate_impl (GET,id,featureItem.Marketing,featureItem,featureItem.Marketing_props)
	featureItem.Contingency,featureItem.Contingency_props = Feature_Contingency_validate_impl (GET,id,featureItem.Contingency,featureItem,featureItem.Contingency_props)
	featureItem.TrackerID,featureItem.TrackerID_props = Feature_TrackerID_validate_impl (GET,id,featureItem.TrackerID,featureItem,featureItem.TrackerID_props)
	featureItem.AdoID,featureItem.AdoID_props = Feature_AdoID_validate_impl (GET,id,featureItem.AdoID,featureItem,featureItem.AdoID_props)
	featureItem.FreshdeskID,featureItem.FreshdeskID_props = Feature_FreshdeskID_validate_impl (GET,id,featureItem.FreshdeskID,featureItem,featureItem.FreshdeskID_props)
	featureItem.ExtRef,featureItem.ExtRef_props = Feature_ExtRef_validate_impl (GET,id,featureItem.ExtRef,featureItem,featureItem.ExtRef_props)
	featureItem.ExtRef2,featureItem.ExtRef2_props = Feature_ExtRef2_validate_impl (GET,id,featureItem.ExtRef2,featureItem,featureItem.ExtRef2_props)
	featureItem.DeveloperResource,featureItem.DeveloperResource_props = Feature_DeveloperResource_validate_impl (GET,id,featureItem.DeveloperResource,featureItem,featureItem.DeveloperResource_props)
	featureItem.ApproverResource,featureItem.ApproverResource_props = Feature_ApproverResource_validate_impl (GET,id,featureItem.ApproverResource,featureItem,featureItem.ApproverResource_props)
	featureItem.OffProfileJustification,featureItem.OffProfileJustification_props = Feature_OffProfileJustification_validate_impl (GET,id,featureItem.OffProfileJustification,featureItem,featureItem.OffProfileJustification_props)
	featureItem.Total,featureItem.Total_props = Feature_Total_validate_impl (GET,id,featureItem.Total,featureItem,featureItem.Total_props)
	featureItem.Training,featureItem.Training_props = Feature_Training_validate_impl (GET,id,featureItem.Training,featureItem,featureItem.Training_props)
	featureItem.EstimateEffort,featureItem.EstimateEffort_props = Feature_EstimateEffort_validate_impl (GET,id,featureItem.EstimateEffort,featureItem,featureItem.EstimateEffort_props)
	featureItem.AnalystEstimate,featureItem.AnalystEstimate_props = Feature_AnalystEstimate_validate_impl (GET,id,featureItem.AnalystEstimate,featureItem,featureItem.AnalystEstimate_props)
	featureItem.OriginName,featureItem.OriginName_props = Feature_OriginName_validate_impl (GET,id,featureItem.OriginName,featureItem,featureItem.OriginName_props)
	featureItem.ProjectName,featureItem.ProjectName_props = Feature_ProjectName_validate_impl (GET,id,featureItem.ProjectName,featureItem,featureItem.ProjectName_props)
	featureItem.OriginID,featureItem.OriginID_props = Feature_OriginID_validate_impl (GET,id,featureItem.OriginID,featureItem,featureItem.OriginID_props)
	featureItem.ProjectID,featureItem.ProjectID_props = Feature_ProjectID_validate_impl (GET,id,featureItem.ProjectID,featureItem,featureItem.ProjectID_props)
	featureItem.ProfileSelectedOld,featureItem.ProfileSelectedOld_props = Feature_ProfileSelectedOld_validate_impl (GET,id,featureItem.ProfileSelectedOld,featureItem,featureItem.ProfileSelectedOld_props)
	featureItem.CCY,featureItem.CCY_props = Feature_CCY_validate_impl (GET,id,featureItem.CCY,featureItem,featureItem.CCY_props)
	featureItem.RoundTo,featureItem.RoundTo_props = Feature_RoundTo_validate_impl (GET,id,featureItem.RoundTo,featureItem,featureItem.RoundTo_props)
	featureItem.OriginCode,featureItem.OriginCode_props = Feature_OriginCode_validate_impl (GET,id,featureItem.OriginCode,featureItem,featureItem.OriginCode_props)
	featureItem.EstimationSessionName,featureItem.EstimationSessionName_props = Feature_EstimationSessionName_validate_impl (GET,id,featureItem.EstimationSessionName,featureItem,featureItem.EstimationSessionName_props)
	// 
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return featureItem
}




// Feature_DeleteByID() deletes a single Feature record
func Feature_Delete(id string) {
		Feature_SoftDelete(id)
	}
// Feature_SoftDelete(id string) soft deletes a single Feature record
func Feature_SoftDelete(id string) {
	//Uses Soft Delete
		_,featureItem,_ := Feature_GetByID(id)
		featureItem.SYSDeletedBy = Audit_Update("", Audit_Host())
		featureItem.SYSDeleted = Audit_Update("", Audit_TimeStamp())
		featureItem.SYSDeletedHost = Audit_Update("", Audit_Host())
		_,err := Feature_StoreSystem(featureItem)
		if err != nil {
			logs.Error("Feature_SoftDelete()",err)
		}
}
	

// Feature_HardDelete(id string) soft deletes a single Feature record
func Feature_HardDelete(id string) {
			// Uses Hard Delete
		object_Table := Feature_QualifiedName
		tsql := das.DELETE + das.FROM + object_Table
		tsql = tsql + " " + das.WHERE + dm.Feature_SQLSearchID + das.EQ + das.ID(id)
		das.Execute(tsql)
		//if err != nil {
		//	logs.Error("Feature_SoftDelete()",err)
		//}
}


// Feature_Store() saves/stores a Feature record to the database
func Feature_Store(r dm.Feature,req *http.Request) (dm.Feature,error) {

	r, err := Feature_Validate(r)
	if err == nil {
		err = feature_Save(r, Audit_User(req))
	} else {
		logs.Information("Feature_Store()", err.Error())
	}

	return r, err
}

// Feature_StoreSystem() saves/stores a Feature record to the database
func Feature_StoreSystem(r dm.Feature) (dm.Feature,error) {
	
	r, err := Feature_Validate(r)
	if err == nil {
		err = feature_Save(r, Audit_Host())
	} else {
		logs.Information("Feature_Store()", err.Error())
	}

	return r, err
}

// Feature_StoreProcess() saves/stores a Feature record to the database
func Feature_StoreProcess(r dm.Feature, operator string) (dm.Feature,error) {
	
	r, err := Feature_Validate(r)
	if err == nil {
		err = feature_Save(r, operator)
	} else {
		logs.Information("Feature_StoreProcess()", err.Error())
	}

	return r, err
}

// Feature_Validate() validates for saves/stores a Feature record to the database
func Feature_Validate(r dm.Feature) (dm.Feature, error) {
	var err error
	// START
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	r.FeatureID,r.FeatureID_props = Feature_FeatureID_validate_impl (PUT,r.FeatureID,r.FeatureID,r,r.FeatureID_props)
	if r.FeatureID_props.MsgMessage != "" {
		err = errors.New(r.FeatureID_props.MsgMessage)
	}
	r.EstimationSessionID,r.EstimationSessionID_props = Feature_EstimationSessionID_validate_impl (PUT,r.FeatureID,r.EstimationSessionID,r,r.EstimationSessionID_props)
	if r.EstimationSessionID_props.MsgMessage != "" {
		err = errors.New(r.EstimationSessionID_props.MsgMessage)
	}
	r.Name,r.Name_props = Feature_Name_validate_impl (PUT,r.FeatureID,r.Name,r,r.Name_props)
	if r.Name_props.MsgMessage != "" {
		err = errors.New(r.Name_props.MsgMessage)
	}
	r.DevelopmentEstimate,r.DevelopmentEstimate_props = Feature_DevelopmentEstimate_validate_impl (PUT,r.FeatureID,r.DevelopmentEstimate,r,r.DevelopmentEstimate_props)
	if r.DevelopmentEstimate_props.MsgMessage != "" {
		err = errors.New(r.DevelopmentEstimate_props.MsgMessage)
	}
	r.DevelopmentFactored,r.DevelopmentFactored_props = Feature_DevelopmentFactored_validate_impl (PUT,r.FeatureID,r.DevelopmentFactored,r,r.DevelopmentFactored_props)
	if r.DevelopmentFactored_props.MsgMessage != "" {
		err = errors.New(r.DevelopmentFactored_props.MsgMessage)
	}
	r.Requirements,r.Requirements_props = Feature_Requirements_validate_impl (PUT,r.FeatureID,r.Requirements,r,r.Requirements_props)
	if r.Requirements_props.MsgMessage != "" {
		err = errors.New(r.Requirements_props.MsgMessage)
	}
	r.AnalystTest,r.AnalystTest_props = Feature_AnalystTest_validate_impl (PUT,r.FeatureID,r.AnalystTest,r,r.AnalystTest_props)
	if r.AnalystTest_props.MsgMessage != "" {
		err = errors.New(r.AnalystTest_props.MsgMessage)
	}
	r.Documentation,r.Documentation_props = Feature_Documentation_validate_impl (PUT,r.FeatureID,r.Documentation,r,r.Documentation_props)
	if r.Documentation_props.MsgMessage != "" {
		err = errors.New(r.Documentation_props.MsgMessage)
	}
	r.Management,r.Management_props = Feature_Management_validate_impl (PUT,r.FeatureID,r.Management,r,r.Management_props)
	if r.Management_props.MsgMessage != "" {
		err = errors.New(r.Management_props.MsgMessage)
	}
	r.Marketing,r.Marketing_props = Feature_Marketing_validate_impl (PUT,r.FeatureID,r.Marketing,r,r.Marketing_props)
	if r.Marketing_props.MsgMessage != "" {
		err = errors.New(r.Marketing_props.MsgMessage)
	}
	r.Contingency,r.Contingency_props = Feature_Contingency_validate_impl (PUT,r.FeatureID,r.Contingency,r,r.Contingency_props)
	if r.Contingency_props.MsgMessage != "" {
		err = errors.New(r.Contingency_props.MsgMessage)
	}
	r.TrackerID,r.TrackerID_props = Feature_TrackerID_validate_impl (PUT,r.FeatureID,r.TrackerID,r,r.TrackerID_props)
	if r.TrackerID_props.MsgMessage != "" {
		err = errors.New(r.TrackerID_props.MsgMessage)
	}
	r.AdoID,r.AdoID_props = Feature_AdoID_validate_impl (PUT,r.FeatureID,r.AdoID,r,r.AdoID_props)
	if r.AdoID_props.MsgMessage != "" {
		err = errors.New(r.AdoID_props.MsgMessage)
	}
	r.FreshdeskID,r.FreshdeskID_props = Feature_FreshdeskID_validate_impl (PUT,r.FeatureID,r.FreshdeskID,r,r.FreshdeskID_props)
	if r.FreshdeskID_props.MsgMessage != "" {
		err = errors.New(r.FreshdeskID_props.MsgMessage)
	}
	r.ExtRef,r.ExtRef_props = Feature_ExtRef_validate_impl (PUT,r.FeatureID,r.ExtRef,r,r.ExtRef_props)
	if r.ExtRef_props.MsgMessage != "" {
		err = errors.New(r.ExtRef_props.MsgMessage)
	}
	r.ExtRef2,r.ExtRef2_props = Feature_ExtRef2_validate_impl (PUT,r.FeatureID,r.ExtRef2,r,r.ExtRef2_props)
	if r.ExtRef2_props.MsgMessage != "" {
		err = errors.New(r.ExtRef2_props.MsgMessage)
	}
	r.DeveloperResource,r.DeveloperResource_props = Feature_DeveloperResource_validate_impl (PUT,r.FeatureID,r.DeveloperResource,r,r.DeveloperResource_props)
	if r.DeveloperResource_props.MsgMessage != "" {
		err = errors.New(r.DeveloperResource_props.MsgMessage)
	}
	r.ApproverResource,r.ApproverResource_props = Feature_ApproverResource_validate_impl (PUT,r.FeatureID,r.ApproverResource,r,r.ApproverResource_props)
	if r.ApproverResource_props.MsgMessage != "" {
		err = errors.New(r.ApproverResource_props.MsgMessage)
	}
	r.OffProfileJustification,r.OffProfileJustification_props = Feature_OffProfileJustification_validate_impl (PUT,r.FeatureID,r.OffProfileJustification,r,r.OffProfileJustification_props)
	if r.OffProfileJustification_props.MsgMessage != "" {
		err = errors.New(r.OffProfileJustification_props.MsgMessage)
	}
	r.Total,r.Total_props = Feature_Total_validate_impl (PUT,r.FeatureID,r.Total,r,r.Total_props)
	if r.Total_props.MsgMessage != "" {
		err = errors.New(r.Total_props.MsgMessage)
	}
	r.Training,r.Training_props = Feature_Training_validate_impl (PUT,r.FeatureID,r.Training,r,r.Training_props)
	if r.Training_props.MsgMessage != "" {
		err = errors.New(r.Training_props.MsgMessage)
	}
	r.EstimateEffort,r.EstimateEffort_props = Feature_EstimateEffort_validate_impl (PUT,r.FeatureID,r.EstimateEffort,r,r.EstimateEffort_props)
	if r.EstimateEffort_props.MsgMessage != "" {
		err = errors.New(r.EstimateEffort_props.MsgMessage)
	}
	r.AnalystEstimate,r.AnalystEstimate_props = Feature_AnalystEstimate_validate_impl (PUT,r.FeatureID,r.AnalystEstimate,r,r.AnalystEstimate_props)
	if r.AnalystEstimate_props.MsgMessage != "" {
		err = errors.New(r.AnalystEstimate_props.MsgMessage)
	}
	r.OriginName,r.OriginName_props = Feature_OriginName_validate_impl (PUT,r.FeatureID,r.OriginName,r,r.OriginName_props)
	if r.OriginName_props.MsgMessage != "" {
		err = errors.New(r.OriginName_props.MsgMessage)
	}
	r.ProjectName,r.ProjectName_props = Feature_ProjectName_validate_impl (PUT,r.FeatureID,r.ProjectName,r,r.ProjectName_props)
	if r.ProjectName_props.MsgMessage != "" {
		err = errors.New(r.ProjectName_props.MsgMessage)
	}
	r.OriginID,r.OriginID_props = Feature_OriginID_validate_impl (PUT,r.FeatureID,r.OriginID,r,r.OriginID_props)
	if r.OriginID_props.MsgMessage != "" {
		err = errors.New(r.OriginID_props.MsgMessage)
	}
	r.ProjectID,r.ProjectID_props = Feature_ProjectID_validate_impl (PUT,r.FeatureID,r.ProjectID,r,r.ProjectID_props)
	if r.ProjectID_props.MsgMessage != "" {
		err = errors.New(r.ProjectID_props.MsgMessage)
	}
	r.ProfileSelectedOld,r.ProfileSelectedOld_props = Feature_ProfileSelectedOld_validate_impl (PUT,r.FeatureID,r.ProfileSelectedOld,r,r.ProfileSelectedOld_props)
	if r.ProfileSelectedOld_props.MsgMessage != "" {
		err = errors.New(r.ProfileSelectedOld_props.MsgMessage)
	}
	r.CCY,r.CCY_props = Feature_CCY_validate_impl (PUT,r.FeatureID,r.CCY,r,r.CCY_props)
	if r.CCY_props.MsgMessage != "" {
		err = errors.New(r.CCY_props.MsgMessage)
	}
	r.RoundTo,r.RoundTo_props = Feature_RoundTo_validate_impl (PUT,r.FeatureID,r.RoundTo,r,r.RoundTo_props)
	if r.RoundTo_props.MsgMessage != "" {
		err = errors.New(r.RoundTo_props.MsgMessage)
	}
	r.OriginCode,r.OriginCode_props = Feature_OriginCode_validate_impl (PUT,r.FeatureID,r.OriginCode,r,r.OriginCode_props)
	if r.OriginCode_props.MsgMessage != "" {
		err = errors.New(r.OriginCode_props.MsgMessage)
	}
	r.EstimationSessionName,r.EstimationSessionName_props = Feature_EstimationSessionName_validate_impl (PUT,r.FeatureID,r.EstimationSessionName,r,r.EstimationSessionName_props)
	if r.EstimationSessionName_props.MsgMessage != "" {
		err = errors.New(r.EstimationSessionName_props.MsgMessage)
	}
	// 

	
	// Cross Validation
	var errVal error
	r, _, errVal = Feature_ObjectValidation_impl(PUT, r.FeatureID, r)
	if errVal != nil {
		err = errVal
	}
	

	return r,err
}
//

// feature_Save() saves/stores a Feature record to the database
func feature_Save(r dm.Feature,usr string) error {

    var err error

	if len(r.FeatureID) == 0 {
		r.FeatureID = Feature_NewID(r)
	}

// If there are fields below, create the methods in dao\feature_impl.go
    r.FeatureID,err = Feature_FeatureID_OnStore_impl (r.FeatureID,r,usr)
    r.EstimationSessionID,err = Feature_EstimationSessionID_OnStore_impl (r.EstimationSessionID,r,usr)
    r.Name,err = Feature_Name_OnStore_impl (r.Name,r,usr)
    r.DevelopmentEstimate,err = Feature_DevelopmentEstimate_OnStore_impl (r.DevelopmentEstimate,r,usr)
    r.DevelopmentFactored,err = Feature_DevelopmentFactored_OnStore_impl (r.DevelopmentFactored,r,usr)
    r.Requirements,err = Feature_Requirements_OnStore_impl (r.Requirements,r,usr)
    r.AnalystTest,err = Feature_AnalystTest_OnStore_impl (r.AnalystTest,r,usr)
    r.Documentation,err = Feature_Documentation_OnStore_impl (r.Documentation,r,usr)
    r.Management,err = Feature_Management_OnStore_impl (r.Management,r,usr)
    r.Marketing,err = Feature_Marketing_OnStore_impl (r.Marketing,r,usr)
    r.Contingency,err = Feature_Contingency_OnStore_impl (r.Contingency,r,usr)
    r.TrackerID,err = Feature_TrackerID_OnStore_impl (r.TrackerID,r,usr)
    r.AdoID,err = Feature_AdoID_OnStore_impl (r.AdoID,r,usr)
    r.FreshdeskID,err = Feature_FreshdeskID_OnStore_impl (r.FreshdeskID,r,usr)
    r.ExtRef,err = Feature_ExtRef_OnStore_impl (r.ExtRef,r,usr)
    r.ExtRef2,err = Feature_ExtRef2_OnStore_impl (r.ExtRef2,r,usr)
    r.DeveloperResource,err = Feature_DeveloperResource_OnStore_impl (r.DeveloperResource,r,usr)
    r.ApproverResource,err = Feature_ApproverResource_OnStore_impl (r.ApproverResource,r,usr)
    r.OffProfileJustification,err = Feature_OffProfileJustification_OnStore_impl (r.OffProfileJustification,r,usr)
    r.Total,err = Feature_Total_OnStore_impl (r.Total,r,usr)
    r.Training,err = Feature_Training_OnStore_impl (r.Training,r,usr)
    r.EstimateEffort,err = Feature_EstimateEffort_OnStore_impl (r.EstimateEffort,r,usr)
    r.AnalystEstimate,err = Feature_AnalystEstimate_OnStore_impl (r.AnalystEstimate,r,usr)
    r.OriginName,err = Feature_OriginName_OnStore_impl (r.OriginName,r,usr)
    r.ProjectName,err = Feature_ProjectName_OnStore_impl (r.ProjectName,r,usr)
    r.OriginID,err = Feature_OriginID_OnStore_impl (r.OriginID,r,usr)
    r.ProjectID,err = Feature_ProjectID_OnStore_impl (r.ProjectID,r,usr)
    r.ProfileSelectedOld,err = Feature_ProfileSelectedOld_OnStore_impl (r.ProfileSelectedOld,r,usr)
    r.CCY,err = Feature_CCY_OnStore_impl (r.CCY,r,usr)
    r.RoundTo,err = Feature_RoundTo_OnStore_impl (r.RoundTo,r,usr)
    r.OriginCode,err = Feature_OriginCode_OnStore_impl (r.OriginCode,r,usr)
    r.EstimationSessionName,err = Feature_EstimationSessionName_OnStore_impl (r.EstimationSessionName,r,usr)

	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	r.SYSDbVersion = core.DB_Version()
	
logs.Storing("Feature",fmt.Sprintf("%v", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.Feature_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.Feature_FeatureID_sql, r.FeatureID)
	ts = addData(ts, dm.Feature_EstimationSessionID_sql, r.EstimationSessionID)
	ts = addData(ts, dm.Feature_ConfidenceCODE_sql, r.ConfidenceCODE)
	ts = addData(ts, dm.Feature_Name_sql, r.Name)
	ts = addData(ts, dm.Feature_DevelopmentEstimate_sql, r.DevelopmentEstimate)
	ts = addData(ts, dm.Feature_DevelopmentFactored_sql, r.DevelopmentFactored)
	ts = addData(ts, dm.Feature_Requirements_sql, r.Requirements)
	ts = addData(ts, dm.Feature_AnalystTest_sql, r.AnalystTest)
	ts = addData(ts, dm.Feature_Documentation_sql, r.Documentation)
	ts = addData(ts, dm.Feature_Management_sql, r.Management)
	ts = addData(ts, dm.Feature_UatSupport_sql, r.UatSupport)
	ts = addData(ts, dm.Feature_Marketing_sql, r.Marketing)
	ts = addData(ts, dm.Feature_Contingency_sql, r.Contingency)
	ts = addData(ts, dm.Feature_TrackerID_sql, r.TrackerID)
	ts = addData(ts, dm.Feature_AdoID_sql, r.AdoID)
	ts = addData(ts, dm.Feature_FreshdeskID_sql, r.FreshdeskID)
	ts = addData(ts, dm.Feature_ExtRef_sql, r.ExtRef)
	ts = addData(ts, dm.Feature_ExtRef2_sql, r.ExtRef2)
	ts = addData(ts, dm.Feature_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.Feature_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.Feature_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.Feature_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.Feature_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.Feature_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.Feature_SYSDeleted_sql, r.SYSDeleted)
	ts = addData(ts, dm.Feature_SYSDeletedBy_sql, r.SYSDeletedBy)
	ts = addData(ts, dm.Feature_SYSDeletedHost_sql, r.SYSDeletedHost)
	ts = addData(ts, dm.Feature_DeveloperResource_sql, r.DeveloperResource)
	ts = addData(ts, dm.Feature_ApproverResource_sql, r.ApproverResource)
	ts = addData(ts, dm.Feature_Activity_sql, r.Activity)
	ts = addData(ts, dm.Feature_OffProfile_sql, r.OffProfile)
	ts = addData(ts, dm.Feature_OffProfileJustification_sql, r.OffProfileJustification)
	ts = addData(ts, dm.Feature_SYSActivity_sql, r.SYSActivity)
	ts = addData(ts, dm.Feature_RequirementsDefault_sql, r.RequirementsDefault)
	ts = addData(ts, dm.Feature_AnalystTestDefault_sql, r.AnalystTestDefault)
	ts = addData(ts, dm.Feature_DocumentationDefault_sql, r.DocumentationDefault)
	ts = addData(ts, dm.Feature_ManagementDefault_sql, r.ManagementDefault)
	ts = addData(ts, dm.Feature_UatSupportDefault_sql, r.UatSupportDefault)
	ts = addData(ts, dm.Feature_MarketingDefault_sql, r.MarketingDefault)
	ts = addData(ts, dm.Feature_ContingencyDefault_sql, r.ContingencyDefault)
	ts = addData(ts, dm.Feature_DevelopmentFactoredDefault_sql, r.DevelopmentFactoredDefault)
	ts = addData(ts, dm.Feature_Total_sql, r.Total)
	ts = addData(ts, dm.Feature_SYSDbVersion_sql, r.SYSDbVersion)
	ts = addData(ts, dm.Feature_Comments_sql, r.Comments)
	ts = addData(ts, dm.Feature_Description_sql, r.Description)
	ts = addData(ts, dm.Feature_AnalystResource_sql, r.AnalystResource)
	ts = addData(ts, dm.Feature_ProductManagerResource_sql, r.ProductManagerResource)
	ts = addData(ts, dm.Feature_ProjectManagerResource_sql, r.ProjectManagerResource)
	ts = addData(ts, dm.Feature_Training_sql, r.Training)
	ts = addData(ts, dm.Feature_TrainingDefault_sql, r.TrainingDefault)
	ts = addData(ts, dm.Feature_ProfileDefault_sql, r.ProfileDefault)
	ts = addData(ts, dm.Feature_ProfileSelected_sql, r.ProfileSelected)
	ts = addData(ts, dm.Feature_EstimateEffort_sql, r.EstimateEffort)
	ts = addData(ts, dm.Feature_EstimateEffortDefault_sql, r.EstimateEffortDefault)
	ts = addData(ts, dm.Feature_AnalystEstimate_sql, r.AnalystEstimate)
	ts = addData(ts, dm.Feature_TotalDefault_sql, r.TotalDefault)
	
	
	
	
	
	
	
	
	
		
	// 
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := das.INSERT + das.INTO + Feature_QualifiedName
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " "+das.VALUES +"(" + values(ts) + ")"

	Feature_HardDelete(r.FeatureID)
	das.Execute(tsql)

	
		feature_PostPutAction_impl(r.FeatureID,r,usr)
	



	return err

}



// feature_Fetch read all Feature's
func feature_Fetch(tsql string) (int, []dm.Feature, dm.Feature, error) {

	var recItem dm.Feature
	var recList []dm.Feature

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.Feature_SYSId_sql, "0")
	   recItem.FeatureID  = get_String(rec, dm.Feature_FeatureID_sql, "")
	   recItem.EstimationSessionID  = get_String(rec, dm.Feature_EstimationSessionID_sql, "")
	   recItem.ConfidenceCODE  = get_String(rec, dm.Feature_ConfidenceCODE_sql, "")
	   recItem.Name  = get_String(rec, dm.Feature_Name_sql, "")
	   recItem.DevelopmentEstimate  = get_String(rec, dm.Feature_DevelopmentEstimate_sql, "")
	   recItem.DevelopmentFactored  = get_String(rec, dm.Feature_DevelopmentFactored_sql, "")
	   recItem.Requirements  = get_String(rec, dm.Feature_Requirements_sql, "")
	   recItem.AnalystTest  = get_String(rec, dm.Feature_AnalystTest_sql, "")
	   recItem.Documentation  = get_String(rec, dm.Feature_Documentation_sql, "")
	   recItem.Management  = get_String(rec, dm.Feature_Management_sql, "")
	   recItem.UatSupport  = get_String(rec, dm.Feature_UatSupport_sql, "")
	   recItem.Marketing  = get_String(rec, dm.Feature_Marketing_sql, "")
	   recItem.Contingency  = get_String(rec, dm.Feature_Contingency_sql, "")
	   recItem.TrackerID  = get_String(rec, dm.Feature_TrackerID_sql, "")
	   recItem.AdoID  = get_String(rec, dm.Feature_AdoID_sql, "")
	   recItem.FreshdeskID  = get_String(rec, dm.Feature_FreshdeskID_sql, "")
	   recItem.ExtRef  = get_String(rec, dm.Feature_ExtRef_sql, "")
	   recItem.ExtRef2  = get_String(rec, dm.Feature_ExtRef2_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.Feature_SYSCreated_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.Feature_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.Feature_SYSCreatedHost_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.Feature_SYSUpdated_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.Feature_SYSUpdatedBy_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.Feature_SYSUpdatedHost_sql, "")
	   recItem.SYSDeleted  = get_String(rec, dm.Feature_SYSDeleted_sql, "")
	   recItem.SYSDeletedBy  = get_String(rec, dm.Feature_SYSDeletedBy_sql, "")
	   recItem.SYSDeletedHost  = get_String(rec, dm.Feature_SYSDeletedHost_sql, "")
	   recItem.DeveloperResource  = get_String(rec, dm.Feature_DeveloperResource_sql, "")
	   recItem.ApproverResource  = get_String(rec, dm.Feature_ApproverResource_sql, "")
	   recItem.Activity  = get_String(rec, dm.Feature_Activity_sql, "")
	   recItem.OffProfile  = get_String(rec, dm.Feature_OffProfile_sql, "")
	   recItem.OffProfileJustification  = get_String(rec, dm.Feature_OffProfileJustification_sql, "")
	   recItem.SYSActivity  = get_String(rec, dm.Feature_SYSActivity_sql, "")
	   recItem.RequirementsDefault  = get_String(rec, dm.Feature_RequirementsDefault_sql, "")
	   recItem.AnalystTestDefault  = get_String(rec, dm.Feature_AnalystTestDefault_sql, "")
	   recItem.DocumentationDefault  = get_String(rec, dm.Feature_DocumentationDefault_sql, "")
	   recItem.ManagementDefault  = get_String(rec, dm.Feature_ManagementDefault_sql, "")
	   recItem.UatSupportDefault  = get_String(rec, dm.Feature_UatSupportDefault_sql, "")
	   recItem.MarketingDefault  = get_String(rec, dm.Feature_MarketingDefault_sql, "")
	   recItem.ContingencyDefault  = get_String(rec, dm.Feature_ContingencyDefault_sql, "")
	   recItem.DevelopmentFactoredDefault  = get_String(rec, dm.Feature_DevelopmentFactoredDefault_sql, "")
	   recItem.Total  = get_String(rec, dm.Feature_Total_sql, "")
	   recItem.SYSDbVersion  = get_String(rec, dm.Feature_SYSDbVersion_sql, "")
	   recItem.Comments  = get_String(rec, dm.Feature_Comments_sql, "")
	   recItem.Description  = get_String(rec, dm.Feature_Description_sql, "")
	   recItem.AnalystResource  = get_String(rec, dm.Feature_AnalystResource_sql, "")
	   recItem.ProductManagerResource  = get_String(rec, dm.Feature_ProductManagerResource_sql, "")
	   recItem.ProjectManagerResource  = get_String(rec, dm.Feature_ProjectManagerResource_sql, "")
	   recItem.Training  = get_String(rec, dm.Feature_Training_sql, "")
	   recItem.TrainingDefault  = get_String(rec, dm.Feature_TrainingDefault_sql, "")
	   recItem.ProfileDefault  = get_String(rec, dm.Feature_ProfileDefault_sql, "")
	   recItem.ProfileSelected  = get_String(rec, dm.Feature_ProfileSelected_sql, "")
	   recItem.EstimateEffort  = get_String(rec, dm.Feature_EstimateEffort_sql, "")
	   recItem.EstimateEffortDefault  = get_String(rec, dm.Feature_EstimateEffortDefault_sql, "")
	   recItem.AnalystEstimate  = get_String(rec, dm.Feature_AnalystEstimate_sql, "")
	   recItem.TotalDefault  = get_String(rec, dm.Feature_TotalDefault_sql, "")
	
	
	
	
	
	
	
	
	
	
	// If there are fields below, create the methods in dao\Feature_adaptor.go
	   recItem.FeatureID  = Feature_FeatureID_OnFetch_impl (recItem)
	   recItem.EstimationSessionID  = Feature_EstimationSessionID_OnFetch_impl (recItem)
	   recItem.Name  = Feature_Name_OnFetch_impl (recItem)
	   recItem.DevelopmentEstimate  = Feature_DevelopmentEstimate_OnFetch_impl (recItem)
	   recItem.DevelopmentFactored  = Feature_DevelopmentFactored_OnFetch_impl (recItem)
	   recItem.Requirements  = Feature_Requirements_OnFetch_impl (recItem)
	   recItem.AnalystTest  = Feature_AnalystTest_OnFetch_impl (recItem)
	   recItem.Documentation  = Feature_Documentation_OnFetch_impl (recItem)
	   recItem.Management  = Feature_Management_OnFetch_impl (recItem)
	   recItem.Marketing  = Feature_Marketing_OnFetch_impl (recItem)
	   recItem.Contingency  = Feature_Contingency_OnFetch_impl (recItem)
	   recItem.TrackerID  = Feature_TrackerID_OnFetch_impl (recItem)
	   recItem.AdoID  = Feature_AdoID_OnFetch_impl (recItem)
	   recItem.FreshdeskID  = Feature_FreshdeskID_OnFetch_impl (recItem)
	   recItem.ExtRef  = Feature_ExtRef_OnFetch_impl (recItem)
	   recItem.ExtRef2  = Feature_ExtRef2_OnFetch_impl (recItem)
	   recItem.DeveloperResource  = Feature_DeveloperResource_OnFetch_impl (recItem)
	   recItem.ApproverResource  = Feature_ApproverResource_OnFetch_impl (recItem)
	   recItem.OffProfileJustification  = Feature_OffProfileJustification_OnFetch_impl (recItem)
	   recItem.Total  = Feature_Total_OnFetch_impl (recItem)
	   recItem.Training  = Feature_Training_OnFetch_impl (recItem)
	   recItem.EstimateEffort  = Feature_EstimateEffort_OnFetch_impl (recItem)
	   recItem.AnalystEstimate  = Feature_AnalystEstimate_OnFetch_impl (recItem)
	   recItem.OriginName  = Feature_OriginName_OnFetch_impl (recItem)
	   recItem.ProjectName  = Feature_ProjectName_OnFetch_impl (recItem)
	   recItem.OriginID  = Feature_OriginID_OnFetch_impl (recItem)
	   recItem.ProjectID  = Feature_ProjectID_OnFetch_impl (recItem)
	   recItem.ProfileSelectedOld  = Feature_ProfileSelectedOld_OnFetch_impl (recItem)
	   recItem.CCY  = Feature_CCY_OnFetch_impl (recItem)
	   recItem.RoundTo  = Feature_RoundTo_OnFetch_impl (recItem)
	   recItem.OriginCode  = Feature_OriginCode_OnFetch_impl (recItem)
	   recItem.EstimationSessionName  = Feature_EstimationSessionName_OnFetch_impl (recItem)
	// 
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func Feature_NewID(r dm.Feature) string {
	
	id := uuid.New().String()
	
	return id
}



// feature_Fetch read all Feature's
func Feature_New() (int, []dm.Feature, dm.Feature, error) {

	var r = dm.Feature{}
	var rList []dm.Feature
	

	// START
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	r.FeatureID,r.FeatureID_props = Feature_FeatureID_validate_impl (NEW,r.FeatureID,r.FeatureID,r,r.FeatureID_props)
	r.EstimationSessionID,r.EstimationSessionID_props = Feature_EstimationSessionID_validate_impl (NEW,r.FeatureID,r.EstimationSessionID,r,r.EstimationSessionID_props)
	r.Name,r.Name_props = Feature_Name_validate_impl (NEW,r.FeatureID,r.Name,r,r.Name_props)
	r.DevelopmentEstimate,r.DevelopmentEstimate_props = Feature_DevelopmentEstimate_validate_impl (NEW,r.FeatureID,r.DevelopmentEstimate,r,r.DevelopmentEstimate_props)
	r.DevelopmentFactored,r.DevelopmentFactored_props = Feature_DevelopmentFactored_validate_impl (NEW,r.FeatureID,r.DevelopmentFactored,r,r.DevelopmentFactored_props)
	r.Requirements,r.Requirements_props = Feature_Requirements_validate_impl (NEW,r.FeatureID,r.Requirements,r,r.Requirements_props)
	r.AnalystTest,r.AnalystTest_props = Feature_AnalystTest_validate_impl (NEW,r.FeatureID,r.AnalystTest,r,r.AnalystTest_props)
	r.Documentation,r.Documentation_props = Feature_Documentation_validate_impl (NEW,r.FeatureID,r.Documentation,r,r.Documentation_props)
	r.Management,r.Management_props = Feature_Management_validate_impl (NEW,r.FeatureID,r.Management,r,r.Management_props)
	r.Marketing,r.Marketing_props = Feature_Marketing_validate_impl (NEW,r.FeatureID,r.Marketing,r,r.Marketing_props)
	r.Contingency,r.Contingency_props = Feature_Contingency_validate_impl (NEW,r.FeatureID,r.Contingency,r,r.Contingency_props)
	r.TrackerID,r.TrackerID_props = Feature_TrackerID_validate_impl (NEW,r.FeatureID,r.TrackerID,r,r.TrackerID_props)
	r.AdoID,r.AdoID_props = Feature_AdoID_validate_impl (NEW,r.FeatureID,r.AdoID,r,r.AdoID_props)
	r.FreshdeskID,r.FreshdeskID_props = Feature_FreshdeskID_validate_impl (NEW,r.FeatureID,r.FreshdeskID,r,r.FreshdeskID_props)
	r.ExtRef,r.ExtRef_props = Feature_ExtRef_validate_impl (NEW,r.FeatureID,r.ExtRef,r,r.ExtRef_props)
	r.ExtRef2,r.ExtRef2_props = Feature_ExtRef2_validate_impl (NEW,r.FeatureID,r.ExtRef2,r,r.ExtRef2_props)
	r.DeveloperResource,r.DeveloperResource_props = Feature_DeveloperResource_validate_impl (NEW,r.FeatureID,r.DeveloperResource,r,r.DeveloperResource_props)
	r.ApproverResource,r.ApproverResource_props = Feature_ApproverResource_validate_impl (NEW,r.FeatureID,r.ApproverResource,r,r.ApproverResource_props)
	r.OffProfileJustification,r.OffProfileJustification_props = Feature_OffProfileJustification_validate_impl (NEW,r.FeatureID,r.OffProfileJustification,r,r.OffProfileJustification_props)
	r.Total,r.Total_props = Feature_Total_validate_impl (NEW,r.FeatureID,r.Total,r,r.Total_props)
	r.Training,r.Training_props = Feature_Training_validate_impl (NEW,r.FeatureID,r.Training,r,r.Training_props)
	r.EstimateEffort,r.EstimateEffort_props = Feature_EstimateEffort_validate_impl (NEW,r.FeatureID,r.EstimateEffort,r,r.EstimateEffort_props)
	r.AnalystEstimate,r.AnalystEstimate_props = Feature_AnalystEstimate_validate_impl (NEW,r.FeatureID,r.AnalystEstimate,r,r.AnalystEstimate_props)
	r.OriginName,r.OriginName_props = Feature_OriginName_validate_impl (NEW,r.FeatureID,r.OriginName,r,r.OriginName_props)
	r.ProjectName,r.ProjectName_props = Feature_ProjectName_validate_impl (NEW,r.FeatureID,r.ProjectName,r,r.ProjectName_props)
	r.OriginID,r.OriginID_props = Feature_OriginID_validate_impl (NEW,r.FeatureID,r.OriginID,r,r.OriginID_props)
	r.ProjectID,r.ProjectID_props = Feature_ProjectID_validate_impl (NEW,r.FeatureID,r.ProjectID,r,r.ProjectID_props)
	r.ProfileSelectedOld,r.ProfileSelectedOld_props = Feature_ProfileSelectedOld_validate_impl (NEW,r.FeatureID,r.ProfileSelectedOld,r,r.ProfileSelectedOld_props)
	r.CCY,r.CCY_props = Feature_CCY_validate_impl (NEW,r.FeatureID,r.CCY,r,r.CCY_props)
	r.RoundTo,r.RoundTo_props = Feature_RoundTo_validate_impl (NEW,r.FeatureID,r.RoundTo,r,r.RoundTo_props)
	r.OriginCode,r.OriginCode_props = Feature_OriginCode_validate_impl (NEW,r.FeatureID,r.OriginCode,r,r.OriginCode_props)
	r.EstimationSessionName,r.EstimationSessionName_props = Feature_EstimationSessionName_validate_impl (NEW,r.FeatureID,r.EstimationSessionName,r,r.EstimationSessionName_props)
	
	// 
	// Dynamically generated 10/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	rList = append(rList, r)
	return 1, rList, r, nil
}