package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/externalmessage.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : ExternalMessage (externalmessage)
// Endpoint 	        : ExternalMessage (MessageID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/12/2021 at 09:31:50
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	core "github.com/mt1976/ebEstimates/core"
	logs "github.com/mt1976/ebEstimates/logs"

	dm "github.com/mt1976/ebEstimates/datamodel"
)

const (
	Message_FormatXML            = "XML"
	Message_TimeoutAction_Notify = "Notify"
)

// ExternalMessage_GetByID() returns a single ExternalMessage record
func ExternalMessage_GetActive() (int, []dm.ExternalMessage, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.ExternalMessage_SQLTable)
	tsql = tsql + " WHERE " + dm.ExternalMessage_ResponseDate_sql + "= ''"
	tsql = tsql + " AND " + dm.ExternalMessage_MessageACKNAK_sql + "= ''"
	noItems, externalmessageList, _, _ := externalmessage_Fetch(tsql)

	return noItems, externalmessageList, nil
}

func ExternalMessage_Sent(id string, format string, msgClass string, destination string, body []byte, filename string, lifeSpan int, timeoutAction string) error {
	//logs.Success("Message Sent:" + id)
	var msItem dm.ExternalMessage
	msItem.MessageID = id
	msItem.MessageFormat = format
	msItem.MessageClass = msgClass
	msItem.MessageDeliveredTo = destination
	msItem.MessageBody = string(body)
	msItem.MessageFilename = filename
	msItem.MessageLife = strconv.Itoa(lifeSpan)
	now := time.Now()
	msItem.MessageEmitted = now.Format(core.DATETIME)
	msItem.MessageDate = now.Format(core.DATEFORMAT)
	msItem.MessageTime = now.Format(core.TIMEFORMAT)
	msItem.MessageTimeout = now.Add(time.Duration(lifeSpan) * time.Second).Format(core.DATETIME)
	msItem.MessageTimeoutAction = timeoutAction
	msItem.AppID = core.ApplicationToken()
	var err error
	//fmt.Printf("msItem: %v\n", msItem)
	//spew.Dump(msItem)
	//err = dao.ExternalMessage_StoreSystem(msItem)

	json_data, err := json.Marshal(msItem)
	//fmt.Println(string(json_data))
	if err != nil {
		return err
	}

	// get current ip address
	ip := "localhost"
	uri := core.ApplicationHTTPProtocol() + "://" + ip + ":" + core.ApplicationHTTPPort() + dm.ExternalMessage_Path

	//logs.Information("uri", uri)

	resp, err := http.Post(uri, "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	//var res map[string]interface{}
	//logs.Information("response Status:", resp.Status)
	//json.NewDecoder(resp.Body).Decode(&res)
	if resp.StatusCode != 200 {
		logs.Error("Error: "+strconv.Itoa(resp.StatusCode), nil)
		return err
	}
	if resp.StatusCode == 404 {
		logs.Warning("Unable to Store Message:" + id)
	}

	logs.Post(uri + id)
	// logs.Accessing(uri + id)

	// params := url.Values{}
	// params.Add(dm.ExternalMessage_QueryString, id)

	// //	fmt.Println(res["json"])
	// gesp, err2 := http.Get(uri + "?" + params.Encode())

	// if err2 != nil {
	// 	log.Fatal(err2)
	// }

	// defer gesp.Body.Close()

	// gbody, err2 := ioutil.ReadAll(gesp.Body)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// //logs.Information("response Status:", gesp.Status)

	// //spew.Dump(gesp)
	// //json.NewDecoder(resp.Body).Decode(&res)
	// if gesp.StatusCode != 200 {
	// 	logs.Error("Error: "+strconv.Itoa(gesp.StatusCode), nil)
	// 	return err
	// }
	// if gesp.StatusCode == 404 {
	// 	logs.Warning("Unable to Store Message:" + id)
	// }

	// //fmt.Printf("gesp.Body: %v\n", gesp.Body)
	// //fmt.Println("RESPONSE>>" + string(gbody))

	// var data *dm.ExternalMessage
	// json.Unmarshal(gbody, &data)

	//json.NewDecoder(gesp.Body).Decode(data)

	//fmt.Printf("Results: %v\n", data)
	//spew.Dump(data)
	return nil
}
