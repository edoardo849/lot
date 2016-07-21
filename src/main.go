package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/apex/go-apex"
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

const funcName = "masterfunc"

func init() {
	log.SetHandler(text.New(os.Stderr))
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {

	logger := log.WithFields(log.Fields{
		"function": funcName,
	})

	apex.HandleFunc(HandleEvent(logger, funcName))
}
