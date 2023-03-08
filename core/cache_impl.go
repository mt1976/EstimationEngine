package core

import (
	"fmt"
	"time"

	"github.com/mt1976/ebEstimates/logs"
)

type Cache struct {
	content map[string]interface{}
	expiry  map[string]time.Time
	maxsize int
	payload int
	life    int
}
type CacheInterface interface {
	Add(key string, val string)
	OK(key string) bool
	Get(key string) interface{}
	KeyGen(object string, id string) string
	KeyExists(key string) bool
	Stat() string
}

var cache Cache

// var cache map[string]interface{}
// var cache_expiry map[string]time.Time
// var cache_maxsize int
// var cache_payload int
// var cache_life int

func (c Cache) Add(key string, val interface{}) {
	if c.payload >= c.maxsize {
		c.init()
	}
	if !c.KeyExists(key) {
		c.payload++
		c.expiry[key] = time.Now().Add(time.Duration(c.life) * time.Second)
	}
	c.content[key] = val
}

func (c Cache) OK(key string) bool {
	if !c.KeyExists(key) {
		return false
	}
	if time.Now().Before(c.expiry[key]) {
		return false
	}
	return true
}

func (c Cache) Get(key string) interface{} {
	return cache.content[key]
}

func (c Cache) KeyGen(object string, id string) string {
	return object + ID_SEP + id
}

func (c Cache) KeyExists(key string) bool {
	_, ok := c.content[key]
	return ok
}

func (c Cache) init() {
	logs.Information("Cache", "Initialising")
	cache := Cache{}
	cache.content = make(map[string]interface{})
	cache.expiry = make(map[string]time.Time)

	cache.payload = 0

	cm_val := StringToInt(GetApplicationProperty("cache_maxsize"))
	if cm_val > 0 {
		cache.maxsize = cm_val
	} else {
		cache.maxsize = 1000
	}

	cl_val := StringToInt(GetApplicationProperty("cache_life"))
	if cl_val > 0 {
		cache.life = cl_val
	} else {
		cache.life = 60
	}
	//fmt.Printf("cache: %v\n", cache)
	logs.Information("Cache", fmt.Sprintf("Initialised - Size: %v Max: %v Life: %v", cache.payload, cache.maxsize, cache.life))
}

func (c Cache) Stat() string {
	return fmt.Sprintf("Cache Status: %v/%v %v", cache.payload, cache.maxsize, cache.life)
}
