package jobs

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
	core "github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	"github.com/mt1976/ebEstimates/logs"
)

func Data_Job_impl(j dm.JobDefinition) dm.JobDefinition {
	j.Name = "ConfigImportExport"
	j.Period = ""
	j.Description = "Import/Export Soft Configuration to/from CSV"
	j.Type = core.Adhoc
	j.ID = j.Name
	return j
}

func Data_Run_impl() (string, error) {

	message := ""
	jobName := Data_Job().Name
	inPath := ""
	outPath := ""
	ssMode := ""
	noWorkItems := 0
	action := ""

	//Get Mode from Data Settings and branch
	ssMode, err := dao.Data_Get(jobName, JOB_MODE, dm.Data_Category_State, "The mode of operation. Trial or Live. Done = Finished")
	if err != nil {
		return "No Mode Specified for " + jobName, err
	}

	//Get Action from Data
	action, err = dao.Data_Get(jobName, JOB_ACTION, dm.Data_Category_State, "The action to perform. Import or Export")
	if err != nil {
		return "No Action Specified for " + jobName, err
	}

	//Get inPath from Data
	msg := "The path to the input file. " + core.TEXTAREA_CR + "{{ENTITY}} is replaced with the entity name being processed."
	inPath, err = dao.Data_Get(jobName, JOB_IN_PATH, dm.Data_Category_Path, msg)
	if err != nil {
		return "No Input Path Specified for " + jobName, err
	}

	//Get outPath from Data
	outPath, err = dao.Data_Get(jobName, JOB_OUT_PATH, dm.Data_Category_Path, "The root path to the output file. Folders will be created as required.")
	if err != nil {
		return "No Output Path Specified for " + jobName, err
	}

	//Get What to do from Data
	msg = "The entity being proceseed" + core.TEXTAREA_CR + "DATA = Settings" + core.TEXTAREA_CR + "ORIGIN = Origin State" + core.TEXTAREA_CR + "DOCUMENT = Document Type" + core.TEXTAREA_CR + "PROJECT = Project State" + core.TEXTAREA_CR + "ESTIMATION = Estimation State" + core.TEXTAREA_CR + "TRANSLATION = Translations" + core.TEXTAREA_CR + "ALL = All of the above"
	what, err := dao.Data_Get(jobName, JOB_WHAT, dm.Data_Category_State, msg)
	what = strings.ToUpper(what)
	actionUpper := strings.ToUpper(action)
	switch actionUpper {
	case JOB_ACTION_EXPORT:

		var shouldReturn bool
		var exportMessage string
		var exportError error
		switch what {
		case "DATA":
			logs.Processing("Exporting Data")
			noWorkItems, message, shouldReturn, exportMessage, exportError = processDataExport(outPath, jobName, err, ssMode)
		case "ORIGIN":
			logs.Processing("Exporting Origin State")
			noWorkItems, message, shouldReturn, exportMessage, exportError = processOriginStateExport(outPath, jobName, err, ssMode)
		case "DOCUMENT":
			logs.Processing("Exporting Document Type")
			noWorkItems, message, shouldReturn, exportMessage, exportError = processDocumentTypeExport(outPath, jobName, err, ssMode)
		case "PROJECT":
			logs.Processing("Exporting Project State")
			noWorkItems, message, shouldReturn, exportMessage, exportError = processProjectStateExport(outPath, jobName, err, ssMode)
		case "ESTIMATION":
			logs.Processing("Exporting Estimation State")
			noWorkItems, message, shouldReturn, exportMessage, exportError = processEstimationStateExport(outPath, jobName, err, ssMode)
		case "TRANSLATION":
			logs.Processing("Exporting Translations")
			noWorkItems, message, shouldReturn, exportMessage, exportError = processTranslationExport(outPath, jobName, err, ssMode)
		case "ALL":
			logs.Processing("Exporting All Data")
			noWorkItems_0, message_0, _, _, _ := processDataExport(outPath, jobName, err, ssMode)
			noWorkItems_1, message_1, _, _, _ := processOriginStateExport(outPath, jobName, err, ssMode)
			noWorkItems_2, message_2, _, _, _ := processDocumentTypeExport(outPath, jobName, err, ssMode)
			noWorkItems_3, message_3, _, _, _ := processProjectStateExport(outPath, jobName, err, ssMode)
			noWorkItems_4, message_4, _, _, _ := processEstimationStateExport(outPath, jobName, err, ssMode)
			noWorkItems_5, message_5, _, _, _ := processTranslationExport(outPath, jobName, err, ssMode)
			logs.Processing("Exporting DONE")

			noWorkItems = noWorkItems_0 + noWorkItems_1 + noWorkItems_2 + noWorkItems_3 + noWorkItems_4 + noWorkItems_5
			message = message_0 + message_1 + message_2 + message_3 + message_4 + message_5

		default:
			return "Invalid What Specified for " + jobName + " " + core.DQuote(what), err
		}
		if shouldReturn {
			return exportMessage, exportError
		}
	case JOB_ACTION_IMPORT:
		//Get Data from Spreadsheet and Update Data
		if inPath == "" {
			return "No Input Path Specified for " + jobName, err
		}

		//Get Data from Spreadsheet and Update Data
		var importMessage string
		var importError error

		switch what {
		case "DATA":
			noWorkItems, message, importMessage, importError = processDataImport(inPath, jobName, err, ssMode)
		case "ORIGIN":
			noWorkItems, message, importMessage, importError = processOriginStateImport(inPath, jobName, err, ssMode)
		case "DOCUMENT":
			noWorkItems, message, importMessage, importError = processDocumentTypeImport(inPath, jobName, err, ssMode)
		case "PROJECT":
			noWorkItems, message, importMessage, importError = processProjectStateImport(inPath, jobName, err, ssMode)
		case "ESTIMATION":
			noWorkItems, message, importMessage, importError = processEstimationStateImport(inPath, jobName, err, ssMode)
		case "TRANSLATION":
			noWorkItems, message, importMessage, importError = processTranslationImport(inPath, jobName, err, ssMode)
		case "ALL":
			logs.Processing("Exporting All Data")
			noWorkItems_0, message_0, importMessage_0, _ := processDataImport(inPath, jobName, err, ssMode)
			noWorkItems_1, message_1, importMessage_1, _ := processOriginStateImport(inPath, jobName, err, ssMode)
			noWorkItems_2, message_2, importMessage_2, _ := processDocumentTypeImport(inPath, jobName, err, ssMode)
			noWorkItems_3, message_3, importMessage_3, _ := processProjectStateImport(inPath, jobName, err, ssMode)
			noWorkItems_4, message_4, importMessage_4, _ := processEstimationStateImport(inPath, jobName, err, ssMode)
			noWorkItems_5, message_5, importMessage_5, _ := processTranslationImport(inPath, jobName, err, ssMode)
			logs.Processing("Exporting DONE")

			noWorkItems = noWorkItems_0 + noWorkItems_1 + noWorkItems_2 + noWorkItems_3 + noWorkItems_4 + noWorkItems_5
			message = message_0 + message_1 + message_2 + message_3 + message_4 + message_5
			importMessage = importMessage_0 + importMessage_1 + importMessage_2 + importMessage_3 + importMessage_4 + importMessage_5

		default:
			return "Invalid What Specified for " + jobName + " " + core.DQuote(what), err
		}

		if importMessage != "" {
			return importMessage, importError
		}

	default:
		return "Invalid Mode Specified for " + jobName + " " + core.DQuote(ssMode), err
	}

	ret := RecordJobStatusInfo(message, noWorkItems, action, jobName, inPath, outPath, ssMode)
	if ret != "" {
		message = "Unable to Update Data for " + jobName + " " + ret
	}
	return message, nil
}

