package jobs

import (
	"fmt"
	"runtime"
	"strings"

	core "github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/dao"
	"github.com/mt1976/ebEstimates/logs"

	cron "github.com/robfig/cron/v3"
)

//CONST_CONFIG_FILE is cheese

// Start Initialises the Required Jobs
func Start() {
	//funcName = "Start"
	//logit("JobHandler", "CRON SCHEDULER")
	// Verbose c := cron.New(cron.WithLogger(
	//	cron.VerbosePrintfLogger(log.New(os.Stdout, "", log.LstdFlags))))
	c := cron.New()

	// Insert Jobs Here
	//var period string
	//var runType string

	//logit("info", c.Location().String())

	//logit("cron Locale", c.Location().String())

	if !core.IsChildInstance {
		logs.Information("Scheduler", "Clearing Schedule")
		_, items, _ := dao.Schedule_GetList()
		for _, item := range items {
			dao.Schedule_HardDelete(item.Id)
		}
		//		RatesFXSpot_Register(c)
		Origin_Register(c)
		EstimationSession_Register(c)
		Project_Register(c)
		ssloader_Register(c)
		Index_Register(c)
		Data_Register(c)
		SessionHouseKeeping_Register(c)
		Credentials_Register(c)
		if core.ApplicationEnvironment() != "production" {
			ClearFiles_Register(c)
		}
	}
	//ExternalMessage_Register(c)
	HeartBeat_Register(c)
	Database_Register(c)

	//DataDispatcher_Register(c, "MARKET", "*/10 6-21 * * 1-5")

	c.Start()

}

func logit(actionType string, data string) {
	_, caller, _, _ := runtime.Caller(1)
	outcall := strings.Split(caller, "/")
	depth := len(outcall) - 1
	depth2 := depth - 1
	//log.Println(len(outcall), depth, depth2)
	callerName := outcall[depth2] + "/" + outcall[depth]
	op := fmt.Sprintf("%v '%v' {%v}", actionType, data, callerName)
	logs.Message("Scheduler", op)
}
