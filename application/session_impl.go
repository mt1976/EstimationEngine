package application

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/nullrocks/identicon"

	core "github.com/mt1976/ebEstimates/core"
	dao "github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

type sessionToken struct {
	SecurityViolation string
	ResponseCode      string
	UID               string
}

// Session_Publish annouces the endpoints available for this object
func Session_Publish_Impl(mux http.ServeMux) {
	mux.HandleFunc("/login", Session_HandlerValidateLogin)
	mux.HandleFunc("/request", Session_HandlerRegister)

	logs.Publish("Application", dm.Session_Title+" Impl")
}

func Session_HandlerValidateLogin(w http.ResponseWriter, r *http.Request) {

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	uName := r.FormValue("username")
	uPassword := r.FormValue("password")
	appToken := core.ApplicationToken()

	tok := session_ValidateLogin(appToken, uName, uPassword, r)
	//fmt.Printf("appToken: %v\n", appToken)
	//fmt.Printf("tok: %v\n", tok)
	//fmt.Printf("uName: %v\n", uName)
	//fmt.Printf("uPassword: %v\n", uPassword)
	lastLoginUsageText := "The last sucessful login for a user."
	faildLoginUsageText := "The last failed login for a user."

	if tok.ResponseCode == "200" {
		core.SecurityViolation = ""
		core.ServiceMessageAction("ACCESS GRANTED", Session_GetUserName(r), tok.ResponseCode)
		logs.Result("Redirecting to: /", "/home")
		//fmt.Printf("core.SessionManager.Get(r.Context(), core.SessionUserName): %v\n", core.SessionManager.Get(r.Context(), core.SessionUserName))
		//spew.Dump(r.Context())
		// Generate a avatar
		GenerateAvatar(uName, tok.UID)

		// Store Last Login Details
		dao.Data_Put("Credentials", uName, dm.Data_Category_LastSuccesfulLogin, time.Now().Format(core.DATEMSG), lastLoginUsageText)
		http.Redirect(w, r, "/home", http.StatusFound)

	} else {
		if tok.ResponseCode == "303" {
			// Special Case - User neets to change password
			core.SecurityViolation = ""
			core.ServiceMessageAction("ACCESS GRANTED TO CHANGE PASSWORD", Session_GetUserName(r), tok.ResponseCode)
			GenerateAvatar(uName, tok.UID)
			logs.Result("Redirecting to: ", dm.CredentialsPassword_PathChange)
			dao.Data_Put("Credentials", uName, dm.Data_Category_LastSuccesfulLogin, time.Now().Format(core.DATEMSG), lastLoginUsageText)
			http.Redirect(w, r, dm.CredentialsPassword_PathChange+"?ID=", http.StatusFound)
		} else {
			core.SecurityViolation = tok.SecurityViolation
			core.ServiceMessageAction(tok.SecurityViolation, Session_GetUserName(r), tok.ResponseCode)
			dao.Data_Put("Credentials", uName, dm.Data_Category_LastFailedLogin, time.Now().Format(core.DATEMSG), faildLoginUsageText)
			http.Redirect(w, r, "/logout", http.StatusFound)
		}
	}
}

func GenerateAvatar(seed string, id string) {
	logs.Information("Generating Avatar for: ", seed+" "+id)
	// New Generator: Rehuse
	ig, err := identicon.New(
		core.ApplicationToken(), // Namespace
		7,                       // Number of blocks (Size)
		7,                       // Density
	)

	if err != nil {
		panic(err) // Invalid Size or Density
	}

	username := seed             // Text - decides the resulting figure
	ii, err := ig.Draw(username) // Generate an IdentIcon

	if err != nil {
		panic(err) // Text is empty
	}

	// File writer
	pwd, _ := os.Getwd()
	filenamepath := pwd + "/data/images/avatars/" + id + ".png"
	// check if the file exists
	//logs.Information("Avatar File: ", filenamepath)

	img, _ := os.Create(filenamepath)
	defer img.Close()

	// Takes the size in pixels and any io.Writer
	ii.Png(300, img) // 300px * 300px
	logs.Information("Avatar Used: ", filenamepath)
}

