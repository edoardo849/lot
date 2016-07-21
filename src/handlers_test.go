package main_test

import (
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/apex/go-apex"
	"github.com/apex/log"
	"github.com/apex/log/handlers/discard"
	"github.com/edoardo849/lot/src"
)

func init() {
	// Discard Logs for the tests
	log.SetHandler(discard.New())
	rand.Seed(time.Now().UTC().UnixNano())
}

func TestHandler(t *testing.T) {
	name := "masterfunc"
	logger := log.WithField("name", name)
	handler := main.HandleEvent(logger, name)

	msg := []byte(`{ "value": "What if" }`)

	ctx := apex.Context{}

	res, err := handler(msg, &ctx)

	if err != nil {
		t.Error("When given a valid message, an error was returned")
	}

	respMsg := res.(main.Message)
	if !strings.Contains(respMsg.Value, "test") {
		t.Error("The returned message should have appended the value of the .json file")
	}

}

func TestHandlerErrorFile(t *testing.T) {
	name := "somefile"
	logger := log.WithField("name", name)
	handler := main.HandleEvent(logger, name)

	msg := []byte(`{ "value": "What if" }`)

	ctx := apex.Context{}

	_, err := handler(msg, &ctx)

	if err == nil {
		t.Error("When given an invalid file name, an error shuold be expected")
	}

}
