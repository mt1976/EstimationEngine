package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/data.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Data (data)
// Endpoint 	        : Data (DataID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 04/01/2023 at 13:23:07
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"strconv"
	"strings"

	core "github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/logs"

	dm "github.com/mt1976/ebEstimates/datamodel"
)

// Data_GetByID() returns a single Data record
func Data_Get(class string, field string) (string, error) {

	id := data_BuildID(class, field)

	tsql := Data_SQLbase
	tsql = tsql + " " + core.DB_WHERE + " " + dm.Data_SQLSearchID + core.DB_EQ + "'" + id + "'"
	logs.Information("tsql", tsql)
	_, _, dataItem, err := data_Fetch(tsql)
	logs.Information("dataItem", dataItem.Value)
	return dataItem.Value, err
}

// Data_GetString() returns a single Data record
func Data_GetString(class string, field string) (string, error) {
	value, err := Data_Get(class, field)
	return value, err
}

// Data_GetByID() returns a single Data record
func Data_GetInt(class string, field string) (int, error) {

	value, err := Data_Get(class, field)
	if err != nil {
		return 0, err
	}
	rtnVal, err2 := strconv.Atoi(value)
	if err2 != nil {
		return 0, err2
	}

	return rtnVal, nil
}

func data_BuildID(class string, field string) string {
	fieldval := class + "-" + field
	fieldval = core.EncodeString(fieldval)
	return fieldval
}

func Data_GetArray(class string, field string) ([]string, error) {
	value, err := Data_Get(class, field)
	if err != nil {
		return strings.Split(value, ","), nil
	}
	return nil, err
}