func processDataExport(outPath string, jobName string, err error, ssMode string) (int, string, bool, string, error) {
	var export []DataItemIO
	what := "Settings"

	now, fileName, resultsFile := getExportFileNames(what)

	if outPath == "" {
		return 0, "", true, "No Output Path Specified for " + jobName + " " + what, err
	}

	noitems, items, err := dao.Data_GetMigrateable()
	noWorkItems := noitems
	if err != nil || noWorkItems == 0 {
		return 0, "", true, "Unable to Get Data for " + jobName + " " + what, err
	}

	for _, item := range items {
		exportItem := DataItemIO{}
		exportItem.DataID = item.DataID
		exportItem.Category = item.Category
		exportItem.Class = item.Class
		exportItem.Field = item.Field
		exportItem.Value = item.Value
		exportItem.Migrate = item.Migrate
		if item.Value != "" {
			export = append(export, exportItem)
		}
	}

	exportFile, results := openExportFiles(outPath, fileName, resultsFile)

	err = gocsv.MarshalFile(&export, exportFile)
	if err != nil {
		panic(err)
	}
	//fmt.Println(csvContent)
	resultsInfo(results, noWorkItems, now, fileName, what)

	fmt.Printf("export: %v\n", export)
	message := "Data Exported to " + outPath + fileName
	trialMode := TrialMode(ssMode)
	if !trialMode {
		logs.Storing("Exporting Data to ", outPath+fileName)

	} else {
		message = "Trial Mode - No Data Written"
		logs.Storing("Exporting Data to ", outPath+fileName+" (Trial Mode - No Data Written)")

	}
	exportFile.Close()
	results.Close()
	return noWorkItems, message, false, "", nil
}

