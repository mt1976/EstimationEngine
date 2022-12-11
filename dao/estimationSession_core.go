package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/estimationsession.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : EstimationSession (estimationsession)
// Endpoint 	        : EstimationSession (EstimationSessionID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 11/12/2022 at 14:24:08
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (

	"fmt"
	"net/http"
	core "github.com/mt1976/ebEstimates/core"
	"github.com/google/uuid"
	das  "github.com/mt1976/ebEstimates/das"
	
		 
		// Does Lookup
		adaptor   "github.com/mt1976/ebEstimates/adaptor"
	
	dm   "github.com/mt1976/ebEstimates/datamodel"
	logs   "github.com/mt1976/ebEstimates/logs"
)

var EstimationSession_SQLbase string
func init(){
	EstimationSession_SQLbase =  core.DB_SELECT + " "+ core.DB_ALL + " " + core.DB_FROM + " " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.EstimationSession_SQLTable)
}

// EstimationSession_GetList() returns a list of all EstimationSession records
func EstimationSession_GetList() (int, []dm.EstimationSession, error) {
	
	tsql := EstimationSession_SQLbase
	count, estimationsessionList, _, _ := estimationsession_Fetch(tsql)
	
	return count, estimationsessionList, nil
}


// EstimationSession_GetLookup() returns a lookup list of all EstimationSession items in lookup format
func EstimationSession_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, estimationsessionList, _ := EstimationSession_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: estimationsessionList[i].EstimationSessionID, Name: estimationsessionList[i].Name})
	}
	return returnList
}


