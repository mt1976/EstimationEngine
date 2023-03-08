package dao

import (
	core "github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
	"golang.org/x/exp/slices"
)

var validObjects = []string{dm.Origin_Name, dm.Project_Name, dm.EstimationSession_Name}

func CACHE_Read(object string, id string) interface{} {
	cacheID := core.CACHE_KeyGen(object, id)

	// Check for valid object
	if !slices.Contains(validObjects, object) {
		panic("Invalid Cache object type: " + object)
	}

	if !core.CACHE_KeyExists(cacheID) {
		// Branch to appropriate DAO
		// Read from DB
		cacheItem := cache_Get_From_DB(object, id)
		core.CACHE_Add(cacheID, cacheItem)
		return cacheItem
	}
	return core.CACHE_Get(cacheID)
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
