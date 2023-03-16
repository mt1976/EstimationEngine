package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/profile.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Profile (profile)
// Endpoint 	        : Profile (ProfileID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 15/03/2023 at 19:24:49
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	core "github.com/mt1976/ebEstimates/core"
	das "github.com/mt1976/ebEstimates/das"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

var Profile_SQLbase string
var Profile_QualifiedName string

func init() {
	Profile_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.Profile_SQLTable)
	Profile_SQLbase = das.SELECTALL + das.FROM + Profile_QualifiedName
}

// Profile_GetList() returns a list of all Profile records
func Profile_GetList() (int, []dm.Profile, error) {
	count, profileList, err := Profile_GetListFiltered("")
	return count, profileList, err
}

// Profile_GetListFiltered() returns a filtered list of all Profile records
func Profile_GetListFiltered(filter string) (int, []dm.Profile, error) {

	tsql := Profile_SQLbase
	if filter != "" {
		tsql = tsql + " " + das.WHERE + filter
	}
	count, profileList, _, _ := profile_Fetch(tsql)

	return count, profileList, nil
}

// Profile_GetLookup() returns a lookup list of all Profile items in lookup format
func Profile_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, profileList, _ := Profile_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: profileList[i].Code, Name: profileList[i].Name})
	}
	return returnList
}

// Profile_GetFilteredLookup() returns a lookup list of all Profile items in lookup format
func Profile_GetFilteredLookup(requestObject string, requestField string) []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	objectName := Translate("ObjectName", requestObject)
	reqField := requestField + "_Profile_Filter"

	usage := "Defines a filter for a lookup list of Profile records, when requested by " + requestField + "." + core.TEXTAREA_CR
	usage = usage + "Fields can be any of those in the underlying DB table." + core.TEXTAREA_CR
	usage = usage + "Examples Below:" + core.TEXTAREA_CR
	usage = usage + "* datalength(_deleted) = 0 or " + core.TEXTAREA_CR
	usage = usage + "* class IN ('x','y','z')"

	filter, _ := Data_GetString(objectName, reqField, dm.Data_Category_FilterRule, usage)
	if filter == "" {
		logs.Warning("Profile_GetFilteredLookup() - No filter found : " + reqField + " for Object: " + objectName)
	}
	count, profileList, _ := Profile_GetListFiltered(filter)
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: profileList[i].Code, Name: profileList[i].Name})
	}
	return returnList
}

// Profile_GetByID() returns a single Profile record
func Profile_GetByID(id string) (int, dm.Profile, error) {

	tsql := Profile_SQLbase
	tsql = tsql + " " + das.WHERE + dm.Profile_SQLSearchID + das.EQ + das.ID(id)
	_, _, profileItem, _ := profile_Fetch(tsql)

	profileItem = Profile_PostGet(profileItem, id)

	return 1, profileItem, nil
}

func Profile_PostGet(profileItem dm.Profile, id string) dm.Profile {

	return profileItem
}

// Profile_DeleteByID() deletes a single Profile record
func Profile_Delete(id string) {
	Profile_SoftDelete(id)
}

// Profile_SoftDelete(id string) soft deletes a single Profile record
func Profile_SoftDelete(id string) {
	//Uses Soft Delete
	_, profileItem, _ := Profile_GetByID(id)
	profileItem.SYSDeletedBy = Audit_Update("", Audit_Host())
	profileItem.SYSDeleted = Audit_Update("", Audit_TimeStamp())
	profileItem.SYSDeletedHost = Audit_Update("", Audit_Host())
	_, err := Profile_StoreSystem(profileItem)
	if err != nil {
		logs.Error("Profile_SoftDelete()", err)
	}
}

// Profile_HardDelete(id string) soft deletes a single Profile record
func Profile_HardDelete(id string) {
	// Uses Hard Delete
	object_Table := Profile_QualifiedName
	tsql := das.DELETE + das.FROM + object_Table
	tsql = tsql + " " + das.WHERE + dm.Profile_SQLSearchID + das.EQ + das.ID(id)
	das.Execute(tsql)
	//if err != nil {
	//	logs.Error("Profile_SoftDelete()",err)
	//}
}

