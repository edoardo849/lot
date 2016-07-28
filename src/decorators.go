package main

import (
	"encoding/json"
	"strings"

	"github.com/apex/go-apex"
	"github.com/apex/log"
)

// WithLogger is a decorator that adds log capabilities to handlers
func WithLogger(fn apex.HandlerFunc) apex.HandlerFunc {
	return func(e json.RawMessage, c *apex.Context) (interface{}, error) {
		logger := log.WithFields(log.Fields{
			"functionName": strings.Replace(c.FunctionName, "_", "", -1), // so that ELK analyzers will not split it
			"requestID":    c.RequestID,
		})
		SetVar(c, "Logger", logger)
		return fn(e, c)
	}
}

// WithVars is a decorator that enables the decorated handlers to access
// the variables for their own apex.Context.
func WithVars(fn apex.HandlerFunc) apex.HandlerFunc {
	return func(e json.RawMessage, c *apex.Context) (interface{}, error) {
		OpenVars(c)
		defer CloseVars(c)
		return fn(e, c)
	}
}
