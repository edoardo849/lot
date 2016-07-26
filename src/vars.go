package main

import (
	"sync"

	"github.com/apex/go-apex"
)

var vars map[*apex.Context]map[string]interface{}
var varsLock sync.RWMutex

// OpenVars prepares the global vars mapping the apex.Context
// as a key parameter.
func OpenVars(r *apex.Context) {
	varsLock.Lock()
	if vars == nil {
		vars = map[*apex.Context]map[string]interface{}{}

	}
	vars[r] = map[string]interface{}{}
	varsLock.Unlock()
}

// CloseVars destroys the specific key=>value association
// when the server completes the response.
func CloseVars(r *apex.Context) {
	varsLock.Lock()
	delete(vars, r)
	varsLock.Unlock()
}

// GetVar gets the requested var for the specific apex.Context and
// key
func GetVar(r *apex.Context, key string) interface{} {
	varsLock.RLock()
	value := vars[r][key]
	varsLock.RUnlock()
	return value
}

//SetVar sets a generic sets of values inside the global vars variable
func SetVar(r *apex.Context, key string, value interface{}) {
	varsLock.Lock()
	vars[r][key] = value
	varsLock.Unlock()
}
