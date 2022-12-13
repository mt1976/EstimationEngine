package core

import (
	"crypto/tls"
	"fmt"
	"strconv"
	"time"

	"gopkg.in/gomail.v2"
)

var EMAIL_From string
var EMAIL_Footer string

const (
	ES_FROM    = "From"
	ES_TO      = "To"
	ES_CC      = "Cc"
	ES_SUBJECT = "Subject"
	ES_TYPE    = "text/html"
)

func Email_init() *gomail.Dialer {
	//fmt.Println("Email Init")

	emailService := GetApplicationProperty("emailservice")
	emailPort, _ := strconv.Atoi(GetApplicationProperty("emailPort"))
	emailUser := GetApplicationProperty("emailUser")
	emailPassword := GetApplicationProperty("emailPassword")
	EMAIL_From = GetApplicationProperty("emailFrom")
	EMAIL_Footer = GetApplicationProperty("emailFooter")
	if EMAIL_Footer == "" {
		EMAIL_Footer = "This is an automated email. Please do not reply!"
	}
	EMAIL_Footer = EMAIL_Footer + "<br><br>--"
	EMAIL_Footer = EMAIL_Footer + fmt.Sprintf("<br><br><small><i>Generated by: %s (%s) on %s</i></small>", ApplicationName(), ApplicationHostname(), time.Now().Format(DATEMSG))

	Emailer = gomail.NewDialer(emailService, emailPort, emailUser, emailPassword)
	Emailer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return Emailer
}

func SendEmail(to string, name string, subject string, body string) {
	m := gomail.NewMessage()
	m.SetHeader(ES_FROM, EMAIL_From)
	m.SetHeader(ES_TO, to)
	m.SetAddressHeader(ES_CC, GetApplicationProperty("admin"), "Admin")
	m.SetHeader(ES_SUBJECT, subject)

	MSG_BODY := name + ",<br><br>" + body + "<br><br>" + EMAIL_Footer
	m.SetBody(ES_TYPE, MSG_BODY)
	//m.Attach("/home/Alex/lolcat.jpg")

	// Send the email to Bob, Cora and Dan.
	if err := Emailer.DialAndSend(m); err != nil {
		panic(err)
	}
}
