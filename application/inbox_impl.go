package application

// ----------------------------------------------------------------
// Automatically generated  "/application/inbox.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Inbox (inbox)
// Endpoint 	        : Inbox (Message)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/06/2022 at 17:00:15
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	core "github.com/mt1976/ebEstimates/core"
	dao "github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// Inbox_Publish annouces the endpoints available for this object
func Inbox_Publish_Impl(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.Inbox_PathMyList, Inbox_HandlerMyInboxList)
	logs.Publish("Application", dm.Inbox_Title+"_Impl")
	//No API
}

// Inbox_HandlerList is the handler for the list page
func Inbox_HandlerMyInboxList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Inbox
	noItems, returnList, _ := dao.Inbox_GetListByUser(Session_GetUserName(r))
	logs.Processing(Session_GetUserName(r))

	pageDetail := dm.Inbox_PageList{
		Title:       CardTitle(dm.Inbox_Title, core.Action_List),
		PageTitle:   PageTitle(dm.Inbox_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:    returnList,
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Inbox_TemplateList, w, r, pageDetail)

}

//TODO new func SendMail - Send mail to 1 user
//TODO new func SendBroadcastMail - Send mail to a group

func Inbox_SendMailSystem(mailto string, mailSubject string, mailContent string) error {
	var mail dm.Inbox
	mail.MailTo = mailto
	mail.MailFrom = dao.Audit_Host()
	mail.MailId = uuid.New().String()
	mail.MailSubject = mailSubject
	mail.MailContent = mailContent
	mail.MailSent = time.Now().Format(core.DATETIME)

	_, err := dao.Inbox_StoreSystem(mail)
	return err
}
