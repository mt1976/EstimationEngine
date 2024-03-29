package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/featurenew.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : FeatureNew (featurenew)
// Endpoint 	        : FeatureNew (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/03/2023 at 19:24:48
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"fmt"
	"net/http"
	"errors"
	core "github.com/mt1976/ebEstimates/core"
	"github.com/google/uuid"
	das  "github.com/mt1976/ebEstimates/das"
	dm   "github.com/mt1976/ebEstimates/datamodel"
	logs   "github.com/mt1976/ebEstimates/logs"
)

var FeatureNew_SQLbase string
var FeatureNew_QualifiedName string
func init(){
	FeatureNew_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.FeatureNew_SQLTable)
	FeatureNew_SQLbase =  das.SELECTALL + das.FROM + FeatureNew_QualifiedName
}

// FeatureNew_GetList() returns a list of all FeatureNew records
func FeatureNew_GetList() (int, []dm.FeatureNew, error) {
	count, featurenewList, err := FeatureNew_GetListFiltered("")
	return count, featurenewList, err
}

// FeatureNew_GetListFiltered() returns a filtered list of all FeatureNew records
func FeatureNew_GetListFiltered(filter string) (int, []dm.FeatureNew, error) {
	
	count, featurenewList, _ := FeatureNew_GetList_impl(filter)
	
	return count, featurenewList, nil
}



// FeatureNew_GetByID() returns a single FeatureNew record
func FeatureNew_GetByID(id string) (int, dm.FeatureNew, error) {


	 _, featurenewItem, _ := FeatureNew_GetByID_impl(id)
	

	featurenewItem = FeatureNew_PostGet(featurenewItem,id)

	return 1, featurenewItem, nil
}

func FeatureNew_PostGet(featurenewItem dm.FeatureNew,id string) dm.FeatureNew {
	featurenewItem.DeveloperEstimate,featurenewItem.DeveloperEstimate_props = FeatureNew_DeveloperEstimate_validate_impl (GET,id,featurenewItem.DeveloperEstimate,featurenewItem,featurenewItem.DeveloperEstimate_props)
	featurenewItem.DeveloperResource,featurenewItem.DeveloperResource_props = FeatureNew_DeveloperResource_validate_impl (GET,id,featurenewItem.DeveloperResource,featurenewItem,featurenewItem.DeveloperResource_props)
	featurenewItem.ConfidenceCODE,featurenewItem.ConfidenceCODE_props = FeatureNew_ConfidenceCODE_validate_impl (GET,id,featurenewItem.ConfidenceCODE,featurenewItem,featurenewItem.ConfidenceCODE_props)
	featurenewItem.AnalystResource,featurenewItem.AnalystResource_props = FeatureNew_AnalystResource_validate_impl (GET,id,featurenewItem.AnalystResource,featurenewItem,featurenewItem.AnalystResource_props)
	featurenewItem.EstimateEffort,featurenewItem.EstimateEffort_props = FeatureNew_EstimateEffort_validate_impl (GET,id,featurenewItem.EstimateEffort,featurenewItem,featurenewItem.EstimateEffort_props)
	featurenewItem.EstimationSessionName,featurenewItem.EstimationSessionName_props = FeatureNew_EstimationSessionName_validate_impl (GET,id,featurenewItem.EstimationSessionName,featurenewItem,featurenewItem.EstimationSessionName_props)
	featurenewItem.ProjectID,featurenewItem.ProjectID_props = FeatureNew_ProjectID_validate_impl (GET,id,featurenewItem.ProjectID,featurenewItem,featurenewItem.ProjectID_props)
	featurenewItem.OriginCode,featurenewItem.OriginCode_props = FeatureNew_OriginCode_validate_impl (GET,id,featurenewItem.OriginCode,featurenewItem,featurenewItem.OriginCode_props)
	featurenewItem.OriginID,featurenewItem.OriginID_props = FeatureNew_OriginID_validate_impl (GET,id,featurenewItem.OriginID,featurenewItem,featurenewItem.OriginID_props)
	featurenewItem.ProjectName,featurenewItem.ProjectName_props = FeatureNew_ProjectName_validate_impl (GET,id,featurenewItem.ProjectName,featurenewItem,featurenewItem.ProjectName_props)

	return featurenewItem
}




// FeatureNew_DeleteByID() deletes a single FeatureNew record
func FeatureNew_Delete(id string) {
	FeatureNew_Delete_impl(id)
}

// FeatureNew_HardDelete(id string) soft deletes a single FeatureNew record
func FeatureNew_HardDelete(id string) {
			// Uses Hard Delete
		object_Table := FeatureNew_QualifiedName
		tsql := das.DELETE + das.FROM + object_Table
		tsql = tsql + " " + das.WHERE + dm.FeatureNew_SQLSearchID + das.EQ + das.ID(id)
		das.Execute(tsql)
		//if err != nil {
		//	logs.Error("FeatureNew_SoftDelete()",err)
		//}
}


