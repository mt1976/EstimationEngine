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
	"errors"
	"strconv"
	"strings"

	core "github.com/mt1976/ebEstimates/core"
	das "github.com/mt1976/ebEstimates/das"
	"github.com/mt1976/ebEstimates/logs"
	"golang.org/x/exp/slices"

	dm "github.com/mt1976/ebEstimates/datamodel"
)

// Data_Get() returns a single Data record
func Data_Get(class string, field string, category string) (string, error) {

	if class == "" || field == "" || category == "" {
		return "", errors.New("invalid parameters passed to Data_Get" + class + "-" + field + "-" + category)
	}
	// If the category passed is not in the valid list then return error
	if !slices.Contains(dm.Data_Category_List, category) {
		msg := "data_Get - invalid category passed to Data_Get - " + category
		error := errors.New(msg)
		logs.Fatal(msg, error)
		return "", error
	}

	logs.Query("Data_Get " + class + "-" + field + "-" + category)
	id := data_BuildID(class, field, category)

	tsql := Data_SQLbase
	tsql = tsql + " " + das.WHERE + dm.Data_SQLSearchID + das.EQ + das.ID(id)
	//logs.Information("tsql", tsql)
	_, _, dataItem, err := data_Fetch(tsql)
	//logs.Information("dataItem", dataItem.Value)

	if err != nil {
		return dataItem.Value, err
	}
	if dataItem.DataID == "" {
		logs.Warning("Data_Get - No Content found for Data " + core.DQuote(id))
		_, err := Data_Put(class, field, category, "")

		return "", err
	}

	return dataItem.Value, err
}

// Data_Put() returns a single Data record
func Data_Put(class string, field string, category string, value string) (string, error) {
	logs.Information("Data_Put", class+"-"+category+"-"+field+"-"+value)
	if class == "" || field == "" || category == "" {
		return "", errors.New("invalid parameters passed to Data_Put" + class + "-" + field + "-" + category)
	}
	// If the category passed is not in the valid list then return error
	if !slices.Contains(dm.Data_Category_List, category) {
		return "", errors.New("invalid category passed to Data_Put - " + category)
	}

	logs.Information("Data_Put", class+"-"+category+"-"+field)
	id := data_BuildID(class, field, category)
	//logs.Information("Data_Put ID", id)
	dataItem := dm.Data{}
	dataItem.DataID = id
	dataItem.Class = class
	dataItem.Field = field
	dataItem.Value = value
	dataItem.Category = category
	_, err2 := Data_StoreSystem(dataItem)
	if err2 != nil {
		return "error", err2
	}
	return "ok", err2
}

func data_GetSEQ(docType string) (int, string, error) {
	if docType == "" {
		return 0, "", nil
	}

	rtnVal, err := Data_GetInt("System", docType, dm.Data_Category_Sequence)
	if err != nil {
		return 0, "", err
	}
	rtnVal++
	return rtnVal, strconv.Itoa(rtnVal), nil
}

func data_PutSEQ(docType string, seq int) error {
	if docType == "" || seq == 0 {
		return errors.New("invalid parameters passed to data_PutSEQ")
	}

	_, err := Data_Put("System", docType, dm.Data_Category_Sequence, strconv.Itoa(seq))
	if err != nil {
		return err
	}
	return err
}

func Data_NextSEQ(docType string) (int, string, error) {
	if docType == "" {
		return 0, "", nil
	}
	if docType == "SOW" {
		docType = "RSC"
	}
	seq, seqStr, err := data_GetSEQ(docType)
	if err != nil {
		return 0, "", err
	}
	err = data_PutSEQ(docType, seq)
	if err != nil {
		return 0, "", err
	}
	return seq, docType + seqStr, nil
}

// Data_GetString() returns a single Data record
func Data_GetString(class string, field string, category string) (string, error) {
	value, err := Data_Get(class, field, category)
	return value, err
}

// Data_GetByID() returns a single Data record
func Data_GetInt(class string, field string, category string) (int, error) {

	value, err := Data_Get(class, field, category)
	if err != nil {
		return 0, err
	}
	rtnVal, err2 := strconv.Atoi(value)
	if err2 != nil {
		return 0, nil
	}

	return rtnVal, nil
}

func Data_GetFloat(class string, field string, category string) (float64, error) {
	value, err := Data_Get(class, field, category)
	if err != nil {
		return 0, err
	}
	if value == "" {
		return 0, nil
	}
	rtnVal, err2 := core.Numeric(value)
	if err2 != nil {
		return 0, nil
	}

	return rtnVal, nil
}

func data_BuildID(class string, field string, category string) string {
	fieldval := class + "-" + field + "-" + category
	fieldval = core.EncodeString(fieldval)
	return fieldval
}

func Data_GetArray(class string, field string, category string) ([]string, error) {

	value, err := Data_Get(class, field, category)
	//logs.Information("Data_GetArray", value)
	//logs.Information("Data_GetArray", err.Error())
	if err != nil {
		return nil, nil
	}
	return strings.Split(value, ","), err
}

// Data_GetMigrateable returns a list of Data records that can be migrated
func Data_GetMigrateable() (int, []dm.Data, error) {

	tsql := Data_SQLbase
	tsql = tsql + " " + das.WHERE + dm.Data_Migrate_sql + das.EQ + das.VALUE(das.TRUE)
	logs.Query(tsql)
	noItems, dataItems, _, err := data_Fetch(tsql)

	return noItems, dataItems, err
}
