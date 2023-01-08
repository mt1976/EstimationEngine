package core

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	uuid "github.com/google/uuid"

	"github.com/mt1976/ebEstimates/logs"
)

// // Converts a date, as a user readable date
// func dateToUserDate(in string) string {

// 	t, err := time.Parse(DATEFORMAT, in)
// 	if err != nil {
// 		log.Println(err.Error())
// 	}

// 	ext := t.Format(DATEFORMATUSER)

// 	return ext
// }

// GetURLparam returns a selected parmeter value from the a given URI
func GetURLparam(r *http.Request, paramID string) string {
	//fmt.Println(paramID)
	//fmt.Println(r.URL)
	key := r.FormValue(paramID)
	//log.Printf("URL Parameter : Key=%q Value=%q", paramID, string(key))
	logs.Information("URL Parameter :", fmt.Sprintf("Key=%q Value=%q", paramID, string(key)))
	return key
}

// RemoveContents clears the contents of a specified directory
func RemoveContents(dir string) error {
	log.Println("TRASH", dir)
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		log.Println(err)
		return err
	}
	//	fmt.Println("do Clear", files)
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return err
}

// GetTemplateID returns the template to use when providing HTML page templates
func GetTemplateID(tmpl string, userRole string) string {
	returnTemplate := ""
	templateName := "html/base/" + tmpl + ".html"
	roleTemplate := "html/roles/" + userRole + "/" + tmpl + ".html"
	versionTemplate := "html/impl/" + tmpl + ".html"
	fallbackTemplate := "html/" + tmpl + ".html"
	helpTemplate := "html/Impl_Oops.html"

	//logs.Information("Template Checks:", tmpl+" "+templateName+" "+roleTemplate+" "+versionTemplate)
	//log.Println("Testing", roleTemplate, FileExists(roleTemplate))
	//log.Println("Testing", templateName, FileExists(templateName))

	if FileExists(fallbackTemplate) {
		returnTemplate = fallbackTemplate
	}

	if FileExists(templateName) {
		returnTemplate = templateName
	}

	if FileExists(versionTemplate) {
		returnTemplate = versionTemplate
	}

	if userRole != "" {
		if FileExists(roleTemplate) {
			returnTemplate = roleTemplate
		}
	}
	if returnTemplate == "" {
		log.Println("Template not found:", tmpl)
		pErr := errors.New("Template not found : " + tmpl + " " + templateName + " " + roleTemplate + " " + versionTemplate + " " + fallbackTemplate + " " + helpTemplate)
		//panic(pErr)
		logs.Warning(pErr.Error())
		returnTemplate = helpTemplate
	}
	//log.Printf("Using Template: Source %q", templateName)
	logs.Template(returnTemplate)
	return returnTemplate
}

// GetMenuID returns the menu template to use when providing HTML meny templates
func GetMenuID(tmpl string, userRole string) string {
	templateName := "config/menu/" + tmpl + ".json"
	roleTemplate := templateName
	if len(userRole) != 0 {
		roleTemplate = "config/menu/" + userRole + "/" + tmpl + ".json"
	}
	//log.Println("Testing", roleTemplate, FileExists(roleTemplate))
	//log.Println("Testing", templateName, FileExists(templateName))
	if FileExists(roleTemplate) {
		templateName = roleTemplate
	}
	//log.Printf("Using Menu    : Source %q", templateName)
	logs.Menu(templateName)
	return templateName
}

// GetNavigationID returns the navigation pane template to use when providing navigation templates
func GetNavigationID(inUserRole string) string {
	templateName := "../assets/navigation.html"
	roleTemplate := "../assets/" + inUserRole + "_navigation.html"
	//log.Println("Testing", templateName, FileExists(templateName))
	//log.Println("Testing", roleTemplate, FileExists(roleTemplate))
	if FileExists(roleTemplate) {
		templateName = roleTemplate
	}
	//log.Println("NAVIGATION", templateName)
	return templateName
}

// FileExists returns true if the specified file existing on the filesystem
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// ReadUserIP returns the end-users IP address
func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

