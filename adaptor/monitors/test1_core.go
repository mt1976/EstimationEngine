package monitors

import (
	core "github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// Simulator_Test1_Job defines the job properties, name, period etc..
func Simulator_Test1_Job() dm.JobDefinition {
	// ----------------------------------------------------------------
	// Job Definition
	// ----------------------------------------------------------------
	var j dm.JobDefinition
	j.ID = "Simulator_Test1"
	j.Name = "Simulator_Test1"
	j.Period = ""
	j.Description = "Simulator Test1 Monitor"
	j.Type = core.Monitor
	return j
}

func Simulator_Test1_Watch() {
	logs.Success("Simulator_Test1_Start")

	// loc := core.AppProperties["swift_recv"]
	// // if first char of loc is . then remove it
	// if loc[0] == '.' {
	// 	loc = loc[1:]
	// }
	// loc = core.PWD + loc

	// logs.Information("Watching ", loc)

	// watcher, err := fsnotify.NewWatcher()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer watcher.Close()

	// application.Schedule_Register(Simulator_Test1_Job())

	// done := make(chan bool)
	// go func() {
	// 	for {
	// 		select {
	// 		case event, ok := <-watcher.Events:
	// 			if !ok {
	// 				return
	// 			}

	// 			if event.Op&fsnotify.Create == fsnotify.Create {
	// 				logs.Event(event.Name)
	// 				// Create a Test1_Simulator_ProcessResponse_impl job in  adaptor/monitor/Simulator_Test1_Impl.go
	// 				err := adaptor.Test1_Simulator_ProcessResponse_impl(event.Name)
	// 				if err != nil {
	// 					logs.Error("Test1", err)
	// 				}
	// 			}
	// 		case err, ok := <-watcher.Errors:
	// 			if !ok {
	// 				return
	// 			}
	// 			log.Println("error:", err)
	// 		}
	// 	}
	// }()

	// err = watcher.Add(loc)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// <-done

}