// FeatureNew_Store() saves/stores a FeatureNew record to the database
func FeatureNew_Store(r dm.FeatureNew,req *http.Request) (dm.FeatureNew,error) {

	r, err := FeatureNew_Validate(r)
	if err == nil {
		err = featurenew_Save(r, Audit_User(req))
	} else {
		logs.Information("FeatureNew_Store()", err.Error())
	}

	return r, err
}

// FeatureNew_StoreSystem() saves/stores a FeatureNew record to the database
func FeatureNew_StoreSystem(r dm.FeatureNew) (dm.FeatureNew,error) {
	
	r, err := FeatureNew_Validate(r)
	if err == nil {
		err = featurenew_Save(r, Audit_Host())
	} else {
		logs.Information("FeatureNew_Store()", err.Error())
	}

	return r, err
}

// FeatureNew_StoreProcess() saves/stores a FeatureNew record to the database
func FeatureNew_StoreProcess(r dm.FeatureNew, operator string) (dm.FeatureNew,error) {
	
	r, err := FeatureNew_Validate(r)
	if err == nil {
		err = featurenew_Save(r, operator)
	} else {
		logs.Information("FeatureNew_StoreProcess()", err.Error())
	}

	return r, err
}

// FeatureNew_Validate() validates for saves/stores a FeatureNew record to the database
func FeatureNew_Validate(r dm.FeatureNew) (dm.FeatureNew, error) {
	var err error
	r.DeveloperEstimate,r.DeveloperEstimate_props = FeatureNew_DeveloperEstimate_validate_impl (PUT,r.ID,r.DeveloperEstimate,r,r.DeveloperEstimate_props)
	if r.DeveloperEstimate_props.MsgMessage != "" {
		err = errors.New(r.DeveloperEstimate_props.MsgMessage)
	}
	r.DeveloperResource,r.DeveloperResource_props = FeatureNew_DeveloperResource_validate_impl (PUT,r.ID,r.DeveloperResource,r,r.DeveloperResource_props)
	if r.DeveloperResource_props.MsgMessage != "" {
		err = errors.New(r.DeveloperResource_props.MsgMessage)
	}
	r.ConfidenceCODE,r.ConfidenceCODE_props = FeatureNew_ConfidenceCODE_validate_impl (PUT,r.ID,r.ConfidenceCODE,r,r.ConfidenceCODE_props)
	if r.ConfidenceCODE_props.MsgMessage != "" {
		err = errors.New(r.ConfidenceCODE_props.MsgMessage)
	}
	r.AnalystResource,r.AnalystResource_props = FeatureNew_AnalystResource_validate_impl (PUT,r.ID,r.AnalystResource,r,r.AnalystResource_props)
	if r.AnalystResource_props.MsgMessage != "" {
		err = errors.New(r.AnalystResource_props.MsgMessage)
	}
	r.EstimateEffort,r.EstimateEffort_props = FeatureNew_EstimateEffort_validate_impl (PUT,r.ID,r.EstimateEffort,r,r.EstimateEffort_props)
	if r.EstimateEffort_props.MsgMessage != "" {
		err = errors.New(r.EstimateEffort_props.MsgMessage)
	}
	r.EstimationSessionName,r.EstimationSessionName_props = FeatureNew_EstimationSessionName_validate_impl (PUT,r.ID,r.EstimationSessionName,r,r.EstimationSessionName_props)
	if r.EstimationSessionName_props.MsgMessage != "" {
		err = errors.New(r.EstimationSessionName_props.MsgMessage)
	}
	r.ProjectID,r.ProjectID_props = FeatureNew_ProjectID_validate_impl (PUT,r.ID,r.ProjectID,r,r.ProjectID_props)
	if r.ProjectID_props.MsgMessage != "" {
		err = errors.New(r.ProjectID_props.MsgMessage)
	}
	r.OriginCode,r.OriginCode_props = FeatureNew_OriginCode_validate_impl (PUT,r.ID,r.OriginCode,r,r.OriginCode_props)
	if r.OriginCode_props.MsgMessage != "" {
		err = errors.New(r.OriginCode_props.MsgMessage)
	}
	r.OriginID,r.OriginID_props = FeatureNew_OriginID_validate_impl (PUT,r.ID,r.OriginID,r,r.OriginID_props)
	if r.OriginID_props.MsgMessage != "" {
		err = errors.New(r.OriginID_props.MsgMessage)
	}
	r.ProjectName,r.ProjectName_props = FeatureNew_ProjectName_validate_impl (PUT,r.ID,r.ProjectName,r,r.ProjectName_props)
	if r.ProjectName_props.MsgMessage != "" {
		err = errors.New(r.ProjectName_props.MsgMessage)
	}

	// Cross Validation
	var errVal error
	r, _, errVal = FeatureNew_ObjectValidation_impl(PUT, r.ID, r)
	if errVal != nil {
		err = errVal
	}
	return r,err
}
//