// GetLocalIP returns the local IP address
func GetLocalIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	//handle err...

	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.String()
}

// SqlDateToHTMLDate returns a string version of a SQL Date
func SqlDateToHTMLDate(inDate string) (outDate string) {
	var rtnDate string
	if inDate != "" {
		rtnDate = inDate[0:10]
	}
	return rtnDate
}

func GetHostIP() string {
	var ip string
	ip = ""
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		//fmt.Printf("%d %v\n", i, addr)
		ip = addr.String()
		return ip
	}
	return ""
}

// GetIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func GetIncomingRequestIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func serviceMessage(i string) {
	//msg := "Servicing     : %q"
	//log.Printf(msg, i)
	logs.Servicing(i)
}
func serviceMessageAction(i string, act string, id string) {
	msgSuffix := fmt.Sprintf("%s %s %q", i, act, id)
	serviceMessage(msgSuffix)
}
func ServiceMessageAction(i string, act string, id string) {
	serviceMessageAction(i, act, id)
}
func ServiceMessage(i string) {
	serviceMessage(i)
}

func ReadDataFile(fileName string, path string) (string, error) {
	pwd, _ := os.Getwd()
	filePath := pwd + "/" + fileName
	if len(path) != 0 {
		filePath = pwd + path + "/" + fileName
	}

	// Check it exists - If not create it
	if !(FileExists(filePath)) {
		FileSystem_WriteData(fileName, path, "")
	}

	//log.Println("Read          :", filePath)
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	// Convert []byte to string and print to screen
	return string(content), err
}

func GetDirectoryContentAbsolute(path string) ([]fs.DirEntry, error) {
	files, err := os.ReadDir(path + "/")
	if err != nil {
		log.Fatal(err)
	}
	return files, err
}

func ReadDataFileAbsolute(fileName string, path string) ([]byte, string, error) {

	pwd, _ := os.Getwd()
	//first char is a . replace with path
	path = strings.Replace(path, "..", pwd, 1)
	path = strings.Replace(path, ".", pwd, 1)

	absFileName := path + "/" + fileName
	//logs.Accessing(absFileName)
	// Check it exists - If not create it
	if !(FileExists(absFileName)) {
		FileSystem_WriteData_Absolute(fileName, path, "")
	}

	//log.Println("Read          :", filePath)
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := os.ReadFile(absFileName)
	if err != nil {
		log.Fatal(err)
	}
	// Convert []byte to string and print to screen
	return content, string(content), err
}

func FileSystem_WriteData_Absolute(fileName string, path string, content string) int {

	absFileName := path + "/" + fileName

	message := []byte(content)
	err := os.WriteFile(absFileName, message, 0644)
	if err != nil {
		log.Fatal(err)
		return -1
	}

	//log.Println("File Write : " + fileName + " in " + path + "[" + absFileName + "]")
	logs.Saving(fileName, absFileName)
	return 1
}

func FileSystem_WriteData(fileName string, path string, content string) int {

	return FileSystem_WriteData_Absolute(fileName, path, content)

}

func DeleteDataFileAbsolute(fileName string, path string) int {

	absFileName := path + "/" + fileName
	if FileExists(absFileName) {
		var err = os.Remove(absFileName)
		if err != nil {
			log.Fatal(err.Error())
			return -1
		}
	}
	log.Println("File Deletion : " + fileName + " in " + path + "[" + absFileName + "]")
	return 1

}

func DeleteDataFile(fileName string, path string) int {
	pwd, _ := os.Getwd()
	filePath := pwd + "/" + fileName
	if len(path) != 0 {
		filePath = pwd + path + "/" + fileName
	}
	//log.Println("Delete        :", filePath)

	// delete file

	if FileExists(filePath) {
		var err = os.Remove(filePath)
		if err != nil {
			log.Fatal(err.Error())
			return -1
		}
	}
	log.Println("File Deletion : " + fileName + " in " + path + "[" + filePath + "]")
	return 1
}

