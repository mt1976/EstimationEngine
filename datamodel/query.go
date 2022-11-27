package datamodel

// Defines the Fields to Fetch from SQL
// var appScheduleStoreSQL = "id, 	name, 	description, 	schedule, 	started, 	lastrun, 	message, 	_created, 	_who, 	_host, 	_updated, type"

// var sqlScheduleStoreId, sqlScheduleStoreName, sqlScheduleStoreDescription, sqlScheduleStoreSchedule, sqlScheduleStoreStarted, sqlScheduleStoreLastrun, sqlScheduleStoreMessage, sqlScheduleStoreSYSCreated, sqlScheduleStoreSYSWho, sqlScheduleStoreSYSHost, sqlScheduleStoreSYSUpdated, sqlScheduleStoreType sql.NullString

// var appScheduleStoreSQLINSERT = "INSERT INTO %s.scheduleStore(%s) VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');"
// var appScheduleStoreSQLDELETE = "DELETE FROM %s.scheduleStore WHERE id='%s';"
// var appScheduleStoreSQLSELECT = "SELECT %s FROM %s.scheduleStore;"
// var appScheduleStoreSQLGET = "SELECT %s FROM %s.scheduleStore WHERE id='%s';"

// DatabaseStoreMessages contains all the messages in DB source language used to interact with the DB
type DataStoreMessages struct {
	//Defines the Columns in the DB
	Table     string ""
	Columns   string ""
	Insert    string ""
	Delete    string ""
	DeleteAlt string ""
	Select    string ""
	Get       string ""
	GetAlt    string ""
}