func openExportFiles(outPath string, fileName string, resultsFile string) (*os.File, *os.File) {
	logs.Information("Exporting Data", "")
	logs.Information("Opening Export to ", outPath+fileName)
	exportFile, err := os.Create(outPath + fileName)
	if err != nil {
		panic(err)
	}
	logs.Information("Opening Results File ", outPath+resultsFile)
	results, err := os.Create(outPath + resultsFile)
	if err != nil {
		panic(err)
	}
	logs.Information("Exporting Data", "Opened Files")
	exportFile.Chmod(0777)
	results.Chmod(0777)
	return exportFile, results
}

func processDataImport(inPath string, jobName string, err error, ssMode string) (int, string, string, error) {

	what := "Settings"

	logs.Processing("Importing Data from " + what)

	csvFile, err := openImportFile(inPath, what)

	fmt.Printf("csvFile: %v\n", csvFile)

	processData := []*DataItemIO{}

	if err := gocsv.UnmarshalFile(csvFile, &processData); err != nil {
		panic(err)
	}
	for _, item := range processData {
		//fmt.Println("Hello, ", item.DataID)
		if !TrialMode(ssMode) {
			var data dm.Data
			data.DataID = item.DataID
			data.Category = item.Category
			data.Class = item.Class
			data.Field = item.Field
			data.Value = item.Value
			data.Migrate = item.Migrate

			_, err := dao.Data_StoreProcess(data, Data_Job().Name)
			if err != nil {
				return 0, "Unable to Store Data", "", err
			}

		} else {
			logs.Storing("Importing Data from ", inPath+" (Trial Mode - No Data Written)")
			fmt.Printf("item: %v\n", item)
		}
	}

	return 0, "", "", nil
}

func processOriginStateExport(outPath string, jobName string, err error, ssMode string) (int, string, bool, string, error) {
	var export []OriginStateIO

	what := "OriginState"
	logs.Processing("Exporting Data from " + what)

	now, fileName, resultsFile := getExportFileNames(what)

	if outPath == "" {
		return 0, "", true, "No Output Path Specified for " + jobName, err
	}

	noitems, items, err := dao.OriginState_GetMigrateable()
	noWorkItems := noitems
	if err != nil || noWorkItems == 0 {
		return 0, "", true, "Unable to Get Data for " + jobName + "OriginState", err
	}

	for _, item := range items {
		exportItem := OriginStateIO{}
		exportItem.OriginStateID = item.OriginStateID
		exportItem.Code = item.Code
		exportItem.Name = item.Name
		exportItem.IsLocked = item.IsLocked
		exportItem.Notify = item.Notify
		exportItem.Migrate = item.Migrate
		if item.Code != "" {
			export = append(export, exportItem)
		}
	}

	exportFile, results := openExportFiles(outPath, fileName, resultsFile)

	//csvContent, err := gocsv.MarshalString(&export)
	//if err != nil {
	//	panic(err)
	//}
	err = gocsv.MarshalFile(&export, exportFile)
	if err != nil {
		panic(err)
	}
	//fmt.Println(csvContent)
	resultsInfo(results, noWorkItems, fileName, now, what)

	fmt.Printf("export: %v\n", export)
	message := "Data Exported to " + outPath + fileName
	trialMode := TrialMode(ssMode)
	if !trialMode {
		logs.Storing("Exporting Data to ", outPath+fileName)

	} else {
		message = "Trial Mode - No Data Written"
		logs.Storing("Exporting Data to ", outPath+fileName+" (Trial Mode - No Data Written)")

	}
	exportFile.Close()
	results.Close()
	return noWorkItems, message, false, "", nil
}

