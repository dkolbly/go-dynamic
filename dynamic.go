package dynamic

import (
	"runtime"
	"sync/atomic"
)

type dynamicBinding struct {
	next *dynamicBinding
	key string
	value interface{}
}

var dynamicScopes map[int64]*dynamicBinding
var nextScopeID int64

func init() {
	dynamicScopes = make(map[int64]*dynamicBinding, 1000)
}

func get() *dynamicBinding {
	gid := runtime.GetGroupID()
	return dynamicScopes[gid]
}

func CallWithDynamicScope(thunk func()) {
	scopeID := atomic.AddInt64(&nextScopeID, 1)
	gid := runtime.GetGroupID()
	dynamicScopes[scopeID] = dynamicScopes[gid]
	defer delete(dynamicScopes, scopeID)
	runtime.SetGroupID(scopeID)
	thunk()
}

func Set(key string, value interface{}) {
	gid := runtime.GetGroupID()
	b := &dynamicBinding{
		next: dynamicScopes[gid],
		key: key,
		value: value,
	}
	dynamicScopes[gid] = b
}

func Get(key string) interface{} {
	gid := runtime.GetGroupID()
	b := dynamicScopes[gid]
	for b != nil {
		if b.key == key {
			return b.value
		}
	}
	return nil
}
