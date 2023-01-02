package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/profile.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Profile (profile)
// Endpoint 	        : Profile (ProfileID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 02/01/2023 at 15:17:14
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
)

var Profile_SQLbase string
var Profile_QualifiedName string
func init(){
	Profile_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.Profile_SQLTable)
	Profile_SQLbase =  core.DB_SELECT + " "+ core.DB_ALL + " " + core.DB_FROM + " " + Profile_QualifiedName
}

// Profile_GetList() returns a list of all Profile records
func Profile_GetList() (int, []dm.Profile, error) {
	
	tsql := Profile_SQLbase
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


// Profile_GetByID() returns a single Profile record
func Profile_GetByID(id string) (int, dm.Profile, error) {


	tsql := Profile_SQLbase
	tsql = tsql + " " + core.DB_WHERE + " " + dm.Profile_SQLSearchID + core.DB_EQ + "'" + id + "'"
	_, _, profileItem, _ := profile_Fetch(tsql)

	// START
	// Dynamically generated 02/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 02/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, profileItem, nil
}



// Profile_DeleteByID() deletes a single Profile record
func Profile_Delete(id string) {




// Uses Hard Delete
	object_Table := Profile_QualifiedName
	tsql := core.DB_DELETE+" "+core.DB_FROM+" " + object_Table
	tsql = tsql + " " + core.DB_WHERE + " " + dm.Profile_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
	
	

	
}














	
	// Profile_SoftDeleteByID() soft deletes a single Profile record
func Profile_SoftDelete(id string) {
	//Uses Soft Delete
		_,profileItem,_ := Profile_GetByID(id)
		profileItem.SYSDeletedBy = Audit_Update("", Audit_Host())
		profileItem.SYSDeleted = Audit_Update("", Audit_TimeStamp())
		profileItem.SYSDeletedHost = Audit_Update("", Audit_Host())
		Profile_StoreSystem(profileItem)
}
	


// Profile_Store() saves/stores a Profile record to the database
func Profile_Store(r dm.Profile,req *http.Request) error {

	r, err := Profile_Validate(r)
	if err == nil {
		err = profile_Save(r, Audit_User(req))
	} else {
		logs.Information("Profile_Store()", err.Error())
	}

	return err
}

// Profile_StoreSystem() saves/stores a Profile record to the database
func Profile_StoreSystem(r dm.Profile) error {
	
	r, err := Profile_Validate(r)
	if err == nil {
		err = profile_Save(r, Audit_Host())
	} else {
		logs.Information("Profile_Store()", err.Error())
	}

	return err
}

// Profile_Validate() validates for saves/stores a Profile record to the database
func Profile_Validate(r dm.Profile) (dm.Profile, error) {
	var err error
	// START
	// Dynamically generated 02/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 02/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	

	return r,err
}
//

// profile_Save() saves/stores a Profile record to the database
func profile_Save(r dm.Profile,usr string) error {

    var err error



	

	if len(r.ProfileID) == 0 {
		r.ProfileID = Profile_NewID(r)
	}

// If there are fields below, create the methods in dao\profile_impl.go



































	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	r.SYSDbVersion = core.DB_Version()
	
logs.Storing("Profile",fmt.Sprintf("%v", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 02/01/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	// Dynamically generated 02/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := core.DB_INSERT + " " + core.DB_INTO + " " + Profile_QualifiedName
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " "+core.DB_VALUES +" (" + values(ts) + ")"

	Profile_Delete(r.ProfileID)
	das.Execute(tsql)



	return err

}



// profile_Fetch read all Profile's
func profile_Fetch(tsql string) (int, []dm.Profile, dm.Profile, error) {

	var recItem dm.Profile
	var recList []dm.Profile

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 02/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.Profile_SYSId_sql, "0")
	   recItem.ProfileID  = get_String(rec, dm.Profile_ProfileID_sql, "")
	   recItem.Code  = get_String(rec, dm.Profile_Code_sql, "")
	   recItem.Name  = get_String(rec, dm.Profile_Name_sql, "")
	   recItem.StartDate  = get_String(rec, dm.Profile_StartDate_sql, "")
	   recItem.EndDate  = get_String(rec, dm.Profile_EndDate_sql, "")
	   recItem.DefaultReleases  = get_String(rec, dm.Profile_DefaultReleases_sql, "")
	   recItem.DefaultReleaseHours  = get_String(rec, dm.Profile_DefaultReleaseHours_sql, "")
	   recItem.BlendedRate  = get_String(rec, dm.Profile_BlendedRate_sql, "")
	   recItem.Rounding  = get_String(rec, dm.Profile_Rounding_sql, "")
	   recItem.HoursPerDay  = get_String(rec, dm.Profile_HoursPerDay_sql, "")
	   recItem.REQPerc  = get_String(rec, dm.Profile_REQPerc_sql, "")
	   recItem.ANAPerc  = get_String(rec, dm.Profile_ANAPerc_sql, "")
	   recItem.DOCPerc  = get_String(rec, dm.Profile_DOCPerc_sql, "")
	   recItem.PMPerc  = get_String(rec, dm.Profile_PMPerc_sql, "")
	   recItem.UATPerc  = get_String(rec, dm.Profile_UATPerc_sql, "")
	   recItem.GTMPerc  = get_String(rec, dm.Profile_GTMPerc_sql, "")
	   recItem.SupportUplift  = get_String(rec, dm.Profile_SupportUplift_sql, "")
	   recItem.ContigencyPerc  = get_String(rec, dm.Profile_ContigencyPerc_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.Profile_SYSCreated_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.Profile_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.Profile_SYSCreatedHost_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.Profile_SYSUpdated_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.Profile_SYSUpdatedBy_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.Profile_SYSUpdatedHost_sql, "")
	   recItem.SYSDeleted  = get_String(rec, dm.Profile_SYSDeleted_sql, "")
	   recItem.SYSDeletedBy  = get_String(rec, dm.Profile_SYSDeletedBy_sql, "")
	   recItem.SYSDeletedHost  = get_String(rec, dm.Profile_SYSDeletedHost_sql, "")
	   recItem.SYSActivity  = get_String(rec, dm.Profile_SYSActivity_sql, "")
	   recItem.Notes  = get_String(rec, dm.Profile_Notes_sql, "")
	   recItem.SYSDbVersion  = get_String(rec, dm.Profile_SYSDbVersion_sql, "")
	   recItem.Comments  = get_String(rec, dm.Profile_Comments_sql, "")
	   recItem.TrainingPerc  = get_String(rec, dm.Profile_TrainingPerc_sql, "")
	
	// If there are fields below, create the methods in adaptor\Profile_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 02/01/2023 by matttownsend (Matt Townsend) on silicon.local 
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
	// Dynamically generated 02/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 02/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}