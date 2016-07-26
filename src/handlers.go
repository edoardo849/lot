package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"

	"github.com/apex/go-apex"
	"github.com/apex/log"
)

// Message needs comments
type Message struct {
	Value string `json:"value"`
}

// HandleEvent needs comments
func HandleEvent(event json.RawMessage, c *apex.Context) (interface{}, error) {

	var m Message
	var r []string
	logger := GetVar(c, "Logger").(*log.Entry)
	logger.Info("Started")

	fileName := "./data.json"
	logger.Info("Loading file")
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		logger.WithError(err).Error("Could not open the file")
		return nil, err
	}

	logger.Info("File loaded")

	logger.Info("Unmarshaling the event")
	if err := json.Unmarshal(event, &m); err != nil {
		logger.WithError(err).Error("Could not unmarshal the event")
		return nil, err
	}
	logger.WithField("message", m).Info("Event unmarshaled")
	logger.Info("Unmarshaling the file")
	if err := json.Unmarshal(file, &r); err != nil {
		logger.WithError(err).Error("Could not unmarshal the file")
		return nil, err
	}
	logger.Info("File unmarshaled")

	randomWord := r[rand.Intn(len(r))]

	logger.WithField("word", randomWord).Info("Random word generated")

	newValue := fmt.Sprintf("%s %s", m.Value, randomWord)

	logger.WithField("value", newValue).Info("New Value generated")

	m.Value = newValue

	return m, nil
}
