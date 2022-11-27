package application

import (
	"time"

	core "github.com/mt1976/ebEstimates/core"
	dao "github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

const (
	Expired = "EXPIRED"
	ACK     = "ACK"
	NACK    = "NACK"
)

func ExternalMessage_Expired(msg dm.ExternalMessage) error {
	msg.MessageACKNAK = Expired

	message := "No Response from External System : " + msg.MessageID
	core.Notification_Normal("External Messaging", message)

	dao.ExternalMessage_StoreSystem(msg)

	return nil
}

func ExternalMessage_Notify(msg dm.ExternalMessage) error {

	logs.Information("message", msg.MessageID)

	return nil
}

func ExternalMessage_ACK(id string, inFile string, body string, additionalInfo string, inTime time.Time) error {

	logs.Information("message", id)

	_, msg, err := dao.ExternalMessage_GetByID(id)

	if err != nil {
		// Cannot find a message to link to.
		var msg dm.ExternalMessage
		msg.MessageID = id
	}

	msg.MessageACKNAK = ACK
	msg = externalMessage_Recv(msg, inFile, body, additionalInfo, inTime)

	dao.ExternalMessage_StoreSystem(msg)

	return nil
}

func externalMessage_Recv(msg dm.ExternalMessage, inFile string, body string, additionalInfo string, inTime time.Time) dm.ExternalMessage {
	now := time.Now()
	msg.ResponseDate = now.Format(core.DATEFORMAT)
	msg.ResponseTime = now.Format(core.TIMEFORMAT)
	msg.ResponseFilename = inFile
	msg.ResponseRecieved = inTime.Format(core.DATETIME)
	msg.ResponseBody = body
	msg.ResponseAdditionalInfo = additionalInfo
	return msg
}

func ExternalMessage_NACK(id string, inFile string, body string, additionalInfo string, inTime time.Time) error {
	logs.Information("message", id)

	_, msg, err := dao.ExternalMessage_GetByID(id)

	if err != nil {
		// Cannot find a message to link to.
		var msg dm.ExternalMessage
		msg.MessageID = id
	}

	msg.MessageACKNAK = NACK
	msg = externalMessage_Recv(msg, inFile, body, additionalInfo, inTime)

	dao.ExternalMessage_StoreSystem(msg)

	return nil
}
