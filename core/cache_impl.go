package core

var cache map[string]interface{}
var cache_maxsize int
var cache_payload int

func CACHE_Add(key string, val interface{}) {
	if cache_payload >= cache_maxsize {
		CACHE_Clear()
	}
	if !CACHE_KeyExists(key) {
		cache_payload++
	}
	cache[key] = val
}

func CACHE_Get(key string) interface{} {
	return cache[key]
}

func CACHE_KeyGen(object string, id string) string {
	return object + ID_SEP + id
}

func CACHE_KeyExists(key string) bool {
	_, ok := cache[key]
	return ok
}

func CACHE_init() {
	cache = make(map[string]interface{})
	cache_maxsize = 1000
	cache_payload = 0
}

func CACHE_Clear() {
	CACHE_init()
}

func CACHE_Refresh(key string, val interface{}) {
	CACHE_Add(key, val)
}