func session_ValidateLogin(appToken string, username string, password string, r *http.Request) sessionToken {

	//logs.Accessing("Session_ValidateLogin")
	//logs.Processing("appToken: " + appToken)
	//logs.Processing("username: " + username)
	//logs.Processing("password: " + password)

	var s sessionToken
	s.SecurityViolation = ""
	s.ResponseCode = ""
	s.UID = ""

	_, cred, _ := dao.Credentials_GetByUserName(username)
	if len(cred.Id) == 0 {
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION"
		logs.Warning("Username not found: " + username)
		return s
	}

	if cred.Username != username {
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION"
		logs.Warning("Username does not match: " + cred.Username)
		return s
	}
	if len(cred.Expiry) == 0 {
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION"
		logs.Warning("Expiry Date is empty: " + cred.Expiry)
		return s
	}

	now := time.Now()
	expiry, err := time.Parse(core.DATETIMEFORMATUSER, cred.Expiry)
	if err != nil {
		logs.Information("Unable to Parse Date: "+cred.Expiry+" trying another format", "")
		//logs.Warning(err.Error())
		// Try alternative format
		expiry, err = time.Parse("Monday, 2 Jan 2006 @ 15:04:05", cred.Expiry)
		if err != nil {
			//logs.Warning(err.Error())
			s.ResponseCode = "512"
			s.SecurityViolation = "SECURITY VIOLATION"
			logs.Warning("Unable to Parse Date: " + cred.Expiry)
			return s
		}
	}
	//fmt.Printf("now: %v\n", now)
	//fmt.Printf("expiry: %v\n", expiry)
	if now.After(expiry) {
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION"
		logs.Warning("Session Expired " + now.String() + " After " + expiry.String())
		return s
	}

	//TODO: insert password check here - start
	if cred.Password == "" && cred.PasswordExpiry == "" {
		s.ResponseCode = "303"
		s.SecurityViolation = "SET VIOLATION"
		logs.Warning("Password and Password Expiry are empty")
		activateSession(r, cred)
		return s
	}

	if cred.PasswordExpiry == "" {
		s.ResponseCode = "303"
		s.SecurityViolation = "RESET VIOLATION"
		logs.Warning("Password Expiry is empty")
		activateSession(r, cred)
		//Redirect to change password screen (need to set a special response code)
		return s
	}
	// Check Password Expiry
	pwdExpiry, err := time.Parse(core.DATEFORMAT, cred.PasswordExpiry)
	if pwdExpiry.Before(now) {
		s.ResponseCode = "303"
		s.SecurityViolation = "PASSWORD EXPIRED"
		logs.Warning("Password Expired " + now.String() + " After " + pwdExpiry.Format(core.DATEFORMAT))
		//Redirect to change password screen (need to set a special response code)
		activateSession(r, cred)
		return s
	}

	// Check Password Hash matches
	if cred.Password != "" {
		// Check Password Hash matches

		if !core.CheckPasswordHash(password, cred.Password) {
			s.ResponseCode = "512"
			s.SecurityViolation = "SECURITY VIOLATION"
			logs.Warning("Password does not match")
			return s
		}
	}

	//
	//TODO: insert password check here - end

	s.SecurityViolation = ""
	s.ResponseCode = "200"
	s.UID = cred.Id

	activateSession(r, cred)

	//fmt.Printf("cred.RoleType: %v\n", cred.RoleType)
	//fmt.Printf("core.GetNavigationID(cred.RoleType): %v\n", core.GetNavigationID(cred.RoleType))
	//fmt.Printf("cred.Knownas: %v\n", cred.Knownas)
	//fmt.Printf("cred.Username: %v\n", cred.Username)
	//fmt.Printf("core.ApplicationToken(): %v\n", core.ApplicationToken())
	//fmt.Printf("cred.Id: %v\n", cred.Id)
	////fmt.Printf("Session_CreateToken(r): %v\n", Session_CreateToken(r))

	//fmt.Printf("s: %v\n", s)
	//fmt.Printf("core.SessionManager: %v\n", core.SessionManager)
	//fmt.Printf("r.Context(): %v\n", r.Context())

	//fmt.Printf("core.SessionManager.Get(r.Context(), core.SessionUserName): %v\n", core.SessionManager.Get(r.Context(), core.SessionUserName))

	return s
}

func activateSession(r *http.Request, cred dm.Credentials) {
	core.SessionManager.Put(r.Context(), core.SessionRole, cred.RoleType)
	core.SessionManager.Put(r.Context(), core.SessionNavi, core.GetNavigationID(cred.RoleType))
	core.SessionManager.Put(r.Context(), core.SessionKnowAs, cred.Knownas)
	core.SessionManager.Put(r.Context(), core.SessionUserName, cred.Username)
	core.SessionManager.Put(r.Context(), core.SessionAppToken, core.ApplicationToken())
	core.SessionManager.Put(r.Context(), core.SessionUUID, cred.Id)
	core.SessionManager.Put(r.Context(), core.SessionSecurityViolation, "")
	core.SessionManager.Put(r.Context(), core.SessionTokenID, Session_CreateToken(r))
}

