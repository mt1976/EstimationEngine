package adaptor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

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
