package dao

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	core "github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
)

func credentials_NewIDImpl(r dm.Credentials) string {
	return uuid.New().String()
}

// Credentials_GetByUserName() returns a single Credentials record
func Credentials_GetByUserName(id string) (int, dm.Credentials, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.Credentials_SQLTable)
	tsql = tsql + " WHERE " + dm.Credentials_Username_sql + "='" + id + "'"

	_, _, credentialsItem, _ := credentials_Fetch(tsql)
	return 1, credentialsItem, nil
}

// Credentials_GetByUUID() returns a single Credentials record
func Credentials_GetByUUID(id string) (int, dm.Credentials, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.Credentials_SQLTable)
	tsql = tsql + " WHERE " + dm.Credentials_Id_sql + "='" + id + "'"

	_, _, credentialsItem, _ := credentials_Fetch(tsql)
	return 1, credentialsItem, nil
}

// Credentials_GetByUUID() returns a single Credentials record
func Credentials_GetByEmail(id string) (int, dm.Credentials, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.Credentials_SQLTable)
	tsql = tsql + " WHERE " + dm.Credentials_Email_sql + "='" + id + "'"

	noItems, _, credentialsItem, _ := credentials_Fetch(tsql)
	return noItems, credentialsItem, nil
}

func getExpiryDate() string {
	expiryDate := time.Now()
	life, _ := strconv.Atoi(core.ApplicationCredentialsLife())
	expiryDate = expiryDate.AddDate(0, 0, life)
	return expiryDate.Format(core.DATETIMEFORMATUSER)
}