// Session_Validate is cheese
func Session_Validate(w http.ResponseWriter, r *http.Request) bool {
	var s sessionToken
	///TODO FIX Somewhere the context is getting lost!
	s.SecurityViolation = ""
	s.ResponseCode = ""
	//fmt.Printf("Session_GetUserName(r): %v\n", Session_GetUserName(r))
	//fmt.Printf("Session_GetUserSessionToken(r): %v\n", Session_GetUserSessionToken(r))
	session_uuid := Session_GetUserUUID(r)
	//fmt.Printf("session_uuid: %v\n", session_uuid)

	cookie_UserName := core.SessionManager.GetString(r.Context(), core.SessionUserName)
	cookie_UserUUID := core.SessionManager.GetString(r.Context(), core.SessionUUID)

	//fmt.Printf("cookie_UserName: %v\n", cookie_UserName)
	//fmt.Printf("cookie_UserUUID: %v\n", cookie_UserUUID)
	//fmt.Printf("session_uuid: %v\n", session_uuid)
	_, cred, err := dao.Credentials_GetByUUID(session_uuid)
	if err != nil {
		log.Panicf("ERROR %e", err)
	}

	if cookie_UserName == cred.Username {
		if cookie_UserName == "" {
			s.ResponseCode = "512"
			s.SecurityViolation = "SECURITY VIOLATION"
			logs.Warning("cookie_UserName & cred.Username are null")
			return false
		}
	}

	if cookie_UserName != cred.Username {
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION"
		logs.Warning("cookie_UserName != cred.Username")
		return false
	}

	if cookie_UserUUID != session_uuid {
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION"
		logs.Warning("cookie_UserUUID != session_uuid")
		return false
	}

	if cookie_UserUUID != cred.Id {
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION"
		logs.Warning("cookie_UserUUID != cred.Id")
		return false
	}

	if len(cred.Expiry) == 0 {
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION"
		logs.Warning("Session_Validate: Expiry is empty")
		return false
	}

	now := time.Now()
	expiry, err := time.Parse(core.DATETIMEFORMATUSER, cred.Expiry)
	if err != nil {
		logs.Information("Unable to Parse Date: "+cred.Expiry+" trying another format", "")
		// Try alternative format
		expiry, err = time.Parse("Monday, 2 Jan 2006 @ 15:04:05", cred.Expiry)
		if err != nil {
			//logs.Warning(err.Error())
			s.ResponseCode = "512"
			s.SecurityViolation = "SECURITY VIOLATION"
			return false
		}
	}

	if now.After(expiry) {
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION"
		logs.Warning("Session Expired " + now.String() + " After " + expiry.String())
		return false
	}

	// TODO: insert password check
	// TODO: insert server cred check

	s.SecurityViolation = ""
	s.ResponseCode = "200"

	return true
}

// getappMenuData
func Session_GetUserRole(r *http.Request) string {
	return core.SessionManager.GetString(r.Context(), core.SessionRole)
}

// getappMenuData
func Session_GetUserName(r *http.Request) string {
	return core.SessionManager.GetString(r.Context(), core.SessionUserName)
}

// getappMenuData
func Session_GetUserKnownAs(r *http.Request) string {
	return core.SessionManager.GetString(r.Context(), core.SessionKnowAs)
}

// getappMenuData
func Session_GetUserSessionToken(r *http.Request) string {
	return core.SessionManager.GetString(r.Context(), core.SessionTokenID)
}

// getappMenuData
func Session_GetUserUUID(r *http.Request) string {
	//logs.Accessing("Session_GetUserUUID")
	return core.SessionManager.GetString(r.Context(), core.SessionUUID)
}

func Session_CreateToken(req *http.Request) string {
	//logs.Accessing("Session_CreateToken")
	id := uuid.New().String()
	now := time.Now()
	//currentUserID, _ := user.Current()
	//userID := currentUserID.Name
	host, _ := os.Hostname()
	var r dm.Session
	r.Id = id
	r.Sessiontoken = id
	r.Apptoken = core.ApplicationToken()
	r.Createdate = now.Format(core.DATEFORMAT)
	r.Createtime = now.Format(core.TIMEHMS)
	r.Uniqueid = Session_GetUserUUID(req)
	r.Username = Session_GetUserName(req)
	r.Password = ""
	r.Userip = req.Referer()
	r.Userhost = core.GetIncomingRequestIP(req)
	r.Appip = core.GetHostIP()
	r.Apphost = host
	r.Issued = now.Format(core.DATETIMEFORMATUSER)

	//addTime, _ := strconv.Atoi(globals.core.ApplicationSessionLife())
	expiry := now.Add(time.Minute * 20)

	r.Expiry = expiry.Format(core.DATETIMEFORMATUSER)
	r.Expiryraw = expiry.String()
	r.SessionRole = Session_GetUserRole(req)
	r.Brand = ""
	r.SYSCreated = time.Now().Format(core.DATETIMEFORMATUSER)
	r.Expires = expiry.Format(core.DATETIMEFORMATSQLSERVER)

	dao.Session_Store(r, req)

	//fmt.Printf("id: %v\n", id)
	//fmt.Printf("r: %v\n", r)
	//fmt.Printf("id: %v\n", id)

	return id
}