// EstimationSession_GetByID() returns a single EstimationSession record
func EstimationSession_GetByID(id string) (int, dm.EstimationSession, error) {


	tsql := EstimationSession_SQLbase
	tsql = tsql + " " + core.DB_WHERE + " " + dm.EstimationSession_SQLSearchID + core.DB_EQ + "'" + id + "'"
	_, _, estimationsessionItem, _ := estimationsession_Fetch(tsql)

	// START
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	estimationsessionItem.Origin,estimationsessionItem.Origin_props = adaptor.EstimationSession_Origin_impl (adaptor.GET,id,estimationsessionItem.Origin,estimationsessionItem,estimationsessionItem.Origin_props)
	estimationsessionItem.OriginStateID,estimationsessionItem.OriginStateID_props = adaptor.EstimationSession_OriginStateID_impl (adaptor.GET,id,estimationsessionItem.OriginStateID,estimationsessionItem,estimationsessionItem.OriginStateID_props)
	estimationsessionItem.OriginState,estimationsessionItem.OriginState_props = adaptor.EstimationSession_OriginState_impl (adaptor.GET,id,estimationsessionItem.OriginState,estimationsessionItem,estimationsessionItem.OriginState_props)
	estimationsessionItem.OriginDocTypeID,estimationsessionItem.OriginDocTypeID_props = adaptor.EstimationSession_OriginDocTypeID_impl (adaptor.GET,id,estimationsessionItem.OriginDocTypeID,estimationsessionItem,estimationsessionItem.OriginDocTypeID_props)
	estimationsessionItem.OriginDocType,estimationsessionItem.OriginDocType_props = adaptor.EstimationSession_OriginDocType_impl (adaptor.GET,id,estimationsessionItem.OriginDocType,estimationsessionItem,estimationsessionItem.OriginDocType_props)
	estimationsessionItem.OriginCode,estimationsessionItem.OriginCode_props = adaptor.EstimationSession_OriginCode_impl (adaptor.GET,id,estimationsessionItem.OriginCode,estimationsessionItem,estimationsessionItem.OriginCode_props)
	estimationsessionItem.OriginName,estimationsessionItem.OriginName_props = adaptor.EstimationSession_OriginName_impl (adaptor.GET,id,estimationsessionItem.OriginName,estimationsessionItem,estimationsessionItem.OriginName_props)
	estimationsessionItem.OriginRate,estimationsessionItem.OriginRate_props = adaptor.EstimationSession_OriginRate_impl (adaptor.GET,id,estimationsessionItem.OriginRate,estimationsessionItem,estimationsessionItem.OriginRate_props)
	estimationsessionItem.ProjectProfileID,estimationsessionItem.ProjectProfileID_props = adaptor.EstimationSession_ProjectProfileID_impl (adaptor.GET,id,estimationsessionItem.ProjectProfileID,estimationsessionItem,estimationsessionItem.ProjectProfileID_props)
	estimationsessionItem.ProjectProfile,estimationsessionItem.ProjectProfile_props = adaptor.EstimationSession_ProjectProfile_impl (adaptor.GET,id,estimationsessionItem.ProjectProfile,estimationsessionItem,estimationsessionItem.ProjectProfile_props)
	estimationsessionItem.ProjectDefaultReleases,estimationsessionItem.ProjectDefaultReleases_props = adaptor.EstimationSession_ProjectDefaultReleases_impl (adaptor.GET,id,estimationsessionItem.ProjectDefaultReleases,estimationsessionItem,estimationsessionItem.ProjectDefaultReleases_props)
	estimationsessionItem.ProjectDefaultReleaseHours,estimationsessionItem.ProjectDefaultReleaseHours_props = adaptor.EstimationSession_ProjectDefaultReleaseHours_impl (adaptor.GET,id,estimationsessionItem.ProjectDefaultReleaseHours,estimationsessionItem,estimationsessionItem.ProjectDefaultReleaseHours_props)
	estimationsessionItem.ProjectBlendedRate,estimationsessionItem.ProjectBlendedRate_props = adaptor.EstimationSession_ProjectBlendedRate_impl (adaptor.GET,id,estimationsessionItem.ProjectBlendedRate,estimationsessionItem,estimationsessionItem.ProjectBlendedRate_props)
	estimationsessionItem.ProjectStateID,estimationsessionItem.ProjectStateID_props = adaptor.EstimationSession_ProjectStateID_impl (adaptor.GET,id,estimationsessionItem.ProjectStateID,estimationsessionItem,estimationsessionItem.ProjectStateID_props)
	estimationsessionItem.ProjectState,estimationsessionItem.ProjectState_props = adaptor.EstimationSession_ProjectState_impl (adaptor.GET,id,estimationsessionItem.ProjectState,estimationsessionItem,estimationsessionItem.ProjectState_props)
	estimationsessionItem.ProjectName,estimationsessionItem.ProjectName_props = adaptor.EstimationSession_ProjectName_impl (adaptor.GET,id,estimationsessionItem.ProjectName,estimationsessionItem,estimationsessionItem.ProjectName_props)
	estimationsessionItem.ProjectStartDate,estimationsessionItem.ProjectStartDate_props = adaptor.EstimationSession_ProjectStartDate_impl (adaptor.GET,id,estimationsessionItem.ProjectStartDate,estimationsessionItem,estimationsessionItem.ProjectStartDate_props)
	estimationsessionItem.ProjectEndDate,estimationsessionItem.ProjectEndDate_props = adaptor.EstimationSession_ProjectEndDate_impl (adaptor.GET,id,estimationsessionItem.ProjectEndDate,estimationsessionItem,estimationsessionItem.ProjectEndDate_props)
	estimationsessionItem.ProfileSupportUpliftPerc,estimationsessionItem.ProfileSupportUpliftPerc_props = adaptor.EstimationSession_ProfileSupportUpliftPerc_impl (adaptor.GET,id,estimationsessionItem.ProfileSupportUpliftPerc,estimationsessionItem,estimationsessionItem.ProfileSupportUpliftPerc_props)
	estimationsessionItem.CCY,estimationsessionItem.CCY_props = adaptor.EstimationSession_CCY_impl (adaptor.GET,id,estimationsessionItem.CCY,estimationsessionItem,estimationsessionItem.CCY_props)
	estimationsessionItem.CCYCode,estimationsessionItem.CCYCode_props = adaptor.EstimationSession_CCYCode_impl (adaptor.GET,id,estimationsessionItem.CCYCode,estimationsessionItem,estimationsessionItem.CCYCode_props)
	estimationsessionItem.EffortTotal,estimationsessionItem.EffortTotal_props = adaptor.EstimationSession_EffortTotal_impl (adaptor.GET,id,estimationsessionItem.EffortTotal,estimationsessionItem,estimationsessionItem.EffortTotal_props)
	estimationsessionItem.FreshDeskURI,estimationsessionItem.FreshDeskURI_props = adaptor.EstimationSession_FreshDeskURI_impl (adaptor.GET,id,estimationsessionItem.FreshDeskURI,estimationsessionItem,estimationsessionItem.FreshDeskURI_props)
	estimationsessionItem.ADOURI,estimationsessionItem.ADOURI_props = adaptor.EstimationSession_ADOURI_impl (adaptor.GET,id,estimationsessionItem.ADOURI,estimationsessionItem,estimationsessionItem.ADOURI_props)
	estimationsessionItem.NoActiveFeatures,estimationsessionItem.NoActiveFeatures_props = adaptor.EstimationSession_NoActiveFeatures_impl (adaptor.GET,id,estimationsessionItem.NoActiveFeatures,estimationsessionItem,estimationsessionItem.NoActiveFeatures_props)
	// 
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, estimationsessionItem, nil
}



