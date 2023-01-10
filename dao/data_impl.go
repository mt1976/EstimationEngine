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

	dm "github.com/mt1976/ebEstimates/datamodel"
)

// Data_GetByID() returns a single Data record
func Data_Get(class string, field string) (string, error) {

	id := data_BuildID(class, field)

	tsql := Data_SQLbase
	tsql = tsql + " " + core.DB_WHERE + " " + dm.Data_SQLSearchID + core.DB_EQ + "'" + id + "'"
	//logs.Information("tsql", tsql)
	_, _, dataItem, err := data_Fetch(tsql)
	//logs.Information("dataItem", dataItem.Value)
	return dataItem.Value, err
}

// Data_Put() returns a single Data record
func Data_Put(class string, field string, value string) (string, error) {
	//logs.Information("Data_Put", class+"-"+field+"-"+value)
	id := data_BuildID(class, field)
	//logs.Information("Data_Put ID", id)
	dataItem := dm.Data{}
	dataItem.DataID = id
	dataItem.Class = class
	dataItem.Field = field
	dataItem.Value = value
	err2 := Data_StoreSystem(dataItem)
	if err2 != nil {
		return "error", err2
	}
	return "ok", err2
}

func Data_GetSEQ(docType string) (int, string, error) {
	if docType == "" {
		return 0, "", nil
	}
	if docType == "SOW" {
		docType = "RSC"
	}
	rtnVal, err := Data_GetInt("SEQ", docType)
	if err != nil {
		return 0, "", err
	}
	rtnVal++
	return rtnVal, strconv.Itoa(rtnVal), nil
}

func Data_PutSEQ(docType string, seq int) error {
	if docType == "SOW" {
		docType = "RSC"
	}
	_, err := Data_Put("SEQ", docType, strconv.Itoa(seq))
	if err != nil {
		return err
	}
	return err
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