func Logit(actionType string, data string) {
	//_, caller, _, _ := runtime.Caller(1)
	//outcall := strings.Split(caller, "/")
	///depth := len(outcall) - 1
	//depth2 := depth - 1
	//log.Println(len(outcall), depth, depth2)
	//callerName := outcall[depth2] + "/" + outcall[depth]
	log.Println("Information   :", data)
	logs.Information(data, "")
	//	log.Println(callerName, actionType, data)
}

func GetUUID() string {
	return uuid.NewString()
}

func ReleaseIdentity() string {
	return ProjectID() + ReleaseID()
}

func ReleaseIdentityVerbose() string {
	return fmt.Sprintf("%s [r%s-%s]", ReleaseIdentity(), ReleaseLevel(), ReleaseNumber())
}

func GetCookieIdentity() string {
	newidentity := EncodeString(ReleaseIdentity())
	return strings.ReplaceAll(RemoveSpecialChars(newidentity), "=", "")
}

func ReleaseID() string {
	return GetApplicationProperty("releaseid")
}

func ProjectID() string {
	return GetApplicationProperty("project")
}

func ReleaseLevel() string {
	return GetApplicationProperty("releaselevel")
}

func ReleaseNumber() string {
	return GetApplicationProperty("releasenumber")
}

func ApplicationName() string {
	return GetApplicationProperty("appname")
}

func ApplicationCompanyName() string {
	return GetApplicationProperty("companyname")
}

func ApplicationLicenseLink() string {
	return GetApplicationProperty("liclink")
}

func ApplicationLicenseName() string {
	return GetApplicationProperty("licname")
}

func ApplicationHostname() string {
	tmpHostname, _ := os.Hostname()
	return tmpHostname
}

func ApplicationSQLServer() string {
	return GetDatabaseProperty("server")
}

func ApplicationSQLSchemaParent() string {
	return GetDatabaseProperty("parentschema")
}

func GetSQLSchema(in map[string]string) string {
	return GetDatabaseProperty("schema")
}

func ApplicationSQLSchema() string {
	return GetDatabaseProperty("schema")
}

func ApplicationSQLDatabase() string {
	return GetDatabaseProperty("database")
}

func ApplicationSessionLife() string {
	return GetApplicationProperty("sessionlife")
}

func ApplicationHTTPProtocol() string {
	return GetApplicationProperty("protocol")
}

func ApplicationHTTPPort() string {
	return GetApplicationProperty("port")
}

func ApplicationEnvironment() string {
	return GetApplicationProperty("environment")
}

func ApplicationCertificatePath() string {
	return GetApplicationProperty("certpath")
}

func ApplicationCertificateName() string {
	return GetApplicationProperty("certname")
}

func ApplicationCredentialsLife() string {
	return GetApplicationProperty("credentialslife")
}

func ApplicationGetLicenseName() string {
	return GetApplicationProperty("licname")
}

func ApplicationGetLicenseLink() string {
	return GetApplicationProperty("liclink")
}

func ApplicationToken() string {
	return GetApplicationProperty("applicationtoken")
}

func SetApplicationSQLDatabase(db string) {
	ApplicationPropertiesDB["database"] = db
}

func DataLoaderArtifactRepository() string {
	return ""
}

func OBSOLETE_MessageQueueDeliveryPath() string {
	return ""
}

func OBSOLETE_MessageQueueRecievePath() string {
	return ""
}

func OBSOLETE_MessageQueueProcessedPath() string {
	return ""
}

func AddActivity(in string, what string) string {
	return AddActivity_ForUser(in, what, "System")
}

func AddActivity_ForProcess(in string, what string, process string) string {
	return AddActivity_ForUser(in, what, process)
}

func AddActivity_ForUser(in string, what string, un string) string {
	if what == "" {
		return in
	}

	logs.Information("Activity", what+" "+un)

	tm := time.Now().Format("02/01/2006 15:04:05")
	if in == "" {
		return tm + " " + un + " : " + what
	}
	out := in + "\n" + tm + " " + un + " : " + what
	return out
}

func RunningInDockerContainer() bool {
	// docker creates a .dockerenv file at the root
	// of the directory tree inside the container.
	// if this file exists then the viewer is running
	// from inside a container so return true

	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}
	return false
}