// EstimationSession_DeleteByID() deletes a single EstimationSession record
func EstimationSession_Delete(id string) {




// Uses Hard Delete
	object_Table := core.GetSQLSchema(core.ApplicationPropertiesDB) + "." + dm.EstimationSession_SQLTable
	tsql := core.DB_DELETE+" "+core.DB_FROM+" " + object_Table
	tsql = tsql + " " + core.DB_WHERE + " " + dm.EstimationSession_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
	
	

	
}














	
	// EstimationSession_SoftDeleteByID() soft deletes a single EstimationSession record
func EstimationSession_SoftDelete(id string) {
	//Uses Soft Delete
		_,estimationsessionItem,_ := EstimationSession_GetByID(id)
		estimationsessionItem.SYSDeletedBy = Audit_Update("", Audit_Host())
		estimationsessionItem.SYSDeleted = Audit_Update("", Audit_TimeStamp())
		estimationsessionItem.SYSDeletedHost = Audit_Update("", Audit_Host())
		EstimationSession_StoreSystem(estimationsessionItem)
}
	


// EstimationSession_Store() saves/stores a EstimationSession record to the database
func EstimationSession_Store(r dm.EstimationSession,req *http.Request) error {

	r, err := EstimationSession_Validate(r)
	if err == nil {
		err = estimationsession_Save(r, Audit_User(req))
	} else {
		logs.Information("EstimationSession_Store()", err.Error())
	}

	return err
}

// EstimationSession_StoreSystem() saves/stores a EstimationSession record to the database
func EstimationSession_StoreSystem(r dm.EstimationSession) error {
	
	r, err := EstimationSession_Validate(r)
	if err == nil {
		err = estimationsession_Save(r, Audit_Host())
	} else {
		logs.Information("EstimationSession_Store()", err.Error())
	}

	return err
}