func resultsInfo(results *os.File, noWorkItems int, fileName string, now string, what string) {
	results.WriteString(fmt.Sprintf("Application: %s\n", core.ApplicationName()))
	results.WriteString(fmt.Sprintf("Version: %s\n", core.ReleaseIdentityVerbose()))
	results.WriteString(fmt.Sprintf("No of Items: %d \n", noWorkItems))
	results.WriteString(fmt.Sprintf("Source: %s \n", what))
	results.WriteString(fmt.Sprintf("File: %s \n", fileName))
	results.WriteString(fmt.Sprintf("Generated On: %s\n", core.ApplicationHostname()))
	results.WriteString(fmt.Sprintf("Generated By: %s\n", core.GetHostIP()))
	results.WriteString(fmt.Sprintf("Generated At: %s\n", now))
	results.WriteString(fmt.Sprintf("Generated From: %s\n", core.ApplicationName()))
	results.WriteString(fmt.Sprintf("Generated Company: %s\n", core.ApplicationCompanyName()))
	results.WriteString(fmt.Sprintf("Generated Job: %s\n", Data_Job().Name))
	results.WriteString(fmt.Sprintf("Generated DB Server: %s\n", core.ApplicationSQLServer()))
	results.WriteString(fmt.Sprintf("Generated DB Database: %s\n", core.ApplicationSQLDatabase()))
	results.WriteString(fmt.Sprintf("Generated DB Level: %s\n", core.DB_Version()))
}

func getExportFileNames(what string) (string, string, string) {
	now := time.Now().Format("20060102_150405")
	pre := "/" + what + "/"
	fileName := pre + now + ".csv"
	resultsFile := pre + now + ".nfo"
	return now, fileName, resultsFile
}

func processDocumentTypeExport(outPath string, jobName string, err error, ssMode string) (int, string, bool, string, error) {
	var export []DocumentTypeIO

	what := "DocumentType"
	logs.Processing("Exporting Data from " + what)

	now, fileName, resultsFile := getExportFileNames(what)

	if outPath == "" {
		return 0, "", true, "No Output Path Specified for " + jobName + " " + what, err
	}

	noitems, items, err := dao.DocType_GetMigrateable()
	noWorkItems := noitems
	if err != nil || noWorkItems == 0 {
		return 0, "", true, "Unable to Get Data for " + jobName + " " + what, err
	}

	for _, item := range items {
		exportItem := DocumentTypeIO{}
		exportItem.DocTypeID = item.DocTypeID
		exportItem.Code = item.Code
		exportItem.Name = item.Name
		exportItem.Comments = item.Comments
		exportItem.Migrate = item.Migrate
		if item.Code != "" {
			export = append(export, exportItem)
		}
	}

	exportFile, results := openExportFiles(outPath, fileName, resultsFile)

	//csvContent, err := gocsv.MarshalString(&export)
	//if err != nil {
	//	panic(err)
	//}
	err = gocsv.MarshalFile(&export, exportFile)
	if err != nil {
		panic(err)
	}
	//fmt.Println(csvContent)
	resultsInfo(results, noWorkItems, fileName, now, what)

	fmt.Printf("export: %v\n", export)
	message := "Data Exported to " + outPath + fileName
	trialMode := TrialMode(ssMode)
	if !trialMode {
		logs.Storing("Exporting Data to ", outPath+fileName)

	} else {
		message = "Trial Mode - No Data Written"
		logs.Storing("Exporting Data to ", outPath+fileName+" (Trial Mode - No Data Written)")

	}
	exportFile.Close()
	results.Close()
	return noWorkItems, message, false, "", nil
}

