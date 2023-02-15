package upgrade

import (
	"fmt"
	"strings"

	"github.com/mt1976/ebEstimates/dao"
	"github.com/mt1976/ebEstimates/datamodel"
	"github.com/mt1976/ebEstimates/logs"
)

func upgrade_2(dbVer_str string) (string, error) {
	// Add your upgrade code here

	msg, e := upgrade_2_data(dbVer_str)
	return "Upgrade 2 " + dbVer_str + msg, e

}

func upgrade_2_data(dbVer_str string) (string, error) {
	// Add your upgrade code here

	//get all data
	_, processList, e := dao.Data_GetList()
	if e != nil {
		return "", e
	}
	//for each data
	for _, process := range processList {
		if process.SYSDbVersion != dbVer_str {

			msg, e := upgrade_2_data_item(dbVer_str, process)
			if e != nil {
				return "", e
			}
			logs.Information("Upgrade 2", msg)
			//filter := ""
		}
	}

	return "Upgrade 2 - Complete", nil
}

func upgrade_2_data_item(dbVer_str string, process datamodel.Data) (string, error) {
	// Add your upgrade code here
	b4 := process
	process.Class = strings.ReplaceAll(process.Class, "_", "-")
	items := strings.Split(process.Class, "-")
	newClass := ""
	newCategory := ""
	if len(items) > 1 {
		newClass = items[0]
		newCategory = items[1]
	} else {
		newClass = items[0]
		newCategory = ""
	}
	if newCategory == "" && process.Class == "SEQ" {
		newClass = "System"
		newCategory = "SEQ"
	}
	if newCategory == "" {
		newCategory = newClass
		newClass = "System"
	}

	process.Class = newClass
	process.Category = newCategory
	process.SYSDbVersion = dbVer_str
	process.DataID = dao.Data_NewID(process)
	//dumpProcess(b4, process)
	logs.Information("Upgrade 2", fmt.Sprintf("OldClass: %s NewClass: %s NewCategory: %s", process.Class, newClass, newCategory))
	dao.Data_Delete(b4.DataID)
	_, err := dao.Data_StoreSystem(process)
	if err != nil {
		return "", err
	}

	return "Upgrade 2 - Complete", nil
}

func dumpProcess(before datamodel.Data, process datamodel.Data) {
	fmt.Printf("before: %v\n", before)
	fmt.Printf("process: %v\n", process)
}
