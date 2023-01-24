package upgrade

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/dao"
	"github.com/mt1976/ebEstimates/logs"
)

var upSystem string
var upVersion string
var upPreviousVersion string
var upLastTime string
var upUpgrade string
var upDatabase string

func init() {
	upUpgrade = "Upgrade"
	logs.Information(upUpgrade, "INIT")
	upSystem = "System"
	upDatabase = "Database"
	upVersion = "Version"
	upPreviousVersion = "Previous" + upVersion
	upLastTime = "Last" + upUpgrade + "Time"

}

func Upgrade() (string, error) {
	// Get Application DB Release Level
	appRelease_str := core.DB_Version()
	// Get Current DB Release Level

	dbRelease_int, _ := getDBReleaseLevel(appRelease_str)
	appRelease_int, _ := strconv.ParseInt(appRelease_str, 10, 64)
	// If Versions are Equal do Nothing
	if dbRelease_int == appRelease_int {
		//logs.Information("Upgrade", "No Upgrade Required")
		return "No Upgrade Required", nil
	}
	// If DB Version is Greater than Application Version
	// This is a problem
	if dbRelease_int > appRelease_int {
		errMSG := fmt.Sprintf("Database Version is Greater than Application Version - DB: %s App: %s", dbVersionFormat(dbRelease_int), dbVersionFormat(appRelease_int))
		logs.Error("Upgrade", errors.New(errMSG))
		return errMSG, errors.New(errMSG)
	}
	// Add your upgrade code here
	// Upgrade from dbVersion to appRelease version
	logs.Information("Upgrade", "START")

	for version := dbRelease_int; version <= appRelease_int; version++ {
		msg := ""
		e := error(nil)
		dbVer_str := dbVersionFormat(version)
		switch version {
		case 1:
			// Add your upgrade code here
			msg, e = upgrade_1(dbVer_str)
		case 2:
			// Add your upgrade code here
			msg, e = upgrade_2(dbVer_str)
		default:
			// Nothing to do for this release
			msg = fmt.Sprintf("Upgrade Version %s not found", dbVer_str)
		}
		if e != nil {
			logs.Error("Upgrade", e)
			return msg, e
		}

		upgString := fmt.Sprintf("%s - %s", dbVer_str, msg)
		logs.Upgrade("Upgrade - " + upgString)
	}

	// Update Database Version
	shouldReturn, returnValue, returnValue1 := updateVersions(dbRelease_int, appRelease_str)
	if shouldReturn {
		return returnValue, returnValue1
	}

	logs.Information("Upgrade", "END")
	MSG := "Upgrade from '%s' to '%s'"
	MSG = fmt.Sprintf(MSG, dbVersionFormat(dbRelease_int), appRelease_str)
	logs.Upgrade(MSG)
	return MSG, nil
}

func updateVersions(dbRelease_int int64, appRelease_str string) (bool, string, error) {
	_, e := dao.Data_Put(upSystem, upPreviousVersion, upDatabase, dbVersionFormat(dbRelease_int))
	if e != nil {
		logs.Error("Upgrade", e)
		return true, "Upgrade - " + e.Error(), e
	}
	_, e = dao.Data_Put(upSystem, upVersion, upDatabase, appRelease_str)
	if e != nil {
		logs.Error("Upgrade", e)
		return true, "Upgrade - " + e.Error(), e
	}
	_, e = dao.Data_Put(upSystem, upLastTime, upDatabase, time.Now().Format(core.DATEMSG))
	if e != nil {
		logs.Error("Upgrade", e)
		return true, "Upgrade - " + e.Error(), e
	}
	return false, "", nil
}

func getDBReleaseLevel(appRel string) (int64, error) {
	dbRelease, _ := dao.Data_Get(upSystem, upVersion, upDatabase)
	if dbRelease == "" {
		logs.Warning("Upgrade - No Database Version Found")
		dbRelease = "0000"
		dao.Data_Put(upSystem, upVersion, upDatabase, dbRelease)
	}
	rtnValue, e := strconv.ParseInt(dbRelease, 10, 64)
	if e != nil {
		logs.Warning("Upgrade - " + e.Error())
		rtnValue = 0
		return 0, e
	}
	return rtnValue, nil
}

func dbVersionFormat(version int64) string {
	return fmt.Sprintf("%04d", version)
}