// featurenew_Save() saves/stores a FeatureNew record to the database
func featurenew_Save(r dm.FeatureNew,usr string) error {

    var err error

	if len(r.ID) == 0 {
		r.ID = FeatureNew_NewID(r)
	}

// If there are fields below, create the methods in dao\featurenew_impl.go
    r.DeveloperEstimate,err = FeatureNew_DeveloperEstimate_OnStore_impl (r.DeveloperEstimate,r,usr)
    r.DeveloperResource,err = FeatureNew_DeveloperResource_OnStore_impl (r.DeveloperResource,r,usr)
    r.ConfidenceCODE,err = FeatureNew_ConfidenceCODE_OnStore_impl (r.ConfidenceCODE,r,usr)
    r.AnalystResource,err = FeatureNew_AnalystResource_OnStore_impl (r.AnalystResource,r,usr)
    r.EstimateEffort,err = FeatureNew_EstimateEffort_OnStore_impl (r.EstimateEffort,r,usr)
    r.EstimationSessionName,err = FeatureNew_EstimationSessionName_OnStore_impl (r.EstimationSessionName,r,usr)
    r.ProjectID,err = FeatureNew_ProjectID_OnStore_impl (r.ProjectID,r,usr)
    r.OriginCode,err = FeatureNew_OriginCode_OnStore_impl (r.OriginCode,r,usr)
    r.OriginID,err = FeatureNew_OriginID_OnStore_impl (r.OriginID,r,usr)
    r.ProjectName,err = FeatureNew_ProjectName_OnStore_impl (r.ProjectName,r,usr)

logs.Storing("FeatureNew",fmt.Sprintf("%v", r))

// Please Create Functions Below in the adaptor/FeatureNew_impl.go file
	err1 := FeatureNew_Delete_impl(r.ID)
	err2 := FeatureNew_Update_impl(r.ID,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



	// featurenew_Fetch is not required as GetByID, GetAll etc... have been diverted to _impl
	


func FeatureNew_NewID(r dm.FeatureNew) string {
	
	id := uuid.New().String()
	
	return id
}



// featurenew_Fetch read all FeatureNew's
func FeatureNew_New() (int, []dm.FeatureNew, dm.FeatureNew, error) {

	var r = dm.FeatureNew{}
	var rList []dm.FeatureNew
	

	// START
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	r.DeveloperEstimate,r.DeveloperEstimate_props = FeatureNew_DeveloperEstimate_validate_impl (NEW,r.ID,r.DeveloperEstimate,r,r.DeveloperEstimate_props)
	r.DeveloperResource,r.DeveloperResource_props = FeatureNew_DeveloperResource_validate_impl (NEW,r.ID,r.DeveloperResource,r,r.DeveloperResource_props)
	r.ConfidenceCODE,r.ConfidenceCODE_props = FeatureNew_ConfidenceCODE_validate_impl (NEW,r.ID,r.ConfidenceCODE,r,r.ConfidenceCODE_props)
	r.AnalystResource,r.AnalystResource_props = FeatureNew_AnalystResource_validate_impl (NEW,r.ID,r.AnalystResource,r,r.AnalystResource_props)
	r.EstimateEffort,r.EstimateEffort_props = FeatureNew_EstimateEffort_validate_impl (NEW,r.ID,r.EstimateEffort,r,r.EstimateEffort_props)
	r.EstimationSessionName,r.EstimationSessionName_props = FeatureNew_EstimationSessionName_validate_impl (NEW,r.ID,r.EstimationSessionName,r,r.EstimationSessionName_props)
	r.ProjectID,r.ProjectID_props = FeatureNew_ProjectID_validate_impl (NEW,r.ID,r.ProjectID,r,r.ProjectID_props)
	r.OriginCode,r.OriginCode_props = FeatureNew_OriginCode_validate_impl (NEW,r.ID,r.OriginCode,r,r.OriginCode_props)
	r.OriginID,r.OriginID_props = FeatureNew_OriginID_validate_impl (NEW,r.ID,r.OriginID,r,r.OriginID_props)
	r.ProjectName,r.ProjectName_props = FeatureNew_ProjectName_validate_impl (NEW,r.ID,r.ProjectName,r,r.ProjectName_props)
	
	// 
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	rList = append(rList, r)
	return 1, rList, r, nil
}