package core

import (
	"database/sql"
	"io/ioutil"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jimlawless/cfg"
	"github.com/mt1976/ebEstimates/logs"
)

var startTime = time.Now()
var SessionToken = ""
var UUID = "authorAdjust"
var PWD = ""
var SecurityViolation = ""
var DB *sql.DB
var SystemHostname string

var ApplicationProperties map[string]string
var ApplicationPropertiesDB map[string]string
var ApplicationStubLists map[string]string
var ApplicationDB *sql.DB

var InstanceProperties map[string]string
var MasterPropertiesDB map[string]string
var MasterDB *sql.DB

var SessionManager *scs.SessionManager

var IsChildInstance bool

type DBConnectionString struct {
	ID         string
	Server     string
	Port       string
	User       string
	Password   string
	Database   string
	Schema     string
	Active     string
	SYSCreated string
	SYSWho     string
	SYSHost    string
	SYSUpdated string
}

type DateItem struct {
	Today     string
	Internal  time.Time
	Default   string
	YYYYMMDD  string
	PICKEpoch string
}

func GetApplicationProperty(inProperty string) string {
	return ApplicationProperties[inProperty]
}

// Load a Properties File
func getProperties(inPropertiesFile string) map[string]string {
	wctProperties := make(map[string]string)
	//machineName, _ := os.Hostname()
	// For docker - if can't find properties file (create one from the template properties file)
	propertiesFileName := "config/" + inPropertiesFile
	if fileExists(propertiesFileName) {
		// Do nothign this is ok
	} else {

		ok := copyDataFile(inPropertiesFile, "config/", "config/fileSystem/config")
		if !ok {
			logs.Error("Issue in Copy Function", nil)
		}
	}

	err := cfg.Load(propertiesFileName, wctProperties)
	if err != nil {
		logs.Fatal("Cannot Load Properties File "+propertiesFileName, err)
	}
	return wctProperties
}

// FileExists returns true if the specified file existing on the filesystem
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func PreInitialise() {

}

func writeDataFile(fileName string, path string, content string) (bool, error) {
	pwd, _ := os.Getwd()
	filePath := pwd + "/" + fileName
	if len(path) != 0 {
		filePath = pwd + path + "/" + fileName
	}
	//log.Println("Write         :", filePath)

	message := []byte(content)
	err := ioutil.WriteFile(filePath, message, 0644)
	if err != nil {
		logs.Fatal("Write Error", err)
		return false, err
	}
	return false, nil
}

func deleteDataFile(fileName string, path string) int {
	pwd, _ := os.Getwd()
	filePath := pwd + "/" + fileName
	if len(path) != 0 {
		filePath = pwd + path + "/" + fileName
	}
	//log.Println("Delete        :", filePath)

	// delete file

	if fileExists(filePath) {
		var err = os.Remove(filePath)
		if err != nil {
			logs.Fatal("File Error", err)
			return -1
		}
	}
	logs.Information("File Deleted", fileName+" - "+path)
	return 1
}

func copyDataFile(fileName string, fromPath string, toPath string) bool {

	logs.Warning("Copying " + fileName + " from " + fromPath + " to " + toPath)

	content, err := ReadDataFile(fileName, fromPath)
	if err != nil {
		logs.Fatal("File Read Error", err)
	}

	ok, err2 := writeDataFile(fileName, toPath, content)
	if err2 != nil {
		logs.Fatal("File Write Error", err2)
	}

	if !ok {
		logs.Fatal("Unable to Copy "+fileName+" from "+fromPath+" to "+toPath, nil)
	}

	return true
}

// getFundsCheckList read all employees
func GetDataList(basePath string) (int, []string, error) {

	var listing []string
	//	log.Println(basePath, kind, direction, requestPath)
	pwd, _ := os.Getwd()
	//logs.Accessing(pwd + basePath)
	files, err := ioutil.ReadDir(pwd + basePath)
	if err != nil {
		logs.Fatal("Directory Error", err)
	}

	//logs.Information("Files Found", strconv.Itoa(len(files)))

	for _, k := range files {
		//fmt.Println("key:", k)
		listing = append(listing, k.Name())
	}

	//count, simFundsCheckList, _, _ := fetchFundsCheckData("")
	return len(files), listing, nil
}