func processProjectStateExport(outPath string, jobName string, err error, ssMode string) (int, string, bool, string, error) {
	var export []ProjectStateIO

	what := "ProjectState"
	logs.Processing("Exporting Data from " + what)
	now, fileName, resultsFile := getExportFileNames(what)

	if outPath == "" {
		return 0, "", true, "No Output Path Specified for " + jobName, err
	}

	noitems, items, err := dao.ProjectState_GetMigrateable()
	noWorkItems := noitems
	if err != nil || noWorkItems == 0 {
		return 0, "", true, "Unable to Get Data for " + jobName + "OriginState", err
	}

	for _, item := range items {
		exportItem := ProjectStateIO{}
		exportItem.ProjectStateID = item.ProjectStateID
		exportItem.Code = item.Code
		exportItem.Name = item.Name
		exportItem.IsLocked = item.IsLocked
		exportItem.Notify = item.Notify
		exportItem.Migrate = item.Migrate
		if item.Code != "" {
			export = append(export, exportItem)
		}
	}

	exportFile, results := openExportFiles(outPath, fileName, resultsFile)

	//csvContent, err := gocsv.MarshalString(&export)
	//if err != nil {
	//	panic(err)
	//}
	err = gocsv.MarshalFile(&export, exportFile)
	if err != nil {
		panic(err)
	}
	//fmt.Println(csvContent)
	resultsInfo(results, noWorkItems, fileName, now, what)

	fmt.Printf("export: %v\n", export)
	message := "Data Exported to " + outPath + fileName
	trialMode := TrialMode(ssMode)
	if !trialMode {
		logs.Storing("Exporting Data to ", outPath+fileName)

	} else {
		message = "Trial Mode - No Data Written"
		logs.Storing("Exporting Data to ", outPath+fileName+" (Trial Mode - No Data Written)")

	}
	exportFile.Close()
	results.Close()
	return noWorkItems, message, false, "", nil
}

func processEstimationStateExport(outPath string, jobName string, err error, ssMode string) (int, string, bool, string, error) {
	var export []EstimationStateIO

	what := "EstimationState"
	logs.Processing("Exporting Data from " + what)
	now, fileName, resultsFile := getExportFileNames(what)

	if outPath == "" {
		return 0, "", true, "No Output Path Specified for " + jobName + " " + what, err
	}

	noitems, items, err := dao.EstimationState_GetMigrateable()
	noWorkItems := noitems
	if err != nil || noWorkItems == 0 {
		return 0, "", true, "Unable to Get Data for " + jobName + " " + what, err
	}

	for _, item := range items {
		exportItem := EstimationStateIO{}
		exportItem.EstimationStateID = item.EstimationStateID
		exportItem.Code = item.Code
		exportItem.Name = item.Name
		exportItem.IsLocked = item.IsLocked
		exportItem.Notify = item.Notify
		exportItem.Migrate = item.Migrate
		if exportItem.IsLocked == "" {
			exportItem.IsLocked = core.FALSE
		}
		if exportItem.Notify == "" {
			exportItem.Notify = core.FALSE
		}
		if item.Code != "" {
			export = append(export, exportItem)
		}
	}

	exportFile, results := openExportFiles(outPath, fileName, resultsFile)

	//csvContent, err := gocsv.MarshalString(&export)
	//if err != nil {
	//	panic(err)
	//}
	err = gocsv.MarshalFile(&export, exportFile)
	if err != nil {
		panic(err)
	}
	//fmt.Println(csvContent)
	resultsInfo(results, noWorkItems, fileName, now, what)

	fmt.Printf("export: %v\n", export)
	message := "Data Exported to " + outPath + fileName
	trialMode := TrialMode(ssMode)
	if !trialMode {
		logs.Storing("Exporting Data to ", outPath+fileName)

	} else {
		message = "Trial Mode - No Data Written"
		logs.Storing("Exporting Data to ", outPath+fileName+" (Trial Mode - No Data Written)")

	}
	exportFile.Close()
	results.Close()
	return noWorkItems, message, false, "", nil
}

