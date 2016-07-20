package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/apex/go-apex"
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

const funcName = "masterfunc"

type message struct {
	Value string `json:"value"`
}

func init() {
	log.SetHandler(text.New(os.Stderr))
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var m message
		var r []string

		logger := log.WithFields(log.Fields{
			"function": funcName,
		})

		logger.Info("Started")

		fileName := fmt.Sprintf("./%s.json", funcName)

		logger.WithField("file", fileName).Info("Loading file")
		file, err := ioutil.ReadFile(fileName)
		if err != nil {
			logger.WithError(err).Error("Could not open the file")
			os.Exit(1)
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
	})
}
