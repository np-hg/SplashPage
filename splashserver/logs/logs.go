package logs

import (
	"os"

	log "gopkg.in/inconshreveable/log15.v2"
)

// Logger for splash server
var Logger log.Logger

// CreateLogger inits the server logger
func CreateLogger() {
	l := log.New()
	l.SetHandler(log.StreamHandler(os.Stderr, log.JsonFormat()))

	Logger = l
}
