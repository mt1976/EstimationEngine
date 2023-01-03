package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/inbox.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Inbox (inbox)
// Endpoint 	        : Inbox (Message)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/06/2022 at 17:00:15
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	core "github.com/mt1976/ebEstimates/core"
	logs "github.com/mt1976/ebEstimates/logs"

	dm "github.com/mt1976/ebEstimates/datamodel"
)

// Inbox_GetList() returns a list of all Inbox records
func Inbox_GetListByUser(inUser string) (int, []dm.Inbox, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.Inbox_SQLTable)
	tsql = tsql + " WHERE " + dm.Inbox_MailTo_sql + "='" + inUser + "'"

	count, inboxList, _, _ := inbox_Fetch(tsql)

	return count, inboxList, nil
}

// Inbox_GetList() returns a list of all Inbox records
func Inbox_GetUnreadByUser(inUser string) (int, []dm.Inbox, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.Inbox_SQLTable)
	tsql = tsql + " WHERE " + dm.Inbox_MailTo_sql + "='" + inUser + "'"
	tsql = tsql + " AND " + dm.Inbox_MailUnread_sql + "='" + "true" + "'"

	count, inboxList, _, _ := inbox_Fetch(tsql)

	return count, inboxList, nil
}

const (
	Inbox_FormatXML            = "XML"
	Inbox_TimeoutAction_Notify = "Notify"
)

func Inbox_SendMail(mailTo string, mailSubject string, mailContent string, mailFrom string, source string) error {
	//logs.Success("Message Sent:" + id)
	fmt.Printf("mailTo: %v\n", mailTo)
	fmt.Printf("mailSubject: %v\n", mailSubject)
	fmt.Printf("mailContent: %v\n", mailContent)
	fmt.Printf("mailFrom: %v\n", mailFrom)
	fmt.Printf("source: %v\n", source)

	now := time.Now()

	var mail dm.Inbox
	mail.MailId = getNewFileID()
	mail.MailTo = mailTo
	mail.MailSubject = mailSubject
	mail.MailContent = mailContent
	mail.MailFrom = mailFrom
	mail.MailSent = now.Format(core.DATETIME)
	mail.MailSource = source

	var err error
	//fmt.Printf("msItem: %v\n", mail)
	//spew.Dump(mail)

	//err = dao.ExternalInbox_StoreSystem(msItem)

	json_data, err := json.Marshal(mail)
	fmt.Println(string(json_data))
	if err != nil {
		return err
	}

	// get current ip address
	ip := "localhost"
	uri := core.ApplicationHTTPProtocol() + "://" + ip + ":" + core.ApplicationHTTPPort() + dm.Inbox_Path

	logs.Information("uri", uri)

	resp, err := http.Post(uri, "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}
	logs.Information("response Status:", resp.Status)
	json.NewDecoder(resp.Body).Decode(&res)
	if resp.StatusCode != 200 {
		logs.Error("Error: "+strconv.Itoa(resp.StatusCode), nil)
		return err
	}
	if resp.StatusCode == 404 {
		logs.Warning("Unable to Store Message:" + mail.MailId)
	}

	logs.Post(uri + mail.MailId)
	return nil
}
