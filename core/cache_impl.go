package core

import (
	"fmt"
	"time"

	logs "github.com/mt1976/ebEstimates/logs"
)

type Cache struct {
	content map[string]interface{}
	used    map[string]bool
	expiry  map[string]time.Time
	keys    map[string]string
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
	Start()
	CheckExpiry(key string) bool
	GetKeys() []string
	Housekeep()
}

// var cache Cache
func (c Cache) Housekeep() Cache {
	//logs.Warning("Housekeep " + strconv.Itoa(len(c.keys)))
	for _, v := range c.keys {
		//fmt.Printf("i: %v\n", i)
		//fmt.Printf("v: %v\n", v)
		c.ExpiryProcessing(v)
	}
	//logs.Warning("Housekeep " + strconv.Itoa(len(c.keys)) + " complete")
	logs.Information("Cache", c.Stat())
	return c
}

func (c Cache) Add(key string, val interface{}) Cache {
	//logs.Information("Add", key)
	//fmt.Printf("c.payload: %v\n", c.payload)
	//fmt.Printf("c.maxsize: %v\n", c.maxsize)
	//logs.Warning("Cache Add: " + key + " " + c.Stat())
	if len(c.content) >= c.maxsize {
		c.Housekeep()
	}
	if len(c.content) >= c.maxsize {
		c.Start()
	}
	if !c.KeyExists(key) {
		c.keys[key] = key
		c.payload = c.payload + 1
		c.expiry[key] = time.Now().Add(time.Duration(c.life) * time.Second)
		c.content[key] = val
		c.used[key] = true
		//fmt.Printf("c.payload: %v\n", c.payload)
		//fmt.Printf("c.expiry[key]: %v\n", c.expiry[key])
		logs.Information("Cache", c.Stat())
	}

	//fmt.Printf("val: %v\n", val)
	//fmt.Printf("c: %v\n", c)
	return c
}

func (c Cache) OK(key string) bool {
	//spew.Dump(c)
	//logs.Information("OK", key)

	if !c.KeyExists(key) {
		//logs.Information("OK", "Not Found")
		return false
	}
	//logs.Information("OK", "Found")
	//fmt.Printf("c.expiry[key]: %v\n", c.expiry[key])
	_, rtnval := c.ExpiryProcessing(key)
	if rtnval {
		return !rtnval
	}
	//logs.Information("OK", c.Stat())
	return true
}

func (c Cache) ExpiryProcessing(key string) (Cache, bool) {
	if time.Now().After(c.expiry[key]) {
		logs.Warning("Cache Expiry: " + key + c.expiry[key].String())
		delete(c.content, key)
		delete(c.expiry, key)
		delete(c.used, key)
		delete(c.keys, key)
		c.payload = c.payload - 1
		return c, true
	}
	logs.Information("Cache", c.Stat())
	return c, false
}

func (c Cache) Get(key string) interface{} {
	//logs.Information("Get", key)
	return c.content[key]
}

func (c Cache) KeyGen(object string, id string) string {
	//logs.Information("KeyGen", object+ID_SEP+id)
	return object + ID_SEP + id
}

func (c Cache) KeyExists(key string) bool {
	return c.used[key]
}

func (c Cache) Start() Cache {
	//logs.Information("Start", "Setting up Cache")
	//cache := Cache{}
	c.content = make(map[string]interface{})
	c.expiry = make(map[string]time.Time)
	c.used = make(map[string]bool)
	c.payload = 0
	c.keys = make(map[string]string)

	cm_val := StringToInt(ApplicationProperties.Get("cache_maxsize"))
	if cm_val > 0 {
		c.maxsize = cm_val
	} else {
		c.maxsize = 1000
	}
	//logs.Information("Start Cache MaxSize", strconv.Itoa(c.maxsize))
	cl_val := StringToInt(ApplicationProperties.Get("cache_life"))
	if cl_val > 0 {
		c.life = cl_val
	} else {
		c.life = 10
	}
	//logs.Information("Start Cache Life", strconv.Itoa(c.life))
	//fmt.Printf("cache: %v\n", cache)
	logs.Information("Start Cache", c.Stat())
	return c
}

func (c Cache) Stat() string {
	// fmt.Printf("c.payload: %v\n", c.payload)
	// fmt.Printf("c.maxsize: %v\n", c.maxsize)
	// fmt.Printf("c.life: %v\n", c.life)
	// //fmt.Printf("c.content: %v\n", c.content)

	// x := len(c.content)
	// fmt.Printf("len(c.content): %v\n", x)
	// for k, v := range c.content {
	// 	fmt.Printf("k: %v, v: %v\n", k, v)
	// }
	// y := len(c.used)
	// fmt.Printf("len(c.used): %v\n", y)
	// for k, v := range c.used {
	// 	fmt.Printf("k: %v, v: %v\n", k, v)
	// }
	// z := len(c.expiry)
	// for k, t := range c.expiry {
	// 	fmt.Printf("k: %v, t: %v\n", k, t)
	// }
	//fmt.Printf("len(c.expiry): %v\n", z)
	return fmt.Sprintf("Cache Status: %v/%v %vsecs", len(c.content), c.maxsize, c.life)
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