// EstimationSession_Validate() validates for saves/stores a EstimationSession record to the database
func EstimationSession_Validate(r dm.EstimationSession) (dm.EstimationSession, error) {
	var err error
	// START
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	r.Origin,r.Origin_props = adaptor.EstimationSession_Origin_impl (adaptor.PUT,r.EstimationSessionID,r.Origin,r,r.Origin_props)
	r.OriginStateID,r.OriginStateID_props = adaptor.EstimationSession_OriginStateID_impl (adaptor.PUT,r.EstimationSessionID,r.OriginStateID,r,r.OriginStateID_props)
	r.OriginState,r.OriginState_props = adaptor.EstimationSession_OriginState_impl (adaptor.PUT,r.EstimationSessionID,r.OriginState,r,r.OriginState_props)
	r.OriginDocTypeID,r.OriginDocTypeID_props = adaptor.EstimationSession_OriginDocTypeID_impl (adaptor.PUT,r.EstimationSessionID,r.OriginDocTypeID,r,r.OriginDocTypeID_props)
	r.OriginDocType,r.OriginDocType_props = adaptor.EstimationSession_OriginDocType_impl (adaptor.PUT,r.EstimationSessionID,r.OriginDocType,r,r.OriginDocType_props)
	r.OriginCode,r.OriginCode_props = adaptor.EstimationSession_OriginCode_impl (adaptor.PUT,r.EstimationSessionID,r.OriginCode,r,r.OriginCode_props)
	r.OriginName,r.OriginName_props = adaptor.EstimationSession_OriginName_impl (adaptor.PUT,r.EstimationSessionID,r.OriginName,r,r.OriginName_props)
	r.OriginRate,r.OriginRate_props = adaptor.EstimationSession_OriginRate_impl (adaptor.PUT,r.EstimationSessionID,r.OriginRate,r,r.OriginRate_props)
	r.ProjectProfileID,r.ProjectProfileID_props = adaptor.EstimationSession_ProjectProfileID_impl (adaptor.PUT,r.EstimationSessionID,r.ProjectProfileID,r,r.ProjectProfileID_props)
	r.ProjectProfile,r.ProjectProfile_props = adaptor.EstimationSession_ProjectProfile_impl (adaptor.PUT,r.EstimationSessionID,r.ProjectProfile,r,r.ProjectProfile_props)
	r.ProjectDefaultReleases,r.ProjectDefaultReleases_props = adaptor.EstimationSession_ProjectDefaultReleases_impl (adaptor.PUT,r.EstimationSessionID,r.ProjectDefaultReleases,r,r.ProjectDefaultReleases_props)
	r.ProjectDefaultReleaseHours,r.ProjectDefaultReleaseHours_props = adaptor.EstimationSession_ProjectDefaultReleaseHours_impl (adaptor.PUT,r.EstimationSessionID,r.ProjectDefaultReleaseHours,r,r.ProjectDefaultReleaseHours_props)
	r.ProjectBlendedRate,r.ProjectBlendedRate_props = adaptor.EstimationSession_ProjectBlendedRate_impl (adaptor.PUT,r.EstimationSessionID,r.ProjectBlendedRate,r,r.ProjectBlendedRate_props)
	r.ProjectStateID,r.ProjectStateID_props = adaptor.EstimationSession_ProjectStateID_impl (adaptor.PUT,r.EstimationSessionID,r.ProjectStateID,r,r.ProjectStateID_props)
	r.ProjectState,r.ProjectState_props = adaptor.EstimationSession_ProjectState_impl (adaptor.PUT,r.EstimationSessionID,r.ProjectState,r,r.ProjectState_props)
	r.ProjectName,r.ProjectName_props = adaptor.EstimationSession_ProjectName_impl (adaptor.PUT,r.EstimationSessionID,r.ProjectName,r,r.ProjectName_props)
	r.ProjectStartDate,r.ProjectStartDate_props = adaptor.EstimationSession_ProjectStartDate_impl (adaptor.PUT,r.EstimationSessionID,r.ProjectStartDate,r,r.ProjectStartDate_props)
	r.ProjectEndDate,r.ProjectEndDate_props = adaptor.EstimationSession_ProjectEndDate_impl (adaptor.PUT,r.EstimationSessionID,r.ProjectEndDate,r,r.ProjectEndDate_props)
	r.ProfileSupportUpliftPerc,r.ProfileSupportUpliftPerc_props = adaptor.EstimationSession_ProfileSupportUpliftPerc_impl (adaptor.PUT,r.EstimationSessionID,r.ProfileSupportUpliftPerc,r,r.ProfileSupportUpliftPerc_props)
	r.CCY,r.CCY_props = adaptor.EstimationSession_CCY_impl (adaptor.PUT,r.EstimationSessionID,r.CCY,r,r.CCY_props)
	r.CCYCode,r.CCYCode_props = adaptor.EstimationSession_CCYCode_impl (adaptor.PUT,r.EstimationSessionID,r.CCYCode,r,r.CCYCode_props)
	r.EffortTotal,r.EffortTotal_props = adaptor.EstimationSession_EffortTotal_impl (adaptor.PUT,r.EstimationSessionID,r.EffortTotal,r,r.EffortTotal_props)
	r.FreshDeskURI,r.FreshDeskURI_props = adaptor.EstimationSession_FreshDeskURI_impl (adaptor.PUT,r.EstimationSessionID,r.FreshDeskURI,r,r.FreshDeskURI_props)
	r.ADOURI,r.ADOURI_props = adaptor.EstimationSession_ADOURI_impl (adaptor.PUT,r.EstimationSessionID,r.ADOURI,r,r.ADOURI_props)
	r.NoActiveFeatures,r.NoActiveFeatures_props = adaptor.EstimationSession_NoActiveFeatures_impl (adaptor.PUT,r.EstimationSessionID,r.NoActiveFeatures,r,r.NoActiveFeatures_props)
	// 
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	
	// START
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	r, _, err = adaptor.EstimationSession_ObjectValidation_impl(adaptor.PUT, r.EstimationSessionID, r)
	// 
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	

	return r,err
}
//

