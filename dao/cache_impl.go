package dao

import (
	core "github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
	"golang.org/x/exp/slices"
)

var validObjects = []string{dm.Origin_Name, dm.Project_Name, dm.EstimationSession_Name, dm.Data_Name, dm.Profile_Name}

func CacheRead(object string, id string) interface{} {
	//logs.Information("CacheRead", "CacheRead: "+object+core.ID_SEP+id)
	//name := core.ApplicationProperties.Get("appname")
	//logs.Information("CacheRead", name)
	//logs.Warning("CacheRead:" + object + core.ID_SEP + id)
	//logs.Information("CacheRead", core.ApplicationCache.Stat())
	cacheID := core.ApplicationCache.KeyGen(object, id)
	//logs.Information("CacheID", cacheID)
	// Check for valid object
	if !slices.Contains(validObjects, object) {
		logs.Warning("Invalid Cache object type: " + object)
		panic("Invalid Cache object type: " + object)
	}
	//logs.Information("CacheRead", "Valid object: "+object)
	if core.ApplicationCache.OK(cacheID) {
		//logs.Warning("CacheRead : Cache hit: " + cacheID)
		//logs.Information("Information : ", core.ApplicationCache.Stat())
		return core.ApplicationCache.Get(cacheID)
	}
	//logs.Information("CacheRead", "Cache Not Found/Not OK: "+cacheID)
	cacheItem := cache_Get_From_DB(object, id)
	//logs.Warning("Loading Cache: " + cacheID)
	//logs.Information("CacheRead", "Adding to cache: "+cacheID)
	core.ApplicationCache.Add(cacheID, cacheItem)
	//logs.Information("Information : ", core.ApplicationCache.Stat())
	return cacheItem
}

func cache_Get_From_DB(object string, id string) interface{} {
	//logs.Warning("cache_Get_From_DB " + object + " " + id)
	switch object {
	case dm.Origin_Name:
		//logs.Information("cache_Get_From_DB", "Origin_GetByCode: "+id)
		_, origin, err := Origin_GetByCode(id)
		//fmt.Printf("no: %v\n", no)
		//fmt.Printf("origin: %v\n", origin)
		//fmt.Printf("err: %v\n", err)
		if err != nil {
			panic("Origin not found: " + id + err.Error())
		}
		//fmt.Printf("origin: %v\n", origin)
		return origin
	case dm.Project_Name:
		_, project, _ := Project_GetByID(id)
		return project
	case dm.EstimationSession_Name:
		_, session, _ := EstimationSession_GetByID(id)
		return session
	case dm.Data_Name:
		_, data, _ := Data_GetByID(id)
		return data
	case dm.Profile_Name:
		_, profile, _ := Profile_GetByCode(id)
		return profile
	default:
		panic("Invalid Cache object type: " + object)
	}
	return nil
}
