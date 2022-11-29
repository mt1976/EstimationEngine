package application

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	core "github.com/mt1976/ebEstimates/core"
	dao "github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
)

func credentials_PublishImpl(mux http.ServeMux) http.ServeMux {
	return mux
}

func Credentials_HandlerBan(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := core.GetURLparam(r, dm.Credentials_QueryString)
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	core.ServiceMessageAction(inUTL, "Ban", searchID)
	banCredentialsStore(searchID, r)
	http.Redirect(w, r, dm.Credentials_Redirect, http.StatusFound)
}

func Credentials_HandlerActivate(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := core.GetURLparam(r, dm.Credentials_QueryString)
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	core.ServiceMessageAction(inUTL, "Activate", searchID)
	activateCredentialsStore(searchID, r)
	http.Redirect(w, r, dm.Credentials_Redirect, http.StatusFound)

}

func getExpiryDate() string {
	expiryDate := time.Now()
	life, _ := strconv.Atoi(core.ApplicationCredentialsLife())
	expiryDate = expiryDate.AddDate(0, 0, life)
	return expiryDate.Format(core.DATETIMEFORMATUSER)
}

func banCredentialsStore(id string, req *http.Request) {

	_, r, _ := dao.Credentials_GetByID(id)
	r.Expiry = ""
	r.Password = ""
	dao.Credentials_Store(r, req)
}

func activateCredentialsStore(id string, req *http.Request) {

	_, r, _ := dao.Credentials_GetByID(id)
	r.Expiry = getExpiryDate()
	dao.Credentials_Store(r, req)
}

func credentials_HandlerViewImpl(pageDetail dm.Credentials_Page) dm.Credentials_Page {
	return pageDetail
}

func credentials_HandlerEditImpl(pageDetail dm.Credentials_Page) dm.Credentials_Page {
	return pageDetail
}

func credentials_HandlerSaveImpl(item dm.Credentials) dm.Credentials {
	if item.Issued == "" {
		item.Issued = time.Now().Format(core.DATETIMEFORMATUSER)
	}
	if item.Expiry == "" {
		item.Expiry = getExpiryDate()
	}
	return item
}

// TODO Build a new request function. Check for duplicates etc.
func Credentials_NewRequest(firstName string, lastName string, email string, username string, password string) error {

	var rec dm.Credentials
	rec.Firstname = firstName
	rec.Lastname = lastName
	rec.Email = email
	rec.Username = username
	rec.State = "REQD"
	rec.Password = core.EncodeString(password)
	//fmt.Printf("rec: %v\n", rec)
	//spew.Dump(rec)

	return dao.Credentials_StoreSystem(rec)
}

// TODO Build a new request function. Check for duplicates etc.
func Credentials_RejectRequest(userID string) error {
	err, _ := credentials_UpdateStatus(userID, "REJD")
	return err
}

// TODO Build a new request function. Check for duplicates etc.
func Credentials_ApproveRequest(userID string) error {
	err, rec := credentials_UpdateStatus(userID, "APPR")
	if err != nil {
		return err
	}
	err = Inbox_SendMailSystem(rec.Email, "Welcome to the system", "emailRequestApproved")
	if err != nil {
		return err
	}
	return nil
}

// TODO Build a new request function. Check for duplicates etc.
func credentials_UpdateStatus(userID string, state string) (error, dm.Credentials) {

	//var rec dm.Credentials
	_, rec, err := dao.Credentials_GetByID(userID)
	if err != nil {
		return err, rec
	}
	rec.State = state
	//fmt.Printf("rec: %v\n", rec)
	//spew.Dump(rec)
	dao.Credentials_StoreSystem(rec)

	return nil, rec
}

func Credentials_DuplicateCheck(email string) error {

	noItems, _, err := dao.Credentials_GetByEmail(email)
	if err != nil {
		return err
	}
	if noItems > 0 {
		return fmt.Errorf("Email already exists")
	}
	return nil
}