func Session_GetSessionInfo(r *http.Request) (dm.SessionInfo, error) {
	//logs.Accessing("Session_GetSessionInfo")
	var s dm.SessionInfo
	s.UserName = Session_GetUserName(r)
	s.AppDB = core.ApplicationSQLDatabase()
	s.AppDBHost = core.ApplicationSQLServer()
	s.AppServerDate = time.Now().Format(core.DATEFORMAT)
	s.HostName = core.SystemHostname
	s.DateSyncIssue = ""
	s.Avatar = ""

	s.Type = "Primary"
	if core.IsChildInstance {
		s.Type = "Secondary"
	}
	s.AppName = core.ApplicationName()
	s.AppID = core.ReleaseIdentityVerbose()
	urMessages, _, _ := dao.Inbox_GetListByUser(s.UserName)

	s.UnreadMessages = ""
	if urMessages > 0 {
		s.UnreadMessages = strconv.Itoa(urMessages)
	}
	////fmt.Printf("s: %v\n", s)
	//spew.Dump(s)
	//fmt.Printf("s: %v\n", s)
	UID := Session_GetUserUUID(r)
	s.Avatar = "/data/images/avatars/" + UID + ".png"
	logs.Information("Acccesing", s.Avatar)

	return s, nil
}

type requestPage struct {
	AppName          string
	UserName         string
	UserPassword     string
	WebServerVersion string
	LicenceType      string
	LicenceLink      string
	ResponseMessage  string
	CompanyName      string
}

func Session_HandlerRegister(w http.ResponseWriter, r *http.Request) {
	tmpl := "Impl_Request"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	appName := core.ApplicationName()

	appServerVersion := core.ReleaseIdentityVerbose()

	loginPageContent := requestPage{
		AppName:          appName,
		UserName:         "",
		UserPassword:     "",
		WebServerVersion: appServerVersion,
		LicenceType:      core.ApplicationGetLicenseName(),
		LicenceLink:      core.ApplicationGetLicenseLink(),
		ResponseMessage:  "",
		CompanyName:      core.ApplicationCompanyName(),
	}

	message := core.GetURLparam(r, "msg")
	if message != "" {
		loginPageContent.ResponseMessage = message
	}
	////fmt.Println("Page Data", loginPageContent)

	//t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.SessionManager.GetString(r.Context(), core.SessionRole)))
	// Does not user ExecuteTemplate because this is a special case

	if message != "" {
		loginPageContent.ResponseMessage = message

		if message == "NEW" {
			loginPageContent.ResponseMessage = ""
		}
		//fmt.Printf("message: %v\n", message)
		ExecuteTemplate(tmpl, w, r, loginPageContent)

	} else {

		message = processRequest(r)
		//fmt.Printf("message: %v\n", message)
		message := html.EscapeString(message)
		if message != "" {
			http.Redirect(w, r, "/request?msg="+message, http.StatusFound)
		} else {
			http.Redirect(w, r, "/reqcomplete", http.StatusFound)
		}
	}

}

func processRequest(r *http.Request) string {
	//fmt.Printf("r.ParseForm(): %v\n", r.ParseForm())
	//fmt.Println("Process Registration Request")
	firstName := r.FormValue("firstname")
	lastName := r.FormValue("lastname")
	email := r.FormValue("email")
	username := email
	password := r.FormValue("password")
	passwordConfirm := r.FormValue("passwordconfirm")

	//fmt.Printf("firstName: %v\n", firstName)
	//fmt.Printf("lastName: %v\n", lastName)
	//fmt.Printf("email: %v\n", email)
	//fmt.Printf("username: %v\n", username)
	//fmt.Printf("password: %v\n", password)
	//fmt.Printf("passwordConfirm: %v\n", passwordConfirm)

	if password != passwordConfirm {
		return "Password and Confirm Password do not match"
	}

	if len(password) < 6 {
		return "Password must be at least 6 characters long"
	}

	//TODO Duplicate Check
	err := Credentials_DuplicateCheck(email)
	if err != nil {
		return err.Error()
	}

	err = Credentials_NewRequest(firstName, lastName, email, username, password)
	if err != nil {
		return err.Error()
	}

	adminID := core.GetApplicationProperty("admin")

	msgTXT := "A registration request has been received from %s %s for %s. Please review the request and approve or reject it."
	msgTXT = dao.Translate("Message", msgTXT)

	mailMessage := fmt.Sprintf(msgTXT, firstName, lastName, username)
	err = Inbox_SendMailSystem(adminID, "Registration Request", mailMessage)
	if err != nil {
		return err.Error()
	}

	return ""
}
