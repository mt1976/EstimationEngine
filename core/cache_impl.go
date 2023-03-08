package core

var CACHE map[string]interface{}

func CACHE_Add(key string, val interface{}) {
	CACHE[key] = val
}

func CACHE_Get(key string) interface{} {
	return CACHE[key]
}

func CACHE_KeyGen(object string, id string) string {
	return object + ID_SEP + id
}

func CACHE_KeyExists(key string) bool {
	_, ok := CACHE[key]
	return ok
}
