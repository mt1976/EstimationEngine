package dao

import (
	core "github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
	"golang.org/x/exp/slices"
)

var validObjects = []string{dm.Origin_Name, dm.Project_Name, dm.EstimationSession_Name}

func CacheRead(object string, id string) interface{} {
	cacheID := core.ApplicationCache.KeyGen(object, id)

	// Check for valid object
	if !slices.Contains(validObjects, object) {
		panic("Invalid Cache object type: " + object)
	}

	if core.ApplicationCache.OK(cacheID) {
		return core.ApplicationCache.Get(cacheID)
	}
	cacheItem := cache_Get_From_DB(object, id)
	core.ApplicationCache.Add(cacheID, cacheItem)
	logs.Information("Information : ", core.ApplicationCache.Stat())
	return cacheItem
}

func cache_Get_From_DB(object string, id string) interface{} {

	switch object {
	case dm.Origin_Name:
		_, origin, _ := Origin_GetByCode(id)
		return origin
	case dm.Project_Name:
		_, project, _ := Project_GetByID(id)
		return project
	case dm.EstimationSession_Name:
		_, session, _ := EstimationSession_GetByID(id)
		return session
	default:
		panic("Invalid Cache object type: " + object)
	}
	return nil
}
