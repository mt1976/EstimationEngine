package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/schedule.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Schedule (schedule)
// Endpoint 	        : Schedule (Schedule)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 24/01/2023 at 13:18:11
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

var Schedule_SQLbase string
var Schedule_QualifiedName string
func init(){
	Schedule_QualifiedName = get_TableName(core.ApplicationSQLSchema(), dm.Schedule_SQLTable)
	Schedule_SQLbase =  core.DB_SELECT + " "+ core.DB_ALL + " " + core.DB_FROM + " " + Schedule_QualifiedName
}

// Schedule_GetList() returns a list of all Schedule records
func Schedule_GetList() (int, []dm.Schedule, error) {
	
	count, scheduleList, err := Schedule_GetListFiltered("")
	
	return count, scheduleList, err
}

// Schedule_GetListFiltered() returns a filtered list of all Schedule records
func Schedule_GetListFiltered(filter string) (int, []dm.Schedule, error) {
	
	tsql := Schedule_SQLbase
	if filter != "" {
		tsql = tsql + " " + core.DB_WHERE + " " + filter
	}
	count, scheduleList, _, _ := schedule_Fetch(tsql)
	
	return count, scheduleList, nil
}



// Schedule_GetByID() returns a single Schedule record
func Schedule_GetByID(id string) (int, dm.Schedule, error) {


	tsql := Schedule_SQLbase
	tsql = tsql + " " + core.DB_WHERE + " " + dm.Schedule_SQLSearchID + core.DB_EQ + "'" + id + "'"
	_, _, scheduleItem, _ := schedule_Fetch(tsql)

	// START
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, scheduleItem, nil
}



// Schedule_DeleteByID() deletes a single Schedule record
func Schedule_Delete(id string) {




// Uses Hard Delete
	object_Table := Schedule_QualifiedName
	tsql := core.DB_DELETE+" "+core.DB_FROM+" " + object_Table
	tsql = tsql + " " + core.DB_WHERE + " " + dm.Schedule_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
	
	

	
}














	


// Schedule_Store() saves/stores a Schedule record to the database
func Schedule_Store(r dm.Schedule,req *http.Request) error {

	r, err := Schedule_Validate(r)
	if err == nil {
		err = schedule_Save(r, Audit_User(req))
	} else {
		logs.Information("Schedule_Store()", err.Error())
	}

	return err
}

// Schedule_StoreSystem() saves/stores a Schedule record to the database
func Schedule_StoreSystem(r dm.Schedule) error {
	
	r, err := Schedule_Validate(r)
	if err == nil {
		err = schedule_Save(r, Audit_Host())
	} else {
		logs.Information("Schedule_Store()", err.Error())
	}

	return err
}

// Schedule_Validate() validates for saves/stores a Schedule record to the database
func Schedule_Validate(r dm.Schedule) (dm.Schedule, error) {
	var err error
	// START
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	

	return r,err
}
//

// schedule_Save() saves/stores a Schedule record to the database
func schedule_Save(r dm.Schedule,usr string) error {

    var err error



	

	if len(r.Id) == 0 {
		r.Id = Schedule_NewID(r)
	}

// If there are fields below, create the methods in dao\schedule_impl.go






















	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	r.SYSDbVersion = core.DB_Version()
	
logs.Storing("Schedule",fmt.Sprintf("%v", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.Schedule_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.Schedule_Id_sql, r.Id)
	ts = addData(ts, dm.Schedule_Name_sql, r.Name)
	ts = addData(ts, dm.Schedule_Description_sql, r.Description)
	ts = addData(ts, dm.Schedule_Schedule_sql, r.Schedule)
	ts = addData(ts, dm.Schedule_Started_sql, r.Started)
	ts = addData(ts, dm.Schedule_Lastrun_sql, r.Lastrun)
	ts = addData(ts, dm.Schedule_Message_sql, r.Message)
	ts = addData(ts, dm.Schedule_Type_sql, r.Type)
	ts = addData(ts, dm.Schedule_Human_sql, r.Human)
	ts = addData(ts, dm.Schedule_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.Schedule_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.Schedule_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.Schedule_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.Schedule_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.Schedule_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.Schedule_SYSDeleted_sql, r.SYSDeleted)
	ts = addData(ts, dm.Schedule_SYSDeletedBy_sql, r.SYSDeletedBy)
	ts = addData(ts, dm.Schedule_SYSDeletedHost_sql, r.SYSDeletedHost)
	ts = addData(ts, dm.Schedule_SYSDbVersion_sql, r.SYSDbVersion)
		
	// 
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := core.DB_INSERT + " " + core.DB_INTO + " " + Schedule_QualifiedName
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " "+core.DB_VALUES +" (" + values(ts) + ")"

	Schedule_Delete(r.Id)
	das.Execute(tsql)



	return err

}



// schedule_Fetch read all Schedule's
func schedule_Fetch(tsql string) (int, []dm.Schedule, dm.Schedule, error) {

	var recItem dm.Schedule
	var recList []dm.Schedule

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.Schedule_SYSId_sql, "0")
	   recItem.Id  = get_String(rec, dm.Schedule_Id_sql, "")
	   recItem.Name  = get_String(rec, dm.Schedule_Name_sql, "")
	   recItem.Description  = get_String(rec, dm.Schedule_Description_sql, "")
	   recItem.Schedule  = get_String(rec, dm.Schedule_Schedule_sql, "")
	   recItem.Started  = get_String(rec, dm.Schedule_Started_sql, "")
	   recItem.Lastrun  = get_String(rec, dm.Schedule_Lastrun_sql, "")
	   recItem.Message  = get_String(rec, dm.Schedule_Message_sql, "")
	   recItem.Type  = get_String(rec, dm.Schedule_Type_sql, "")
	   recItem.Human  = get_String(rec, dm.Schedule_Human_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.Schedule_SYSCreated_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.Schedule_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.Schedule_SYSCreatedHost_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.Schedule_SYSUpdated_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.Schedule_SYSUpdatedBy_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.Schedule_SYSUpdatedHost_sql, "")
	   recItem.SYSDeleted  = get_String(rec, dm.Schedule_SYSDeleted_sql, "")
	   recItem.SYSDeletedBy  = get_String(rec, dm.Schedule_SYSDeletedBy_sql, "")
	   recItem.SYSDeletedHost  = get_String(rec, dm.Schedule_SYSDeletedHost_sql, "")
	   recItem.SYSDbVersion  = get_String(rec, dm.Schedule_SYSDbVersion_sql, "")
	
	// If there are fields below, create the methods in adaptor\Schedule_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func Schedule_NewID(r dm.Schedule) string {
	
			id := uuid.New().String()
	
	return id
}



// schedule_Fetch read all Schedule's
func Schedule_New() (int, []dm.Schedule, dm.Schedule, error) {

	var r = dm.Schedule{}
	var rList []dm.Schedule
	

	// START
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 24/01/2023 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}