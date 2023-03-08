package core

import (
	"database/sql"
	"os"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jimlawless/cfg"
	"github.com/mt1976/ebEstimates/logs"
	"gopkg.in/gomail.v2"
)

var startTime = time.Now()
var SessionToken = ""
var UUID = "authorAdjust"
var PWD = ""
var SecurityViolation = ""
var DB *sql.DB
var SystemHostname string

var ApplicationProperties Congiguration
var ApplicationPropertiesDB Congiguration
var ApplicationStubLists Congiguration
var ApplicationDB *sql.DB

var InstanceProperties Congiguration
var MasterPropertiesDB Congiguration
var MasterDB *sql.DB

var SessionManager *scs.SessionManager
var Emailer *gomail.Dialer

var IsChildInstance bool
var ApplicationCache Cache
var API Catalog

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
type Congiguration struct {
	properties map[string]string
}
type CongigurationInterface interface {
	Load(string) map[string]string
	Get(string) string
	Override(string, string)
}

func (c *Congiguration) Load(inPropertiesFile string) map[string]string {
	c.properties = getPropertiesFromFile(inPropertiesFile)
	return c.properties
}
func (c *Congiguration) Get(inProperty string) string {
	return c.properties[inProperty]
}
func (c *Congiguration) Override(inProperty string, inValue string) {
	c.properties[inProperty] = inValue
}

// GetApplicationProperty returns the value of a property configuration for the application
// GetApplicationProperty checks the environment variable APP_<PROPERTY> first and if not found uses the value in the properties file
func GetApplicationProperty(inProperty string) string {
	low_var := strings.ToLower(inProperty)
	rtn_var := ApplicationProperties.Get(low_var)
	env_var := "APP_" + strings.ToUpper(inProperty)

	//logs.Accessing("GetApplicationProperty : " + inProperty + " " + env_var)

	xxx := os.Getenv(env_var)
	if xxx != "" {
		logs.Override(inProperty, rtn_var, env_var, xxx)
		return xxx
	}

	//logs.Information("GetApplicationProperty", rtn_var)
	return rtn_var
}

// GetDatabaseProperty returns the value of a property configuration for the database
// GetDatabaseProperty checks the environment variable DB_<PROPERTY> first and if not found uses the value in the properties file
func GetDatabaseProperty(inProperty string) string {
	//logs.Accessing("GetDatabaseProperty : " + inProperty)
	rtn_var := ApplicationPropertiesDB.Get(inProperty)
	env_var := "DB_" + strings.ToUpper(inProperty)
	logs.Accessing("GetDBProperty : " + inProperty + " " + env_var)
	xxx := os.Getenv(env_var)
	if xxx != "" {
		logs.Override(inProperty, rtn_var, env_var, xxx)
		return xxx
	}
	return rtn_var
}

// Load a Properties File
func getPropertiesFromFile(inPropertiesFile string) map[string]string {
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
	err := os.WriteFile(filePath, message, 0644)
	if err != nil {
		logs.Fatal("Write Error", err)
		return false, err
	}
	return false, nil
}

// deleteDataFile deletes a file - For Future Use
// func deleteDataFile(fileName string, path string) int {
// 	pwd, _ := os.Getwd()
// 	filePath := pwd + "/" + fileName
// 	if len(path) != 0 {
// 		filePath = pwd + path + "/" + fileName
// 	}
// 	//log.Println("Delete        :", filePath)

// 	// delete file

// 	if fileExists(filePath) {
// 		var err = os.Remove(filePath)
// 		if err != nil {
// 			logs.Fatal("File Error", err)
// 			return -1
// 		}
// 	}
// 	logs.Information("File Deleted", fileName+" - "+path)
// 	return 1
// }

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

// GetDataList gets a list from a properties file
func GetDataList(basePath string) (int, []string, error) {

	var listing []string
	pwd, _ := os.Getwd()
	files, err := os.ReadDir(pwd + basePath)
	if err != nil {
		logs.Fatal("Directory Error", err)
	}

	for _, k := range files {
		listing = append(listing, k.Name())
	}

	return len(files), listing, nil
}

func (dbconnectionstring *DBConnectionString) GetID() string {
	return dbconnectionstring.ID
}

func (dbconnectionstring *DBConnectionString) GetServer() string {
	return dbconnectionstring.Server
}

func (dbconnectionstring *DBConnectionString) GetPort() string {
	return dbconnectionstring.Port
}