func processTranslationExport(outPath string, jobName string, err error, ssMode string) (int, string, bool, string, error) {
	var export []TranslationIO

	what := "Translation"
	logs.Processing("Exporting Data from " + what)
	now, fileName, resultsFile := getExportFileNames(what)

	if outPath == "" {
		return 0, "", true, "No Output Path Specified for " + jobName + " " + what, err
	}

	noitems, items, err := dao.Translation_GetMigrateable()
	noWorkItems := noitems
	if err != nil || noWorkItems == 0 {
		return 0, "", true, "Unable to Get Data for " + jobName + " " + what, err
	}

	for _, item := range items {
		exportItem := TranslationIO{}
		exportItem.TranslationID = safe(item.Id)
		exportItem.Class = safe(item.Class)
		exportItem.Message = safe(item.Message)
		exportItem.Translation = safe(item.Translation)
		exportItem.Migrate = safe(item.Migrate)
		if item.Id != "" {
			export = append(export, exportItem)
		}
	}

	exportFile, results := openExportFiles(outPath, fileName, resultsFile)

	//csvContent, err := gocsv.MarshalString(&export)
	//if err != nil {
	//	panic(err)
	//}
	err = gocsv.MarshalFile(&export, exportFile)
	if err != nil {
		panic(err)
	}
	//fmt.Println(csvContent)
	resultsInfo(results, noWorkItems, fileName, now, what)

	fmt.Printf("export: %v\n", export)
	message := "Data Exported to " + outPath + fileName
	trialMode := TrialMode(ssMode)
	if !trialMode {
		logs.Storing("Exporting Data to ", outPath+fileName)

	} else {
		message = "Trial Mode - No Data Written"
		logs.Storing("Exporting Data to ", outPath+fileName+" (Trial Mode - No Data Written)")

	}
	exportFile.Close()
	results.Close()
	return noWorkItems, message, false, "", nil
}

func openImportFile(inPath string, what string) (*os.File, error) {

	inPath = core.ReplaceWildcard(inPath, "ENTITY", what)
	logs.Processing("Opening File for Import : " + inPath)
	// Open CSV file
	test, e := os.ReadFile(inPath)
	fmt.Println("test", test)
	fmt.Println("e", e)
	if e != nil {
		return nil, errors.New("cannot open source file : " + inPath)
	}
	csvFile, err := os.Open(inPath)
	if err != nil {
		return nil, errors.New("cannot open source file : " + inPath)
	}
	//defer csvFile.Close()
	logs.Information("Opened File for Import : ", inPath)
	return csvFile, nil
}

func processOriginStateImport(inPath string, jobName string, err error, ssMode string) (int, string, string, error) {

	what := "OriginState"
	logs.Processing("Importing Data from " + what)
	file, err := openImportFile(inPath, what)

	processData := []*OriginStateIO{}

	if err := gocsv.UnmarshalFile(file, &processData); err != nil {
		panic(err)
	}
	for _, item := range processData {
		//fmt.Println("Hello, ", item.DataID)
		if !TrialMode(ssMode) {
			var data dm.OriginState
			data.OriginStateID = item.OriginStateID
			data.Code = item.Code
			data.Name = item.Name
			data.Notify = item.Notify
			data.IsLocked = item.IsLocked
			data.Migrate = item.Migrate

			_, err := dao.OriginState_StoreProcess(data, Data_Job().Name)
			if err != nil {
				return 0, "Unable to Store Data", "", err
			}

		} else {
			logs.Storing("Importing Data from ", inPath+" (Trial Mode - No Data Written)")
			fmt.Printf("item: %v\n", item)
		}
	}

	return 0, "", "", nil
}