// estimationsession_Save() saves/stores a EstimationSession record to the database
func estimationsession_Save(r dm.EstimationSession,usr string) error {

    var err error



	

	if len(r.EstimationSessionID) == 0 {
		r.EstimationSessionID = EstimationSession_NewID(r)
	}

// If there are fields below, create the methods in dao\estimationsession_impl.go





































  r.Origin,err = adaptor.EstimationSession_Origin_OnStore_impl (r.Origin,r,usr)
  r.OriginStateID,err = adaptor.EstimationSession_OriginStateID_OnStore_impl (r.OriginStateID,r,usr)
  r.OriginState,err = adaptor.EstimationSession_OriginState_OnStore_impl (r.OriginState,r,usr)
  r.OriginDocTypeID,err = adaptor.EstimationSession_OriginDocTypeID_OnStore_impl (r.OriginDocTypeID,r,usr)
  r.OriginDocType,err = adaptor.EstimationSession_OriginDocType_OnStore_impl (r.OriginDocType,r,usr)
  r.OriginCode,err = adaptor.EstimationSession_OriginCode_OnStore_impl (r.OriginCode,r,usr)
  r.OriginName,err = adaptor.EstimationSession_OriginName_OnStore_impl (r.OriginName,r,usr)
  r.OriginRate,err = adaptor.EstimationSession_OriginRate_OnStore_impl (r.OriginRate,r,usr)
  r.ProjectProfileID,err = adaptor.EstimationSession_ProjectProfileID_OnStore_impl (r.ProjectProfileID,r,usr)
  r.ProjectProfile,err = adaptor.EstimationSession_ProjectProfile_OnStore_impl (r.ProjectProfile,r,usr)
  r.ProjectDefaultReleases,err = adaptor.EstimationSession_ProjectDefaultReleases_OnStore_impl (r.ProjectDefaultReleases,r,usr)
  r.ProjectDefaultReleaseHours,err = adaptor.EstimationSession_ProjectDefaultReleaseHours_OnStore_impl (r.ProjectDefaultReleaseHours,r,usr)
  r.ProjectBlendedRate,err = adaptor.EstimationSession_ProjectBlendedRate_OnStore_impl (r.ProjectBlendedRate,r,usr)
  r.ProjectStateID,err = adaptor.EstimationSession_ProjectStateID_OnStore_impl (r.ProjectStateID,r,usr)
  r.ProjectState,err = adaptor.EstimationSession_ProjectState_OnStore_impl (r.ProjectState,r,usr)
  r.ProjectName,err = adaptor.EstimationSession_ProjectName_OnStore_impl (r.ProjectName,r,usr)
  r.ProjectStartDate,err = adaptor.EstimationSession_ProjectStartDate_OnStore_impl (r.ProjectStartDate,r,usr)
  r.ProjectEndDate,err = adaptor.EstimationSession_ProjectEndDate_OnStore_impl (r.ProjectEndDate,r,usr)
  r.ProfileSupportUpliftPerc,err = adaptor.EstimationSession_ProfileSupportUpliftPerc_OnStore_impl (r.ProfileSupportUpliftPerc,r,usr)
  r.CCY,err = adaptor.EstimationSession_CCY_OnStore_impl (r.CCY,r,usr)
  r.CCYCode,err = adaptor.EstimationSession_CCYCode_OnStore_impl (r.CCYCode,r,usr)
  r.EffortTotal,err = adaptor.EstimationSession_EffortTotal_OnStore_impl (r.EffortTotal,r,usr)
  r.FreshDeskURI,err = adaptor.EstimationSession_FreshDeskURI_OnStore_impl (r.FreshDeskURI,r,usr)
  r.ADOURI,err = adaptor.EstimationSession_ADOURI_OnStore_impl (r.ADOURI,r,usr)
  r.NoActiveFeatures,err = adaptor.EstimationSession_NoActiveFeatures_OnStore_impl (r.NoActiveFeatures,r,usr)


	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	r.SYSDbVersion = core.DB_Version()
	
logs.Storing("EstimationSession",fmt.Sprintf("%v", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.EstimationSession_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.EstimationSession_EstimationSessionID_sql, r.EstimationSessionID)
	ts = addData(ts, dm.EstimationSession_ProjectID_sql, r.ProjectID)
	ts = addData(ts, dm.EstimationSession_EstimationStateID_sql, r.EstimationStateID)
	ts = addData(ts, dm.EstimationSession_Notes_sql, r.Notes)
	ts = addData(ts, dm.EstimationSession_Releases_sql, r.Releases)
	ts = addData(ts, dm.EstimationSession_Total_sql, r.Total)
	ts = addData(ts, dm.EstimationSession_Contingency_sql, r.Contingency)
	ts = addData(ts, dm.EstimationSession_ReqDays_sql, r.ReqDays)
	ts = addData(ts, dm.EstimationSession_RegCost_sql, r.RegCost)
	ts = addData(ts, dm.EstimationSession_ImpDays_sql, r.ImpDays)
	ts = addData(ts, dm.EstimationSession_ImpCost_sql, r.ImpCost)
	ts = addData(ts, dm.EstimationSession_UatDays_sql, r.UatDays)
	ts = addData(ts, dm.EstimationSession_UatCost_sql, r.UatCost)
	ts = addData(ts, dm.EstimationSession_MgtDays_sql, r.MgtDays)
	ts = addData(ts, dm.EstimationSession_MgtCost_sql, r.MgtCost)
	ts = addData(ts, dm.EstimationSession_RelDays_sql, r.RelDays)
	ts = addData(ts, dm.EstimationSession_RelCost_sql, r.RelCost)
	ts = addData(ts, dm.EstimationSession_SupportUplift_sql, r.SupportUplift)
	ts = addData(ts, dm.EstimationSession_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.EstimationSession_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.EstimationSession_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.EstimationSession_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.EstimationSession_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.EstimationSession_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.EstimationSession_SYSDeleted_sql, r.SYSDeleted)
	ts = addData(ts, dm.EstimationSession_SYSDeletedBy_sql, r.SYSDeletedBy)
	ts = addData(ts, dm.EstimationSession_SYSDeletedHost_sql, r.SYSDeletedHost)
	ts = addData(ts, dm.EstimationSession_Name_sql, r.Name)
	ts = addData(ts, dm.EstimationSession_AdoID_sql, r.AdoID)
	ts = addData(ts, dm.EstimationSession_FreshdeskID_sql, r.FreshdeskID)
	ts = addData(ts, dm.EstimationSession_TrackerID_sql, r.TrackerID)
	ts = addData(ts, dm.EstimationSession_EstRef_sql, r.EstRef)
	ts = addData(ts, dm.EstimationSession_ExtRef_sql, r.ExtRef)
	ts = addData(ts, dm.EstimationSession_SYSActivity_sql, r.SYSActivity)
	ts = addData(ts, dm.EstimationSession_SYSDbVersion_sql, r.SYSDbVersion)
	ts = addData(ts, dm.EstimationSession_Comments_sql, r.Comments)
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
		
	// 
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := core.DB_INSERT + " " + core.DB_INTO + " " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.EstimationSession_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " "+core.DB_VALUES +" (" + values(ts) + ")"

	EstimationSession_Delete(r.EstimationSessionID)
	das.Execute(tsql)



	return err

}



// estimationsession_Fetch read all EstimationSession's
func estimationsession_Fetch(tsql string) (int, []dm.EstimationSession, dm.EstimationSession, error) {

	var recItem dm.EstimationSession
	var recList []dm.EstimationSession

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.EstimationSession_SYSId_sql, "0")
	   recItem.EstimationSessionID  = get_String(rec, dm.EstimationSession_EstimationSessionID_sql, "")
	   recItem.ProjectID  = get_String(rec, dm.EstimationSession_ProjectID_sql, "")
	   recItem.EstimationStateID  = get_String(rec, dm.EstimationSession_EstimationStateID_sql, "")
	   recItem.Notes  = get_String(rec, dm.EstimationSession_Notes_sql, "")
	   recItem.Releases  = get_String(rec, dm.EstimationSession_Releases_sql, "")
	   recItem.Total  = get_String(rec, dm.EstimationSession_Total_sql, "")
	   recItem.Contingency  = get_String(rec, dm.EstimationSession_Contingency_sql, "")
	   recItem.ReqDays  = get_String(rec, dm.EstimationSession_ReqDays_sql, "")
	   recItem.RegCost  = get_String(rec, dm.EstimationSession_RegCost_sql, "")
	   recItem.ImpDays  = get_String(rec, dm.EstimationSession_ImpDays_sql, "")
	   recItem.ImpCost  = get_String(rec, dm.EstimationSession_ImpCost_sql, "")
	   recItem.UatDays  = get_String(rec, dm.EstimationSession_UatDays_sql, "")
	   recItem.UatCost  = get_String(rec, dm.EstimationSession_UatCost_sql, "")
	   recItem.MgtDays  = get_String(rec, dm.EstimationSession_MgtDays_sql, "")
	   recItem.MgtCost  = get_String(rec, dm.EstimationSession_MgtCost_sql, "")
	   recItem.RelDays  = get_String(rec, dm.EstimationSession_RelDays_sql, "")
	   recItem.RelCost  = get_String(rec, dm.EstimationSession_RelCost_sql, "")
	   recItem.SupportUplift  = get_String(rec, dm.EstimationSession_SupportUplift_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.EstimationSession_SYSCreated_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.EstimationSession_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.EstimationSession_SYSCreatedHost_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.EstimationSession_SYSUpdated_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.EstimationSession_SYSUpdatedBy_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.EstimationSession_SYSUpdatedHost_sql, "")
	   recItem.SYSDeleted  = get_String(rec, dm.EstimationSession_SYSDeleted_sql, "")
	   recItem.SYSDeletedBy  = get_String(rec, dm.EstimationSession_SYSDeletedBy_sql, "")
	   recItem.SYSDeletedHost  = get_String(rec, dm.EstimationSession_SYSDeletedHost_sql, "")
	   recItem.Name  = get_String(rec, dm.EstimationSession_Name_sql, "")
	   recItem.AdoID  = get_String(rec, dm.EstimationSession_AdoID_sql, "")
	   recItem.FreshdeskID  = get_String(rec, dm.EstimationSession_FreshdeskID_sql, "")
	   recItem.TrackerID  = get_String(rec, dm.EstimationSession_TrackerID_sql, "")
	   recItem.EstRef  = get_String(rec, dm.EstimationSession_EstRef_sql, "")
	   recItem.ExtRef  = get_String(rec, dm.EstimationSession_ExtRef_sql, "")
	   recItem.SYSActivity  = get_String(rec, dm.EstimationSession_SYSActivity_sql, "")
	   recItem.SYSDbVersion  = get_String(rec, dm.EstimationSession_SYSDbVersion_sql, "")
	   recItem.Comments  = get_String(rec, dm.EstimationSession_Comments_sql, "")
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// If there are fields below, create the methods in adaptor\EstimationSession_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	   recItem.Origin  = adaptor.EstimationSession_Origin_OnFetch_impl (recItem)
	   recItem.OriginStateID  = adaptor.EstimationSession_OriginStateID_OnFetch_impl (recItem)
	   recItem.OriginState  = adaptor.EstimationSession_OriginState_OnFetch_impl (recItem)
	   recItem.OriginDocTypeID  = adaptor.EstimationSession_OriginDocTypeID_OnFetch_impl (recItem)
	   recItem.OriginDocType  = adaptor.EstimationSession_OriginDocType_OnFetch_impl (recItem)
	   recItem.OriginCode  = adaptor.EstimationSession_OriginCode_OnFetch_impl (recItem)
	   recItem.OriginName  = adaptor.EstimationSession_OriginName_OnFetch_impl (recItem)
	   recItem.OriginRate  = adaptor.EstimationSession_OriginRate_OnFetch_impl (recItem)
	   recItem.ProjectProfileID  = adaptor.EstimationSession_ProjectProfileID_OnFetch_impl (recItem)
	   recItem.ProjectProfile  = adaptor.EstimationSession_ProjectProfile_OnFetch_impl (recItem)
	   recItem.ProjectDefaultReleases  = adaptor.EstimationSession_ProjectDefaultReleases_OnFetch_impl (recItem)
	   recItem.ProjectDefaultReleaseHours  = adaptor.EstimationSession_ProjectDefaultReleaseHours_OnFetch_impl (recItem)
	   recItem.ProjectBlendedRate  = adaptor.EstimationSession_ProjectBlendedRate_OnFetch_impl (recItem)
	   recItem.ProjectStateID  = adaptor.EstimationSession_ProjectStateID_OnFetch_impl (recItem)
	   recItem.ProjectState  = adaptor.EstimationSession_ProjectState_OnFetch_impl (recItem)
	   recItem.ProjectName  = adaptor.EstimationSession_ProjectName_OnFetch_impl (recItem)
	   recItem.ProjectStartDate  = adaptor.EstimationSession_ProjectStartDate_OnFetch_impl (recItem)
	   recItem.ProjectEndDate  = adaptor.EstimationSession_ProjectEndDate_OnFetch_impl (recItem)
	   recItem.ProfileSupportUpliftPerc  = adaptor.EstimationSession_ProfileSupportUpliftPerc_OnFetch_impl (recItem)
	   recItem.CCY  = adaptor.EstimationSession_CCY_OnFetch_impl (recItem)
	   recItem.CCYCode  = adaptor.EstimationSession_CCYCode_OnFetch_impl (recItem)
	   recItem.EffortTotal  = adaptor.EstimationSession_EffortTotal_OnFetch_impl (recItem)
	   recItem.FreshDeskURI  = adaptor.EstimationSession_FreshDeskURI_OnFetch_impl (recItem)
	   recItem.ADOURI  = adaptor.EstimationSession_ADOURI_OnFetch_impl (recItem)
	   recItem.NoActiveFeatures  = adaptor.EstimationSession_NoActiveFeatures_OnFetch_impl (recItem)
	
	// 
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func EstimationSession_NewID(r dm.EstimationSession) string {
	
			id := uuid.New().String()
	
	return id
}



// estimationsession_Fetch read all EstimationSession's
func EstimationSession_New() (int, []dm.EstimationSession, dm.EstimationSession, error) {

	var r = dm.EstimationSession{}
	var rList []dm.EstimationSession
	

	// START
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	r.Origin,r.Origin_props = adaptor.EstimationSession_Origin_impl (adaptor.NEW,r.EstimationSessionID,r.Origin,r,r.Origin_props)
	r.OriginStateID,r.OriginStateID_props = adaptor.EstimationSession_OriginStateID_impl (adaptor.NEW,r.EstimationSessionID,r.OriginStateID,r,r.OriginStateID_props)
	r.OriginState,r.OriginState_props = adaptor.EstimationSession_OriginState_impl (adaptor.NEW,r.EstimationSessionID,r.OriginState,r,r.OriginState_props)
	r.OriginDocTypeID,r.OriginDocTypeID_props = adaptor.EstimationSession_OriginDocTypeID_impl (adaptor.NEW,r.EstimationSessionID,r.OriginDocTypeID,r,r.OriginDocTypeID_props)
	r.OriginDocType,r.OriginDocType_props = adaptor.EstimationSession_OriginDocType_impl (adaptor.NEW,r.EstimationSessionID,r.OriginDocType,r,r.OriginDocType_props)
	r.OriginCode,r.OriginCode_props = adaptor.EstimationSession_OriginCode_impl (adaptor.NEW,r.EstimationSessionID,r.OriginCode,r,r.OriginCode_props)
	r.OriginName,r.OriginName_props = adaptor.EstimationSession_OriginName_impl (adaptor.NEW,r.EstimationSessionID,r.OriginName,r,r.OriginName_props)
	r.OriginRate,r.OriginRate_props = adaptor.EstimationSession_OriginRate_impl (adaptor.NEW,r.EstimationSessionID,r.OriginRate,r,r.OriginRate_props)
	r.ProjectProfileID,r.ProjectProfileID_props = adaptor.EstimationSession_ProjectProfileID_impl (adaptor.NEW,r.EstimationSessionID,r.ProjectProfileID,r,r.ProjectProfileID_props)
	r.ProjectProfile,r.ProjectProfile_props = adaptor.EstimationSession_ProjectProfile_impl (adaptor.NEW,r.EstimationSessionID,r.ProjectProfile,r,r.ProjectProfile_props)
	r.ProjectDefaultReleases,r.ProjectDefaultReleases_props = adaptor.EstimationSession_ProjectDefaultReleases_impl (adaptor.NEW,r.EstimationSessionID,r.ProjectDefaultReleases,r,r.ProjectDefaultReleases_props)
	r.ProjectDefaultReleaseHours,r.ProjectDefaultReleaseHours_props = adaptor.EstimationSession_ProjectDefaultReleaseHours_impl (adaptor.NEW,r.EstimationSessionID,r.ProjectDefaultReleaseHours,r,r.ProjectDefaultReleaseHours_props)
	r.ProjectBlendedRate,r.ProjectBlendedRate_props = adaptor.EstimationSession_ProjectBlendedRate_impl (adaptor.NEW,r.EstimationSessionID,r.ProjectBlendedRate,r,r.ProjectBlendedRate_props)
	r.ProjectStateID,r.ProjectStateID_props = adaptor.EstimationSession_ProjectStateID_impl (adaptor.NEW,r.EstimationSessionID,r.ProjectStateID,r,r.ProjectStateID_props)
	r.ProjectState,r.ProjectState_props = adaptor.EstimationSession_ProjectState_impl (adaptor.NEW,r.EstimationSessionID,r.ProjectState,r,r.ProjectState_props)
	r.ProjectName,r.ProjectName_props = adaptor.EstimationSession_ProjectName_impl (adaptor.NEW,r.EstimationSessionID,r.ProjectName,r,r.ProjectName_props)
	r.ProjectStartDate,r.ProjectStartDate_props = adaptor.EstimationSession_ProjectStartDate_impl (adaptor.NEW,r.EstimationSessionID,r.ProjectStartDate,r,r.ProjectStartDate_props)
	r.ProjectEndDate,r.ProjectEndDate_props = adaptor.EstimationSession_ProjectEndDate_impl (adaptor.NEW,r.EstimationSessionID,r.ProjectEndDate,r,r.ProjectEndDate_props)
	r.ProfileSupportUpliftPerc,r.ProfileSupportUpliftPerc_props = adaptor.EstimationSession_ProfileSupportUpliftPerc_impl (adaptor.NEW,r.EstimationSessionID,r.ProfileSupportUpliftPerc,r,r.ProfileSupportUpliftPerc_props)
	r.CCY,r.CCY_props = adaptor.EstimationSession_CCY_impl (adaptor.NEW,r.EstimationSessionID,r.CCY,r,r.CCY_props)
	r.CCYCode,r.CCYCode_props = adaptor.EstimationSession_CCYCode_impl (adaptor.NEW,r.EstimationSessionID,r.CCYCode,r,r.CCYCode_props)
	r.EffortTotal,r.EffortTotal_props = adaptor.EstimationSession_EffortTotal_impl (adaptor.NEW,r.EstimationSessionID,r.EffortTotal,r,r.EffortTotal_props)
	r.FreshDeskURI,r.FreshDeskURI_props = adaptor.EstimationSession_FreshDeskURI_impl (adaptor.NEW,r.EstimationSessionID,r.FreshDeskURI,r,r.FreshDeskURI_props)
	r.ADOURI,r.ADOURI_props = adaptor.EstimationSession_ADOURI_impl (adaptor.NEW,r.EstimationSessionID,r.ADOURI,r,r.ADOURI_props)
	r.NoActiveFeatures,r.NoActiveFeatures_props = adaptor.EstimationSession_NoActiveFeatures_impl (adaptor.NEW,r.EstimationSessionID,r.NoActiveFeatures,r,r.NoActiveFeatures_props)
	// 
	// Dynamically generated 11/12/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}