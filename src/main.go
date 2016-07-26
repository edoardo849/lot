package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/apex/go-apex"
	"github.com/apex/log"
	jsonLog "github.com/apex/log/handlers/json"
)

func init() {
	log.SetHandler(jsonLog.New(os.Stderr))
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {

	apex.HandleFunc(
		WithVars(
			WithLogger(HandleEvent),
		),
	)
}