// Profile_Store() saves/stores a Profile record to the database
func Profile_Store(r dm.Profile, req *http.Request) (dm.Profile, error) {

	r, err := Profile_Validate(r)
	if err == nil {
		err = profile_Save(r, Audit_User(req))
	} else {
		logs.Information("Profile_Store()", err.Error())
	}

	return r, err
}

// Profile_StoreSystem() saves/stores a Profile record to the database
func Profile_StoreSystem(r dm.Profile) (dm.Profile, error) {

	r, err := Profile_Validate(r)
	if err == nil {
		err = profile_Save(r, Audit_Host())
	} else {
		logs.Information("Profile_Store()", err.Error())
	}

	return r, err
}

// Profile_StoreProcess() saves/stores a Profile record to the database
func Profile_StoreProcess(r dm.Profile, operator string) (dm.Profile, error) {

	r, err := Profile_Validate(r)
	if err == nil {
		err = profile_Save(r, operator)
	} else {
		logs.Information("Profile_StoreProcess()", err.Error())
	}

	return r, err
}

// Profile_Validate() validates for saves/stores a Profile record to the database
func Profile_Validate(r dm.Profile) (dm.Profile, error) {
	var err error

	// Cross Validation
	var errVal error
	r, _, errVal = Profile_ObjectValidation_impl(PUT, r.ProfileID, r)
	if errVal != nil {
		err = errVal
	}
	return r, err
}

//

// profile_Save() saves/stores a Profile record to the database
func profile_Save(r dm.Profile, usr string) error {

	var err error

	if len(r.ProfileID) == 0 {
		r.ProfileID = Profile_NewID(r)
	}

	// If there are fields below, create the methods in dao\profile_impl.go

	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost, Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("", usr)
	r.SYSUpdatedHost = Audit_Update("", Audit_Host())
	r.SYSDbVersion = core.DB_Version()

	logs.Storing("Profile", fmt.Sprintf("%v", r))

	//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	ts = addData(ts, dm.Profile_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.Profile_ProfileID_sql, r.ProfileID)
	ts = addData(ts, dm.Profile_Code_sql, r.Code)
	ts = addData(ts, dm.Profile_Name_sql, r.Name)
	ts = addData(ts, dm.Profile_StartDate_sql, r.StartDate)
	ts = addData(ts, dm.Profile_EndDate_sql, r.EndDate)
	ts = addData(ts, dm.Profile_DefaultReleases_sql, r.DefaultReleases)
	ts = addData(ts, dm.Profile_DefaultReleaseHours_sql, r.DefaultReleaseHours)
	ts = addData(ts, dm.Profile_BlendedRate_sql, r.BlendedRate)
	ts = addData(ts, dm.Profile_Rounding_sql, r.Rounding)
	ts = addData(ts, dm.Profile_HoursPerDay_sql, r.HoursPerDay)
	ts = addData(ts, dm.Profile_REQPerc_sql, r.REQPerc)
	ts = addData(ts, dm.Profile_ANAPerc_sql, r.ANAPerc)
	ts = addData(ts, dm.Profile_DOCPerc_sql, r.DOCPerc)
	ts = addData(ts, dm.Profile_PMPerc_sql, r.PMPerc)
	ts = addData(ts, dm.Profile_UATPerc_sql, r.UATPerc)
	ts = addData(ts, dm.Profile_GTMPerc_sql, r.GTMPerc)
	ts = addData(ts, dm.Profile_SupportUplift_sql, r.SupportUplift)
	ts = addData(ts, dm.Profile_ContigencyPerc_sql, r.ContigencyPerc)
	ts = addData(ts, dm.Profile_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.Profile_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.Profile_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.Profile_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.Profile_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.Profile_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.Profile_SYSDeleted_sql, r.SYSDeleted)
	ts = addData(ts, dm.Profile_SYSDeletedBy_sql, r.SYSDeletedBy)
	ts = addData(ts, dm.Profile_SYSDeletedHost_sql, r.SYSDeletedHost)
	ts = addData(ts, dm.Profile_SYSActivity_sql, r.SYSActivity)
	ts = addData(ts, dm.Profile_Notes_sql, r.Notes)
	ts = addData(ts, dm.Profile_SYSDbVersion_sql, r.SYSDbVersion)
	ts = addData(ts, dm.Profile_Comments_sql, r.Comments)
	ts = addData(ts, dm.Profile_TrainingPerc_sql, r.TrainingPerc)

	//
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END

	tsql := das.INSERT + das.INTO + Profile_QualifiedName
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " " + das.VALUES + "(" + values(ts) + ")"

	Profile_HardDelete(r.ProfileID)
	das.Execute(tsql)

	return err

}