func processDocumentTypeImport(inPath string, jobName string, err error, ssMode string) (int, string, string, error) {

	what := "DocumentType"
	logs.Processing("Importing Data from " + what)
	csvFile, err := openImportFile(inPath, what)

	processData := []*DocumentTypeIO{}

	if err := gocsv.UnmarshalFile(csvFile, &processData); err != nil {
		panic(err)
	}
	for _, item := range processData {
		//fmt.Println("Hello, ", item.DataID)
		if !TrialMode(ssMode) {
			var data dm.DocType
			data.DocTypeID = item.DocTypeID
			data.Code = item.Code
			data.Name = item.Name
			data.Comments = item.Comments
			data.Migrate = item.Migrate

			_, err := dao.DocType_StoreProcess(data, Data_Job().Name)
			if err != nil {
				return 0, "Unable to Store Data", "", err
			}

		} else {
			logs.Storing("Importing Data from ", inPath+" (Trial Mode - No Data Written)")
			fmt.Printf("item: %v\n", item)
		}
	}

	return 0, "", "", nil
}

func processProjectStateImport(inPath string, jobName string, err error, ssMode string) (int, string, string, error) {

	what := "ProjectState"
	logs.Processing("Importing Data from " + what)
	csvFile, err := openImportFile(inPath, what)

	processData := []*ProjectStateIO{}

	if err := gocsv.UnmarshalFile(csvFile, &processData); err != nil {
		panic(err)
	}
	for _, item := range processData {
		//fmt.Println("Hello, ", item.DataID)
		if !TrialMode(ssMode) {
			var data dm.ProjectState
			data.ProjectStateID = item.ProjectStateID
			data.Code = item.Code
			data.Name = item.Name
			data.IsLocked = item.IsLocked
			data.Notify = item.Notify
			data.Migrate = item.Migrate

			_, err := dao.ProjectState_StoreProcess(data, Data_Job().Name)
			if err != nil {
				return 0, "Unable to Store Data", "", err
			}

		} else {
			logs.Storing("Importing Data from ", inPath+" (Trial Mode - No Data Written)")
			fmt.Printf("item: %v\n", item)
		}
	}

	return 0, "", "", nil
}

func processEstimationStateImport(inPath string, jobName string, err error, ssMode string) (int, string, string, error) {

	what := "EstimationState"
	logs.Processing("Importing Data from " + what)
	csvFile, err := openImportFile(inPath, what)

	processData := []*EstimationStateIO{}

	if err := gocsv.UnmarshalFile(csvFile, &processData); err != nil {
		panic(err)
	}
	for _, item := range processData {
		//fmt.Println("Hello, ", item.DataID)
		if !TrialMode(ssMode) {
			var data dm.EstimationState
			data.EstimationStateID = item.EstimationStateID
			data.Code = item.Code
			data.Name = item.Name
			data.IsLocked = item.IsLocked
			data.Notify = item.Notify
			data.Migrate = item.Migrate

			_, err := dao.EstimationState_StoreProcess(data, Data_Job().Name)
			if err != nil {
				return 0, "Unable to Store Data", "", err
			}

		} else {
			logs.Storing("Importing Data from ", inPath+" (Trial Mode - No Data Written)")
			fmt.Printf("item: %v\n", item)
		}
	}

	return 0, "", "", nil
}

func processTranslationImport(inPath string, jobName string, err error, ssMode string) (int, string, string, error) {

	what := "Translation"
	logs.Processing("Importing Data from " + what)
	csvFile, err := openImportFile(inPath, what)

	processData := []*TranslationIO{}

	if err := gocsv.UnmarshalFile(csvFile, &processData); err != nil {
		panic(err)
	}
	for _, item := range processData {
		//fmt.Println("Hello, ", item.DataID)
		if !TrialMode(ssMode) {
			var data dm.Translation
			data.Id = item.TranslationID
			data.Class = item.Class
			data.Message = item.Message
			data.Translation = item.Translation
			data.Migrate = item.Migrate

			_, err := dao.Translation_StoreProcess(data, Data_Job().Name)
			if err != nil {
				return 0, "Unable to Store Data", "", err
			}

		} else {
			logs.Storing("Importing Data from ", inPath+" (Trial Mode - No Data Written)")
			fmt.Printf("item: %v\n", item)
		}
	}

	return 0, "", "", nil
}

func safe(in string) string {
	return in
}
