package dao

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	core "github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

const GET = "GET"
const PUT = "PUT"
const VAL = "VALIDATE"
const NEW = "NEW"

type SQLData struct {
	fields []string
	values []string
}

func addData(d SQLData, f string, v string) SQLData {
	if f != "_id" {
		d.fields = append(d.fields, f)
		d.values = append(d.values, sq(v))
	}
	return d
}

func fields(ts SQLData) string {
	return strings.Join(ts.fields[:], ",")
}

func values(ts SQLData) string {
	return strings.Join(ts.values[:], ",")
}

func sq(in string) string {
	return "'" + in + "'"
}

func get_TableName(in string, in2 string) string {
	return in + "." + in2
}

// Gets a string from the interface map & default a value if not found
func get_String(t map[string]interface{}, v string, d string) string {
	i := t[v]
	//fmt.Printf("%q: %q %q\n", v, i, d)
	if i == nil {
		return d
	}
	return i.(string)
}

// Gets a time.Time from the interface map & default a value if not found
func get_Time(t map[string]interface{}, v string, d string) string {
	i := t[v]

	//fmt.Printf("%q: %q %q\n", v, i, d)
	if i == nil {
		return d
	}
	return core.TimeToString(i.(time.Time))
}

// Convert time.Time to string
func get_Int(t map[string]interface{}, v string, d string) string {
	i := t[v]
	//fmt.Printf("%q: %q %q\n", v, i, d)

	if i != nil {
		return fmt.Sprintf("%d", i.(int64))
	}
	return d
}

// Convert float64 to string
func get_Float(t map[string]interface{}, v string, d string) string {
	i := t[v]
	//fmt.Printf("%q: %q %q\n", v, i, d)
	if i != nil {
		return fmt.Sprintf("%f", i.(float64))
	}
	return d
}

// Convert bool to string
func get_Bool(t map[string]interface{}, v string, d string) string {
	i := t[v]
	//fmt.Printf("%q: %v %q\n", v, i, d)
	if i != nil {
		return fmt.Sprintf("%v", i.(bool))
	}
	return d
}

func Audit_Update(val string, replace string) string {
	//	fmt.Printf("val: %v\n", val)
	//	fmt.Printf("replace: %v\n", replace)
	if val == "" {
		return replace
	}
	return val
}

func Audit_User(r *http.Request) string {
	//thisUser, _ := user.Current()
	thisUser := Session_GetUserName(r)
	return thisUser
}

func Audit_TimeStamp() string {
	return time.Now().Format(core.DATETIMEFORMATUSER)
}

func Audit_Host() string {
	host, _ := os.Hostname()
	return host
}

func tsql_AND(tsql string, inField string, inValue string) string {
	tsql = tsql + " AND " + inField + " = " + sq(inValue)
	return tsql
}

func StubLists_Get(listName string) []dm.Lookup_Item {
	listid := strings.ToLower(listName)
	ctList := core.ApplicationStubLists[listid]
	//logs.Success("listName=" + listName)
	//logs.Success("List=" + ctList)
	// Turn comma-separated string into a slice of strings
	ctSlice := strings.Split(ctList, core.LIST_SEP)
	//fmt.Printf("ctSlice: %v\n", ctSlice)
	// Create a lookup list of all CounterpartyType items
	ctLookup := make([]dm.Lookup_Item, len(ctSlice))
	for i := 0; i < len(ctSlice); i++ {

		listID := ctSlice[i]
		listItem := ctSlice[i]
		if strings.Contains(listID, core.LIST_ID_SEP) {
			ctExploded := strings.Split(ctSlice[i], core.LIST_ID_SEP)
			listID = ctExploded[0]
			listItem = ctExploded[1]
		}

		ctLookup[i] = dm.Lookup_Item{ID: listID, Name: listItem}
	}
	//fmt.Printf("ctLookup: %v\n", ctLookup)
	return ctLookup
}

func getNewFileID() string {
	nID := uuid.New().String()
	nID = core.ReplaceSpecialChars(nID)
	nID = strings.ReplaceAll(nID, "-", "")
	return nID
}

func SendMailToResource(resourceID string, MSG_SUBJECT string, MSG_BODY string) {
	_, who, _ := Resource_GetByCode(resourceID)
	if resourceID == "" {
		logs.Warning("SendMail: No resourceID - not sending email")
		return
	}
	if MSG_BODY == "" || MSG_SUBJECT == "" {
		logs.Warning("SendMail: No message body or subject - not sending email")
		return
	}
	if who.Email == "" {
		logs.Warning("SendMail: No email address for " + core.DQuote(who.Name) + " " + core.DQuote(resourceID) + " - not sending email")
		return
	}
	if who.UserActive == core.FALSE {
		logs.Warning("SendMail: User " + core.DQuote(who.Name) + " " + core.DQuote(resourceID) + " is not active - not sending email")
		return
	}
	core.SendEmail(who.Email, who.Name, MSG_SUBJECT, MSG_BODY)
}

func SetFieldError(fP dm.FieldProperties, msg string) dm.FieldProperties {
	fP.MsgType = "is-invalid"
	if len(fP.MsgMessage) > 0 {
		fP.MsgMessage = fP.MsgMessage + ", " + msg
	} else {
		fP.MsgMessage = msg
	}
	fP.MsgFeedBackType = "invalid-feedback"
	return fP
}

func SetFieldMandatory(fP dm.FieldProperties, msg string) dm.FieldProperties {
	fP.MsgType = "is-invalid"
	if len(fP.MsgMessage) > 0 {
		fP.MsgMessage = fP.MsgMessage + ", " + msg
	} else {
		fP.MsgMessage = msg
	}
	fP.MsgFeedBackType = "invalid-feedback"
	return fP
}

func SetFieldGood(fP dm.FieldProperties, msg string) dm.FieldProperties {
	fP.MsgType = "is-invalid"
	if len(fP.MsgMessage) > 0 {
		fP.MsgMessage = fP.MsgMessage + ", " + msg
	} else {
		fP.MsgMessage = msg
	}
	fP.MsgFeedBackType = "invalid-feedback"
	return fP
}