func (dbconnectionstring *DBConnectionString) GetUser() string {
	return dbconnectionstring.User
}

func (dbconnectionstring *DBConnectionString) GetPassword() string {
	return dbconnectionstring.Password
}

func (dbconnectionstring *DBConnectionString) GetDatabase() string {
	return dbconnectionstring.Database
}

func (dbconnectionstring *DBConnectionString) GetSchema() string {
	return dbconnectionstring.Schema
}

func (dbconnectionstring *DBConnectionString) GetActive() string {
	return dbconnectionstring.Active
}

func (dbconnectionstring *DBConnectionString) GetSYSCreated() string {
	return dbconnectionstring.SYSCreated
}

func (dbconnectionstring *DBConnectionString) GetSYSWho() string {
	return dbconnectionstring.SYSWho
}

func (dbconnectionstring *DBConnectionString) GetSYSHost() string {
	return dbconnectionstring.SYSHost
}

func (dbconnectionstring *DBConnectionString) GetSYSUpdated() string {
	return dbconnectionstring.SYSUpdated
}

func (dbconnectionstring *DBConnectionString) SetID(ID string) *DBConnectionString {
	dbconnectionstring.ID = ID
	return dbconnectionstring
}

func (dbconnectionstring *DBConnectionString) SetServer(Server string) *DBConnectionString {
	dbconnectionstring.Server = Server
	return dbconnectionstring
}

func (dbconnectionstring *DBConnectionString) SetPort(Port string) *DBConnectionString {
	dbconnectionstring.Port = Port
	return dbconnectionstring
}

func (dbconnectionstring *DBConnectionString) SetUser(User string) *DBConnectionString {
	dbconnectionstring.User = User
	return dbconnectionstring
}

func (dbconnectionstring *DBConnectionString) SetPassword(Password string) *DBConnectionString {
	dbconnectionstring.Password = Password
	return dbconnectionstring
}

func (dbconnectionstring *DBConnectionString) SetDatabase(Database string) *DBConnectionString {
	dbconnectionstring.Database = Database
	return dbconnectionstring
}

func (dbconnectionstring *DBConnectionString) SetSchema(Schema string) *DBConnectionString {
	dbconnectionstring.Schema = Schema
	return dbconnectionstring
}

func (dbconnectionstring *DBConnectionString) SetActive(Active string) *DBConnectionString {
	dbconnectionstring.Active = Active
	return dbconnectionstring
}

func (dbconnectionstring *DBConnectionString) SetSYSCreated(SYSCreated string) *DBConnectionString {
	dbconnectionstring.SYSCreated = SYSCreated
	return dbconnectionstring
}

func (dbconnectionstring *DBConnectionString) SetSYSWho(SYSWho string) *DBConnectionString {
	dbconnectionstring.SYSWho = SYSWho
	return dbconnectionstring
}

func (dbconnectionstring *DBConnectionString) SetSYSHost(SYSHost string) *DBConnectionString {
	dbconnectionstring.SYSHost = SYSHost
	return dbconnectionstring
}

func (dbconnectionstring *DBConnectionString) SetSYSUpdated(SYSUpdated string) *DBConnectionString {
	dbconnectionstring.SYSUpdated = SYSUpdated
	return dbconnectionstring
}
func (dateitem *DateItem) GetToday() string {
	return dateitem.Today
}

func (dateitem *DateItem) GetInternal() time.Time {
	return dateitem.Internal
}

func (dateitem *DateItem) GetDefault() string {
	return dateitem.Default
}

func (dateitem *DateItem) GetYYYYMMDD() string {
	return dateitem.YYYYMMDD
}

func (dateitem *DateItem) GetPICKEpoch() string {
	return dateitem.PICKEpoch
}

func (dateitem *DateItem) SetToday(Today string) *DateItem {
	dateitem.Today = Today
	return dateitem
}

func (dateitem *DateItem) SetInternal(Internal time.Time) *DateItem {
	dateitem.Internal = Internal
	return dateitem
}

func (dateitem *DateItem) SetDefault(Default string) *DateItem {
	dateitem.Default = Default
	return dateitem
}

func (dateitem *DateItem) SetYYYYMMDD(YYYYMMDD string) *DateItem {
	dateitem.YYYYMMDD = YYYYMMDD
	return dateitem
}

func (dateitem *DateItem) SetPICKEpoch(PICKEpoch string) *DateItem {
	dateitem.PICKEpoch = PICKEpoch
	return dateitem
}