// profile_Fetch read all Profile's
func profile_Fetch(tsql string) (int, []dm.Profile, dm.Profile, error) {

	var recItem dm.Profile
	var recList []dm.Profile

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(), err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
		// START
		// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local
		//
		recItem.SYSId = get_Int(rec, dm.Profile_SYSId_sql, "0")
		recItem.ProfileID = get_String(rec, dm.Profile_ProfileID_sql, "")
		recItem.Code = get_String(rec, dm.Profile_Code_sql, "")
		recItem.Name = get_String(rec, dm.Profile_Name_sql, "")
		recItem.StartDate = get_String(rec, dm.Profile_StartDate_sql, "")
		recItem.EndDate = get_String(rec, dm.Profile_EndDate_sql, "")
		recItem.DefaultReleases = get_String(rec, dm.Profile_DefaultReleases_sql, "")
		recItem.DefaultReleaseHours = get_String(rec, dm.Profile_DefaultReleaseHours_sql, "")
		recItem.BlendedRate = get_String(rec, dm.Profile_BlendedRate_sql, "")
		recItem.Rounding = get_String(rec, dm.Profile_Rounding_sql, "")
		recItem.HoursPerDay = get_String(rec, dm.Profile_HoursPerDay_sql, "")
		recItem.REQPerc = get_String(rec, dm.Profile_REQPerc_sql, "")
		recItem.ANAPerc = get_String(rec, dm.Profile_ANAPerc_sql, "")
		recItem.DOCPerc = get_String(rec, dm.Profile_DOCPerc_sql, "")
		recItem.PMPerc = get_String(rec, dm.Profile_PMPerc_sql, "")
		recItem.UATPerc = get_String(rec, dm.Profile_UATPerc_sql, "")
		recItem.GTMPerc = get_String(rec, dm.Profile_GTMPerc_sql, "")
		recItem.SupportUplift = get_String(rec, dm.Profile_SupportUplift_sql, "")
		recItem.ContigencyPerc = get_String(rec, dm.Profile_ContigencyPerc_sql, "")
		recItem.SYSCreated = get_String(rec, dm.Profile_SYSCreated_sql, "")
		recItem.SYSCreatedBy = get_String(rec, dm.Profile_SYSCreatedBy_sql, "")
		recItem.SYSCreatedHost = get_String(rec, dm.Profile_SYSCreatedHost_sql, "")
		recItem.SYSUpdated = get_String(rec, dm.Profile_SYSUpdated_sql, "")
		recItem.SYSUpdatedBy = get_String(rec, dm.Profile_SYSUpdatedBy_sql, "")
		recItem.SYSUpdatedHost = get_String(rec, dm.Profile_SYSUpdatedHost_sql, "")
		recItem.SYSDeleted = get_String(rec, dm.Profile_SYSDeleted_sql, "")
		recItem.SYSDeletedBy = get_String(rec, dm.Profile_SYSDeletedBy_sql, "")
		recItem.SYSDeletedHost = get_String(rec, dm.Profile_SYSDeletedHost_sql, "")
		recItem.SYSActivity = get_String(rec, dm.Profile_SYSActivity_sql, "")
		recItem.Notes = get_String(rec, dm.Profile_Notes_sql, "")
		recItem.SYSDbVersion = get_String(rec, dm.Profile_SYSDbVersion_sql, "")
		recItem.Comments = get_String(rec, dm.Profile_Comments_sql, "")
		recItem.TrainingPerc = get_String(rec, dm.Profile_TrainingPerc_sql, "")

		// If there are fields below, create the methods in dao\Profile_adaptor.go
		//
		// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local
		// END
		///
		//Add to the list
		//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}

func Profile_NewID(r dm.Profile) string {

	id := uuid.New().String()

	return id
}

// profile_Fetch read all Profile's
func Profile_New() (int, []dm.Profile, dm.Profile, error) {

	var r = dm.Profile{}
	var rList []dm.Profile

	// START
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//

	//
	// Dynamically generated 15/03/2023 by matttownsend (Matt Townsend) on silicon.local
	// END
	rList = append(rList, r)
	return 1, rList, r, nil
}
