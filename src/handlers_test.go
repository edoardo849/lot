package main_test

import (
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/apex/go-apex"
	"github.com/apex/log"
	"github.com/apex/log/handlers/discard"
	"github.com/edoardo849/lot/functions/who"
)

func init() {
	// Discard Logs for the tests
	log.SetHandler(discard.New())
	rand.Seed(time.Now().UTC().UnixNano())
}

func TestHandler(t *testing.T) {
	handler := main.WithVars(main.WithLogger(main.HandleEvent))

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
