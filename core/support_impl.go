package core

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
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

// Converts a date, as a user readable date
func dateToUserDate(in string) string {

	t, err := time.Parse(DATEFORMAT, in)
	if err != nil {
		log.Println(err.Error())
	}

	ext := t.Format(DATEFORMATUSER)

	return ext
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

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
		//templateName = roleTemplate
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

// ipRange - a structure that holds the start and end of a range of ip addresses
type ipRange struct {
	start net.IP
	end   net.IP
}

// inRange - check to see if a given ip address is within a range given
func inRange(r ipRange, ipAddress net.IP) bool {
	// strcmp type byte comparison
	if bytes.Compare(ipAddress, r.start) >= 0 && bytes.Compare(ipAddress, r.end) < 0 {
		return true
	}
	return false
}

var privateRanges = []ipRange{
	ipRange{
		start: net.ParseIP("10.0.0.0"),
		end:   net.ParseIP("10.255.255.255"),
	},
	ipRange{
		start: net.ParseIP("100.64.0.0"),
		end:   net.ParseIP("100.127.255.255"),
	},
	ipRange{
		start: net.ParseIP("172.16.0.0"),
		end:   net.ParseIP("172.31.255.255"),
	},
	ipRange{
		start: net.ParseIP("192.0.0.0"),
		end:   net.ParseIP("192.0.0.255"),
	},
	ipRange{
		start: net.ParseIP("192.168.0.0"),
		end:   net.ParseIP("192.168.255.255"),
	},
	ipRange{
		start: net.ParseIP("198.18.0.0"),
		end:   net.ParseIP("198.19.255.255"),
	},
}

// isPrivateSubnet - check to see if this ip is in a private subnet
func isPrivateSubnet(ipAddress net.IP) bool {
	// my use case is only concerned with ipv4 atm
	if ipCheck := ipAddress.To4(); ipCheck != nil {
		// iterate over all our ranges
		for _, r := range privateRanges {
			// check if this ip is in a private range
			if inRange(r, ipAddress) {
				return true
			}
		}
	}
	return false
}

// GetIPAdress returns the current IP address
func GetIPAdress(r *http.Request) string {
	var ipAddress string
	for _, h := range []string{"X-Forwarded-For", "X-Real-Ip"} {
		for _, ip := range strings.Split(r.Header.Get(h), ",") {
			// header can contain spaces too, strip those out.
			ip = strings.TrimSpace(ip)
			realIP := net.ParseIP(ip)
			if !realIP.IsGlobalUnicast() || isPrivateSubnet(realIP) {
				// bad address, go to next
				continue
			} else {
				ipAddress = ip
				goto Done
			}
		}
	}
Done:
	return ipAddress
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
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	// Convert []byte to string and print to screen
	return string(content), err
}

func GetDirectoryContentAbsolute(path string) ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(path + "/")
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
	logs.Accessing(absFileName)
	// Check it exists - If not create it
	if !(FileExists(absFileName)) {
		FileSystem_WriteData_Absolute(fileName, path, "")
	}

	//log.Println("Read          :", filePath)
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(absFileName)
	if err != nil {
		log.Fatal(err)
	}
	// Convert []byte to string and print to screen
	return content, string(content), err
}

func FileSystem_WriteData_Absolute(fileName string, path string, content string) int {

	absFileName := path + "/" + fileName

	message := []byte(content)
	err := ioutil.WriteFile(absFileName, message, 0644)
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
	return ApplicationProperties["releaseid"]
}

func ProjectID() string {
	return ApplicationProperties["project"]
}

func ReleaseLevel() string {
	return ApplicationProperties["releaselevel"]
}

func ReleaseNumber() string {
	return ApplicationProperties["releasenumber"]
}

func ApplicationName() string {
	return ApplicationProperties["appname"]
}

func ApplicationHostname() string {
	tmpHostname, _ := os.Hostname()
	return tmpHostname
}

func ApplicationSQLServer() string {
	return ApplicationPropertiesDB["server"]
}

func ApplicationSQLSchemaParent() string {
	return ApplicationPropertiesDB["parentschema"]
}

func GetSQLSchema(in map[string]string) string {
	return in["schema"]
}
func ApplicationSQLSchema() string {
	return GetSQLSchema(ApplicationPropertiesDB)
}

func ApplicationSQLDatabase() string {
	return ApplicationPropertiesDB["database"]
}

func ApplicationSessionLife() string {
	return ApplicationProperties["sessionlife"]
}

func ApplicationHTTPProtocol() string {
	return ApplicationProperties["protocol"]
}

func ApplicationHTTPPort() string {
	return ApplicationProperties["port"]
}

func ApplicationEnvironment() string {
	return ApplicationProperties["environment"]
}

func ApplicationCertificatePath() string {
	return ApplicationProperties["certpath"]
}

func ApplicationCertificateName() string {
	return ApplicationProperties["certname"]
}

func ApplicationCredentialsLife() string {
	return ApplicationProperties["credentialslife"]
}

func ApplicationGetLicenseName() string {
	return ApplicationProperties["licname"]
}

func ApplicationGetLicenseLink() string {
	return ApplicationProperties["liclink"]
}

func ApplicationToken() string {
	return ApplicationProperties["applicationtoken"]
}

func SetApplicationSQLDatabase(db string) {
	ApplicationPropertiesDB["database"] = db
}

func DataLoaderArtifactRepository() string {
	return DataLoaderArtifactRepository()
}

func OBSOLETE_MessageQueueDeliveryPath() string {
	return ApplicationProperties["deliverpath"]
}

func OBSOLETE_MessageQueueRecievePath() string {
	return ApplicationProperties["receivepath"]
}

func OBSOLETE_MessageQueueProcessedPath() string {
	return ApplicationProperties["processedpath"]
}
